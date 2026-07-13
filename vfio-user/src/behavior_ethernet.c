#include "device_behavior.h"

#include <errno.h>
#include <stdlib.h>
#include <string.h>

#define ETH_DESC_DD 0x01u
#define ETH_DESC_EOP 0x02u
#define ETH_MAX_PACKET 2048u
#define ETH_MIN_PACKET 60u
#define ETH_ICR_TXDW 0x00000001u
#define ETH_ICR_RXT0 0x00000080u
#define ETH_REG_ICR 0x00c0u
#define ETH_REG_ICS 0x00c8u
#define ETH_REG_IMS 0x00d0u
#define ETH_REG_IMC 0x00d8u
#define ETH_REG_RDBAL 0x2800u
#define ETH_REG_RDBAH 0x2804u
#define ETH_REG_RDLEN 0x2808u
#define ETH_REG_RDH 0x2810u
#define ETH_REG_RDT 0x2818u
#define ETH_REG_TDBAL 0x3800u
#define ETH_REG_TDBAH 0x3804u
#define ETH_REG_TDLEN 0x3808u
#define ETH_REG_TDH 0x3810u
#define ETH_REG_TDT 0x3818u

struct ethernet_state {
    struct device_behavior registers;
    struct behavior_host_ops host;
    uint32_t ctrl, status, eerd, mdic, icr, ims;
    uint64_t rdbal, tdbal;
    uint32_t rdlen, rdh, rdt, tdlen, tdh, tdt;
    uint64_t rx_packets, tx_packets;
};

static uint32_t *reg32(struct ethernet_state *s, uint64_t off)
{
    switch (off) {
    case 0x0000: return &s->ctrl;
    case 0x0008: return &s->status;
    case 0x0014: return &s->eerd;
    case 0x0020: return &s->mdic;
    case ETH_REG_ICR: return &s->icr;
    case ETH_REG_IMS: return &s->ims;
    case ETH_REG_RDLEN: return &s->rdlen;
    case ETH_REG_RDH: return &s->rdh;
    case ETH_REG_RDT: return &s->rdt;
    case ETH_REG_TDLEN: return &s->tdlen;
    case ETH_REG_TDH: return &s->tdh;
    case ETH_REG_TDT: return &s->tdt;
    default: return NULL;
    }
}

static int ring_ready(uint64_t base, uint32_t length,
                      uint32_t head, uint32_t tail)
{
    uint32_t entries;

    if (base == 0 || (base & 0x7fu) != 0 ||
        length < 128 || (length & 0x7fu) != 0)
        return 0;
    entries = length / 16;
    return head < entries && tail < entries;
}

static int rx_ring_ready(const struct ethernet_state *state)
{
    return ring_ready(state->rdbal, state->rdlen, state->rdh, state->rdt);
}

static int tx_ring_ready(const struct ethernet_state *state)
{
    return ring_ready(state->tdbal, state->tdlen, state->tdh, state->tdt);
}

static uint64_t reg64(struct ethernet_state *s, uint64_t off)
{
    if (off == ETH_REG_RDBAL || off == ETH_REG_RDBAH) return s->rdbal;
    if (off == ETH_REG_TDBAL || off == ETH_REG_TDBAH) return s->tdbal;
    return 0;
}

static void ethernet_raise_irq(struct ethernet_state *state, uint32_t causes)
{
    state->icr |= causes;
    if ((state->ims & causes) != 0 && state->host.irq != NULL)
        state->host.irq(state->host.opaque, 0);
}

static int ethernet_complete_rx(struct ethernet_state *state, uint64_t desc,
                                const uint8_t *payload, uint16_t length)
{
    uint8_t rx[8];
    uint8_t status = ETH_DESC_DD | ETH_DESC_EOP;
    uint64_t buffer;

    if (state->host.dma_read(state->host.opaque, desc, rx, sizeof(rx)) < 0)
        return -1;
    memcpy(&buffer, rx, sizeof(buffer));
    if (buffer == 0 || length == 0 || length > ETH_MAX_PACKET)
        return -1;
    if (state->host.dma_write(state->host.opaque, buffer, payload, length) < 0)
        return -1;
    if (state->host.dma_write(state->host.opaque, desc + 8, &length,
                              sizeof(length)) < 0 ||
        state->host.dma_write(state->host.opaque, desc + 12, &status,
                              sizeof(status)) < 0)
        return -1;
    state->rx_packets++;
    return 0;
}

static int ethernet_complete_tx(struct ethernet_state *state, uint64_t desc)
{
    uint8_t status = ETH_DESC_DD;

    if (state->host.dma_write(state->host.opaque, desc + 12, &status,
                              sizeof(status)) < 0)
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
    ethernet_raise_irq(state, ETH_ICR_RXT0);
    return 0;
}


