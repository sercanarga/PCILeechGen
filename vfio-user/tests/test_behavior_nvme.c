#include <stdarg.h>
#include <stddef.h>
#include <setjmp.h>
#include <stdint.h>
#include <stdlib.h>

#include <cmocka.h>

#include "device_behavior.h"
#include "device_model.h"


struct fixture {
    struct device_model *model;
    struct device_behavior behavior;
};


static int setup(void **state)
{
    struct fixture *fixture = calloc(1, sizeof(*fixture));
    char err[256] = {0};

    assert_non_null(fixture);
    assert_int_equal(device_model_load("../tests/cocotb/out_nvme", &fixture->model, err, sizeof(err)), 0);
    assert_int_equal(behavior_nvme_create(fixture->model, &fixture->behavior, err, sizeof(err)), 0);
    *state = fixture;
    return 0;
}


static int teardown(void **state)
{
    struct fixture *fixture = *state;

    fixture->behavior.destroy(fixture->behavior.state);
    device_model_free(fixture->model);
    free(fixture);
    return 0;
}


static uint32_t read32(struct fixture *fixture, uint64_t offset)
{
    uint32_t value = 0;

    assert_int_equal(fixture->behavior.read(fixture->behavior.state, 0, offset, &value, sizeof(value)), 4);
    return value;
}


static void exposes_generated_capability_registers(void **state)
{
    struct fixture *fixture = *state;

    assert_int_equal(read32(fixture, 0x00), 0x0040ff17);
    assert_int_equal(read32(fixture, 0x04), 0x00000020);
    assert_int_equal(read32(fixture, 0x08), 0x00010400);
}


static void transitions_enable_and_shutdown_state(void **state)
{
    struct fixture *fixture = *state;
    uint32_t cc = 1;

    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x14, &cc, sizeof(cc)), 4);
    assert_int_equal(read32(fixture, 0x1c) & 1, 1);

    cc = 1 | (1u << 14);
    assert_int_equal(fixture->behavior.write(fixture->behavior.state, 0, 0x14, &cc, sizeof(cc)), 4);
    assert_int_equal((read32(fixture, 0x1c) >> 2) & 3, 2);

    assert_int_equal(fixture->behavior.reset(fixture->behavior.state), 0);
    assert_int_equal(read32(fixture, 0x1c) & 0x0d, 0);
}


int main(void)
{
    const struct CMUnitTest tests[] = {
        cmocka_unit_test_setup_teardown(exposes_generated_capability_registers, setup, teardown),
        cmocka_unit_test_setup_teardown(transitions_enable_and_shutdown_state, setup, teardown),
    };

    return cmocka_run_group_tests(tests, NULL, NULL);
}
