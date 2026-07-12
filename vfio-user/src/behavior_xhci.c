#include "device_behavior.h"

#include <errno.h>
#include <stdlib.h>
#include <string.h>


struct xhci_state {
    struct device_behavior registers;
};


static int xhci_bind(void *opaque, const struct behavior_host_ops *ops)
{
    struct xhci_state *state = opaque;
    return state->registers.bind_host(state->registers.state, ops);
}


static int xhci_reset(void *opaque)
{
    struct xhci_state *state = opaque;
    return state->registers.reset(state->registers.state);
}


static ssize_t xhci_read(void *opaque, unsigned bir, uint64_t offset, void *data, size_t length)
{
    struct xhci_state *state = opaque;
    return state->registers.read(state->registers.state, bir, offset, data, length);
}


static int xhci_set32(struct xhci_state *state, uint64_t offset, uint32_t value)
{
    return state->registers.write(state->registers.state, 0, offset, &value, 4) == 4 ? 0 : -EIO;
}


static ssize_t xhci_write(void *opaque, unsigned bir, uint64_t offset,
                          const void *data, size_t length)
{
    struct xhci_state *state = opaque;
    ssize_t result = state->registers.write(state->registers.state, bir, offset, data, length);

    if (result >= 0 && bir == 0 && offset == 0x20 && length == 4) {
        uint32_t command;
        uint32_t status = 0;

        memcpy(&command, data, sizeof(command));
        if ((command & 2) != 0) {
            command &= ~2u;
            status = 1;
        } else if ((command & 1) == 0) {
            status = 1;
        }
        if (xhci_set32(state, 0x20, command) < 0 || xhci_set32(state, 0x24, status) < 0) {
            return -EIO;
        }
    }
    return result;
}


static void xhci_destroy(void *opaque)
{
    struct xhci_state *state = opaque;
    if (state != NULL) {
        state->registers.destroy(state->registers.state);
        free(state);
    }
}


int behavior_xhci_create(const struct device_model *model,
                         struct device_behavior *out, char *err, size_t err_len)
{
    struct xhci_state *state;

    if (model == NULL || out == NULL || model->class_code != 0x0c0330) {
        return -EINVAL;
    }
    state = calloc(1, sizeof(*state));
    if (state == NULL || behavior_static_create(model, &state->registers, err, err_len) < 0) {
        free(state);
        return -1;
    }
    *out = (struct device_behavior){
        .state = state, .bind_host = xhci_bind, .reset = xhci_reset,
        .read = xhci_read, .write = xhci_write, .destroy = xhci_destroy,
    };
    return 0;
}
