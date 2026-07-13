#include <stdarg.h>
#include <stddef.h>
#include <setjmp.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

#include <cmocka.h>

#include "device_behavior.h"
#include "device_model.h"

struct ethernet_host {
    uint8_t memory[0x10000];
    unsigned irqs;
};

static int ethernet_dma_read(void *opaque, uint64_t address, void *data, size_t length)
{
    struct ethernet_host *host = opaque;
    if (address > sizeof(host->memory) || length > sizeof(host->memory) - address) return -1;
    memcpy(data, host->memory + address, length);
    return 0;
}

static int ethernet_dma_write(void *opaque, uint64_t address, const void *data, size_t length)
{
    struct ethernet_host *host = opaque;
    if (address > sizeof(host->memory) || length > sizeof(host->memory) - address) return -1;
    memcpy(host->memory + address, data, length);
    return 0;
}

static int ethernet_irq(void *opaque, unsigned vector)
{
    struct ethernet_host *host = opaque;
    assert_int_equal(vector, 0);
    host->irqs++;
    return 0;
}

static uint16_t ethernet_test_checksum(const uint8_t *data, size_t length)
{
    uint32_t sum = 0;

    while (length > 1) {
        sum += ((uint16_t)data[0] << 8) | data[1];
        data += 2;
        length -= 2;
    }
    while ((sum >> 16) != 0)
        sum = (sum & 0xffffu) + (sum >> 16);
    return (uint16_t)sum;
}


static uint32_t read32(struct device_behavior *behavior, uint64_t offset)
{
    uint32_t value = 0;

    assert_int_equal(behavior->read(behavior->state, 0, offset, &value, 4), 4);
    return value;
}


static void ahci_reset_self_clears(void **state)
{
    struct device_model *model = NULL;
    struct device_behavior behavior = {0};
    char err[256] = {0};
    uint32_t ghc;

    (void)state;
    assert_int_equal(device_model_load("../tests/cocotb/out_sata", &model, err, sizeof(err)), 0);
    assert_int_equal(behavior_ahci_create(model, &behavior, err, sizeof(err)), 0);
    uint32_t value = 0;
    assert_int_equal(behavior.read(behavior.state, 5, 0x0c, &value, 4), 4);
    assert_int_equal(value, 1);
    ghc = 0x80000001;
    assert_int_equal(behavior.write(behavior.state, 5, 0x04, &ghc, 4), 4);
    assert_int_equal(behavior.read(behavior.state, 5, 0x04, &value, 4), 4);
    assert_int_equal(value & 1, 0);
    behavior.destroy(behavior.state);
    device_model_free(model);
}


static void xhci_reset_and_run_state(void **state)
{
    struct device_model *model = NULL;
    struct device_behavior behavior = {0};
    char err[256] = {0};
    uint32_t command;

    (void)state;
    assert_int_equal(device_model_load("../tests/cocotb/out_xhci", &model, err, sizeof(err)), 0);
    assert_int_equal(behavior_xhci_create(model, &behavior, err, sizeof(err)), 0);
    assert_int_equal(read32(&behavior, 0x00), 0x01100020);
    command = 2;
    assert_int_equal(behavior.write(behavior.state, 0, 0x20, &command, 4), 4);
    assert_int_equal(read32(&behavior, 0x20) & 2, 0);
    assert_int_equal(read32(&behavior, 0x24) & 0x801, 1);
    command = 1;
    assert_int_equal(behavior.write(behavior.state, 0, 0x20, &command, 4), 4);
    assert_int_equal(read32(&behavior, 0x24) & 1, 0);
    assert_int_equal(read32(&behavior, 0x10), 0x00100001);
    assert_int_equal(read32(&behavior, 0x40), 2);
    behavior.destroy(behavior.state);
    device_model_free(model);
}


