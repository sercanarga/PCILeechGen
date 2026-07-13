#include <stdarg.h>
#include <stddef.h>
#include <setjmp.h>
#include <errno.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

#include <cmocka.h>

#include "device_behavior.h"
#include "device_model.h"


struct fixture {
    struct device_model *model;
    struct device_behavior behavior;
};

struct fake_host {
    uint8_t memory[0x10000];
    unsigned irqs;
};

struct nvme_sqe {
    uint8_t opcode;
    uint8_t flags;
    uint16_t cid;
    uint32_t nsid;
    uint64_t reserved;
    uint64_t mptr;
    uint64_t prp1;
    uint64_t prp2;
    uint32_t cdw10;
    uint32_t cdw11;
    uint32_t cdw12;
    uint32_t cdw13;
    uint32_t cdw14;
    uint32_t cdw15;
} __attribute__((packed));

struct nvme_cqe {
    uint32_t result;
    uint32_t reserved;
    uint16_t sq_head;
    uint16_t sq_id;
    uint16_t cid;
    uint16_t status;
} __attribute__((packed));


static int fake_read(void *opaque, uint64_t address, void *data, size_t length)
{
    struct fake_host *host = opaque;

    if (address > sizeof(host->memory) || length > sizeof(host->memory) - address) {
        return -1;
    }
    memcpy(data, host->memory + address, length);
    return 0;
}


static int fake_write(void *opaque, uint64_t address, const void *data, size_t length)
{
    struct fake_host *host = opaque;

    if (address > sizeof(host->memory) || length > sizeof(host->memory) - address) {
        return -1;
    }
    memcpy(host->memory + address, data, length);
    return 0;
}


static int fake_irq(void *opaque, unsigned vector)
{
    struct fake_host *host = opaque;

    assert_int_equal(vector, 0);
    host->irqs++;
    return 0;
}


static int setup(void **state)
{
    struct fixture *fixture = calloc(1, sizeof(*fixture));
    char err[256] = {0};

    assert_non_null(fixture);
    assert_int_equal(device_model_load("../tests/cocotb/out_nvme", &fixture->model, err, sizeof(err)), 0);
    assert_int_equal(behavior_nvme_create(fixture->model, &fixture->behavior, err, sizeof(err)), 0);
    *state = fixture;
    return 0;
}


static int teardown(void **state)
{
    struct fixture *fixture = *state;

    fixture->behavior.destroy(fixture->behavior.state);
    device_model_free(fixture->model);
    free(fixture);
    return 0;
}


static uint32_t read32(struct fixture *fixture, uint64_t offset)
{
    uint32_t value = 0;

    assert_int_equal(fixture->behavior.read(fixture->behavior.state, 0, offset, &value, sizeof(value)), 4);
    return value;
}

static void configures_admin_queues(struct fixture *, struct fake_host *,
                                    struct nvme_sqe **, struct nvme_cqe **,
                                    uint32_t);


static void exposes_generated_capability_registers(void **state)
{
    struct fixture *fixture = *state;

    assert_int_equal(read32(fixture, 0x00), 0x0040ff17);
    assert_int_equal(read32(fixture, 0x04), 0x00000020);
    assert_int_equal(read32(fixture, 0x08), 0x00010400);
}


static void transitions_enable_and_shutdown_state(void **state)
{
    struct fixture *fixture = *state;
    uint32_t cc = 1;

    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x14, &cc, sizeof(cc)), 4);
    assert_int_equal(read32(fixture, 0x1c) & 1, 1);

    cc = 1 | (1u << 14);
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x14, &cc, sizeof(cc)), 4);
    assert_int_equal((read32(fixture, 0x1c) >> 2) & 3, 2);

    assert_int_equal(fixture->behavior.reset(fixture->behavior.state), 0);
    assert_int_equal(read32(fixture, 0x1c) & 0x0d, 0);
}


static void rejects_completion_doorbell_outside_queue(void **state)
{
    struct fixture *fixture = *state;
    struct fake_host host = {0};
    struct nvme_sqe *sqe;
    struct nvme_cqe *cq;
    uint32_t head = 2;

    configures_admin_queues(fixture, &host, &sqe, &cq, 2);
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x1004,
                                              &head, sizeof(head)), -EINVAL);
}


