#ifndef PCILEECH_VFIO_DEVICE_MODEL_H
#define PCILEECH_VFIO_DEVICE_MODEL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>


#define DEVICE_CONFIG_SPACE_SIZE 4096
#define DEVICE_MAX_BARS 6

enum device_bar_type {
    DEVICE_BAR_MEMORY,
    DEVICE_BAR_IO,
};

struct device_bar {
    unsigned bir;
    enum device_bar_type type;
    uint64_t size;
    bool prefetchable;
    bool is_64bit;
    uint8_t *reset_image;
};

struct device_model {
    uint16_t vendor_id;
    uint16_t device_id;
    uint16_t subsystem_vendor_id;
    uint16_t subsystem_device_id;
    uint8_t revision_id;
    uint32_t class_code;
    uint8_t header_type;
    uint8_t config_space[DEVICE_CONFIG_SPACE_SIZE];
    size_t config_space_size;
    struct device_bar bars[DEVICE_MAX_BARS];
    size_t bar_count;
    unsigned msi_vectors;
    unsigned msix_vectors;
    unsigned msix_table_bir;
    uint64_t msix_table_offset;
    unsigned msix_pba_bir;
    uint64_t msix_pba_offset;
};

int device_model_load(const char *artifact_dir, struct device_model **out,
                      char *err, size_t err_len);
int device_model_validate(const struct device_model *model, char *err, size_t err_len);
void device_model_free(struct device_model *model);

#endif