static void ethernet_link_status_and_reset(void **state)
{
    struct device_model *model = NULL;
    struct device_behavior behavior = {0};
    char err[256] = {0};
    uint32_t value;
    uint32_t command = 0x10000000;

    (void)state;
    assert_int_equal(device_model_load("../tests/cocotb/out_ethernet", &model,
                                      err, sizeof(err)), 0);
    assert_int_equal(behavior_ethernet_create(model, &behavior, err, sizeof(err)), 0);
    assert_int_equal(behavior.reset(behavior.state), 0);
    assert_int_equal(behavior.read(behavior.state, 0, 0x08, &value, 4), 4);
    assert_int_equal(value, 0x80080783);
    command = (1u << 26) | (1u << 21);
    assert_int_equal(behavior.write(behavior.state, 0, 0x20, &command, 4), 4);
    assert_int_equal(behavior.read(behavior.state, 0, 0x20, &value, 4), 4);
    assert_true((value & 0x10000000u) != 0);
    command = 1u << 7;
    assert_int_equal(behavior.write(behavior.state, 0, 0x00d0, &command, 4), 4);
    assert_int_equal(behavior.write(behavior.state, 0, 0x00c8, &command, 4), 4);
    assert_int_equal(behavior.reset(behavior.state), 0);
    assert_int_equal(behavior.read(behavior.state, 0, 0x00c0, &value, 4), 4);
    assert_int_equal(value, 0);
    assert_int_equal(behavior.read(behavior.state, 0, 0x00d0, &value, 4), 4);
    assert_int_equal(value, 0);
    behavior.destroy(behavior.state);
    device_model_free(model);
}

static void ethernet_descriptor_loopback(void **state)
{
    struct device_model *model = NULL;
    struct device_behavior behavior = {0};
    struct ethernet_host host = {0};
    struct behavior_host_ops ops = {
        .opaque = &host, .dma_read = ethernet_dma_read,
        .dma_write = ethernet_dma_write, .irq = ethernet_irq,
    };
    char err[256] = {0};
    uint32_t base;
    uint32_t value;
    uint8_t tx_desc[16] = {0};
    uint8_t rx_desc[16] = {0};
    uint16_t received_len;
    uint64_t tx_buffer = 0x3000;
    uint64_t rx_buffer = 0x4000;
    uint16_t packet_len = 4;

    (void)state;
    assert_int_equal(device_model_load("../tests/cocotb/out_ethernet", &model, err, sizeof(err)), 0);
    assert_int_equal(behavior_ethernet_create(model, &behavior, err, sizeof(err)), 0);
    assert_int_equal(behavior.bind_host(behavior.state, &ops), 0);
    base = 0x1000;
    assert_int_equal(behavior.write(behavior.state, 0, 0x2800, &base, 4), 4);
    base = 0x2000;
    assert_int_equal(behavior.write(behavior.state, 0, 0x3800, &base, 4), 4);
    value = 128;
    assert_int_equal(behavior.write(behavior.state, 0, 0x2808, &value, 4), 4);
    assert_int_equal(behavior.write(behavior.state, 0, 0x3808, &value, 4), 4);
    memcpy(tx_desc, &tx_buffer, sizeof(tx_buffer));
    memcpy(tx_desc + 8, &packet_len, sizeof(packet_len));
    tx_desc[11] = 0x02;
    tx_desc[13] = 0xa5;
    memcpy(rx_desc, &rx_buffer, sizeof(rx_buffer));
    rx_desc[10] = 0x5a;
    rx_desc[13] = 0xa5;
    memcpy(host.memory + 0x2000, tx_desc, sizeof(tx_desc));
    memcpy(host.memory + 0x1000, rx_desc, sizeof(rx_desc));
    memcpy(host.memory + 0x3000, "PING", 4);
    value = 1u << 7;
    assert_int_equal(behavior.write(behavior.state, 0, 0x00d0, &value, 4), 4);
    value = 1;
    assert_int_equal(behavior.write(behavior.state, 0, 0x3818, &value, 4), 4);
    assert_int_equal(behavior.write(behavior.state, 0, 0x2818, &value, 4), 4);
    assert_memory_equal(host.memory + 0x4000, "PING", 4);
    memcpy(tx_desc, host.memory + 0x2000, sizeof(tx_desc));
    memcpy(rx_desc, host.memory + 0x1000, sizeof(rx_desc));
    assert_true((tx_desc[11] & 0x02) != 0);
    assert_int_equal(tx_desc[12] & 0x01, 1);
    assert_int_equal(tx_desc[13], 0xa5);
    assert_int_equal(rx_desc[10], 0x5a);
    assert_int_equal(rx_desc[13], 0xa5);
    memcpy(&received_len, rx_desc + 8, sizeof(received_len));
    assert_int_equal(received_len, packet_len);
    assert_int_equal(rx_desc[12], 0x03);
    assert_int_equal(host.irqs, 1);
    assert_int_equal(behavior.read(behavior.state, 0, 0x2818, &value, 4), 4);
    assert_int_equal(value, 1);
    assert_int_equal(behavior.read(behavior.state, 0, 0x2810, &value, 4), 4);
    assert_int_equal(value, 1);
    assert_int_equal(behavior.read(behavior.state, 0, 0x00c0, &value, 4), 4);
    assert_int_equal(value, 0x81);
    assert_int_equal(behavior.read(behavior.state, 0, 0x00c0, &value, 4), 4);
    assert_int_equal(value, 0);
    behavior.destroy(behavior.state);
    device_model_free(model);
}

