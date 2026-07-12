#include <stdarg.h>
#include <stddef.h>
#include <setjmp.h>
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


int main(void)
{
    const struct CMUnitTest tests[] = {
        cmocka_unit_test_setup_teardown(exposes_generated_capability_registers, setup, teardown),
        cmocka_unit_test_setup_teardown(transitions_enable_and_shutdown_state, setup, teardown),
        cmocka_unit_test_setup_teardown(completes_identify_controller, setup, teardown),
        cmocka_unit_test_setup_teardown(completes_identify_namespace_across_prp_pages, setup, teardown),
        cmocka_unit_test_setup_teardown(wraps_admin_queue_phase, setup, teardown),
        cmocka_unit_test_setup_teardown(defers_submission_when_completion_queue_is_full, setup, teardown),
    };

    return cmocka_run_group_tests(tests, NULL, NULL);
}
