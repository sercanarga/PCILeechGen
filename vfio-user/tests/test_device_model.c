#include <stdarg.h>
#include <stddef.h>
#include <setjmp.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include <cmocka.h>

#include "device_model.h"


static void loads_generated_nvme_model(void **state)
{
    struct device_model *model = NULL;
    char err[256] = {0};

    (void)state;
    assert_int_equal(device_model_load("../tests/cocotb/out_nvme", &model, err, sizeof(err)), 0);
    assert_non_null(model);
    assert_int_equal(model->vendor_id, 0x144d);
    assert_int_equal(model->device_id, 0xa809);
    assert_int_equal(model->class_code, 0x010802);
    assert_int_equal(model->bars[0].size, 16384);
    assert_non_null(model->bars[0].reset_image);
    assert_memory_equal(model->bars[0].reset_image, "\x17\xff\x40\x00", 4);
    assert_memory_equal(model->config_space, "\x4d\x14\x09\xa8", 4);
    device_model_free(model);
}


static void rejects_non_power_of_two_bar(void **state)
{
    struct device_model model = {0};
    char err[256] = {0};

    (void)state;
    model.config_space_size = 256;
    model.bar_count = 1;
    model.bars[0].bir = 0;
    model.bars[0].size = 6000;
    assert_true(device_model_validate(&model, err, sizeof(err)) < 0);
    assert_non_null(strstr(err, "power of two"));
}


static void rejects_capability_cycle(void **state)
{
    struct device_model model = {0};
    char err[256] = {0};

    (void)state;
    model.config_space_size = 256;
    model.config_space[0x06] = 0x10;
    model.config_space[0x34] = 0x40;
    model.config_space[0x40] = 0x01;
    model.config_space[0x41] = 0x40;
    assert_true(device_model_validate(&model, err, sizeof(err)) < 0);
    assert_non_null(strstr(err, "capability chain"));
}


static void loads_every_generated_model(void **state)
{
    static const char *names[] = {
        "audio", "ethernet", "generic", "gpu", "multibar",
        "nvme", "sata", "thunderbolt", "wifi", "xhci",
    };
    size_t index;

    (void)state;
    for (index = 0; index < sizeof(names) / sizeof(names[0]); ++index) {
        struct device_model *model = NULL;
        char path[128];
        char err[256] = {0};

        assert_true(snprintf(path, sizeof(path), "../tests/cocotb/out_%s", names[index]) > 0);
        assert_int_equal(device_model_load(path, &model, err, sizeof(err)), 0);
        assert_non_null(model);
        device_model_free(model);
    }
}


int main(void)
{
    const struct CMUnitTest tests[] = {
        cmocka_unit_test(loads_generated_nvme_model),
        cmocka_unit_test(rejects_non_power_of_two_bar),
        cmocka_unit_test(rejects_capability_cycle),
        cmocka_unit_test(loads_every_generated_model),
    };

    return cmocka_run_group_tests(tests, NULL, NULL);
}