static int ethernet_process_tx(struct ethernet_state *state)
{
    uint8_t tx[16];
    uint64_t tx_desc;
    uint64_t rx_desc;
    uint64_t buffer;
    uint16_t length;
    uint8_t *payload;

    if (state->host.dma_read == NULL || state->host.dma_write == NULL ||
        !tx_ring_ready(state) || !rx_ring_ready(state) ||
        state->tdh == state->tdt || state->rdh == state->rdt)
        return 0;

    tx_desc = state->tdbal + (uint64_t)state->tdh * 16;
    rx_desc = state->rdbal + (uint64_t)state->rdh * 16;
    if (state->host.dma_read(state->host.opaque, tx_desc, tx, sizeof(tx)) < 0)
        return -1;
    memcpy(&buffer, tx, sizeof(buffer));
    memcpy(&length, tx + 8, sizeof(length));
    if (buffer == 0 || length == 0 || length > ETH_MAX_PACKET)
        return -1;

    payload = malloc(length);
    if (payload == NULL)
        return -1;
    if (state->host.dma_read(state->host.opaque, buffer, payload, length) < 0 ||
        ethernet_complete_rx(state, rx_desc, payload, length) < 0 ||
        ethernet_complete_tx(state, tx_desc) < 0) {
        free(payload);
        return -1;
    }
    free(payload);

    state->tdh = (state->tdh + 1) % (state->tdlen / 16);
    state->rdh = (state->rdh + 1) % (state->rdlen / 16);
    ethernet_raise_irq(state, ETH_ICR_TXDW | ETH_ICR_RXT0);
    return 1;
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
    state->rx_packets = state->tx_packets = 0;
    return rc;
}


static ssize_t ethernet_read(void *opaque, unsigned bir, uint64_t offset,
                             void *data, size_t length)
{
    struct ethernet_state *state = opaque;
    uint32_t *reg;
    if (state == NULL) return -EINVAL;
    if (data == NULL) return -EINVAL;
    if (bir == 0 && length == 4 && offset == ETH_REG_ICR) {
        uint32_t value = state->icr;
        state->icr = 0;
        memcpy(data, &value, sizeof(value));
        return 4;
    }
    if (bir == 0 && length == 4 && (reg = reg32(state, offset)) != NULL) {
        memcpy(data, reg, 4);
        return 4;
    }
    if (bir == 0 && length == 4 &&
        (offset == ETH_REG_RDBAL || offset == ETH_REG_RDBAH ||
         offset == ETH_REG_TDBAL || offset == ETH_REG_TDBAH)) {
        uint64_t value64 = reg64(state, offset);
        uint32_t value = (offset == ETH_REG_RDBAH || offset == ETH_REG_TDBAH)
                             ? (uint32_t)(value64 >> 32)
                             : (uint32_t)value64;
        memcpy(data, &value, sizeof(value));
        return 4;
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
    if (bir == 0 && length == 4 &&
        (offset == ETH_REG_RDBAL || offset == ETH_REG_TDBAL)) {
        uint64_t *base = offset == ETH_REG_RDBAL ? &state->rdbal : &state->tdbal;
        memcpy(&value, data, sizeof(value));
        *base = (*base & 0xffffffff00000000ULL) | (value & 0xffffff80u);
        return 4;
    }
    if (bir == 0 && length == 4 &&
        (offset == ETH_REG_RDBAH || offset == ETH_REG_TDBAH)) {
        uint64_t *base = offset == ETH_REG_RDBAH ? &state->rdbal : &state->tdbal;
        memcpy(&value, data, sizeof(value));
        *base = (*base & 0xffffffffULL) | ((uint64_t)value << 32);
        return 4;
    }
    if (bir == 0 && length == 4 && offset == ETH_REG_ICS) {
        memcpy(&value, data, sizeof(value));
        ethernet_raise_irq(state, value);
        return 4;
    }
    if (bir == 0 && length == 4 && offset == ETH_REG_IMC) {
        memcpy(&value, data, sizeof(value));
        state->ims &= ~value;
        return 4;
    }
    if (bir == 0 && length == 4 && reg32(state, offset) != NULL) {
        memcpy(&value, data, 4);
        if (offset == 0x0014) { state->eerd = (value & 1) | (2u << 4) | (0x1100u << 16); return 4; }
        if (offset == 0x0020) {
            uint16_t phy_reg = (uint16_t)((value >> 16) & 0x1f);
            uint16_t phy_data = phy_reg == 1 ? 0x796d : (phy_reg == 0 ? 0x1140 : 0);
            state->mdic = (value & 0x03ff0000u) | phy_data | 0x10000000u;
            return 4;
        }
        if (offset == ETH_REG_ICR) return 4;
        if (offset == ETH_REG_IMS) {
            state->ims |= value;
            if ((state->icr & value) != 0 && state->host.irq != NULL)
                state->host.irq(state->host.opaque, 0);
            return 4;
        }
        if (offset == ETH_REG_RDLEN || offset == ETH_REG_TDLEN)
            value &= 0x000fff80u;
        if (offset == ETH_REG_RDH || offset == ETH_REG_RDT ||
            offset == ETH_REG_TDH || offset == ETH_REG_TDT)
            value &= 0x0000ffffu;
        *reg32(state, offset) = value;
        if (state->host.dma_read != NULL && state->host.dma_write != NULL) {
            uint32_t processed = 0;
            uint32_t limit = tx_ring_ready(state) ? state->tdlen / 16 : 0;
            int tx_result = 0;
            if (offset == ETH_REG_TDT || offset == ETH_REG_RDT) {
                while (processed < limit &&
                       (tx_result = ethernet_process_tx(state)) > 0)
                    processed++;
            }
            if (offset == ETH_REG_RDT && value != previous_rdt &&
                processed == 0 && tx_result == 0 &&
                rx_ring_ready(state) && state->rdh != state->rdt)
                (void)ethernet_inject_packet(state);
        }
        return 4;
    }
    return state->registers.write(state->registers.state, bir, offset,
                                  data, length);
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

    if (model == NULL || out == NULL || model->class_code != 0x020000 ||
        model->vendor_id != 0x8086 || model->device_id != 0x15b7) {
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
