#include <stdint.h>
#include <stdlib.h>

#include <cmocka.h>

#include "device_behavior.h"
#include "device_model.h"

static void probes_mmio_timer_and_fence(void **state)
{
    struct device_model *model = NULL;
    struct device_behavior behavior = {0};
    char err[128] = {0};
    uint32_t value = 0;
    uint32_t fence = 0x12345678;

    (void)state;
    assert_int_equal(device_model_load("../tests/cocotb/out_gpu", &model, err, sizeof(err)), 0);
    assert_int_equal(behavior_gpu_create(model, &behavior, err, sizeof(err)), 0);
    assert_int_equal(behavior.read(behavior.state, 0, 0x200, &value, 4), 4);
    assert_int_equal(value, 0xffffffff);
    assert_int_equal(behavior.read(behavior.state, 0, 0x9400, &value, 4), 4);
    assert_int_equal(value, 0);
    assert_int_equal(behavior.write(behavior.state, 0, 0x1000, &fence, 4), 4);
    assert_int_equal(behavior.read(behavior.state, 0, 0x1004, &value, 4), 4);
    assert_int_equal(value, fence);
    behavior.destroy(behavior.state);
    device_model_free(model);
}

int main(void)
{
    const struct CMUnitTest tests[] = {cmocka_unit_test(probes_mmio_timer_and_fence)};
    return cmocka_run_group_tests(tests, NULL, NULL);
}
