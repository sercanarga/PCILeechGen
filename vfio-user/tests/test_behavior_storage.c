#include <stdarg.h>
#include <stddef.h>
#include <setjmp.h>
#include <stdint.h>
#include <stdlib.h>

#include <cmocka.h>

#include "device_behavior.h"
#include "device_model.h"


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
    assert_int_equal(read32(&behavior, 0x0c), 1);
    ghc = 0x80000001;
    assert_int_equal(behavior.write(behavior.state, 0, 0x04, &ghc, 4), 4);
    assert_int_equal(read32(&behavior, 0x04) & 1, 0);
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
    behavior.destroy(behavior.state);
    device_model_free(model);
}


int main(void)
{
    const struct CMUnitTest tests[] = {
        cmocka_unit_test(ahci_reset_self_clears),
        cmocka_unit_test(xhci_reset_and_run_state),
    };

    return cmocka_run_group_tests(tests, NULL, NULL);
}
