#include "device_behavior.h"

#include <errno.h>
#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>


struct nvme_state {
    struct device_behavior registers;
};


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


static int nvme_reset(void *opaque)
{
    struct nvme_state *state = opaque;
    int result = state->registers.reset(state->registers.state);

    if (result < 0) {
        return result;
    }
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

    if (result < 0 || bir != 0 || offset != 0x14 || len != sizeof(uint32_t)) {
        return result;
    }
    uint32_t cc = *(const uint32_t *)buf;
    uint32_t csts = 0;

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
    if (behavior_static_create(model, &state->registers, err, err_len) < 0) {
        free(state);
        return -1;
    }
    *out = (struct device_behavior){
        .state = state,
        .reset = nvme_reset,
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
