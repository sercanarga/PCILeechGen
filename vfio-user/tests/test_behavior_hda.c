#include <setjmp.h>
#include <stdarg.h>
#include <errno.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

#include <cmocka.h>

#include "device_behavior.h"
#include "device_model.h"

struct host_memory {
    uint8_t data[0x10000];
    unsigned irqs;
};

static struct device_model *fixture_model;

static int dma_read(void *opaque, uint64_t address, void *data, size_t length)
{
    struct host_memory *host = opaque;
    if (address > sizeof(host->data) || length > sizeof(host->data) - address) return -1;
    memcpy(data, host->data + address, length);
    return 0;
}

static int dma_write(void *opaque, uint64_t address, const void *data, size_t length)
{
    struct host_memory *host = opaque;
    if (address > sizeof(host->data) || length > sizeof(host->data) - address) return -1;
    memcpy(host->data + address, data, length);
    return 0;
}

static int irq(void *opaque, unsigned vector)
{
    struct host_memory *host = opaque;
    assert_int_equal(vector, 0);
    host->irqs++;
    return 0;
}

static int setup(void **state)
{
    struct device_behavior *behavior = calloc(1, sizeof(*behavior));
    struct device_model *model = NULL;
    char err[128] = {0};

    assert_non_null(behavior);
    assert_int_equal(device_model_load("../tests/cocotb/out_audio", &model, err, sizeof(err)), 0);
    assert_int_equal(behavior_hda_create(model, behavior, err, sizeof(err)), 0);
    fixture_model = model;
    *state = behavior;
    return 0;
}

static int teardown(void **state)
{
    struct device_behavior *behavior = *state;
    behavior->destroy(behavior->state);
    free(behavior);
    device_model_free(fixture_model);
    fixture_model = NULL;
    return 0;
}

static void reset_and_process_corb(void **state)
{
    struct device_behavior *behavior = *state;
    struct host_memory host = {0};
    struct behavior_host_ops ops = {
        .opaque = &host, .dma_read = dma_read, .dma_write = dma_write, .irq = irq,
    };
    uint64_t base;
    uint32_t value;
    uint64_t response;

    assert_int_equal(behavior->bind_host(behavior->state, NULL), -EINVAL);
    assert_int_equal(behavior->bind_host(behavior->state, &ops), 0);
    base = 0x1000;
    assert_int_equal(behavior->write(behavior->state, 0, 0x40, &base, 8), 8);
    base = 0x2000;
    assert_int_equal(behavior->write(behavior->state, 0, 0x50, &base, 8), 8);
    value = 2;
    assert_int_equal(behavior->write(behavior->state, 0, 0x4c, &value, 4), 4);
    assert_int_equal(behavior->write(behavior->state, 0, 0x5c, &value, 4), 4);
    value = 0x00112233;
    memcpy(&host.data[0x1004], &value, sizeof(value));
    value = 1;
    assert_int_equal(behavior->write(behavior->state, 0, 0x48, &value, 4), 4);
    assert_int_equal(behavior->service(behavior->state), 0);
    memcpy(&response, &host.data[0x2008], sizeof(response));
    assert_int_equal(response, UINT64_C(0x0011223300000000));
    assert_int_equal(host.irqs, 1);

    value = 0;
    assert_int_equal(behavior->write(behavior->state, 0, 0x08, &value, 4), 4);
    assert_int_equal(behavior->read(behavior->state, 0, 0x48, &value, 4), 4);
    assert_int_equal(value, 0);
    assert_int_equal(behavior->read(behavior->state, 0, 0x58, &value, 4), 4);
    assert_int_equal(value, 0);
}

int main(void)
{
    const struct CMUnitTest tests[] = {
        cmocka_unit_test_setup_teardown(reset_and_process_corb, setup, teardown),
    };
    return cmocka_run_group_tests(tests, NULL, NULL);
}
