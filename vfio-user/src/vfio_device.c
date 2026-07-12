#include "vfio_device.h"

#include <errno.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/mman.h>
#include <sys/uio.h>
#include <time.h>
#include <unistd.h>

#include <vfio-user/libvfio-user.h>
#include <vfio-user/pci_defs.h>


struct server_state {
    const struct device_model *model;
    struct device_behavior *behavior;
    vfu_ctx_t *context;
};


static ssize_t access_bir(vfu_ctx_t *context, unsigned bir, char *buf,
                          size_t count, loff_t offset, bool write)
{
    struct server_state *state = vfu_get_private(context);
    ssize_t result = write
        ? state->behavior->write(state->behavior->state, bir, (uint64_t)offset, buf, count)
        : state->behavior->read(state->behavior->state, bir, (uint64_t)offset, buf, count);

    if (result < 0) {
        errno = (int)-result;
        return -1;
    }
    return result;
}


#define DEFINE_BAR_ACCESS(bir) \
static ssize_t bar##bir##_access(vfu_ctx_t *context, char *buf, size_t count, \
                                 loff_t offset, bool write) \
{ \
    return access_bir(context, bir, buf, count, offset, write); \
}

DEFINE_BAR_ACCESS(0)
DEFINE_BAR_ACCESS(1)
DEFINE_BAR_ACCESS(2)
DEFINE_BAR_ACCESS(3)
DEFINE_BAR_ACCESS(4)
DEFINE_BAR_ACCESS(5)


static vfu_region_access_cb_t *const bar_callbacks[DEVICE_MAX_BARS] = {
    bar0_access, bar1_access, bar2_access, bar3_access, bar4_access, bar5_access,
};


static int device_reset(vfu_ctx_t *context, vfu_reset_type_t type)
{
    struct server_state *state = vfu_get_private(context);

    (void)type;
    memcpy(vfu_pci_get_config_space(context), state->model->config_space,
           state->model->config_space_size);
    return state->behavior->reset(state->behavior->state);
}


static ssize_t config_access(vfu_ctx_t *context, char *buf, size_t count,
                             loff_t offset, bool is_write)
{
    struct server_state *state = vfu_get_private(context);

    if (offset < 0 || (uint64_t)offset > state->model->config_space_size ||
        count > state->model->config_space_size - (size_t)offset) {
        errno = EINVAL;
        return -1;
    }
    uint8_t *config = (uint8_t *)vfu_pci_get_config_space(context);

    if (is_write) {
        memcpy(config + offset, buf, count);
    } else {
        memcpy(buf, config + offset, count);
    }
    return (ssize_t)count;
}


static void dma_register(vfu_ctx_t *context, vfu_dma_info_t *info)
{
    (void)context;
    (void)info;
}


static void dma_unregister(vfu_ctx_t *context, vfu_dma_info_t *info)
{
    (void)context;
    (void)info;
}


static int dma_transfer(void *opaque, uint64_t address, void *data, size_t length, bool write)
{
    struct server_state *state = opaque;
    uint8_t *cursor = data;
    dma_sg_t *sg = malloc(dma_sg_size());

    if (sg == NULL) {
        return -1;
    }

    while (length > 0) {
        size_t chunk = 4096 - (size_t)(address & 0xfff);
        int protection = write ? PROT_WRITE : PROT_READ;

        if (chunk > length) {
            chunk = length;
        }
        if (vfu_addr_to_sgl(state->context, (vfu_dma_addr_t)address,
                            chunk, sg, 1, protection) != 1) {
            free(sg);
            return -1;
        }
        struct iovec iov;
        if (vfu_sgl_get(state->context, sg, &iov, 1, 0) == 0) {
            if (iov.iov_len < chunk) {
                vfu_sgl_put(state->context, sg, &iov, 1);
                free(sg);
                return -1;
            }
            if (write) {
                memcpy(iov.iov_base, cursor, chunk);
            } else {
                memcpy(cursor, iov.iov_base, chunk);
            }
            vfu_sgl_put(state->context, sg, &iov, 1);
        } else if ((write ? vfu_sgl_write(state->context, sg, 1, cursor, 0)
                           : vfu_sgl_read(state->context, sg, 1, cursor, 0)) < 0) {
            free(sg);
            return -1;
        }
        address += chunk;
        cursor += chunk;
        length -= chunk;
    }
    free(sg);
    return 0;
}


static int host_dma_read(void *opaque, uint64_t address, void *data, size_t length)
{
    return dma_transfer(opaque, address, data, length, false);
}


static int host_dma_write(void *opaque, uint64_t address, const void *data, size_t length)
{
    return dma_transfer(opaque, address, (void *)data, length, true);
}


static int host_irq(void *opaque, unsigned vector)
{
    struct server_state *state = opaque;
    return vfu_irq_trigger(state->context, vector);
}


static bool config_has_capability(const struct device_model *model, uint8_t id)
{
    uint8_t offset;
    unsigned steps = 0;

    if (model->config_space_size <= 0x35) {
        return false;
    }
    offset = model->config_space[0x34];
    while (offset >= 0x40 && (size_t)offset + 2 <= model->config_space_size && steps++ < 48) {
        if (model->config_space[offset] == id) {
            return true;
        }
        offset = model->config_space[offset + 1];
    }
    return false;
}