static void ethernet_injects_fake_arp(void **state)
{
    struct device_model *model = NULL;
    struct device_behavior behavior = {0};
    struct ethernet_host host = {0};
    struct behavior_host_ops ops = {
        .opaque = &host, .dma_read = ethernet_dma_read,
        .dma_write = ethernet_dma_write, .irq = ethernet_irq,
    };
    char err[256] = {0};
    uint32_t value;
    uint32_t base;
    uint8_t rx_desc[16] = {0};
    uint64_t rx_buffer = 0x5000;
    uint16_t packet_len;

    (void)state;
    assert_int_equal(device_model_load("../tests/cocotb/out_ethernet", &model,
                                      err, sizeof(err)), 0);
    assert_int_equal(behavior_ethernet_create(model, &behavior, err, sizeof(err)), 0);
    assert_int_equal(behavior.bind_host(behavior.state, &ops), 0);
    base = 0x1000;
    assert_int_equal(behavior.write(behavior.state, 0, 0x2800, &base, 4), 4);
    base = 0x2000;
    assert_int_equal(behavior.write(behavior.state, 0, 0x3800, &base, 4), 4);
    value = 128;
    assert_int_equal(behavior.write(behavior.state, 0, 0x2808, &value, 4), 4);
    assert_int_equal(behavior.write(behavior.state, 0, 0x3808, &value, 4), 4);
    memcpy(rx_desc, &rx_buffer, sizeof(rx_buffer));
    memcpy(host.memory + 0x1000, rx_desc, sizeof(rx_desc));
    value = 1u << 7;
    assert_int_equal(behavior.write(behavior.state, 0, 0x00d0, &value, 4), 4);
    value = 1;
    assert_int_equal(behavior.write(behavior.state, 0, 0x2818, &value, 4), 4);
    assert_memory_equal(host.memory + 0x5000,
                        "\xff\xff\xff\xff\xff\xff\x02\x50\x43\x49\x4c\x45\x08\x06", 14);
    memcpy(rx_desc, host.memory + 0x1000, sizeof(rx_desc));
    memcpy(&packet_len, rx_desc + 8, sizeof(packet_len));
    assert_int_equal(packet_len, 60);
    assert_int_equal(rx_desc[12], 0x03);
    assert_int_equal(host.irqs, 1);
    assert_int_equal(behavior.read(behavior.state, 0, 0x2810, &value, 4), 4);
    assert_int_equal(value, 1);
    memcpy(&rx_buffer, &(uint64_t){0x5010}, sizeof(rx_buffer));
    memcpy(host.memory + 0x1010, &rx_buffer, sizeof(rx_buffer));
    value = 2;
    assert_int_equal(behavior.write(behavior.state, 0, 0x2818, &value, 4), 4);
    assert_memory_equal(host.memory + 0x5010,
                        "\x02\x50\x43\x49\x4c\x46\x02\x50\x43\x49\x4c\x45\x08\x00", 14);
    assert_int_equal(ethernet_test_checksum(host.memory + 0x5010 + 14, 20), 0xffff);
    assert_int_equal(ethernet_test_checksum(host.memory + 0x5010 + 34, 12), 0xffff);
    memcpy(rx_desc, host.memory + 0x1010, sizeof(rx_desc));
    memcpy(&packet_len, rx_desc + 8, sizeof(packet_len));
    assert_int_equal(packet_len, 60);
    assert_int_equal(rx_desc[12], 0x03);
    assert_int_equal(host.irqs, 2);
    assert_int_equal(behavior.read(behavior.state, 0, 0x2810, &value, 4), 4);
    assert_int_equal(value, 2);
    behavior.destroy(behavior.state);
    device_model_free(model);
}


int main(void)
{
    const struct CMUnitTest tests[] = {
        cmocka_unit_test(ahci_reset_self_clears),
        cmocka_unit_test(xhci_reset_and_run_state),
        cmocka_unit_test(ethernet_link_status_and_reset),
        cmocka_unit_test(ethernet_descriptor_loopback),
        cmocka_unit_test(ethernet_injects_fake_arp),
    };

    return cmocka_run_group_tests(tests, NULL, NULL);
}
