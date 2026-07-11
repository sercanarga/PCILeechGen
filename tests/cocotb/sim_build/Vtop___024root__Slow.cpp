// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vtop.h for the primary calling header

#include "Vtop__pch.h"

// Parameter definitions for Vtop___024root
constexpr CData/*0:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__NVME_IO_QUEUES_ENABLED;
constexpr CData/*0:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__NVME_IO_DATA_STUBBED;
constexpr CData/*0:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__NVME_IO_READS_ZEROED;
constexpr CData/*0:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__NVME_IO_PRP_ONLY;
constexpr CData/*0:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__NVME_TIMING_ENABLED;
constexpr CData/*2:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__KIND_UNKNOWN;
constexpr CData/*2:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__KIND_MRD;
constexpr CData/*2:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__KIND_MWR;
constexpr CData/*2:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__KIND_IORD;
constexpr CData/*2:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__KIND_IOWR;
constexpr CData/*2:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__KIND_ATOMIC;
constexpr CData/*2:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__KIND_MRDLK;
constexpr CData/*2:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__KIND_CONFIG;
constexpr CData/*0:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__ST_IDLE;
constexpr CData/*0:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__ST_ISSUE;
constexpr CData/*1:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__Q_SELECT;
constexpr CData/*1:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__Q_WAIT;
constexpr CData/*1:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__Q_CHECK;
constexpr CData/*0:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__NVME_ENABLE_IO_QUEUES;
constexpr CData/*0:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__NVME_STUB_IO_DATA;
constexpr CData/*0:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__NVME_ZERO_IO_READS;
constexpr CData/*0:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__NVME_PRP_ONLY_IO;
constexpr CData/*0:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__NVME_TIMING_ENABLE;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_IDLE;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_FETCH_SQE;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_DECODE_CMD;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_EXEC_IDENTIFY;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_EXEC_NSLIST;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_EXEC_LOGPAGE;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_EXEC_ZERO4K;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_BUILD_CQE;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_POST_CQE;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_SEND_MSIX;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_WAIT_DMA_WR;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_IO_WRITE_DMA_REQ;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_IO_WRITE_DMA_WAIT;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_IO_DISK_WR_WAIT;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_IO_READ_DISK_REQ;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_IO_READ_DISK_WAIT;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_IO_ADVANCE;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_IO_ZERO_REQ;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_IO_FLUSH_REQ;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_IO_FLUSH_WAIT;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_SHUTDOWN;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_EXEC_NSID_DESC;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_RESOLVE_IO_PRP;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_PRP_LIST_RD_LO;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_PRP_LIST_RD_HI;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_IO_READ_HOST_WR;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_DSM_FETCH_REQ;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_DSM_FETCH_WAIT;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_DSM_CLEAR_WAIT;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_FORMAT_CLEAR;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_TIMING_DELAY;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_TIMING_PREP;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_TIMING_EVAL;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_AER_POST;
constexpr CData/*5:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__S_AER_FINISH;
constexpr CData/*1:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__IO_DIR_READ;
constexpr CData/*1:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__IO_DIR_WRITE;
constexpr CData/*1:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__IO_DIR_ZERO;
constexpr CData/*3:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__TIMING_ADMIN_SHORT;
constexpr CData/*3:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__TIMING_IDENTIFY;
constexpr CData/*3:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__TIMING_GET_LOG;
constexpr CData/*3:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__TIMING_READ_PENDING;
constexpr CData/*3:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__TIMING_READ_HIT;
constexpr CData/*3:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__TIMING_READ_MISS_ZERO;
constexpr CData/*3:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__TIMING_WRITE;
constexpr CData/*3:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__TIMING_WRITE_ZEROES;
constexpr CData/*3:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__TIMING_FLUSH;
constexpr CData/*3:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__TIMING_DSM;
constexpr CData/*3:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__TIMING_FORMAT;
constexpr CData/*3:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__TIMING_ERROR;
constexpr CData/*7:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__LOG_PAGE_SUPPORTED;
constexpr CData/*7:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__LOG_PAGE_ERROR;
constexpr CData/*7:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__LOG_PAGE_SMART;
constexpr CData/*7:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__LOG_PAGE_FW_SLOT;
constexpr CData/*2:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__S_IDLE;
constexpr CData/*2:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__S_DONE;
constexpr CData/*2:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__S_CLEAR;
constexpr CData/*2:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__S_RAM;
constexpr CData/*2:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__S_RAM_WAIT;
constexpr CData/*3:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__B_IDLE;
constexpr CData/*3:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__B_WR_SEND;
constexpr CData/*3:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__B_WR_SEND2;
constexpr CData/*3:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__B_WR_DONE;
constexpr CData/*3:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__B_RD_SEND;
constexpr CData/*3:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__B_RD_WAIT;
constexpr CData/*3:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__B_RD_EMIT;
constexpr CData/*7:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__NVME_RD_TAG_FIRST;
constexpr CData/*7:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__TAG_FIRST;
constexpr CData/*1:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__OUTCOME_COMPLETED;
constexpr CData/*1:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__OUTCOME_ERROR;
constexpr CData/*1:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__OUTCOME_TIMEOUT;
constexpr CData/*1:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__OUTCOME_CANCELLED;
constexpr SData/*14:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__NVME_SC_DATA_XFER_ERROR;
constexpr SData/*14:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__NVME_SC_INVALID_FIELD;
constexpr SData/*14:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__NVME_SC_INVALID_NS;
constexpr SData/*14:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__NVME_SC_PRP_OFFSET_INVALID;
constexpr SData/*14:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__NVME_SC_LBA_RANGE;
constexpr SData/*14:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__NVME_SC_INVALID_FORMAT;
constexpr SData/*15:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__NVME_NUM_MSIX_VECTORS;
constexpr SData/*15:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__BACKING_COUNT;
constexpr SData/*15:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__BACKING_ONE;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__NVME_BRAM_DISK_WORDS;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__NVME_TIMING_CLOCK_HZ;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__READ_COMPLETION_BOUNDARY;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__MAX_PAYLOAD_BYTES;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__READ_COMPLETION_BOUNDARY;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__MAX_PAYLOAD_BYTES;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__REQUEST_FIFO_DEPTH;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__REQUEST_FIFO_DEPTH;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__FIFO_DEPTH;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_bar0__DOT__BAR_SIZE;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__MIN_LATENCY_CYCLES;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__MAX_LATENCY_CYCLES;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__AVG_LATENCY_CYCLES;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__ENABLE_JITTER;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__BURST_CORRELATION;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__THERMAL_DRIFT_PERIOD;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__WR_MIN_LATENCY;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__WR_MAX_LATENCY;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__CPL_TIMEOUT_CYCLES;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__PRNG_SEED_0;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__PRNG_SEED_1;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__PRNG_SEED_2;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__PRNG_SEED_3;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__NUM_VECTORS;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__DEFER_MSIX_CLEAR;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__INDEX_WIDTH;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__NVME_TIMING_CLK_HZ;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__NVME_TIMING_CYCLES_PER_US;
constexpr IData/*23:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__MAX_XFER_DW;
constexpr IData/*19:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__PRP_LIST_MAX_ENTRIES;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__DISK_WORDS;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__BACKING_LBAS;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__SLOT_WIDTH;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__ADDR_WIDTH;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__CLEAR_WIDTH;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__BANK_BITS;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__BANK_ADDR_WIDTH;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__PIN_WINDOW_LBAS;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__PINNED_SLOTS;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__HASH_SLOTS;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__HASH_WIDTH;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ADDR_WIDTH;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ADDR_WIDTH;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ADDR_WIDTH;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ADDR_WIDTH;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__TAG_COUNT;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__TIMEOUT_WIDTH;
constexpr IData/*23:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__TIMEOUT_CYCLES;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__INDEX_WIDTH;
constexpr QData/*63:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__NVME_ADVERTISED_LBAS;
constexpr QData/*63:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__NVME_ADVERTISED_LBAS;
constexpr QData/*63:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__WIN64;


void Vtop___024root___ctor_var_reset(Vtop___024root* vlSelf);

Vtop___024root::Vtop___024root(Vtop__Syms* symsp, const char* namep)
    : __VdlySched{*symsp->_vm_contextp__}
 {
    vlSymsp = symsp;
    vlNamep = strdup(namep);
    // Reset structure values
    Vtop___024root___ctor_var_reset(this);
}

void Vtop___024root::__Vconfigure(bool first) {
    (void)first;  // Prevent unused variable warning
}

Vtop___024root::~Vtop___024root() {
    VL_DO_DANGLING(std::free(const_cast<char*>(vlNamep)), vlNamep);
}