static int register_standard_capabilities(vfu_ctx_t *context,
                                          const struct device_model *model)
{
    uint8_t *config = (uint8_t *)vfu_pci_get_config_space(context);
    uint8_t offset = model->config_space[0x34];
    unsigned steps = 0;

    while (offset >= 0x40 && (size_t)offset + 2 <= model->config_space_size &&
           steps++ < 48) {
        uint8_t id = config[offset];

        if (id == PCI_CAP_ID_MSI || id == PCI_CAP_ID_MSIX) {
            if (vfu_pci_add_capability(context, offset, 0, config + offset) < 0) {
                return -1;
            }
        }
        offset = config[offset + 1];
    }
    return 0;
}


static int setup_regions(struct server_state *state)
{
    size_t index;

    for (index = 0; index < state->model->bar_count; ++index) {
        const struct device_bar *bar = &state->model->bars[index];
        int flags = VFU_REGION_FLAG_RW;

        if (bar->type == DEVICE_BAR_MEMORY) {
            flags |= VFU_REGION_FLAG_MEM;
        }
        if (bar->is_64bit) {
            flags |= VFU_REGION_FLAG_64_BITS;
        }
        if (bar->prefetchable) {
            flags |= VFU_REGION_FLAG_PREFETCH;
        }
        if (vfu_setup_region(state->context, VFU_PCI_DEV_BAR0_REGION_IDX + (int)bar->bir,
                             (size_t)bar->size, bar_callbacks[bar->bir], flags,
                             NULL, 0, -1, 0) < 0) {
            return -1;
        }
    }
    return 0;
}


int vfio_device_run(const struct device_model *model,
                    struct device_behavior *behavior,
                    const char *socket_path,
                    volatile sig_atomic_t *stop)
{
    struct server_state state = {
        .model = model,
        .behavior = behavior,
    };
    int result = -1;

    if (model == NULL || behavior == NULL || socket_path == NULL || stop == NULL) {
        errno = EINVAL;
        return -1;
    }
    unlink(socket_path);
    state.context = vfu_create_ctx(VFU_TRANS_SOCK, socket_path,
                                   LIBVFIO_USER_FLAG_ATTACH_NB, &state, VFU_DEV_TYPE_PCI);
    if (state.context == NULL ||
        vfu_pci_init(state.context, VFU_PCI_TYPE_EXPRESS, PCI_HEADER_TYPE_NORMAL, 0) < 0) {
        goto done;
    }
    memcpy(vfu_pci_get_config_space(state.context), model->config_space, model->config_space_size);
    if (register_standard_capabilities(state.context, model) < 0) {
        goto done;
    }
    if (vfu_setup_region(state.context, VFU_PCI_DEV_CFG_REGION_IDX,
                         model->config_space_size, config_access,
                         VFU_REGION_FLAG_RW | VFU_REGION_FLAG_ALWAYS_CB,
                         NULL, 0, -1, 0) < 0) {
        goto done;
    }
    struct behavior_host_ops host = {
        .opaque = &state,
        .dma_read = host_dma_read,
        .dma_write = host_dma_write,
        .irq = host_irq,
    };
    if (behavior->bind_host == NULL || behavior->bind_host(behavior->state, &host) < 0) {
        errno = EINVAL;
        goto done;
    }
    if (setup_regions(&state) < 0 ||
        vfu_setup_device_reset_cb(state.context, device_reset) < 0 ||
        vfu_setup_device_dma(state.context, LIBVFIO_USER_MAX_DMA_REGIONS,
                             dma_register, dma_unregister) < 0) {
        goto done;
    }
    if (model->msi_vectors > 0 &&
        vfu_setup_device_nr_irqs(state.context, VFU_DEV_MSI_IRQ, model->msi_vectors) < 0) {
        goto done;
    }
    if (model->config_space_size > 0x3d && model->config_space[0x3d] != 0 &&
        vfu_setup_device_nr_irqs(state.context, VFU_DEV_INTX_IRQ, 1) < 0) {
        goto done;
    }
    if (model->msix_vectors > 0 && config_has_capability(model, PCI_CAP_ID_MSIX) &&
        vfu_setup_device_nr_irqs(state.context, VFU_DEV_MSIX_IRQ, model->msix_vectors) < 0) {
        goto done;
    }
    if (vfu_realize_ctx(state.context) < 0) {
        goto done;
    }
    printf("{\"event\":\"ready\",\"vendor_id\":\"%04x\",\"device_id\":\"%04x\",\"class_code\":\"%06x\",\"bar_count\":%zu}\n",
           model->vendor_id, model->device_id, model->class_code, model->bar_count);
    fflush(stdout);
    while (!*stop) {
        if (vfu_attach_ctx(state.context) == 0) {
            break;
        }
        if (errno != EINTR && errno != EAGAIN && errno != EWOULDBLOCK) {
            goto done;
        }
        struct timespec delay = {.tv_sec = 0, .tv_nsec = 1000000};
        nanosleep(&delay, NULL);
    }
    while (!*stop) {
        int run = vfu_run_ctx(state.context);

        if (run >= 0 || errno == EINTR || errno == EBUSY ||
            errno == EAGAIN || errno == EWOULDBLOCK) {
            if (run >= 0 && behavior->service != NULL &&
                behavior->service(behavior->state) < 0) {
                goto done;
            }
            if (run == 0) {
                struct timespec delay = {.tv_sec = 0, .tv_nsec = 1000000};
                nanosleep(&delay, NULL);
            }
            continue;
        }
        if (errno == ENOTCONN) {
            result = 0;
            goto done;
        }
        goto done;
    }
    result = 0;

done:
    if (state.context != NULL) {
        vfu_destroy_ctx(state.context);
    }
    unlink(socket_path);
    return result;
}
