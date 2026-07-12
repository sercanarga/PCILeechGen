#ifndef PCILEECH_VFIO_DEVICE_H
#define PCILEECH_VFIO_DEVICE_H

#include <signal.h>

#include "device_behavior.h"
#include "device_model.h"


int vfio_device_run(const struct device_model *model,
                    struct device_behavior *behavior,
                    const char *socket_path,
                    volatile sig_atomic_t *stop);

#endif
