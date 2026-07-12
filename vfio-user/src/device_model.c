#include "device_model.h"

#include <errno.h>
#include <limits.h>
#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/stat.h>

#include <json-c/json.h>
#include <openssl/evp.h>


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


static int join_path(const char *dir, const char *name, char *out, size_t out_len,
                     char *err, size_t err_len)
{
    int written = snprintf(out, out_len, "%s/%s", dir, name);

    if (written < 0 || (size_t)written >= out_len) {
        return fail(err, err_len, "artifact path is too long");
    }
    return 0;
}


static int read_file(const char *path, unsigned char **data, size_t *len,
                     char *err, size_t err_len)
{
    FILE *file = fopen(path, "rb");
    long size;
    unsigned char *buffer;

    if (file == NULL) {
        return fail(err, err_len, "open %s: %s", path, strerror(errno));
    }
    if (fseek(file, 0, SEEK_END) != 0 || (size = ftell(file)) < 0 ||
        fseek(file, 0, SEEK_SET) != 0) {
        fclose(file);
        return fail(err, err_len, "measure %s", path);
    }
    buffer = malloc((size_t)size + 1);
    if (buffer == NULL) {
        fclose(file);
        return fail(err, err_len, "allocate %ld bytes", size);
    }
    if (fread(buffer, 1, (size_t)size, file) != (size_t)size) {
        free(buffer);
        fclose(file);
        return fail(err, err_len, "read %s", path);
    }
    fclose(file);
    buffer[size] = '\0';
    *data = buffer;
    *len = (size_t)size;
    return 0;
}


static int sha256_hex(const unsigned char *data, size_t len, char out[65],
                      char *err, size_t err_len)
{
    EVP_MD_CTX *context = EVP_MD_CTX_new();
    unsigned char digest[EVP_MAX_MD_SIZE];
    unsigned digest_len = 0;
    size_t index;

    if (context == NULL || EVP_DigestInit_ex(context, EVP_sha256(), NULL) != 1 ||
        EVP_DigestUpdate(context, data, len) != 1 ||
        EVP_DigestFinal_ex(context, digest, &digest_len) != 1) {
        EVP_MD_CTX_free(context);
        return fail(err, err_len, "calculate SHA-256");
    }
    EVP_MD_CTX_free(context);
    if (digest_len != 32) {
        return fail(err, err_len, "unexpected SHA-256 length %u", digest_len);
    }
    for (index = 0; index < digest_len; ++index) {
        snprintf(out + index * 2, 3, "%02x", digest[index]);
    }
    out[64] = '\0';
    return 0;
}


static json_object *required(json_object *parent, const char *name,
                             char *err, size_t err_len)
{
    json_object *value = NULL;

    if (!json_object_object_get_ex(parent, name, &value)) {
        fail(err, err_len, "missing JSON field %s", name);
        return NULL;
    }
    return value;
}


