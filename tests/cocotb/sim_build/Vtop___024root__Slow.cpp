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
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_bar0__DOT__BAR_STORAGE_BYTES;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_bar0__DOT__BAR_WORDS;
constexpr IData/*31:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__i_bar0__DOT__BAR_WORD_ADDR_BITS;
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
constexpr QData/*63:0*/ Vtop___024root::tb_top__DOT__i_bar__DOT__NVME_ADVERTISED_LBAS;


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