static void completes_identify_controller(void **state)
{
    struct fixture *fixture = *state;
    struct fake_host host = {0};
    struct behavior_host_ops ops = {
        .opaque = &host,
        .dma_read = fake_read,
        .dma_write = fake_write,
        .irq = fake_irq,
    };
    struct nvme_sqe *sqe = (struct nvme_sqe *)&host.memory[0x1000];
    struct nvme_cqe *cqe = (struct nvme_cqe *)&host.memory[0x2000];
    uint32_t value;

    assert_int_equal(fixture->behavior.bind_host(fixture->behavior.state, &ops), 0);
    value = 3 | (3u << 16);
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x24, &value, 4), 4);
    uint64_t queue_address = 0x1000;
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x28,
                                              &queue_address, sizeof(queue_address)), 8);
    queue_address = 0x2000;
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x30,
                                              &queue_address, sizeof(queue_address)), 8);
    sqe->opcode = 0x06;
    sqe->cid = 7;
    sqe->prp1 = 0x3000;
    sqe->cdw10 = 1;
    value = 1;
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x1000, &value, 4), 4);
    assert_non_null(fixture->behavior.service);
    assert_int_equal(fixture->behavior.service(fixture->behavior.state), 0);

    assert_memory_equal(&host.memory[0x3000], "\x4d\x14", 2);
    assert_memory_equal(&host.memory[0x3018], "PCILeechGen NVMe 144D:A809", 25);
    assert_int_equal(cqe->cid, 7);
    assert_int_equal(cqe->status, 1);
    assert_int_equal(host.irqs, 1);
}


static void configures_admin_queues(struct fixture *fixture, struct fake_host *host,
                                    struct nvme_sqe **sq, struct nvme_cqe **cq,
                                    uint32_t queue_size)
{
    uint32_t aqa = (queue_size - 1) | ((queue_size - 1) << 16);
    uint64_t address;

    assert_int_equal(fixture->behavior.bind_host(fixture->behavior.state,
                                                  &(struct behavior_host_ops){
                                                      .opaque = host,
                                                      .dma_read = fake_read,
                                                      .dma_write = fake_write,
                                                      .irq = fake_irq,
                                                  }), 0);
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x24,
                                              &aqa, sizeof(aqa)), 4);
    address = 0x1000;
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x28,
                                              &address, sizeof(address)), 8);
    address = 0x2000;
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x30,
                                              &address, sizeof(address)), 8);
    *sq = (struct nvme_sqe *)&host->memory[0x1000];
    *cq = (struct nvme_cqe *)&host->memory[0x2000];
}


static void completes_identify_namespace_across_prp_pages(void **state)
{
    struct fixture *fixture = *state;
    struct fake_host host = {0};
    struct nvme_sqe *sqe;
    struct nvme_cqe *cqe;
    uint32_t tail = 1;
    uint8_t expected[16] = {0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00,
                            0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00};

    configures_admin_queues(fixture, &host, &sqe, &cqe, 4);
    sqe->opcode = 0x06;
    sqe->cid = 9;
    sqe->nsid = 1;
    sqe->prp1 = 0x3ff0;
    sqe->prp2 = 0x5000;
    sqe->cdw10 = 0;
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x1000,
                                              &tail, sizeof(tail)), 4);
    assert_int_equal(fixture->behavior.service(fixture->behavior.state), 0);
    assert_memory_equal(&host.memory[0x3ff0], expected, sizeof(expected));
    assert_int_equal(cqe->cid, 9);
    assert_int_equal(cqe->status, 1);
    assert_int_equal(host.irqs, 1);
}

static void completes_vendor_log_page(void **state)
{
    struct fixture *fixture = *state;
    struct fake_host host = {0};
    struct nvme_sqe *sqe;
    struct nvme_cqe *cq;
    uint32_t tail = 1;

    configures_admin_queues(fixture, &host, &sqe, &cq, 4);
    sqe->opcode = 0x02;
    sqe->cid = 13;
    sqe->nsid = 0;
    sqe->prp1 = 0x3ff0;
    sqe->prp2 = 0x5000;
    sqe->cdw10 = 0xc0 | (1u << 16);
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x1000,
                                              &tail, sizeof(tail)), 4);
    assert_int_equal(fixture->behavior.service(fixture->behavior.state), 0);
    assert_int_equal(host.memory[0x3ff0], 'N');
    assert_int_equal(host.memory[0x3ff1], 'V');
    assert_int_equal(host.memory[0x3ff2], 'M');
    assert_int_equal(host.memory[0x3ff3], 'D');
    assert_int_equal(cq->status, 1);
}