static int verify_device_model_artifact(const char *artifact_dir,
                                        unsigned char **model_data, size_t *model_len,
                                        char *err, size_t err_len)
{
    char manifest_path[PATH_MAX];
    char model_path[PATH_MAX];
    unsigned char *manifest_data = NULL;
    size_t manifest_len = 0;
    json_object *manifest = NULL;
    json_object *files;
    size_t index;
    int result = -1;

    if (join_path(artifact_dir, "build_manifest.json", manifest_path, sizeof(manifest_path), err, err_len) < 0 ||
        join_path(artifact_dir, "device_model.json", model_path, sizeof(model_path), err, err_len) < 0 ||
        read_file(manifest_path, &manifest_data, &manifest_len, err, err_len) < 0 ||
        read_file(model_path, model_data, model_len, err, err_len) < 0) {
        goto done;
    }
    manifest = json_tokener_parse((const char *)manifest_data);
    if (manifest == NULL || (files = required(manifest, "files", err, err_len)) == NULL ||
        !json_object_is_type(files, json_type_array)) {
        fail(err, err_len, "invalid build manifest");
        goto done;
    }
    for (index = 0; index < json_object_array_length(files); ++index) {
        json_object *entry = json_object_array_get_idx(files, index);
        json_object *name = required(entry, "name", err, err_len);
        json_object *size = required(entry, "size", err, err_len);
        json_object *hash = required(entry, "sha256", err, err_len);
        char actual_hash[65];

        if (name == NULL || size == NULL || hash == NULL) {
            goto done;
        }
        if (strcmp(json_object_get_string(name), "device_model.json") != 0) {
            continue;
        }
        if (json_object_get_int64(size) != (int64_t)*model_len) {
            fail(err, err_len, "device_model.json size differs from manifest");
            goto done;
        }
        if (sha256_hex(*model_data, *model_len, actual_hash, err, err_len) < 0) {
            goto done;
        }
        if (strcmp(actual_hash, json_object_get_string(hash)) != 0) {
            fail(err, err_len, "device_model.json SHA-256 differs from manifest");
            goto done;
        }
        result = 0;
        goto done;
    }
    fail(err, err_len, "device_model.json is absent from manifest");

done:
    json_object_put(manifest);
    free(manifest_data);
    if (result < 0) {
        free(*model_data);
        *model_data = NULL;
        *model_len = 0;
    }
    return result;
}


static int decode_config_image(const char *encoded, uint8_t *output, size_t output_size,
                               size_t expected_size, char *err, size_t err_len)
{
    size_t encoded_len = strlen(encoded);
    unsigned char *decoded;
    int decoded_len;

    if (encoded_len == 0 || encoded_len % 4 != 0 || expected_size > output_size) {
        return fail(err, err_len, "invalid config-space base64 length");
    }
    decoded = malloc(encoded_len);
    if (decoded == NULL) {
        return fail(err, err_len, "allocate config-space decode buffer");
    }
    decoded_len = EVP_DecodeBlock(decoded, (const unsigned char *)encoded, (int)encoded_len);
    if (decoded_len < 0) {
        free(decoded);
        return fail(err, err_len, "invalid config-space base64 data");
    }
    if (encoded_len >= 1 && encoded[encoded_len - 1] == '=') {
        --decoded_len;
    }
    if (encoded_len >= 2 && encoded[encoded_len - 2] == '=') {
        --decoded_len;
    }
    if ((size_t)decoded_len != expected_size) {
        free(decoded);
        return fail(err, err_len, "config-space image is %d bytes, expected %zu", decoded_len, expected_size);
    }
    memcpy(output, decoded, expected_size);
    free(decoded);
    return 0;
}


static int parse_function(json_object *root, struct device_model *model,
                          char *err, size_t err_len)
{
    json_object *functions = required(root, "functions", err, err_len);
    json_object *function;

    if (functions == NULL || !json_object_is_type(functions, json_type_array) ||
        json_object_array_length(functions) != 1) {
        return fail(err, err_len, "device model must contain exactly one PCI function");
    }
    function = json_object_array_get_idx(functions, 0);
    json_object *vendor_id = required(function, "vendor_id", err, err_len);
    json_object *device_id = required(function, "device_id", err, err_len);
    json_object *subsystem_vendor_id = required(function, "subsystem_vendor_id", err, err_len);
    json_object *subsystem_device_id = required(function, "subsystem_device_id", err, err_len);
    json_object *revision_id = required(function, "revision_id", err, err_len);
    json_object *class_code = required(function, "class_code", err, err_len);
    json_object *header_type = required(function, "header_type", err, err_len);

    if (vendor_id == NULL || device_id == NULL || subsystem_vendor_id == NULL ||
        subsystem_device_id == NULL || revision_id == NULL || class_code == NULL ||
        header_type == NULL) {
        return -1;
    }
    model->vendor_id = (uint16_t)json_object_get_int64(vendor_id);
    model->device_id = (uint16_t)json_object_get_int64(device_id);
    model->subsystem_vendor_id = (uint16_t)json_object_get_int64(subsystem_vendor_id);
    model->subsystem_device_id = (uint16_t)json_object_get_int64(subsystem_device_id);
    model->revision_id = (uint8_t)json_object_get_int64(revision_id);
    model->class_code = (uint32_t)json_object_get_int64(class_code);
    model->header_type = (uint8_t)json_object_get_int64(header_type);
    return 0;
}


