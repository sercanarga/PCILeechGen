// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design internal header
// See Vtop.h for the primary calling header

#ifndef VERILATED_VTOP_DEVICE_CONFIG_H_
#define VERILATED_VTOP_DEVICE_CONFIG_H_  // guard

#include "verilated.h"
#include "verilated_timing.h"


class Vtop__Syms;

class alignas(VL_CACHE_LINE_BYTES) Vtop_device_config final {
  public:

    // INTERNAL VARIABLES
    Vtop__Syms* vlSymsp;
    const char* vlNamep;

    // PARAMETERS
    static constexpr CData/*7:0*/ REVISION_ID = 1U;
    static constexpr CData/*7:0*/ CLASS_BASE = 1U;
    static constexpr CData/*7:0*/ CLASS_SUB = 8U;
    static constexpr CData/*7:0*/ CLASS_PROGIF = 2U;
    static constexpr CData/*7:0*/ DONOR_PME_SUPPORT_MASK = 4U;
    static constexpr CData/*2:0*/ DONOR_MSI_MULTIPLE_MSG = 0U;
    static constexpr CData/*1:0*/ DONOR_PCIE_ASPM_CAP = 0U;
    static constexpr CData/*1:0*/ DONOR_PCIE_ASPM_ENABLE = 0U;
    static constexpr CData/*7:0*/ DONOR_PCIELINK_SPEED = 0U;
    static constexpr CData/*7:0*/ DONOR_PCIELINK_WIDTH = 0U;
    static constexpr SData/*15:0*/ VENDOR_ID = 0x144dU;
    static constexpr SData/*15:0*/ DEVICE_ID = 0xa809U;
    static constexpr SData/*15:0*/ SUBSYS_VENDOR_ID = 0U;
    static constexpr SData/*15:0*/ SUBSYS_DEVICE_ID = 0U;
    static constexpr SData/*11:0*/ DONOR_PM_CAP_OFF = 0x0040U;
    static constexpr SData/*11:0*/ DONOR_MSI_CAP_OFF = 0x0048U;
    static constexpr SData/*11:0*/ DONOR_MSIX_CAP_OFF = 0U;
    static constexpr SData/*11:0*/ DONOR_PCIE_CAP_OFF = 0x0054U;
    static constexpr SData/*11:0*/ DONOR_AER_CAP_OFF = 0U;
    static constexpr SData/*11:0*/ DONOR_LTR_CAP_OFF = 0U;
    static constexpr SData/*11:0*/ DONOR_L1PM_CAP_OFF = 0U;
    static constexpr SData/*11:0*/ DONOR_DSN_CAP_OFF = 0U;
    static constexpr IData/*31:0*/ BAR0_PRESENT = 1U;
    static constexpr IData/*31:0*/ BAR0_SIZE = 0x00004000U;
    static constexpr IData/*31:0*/ BAR0_REGS = 0x0000000dU;
    static constexpr IData/*31:0*/ LATENCY_MIN = 3U;
    static constexpr IData/*31:0*/ LATENCY_MAX = 0x0000000cU;
    static constexpr IData/*31:0*/ LATENCY_AVG = 6U;
    static constexpr IData/*31:0*/ HAS_NVME_FSM = 1U;
    static constexpr IData/*31:0*/ HAS_NVME_RESP = 1U;
    static constexpr IData/*31:0*/ HAS_XHCI_FSM = 0U;
    static constexpr IData/*31:0*/ HAS_AUDIO_FSM = 0U;
    static constexpr IData/*31:0*/ HAS_MSIX_INT = 0U;
    static constexpr IData/*31:0*/ HAS_DONOR_PM_CAP = 1U;
    static constexpr IData/*31:0*/ HAS_DONOR_MSI_CAP = 1U;
    static constexpr IData/*31:0*/ HAS_DONOR_MSIX_CAP = 0U;
    static constexpr IData/*31:0*/ HAS_DONOR_PCIE_CAP = 1U;
    static constexpr IData/*31:0*/ HAS_DONOR_AER_CAP = 0U;
    static constexpr IData/*31:0*/ HAS_DONOR_LTR_CAP = 0U;
    static constexpr IData/*31:0*/ HAS_DONOR_L1PM_CAP = 0U;
    static constexpr IData/*31:0*/ HAS_DONOR_DSN_CAP = 0U;
    static constexpr IData/*31:0*/ DONOR_MSI_DISABLE_64 = 1U;
    static constexpr IData/*31:0*/ HAS_LATENCY_EMU = 1U;

    // CONSTRUCTORS
    Vtop_device_config();
    ~Vtop_device_config();
    void ctor(Vtop__Syms* symsp, const char* namep);
    void dtor();
    VL_UNCOPYABLE(Vtop_device_config);

    // INTERNAL METHODS
    void __Vconfigure(bool first);
};


#endif  // guard
