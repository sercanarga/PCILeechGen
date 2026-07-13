#include <stdint.h>
#include <stdlib.h>
#include <string.h>

#include "device_behavior.h"
#include "device_model.h"

struct fuzz_host {
    uint8_t memory[1 << 20];
    unsigned irqs;
};

static int fuzz_read(void *opaque, uint64_t address, void *data, size_t length)
{
    struct fuzz_host *host = opaque;
    if (address > sizeof(host->memory) || length > sizeof(host->memory) - address) return -1;
    memcpy(data, host->memory + address, length);
    return 0;
}

static int fuzz_write(void *opaque, uint64_t address, const void *data, size_t length)
{
    struct fuzz_host *host = opaque;
    if (address > sizeof(host->memory) || length > sizeof(host->memory) - address) return -1;
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

static struct device_behavior *load_behavior(const char *path, int ethernet,
                                              struct device_model **model)
{
    struct device_behavior *behavior = calloc(1, sizeof(*behavior));
    char error[256] = {0};
    if (behavior == NULL || device_model_load(path, model, error, sizeof(error)) < 0) {
        free(behavior);
        return NULL;
    }
    if ((ethernet ? behavior_ethernet_create : behavior_nvme_create)(*model, behavior,
                                                                       error, sizeof(error)) < 0) {
        device_model_free(*model);
        free(behavior);
        return NULL;
    }
    return behavior;
}

int LLVMFuzzerTestOneInput(const uint8_t *data, size_t size)
{
    struct device_model *nvme_model = NULL, *ethernet_model = NULL;
    struct device_behavior *nvme = NULL, *ethernet = NULL;
    struct fuzz_host host = {0};
    struct behavior_host_ops ops = {
        .opaque = &host, .dma_read = fuzz_read, .dma_write = fuzz_write, .irq = fuzz_irq,
    };
    uint32_t value;

    if (data == NULL || size == 0) return 0;
    nvme = load_behavior("../tests/cocotb/out_nvme", 0, &nvme_model);
    ethernet = load_behavior("../tests/cocotb/out_ethernet", 1, &ethernet_model);
    if (nvme == NULL || ethernet == NULL) goto done;
    nvme->bind_host(nvme->state, &ops);
    ethernet->bind_host(ethernet->state, &ops);

    value = (uint32_t)data[0] | ((size > 1 ? data[1] : 0) << 8) |
            ((size > 2 ? data[2] : 0) << 16) | ((size > 3 ? data[3] : 0) << 24);
    if ((data[0] & 1) == 0) {
        uint64_t queue = 0x1000;
        uint32_t aqa = 3 | (3u << 16);
        nvme->write(nvme->state, 0, 0x24, &aqa, 4);
        nvme->write(nvme->state, 0, 0x28, &queue, 8);
        queue = 0x2000;
        nvme->write(nvme->state, 0, 0x30, &queue, 8);
        nvme->write(nvme->state, 0, 0x14, &value, 4);
        nvme->write(nvme->state, 0, 0x1000, &value, 4);
    } else {
        uint64_t ring = 0x3000;
        ethernet->write(ethernet->state, 0, 0x280, &ring, 4);
        ethernet->write(ethernet->state, 0, 0x380, &ring, 4);
        ethernet->write(ethernet->state, 0, 0x288, &value, 4);
        ethernet->write(ethernet->state, 0, 0x388, &value, 4);
        ethernet->write(ethernet->state, 0, 0x3818, &value, 4);
    }

done:
    if (nvme != NULL) {
        nvme->destroy(nvme->state);
        free(nvme);
    }
    if (ethernet != NULL) {
        ethernet->destroy(ethernet->state);
        free(ethernet);
    }
    device_model_free(nvme_model);
    device_model_free(ethernet_model);
    return 0;
}

#ifdef FUZZ_SMOKE
int main(void)
{
    static const uint8_t cases[][8] = {{0}, {1, 0xff, 0x55}, {2, 0xaa, 0x11, 0x7f}};
    for (size_t i = 0; i < sizeof(cases) / sizeof(cases[0]); ++i) {
        if (LLVMFuzzerTestOneInput(cases[i], sizeof(cases[i])) != 0) return 1;
    }
    return 0;
}
#endif
