#include <limits.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "device_behavior.h"
#include "device_model.h"

struct fuzz_host {
    uint8_t memory[1 << 20];
    unsigned irqs;
};

struct fuzz_fixture {
    bool ethernet;
    bool loaded;
    bool failed;
    struct device_model *model;
    struct device_behavior behavior;
};

static struct fuzz_fixture nvme_fixture = {.ethernet = false};
static struct fuzz_fixture ethernet_fixture = {.ethernet = true};

static int fuzz_read(void *opaque, uint64_t address, void *data, size_t length)
{
    struct fuzz_host *host = opaque;

    if (address > sizeof(host->memory) || length > sizeof(host->memory) - address)
        return -1;
    memcpy(data, host->memory + address, length);
    return 0;
}

static int fuzz_write(void *opaque, uint64_t address, const void *data, size_t length)
{
    struct fuzz_host *host = opaque;

    if (address > sizeof(host->memory) || length > sizeof(host->memory) - address)
        return -1;
    memcpy(host->memory + address, data, length);
    return 0;
}

static int fuzz_irq(void *opaque, unsigned vector)
{
    struct fuzz_host *host = opaque;

    (void)vector;
    host->irqs++;
    return 0;
}

static void fixture_destroy(struct fuzz_fixture *fixture)
{
    if (fixture->behavior.destroy != NULL && fixture->behavior.state != NULL)
        fixture->behavior.destroy(fixture->behavior.state);
    device_model_free(fixture->model);
    fixture->model = NULL;
    fixture->behavior = (struct device_behavior){0};
    fixture->loaded = false;
}

static int fixture_load(struct fuzz_fixture *fixture, const char *path)
{
    char error[256] = {0};
    int rc;

    if (fixture->loaded)
        return 0;
    if (fixture->failed)
        return -1;
    if (device_model_load(path, &fixture->model, error, sizeof(error)) < 0)
        goto failed;
    rc = fixture->ethernet
             ? behavior_ethernet_create(fixture->model, &fixture->behavior, error, sizeof(error))
             : behavior_nvme_create(fixture->model, &fixture->behavior, error, sizeof(error));
    if (rc < 0)
        goto failed;
    fixture->loaded = true;
    return 0;

failed:
    fixture_destroy(fixture);
    fixture->failed = true;
    return -1;
}

static int fixture_reset(struct fuzz_fixture *fixture, const struct behavior_host_ops *ops)
{
    if (!fixture->loaded || fixture->behavior.reset == NULL || fixture->behavior.bind_host == NULL)
        return -1;
    if (fixture->behavior.reset(fixture->behavior.state) < 0)
        return -1;
    return fixture->behavior.bind_host(fixture->behavior.state, ops);
}

static int fixture_path(const char *root, const char *name, char path[PATH_MAX])
{
    int length = snprintf(path, PATH_MAX, "%s/%s", root, name);

    return length >= 0 && length < PATH_MAX ? 0 : -1;
}

static int prepare_fixtures(const char *root, const struct behavior_host_ops *ops)
{
    char nvme_path[PATH_MAX];
    char ethernet_path[PATH_MAX];

    if (fixture_path(root, "nvme", nvme_path) < 0 ||
        fixture_path(root, "ethernet", ethernet_path) < 0 ||
        fixture_load(&nvme_fixture, nvme_path) < 0 ||
        fixture_load(&ethernet_fixture, ethernet_path) < 0 ||
        fixture_reset(&nvme_fixture, ops) < 0 ||
        fixture_reset(&ethernet_fixture, ops) < 0) {
        fixture_destroy(&nvme_fixture);
        fixture_destroy(&ethernet_fixture);
        return -1;
    }
    return 0;
}

int LLVMFuzzerTestOneInput(const uint8_t *data, size_t size)
{
    const char *fixture_root = getenv("VFIO_FUZZ_FIXTURE_ROOT");
    struct fuzz_host host = {0};
    struct behavior_host_ops ops = {
        .opaque = &host,
        .dma_read = fuzz_read,
        .dma_write = fuzz_write,
        .irq = fuzz_irq,
    };
    uint32_t value;

    if (data == NULL || size == 0 || fixture_root == NULL || fixture_root[0] == '\0')
        return 0;
    if (prepare_fixtures(fixture_root, &ops) < 0)
        return 0;

    value = (uint32_t)data[0] |
            ((uint32_t)(size > 1 ? data[1] : 0) << 8) |
            ((uint32_t)(size > 2 ? data[2] : 0) << 16) |
            ((uint32_t)(size > 3 ? data[3] : 0) << 24);
    if ((data[0] & 1u) == 0) {
        uint64_t queue = 0x1000;
        uint32_t aqa = 3 | (3u << 16);

        nvme_fixture.behavior.write(nvme_fixture.behavior.state, 0, 0x24, &aqa, 4);
        nvme_fixture.behavior.write(nvme_fixture.behavior.state, 0, 0x28, &queue, 8);
        queue = 0x2000;
        nvme_fixture.behavior.write(nvme_fixture.behavior.state, 0, 0x30, &queue, 8);
        nvme_fixture.behavior.write(nvme_fixture.behavior.state, 0, 0x14, &value, 4);
        nvme_fixture.behavior.write(nvme_fixture.behavior.state, 0, 0x1000, &value, 4);
    } else {
        uint64_t ring = 0x3000;

        ethernet_fixture.behavior.write(ethernet_fixture.behavior.state, 0, 0x280, &ring, 4);
        ethernet_fixture.behavior.write(ethernet_fixture.behavior.state, 0, 0x380, &ring, 4);
        ethernet_fixture.behavior.write(ethernet_fixture.behavior.state, 0, 0x288, &value, 4);
        ethernet_fixture.behavior.write(ethernet_fixture.behavior.state, 0, 0x388, &value, 4);
        ethernet_fixture.behavior.write(ethernet_fixture.behavior.state, 0, 0x3818, &value, 4);
    }
    return 0;
}

__attribute__((destructor)) static void fuzz_cleanup(void)
{
    fixture_destroy(&nvme_fixture);
    fixture_destroy(&ethernet_fixture);
}

#ifdef FUZZ_SMOKE
int main(void)
{
    static const uint8_t cases[][8] = {
        {0},
        {1, 0xff, 0x55},
        {2, 0xaa, 0x11, 0x7f},
    };
    size_t index;

    for (index = 0; index < sizeof(cases) / sizeof(cases[0]); ++index) {
        if (LLVMFuzzerTestOneInput(cases[index], sizeof(cases[index])) != 0)
            return 1;
    }
    return 0;
}
#endif
