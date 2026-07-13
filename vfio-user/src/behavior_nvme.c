#include "device_behavior.h"

#include <errno.h>
#include <stdbool.h>
#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define NVME_PAGE_SIZE 4096u
#define NVME_MDTS_BYTES (32u * 1024u)
#define NVME_NAMESPACE_LBAS (1u << 20)
#define NVME_LBA_BYTES 512u
#define NVME_MAX_PRP_ENTRIES 64u
#define NVME_CACHE_PAGES 1024u


struct nvme_state {
    struct device_behavior registers;
    const struct device_model *model;
    struct behavior_host_ops host;
    uint32_t aqa;
    uint64_t asq;
    uint64_t acq;
    uint16_t sq_head;
    uint16_t cq_tail;
    uint16_t cq_head;
    uint8_t cq_phase;
    uint8_t cq_head_phase;
    uint16_t pending_tail;
    uint16_t io_pending_tail;
    bool admin_pending;
    bool io_pending;
    uint64_t io_sq;
    uint64_t io_cq;
    uint16_t io_sq_head;
    uint16_t io_cq_tail;
    uint16_t io_cq_head;
    uint8_t io_cq_phase;
    uint8_t io_cq_head_phase;
    uint16_t io_sq_size;
    uint16_t io_cq_size;
    bool io_sq_created;
    bool io_cq_created;
    unsigned io_vector;
    bool msix_masked;
    bool msix_pending;
    struct {
        bool valid;
        uint64_t tag;
        uint8_t data[NVME_PAGE_SIZE];
    } namespace_cache[NVME_CACHE_PAGES];
    uint64_t read_commands;
    uint64_t write_commands;
    uint64_t data_units_read;
    uint64_t data_units_written;
    uint64_t flush_cmds;
    uint64_t dataset_cmds;
    uint64_t write_zero_cmds;
    uint64_t format_cmds;
    uint32_t error_log_entries;
    uint32_t unsafe_shutdowns;
    uint32_t power_cycles;
    uint32_t power_on_hours;
    uint32_t feat_arbitration;
    uint32_t feat_power_mgmt;
    uint32_t feat_temp_threshold;
    uint32_t feat_write_cache;
    uint32_t feat_irq_coalescing;
    uint32_t feat_async_event_cfg;
    uint64_t dma_mrd_tlps;
    uint64_t dma_mwr_tlps;
    uint64_t prp_list_fetches;
    uint64_t queue_resets;
    uint64_t shutdowns;
    uint64_t timeout_errors;
    uint64_t cpl_errors;
    uint32_t invalid_cmds;
    uint64_t transport_errors;
    uint64_t aer_events;
};

struct nvme_sqe {
    uint8_t opcode;
    uint8_t flags;
    uint16_t cid;
    uint32_t nsid;
    uint64_t reserved;
    uint64_t mptr;
    uint64_t prp1;
    uint64_t prp2;
    uint32_t cdw10;
    uint32_t cdw11;
    uint32_t cdw12;
    uint32_t cdw13;
    uint32_t cdw14;
    uint32_t cdw15;
} __attribute__((packed));

struct nvme_cqe {
    uint32_t result;
    uint32_t reserved;
    uint16_t sq_head;
    uint16_t sq_id;
    uint16_t cid;
    uint16_t status;
} __attribute__((packed));


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


static int set_csts(struct nvme_state *state, uint32_t value)
{
    ssize_t written = state->registers.write(
        state->registers.state, 0, 0x1c, &value, sizeof(value));

    return written == (ssize_t)sizeof(value) ? 0 : -EIO;
}


static int nvme_bind_host(void *opaque, const struct behavior_host_ops *ops)
{
    struct nvme_state *state = opaque;

    if (ops == NULL || ops->dma_read == NULL || ops->dma_write == NULL || ops->irq == NULL) {
        return -EINVAL;
    }
    state->host = *ops;
    return 0;
}


static void put16(uint8_t *data, size_t offset, uint16_t value)
{
    data[offset] = (uint8_t)value;
    data[offset + 1] = (uint8_t)(value >> 8);
}


static void put32(uint8_t *data, size_t offset, uint32_t value)
{
    put16(data, offset, (uint16_t)value);
    put16(data, offset + 2, (uint16_t)(value >> 16));
}


static void put64(uint8_t *data, size_t offset, uint64_t value)
{
    put32(data, offset, (uint32_t)value);
    put32(data, offset + 4, (uint32_t)(value >> 32));
}

