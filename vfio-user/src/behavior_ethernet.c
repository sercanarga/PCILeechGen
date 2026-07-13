#include "device_behavior.h"

#include <errno.h>
#include <stdlib.h>
#include <string.h>

#define ETH_DESC_DD 0x01u
#define ETH_DESC_EOP 0x02u
#define ETH_MAX_PACKET 16384u
#define ETH_MIN_PACKET 60u

struct ethernet_state {
    struct device_behavior registers;
    struct behavior_host_ops host;
    uint32_t ctrl, status, eerd, mdic, icr, ims;
    uint64_t rdbal, tdbal;
    uint32_t rdlen, rdh, rdt, tdlen, tdh, tdt;
    uint64_t rx_packets, tx_packets, rx_errors;
};

static uint32_t *reg32(struct ethernet_state *s, uint64_t off)
{
    switch (off) {
    case 0x0000: return &s->ctrl;
    case 0x0008: return &s->status;
    case 0x0014: return &s->eerd;
    case 0x0020: return &s->mdic;
    case 0x00c0: return &s->icr;
    case 0x00d0: return &s->ims;
    case 0x0288: return &s->rdlen;
    case 0x2810: return &s->rdh;
    case 0x2818: return &s->rdt;
    case 0x0388: return &s->tdlen;
    case 0x3810: return &s->tdh;
    case 0x3818: return &s->tdt;
    default: return NULL;
    }
}

static int ring_ready(const struct ethernet_state *state)
{
    return state->rdlen >= 16 && state->tdlen >= 16 &&
           (state->rdlen % 16) == 0 && (state->tdlen % 16) == 0 &&
           state->tdh < state->tdlen / 16 && state->rdt < state->rdlen / 16;
}

static uint64_t reg64(struct ethernet_state *s, uint64_t off)
{
    if (off == 0x0280) return s->rdbal;
    if (off == 0x0380) return s->tdbal;
    return 0;
}

static uint16_t phy_register_value(uint16_t reg)
{
    switch (reg) {
    case 0: return 0x1140;
    case 1: return 0x796d;
    case 2: return 0x001c;
    case 3: return 0xc800;
    case 4: return 0x01e1;
    case 5: return 0xcde1;
    case 9: return 0x0300;
    case 10: return 0x3800;
    default: return 0x0000;
    }
}

static int ethernet_read_stat(const struct ethernet_state *state,
                              uint64_t offset, uint32_t *value)
{
    switch (offset) {
    case 0x4000:
        *value = (uint32_t)state->rx_packets;
        return 1;
    case 0x4004:
        *value = (uint32_t)(state->rx_packets >> 32);
        return 1;
    case 0x4008:
        *value = (uint32_t)state->tx_packets;
        return 1;
    case 0x400c:
        *value = (uint32_t)(state->tx_packets >> 32);
        return 1;
    case 0x4010:
        *value = (uint32_t)state->rx_errors;
        return 1;
    default:
        return 0;
    }
}

static void ethernet_raise_rx_irq(struct ethernet_state *state)
{
    state->icr |= 1u << 7;
    if ((state->ims & (1u << 7)) != 0 && state->host.irq != NULL)
        state->host.irq(state->host.opaque, 0);
}

static int ethernet_complete_rx(struct ethernet_state *state, uint64_t desc,
                                const uint8_t *payload, uint16_t length)
{
    uint8_t rx[16];
    uint64_t buffer;

    if (state->host.dma_read(state->host.opaque, desc, rx, sizeof(rx)) < 0)
        return -1;
    memcpy(&buffer, rx, sizeof(buffer));
    if (buffer == 0 || length == 0 || length > ETH_MAX_PACKET)
        return -1;
    if (state->host.dma_write(state->host.opaque, buffer, payload, length) < 0)
        return -1;
    memcpy(rx + 8, &length, sizeof(length));
    rx[12] = ETH_DESC_DD | ETH_DESC_EOP;
    if (state->host.dma_write(state->host.opaque, desc, rx, sizeof(rx)) < 0)
        return -1;
    state->rx_packets++;
    ethernet_raise_rx_irq(state);
    return 0;
}

static int ethernet_complete_tx(struct ethernet_state *state, uint64_t desc,
                                uint8_t *tx)
{
    tx[11] |= ETH_DESC_EOP;
    tx[12] |= ETH_DESC_DD;
    if (state->host.dma_write(state->host.opaque, desc, tx, 16) < 0)
        return -1;
    state->tx_packets++;
    return 0;
}

static size_t ethernet_fake_arp(uint8_t *packet)
{
    static const uint8_t source[6] = {0x02, 0x50, 0x43, 0x49, 0x4c, 0x45};
    static const uint8_t target_ip[4] = {192, 0, 2, 1};

    memset(packet, 0, ETH_MIN_PACKET);
    memset(packet, 0xff, 6);
    memcpy(packet + 6, source, sizeof(source));
    packet[12] = 0x08;
    packet[13] = 0x06;
    packet[14] = 0x00;
    packet[15] = 0x01;
    packet[16] = 0x08;
    packet[17] = 0x00;
    packet[18] = 0x06;
    packet[19] = 0x04;
    packet[20] = 0x00;
    packet[21] = 0x01;
    memcpy(packet + 22, source, sizeof(source));
    memcpy(packet + 38, target_ip, sizeof(target_ip));
    return ETH_MIN_PACKET;
}

static uint16_t ethernet_checksum(const uint8_t *data, size_t length)
{
    uint32_t sum = 0;

    while (length > 1) {
        sum += ((uint16_t)data[0] << 8) | data[1];
        data += 2;
        length -= 2;
    }
    if (length != 0)
        sum += (uint16_t)data[0] << 8;
    while ((sum >> 16) != 0)
        sum = (sum & 0xffffu) + (sum >> 16);
    return (uint16_t)~sum;
}