static int parse_config(json_object *root, struct device_model *model,
                        char *err, size_t err_len)
{
    json_object *config = required(root, "config_space", err, err_len);
    json_object *size;
    json_object *image;

    if (config == NULL || (size = required(config, "size", err, err_len)) == NULL ||
        (image = required(config, "reset_image", err, err_len)) == NULL) {
        return -1;
    }
    model->config_space_size = (size_t)json_object_get_int64(size);
    return decode_config_image(json_object_get_string(image), model->config_space,
                               sizeof(model->config_space), model->config_space_size,
                               err, err_len);
}


static int parse_bars(json_object *root, struct device_model *model,
                      char *err, size_t err_len)
{
    json_object *bars = required(root, "bars", err, err_len);
    size_t index;

    if (bars == NULL || !json_object_is_type(bars, json_type_array) ||
        json_object_array_length(bars) > DEVICE_MAX_BARS) {
        return fail(err, err_len, "invalid BAR array");
    }
    model->bar_count = json_object_array_length(bars);
    for (index = 0; index < model->bar_count; ++index) {
        json_object *bar = json_object_array_get_idx(bars, index);
        json_object *bir = required(bar, "bir", err, err_len);
        json_object *type = required(bar, "type", err, err_len);
        json_object *size = required(bar, "size", err, err_len);
        json_object *prefetchable = required(bar, "prefetchable", err, err_len);
        json_object *width = required(bar, "address_width", err, err_len);
        json_object *reset_image = NULL;
        const char *type_name;

        if (bir == NULL || type == NULL || size == NULL || prefetchable == NULL || width == NULL) {
            return -1;
        }
        model->bars[index].bir = (unsigned)json_object_get_int(bir);
        model->bars[index].size = (uint64_t)json_object_get_int64(size);
        model->bars[index].prefetchable = json_object_get_boolean(prefetchable);
        model->bars[index].is_64bit = json_object_get_int(width) == 64;
        if (json_object_object_get_ex(bar, "reset_image", &reset_image) &&
            json_object_is_type(reset_image, json_type_string) &&
            json_object_get_string_len(reset_image) > 0) {
            if (model->bars[index].size > SIZE_MAX) {
                return fail(err, err_len, "BAR%u reset image is too large", model->bars[index].bir);
            }
            model->bars[index].reset_image = malloc((size_t)model->bars[index].size);
            if (model->bars[index].reset_image == NULL) {
                return fail(err, err_len, "allocate BAR%u reset image", model->bars[index].bir);
            }
            if (decode_config_image(json_object_get_string(reset_image),
                                    model->bars[index].reset_image,
                                    (size_t)model->bars[index].size,
                                    (size_t)model->bars[index].size,
                                    err, err_len) < 0) {
                return -1;
            }
        }
        type_name = json_object_get_string(type);
        if (strcmp(type_name, "io") == 0) {
            model->bars[index].type = DEVICE_BAR_IO;
        } else if (strcmp(type_name, "mem32") == 0 || strcmp(type_name, "mem64") == 0) {
            model->bars[index].type = DEVICE_BAR_MEMORY;
        } else {
            return fail(err, err_len, "unsupported BAR type %s", type_name);
        }
    }
    return 0;
}