static unsigned namespace_slot(uint64_t page)
{
    return (unsigned)((page * 11400714819323198485ull) % NVME_CACHE_PAGES);
}

static uint8_t *namespace_page(struct nvme_state *state, uint64_t page, bool allocate)
{
    unsigned slot = namespace_slot(page);
    if (!state->namespace_cache[slot].valid || state->namespace_cache[slot].tag != page) {
        if (!allocate) {
            return NULL;
        }
        state->namespace_cache[slot].valid = true;
        state->namespace_cache[slot].tag = page;
        memset(state->namespace_cache[slot].data, 0, NVME_PAGE_SIZE);
    }
    return state->namespace_cache[slot].data;
}

static void namespace_read(struct nvme_state *state, uint64_t offset,
                           uint8_t *data, size_t length)
{
    while (length > 0) {
        uint64_t page = offset / NVME_PAGE_SIZE;
        size_t page_offset = (size_t)(offset % NVME_PAGE_SIZE);
        size_t chunk = NVME_PAGE_SIZE - page_offset;
        uint8_t *source = namespace_page(state, page, false);

        if (chunk > length) {
            chunk = length;
        }
        if (source == NULL) {
            memset(data, 0, chunk);
        } else {
            memcpy(data, source + page_offset, chunk);
        }
        offset += chunk;
        data += chunk;
        length -= chunk;
    }
}

static void namespace_write(struct nvme_state *state, uint64_t offset,
                            const uint8_t *data, size_t length)
{
    while (length > 0) {
        uint64_t page = offset / NVME_PAGE_SIZE;
        size_t page_offset = (size_t)(offset % NVME_PAGE_SIZE);
        size_t chunk = NVME_PAGE_SIZE - page_offset;
        uint8_t *target = namespace_page(state, page, true);

        if (chunk > length) {
            chunk = length;
        }
        memcpy(target + page_offset, data, chunk);
        offset += chunk;
        data += chunk;
        length -= chunk;
    }
}

static void namespace_zero(struct nvme_state *state, uint64_t offset, size_t length)
{
    while (length > 0) {
        uint64_t page = offset / NVME_PAGE_SIZE;
        size_t page_offset = (size_t)(offset % NVME_PAGE_SIZE);
        size_t chunk = NVME_PAGE_SIZE - page_offset;
        uint8_t *target;

        if (chunk > length) {
            chunk = length;
        }
        target = namespace_page(state, page, false);
        if (target != NULL) {
            memset(target + page_offset, 0, chunk);
        }
        offset += chunk;
        length -= chunk;
    }
}

static void namespace_reset(struct nvme_state *state)
{
    for (unsigned index = 0; index < NVME_CACHE_PAGES; ++index) {
        state->namespace_cache[index].valid = false;
    }
}


static void identify_controller(const struct nvme_state *state, uint8_t data[4096])
{
    char model[40];

    if (state->model->has_nvme_identify) {
        memcpy(data, state->model->nvme_controller_ident, 4096);
        put16(data, 0x00, state->model->vendor_id);
        put16(data, 0x02, state->model->subsystem_vendor_id);
        data[0x4d] = 3;
        put32(data, 0x50, 0x00010400);
        return;
    }
    memset(data, 0, 4096);
    put16(data, 0x00, state->model->vendor_id);
    put16(data, 0x02, state->model->subsystem_vendor_id);
    memcpy(data + 0x04, "NVME0000000000000000", 20);
    memset(model, ' ', sizeof(model));
    snprintf(model, sizeof(model), "PCILeechGen NVMe %04X:%04X",
             state->model->vendor_id, state->model->device_id);
    memcpy(data + 0x18, model, sizeof(model));
    memcpy(data + 0x40, "1.0     ", 8);
    memcpy(data + 0x300, "nqn.2014.08.org.nvmexpress:pcileechgen", 38);
    data[0x4d] = 3;
    put32(data, 0x50, 0x00010400);
    data[0x6f] = 1;
    data[0x103] = 3;
    data[0x200] = 0x66;
    data[0x201] = 0x44;
    put32(data, 0x204, 1);
    put16(data, 0x208, 0x000c);
}


static void identify_namespace(const struct nvme_state *state, uint8_t data[4096])
{
    if (state->model->has_nvme_identify) {
        memcpy(data, state->model->nvme_namespace_ident, 4096);
        return;
    }
    memset(data, 0, 4096);
    put64(data, 0x00, 0x100000);
    put64(data, 0x08, 0x100000);
    put64(data, 0x10, 0x100000);
    data[0x19] = 0;
    data[0x1a] = 0;
    data[0xc0 + 2] = 9;
}

