#include "device_behavior.h"

#include <errno.h>
#include <stdlib.h>
#include <string.h>

struct hda_state {
    struct device_behavior registers;
    struct behavior_host_ops host;
    uint8_t corb_rp;
    uint8_t rirb_wp;
};

static int hda_bind(void *opaque, const struct behavior_host_ops *ops)
{
    struct hda_state *state = opaque;

    if (ops == NULL || ops->dma_read == NULL || ops->dma_write == NULL || ops->irq == NULL) {
        return -EINVAL;
    }
    state->host = *ops;
    return 0;
}

static ssize_t hda_read(void *opaque, unsigned bir, uint64_t offset,
                        void *data, size_t length)
{
    struct hda_state *state = opaque;
    return state->registers.read(state->registers.state, bir, offset, data, length);
}

static int hda_reg_write(struct hda_state *state, uint64_t offset, uint32_t value)
{
    return state->registers.write(state->registers.state, 0, offset, &value, sizeof(value)) == 4
               ? 0
               : -EIO;
}

static int hda_reset_rings(struct hda_state *state)
{
    state->corb_rp = 0;
    state->rirb_wp = 0;
    if (hda_reg_write(state, 0x48, 0) < 0 || hda_reg_write(state, 0x58, 0) < 0) {
        return -EIO;
    }
    return 0;
}

static ssize_t hda_write(void *opaque, unsigned bir, uint64_t offset,
                         const void *data, size_t length)
{
    struct hda_state *state = opaque;
    ssize_t result = state->registers.write(state->registers.state, bir, offset, data, length);

    if (result < 0 || bir != 0 || length != sizeof(uint32_t)) {
        return result;
    }
    if (offset == 0x08) {
        uint32_t value;
        memcpy(&value, data, sizeof(value));
        if ((value & 1) == 0 && hda_reset_rings(state) < 0) {
            return -EIO;
        }
    } else if (offset == 0x48) {
        uint32_t value;
        memcpy(&value, data, sizeof(value));
        state->corb_rp = (uint8_t)(value >> 16);
    } else if (offset == 0x58) {
        uint32_t value;
        memcpy(&value, data, sizeof(value));
        if ((value & (1u << 15)) != 0 && hda_reset_rings(state) < 0) {
            return -EIO;
        }
    }
    return result;
}

static int hda_service(void *opaque)
{
    struct hda_state *state = opaque;
    uint32_t corb_base_lo = 0, corb_base_hi = 0, rirb_base_lo = 0, rirb_base_hi = 0;
    uint32_t corb_wp_reg = 0, corb_ctl = 0, rirb_ctl = 0;
    uint8_t corb_wp;
    uint32_t verb = 0;
    uint64_t corb_base, rirb_base, entry_address;
    uint64_t response = 0;
    uint32_t wp_reg;

    if (state->host.dma_read == NULL || state->host.dma_write == NULL) {
        return 0;
    }
    if (hda_read(state, 0, 0x40, &corb_base_lo, 4) != 4 ||
        hda_read(state, 0, 0x44, &corb_base_hi, 4) != 4 ||
        hda_read(state, 0, 0x50, &rirb_base_lo, 4) != 4 ||
        hda_read(state, 0, 0x54, &rirb_base_hi, 4) != 4 ||
        hda_read(state, 0, 0x48, &corb_wp_reg, 4) != 4 ||
        hda_read(state, 0, 0x4c, &corb_ctl, 4) != 4 ||
        hda_read(state, 0, 0x5c, &rirb_ctl, 4) != 4) {
        return -EIO;
    }
    corb_wp = (uint8_t)corb_wp_reg;
    if ((corb_ctl & 2) == 0 || (rirb_ctl & 2) == 0 || corb_wp == state->corb_rp) {
        return 0;
    }
    corb_base = ((uint64_t)corb_base_hi << 32) | (corb_base_lo & ~0x7fU);
    rirb_base = ((uint64_t)rirb_base_hi << 32) | (rirb_base_lo & ~0x7fU);
    entry_address = corb_base + (uint64_t)((state->corb_rp + 1) & 0xff) * 4;
    if (state->host.dma_read(state->host.opaque, entry_address, &verb, sizeof(verb)) < 0) {
        return -EIO;
    }
    state->corb_rp = (uint8_t)(state->corb_rp + 1);
    state->rirb_wp = (uint8_t)(state->rirb_wp + 1);
    response = ((uint64_t)verb & 0x00ffffffU) << 32;
    entry_address = rirb_base + (uint64_t)state->rirb_wp * 8;
    if (state->host.dma_write(state->host.opaque, entry_address, &response, sizeof(response)) < 0) {
        return -EIO;
    }
    wp_reg = (uint32_t)state->rirb_wp;
    if (hda_reg_write(state, 0x48, (uint32_t)corb_wp | ((uint32_t)state->corb_rp << 16)) < 0 ||
        hda_reg_write(state, 0x58, wp_reg) < 0 ||
        state->host.irq(state->host.opaque, 0) < 0) {
        return -EIO;
    }
    return 0;
}

static int hda_reset(void *opaque)
{
    struct hda_state *state = opaque;
    if (state->registers.reset(state->registers.state) < 0) {
        return -EIO;
    }
    return hda_reset_rings(state);
}

static void hda_destroy(void *opaque)
{
    struct hda_state *state = opaque;
    if (state != NULL) {
        state->registers.destroy(state->registers.state);
        free(state);
    }
}

int behavior_hda_create(const struct device_model *model,
                        struct device_behavior *out, char *err, size_t err_len)
{
    struct hda_state *state;

    if (model == NULL || out == NULL || model->class_code != 0x040300) {
        return -EINVAL;
    }
    state = calloc(1, sizeof(*state));
    if (state == NULL || behavior_static_create(model, &state->registers, err, err_len) < 0) {
        free(state);
        return -1;
    }
    *out = (struct device_behavior){
        .state = state, .bind_host = hda_bind, .reset = hda_reset,
        .service = hda_service, .read = hda_read, .write = hda_write,
        .destroy = hda_destroy,
    };
    if (hda_reset(state) < 0) {
        hda_destroy(state);
        *out = (struct device_behavior){0};
        return -1;
    }
    return 0;
}
