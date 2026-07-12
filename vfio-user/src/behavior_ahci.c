#include "device_behavior.h"

#include <errno.h>
#include <stdlib.h>
#include <string.h>


struct ahci_state {
    struct device_behavior registers;
};


static int ahci_bind(void *opaque, const struct behavior_host_ops *ops)
{
    struct ahci_state *state = opaque;
    return state->registers.bind_host(state->registers.state, ops);
}


static int ahci_reset(void *opaque)
{
    struct ahci_state *state = opaque;
    return state->registers.reset(state->registers.state);
}


static ssize_t ahci_read(void *opaque, unsigned bir, uint64_t offset, void *data, size_t length)
{
    struct ahci_state *state = opaque;
    return state->registers.read(state->registers.state, bir, offset, data, length);
}


static ssize_t ahci_write(void *opaque, unsigned bir, uint64_t offset,
                          const void *data, size_t length)
{
    struct ahci_state *state = opaque;
    ssize_t result = state->registers.write(state->registers.state, bir, offset, data, length);

    if (result >= 0 && bir == 0 && offset == 0x04 && length == 4) {
        uint32_t value;
        memcpy(&value, data, sizeof(value));
        if ((value & 1) != 0) {
            value &= ~1u;
            result = state->registers.write(state->registers.state, 0, 0x04, &value, 4);
        }
    }
    return result;
}


static void ahci_destroy(void *opaque)
{
    struct ahci_state *state = opaque;
    if (state != NULL) {
        state->registers.destroy(state->registers.state);
        free(state);
    }
}


int behavior_ahci_create(const struct device_model *model,
                         struct device_behavior *out, char *err, size_t err_len)
{
    struct ahci_state *state;

    if (model == NULL || out == NULL || model->class_code != 0x010601) {
        return -EINVAL;
    }
    state = calloc(1, sizeof(*state));
    if (state == NULL || behavior_static_create(model, &state->registers, err, err_len) < 0) {
        free(state);
        return -1;
    }
    *out = (struct device_behavior){
        .state = state, .bind_host = ahci_bind, .reset = ahci_reset,
        .read = ahci_read, .write = ahci_write, .destroy = ahci_destroy,
    };
    return 0;
}