static int validate_transfer(const struct nvme_sqe *sqe, size_t length)
{
    if (length == 0 || length > NVME_MDTS_BYTES || sqe->prp1 == 0 ||
        (length > NVME_PAGE_SIZE && sqe->prp2 == 0)) {
        return -EINVAL;
    }
    return 0;
}

static int dma_prp(struct nvme_state *state, const struct nvme_sqe *sqe,
                   uint8_t *data, size_t length, bool device_to_host)
{
    size_t remaining = length;
    size_t offset = (size_t)(sqe->prp1 & (NVME_PAGE_SIZE - 1));
    uint64_t address = sqe->prp1;
    uint64_t list_address = sqe->prp2;
    unsigned list_index = 0;

    if (validate_transfer(sqe, length) < 0) {
        return -EINVAL;
    }
    while (remaining > 0) {
        size_t chunk = NVME_PAGE_SIZE - offset;

        if (chunk > remaining) {
            chunk = remaining;
        }
        if (device_to_host) {
            state->dma_mwr_tlps++;
            if (state->host.dma_write(state->host.opaque, address, data, chunk) < 0) {
                state->transport_errors++;
                return -EIO;
            }
        } else {
            state->dma_mrd_tlps++;
            if (state->host.dma_read(state->host.opaque, address, data, chunk) < 0) {
                state->transport_errors++;
                return -EIO;
            }
        }
        data += chunk;
        remaining -= chunk;
        offset = 0;
        if (remaining == 0) {
            break;
        }
        if (address == sqe->prp1) {
            if (remaining <= NVME_PAGE_SIZE) {
                address = sqe->prp2;
            } else {
                if ((sqe->prp2 & (NVME_PAGE_SIZE - 1)) != 0) {
                    return -EINVAL;
                }
                list_address = sqe->prp2;
                state->prp_list_fetches++;
                if (state->host.dma_read(state->host.opaque, list_address,
                                         &address, sizeof(address)) < 0) {
                    state->transport_errors++;
                    return -EIO;
                }
                list_index = 1;
                if (address == 0 || (address & (NVME_PAGE_SIZE - 1)) != 0) {
                    return -EINVAL;
                }
            }
        } else {
            if (++list_index >= NVME_MAX_PRP_ENTRIES) {
                return -E2BIG;
            }
            if (state->host.dma_read(state->host.opaque,
                                     list_address + (uint64_t)list_index * sizeof(address),
                                     &address, sizeof(address)) < 0) {
                state->transport_errors++;
                return -EIO;
            }
            state->prp_list_fetches++;
            if (address == 0 || (address & (NVME_PAGE_SIZE - 1)) != 0) {
                return -EINVAL;
            }
        }
    }
    return 0;
}


static int write_prps(struct nvme_state *state, const struct nvme_sqe *sqe,
                      uint8_t *data, size_t length)
{
    return dma_prp(state, sqe, data, length, true);
}

static int read_prps(struct nvme_state *state, const struct nvme_sqe *sqe,
                     uint8_t *data, size_t length)
{
    return dma_prp(state, sqe, data, length, false);
}

