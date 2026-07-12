#include "device_behavior.h"

#include <errno.h>
#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>


struct static_bar {
    uint8_t *data;
    uint8_t *reset;
    size_t size;
};

struct static_state {
    struct static_bar bars[DEVICE_MAX_BARS];
};


static int fail(char *err, size_t err_len, const char *format, ...)
{
    va_list args;

    if (err != NULL && err_len > 0) {
        va_start(args, format);
        vsnprintf(err, err_len, format, args);
        va_end(args);
    }
    return -1;
}


static int static_reset(void *opaque)
{
    struct static_state *state = opaque;
    unsigned bir;

    for (bir = 0; bir < DEVICE_MAX_BARS; ++bir) {
        if (state->bars[bir].data != NULL) {
            if (state->bars[bir].reset != NULL) {
                memcpy(state->bars[bir].data, state->bars[bir].reset, state->bars[bir].size);
            } else {
                memset(state->bars[bir].data, 0, state->bars[bir].size);
            }
        }
    }
    return 0;
}


static int static_bind_host(void *opaque, const struct behavior_host_ops *ops)
{
    (void)opaque;
    (void)ops;
    return 0;
}


static ssize_t static_access(struct static_state *state, unsigned bir, uint64_t offset,
                             void *buf, size_t len, int write)
{
    struct static_bar *bar;

    if (state == NULL || buf == NULL || bir >= DEVICE_MAX_BARS) {
        return -EINVAL;
    }
    bar = &state->bars[bir];
    if (bar->data == NULL || offset > bar->size || len > bar->size - offset) {
        return -EINVAL;
    }
    if (write) {
        memcpy(bar->data + offset, buf, len);
    } else {
        memcpy(buf, bar->data + offset, len);
    }
    return (ssize_t)len;
}


static ssize_t static_read(void *opaque, unsigned bir, uint64_t offset,
                           void *buf, size_t len)
{
    return static_access(opaque, bir, offset, buf, len, 0);
}


static ssize_t static_write(void *opaque, unsigned bir, uint64_t offset,
                            const void *buf, size_t len)
{
    return static_access(opaque, bir, offset, (void *)buf, len, 1);
}


static void static_destroy(void *opaque)
{
    struct static_state *state = opaque;
    unsigned bir;

    if (state == NULL) {
        return;
    }
    for (bir = 0; bir < DEVICE_MAX_BARS; ++bir) {
        free(state->bars[bir].data);
        free(state->bars[bir].reset);
    }
    free(state);
}


int behavior_static_create(const struct device_model *model,
                           struct device_behavior *out, char *err, size_t err_len)
{
    struct static_state *state;
    size_t index;

    if (model == NULL || out == NULL) {
        return fail(err, err_len, "model and behavior output are required");
    }
    state = calloc(1, sizeof(*state));
    if (state == NULL) {
        return fail(err, err_len, "allocate static behavior");
    }
    for (index = 0; index < model->bar_count; ++index) {
        const struct device_bar *source = &model->bars[index];
        struct static_bar *bar;

        if (source->bir >= DEVICE_MAX_BARS || source->size == 0 ||
            state->bars[source->bir].data != NULL) {
            static_destroy(state);
            return fail(err, err_len, "invalid BAR%u definition", source->bir);
        }
        bar = &state->bars[source->bir];
        if (source->size > SIZE_MAX) {
            static_destroy(state);
            return fail(err, err_len, "BAR%u is too large for this host", source->bir);
        }
        bar->data = calloc(1, (size_t)source->size);
        if (bar->data == NULL) {
            static_destroy(state);
            return fail(err, err_len, "allocate BAR%u", source->bir);
        }
        bar->size = (size_t)source->size;
        if (source->reset_image != NULL) {
            memcpy(bar->data, source->reset_image, bar->size);
            bar->reset = malloc(bar->size);
            if (bar->reset == NULL) {
                static_destroy(state);
                return fail(err, err_len, "allocate BAR%u reset snapshot", source->bir);
            }
            memcpy(bar->reset, source->reset_image, bar->size);
        }
    }
    *out = (struct device_behavior){
        .state = state,
        .bind_host = static_bind_host,
        .reset = static_reset,
        .read = static_read,
        .write = static_write,
        .destroy = static_destroy,
    };
    return 0;
}


int behavior_create(const struct device_model *model,
                    struct device_behavior *out, char *err, size_t err_len)
{
    if (model == NULL) {
        return fail(err, err_len, "device model is required");
    }
    if (model->class_code == 0x000000) {
        return behavior_static_create(model, out, err, err_len);
    }
    if (model->class_code == 0x010802) {
        return behavior_nvme_create(model, out, err, err_len);
    }
    if (model->class_code == 0x040300) {
        return behavior_hda_create(model, out, err, err_len);
    }
    if ((model->class_code & 0xffff00) == 0x030000) {
        return behavior_gpu_create(model, out, err, err_len);
    }
    if (model->class_code == 0x010601) {
        return behavior_ahci_create(model, out, err, err_len);
    }
    if (model->class_code == 0x0c0330) {
        return behavior_xhci_create(model, out, err, err_len);
    }
    if (model->class_code == 0x020000) {
        return behavior_ethernet_create(model, out, err, err_len);
    }
    return behavior_static_create(model, out, err, err_len);
}
