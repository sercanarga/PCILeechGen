#ifndef PCILEECH_VFIO_DEVICE_BEHAVIOR_H
#define PCILEECH_VFIO_DEVICE_BEHAVIOR_H

#include <stddef.h>
#include <stdint.h>
#include <sys/types.h>

#include "device_model.h"


struct device_behavior {
    void *state;
    int (*reset)(void *state);
    ssize_t (*read)(void *state, unsigned bir, uint64_t offset, void *buf, size_t len);
    ssize_t (*write)(void *state, unsigned bir, uint64_t offset, const void *buf, size_t len);
    void (*destroy)(void *state);
};

int behavior_create(const struct device_model *model,
                    struct device_behavior *out, char *err, size_t err_len);
int behavior_static_create(const struct device_model *model,
                           struct device_behavior *out, char *err, size_t err_len);

#endif