static int parse_interrupts(json_object *root, struct device_model *model,
                            char *err, size_t err_len)
{
    json_object *interrupts = NULL;
    json_object *msix = NULL;
    size_t index;

    if (json_object_object_get_ex(root, "interrupts", &interrupts) &&
        json_object_is_type(interrupts, json_type_array)) {
        for (index = 0; index < json_object_array_length(interrupts); ++index) {
            json_object *entry = json_object_array_get_idx(interrupts, index);
            json_object *kind = required(entry, "kind", err, err_len);
            json_object *vectors = required(entry, "vectors", err, err_len);

            if (kind == NULL || vectors == NULL) {
                return -1;
            }
            if (strcmp(json_object_get_string(kind), "msi") == 0) {
                model->msi_vectors = (unsigned)json_object_get_int(vectors);
            } else if (strcmp(json_object_get_string(kind), "msix") == 0) {
                model->msix_vectors = (unsigned)json_object_get_int(vectors);
            }
        }
    }
    if (!json_object_object_get_ex(root, "msix", &msix) || json_object_is_type(msix, json_type_null)) {
        return 0;
    }
    json_object *table_size = required(msix, "table_size", err, err_len);
    json_object *table_bir = required(msix, "table_bir", err, err_len);
    json_object *table_offset = required(msix, "table_offset", err, err_len);
    json_object *pba_bir = required(msix, "pba_bir", err, err_len);
    json_object *pba_offset = required(msix, "pba_offset", err, err_len);

    if (table_size == NULL || table_bir == NULL || table_offset == NULL ||
        pba_bir == NULL || pba_offset == NULL) {
        return -1;
    }
    model->msix_vectors = (unsigned)json_object_get_int64(table_size);
    model->msix_table_bir = (unsigned)json_object_get_int64(table_bir);
    model->msix_table_offset = (uint64_t)json_object_get_int64(table_offset);
    model->msix_pba_bir = (unsigned)json_object_get_int64(pba_bir);
    model->msix_pba_offset = (uint64_t)json_object_get_int64(pba_offset);
    return 0;
}


static const struct device_bar *find_bar(const struct device_model *model, unsigned bir)
{
    size_t index;

    for (index = 0; index < model->bar_count; ++index) {
        if (model->bars[index].bir == bir) {
            return &model->bars[index];
        }
    }
    return NULL;
}


static int validate_standard_capabilities(const struct device_model *model,
                                          char *err, size_t err_len)
{
    bool seen[256] = {false};
    unsigned offset;

    if ((model->config_space[0x06] & 0x10) == 0) {
        return 0;
    }
    offset = model->config_space[0x34] & 0xfc;
    while (offset != 0) {
        if (offset < 0x40 || offset + 1 >= model->config_space_size ||
            (offset & 0x03) != 0 || seen[offset]) {
            return fail(err, err_len, "standard capability chain is invalid at 0x%x", offset);
        }
        seen[offset] = true;
        offset = model->config_space[offset + 1] & 0xfc;
    }
    return 0;
}


static int validate_extended_capabilities(const struct device_model *model,
                                          char *err, size_t err_len)
{
    bool seen[DEVICE_CONFIG_SPACE_SIZE / 4] = {false};
    unsigned offset = 0x100;

    if (model->config_space_size < DEVICE_CONFIG_SPACE_SIZE) {
        return 0;
    }
    while (offset != 0) {
        uint32_t header;
        unsigned next;

        if (offset < 0x100 || offset + 3 >= model->config_space_size ||
            (offset & 0x03) != 0 || seen[offset / 4]) {
            return fail(err, err_len, "extended capability chain is invalid at 0x%x", offset);
        }
        seen[offset / 4] = true;
        memcpy(&header, model->config_space + offset, sizeof(header));
        if (header == 0) {
            return 0;
        }
        next = (header >> 20) & 0xfff;
        if (next != 0 && next <= offset) {
            return fail(err, err_len, "extended capability chain does not advance at 0x%x", offset);
        }
        offset = next;
    }
    return 0;
}


