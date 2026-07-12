#include "device_behavior.h"

#include <errno.h>
#include <stdlib.h>
#include <string.h>


struct ethernet_state {
    struct device_behavior registers;
    struct behavior_host_ops host;
    uint32_t ctrl, status, eerd, mdic, icr, ims;
    uint64_t rdbal, tdbal;
    uint32_t rdlen, rdh, rdt, tdlen, tdh, tdt;
};

static uint32_t *reg32(struct ethernet_state *s, uint64_t off)
{
    switch (off) {
    case 0x0000: return &s->ctrl; case 0x0008: return &s->status;
    case 0x0014: return &s->eerd; case 0x0020: return &s->mdic; case 0x00c0: return &s->icr;
    case 0x00d0: return &s->ims; case 0x0288: return &s->rdlen;
    case 0x2810: return &s->rdh; case 0x2818: return &s->rdt;
    case 0x0388: return &s->tdlen; case 0x3810: return &s->tdh;
    case 0x3818: return &s->tdt; default: return NULL;
    }
}

static uint64_t reg64(struct ethernet_state *s, uint64_t off)
{
    if (off == 0x0280) return s->rdbal;
    if (off == 0x0380) return s->tdbal;
    return 0;
}


static int ethernet_bind(void *opaque, const struct behavior_host_ops *ops)
{
    struct ethernet_state *state = opaque;
    if (ops != NULL) state->host = *ops;
    return state->registers.bind_host(state->registers.state, ops);
}


static int ethernet_reset(void *opaque)
{
    struct ethernet_state *state = opaque;
    int rc = state->registers.reset(state->registers.state);
    state->ctrl = 0; state->status = 0x80080783; state->eerd = 0; state->mdic = 0x08000000;
    state->icr = 0; state->ims = 0; state->rdbal = state->tdbal = 0;
    state->rdlen = state->rdh = state->rdt = 0;
    state->tdlen = state->tdh = state->tdt = 0;
    return rc;
}


static ssize_t ethernet_read(void *opaque, unsigned bir, uint64_t offset,
                             void *data, size_t length)
{
    struct ethernet_state *state = opaque;
    uint32_t *reg;
    if (bir == 0 && length == 4 && (reg = reg32(state, offset)) != NULL) {
        memcpy(data, reg, 4); return 4;
    }
    if (bir == 0 && length == 4 && (offset == 0x0280 || offset == 0x0380)) {
        uint32_t value = (uint32_t)reg64(state, offset); memcpy(data, &value, 4); return 4;
    }
    return state->registers.read(state->registers.state, bir, offset, data, length);
}


static ssize_t ethernet_write(void *opaque, unsigned bir, uint64_t offset,
                              const void *data, size_t length)
{
    struct ethernet_state *state = opaque;
    uint32_t value;
    if (bir == 0 && length == 4 && (offset == 0x0280 || offset == 0x0380)) {
        memcpy(&value, data, 4); if (offset == 0x0280) state->rdbal = value; else state->tdbal = value; return 4;
    }
    if (bir == 0 && length == 4 && (offset == 0x0284 || offset == 0x0384)) return 4;
    if (bir == 0 && length == 4 && reg32(state, offset) != NULL) {
        memcpy(&value, data, 4);
        if (offset == 0x0014) { state->eerd = (value & 1) | (2u << 4) | (0x1100u << 16); return 4; }
        if (offset == 0x0020) {
            uint16_t phy_reg = (uint16_t)((value >> 16) & 0x1f);
            uint16_t phy_data = phy_reg == 1 ? 0x796d : (phy_reg == 0 ? 0x1140 : 0);
            state->mdic = (value & 0x03ff0000u) | phy_data | 0x10000000u;
            return 4;
        }
        if (offset == 0x00c0) { state->icr &= ~value; return 4; }
        if (offset == 0x00d0) { state->ims = value; return 4; }
        *reg32(state, offset) = value;
        if (offset == 0x3818 && state->host.dma_read != NULL && state->host.dma_write != NULL &&
            state->tdlen >= 16 && (state->tdlen % 16) == 0 &&
            state->rdlen >= 16 && (state->rdlen % 16) == 0 &&
            state->tdh < state->tdlen / 16 && state->rdt < state->rdlen / 16) {
            uint8_t tx[16]; uint64_t desc = state->tdbal + (uint64_t)state->tdh * 16;
            if (state->host.dma_read(state->host.opaque, desc, tx, sizeof(tx)) == 0) {
                uint64_t buf; uint16_t len; memcpy(&buf, tx, 8); memcpy(&len, tx + 8, 2);
                uint8_t rx[16]; uint64_t rx_desc = state->rdbal + (uint64_t)state->rdt * 16;
                uint64_t rxbuf = 0;
                if (state->host.dma_read(state->host.opaque, rx_desc, rx, sizeof(rx)) == 0) {
                    memcpy(&rxbuf, rx, 8);
                }
                if (len > 0 && len <= 16384 && rxbuf != 0) {
                    uint8_t *payload = malloc(len);
                    if (payload != NULL && state->host.dma_read(state->host.opaque, buf, payload, len) == 0)
                        state->host.dma_write(state->host.opaque, rxbuf, payload, len);
                    free(payload);
                }
                if (state->host.dma_write(state->host.opaque, rx_desc, tx, sizeof(tx)) == 0) {
                    state->rdt = (state->rdt + 1) % (state->rdlen / 16);
                    state->icr |= 1u << 7;
                    if (state->ims & (1u << 7) && state->host.irq) state->host.irq(state->host.opaque, 0);
                }
                (void)buf; (void)len;
            }
        }
        return 4;
    }
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