static void fill_log_page(const struct nvme_state *state, uint8_t page,
                          uint8_t data[4096])
{
    memset(data, 0, 4096);
    switch (page) {
    case 0x00:
        data[0] = 0x01; data[1] = 0x01; data[2] = 0x01; data[3] = 0x01;
        data[48 * 4] = 0x01;
        break;
    case 0x01:
        put32(data, 0, state->error_log_entries);
        put16(data, 6, 0x0002);
        break;
    case 0x02:
        data[0] = 100;
        data[1] = 10;
        put64(data, 32, state->data_units_read);
        put64(data, 48, state->data_units_written);
        put64(data, 64, state->read_commands);
        put64(data, 80, state->write_commands);
        put32(data, 112, state->power_cycles);
        put32(data, 128, state->power_on_hours);
        put32(data, 144, state->unsafe_shutdowns);
        put32(data, 176, state->error_log_entries);
        break;
    case 0x03:
        put32(data, 0, 1);
        memcpy(data + 8, "1.0     ", 8);
        break;
    case 0xc0:
        put32(data, 0, 0x444d564e); /* NVMD */
        put32(data, 4, 0x00000002);
        put32(data, 8, NVME_NAMESPACE_LBAS);
        put32(data, 12, 0x00002000); /* 32 KiB MDTS */
        put32(data, 32, state->dma_mrd_tlps);
        put32(data, 36, state->dma_mrd_tlps >> 32);
        put32(data, 40, state->dma_mwr_tlps);
        put32(data, 44, state->dma_mwr_tlps >> 32);
        put32(data, 48, state->prp_list_fetches);
        put32(data, 52, state->prp_list_fetches >> 32);
        put32(data, 56, state->queue_resets);
        put32(data, 60, state->queue_resets >> 32);
        put32(data, 64, state->shutdowns);
        put32(data, 68, state->shutdowns >> 32);
        put32(data, 72, state->flush_cmds);
        put32(data, 76, state->flush_cmds >> 32);
        put32(data, 80, state->dataset_cmds);
        put32(data, 84, state->dataset_cmds >> 32);
        put32(data, 88, state->write_zero_cmds);
        put32(data, 92, state->write_zero_cmds >> 32);
        put32(data, 96, state->format_cmds);
        put32(data, 100, state->format_cmds >> 32);
        put32(data, 112, state->timeout_errors);
        put32(data, 116, state->timeout_errors >> 32);
        put32(data, 120, state->cpl_errors);
        put32(data, 124, state->cpl_errors >> 32);
        put32(data, 128, state->invalid_cmds);
        put32(data, 132, state->error_log_entries);
        put32(data, 136, 0);
        put32(data, 140, state->unsafe_shutdowns);
        put32(data, 144, 0);
        put32(data, 148, state->transport_errors);
        put32(data, 152, state->transport_errors >> 32);
        put32(data, 156, 0);
        put32(data, 160, state->msix_pending ? 1u : 0u);
        put32(data, 164, state->io_vector);
        put32(data, 168, 1);
        put32(data, 172, state->feat_async_event_cfg & 0xff);
        put32(data, 176, state->power_cycles);
        put32(data, 180, state->feat_write_cache & 1u);
        break;
    default:
        break;
    }
}

static size_t log_length(const struct nvme_sqe *sqe)
{
    uint64_t dwords = ((uint64_t)(sqe->cdw11) << 16) | (sqe->cdw10 >> 16);

    dwords += 1;
    if (dwords > 1024) {
        dwords = 1024;
    }
    return (size_t)dwords * sizeof(uint32_t);
}


