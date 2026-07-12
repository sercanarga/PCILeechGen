#include "device_behavior.h"

#include <errno.h>
#include <stdbool.h>
#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>


struct nvme_state {
    struct device_behavior registers;
    const struct device_model *model;
    struct behavior_host_ops host;
    uint32_t aqa;
    uint64_t asq;
    uint64_t acq;
    uint16_t sq_head;
    uint16_t cq_tail;
    uint8_t cq_phase;
    uint16_t pending_tail;
    bool admin_pending;
};

struct nvme_sqe {
    uint8_t opcode;
    uint8_t flags;
    uint16_t cid;
    uint32_t nsid;
    uint64_t reserved;
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


static int fail(char *err, size_t err_len, const char *format, ...)
{
    va_list args;

    if (err != NULL && err_len > 0) {
        va_start(args, format);
        vsnprintf(err, err_len, format, args);
        va_end(args);
    }
    return -1;
}


static int set_csts(struct nvme_state *state, uint32_t value)
{
    ssize_t written = state->registers.write(
        state->registers.state, 0, 0x1c, &value, sizeof(value));

    return written == (ssize_t)sizeof(value) ? 0 : -EIO;
}


static int nvme_bind_host(void *opaque, const struct behavior_host_ops *ops)
{
    struct nvme_state *state = opaque;

    if (ops == NULL || ops->dma_read == NULL || ops->dma_write == NULL || ops->irq == NULL) {
        return -EINVAL;
    }
    state->host = *ops;
    return 0;
}


static void put16(uint8_t *data, size_t offset, uint16_t value)
{
    data[offset] = (uint8_t)value;
    data[offset + 1] = (uint8_t)(value >> 8);
}


static void put32(uint8_t *data, size_t offset, uint32_t value)
{
    put16(data, offset, (uint16_t)value);
    put16(data, offset + 2, (uint16_t)(value >> 16));
}


static void put64(uint8_t *data, size_t offset, uint64_t value)
{
    put32(data, offset, (uint32_t)value);
    put32(data, offset + 4, (uint32_t)(value >> 32));
}


static void identify_controller(const struct nvme_state *state, uint8_t data[4096])
{
    memset(data, 0, 4096);
    put16(data, 0x00, state->model->vendor_id);
    put16(data, 0x02, state->model->subsystem_vendor_id);
    memcpy(data + 0x04, "PCILEECHGEN000000001", 20);
    memcpy(data + 0x18, "PCILeechGen NVMe Controller", 27);
    memcpy(data + 0x40, "1.0     ", 8);
    data[0x4d] = 3;
    put32(data, 0x50, 0x00010400);
    data[0x6f] = 1;
    data[0x103] = 3;
    data[0x200] = 0x66;
    data[0x201] = 0x44;
    put32(data, 0x204, 1);
    put16(data, 0x208, 0x000c);
}


static void identify_namespace(uint8_t data[4096])
{
    memset(data, 0, 4096);
    put64(data, 0x00, 0x100000);
    put64(data, 0x08, 0x100000);
    put64(data, 0x10, 0x100000);
    data[0x19] = 0;
    data[0x1a] = 0;
    data[0x80 + 2] = 9;
}


static int write_prps(struct nvme_state *state, const struct nvme_sqe *sqe,
                      const uint8_t *data, size_t length)
{
    size_t first = 4096 - (size_t)(sqe->prp1 & 0xfff);

    if (first > length) {
        first = length;
    }
    if (state->host.dma_write(state->host.opaque, sqe->prp1, data, first) < 0) {
        return -1;
    }
    if (first < length &&
        state->host.dma_write(state->host.opaque, sqe->prp2, data + first, length - first) < 0) {
        return -1;
    }
    return 0;
}


static int process_admin(struct nvme_state *state, uint16_t tail)
{
    uint16_t sq_size = (uint16_t)((state->aqa & 0xfff) + 1);
    uint16_t cq_size = (uint16_t)(((state->aqa >> 16) & 0xfff) + 1);

    if (state->host.dma_read == NULL || state->asq == 0 || state->acq == 0 ||
        sq_size < 2 || cq_size < 2 || tail >= sq_size) {
        return -EINVAL;
    }
    while (state->sq_head != tail) {
        struct nvme_sqe sqe;
        struct nvme_cqe cqe = {0};
        uint16_t status = 0;
        uint32_t result = 0;
        uint8_t data[4096];

        if (state->host.dma_read(state->host.opaque,
                                 state->asq + (uint64_t)state->sq_head * sizeof(sqe),
                                 &sqe, sizeof(sqe)) < 0) {
            return -EIO;
        }
        if (sqe.opcode == 0x06) {
            unsigned cns = sqe.cdw10 & 0xff;

            if (cns == 1) {
                identify_controller(state, data);
            } else if (cns == 0 && sqe.nsid == 1) {
                identify_namespace(data);
            } else if (cns == 2) {
                memset(data, 0, sizeof(data));
                put32(data, 0, 1);
            } else {
                status = 0xb;
                memset(data, 0, sizeof(data));
            }
            if (status == 0 && write_prps(state, &sqe, data, sizeof(data)) < 0) {
                return -EIO;
            }
        } else if (sqe.opcode == 0x09 || sqe.opcode == 0x0a) {
            if ((sqe.cdw10 & 0xff) != 0x07) {
                status = 2;
            }
        }
        state->sq_head = (uint16_t)((state->sq_head + 1) % sq_size);
        cqe.result = result;
        cqe.sq_head = state->sq_head;
        cqe.cid = sqe.cid;
        cqe.status = (uint16_t)((status << 1) | state->cq_phase);
        if (state->host.dma_write(state->host.opaque,
                                  state->acq + (uint64_t)state->cq_tail * sizeof(cqe),
                                  &cqe, sizeof(cqe)) < 0) {
            return -EIO;
        }
        state->cq_tail = (uint16_t)((state->cq_tail + 1) % cq_size);
        if (state->cq_tail == 0) {
            state->cq_phase ^= 1;
        }
        if (state->host.irq(state->host.opaque, 0) < 0) {
            return -EIO;
        }
    }
    return 0;
}


static int nvme_service(void *opaque)
{
    struct nvme_state *state = opaque;

    if (!state->admin_pending) {
        return 0;
    }
    state->admin_pending = false;
    return process_admin(state, state->pending_tail);
}


static int nvme_reset(void *opaque)
{
    struct nvme_state *state = opaque;
    int result = state->registers.reset(state->registers.state);

    if (result < 0) {
        return result;
    }
    state->aqa = 0;
    state->asq = 0;
    state->acq = 0;
    state->sq_head = 0;
    state->cq_tail = 0;
    state->cq_phase = 1;
    state->pending_tail = 0;
    state->admin_pending = false;
    return set_csts(state, 0);
}


static ssize_t nvme_read(void *opaque, unsigned bir, uint64_t offset,
                         void *buf, size_t len)
{
    struct nvme_state *state = opaque;

    return state->registers.read(state->registers.state, bir, offset, buf, len);
}


static ssize_t nvme_write(void *opaque, unsigned bir, uint64_t offset,
                          const void *buf, size_t len)
{
    struct nvme_state *state = opaque;
    ssize_t result = state->registers.write(state->registers.state, bir, offset, buf, len);

    if (result >= 0 && bir == 0 && len == sizeof(uint64_t) &&
        (offset == 0x28 || offset == 0x30)) {
        uint64_t value;

        memcpy(&value, buf, sizeof(value));
        if (offset == 0x28) {
            state->asq = value;
        } else {
            state->acq = value;
        }
        return result;
    }
    if (result < 0 || bir != 0 || offset != 0x14 || len != sizeof(uint32_t)) {
        if (result < 0 || bir != 0 || len != sizeof(uint32_t)) {
            return result;
        }
        uint32_t value;
        memcpy(&value, buf, sizeof(value));
        switch (offset) {
        case 0x24: state->aqa = value; break;
        case 0x28: state->asq = (state->asq & 0xffffffff00000000ULL) | value; break;
        case 0x2c: state->asq = (state->asq & 0xffffffffULL) | ((uint64_t)value << 32); break;
        case 0x30: state->acq = (state->acq & 0xffffffff00000000ULL) | value; break;
        case 0x34: state->acq = (state->acq & 0xffffffffULL) | ((uint64_t)value << 32); break;
        case 0x1000:
            state->pending_tail = (uint16_t)value;
            state->admin_pending = true;
            return result;
        default: break;
        }
        return result;
    }
    uint32_t cc;
    uint32_t csts = 0;

    memcpy(&cc, buf, sizeof(cc));

    if ((cc & 1) != 0) {
        csts |= 1;
    }
    if (((cc >> 14) & 3) != 0) {
        csts |= 2u << 2;
    }
    if (set_csts(state, csts) < 0) {
        return -EIO;
    }
    return result;
}


static void nvme_destroy(void *opaque)
{
    struct nvme_state *state = opaque;

    if (state == NULL) {
        return;
    }
    state->registers.destroy(state->registers.state);
    free(state);
}


int behavior_nvme_create(const struct device_model *model,
                         struct device_behavior *out, char *err, size_t err_len)
{
    struct nvme_state *state;

    if (model == NULL || out == NULL || model->class_code != 0x010802) {
        return fail(err, err_len, "NVMe behavior requires PCI class 0x010802");
    }
    state = calloc(1, sizeof(*state));
    if (state == NULL) {
        return fail(err, err_len, "allocate NVMe behavior");
    }
    state->model = model;
    if (behavior_static_create(model, &state->registers, err, err_len) < 0) {
        free(state);
        return -1;
    }
    *out = (struct device_behavior){
        .state = state,
        .bind_host = nvme_bind_host,
        .reset = nvme_reset,
        .service = nvme_service,
        .read = nvme_read,
        .write = nvme_write,
        .destroy = nvme_destroy,
    };
    if (nvme_reset(state) < 0) {
        nvme_destroy(state);
        *out = (struct device_behavior){0};
        return fail(err, err_len, "reset NVMe behavior");
    }
    return 0;
}
