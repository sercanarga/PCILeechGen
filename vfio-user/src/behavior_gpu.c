#include "device_behavior.h"

#include <errno.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

struct gpu_state {
    struct device_behavior registers;
    uint64_t timer;
    uint32_t fence;
};

static int gpu_bind(void *opaque, const struct behavior_host_ops *ops)
{
    (void)opaque;
    return ops == NULL ? -EINVAL : 0;
}

static int gpu_reset(void *opaque)
{
    struct gpu_state *state = opaque;
    state->timer = 0;
    state->fence = 0;
    return state->registers.reset(state->registers.state);
}

static ssize_t gpu_read(void *opaque, unsigned bir, uint64_t offset,
                        void *data, size_t length)
{
    struct gpu_state *state = opaque;
    if (bir == 0 && length == 4 && offset == 0x9400) {
        uint32_t value = (uint32_t)state->timer++;
        memcpy(data, &value, sizeof(value));
        return 4;
    }
    if (bir == 0 && length == 4 && offset == 0x9410) {
        uint32_t value = (uint32_t)(state->timer >> 32);
        memcpy(data, &value, sizeof(value));
        return 4;
    }
    if (bir == 0 && length == 4 && offset == 0x1004) {
        memcpy(data, &state->fence, sizeof(state->fence));
        return 4;
    }
    return state->registers.read(state->registers.state, bir, offset, data, length);
}

static ssize_t gpu_write(void *opaque, unsigned bir, uint64_t offset,
                         const void *data, size_t length)
{
    struct gpu_state *state = opaque;
    if (bir == 0 && length == 4 && offset == 0x1000) {
        memcpy(&state->fence, data, sizeof(state->fence));
        return 4;
    }
    return state->registers.write(state->registers.state, bir, offset, data, length);
}

static void gpu_destroy(void *opaque)
{
    struct gpu_state *state = opaque;
    if (state != NULL) {
        state->registers.destroy(state->registers.state);
        free(state);
    }
}

int behavior_gpu_create(const struct device_model *model,
                        struct device_behavior *out, char *err, size_t err_len)
{
    struct gpu_state *state;
    (void)err;
    (void)err_len;
    if (model == NULL || out == NULL || (model->class_code & 0xffff00) != 0x030000) {
        return -EINVAL;
    }
    state = calloc(1, sizeof(*state));
    if (state == NULL || behavior_static_create(model, &state->registers, err, err_len) < 0) {
        free(state);
        return -1;
    }
    *out = (struct device_behavior){
        .state = state, .bind_host = gpu_bind, .reset = gpu_reset,
        .read = gpu_read, .write = gpu_write, .destroy = gpu_destroy,
    };
    if (gpu_reset(state) < 0) {
        gpu_destroy(state);
        *out = (struct device_behavior){0};
        return -1;
    }
    return 0;
}