static int process_admin(struct nvme_state *state, uint16_t tail)
{
    uint16_t sq_size = (uint16_t)((state->aqa & 0xfff) + 1);
    uint16_t cq_size = (uint16_t)(((state->aqa >> 16) & 0xfff) + 1);

    if (state->host.dma_read == NULL || state->asq == 0 || state->acq == 0 ||
        sq_size < 2 || cq_size < 2 || tail >= sq_size) {
        return -EINVAL;
    }
    while (state->sq_head != tail) {
        if (state->cq_tail == state->cq_head &&
            state->cq_phase != state->cq_head_phase) {
            state->admin_pending = true;
            return 0;
        }
        struct nvme_sqe sqe;
        struct nvme_cqe cqe = {0};
        uint16_t status = 0;
        uint32_t result = 0;
        uint8_t data[4096];

        if (state->host.dma_read(state->host.opaque,
                                 state->asq + (uint64_t)state->sq_head * sizeof(sqe),
                                 &sqe, sizeof(sqe)) < 0) {
            return -EIO;
        }
        if (sqe.opcode == 0x06) {
            unsigned cns = sqe.cdw10 & 0xff;

            if (cns == 1) {
                identify_controller(state, data);
            } else if (cns == 0 && sqe.nsid == 1) {
                identify_namespace(state, data);
            } else if (cns == 2) {
                memset(data, 0, sizeof(data));
                put32(data, 0, 1);
            } else {
                status = 0xb;
                memset(data, 0, sizeof(data));
            }
            if (status == 0 && write_prps(state, &sqe, data, sizeof(data)) < 0) {
                return -EIO;
            }
        } else if (sqe.opcode == 0x02) {
            uint8_t page = (uint8_t)(sqe.cdw10 & 0xff);
            size_t length = log_length(&sqe);

            if (page != 0x00 && page != 0x01 && page != 0x02 &&
                page != 0x03 && page != 0xc0) {
                status = 2;
            } else if (sqe.nsid != 0 && sqe.nsid != 1 && sqe.nsid != 0xffffffffu) {
                status = 0xb;
            } else {
                fill_log_page(state, page, data);
                if (write_prps(state, &sqe, data, length) < 0) {
                    return -EIO;
                }
            }
        } else if (sqe.opcode == 0x01 || sqe.opcode == 0x05) {
            uint16_t qid = (uint16_t)(sqe.cdw10 & 0xffff);
            uint16_t qsize = (uint16_t)(sqe.cdw10 >> 16) + 1;

            if (qid != 1 || qsize < 2 || qsize > 64 ||
                sqe.prp1 == 0 || (sqe.prp1 & (NVME_PAGE_SIZE - 1)) != 0) {
                status = 1;
            } else if (sqe.opcode == 0x01) {
                if ((sqe.cdw11 >> 16) >= state->model->msix_vectors &&
                    state->model->msix_vectors != 0) {
                    status = 0x08;
                } else if (state->io_cq_created) {
                    status = 0x18;
                } else {
                state->io_cq = sqe.prp1;
                state->io_cq_size = qsize;
                state->io_cq_head = 0;
                state->io_cq_tail = 0;
                state->io_cq_phase = 1;
                state->io_cq_head_phase = 1;
                state->io_vector = (unsigned)(sqe.cdw11 >> 16);
                state->io_cq_created = true;
                }
            } else if (!state->io_cq_created) {
                status = 1;
            } else if (state->io_sq_created || qsize != state->io_cq_size) {
                status = 0x18;
            } else {
                state->io_sq = sqe.prp1;
                state->io_sq_size = qsize;
                state->io_sq_head = 0;
                state->io_sq_created = true;
            }
        } else if (sqe.opcode == 0x00) {
            if (sqe.cdw10 != 1 || !state->io_sq_created) {
                status = 1;
            } else {
                state->io_sq_created = false;
            }
        } else if (sqe.opcode == 0x04 || sqe.opcode == 0x0c) {
            if (sqe.cdw10 != 1 && sqe.opcode == 0x04) {
                status = 1;
            } else if (sqe.opcode == 0x04) {
                if (state->io_sq_created) {
                    status = 0x19;
                } else {
                    state->io_cq_created = false;
                    state->io_sq_created = false;
                }
            } else if (state->admin_pending) {
                status = 2;
            }
        } else if (sqe.opcode == 0x09 || sqe.opcode == 0x0a) {
            uint8_t feature = (uint8_t)(sqe.cdw10 & 0xff);

            if (feature != 0x01 && feature != 0x02 && feature != 0x04 &&
                feature != 0x06 && feature != 0x07 && feature != 0x08 &&
                feature != 0x09 && feature != 0x0a && feature != 0x0b) {
                status = 2;
            } else if (sqe.opcode == 0x09) {
                switch (feature) {
                case 0x01: state->feat_arbitration = sqe.cdw11; break;
                case 0x02: state->feat_power_mgmt = sqe.cdw11; break;
                case 0x04: state->feat_temp_threshold = sqe.cdw11; break;
                case 0x06: state->feat_write_cache = sqe.cdw11; break;
                case 0x08: state->feat_irq_coalescing = sqe.cdw11; break;
                case 0x0b: state->feat_async_event_cfg = sqe.cdw11; break;
                case 0x07: break;
                default: break;
                }
            } else {
                switch (feature) {
                case 0x01: result = state->feat_arbitration; break;
                case 0x02: result = state->feat_power_mgmt; break;
                case 0x04: result = state->feat_temp_threshold; break;
                case 0x06: result = state->feat_write_cache; break;
                case 0x07: result = state->io_sq_created ? 0x00010001u : 0; break;
                case 0x08: result = state->feat_irq_coalescing; break;
                case 0x0b: result = state->feat_async_event_cfg; break;
                default: result = 0; break;
                }
            }
        } else if (sqe.opcode == 0x80) {
            if (sqe.nsid != 1 || sqe.cdw10 != 0) {
                status = 0x20a;
            } else {
                state->format_cmds++;
                namespace_reset(state);
            }
        } else if (sqe.opcode == 0x00 || sqe.opcode == 0x01 ||
                   sqe.opcode == 0x02 || sqe.opcode == 0x08 || sqe.opcode == 0x09) {
            status = 1;
        }
        if (status != 0) {
            state->error_log_entries++;
            if (status == 1 || status == 2) {
                state->invalid_cmds++;
            }
        }
        state->sq_head = (uint16_t)((state->sq_head + 1) % sq_size);
        cqe.result = result;
        cqe.sq_head = state->sq_head;
        cqe.cid = sqe.cid;
        cqe.status = (uint16_t)((status << 1) | state->cq_phase);
        if (state->host.dma_write(state->host.opaque,
                                  state->acq + (uint64_t)state->cq_tail * sizeof(cqe),
                                  &cqe, sizeof(cqe)) < 0) {
            return -EIO;
        }
        state->cq_tail = (uint16_t)((state->cq_tail + 1) % cq_size);
        if (state->cq_tail == 0) {
            state->cq_phase ^= 1;
        }
        if (state->host.irq(state->host.opaque, 0) < 0) {
            return -EIO;
        }
    }
    return 0;
}

