// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vtop.h for the primary calling header

#include "Vtop__pch.h"

// Parameter definitions for Vtop_device_config
constexpr CData/*7:0*/ Vtop_device_config::REVISION_ID;
constexpr CData/*7:0*/ Vtop_device_config::CLASS_BASE;
constexpr CData/*7:0*/ Vtop_device_config::CLASS_SUB;
constexpr CData/*7:0*/ Vtop_device_config::CLASS_PROGIF;
constexpr CData/*7:0*/ Vtop_device_config::DONOR_PME_SUPPORT_MASK;
constexpr CData/*2:0*/ Vtop_device_config::DONOR_MSI_MULTIPLE_MSG;
constexpr CData/*1:0*/ Vtop_device_config::DONOR_PCIE_ASPM_CAP;
constexpr CData/*1:0*/ Vtop_device_config::DONOR_PCIE_ASPM_ENABLE;
constexpr CData/*7:0*/ Vtop_device_config::DONOR_PCIELINK_SPEED;
constexpr CData/*7:0*/ Vtop_device_config::DONOR_PCIELINK_WIDTH;
constexpr SData/*15:0*/ Vtop_device_config::VENDOR_ID;
constexpr SData/*15:0*/ Vtop_device_config::DEVICE_ID;
constexpr SData/*15:0*/ Vtop_device_config::SUBSYS_VENDOR_ID;
constexpr SData/*15:0*/ Vtop_device_config::SUBSYS_DEVICE_ID;
constexpr SData/*11:0*/ Vtop_device_config::DONOR_PM_CAP_OFF;
constexpr SData/*11:0*/ Vtop_device_config::DONOR_MSI_CAP_OFF;
constexpr SData/*11:0*/ Vtop_device_config::DONOR_MSIX_CAP_OFF;
constexpr SData/*11:0*/ Vtop_device_config::DONOR_PCIE_CAP_OFF;
constexpr SData/*11:0*/ Vtop_device_config::DONOR_AER_CAP_OFF;
constexpr SData/*11:0*/ Vtop_device_config::DONOR_LTR_CAP_OFF;
constexpr SData/*11:0*/ Vtop_device_config::DONOR_L1PM_CAP_OFF;
constexpr SData/*11:0*/ Vtop_device_config::DONOR_DSN_CAP_OFF;
constexpr IData/*31:0*/ Vtop_device_config::BAR0_PRESENT;
constexpr IData/*31:0*/ Vtop_device_config::BAR0_SIZE;
constexpr IData/*31:0*/ Vtop_device_config::BAR0_REGS;
constexpr IData/*31:0*/ Vtop_device_config::LATENCY_MIN;
constexpr IData/*31:0*/ Vtop_device_config::LATENCY_MAX;
constexpr IData/*31:0*/ Vtop_device_config::LATENCY_AVG;
constexpr IData/*31:0*/ Vtop_device_config::MSIX_NUM_VECTORS;
constexpr IData/*31:0*/ Vtop_device_config::MSIX_TABLE_OFF;
constexpr IData/*31:0*/ Vtop_device_config::MSIX_PBA_OFF;
constexpr IData/*31:0*/ Vtop_device_config::MSIX_TABLE_BIR;
constexpr IData/*31:0*/ Vtop_device_config::MSIX_PBA_BIR;
constexpr IData/*31:0*/ Vtop_device_config::HAS_NVME_FSM;
constexpr IData/*31:0*/ Vtop_device_config::HAS_NVME_RESP;
constexpr IData/*31:0*/ Vtop_device_config::HAS_XHCI_FSM;
constexpr IData/*31:0*/ Vtop_device_config::HAS_AUDIO_FSM;
constexpr IData/*31:0*/ Vtop_device_config::HAS_MSIX_INT;
constexpr IData/*31:0*/ Vtop_device_config::HAS_DONOR_PM_CAP;
constexpr IData/*31:0*/ Vtop_device_config::HAS_DONOR_MSI_CAP;
constexpr IData/*31:0*/ Vtop_device_config::HAS_DONOR_MSIX_CAP;
constexpr IData/*31:0*/ Vtop_device_config::HAS_DONOR_PCIE_CAP;
constexpr IData/*31:0*/ Vtop_device_config::HAS_DONOR_AER_CAP;
constexpr IData/*31:0*/ Vtop_device_config::HAS_DONOR_LTR_CAP;
constexpr IData/*31:0*/ Vtop_device_config::HAS_DONOR_L1PM_CAP;
constexpr IData/*31:0*/ Vtop_device_config::HAS_DONOR_DSN_CAP;
constexpr IData/*31:0*/ Vtop_device_config::DONOR_MSI_DISABLE_64;
constexpr IData/*31:0*/ Vtop_device_config::HAS_LATENCY_EMU;



Vtop_device_config::Vtop_device_config() = default;
Vtop_device_config::~Vtop_device_config() = default;

void Vtop_device_config::ctor(Vtop__Syms* symsp, const char* namep) {
    vlSymsp = symsp;
    vlNamep = strdup(Verilated::catName(vlSymsp->name(), namep));
    // Reset structure values
}

void Vtop_device_config::__Vconfigure(bool first) {
    (void)first;  // Prevent unused variable warning
}

void Vtop_device_config::dtor() {
    VL_DO_DANGLING(std::free(const_cast<char*>(vlNamep)), vlNamep);
}
