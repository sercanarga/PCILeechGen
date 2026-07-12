#ifndef PCILEECH_VFIO_DEVICE_BEHAVIOR_H
#define PCILEECH_VFIO_DEVICE_BEHAVIOR_H

#include <stddef.h>
#include <stdint.h>
#include <sys/types.h>

#include "device_model.h"


struct behavior_host_ops {
    void *opaque;
    int (*dma_read)(void *opaque, uint64_t address, void *data, size_t length);
    int (*dma_write)(void *opaque, uint64_t address, const void *data, size_t length);
    int (*irq)(void *opaque, unsigned vector);
};


struct device_behavior {
    void *state;
    int (*bind_host)(void *state, const struct behavior_host_ops *ops);
    int (*reset)(void *state);
    int (*service)(void *state);
    ssize_t (*read)(void *state, unsigned bir, uint64_t offset, void *buf, size_t len);
    ssize_t (*write)(void *state, unsigned bir, uint64_t offset, const void *buf, size_t len);
    void (*destroy)(void *state);
};

int behavior_create(const struct device_model *model,
                    struct device_behavior *out, char *err, size_t err_len);
int behavior_static_create(const struct device_model *model,
                           struct device_behavior *out, char *err, size_t err_len);
int behavior_nvme_create(const struct device_model *model,
                         struct device_behavior *out, char *err, size_t err_len);
int behavior_hda_create(const struct device_model *model,
                        struct device_behavior *out, char *err, size_t err_len);
int behavior_ahci_create(const struct device_model *model,
                         struct device_behavior *out, char *err, size_t err_len);
int behavior_xhci_create(const struct device_model *model,
                         struct device_behavior *out, char *err, size_t err_len);
int behavior_ethernet_create(const struct device_model *model,
                             struct device_behavior *out, char *err, size_t err_len);

#endif