static int post_io_cqe(struct nvme_state *state, const struct nvme_sqe *sqe,
                       uint16_t status)
{
    struct nvme_cqe cqe = {0};

    if (!state->io_cq_created || state->io_cq_size < 2 ||
        (state->io_cq_tail == state->io_cq_head &&
         state->io_cq_phase != state->io_cq_head_phase)) {
        return -EAGAIN;
    }
    cqe.sq_head = state->io_sq_head;
    cqe.sq_id = 1;
    cqe.cid = sqe->cid;
    cqe.status = (uint16_t)((status << 1) | state->io_cq_phase);
    if (state->host.dma_write(state->host.opaque,
                              state->io_cq + (uint64_t)state->io_cq_tail * sizeof(cqe),
                              &cqe, sizeof(cqe)) < 0) {
        state->cpl_errors++;
        return -EIO;
    }
    state->io_cq_tail = (uint16_t)((state->io_cq_tail + 1) % state->io_cq_size);
    if (state->io_cq_tail == 0) {
        state->io_cq_phase ^= 1;
    }
    if (state->msix_masked) {
        state->msix_pending = true;
        return 0;
    }
    return state->host.irq(state->host.opaque, state->io_vector);
}

static int process_io(struct nvme_state *state, uint16_t tail)
{
    if (!state->io_sq_created || !state->io_cq_created || tail >= state->io_sq_size) {
        return -EINVAL;
    }
    while (state->io_sq_head != tail) {
        struct nvme_sqe sqe;
        uint16_t status = 0;

        if (state->host.dma_read(state->host.opaque,
                                 state->io_sq + (uint64_t)state->io_sq_head * sizeof(sqe),
                                 &sqe, sizeof(sqe)) < 0) {
            return -EIO;
        }
        if (sqe.nsid != 1) {
            status = 0xb;
        } else if (sqe.opcode == 0x00) {
            if (sqe.prp1 || sqe.prp2 || sqe.cdw10 || sqe.cdw11 || sqe.cdw12 ||
                sqe.cdw13 || sqe.cdw14 || sqe.cdw15) {
                status = 2;
            } else {
                state->flush_cmds++;
            }
        } else if (sqe.opcode == 0x01 || sqe.opcode == 0x02 || sqe.opcode == 0x08) {
            uint64_t lba = (uint64_t)sqe.cdw10 | ((uint64_t)sqe.cdw11 << 32);
            size_t length = ((size_t)(sqe.cdw12 & 0xffff) + 1) * NVME_LBA_BYTES;
            uint64_t byte_offset = lba * NVME_LBA_BYTES;
            uint8_t *buffer = NULL;

            if (lba >= NVME_NAMESPACE_LBAS || length > NVME_MDTS_BYTES ||
                lba + (length / NVME_LBA_BYTES) > NVME_NAMESPACE_LBAS ||
                byte_offset > NVME_NAMESPACE_LBAS * NVME_LBA_BYTES - length) {
                status = 0x80;
            } else if (sqe.opcode == 0x08 &&
                       (sqe.prp1 || sqe.prp2 || sqe.cdw13 || sqe.cdw14 || sqe.cdw15)) {
                status = 2;
            } else if (sqe.opcode != 0x08 && validate_transfer(&sqe, length) < 0) {
                status = 0x13;
            } else {
                buffer = malloc(length);
                if (buffer == NULL) {
                    return -ENOMEM;
                }
                if (sqe.opcode == 0x01) {
                    if (read_prps(state, &sqe, buffer, length) < 0) {
                        status = 0x04;
                    } else {
                        namespace_write(state, byte_offset, buffer, length);
                        state->write_commands++;
                        state->data_units_written += (length + 511) / 512;
                    }
                } else if (sqe.opcode == 0x02) {
                    namespace_read(state, byte_offset, buffer, length);
                    if (write_prps(state, &sqe, buffer, length) < 0) {
                        status = 0x04;
                    } else {
                        state->read_commands++;
                        state->data_units_read += (length + 511) / 512;
                    }
                } else {
                    state->write_zero_cmds++;
                    namespace_zero(state, byte_offset, length);
                }
                free(buffer);
            }
        } else if (sqe.opcode == 0x09) {
            unsigned count = (sqe.cdw10 & 0xff) + 1;
            size_t range_length = (size_t)count * 16;
            uint8_t *ranges = malloc(NVME_MDTS_BYTES);
            if (ranges == NULL) {
                return -ENOMEM;
            }
            if (read_prps(state, &sqe, ranges, range_length) < 0) {
                status = 0x13;
            } else {
                state->dataset_cmds++;
                if (count > NVME_MDTS_BYTES / 16) {
                    count = NVME_MDTS_BYTES / 16;
                }
                for (unsigned index = 0; index < count; ++index) {
                    uint64_t lba = (uint64_t)ranges[index * 16] |
                        ((uint64_t)ranges[index * 16 + 1] << 8) |
                        ((uint64_t)ranges[index * 16 + 2] << 16) |
                        ((uint64_t)ranges[index * 16 + 3] << 24) |
                        ((uint64_t)ranges[index * 16 + 4] << 32) |
                        ((uint64_t)ranges[index * 16 + 5] << 40);
                    uint32_t nlb = (uint32_t)ranges[index * 16 + 8] |
                        ((uint32_t)ranges[index * 16 + 9] << 8) |
                        ((uint32_t)ranges[index * 16 + 10] << 16) |
                        ((uint32_t)ranges[index * 16 + 11] << 24);
                    if (lba < NVME_NAMESPACE_LBAS && nlb <= NVME_NAMESPACE_LBAS - lba) {
                        namespace_zero(state, lba * NVME_LBA_BYTES,
                                       (size_t)nlb * NVME_LBA_BYTES);
                    }
                }
            }
            free(ranges);
        } else {
            status = 1;
        }
        if (post_io_cqe(state, &sqe, status) < 0) {
            state->io_pending = true;
            return 0;
        }
        state->io_sq_head = (uint16_t)((state->io_sq_head + 1) % state->io_sq_size);
    }
    return 0;
}