static void rejects_invalid_admin_commands(void **state)
{
    struct fixture *fixture = *state;
    struct fake_host host = {0};
    struct nvme_sqe *sqe;
    struct nvme_cqe *cq;
    uint32_t tail = 1;

    configures_admin_queues(fixture, &host, &sqe, &cq, 4);
    sqe->opcode = 0x02;
    sqe->cid = 15;
    sqe->prp1 = 0x3000;
    sqe->cdw10 = 0x7f | (1u << 16);
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x1000,
                                              &tail, sizeof(tail)), 4);
    assert_int_equal(fixture->behavior.service(fixture->behavior.state), 0);
    assert_int_equal(cq->status, 5);

    sqe->opcode = 0x05;
    sqe->cid = 16;
    sqe->prp1 = 0x6000;
    sqe->cdw10 = (3u << 16) | 1u;
    tail = 2;
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x1000,
                                              &tail, sizeof(tail)), 4);
    assert_int_equal(fixture->behavior.service(fixture->behavior.state), 0);
    assert_int_equal(cq[1].status, 3);
}


static void wraps_admin_queue_phase(void **state)
{
    struct fixture *fixture = *state;
    struct fake_host host = {0};
    struct nvme_sqe *sqe;
    struct nvme_cqe *cq;
    uint32_t tail;

    configures_admin_queues(fixture, &host, &sqe, &cq, 2);
    sqe[0].opcode = 0x06;
    sqe[0].cid = 1;
    sqe[0].prp1 = 0x3000;
    sqe[0].cdw10 = 1;
    tail = 1;
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x1000,
                                              &tail, sizeof(tail)), 4);
    assert_int_equal(fixture->behavior.service(fixture->behavior.state), 0);
    assert_int_equal(cq[0].cid, 1);
    assert_int_equal(cq[0].status, 1);
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x1004,
                                              &tail, sizeof(tail)), 4);

    sqe[1].opcode = 0x06;
    sqe[1].cid = 2;
    sqe[1].prp1 = 0x3000;
    sqe[1].cdw10 = 1;
    tail = 0;
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x1000,
                                              &tail, sizeof(tail)), 4);
    assert_int_equal(fixture->behavior.service(fixture->behavior.state), 0);
    assert_int_equal(cq[1].cid, 2);
    assert_int_equal(cq[1].status, 1);
    tail = 0;
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x1004,
                                              &tail, sizeof(tail)), 4);

    sqe[0].cid = 3;
    tail = 1;
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x1000,
                                              &tail, sizeof(tail)), 4);
    assert_int_equal(fixture->behavior.service(fixture->behavior.state), 0);
    assert_int_equal(cq[0].cid, 3);
    assert_int_equal(cq[0].status, 0);
    assert_int_equal(host.irqs, 3);
}


static void defers_submission_when_completion_queue_is_full(void **state)
{
    struct fixture *fixture = *state;
    struct fake_host host = {0};
    struct nvme_sqe *sqe;
    struct nvme_cqe *cq;
    uint32_t tail;

    configures_admin_queues(fixture, &host, &sqe, &cq, 2);
    sqe[0].opcode = 0x06;
    sqe[0].cid = 1;
    sqe[0].prp1 = 0x3000;
    sqe[0].cdw10 = 1;
    tail = 1;
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x1000,
                                              &tail, sizeof(tail)), 4);
    assert_int_equal(fixture->behavior.service(fixture->behavior.state), 0);

    sqe[1].opcode = 0x06;
    sqe[1].cid = 2;
    sqe[1].prp1 = 0x3000;
    sqe[1].cdw10 = 1;
    tail = 0;
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x1000,
                                              &tail, sizeof(tail)), 4);
    assert_int_equal(fixture->behavior.service(fixture->behavior.state), 0);
    assert_int_equal(cq[1].cid, 2);

    sqe[0].cid = 3;
    tail = 1;
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x1000,
                                              &tail, sizeof(tail)), 4);
    assert_int_equal(fixture->behavior.service(fixture->behavior.state), 0);
    assert_int_equal(cq[0].cid, 1);
    assert_int_equal(host.irqs, 2);

    tail = 1;
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x1004,
                                              &tail, sizeof(tail)), 4);
    assert_int_equal(fixture->behavior.service(fixture->behavior.state), 0);
    assert_int_equal(cq[0].cid, 3);
    assert_int_equal(host.irqs, 3);
}