static void ethernet_put_be16(uint8_t *data, uint16_t value)
{
    data[0] = (uint8_t)(value >> 8);
    data[1] = (uint8_t)value;
}

static size_t ethernet_fake_icmp(uint8_t *packet)
{
    static const uint8_t source[6] = {0x02, 0x50, 0x43, 0x49, 0x4c, 0x45};
    static const uint8_t destination[6] = {0x02, 0x50, 0x43, 0x49, 0x4c, 0x46};
    static const uint8_t payload[] = {'P', 'C', 'L', 'G'};
    uint16_t checksum;

    memset(packet, 0, ETH_MIN_PACKET);
    memcpy(packet, destination, sizeof(destination));
    memcpy(packet + 6, source, sizeof(source));
    packet[12] = 0x08;
    packet[13] = 0x00;
    packet[14] = 0x45;
    packet[16] = 0x00;
    packet[17] = 0x20;
    packet[18] = 0x12;
    packet[19] = 0x34;
    packet[20] = 0x00;
    packet[21] = 0x00;
    packet[22] = 64;
    packet[23] = 1;
    packet[26] = 192;
    packet[27] = 0;
    packet[28] = 2;
    packet[29] = 2;
    packet[30] = 192;
    packet[31] = 0;
    packet[32] = 2;
    packet[33] = 1;
    checksum = ethernet_checksum(packet + 14, 20);
    ethernet_put_be16(packet + 24, checksum);
    packet[34] = 8;
    packet[35] = 0;
    packet[36] = 0;
    packet[37] = 0;
    packet[38] = 0;
    packet[39] = 1;
    packet[40] = 0;
    packet[41] = 1;
    memcpy(packet + 42, payload, sizeof(payload));
    checksum = ethernet_checksum(packet + 34, 12);
    ethernet_put_be16(packet + 36, checksum);
    return ETH_MIN_PACKET;
}

static int ethernet_inject_packet(struct ethernet_state *state)
{
    uint8_t packet[ETH_MIN_PACKET];
    uint64_t desc = state->rdbal + (uint64_t)state->rdh * 16;

    if ((state->rx_packets % 2) == 0) {
        if (ethernet_fake_arp(packet) == 0)
            return -1;
    } else if (ethernet_fake_icmp(packet) == 0) {
        return -1;
    }
    if (ethernet_complete_rx(state, desc, packet, sizeof(packet)) < 0)
        return -1;
    state->rdh = (state->rdh + 1) % (state->rdlen / 16);
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
    state->rx_packets = state->tx_packets = state->rx_errors = 0;
    return rc;
}


static ssize_t ethernet_read(void *opaque, unsigned bir, uint64_t offset,
                             void *data, size_t length)
{
    struct ethernet_state *state = opaque;
    uint32_t *reg;
    uint32_t value;
    if (state == NULL) return -EINVAL;
    if (data == NULL) return -EINVAL;
    if (bir == 0 && length == 4 && ethernet_read_stat(state, offset, &value)) {
        memcpy(data, &value, 4); return 4;
    }
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
    uint32_t previous_rdt;
    if (state == NULL) return -EINVAL;
    if (data == NULL) return -EINVAL;
    previous_rdt = state->rdt;
    if (bir == 0 && length == 4 && (offset == 0x0280 || offset == 0x0380)) {
        memcpy(&value, data, 4); if (offset == 0x0280) state->rdbal = value; else state->tdbal = value; return 4;
    }
    if (bir == 0 && length == 4 && (offset == 0x0284 || offset == 0x0384)) return 4;
    if (bir == 0 && length == 4 && reg32(state, offset) != NULL) {
        memcpy(&value, data, 4);
        if (offset == 0x0014) { state->eerd = (value & 1) | (2u << 4) | (0x1100u << 16); return 4; }
        if (offset == 0x0020) {
            uint16_t phy_reg = (uint16_t)((value >> 16) & 0x1f);
            uint16_t phy_data = phy_register_value(phy_reg);
            state->mdic = (value & 0x03ff0000u) | phy_data | 0x10000000u;
            return 4;
        }
        if (offset == 0x00c0) { state->icr &= ~value; return 4; }
        if (offset == 0x00d0) { state->ims = value; return 4; }
        *reg32(state, offset) = value;
        if (state->host.dma_read != NULL && state->host.dma_write != NULL && ring_ready(state)) {
            if (offset == 0x3818 && state->tdh != state->tdt) {
                uint8_t tx[16];
                uint64_t desc = state->tdbal + (uint64_t)state->tdh * 16;
                uint64_t buf;
                uint16_t len;

                if (state->host.dma_read(state->host.opaque, desc, tx, sizeof(tx)) == 0) {
                    memcpy(&buf, tx, sizeof(buf));
                    memcpy(&len, tx + 8, sizeof(len));
                    if (len > 0 && len <= ETH_MAX_PACKET && buf != 0) {
                        uint8_t *payload = malloc(len);
                        uint64_t rx_desc = state->rdbal + (uint64_t)state->rdt * 16;
                        if (payload != NULL && state->host.dma_read(state->host.opaque, buf,
                                                                     payload, len) == 0 &&
                            ethernet_complete_rx(state, rx_desc, payload, len) == 0 &&
                            ethernet_complete_tx(state, desc, tx) == 0) {
                            state->tdh = (state->tdh + 1) % (state->tdlen / 16);
                            state->rdt = (state->rdt + 1) % (state->rdlen / 16);
                        }
                        free(payload);
                    }
                }
            } else if (offset == 0x2818 && value != previous_rdt) {
                (void)ethernet_inject_packet(state);
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