static int nvme_service(void *opaque)
{
    struct nvme_state *state = opaque;

    int result = 0;

    if (state->admin_pending) {
        state->admin_pending = false;
        result = process_admin(state, state->pending_tail);
    }
    if (result == 0 && state->io_pending) {
        state->io_pending = false;
        result = process_io(state, state->io_pending_tail);
    }
    return result;
}


static int nvme_reset(void *opaque)
{
    struct nvme_state *state = opaque;
    int result = state->registers.reset(state->registers.state);

    if (result < 0) {
        return result;
    }
    state->aqa = 0;
    state->asq = 0;
    state->acq = 0;
    state->sq_head = 0;
    state->cq_tail = 0;
    state->cq_head = 0;
    state->cq_phase = 1;
    state->cq_head_phase = 1;
    state->pending_tail = 0;
    state->io_pending_tail = 0;
    state->admin_pending = false;
    state->io_pending = false;
    state->io_sq = 0;
    state->io_cq = 0;
    state->io_sq_head = 0;
    state->io_cq_tail = 0;
    state->io_cq_head = 0;
    state->io_cq_phase = 1;
    state->io_cq_head_phase = 1;
    state->io_sq_created = false;
    state->io_cq_created = false;
    state->io_vector = 0;
    state->msix_masked = false;
    state->msix_pending = false;
    namespace_reset(state);
    state->queue_resets++;
    return set_csts(state, 0);
}


static ssize_t nvme_read(void *opaque, unsigned bir, uint64_t offset,
                         void *buf, size_t len)
{
    struct nvme_state *state = opaque;

    return state->registers.read(state->registers.state, bir, offset, buf, len);
}


