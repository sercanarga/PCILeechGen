#include <stdarg.h>
#include <stddef.h>
#include <setjmp.h>
#include <stdint.h>
#include <errno.h>
#include <stdlib.h>
#include <string.h>

#include <cmocka.h>

#include "device_behavior.h"


static struct device_model one_bar_model(void)
{
    struct device_model model = {0};

    model.config_space_size = 256;
    model.bar_count = 1;
    model.bars[0].bir = 0;
    model.bars[0].type = DEVICE_BAR_MEMORY;
    model.bars[0].size = 4096;
    return model;
}


static void supports_partial_writes_and_reset(void **state)
{
    struct device_model model = one_bar_model();
    struct device_behavior behavior = {0};
    uint32_t value = 0x11223344;
    uint16_t partial = 0xaabb;
    uint32_t result = 0;
    char err[128] = {0};

    (void)state;
    assert_int_equal(behavior_create(&model, &behavior, err, sizeof(err)), 0);
    assert_int_equal(behavior.write(behavior.state, 0, 4, &value, sizeof(value)), 4);
    assert_int_equal(behavior.write(behavior.state, 0, 5, &partial, sizeof(partial)), 2);
    assert_int_equal(behavior.read(behavior.state, 0, 4, &result, sizeof(result)), 4);
    assert_int_equal(result, 0x11aabb44);
    assert_int_equal(behavior.reset(behavior.state), 0);
    assert_int_equal(behavior.read(behavior.state, 0, 4, &result, sizeof(result)), 4);
    assert_int_equal(result, 0);
    behavior.destroy(behavior.state);
}


static void rejects_out_of_range_access(void **state)
{
    struct device_model model = one_bar_model();
    struct device_behavior behavior = {0};
    uint32_t value = 0;
    char err[128] = {0};

    (void)state;
    assert_int_equal(behavior_create(&model, &behavior, err, sizeof(err)), 0);
    assert_int_equal(behavior.read(behavior.state, 0, 4094, &value, sizeof(value)), -EINVAL);
    assert_int_equal(behavior.write(behavior.state, 1, 0, &value, sizeof(value)), -EINVAL);
    behavior.destroy(behavior.state);
}


static void starts_from_generated_bar_reset_image(void **state)
{
    struct device_model model = one_bar_model();
    struct device_behavior behavior = {0};
    uint32_t result = 0;
    char err[128] = {0};

    (void)state;
    model.bars[0].reset_image = calloc(1, model.bars[0].size);
    assert_non_null(model.bars[0].reset_image);
    memcpy(model.bars[0].reset_image, "\x17\xff\x40\x00", 4);
    assert_int_equal(behavior_static_create(&model, &behavior, err, sizeof(err)), 0);
    assert_int_equal(behavior.read(behavior.state, 0, 0, &result, sizeof(result)), 4);
    assert_int_equal(result, 0x0040ff17);
    result = 0;
    assert_int_equal(behavior.write(behavior.state, 0, 0, &result, sizeof(result)), 4);
    assert_int_equal(behavior.reset(behavior.state), 0);
    assert_int_equal(behavior.read(behavior.state, 0, 0, &result, sizeof(result)), 4);
    assert_int_equal(result, 0x0040ff17);
    behavior.destroy(behavior.state);
    free(model.bars[0].reset_image);
}


int main(void)
{
    const struct CMUnitTest tests[] = {
        cmocka_unit_test(supports_partial_writes_and_reset),
        cmocka_unit_test(rejects_out_of_range_access),
        cmocka_unit_test(starts_from_generated_bar_reset_image),
    };

    return cmocka_run_group_tests(tests, NULL, NULL);
}