static void processes_io_write_read_and_flush(void **state)
{
    struct fixture *fixture = *state;
    struct fake_host host = {0};
    struct nvme_sqe *admin_sq;
    struct nvme_cqe *admin_cq;
    struct nvme_sqe *io_sq;
    struct nvme_cqe *io_cq;
    uint32_t tail;
    uint32_t doorbell;
    const uint8_t payload[512] = {0x50, 0x43, 0x49, 0x4c, 0x45, 0x45, 0x43, 0x48};

    configures_admin_queues(fixture, &host, &admin_sq, &admin_cq, 4);
    admin_sq[0].opcode = 0x01;
    admin_sq[0].cid = 1;
    admin_sq[0].prp1 = 0x6000;
    admin_sq[0].cdw10 = (3u << 16) | 1u;
    admin_sq[1].opcode = 0x05;
    admin_sq[1].cid = 2;
    admin_sq[1].prp1 = 0x7000;
    admin_sq[1].cdw10 = (3u << 16) | 1u;
    tail = 2;
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x1000,
                                              &tail, sizeof(tail)), 4);
    assert_int_equal(fixture->behavior.service(fixture->behavior.state), 0);
    assert_int_equal(admin_cq[0].status, 1);
    assert_int_equal(admin_cq[1].status, 1);

    io_sq = (struct nvme_sqe *)&host.memory[0x7000];
    io_cq = (struct nvme_cqe *)&host.memory[0x6000];
    if (fixture->model->msix_vectors > 0) {
        uint32_t mask = 1u << 30;
        assert_int_equal(fixture->behavior.write(fixture->behavior.state,
                                                  fixture->model->msix_table_bir,
                                                  fixture->model->msix_table_offset + 12,
                                                  &mask, sizeof(mask)), 4);
    }
    memcpy(&host.memory[0x8000], payload, sizeof(payload));
    io_sq[0].opcode = 0x01;
    io_sq[0].cid = 10;
    io_sq[0].nsid = 1;
    io_sq[0].prp1 = 0x8000;
    io_sq[0].cdw12 = 0;
    doorbell = 1;
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x1008,
                                              &doorbell, sizeof(doorbell)), 4);
    assert_int_equal(fixture->behavior.service(fixture->behavior.state), 0);
    assert_int_equal(io_cq[0].status, 1);
    if (fixture->model->msix_vectors > 0) {
        uint32_t unmask = 0;
        assert_int_equal(fixture->behavior.write(fixture->behavior.state,
                                                  fixture->model->msix_table_bir,
                                                  fixture->model->msix_table_offset + 12,
                                                  &unmask, sizeof(unmask)), 4);
    }

    io_sq[1].opcode = 0x02;
    io_sq[1].cid = 11;
    io_sq[1].nsid = 1;
    io_sq[1].prp1 = 0x9000;
    doorbell = 2;
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x1008,
                                              &doorbell, sizeof(doorbell)), 4);
    assert_int_equal(fixture->behavior.service(fixture->behavior.state), 0);
    assert_memory_equal(&host.memory[0x9000], payload, sizeof(payload));
    assert_int_equal(io_cq[1].status, 1);

    io_sq[2].opcode = 0x00;
    io_sq[2].cid = 12;
    io_sq[2].nsid = 1;
    doorbell = 3;
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x1008,
                                              &doorbell, sizeof(doorbell)), 4);
    assert_int_equal(fixture->behavior.service(fixture->behavior.state), 0);
    assert_int_equal(io_cq[2].status, 1);

    io_sq[3].opcode = 0x02;
    io_sq[3].cid = 14;
    io_sq[3].nsid = 1;
    io_sq[3].prp1 = 0xa000;
    io_sq[3].cdw10 = 100000;
    doorbell = 0;
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x1008,
                                              &doorbell, sizeof(doorbell)), 4);
    assert_int_equal(fixture->behavior.service(fixture->behavior.state), 0);
    assert_memory_equal(&host.memory[0xa000], (uint8_t[512]){0}, 512);
    assert_int_equal(io_cq[3].status, 1);
}


int main(void)
{
    const struct CMUnitTest tests[] = {
        cmocka_unit_test_setup_teardown(exposes_generated_capability_registers, setup, teardown),
        cmocka_unit_test_setup_teardown(transitions_enable_and_shutdown_state, setup, teardown),
        cmocka_unit_test_setup_teardown(rejects_completion_doorbell_outside_queue, setup, teardown),
        cmocka_unit_test_setup_teardown(completes_identify_controller, setup, teardown),
        cmocka_unit_test_setup_teardown(completes_identify_namespace_across_prp_pages, setup, teardown),
        cmocka_unit_test_setup_teardown(completes_vendor_log_page, setup, teardown),
        cmocka_unit_test_setup_teardown(rejects_invalid_admin_commands, setup, teardown),
        cmocka_unit_test_setup_teardown(wraps_admin_queue_phase, setup, teardown),
        cmocka_unit_test_setup_teardown(defers_submission_when_completion_queue_is_full, setup, teardown),
        cmocka_unit_test_setup_teardown(processes_io_write_read_and_flush, setup, teardown),
    };

    return cmocka_run_group_tests(tests, NULL, NULL);
}