static ssize_t nvme_write(void *opaque, unsigned bir, uint64_t offset,
                          const void *buf, size_t len)
{
    struct nvme_state *state = opaque;
    ssize_t result = state->registers.write(state->registers.state, bir, offset, buf, len);

    if (result >= 0 && state->model->msix_vectors > 0 &&
        bir == state->model->msix_table_bir && len == sizeof(uint32_t) &&
        offset >= state->model->msix_table_offset + 12 &&
        offset < state->model->msix_table_offset +
                 (uint64_t)state->model->msix_vectors * 16 &&
        (offset - state->model->msix_table_offset - 12) % 16 == 0) {
        uint32_t value;

        memcpy(&value, buf, sizeof(value));
        state->msix_masked = (value & (1u << 30)) != 0;
        if (!state->msix_masked && state->msix_pending) {
            state->msix_pending = false;
            if (state->host.irq(state->host.opaque, state->io_vector) < 0) {
                return -EIO;
            }
        }
    }

    if (result >= 0 && bir == 0 && len == sizeof(uint64_t) &&
        (offset == 0x28 || offset == 0x30)) {
        uint64_t value;

        memcpy(&value, buf, sizeof(value));
        if (offset == 0x28) {
            state->asq = value;
        } else {
            state->acq = value;
        }
        return result;
    }
    if (result < 0 || bir != 0 || offset != 0x14 || len != sizeof(uint32_t)) {
        if (result < 0 || bir != 0 || len != sizeof(uint32_t)) {
            return result;
        }
        uint32_t value;
        memcpy(&value, buf, sizeof(value));
        switch (offset) {
        case 0x24: state->aqa = value; break;
        case 0x28: state->asq = (state->asq & 0xffffffff00000000ULL) | value; break;
        case 0x2c: state->asq = (state->asq & 0xffffffffULL) | ((uint64_t)value << 32); break;
        case 0x30: state->acq = (state->acq & 0xffffffff00000000ULL) | value; break;
        case 0x34: state->acq = (state->acq & 0xffffffffULL) | ((uint64_t)value << 32); break;
        case 0x1004: {
            uint16_t cq_size = (uint16_t)(((state->aqa >> 16) & 0xfff) + 1);

            if (value >= cq_size) {
                return -EINVAL;
            }
            if ((uint16_t)value < state->cq_head) {
                state->cq_head_phase ^= 1;
            }
            state->cq_head = (uint16_t)value;
            break;
        }
        case 0x1000:
            state->pending_tail = (uint16_t)value;
            state->admin_pending = true;
            return result;
        case 0x1008:
            if (!state->io_sq_created || value >= state->io_sq_size) {
                return -EINVAL;
            }
            state->io_pending_tail = (uint16_t)value;
            state->io_pending = true;
            return result;
        case 0x100c:
            if (state->io_cq_size == 0 || value >= state->io_cq_size) {
                return -EINVAL;
            }
            if ((uint16_t)value < state->io_cq_head) {
                state->io_cq_head_phase ^= 1;
            }
            state->io_cq_head = (uint16_t)value;
            return result;
        default: break;
        }
        return result;
    }
    uint32_t cc;
    uint32_t csts = 0;

    memcpy(&cc, buf, sizeof(cc));

    if ((cc & 1) == 0) {
        state->queue_resets++;
        if (((cc >> 14) & 3) != 0) {
            state->shutdowns++;
        }
        state->sq_head = 0;
        state->cq_tail = 0;
        state->cq_head = 0;
        state->cq_phase = 1;
        state->cq_head_phase = 1;
        state->pending_tail = 0;
        state->admin_pending = false;
    }

    if ((cc & 1) != 0) {
        csts |= 1;
    }
    if (((cc >> 14) & 3) != 0) {
        csts |= 2u << 2;
    }
    if (set_csts(state, csts) < 0) {
        return -EIO;
    }
    return result;
}


static void nvme_destroy(void *opaque)
{
    struct nvme_state *state = opaque;

    if (state == NULL) {
        return;
    }
    state->registers.destroy(state->registers.state);
    free(state);
}


int behavior_nvme_create(const struct device_model *model,
                         struct device_behavior *out, char *err, size_t err_len)
{
    struct nvme_state *state;

    if (model == NULL || out == NULL || model->class_code != 0x010802) {
        return fail(err, err_len, "NVMe behavior requires PCI class 0x010802");
    }
    state = calloc(1, sizeof(*state));
    if (state == NULL) {
        return fail(err, err_len, "allocate NVMe behavior");
    }
    state->model = model;
    if (behavior_static_create(model, &state->registers, err, err_len) < 0) {
        free(state);
        return -1;
    }
    *out = (struct device_behavior){
        .state = state,
        .bind_host = nvme_bind_host,
        .reset = nvme_reset,
        .service = nvme_service,
        .read = nvme_read,
        .write = nvme_write,
        .destroy = nvme_destroy,
    };
    if (nvme_reset(state) < 0) {
        nvme_destroy(state);
        *out = (struct device_behavior){0};
        return fail(err, err_len, "reset NVMe behavior");
    }
    return 0;
}
