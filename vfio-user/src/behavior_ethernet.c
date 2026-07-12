#include "device_behavior.h"

#include <errno.h>
#include <stdlib.h>
#include <string.h>


struct ethernet_state {
    struct device_behavior registers;
};


static int ethernet_bind(void *opaque, const struct behavior_host_ops *ops)
{
    struct ethernet_state *state = opaque;
    return state->registers.bind_host(state->registers.state, ops);
}


static int ethernet_reset(void *opaque)
{
    struct ethernet_state *state = opaque;
    return state->registers.reset(state->registers.state);
}


static ssize_t ethernet_read(void *opaque, unsigned bir, uint64_t offset,
                             void *data, size_t length)
{
    struct ethernet_state *state = opaque;
    return state->registers.read(state->registers.state, bir, offset, data, length);
}


static ssize_t ethernet_write(void *opaque, unsigned bir, uint64_t offset,
                              const void *data, size_t length)
{
    struct ethernet_state *state = opaque;
    ssize_t result = state->registers.write(state->registers.state, bir, offset,
                                            data, length);
    uint32_t command;

    if (result < 0 || bir != 0 || offset != 0x34 || length != 4) {
        return result;
    }
    memcpy(&command, data, sizeof(command));
    if ((command & 0x10000000u) == 0) {
        return result;
    }
    command = 0x0c000000u;
    return state->registers.write(state->registers.state, 0, 0x34,
                                  &command, sizeof(command));
}


static void ethernet_destroy(void *opaque)
{
    struct ethernet_state *state = opaque;
    if (state != NULL) {
        state->registers.destroy(state->registers.state);
        free(state);
    }
}


int behavior_ethernet_create(const struct device_model *model,
                             struct device_behavior *out, char *err, size_t err_len)
{
    struct ethernet_state *state;

    if (model == NULL || out == NULL || model->class_code != 0x020000) {
        return -EINVAL;
    }
    state = calloc(1, sizeof(*state));
    if (state == NULL || behavior_static_create(model, &state->registers, err, err_len) < 0) {
        free(state);
        return -1;
    }
    *out = (struct device_behavior){
        .state = state,
        .bind_host = ethernet_bind,
        .reset = ethernet_reset,
        .read = ethernet_read,
        .write = ethernet_write,
        .destroy = ethernet_destroy,
    };
    return 0;
}
