// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Symbol table internal header
//
// Internal details; most calling programs do not need this header,
// unless using verilator public meta comments.

#ifndef VERILATED_VTOP__SYMS_H_
#define VERILATED_VTOP__SYMS_H_  // guard

#include "verilated.h"

// INCLUDE MODEL CLASS

#include "Vtop.h"

// INCLUDE MODULE CLASSES
#include "Vtop___024root.h"
#include "Vtop_device_config.h"
#include "Vtop_IfAXIS128.h"

// DPI TYPES for DPI Export callbacks (Internal use)

// SYMS CLASS (contains all model state)
class alignas(VL_CACHE_LINE_BYTES) Vtop__Syms final : public VerilatedSyms {
  public:
    // INTERNAL STATE
    Vtop* const __Vm_modelp;
    VlDeleter __Vm_deleter;
    bool __Vm_didInit = false;

    // MODULE INSTANCE STATE
    Vtop___024root                 TOP;
    Vtop_device_config             TOP__device_config;
    Vtop_IfAXIS128                 TOP__tb_top__DOT__i_bar__DOT__tlps_cpl;
    Vtop_IfAXIS128                 TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng;
    Vtop_IfAXIS128                 TOP__tb_top__DOT__i_bar__DOT__tlps_ur;
    Vtop_IfAXIS128                 TOP__tb_top__DOT__tlps_dma_out_if;
    Vtop_IfAXIS128                 TOP__tb_top__DOT__tlps_in_if;
    Vtop_IfAXIS128                 TOP__tb_top__DOT__tlps_out_if;

    // SCOPE NAMES
    VerilatedScope* __Vscopep_device_config;
    VerilatedScope* __Vscopep_tb_top;
    VerilatedScope* __Vscopep_tb_top__i_bar;
    VerilatedScope* __Vscopep_tb_top__i_bar__i_bar0;
    VerilatedScope* __Vscopep_tb_top__i_bar__i_bar1;
    VerilatedScope* __Vscopep_tb_top__i_bar__i_bar2;
    VerilatedScope* __Vscopep_tb_top__i_bar__i_bar3;
    VerilatedScope* __Vscopep_tb_top__i_bar__i_bar4;
    VerilatedScope* __Vscopep_tb_top__i_bar__i_bar5;
    VerilatedScope* __Vscopep_tb_top__i_bar__i_bar_rsp_arbiter;
    VerilatedScope* __Vscopep_tb_top__i_bar__i_latency_emu;
    VerilatedScope* __Vscopep_tb_top__i_bar__i_lifecycle_service;
    VerilatedScope* __Vscopep_tb_top__i_bar__i_pcileech_tlps128_bar_rdengine;
    VerilatedScope* __Vscopep_tb_top__i_bar__i_pcileech_tlps128_bar_rdengine__i_fifo_134_134_clk1_bar_rdrsp;
    VerilatedScope* __Vscopep_tb_top__i_bar__i_pcileech_tlps128_bar_wrengine;
    VerilatedScope* __Vscopep_tb_top__i_bar__i_pcileech_tlps128_bar_wrengine__i_fifo_141_141_clk1_bar_wr;
    VerilatedScope* __Vscopep_tb_top__i_bar__i_tlp_normalizer;
    VerilatedScope* __Vscopep_tb_top__i_bar__i_tlp_ur_completer;
    VerilatedScope* __Vscopep_tb_top__i_bar__tlps_cpl;
    VerilatedScope* __Vscopep_tb_top__i_bar__tlps_rdeng;
    VerilatedScope* __Vscopep_tb_top__i_bar__tlps_ur;
    VerilatedScope* __Vscopep_tb_top__tlps_dma_out_if;
    VerilatedScope* __Vscopep_tb_top__tlps_in_if;
    VerilatedScope* __Vscopep_tb_top__tlps_out_if;

    // SCOPE HIERARCHY
    VerilatedHierarchy __Vhier;

    // CONSTRUCTORS
    Vtop__Syms(VerilatedContext* contextp, const char* namep, Vtop* modelp);
    ~Vtop__Syms();

    // METHODS
    const char* name() const { return TOP.vlNamep; }
};

#endif  // guard