int device_model_validate(const struct device_model *model, char *err, size_t err_len)
{
    bool used[DEVICE_MAX_BARS] = {false};
    size_t index;

    if (model == NULL) {
        return fail(err, err_len, "device model is null");
    }
    if (model->config_space_size != 256 && model->config_space_size != 4096) {
        return fail(err, err_len, "config-space size must be 256 or 4096 bytes");
    }
    if (validate_standard_capabilities(model, err, err_len) < 0 ||
        validate_extended_capabilities(model, err, err_len) < 0) {
        return -1;
    }
    for (index = 0; index < model->bar_count; ++index) {
        const struct device_bar *bar = &model->bars[index];

        if (bar->bir >= DEVICE_MAX_BARS || used[bar->bir]) {
            return fail(err, err_len, "BAR index %u is invalid or duplicated", bar->bir);
        }
        if (bar->size < 4 || (bar->size & (bar->size - 1)) != 0) {
            return fail(err, err_len, "BAR%u size must be a power of two", bar->bir);
        }
        if (bar->is_64bit && bar->bir == DEVICE_MAX_BARS - 1) {
            return fail(err, err_len, "64-bit BAR cannot start at BAR5");
        }
        used[bar->bir] = true;
        if (bar->is_64bit) {
            if (used[bar->bir + 1]) {
                return fail(err, err_len, "64-bit BAR%u overlaps another BAR", bar->bir);
            }
            used[bar->bir + 1] = true;
        }
    }
    if (model->msix_vectors > 0) {
        const struct device_bar *table = find_bar(model, model->msix_table_bir);
        const struct device_bar *pba = find_bar(model, model->msix_pba_bir);
        uint64_t table_size = (uint64_t)model->msix_vectors * 16;
        uint64_t pba_size = ((uint64_t)model->msix_vectors + 63) / 64 * 8;

        if (table == NULL || model->msix_table_offset > table->size ||
            table_size > table->size - model->msix_table_offset) {
            return fail(err, err_len, "MSI-X table is outside its BAR");
        }
        if (pba == NULL || model->msix_pba_offset > pba->size ||
            pba_size > pba->size - model->msix_pba_offset) {
            return fail(err, err_len, "MSI-X PBA is outside its BAR");
        }
    }
    return 0;
}


int device_model_load(const char *artifact_dir, struct device_model **out,
                      char *err, size_t err_len)
{
    unsigned char *model_data = NULL;
    size_t model_len = 0;
    json_object *root = NULL;
    json_object *schema;
    struct device_model *model = NULL;
    int result = -1;

    if (artifact_dir == NULL || out == NULL) {
        return fail(err, err_len, "artifact directory and output are required");
    }
    *out = NULL;
    if (verify_device_model_artifact(artifact_dir, &model_data, &model_len, err, err_len) < 0) {
        return -1;
    }
    root = json_tokener_parse((const char *)model_data);
    if (root == NULL || (schema = required(root, "schema_version", err, err_len)) == NULL ||
        json_object_get_int(schema) != 1) {
        fail(err, err_len, "unsupported or missing device-model schema");
        goto done;
    }
    model = calloc(1, sizeof(*model));
    if (model == NULL) {
        fail(err, err_len, "allocate device model");
        goto done;
    }
    if (parse_function(root, model, err, err_len) < 0 ||
        parse_config(root, model, err, err_len) < 0 ||
        parse_bars(root, model, err, err_len) < 0 ||
        parse_interrupts(root, model, err, err_len) < 0 ||
        device_model_validate(model, err, err_len) < 0) {
        goto done;
    }
    *out = model;
    model = NULL;
    result = 0;

done:
    device_model_free(model);
    json_object_put(root);
    free(model_data);
    return result;
}


void device_model_free(struct device_model *model)
{
    size_t index;

    if (model == NULL) {
        return;
    }
    for (index = 0; index < model->bar_count; ++index) {
        free(model->bars[index].reset_image);
    }
    free(model);
}
