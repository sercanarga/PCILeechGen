// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vtop.h for the primary calling header

#include "Vtop__pch.h"

VL_ATTR_COLD void Vtop___024root___eval_static(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_static\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__clk = 0U;
    vlSelfRef.tb_top__DOT__rst = 1U;
    vlSelfRef.tb_top__DOT__tlps_in_tdata[0U] = 0U;
    vlSelfRef.tb_top__DOT__tlps_in_tdata[1U] = 0U;
    vlSelfRef.tb_top__DOT__tlps_in_tdata[2U] = 0U;
    vlSelfRef.tb_top__DOT__tlps_in_tdata[3U] = 0U;
    vlSelfRef.tb_top__DOT__tlps_in_tkeepdw = 0U;
    vlSelfRef.tb_top__DOT__tlps_in_tvalid = 0U;
    vlSelfRef.tb_top__DOT__tlps_in_tlast = 0U;
    vlSelfRef.tb_top__DOT__tlps_in_tuser = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__reset_d = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state = 0U;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__cpl_select_ur__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_select_ur;
    vlSelfRef.__VactTriggered[0U] = (1ULL | vlSelfRef.__VactTriggered[0U]);
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__cpl_select_ur__1 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_select_ur;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__clk__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__clk__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__clk;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__clk__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__clk;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__clk__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__clk;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__clk__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__clk;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__clk__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__clk;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__clk__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__clk;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_bar0__DOT__clk__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__clk;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__clk__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__clk;
    do {
        vlSelfRef.__VactTriggeredAcc[vlSelfRef.__Vi] 
            = vlSelfRef.__VactTriggered[vlSelfRef.__Vi];
        vlSelfRef.__Vi = ((IData)(1U) + vlSelfRef.__Vi);
    } while ((0U >= vlSelfRef.__Vi));
}

VL_ATTR_COLD void Vtop___024root___eval_static__TOP(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_static__TOP\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__clk = 0U;
    vlSelfRef.tb_top__DOT__rst = 1U;
    vlSelfRef.tb_top__DOT__tlps_in_tdata[0U] = 0U;
    vlSelfRef.tb_top__DOT__tlps_in_tdata[1U] = 0U;
    vlSelfRef.tb_top__DOT__tlps_in_tdata[2U] = 0U;
    vlSelfRef.tb_top__DOT__tlps_in_tdata[3U] = 0U;
    vlSelfRef.tb_top__DOT__tlps_in_tkeepdw = 0U;
    vlSelfRef.tb_top__DOT__tlps_in_tvalid = 0U;
    vlSelfRef.tb_top__DOT__tlps_in_tlast = 0U;
    vlSelfRef.tb_top__DOT__tlps_in_tuser = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__reset_d = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state = 0U;
}

VL_ATTR_COLD void Vtop___024root___eval_initial__TOP(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_initial__TOP\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__cfg_command = 6U;
    vlSelfRef.tb_top__DOT__cfg_power_state = 0U;
    vlSelfRef.tb_top__DOT__cfg_flr_in_process = 0U;
    vlSelfRef.tb_top__DOT__cfg_to_turnoff = 0U;
    vlSelfRef.tb_top__DOT__cfg_link_up = 1U;
    vlSelfRef.tb_top__DOT__cfg_msi_enable = 0U;
    vlSelfRef.tb_top__DOT__cfg_msix_enable = 1U;
    vlSelfRef.tb_top__DOT__cfg_msix_function_mask = 0U;
    vlSelfRef.tb_top__DOT__pcie_id = 0U;
    vlSelfRef.tb_top__DOT__bar_en = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_en = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__pcie_id = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__cfg_command = 6U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__cfg_power_state = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__cfg_flr_in_process = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__cfg_to_turnoff = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__cfg_link_up = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__cfg_msi_enable = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__cfg_msix_enable = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__cfg_msix_function_mask = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__flash_sdo_dq1 = 0U;
    vlSelfRef.tb_top__DOT__intr_req = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__cfg_mse = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__cfg_bme = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__hda_crst_falling = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__intr_req = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__flr_in_process = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__pm_dstate = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__memory_space_enable = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__bus_master_enable = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__turnoff_pending = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__link_up = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__generation = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__pcie_id = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd3_enable = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__prog_full = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__prog_empty = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__pcie_id = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__wrengine_ready = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlps_in_ready = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tvalid = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__dout[0U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__dout[1U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__dout[2U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__dout[3U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__dout[4U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__full = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__empty = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__prog_empty = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__valid = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__i = 0U;
    while (VL_GTS_III(32, 0x00000400U, vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__i)) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem[(0x000003ffU 
                                                                 & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__i)] = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__i 
            = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__i);
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__crst_falling = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__correlated_latency = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[0U] = 0x0fU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[1U] = 0x1fU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[2U] = 0x2fU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[3U] = 0x3fU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[4U] = 0x4fU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[5U] = 0x5fU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[6U] = 0x6fU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[7U] = 0x7fU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[8U] = 0x8fU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[9U] = 0x9fU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[10U] = 0xafU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[11U] = 0xbfU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[12U] = 0xcfU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[13U] = 0xdfU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[14U] = 0xefU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[15U] = 0xffU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__latency_range = 0x0dU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_range = 7U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_rsp_ctx[0U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_rsp_ctx[1U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_rsp_ctx[2U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_rsp_data = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_rsp_valid = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__intr_req = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_rsp_ctx[0U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_rsp_ctx[1U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_rsp_ctx[2U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_rsp_data = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_rsp_valid = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__intr_req = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_rsp_ctx[0U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_rsp_ctx[1U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_rsp_ctx[2U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_rsp_data = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_rsp_valid = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__intr_req = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_rsp_ctx[0U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_rsp_ctx[1U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_rsp_ctx[2U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_rsp_data = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_rsp_valid = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__intr_req = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_rsp_ctx[0U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_rsp_ctx[1U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_rsp_ctx[2U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_rsp_data = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_rsp_valid = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__intr_req = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tkeepdw = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[6U][0U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[6U][1U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[6U][2U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[6U][0U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[6U][1U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[6U][2U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[6U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[6U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tuser = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tdata[0U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tdata[1U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tdata[2U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tdata[3U] = 0U;
}

VL_ATTR_COLD void Vtop___024root___eval_final(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_final\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
}

#ifdef VL_DEBUG
VL_ATTR_COLD void Vtop___024root___dump_triggers__stl(const VlUnpacked<QData/*63:0*/, 2> &triggers, const std::string &tag);
#endif  // VL_DEBUG
VL_ATTR_COLD bool Vtop___024root___eval_phase__stl(Vtop___024root* vlSelf);

VL_ATTR_COLD void Vtop___024root___eval_settle(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_settle\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    IData/*31:0*/ __VstlIterCount;
    // Body
    __VstlIterCount = 0U;
    vlSelfRef.__VstlFirstIteration = 1U;
    do {
        if (VL_UNLIKELY(((0x00002710U < __VstlIterCount)))) {
#ifdef VL_DEBUG
            Vtop___024root___dump_triggers__stl(vlSelfRef.__VstlTriggered, "stl"s);
#endif
            VL_FATAL_MT("tb_top.sv", 4, "", "DIDNOTCONVERGE: Settle region did not converge after '--converge-limit' of 10000 tries");
        }
        __VstlIterCount = ((IData)(1U) + __VstlIterCount);
        vlSelfRef.__VstlPhaseResult = Vtop___024root___eval_phase__stl(vlSelf);
        vlSelfRef.__VstlFirstIteration = 0U;
    } while (vlSelfRef.__VstlPhaseResult);
}

VL_ATTR_COLD void Vtop___024root___eval_triggers_vec__stl(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_triggers_vec__stl\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.__VstlTriggered[1U] = ((0xfffffffffffffffeULL 
                                      & vlSelfRef.__VstlTriggered[1U]) 
                                     | (IData)((IData)(vlSelfRef.__VstlFirstIteration)));
    vlSelfRef.__VstlTriggered[0U] = (QData)((IData)(
                                                    ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_select_ur) 
                                                     != (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__cpl_select_ur__0))));
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__cpl_select_ur__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_select_ur;
    if (VL_UNLIKELY(((1U & (~ (IData)(vlSelfRef.__VstlDidInit)))))) {
        vlSelfRef.__VstlDidInit = 1U;
        vlSelfRef.__VstlTriggered[0U] = (1ULL | vlSelfRef.__VstlTriggered[0U]);
    }
}

VL_ATTR_COLD bool Vtop___024root___trigger_anySet__stl(const VlUnpacked<QData/*63:0*/, 2> &in);

#ifdef VL_DEBUG
VL_ATTR_COLD void Vtop___024root___dump_triggers__stl(const VlUnpacked<QData/*63:0*/, 2> &triggers, const std::string &tag) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___dump_triggers__stl\n"); );
    // Body
    if ((1U & (~ (IData)(Vtop___024root___trigger_anySet__stl(triggers))))) {
        VL_DBG_MSGS("         No '" + tag + "' region triggers active\n");
    }
    if ((1U & (IData)(triggers[0U]))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 0 is active: @([hybrid] tb_top.i_bar.cpl_select_ur)\n");
    }
    if ((1U & (IData)(triggers[1U]))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 64 is active: Internal 'stl' trigger - first iteration\n");
    }
}
#endif  // VL_DEBUG

VL_ATTR_COLD bool Vtop___024root___trigger_anySet__stl(const VlUnpacked<QData/*63:0*/, 2> &in) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___trigger_anySet__stl\n"); );
    // Locals
    IData/*31:0*/ n;
    // Body
    n = 0U;
    do {
        if (in[n]) {
            return (1U);
        }
        n = ((IData)(1U) + n);
    } while ((2U > n));
    return (0U);
}

void Vtop___024root___ico_sequent__TOP__0(Vtop___024root* vlSelf);
void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__tlps_in_if__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_ur__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_cpl__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__0(Vtop_IfAXIS128* vlSelf);
void Vtop___024root___ico_sequent__TOP__1(Vtop___024root* vlSelf);
void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__tlps_in_if__1(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__1(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_cpl__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_ur__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__tlps_out_if__0(Vtop_IfAXIS128* vlSelf);
void Vtop___024root___ico_comb__TOP__2(Vtop___024root* vlSelf);

VL_ATTR_COLD void Vtop___024root___eval_stl(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_stl\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    if ((1ULL & vlSelfRef.__VstlTriggered[1U])) {
        Vtop___024root___ico_sequent__TOP__0(vlSelf);
        Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__tlps_in_if__0((&vlSymsp->TOP__tb_top__DOT__tlps_in_if));
        Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_ur__0((&vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_ur));
        Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_cpl__0((&vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_cpl));
        Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__0((&vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng));
        Vtop___024root___ico_sequent__TOP__1(vlSelf);
        Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__tlps_in_if__1((&vlSymsp->TOP__tb_top__DOT__tlps_in_if));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tlps_in_valid 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_rd) 
               & (IData)(vlSymsp->TOP__tb_top__DOT__tlps_in_if.__VdfgRegularize_hebeb780c_0_3));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlps_in_valid 
            = ((IData)(vlSymsp->TOP__tb_top__DOT__tlps_in_if.__VdfgRegularize_hebeb780c_0_3) 
               & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr) 
                  & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_request_supported) 
                     | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr_last))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_pending 
            = ((IData)(vlSymsp->TOP__tb_top__DOT__tlps_in_if.__VdfgRegularize_hebeb780c_0_3) 
               & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_request_present) 
                  & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_ur_required) 
                     & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_non_posted_request))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_push 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tlps_in_valid) 
               & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_can_accept));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__wr_en 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlps_in_valid;
    }
    if ((1ULL & (vlSelfRef.__VstlTriggered[1U] | vlSelfRef.__VstlTriggered[0U]))) {
        Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__0((&vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_en 
            = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng.tready;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__valid 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_en) 
               & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__empty)));
        Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__1((&vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng));
        vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_select_ur 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_grant_active)
                ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_grant_ur)
                : (((IData)(vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_ur.tvalid) 
                    & (IData)(vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng.tvalid))
                    ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_rr_ur)
                    : (IData)(vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_ur.tvalid)));
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_select_ur) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_selected_valid 
                = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_ur.tvalid;
            vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_selected_last 
                = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_ur.tlast;
        } else {
            vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_selected_valid 
                = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng.tvalid;
            vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_selected_last 
                = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng.tlast;
        }
        Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_cpl__0((&vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_cpl));
        Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_ur__0((&vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_ur));
        Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__tlps_out_if__0((&vlSymsp->TOP__tb_top__DOT__tlps_out_if));
        Vtop___024root___ico_comb__TOP__2(vlSelf);
    }
}

VL_ATTR_COLD bool Vtop___024root___eval_phase__stl(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_phase__stl\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    CData/*0:0*/ __VstlExecute;
    // Body
    Vtop___024root___eval_triggers_vec__stl(vlSelf);
#ifdef VL_DEBUG
    if (VL_UNLIKELY(vlSymsp->_vm_contextp__->debug())) {
        Vtop___024root___dump_triggers__stl(vlSelfRef.__VstlTriggered, "stl"s);
    }
#endif
    __VstlExecute = Vtop___024root___trigger_anySet__stl(vlSelfRef.__VstlTriggered);
    if (__VstlExecute) {
        Vtop___024root___eval_stl(vlSelf);
    }
    return (__VstlExecute);
}

bool Vtop___024root___trigger_anySet__ico(const VlUnpacked<QData/*63:0*/, 2> &in);

#ifdef VL_DEBUG
VL_ATTR_COLD void Vtop___024root___dump_triggers__ico(const VlUnpacked<QData/*63:0*/, 2> &triggers, const std::string &tag) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___dump_triggers__ico\n"); );
    // Body
    if ((1U & (~ (IData)(Vtop___024root___trigger_anySet__ico(triggers))))) {
        VL_DBG_MSGS("         No '" + tag + "' region triggers active\n");
    }
    if ((1U & (IData)(triggers[0U]))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 0 is active: @([hybrid] tb_top.i_bar.cpl_select_ur)\n");
    }
    if ((1U & (IData)(triggers[1U]))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 64 is active: Internal 'ico' trigger - first iteration\n");
    }
}
#endif  // VL_DEBUG

bool Vtop___024root___trigger_anySet__act(const VlUnpacked<QData/*63:0*/, 1> &in);

#ifdef VL_DEBUG
VL_ATTR_COLD void Vtop___024root___dump_triggers__act(const VlUnpacked<QData/*63:0*/, 1> &triggers, const std::string &tag) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___dump_triggers__act\n"); );
    // Body
    if ((1U & (~ (IData)(Vtop___024root___trigger_anySet__act(triggers))))) {
        VL_DBG_MSGS("         No '" + tag + "' region triggers active\n");
    }
    if ((1U & (IData)(triggers[0U]))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 0 is active: @([hybrid] tb_top.i_bar.cpl_select_ur)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 1U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 1 is active: @(posedge tb_top.i_bar.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 2U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 2 is active: @(posedge tb_top.i_bar.i_lifecycle_service.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 3U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 3 is active: @(posedge tb_top.i_bar.i_pcileech_tlps128_bar_rdengine.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 4U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 4 is active: @(posedge tb_top.i_bar.i_pcileech_tlps128_bar_rdengine.i_fifo_134_134_clk1_bar_rdrsp.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 5U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 5 is active: @(posedge tb_top.i_bar.i_tlp_ur_completer.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 6U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 6 is active: @(posedge tb_top.i_bar.i_pcileech_tlps128_bar_wrengine.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 7U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 7 is active: @(posedge tb_top.i_bar.i_bar_rsp_arbiter.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 8U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 8 is active: @(posedge tb_top.i_bar.i_bar0.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 9U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 9 is active: @(posedge tb_top.i_bar.i_latency_emu.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 0x0000000aU)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 10 is active: @([true] __VdlySched.awaitingCurrentTime())\n");
    }
}
#endif  // VL_DEBUG

VL_ATTR_COLD void Vtop___024root___ctor_var_reset(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___ctor_var_reset\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelf->tb_top__DOT__cfg_command = 6U;
    ;
    vlSelf->tb_top__DOT__cfg_power_state = 0U;
    ;
    vlSelf->tb_top__DOT__cfg_flr_in_process = 0U;
    ;
    vlSelf->tb_top__DOT__cfg_to_turnoff = 0U;
    ;
    vlSelf->tb_top__DOT__cfg_link_up = 1U;
    ;
    vlSelf->tb_top__DOT__cfg_msi_enable = 0U;
    ;
    vlSelf->tb_top__DOT__cfg_msix_enable = 1U;
    ;
    vlSelf->tb_top__DOT__cfg_msix_function_mask = 0U;
    ;
    vlSelf->tb_top__DOT__pcie_id = 0U;
    ;
    vlSelf->tb_top__DOT__bar_en = 1U;
    ;
    const uint64_t __VscopeHash = VL_MURMUR64_HASH(vlSelf->vlNamep);
    VL_SCOPED_RAND_RESET_W(128, vlSelf->tb_top__DOT__tlps_out_tdata, __VscopeHash, 16747214199302942365ull);
    vlSelf->tb_top__DOT__tlps_out_tkeepdw = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 8264429687089180992ull);
    vlSelf->tb_top__DOT__tlps_out_tvalid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 6395762658107723871ull);
    vlSelf->tb_top__DOT__tlps_out_tlast = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 16450648599733774355ull);
    vlSelf->tb_top__DOT__tlps_out_tuser = VL_SCOPED_RAND_RESET_I(9, __VscopeHash, 17342319213029503465ull);
    vlSelf->tb_top__DOT__tlps_out_has_data = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 11477534266498542208ull);
    vlSelf->tb_top__DOT__tlps_dma_out_tvalid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 15495265200251862157ull);
    vlSelf->tb_top__DOT__intr_req = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__rst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 200354017822851892ull);
    vlSelf->tb_top__DOT__i_bar__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5183640662232166825ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar_en = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__pcie_id = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__cfg_command = 6U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__cfg_power_state = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__cfg_flr_in_process = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__cfg_to_turnoff = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__cfg_link_up = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__cfg_msi_enable = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__cfg_msix_enable = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__cfg_msix_function_mask = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__flash_csn = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14370257197210649650ull);
    vlSelf->tb_top__DOT__i_bar__DOT__flash_sdi_dq0 = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 282976605691636626ull);
    vlSelf->tb_top__DOT__i_bar__DOT__flash_sdo_dq1 = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__flash_wpn_dq2 = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 18418543270268129798ull);
    vlSelf->tb_top__DOT__i_bar__DOT__flash_hldn_dq3 = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1449949257993782067ull);
    vlSelf->tb_top__DOT__i_bar__DOT__intr_req = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__cfg_mse = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__cfg_bme = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__device_reset = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9299937047987487436ull);
    vlSelf->tb_top__DOT__i_bar__DOT__io_enabled = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 800537706129363748ull);
    vlSelf->tb_top__DOT__i_bar__DOT__dma_enabled = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 10219091600365513000ull);
    vlSelf->tb_top__DOT__i_bar__DOT__lifecycle_quiesce = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 15368278941741451304ull);
    vlSelf->tb_top__DOT__i_bar__DOT__lifecycle_generation = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 7993734939244974819ull);
    vlSelf->tb_top__DOT__i_bar__DOT__in_is_wr_ready = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14240863171373311622ull);
    vlSelf->tb_top__DOT__i_bar__DOT__in_is_wr_last = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__in_is_bar = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5182448629857945030ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_request_present = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8646103267730812998ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_request_supported = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4159350575608973831ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_unsupported_request = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5794306240884041947ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_ur_required = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 12848315205201590059ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_non_posted_request = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9340084700738335711ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_malformed_request = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4178648520610408714ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_posted_write = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 10511998048576874945ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_memory_read = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 15056861026200526134ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_memory_write = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 11184839504102593384ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_request_kind = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 791922838427282197ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_header_4dw = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14125253002905450095ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_address = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 1891304145170907554ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_length_dw = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 3845121768432691906ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_first_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 11866947100909131499ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_last_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 15224915428123269348ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_requester_id = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 563109783535678423ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_tag = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 16160087251110571695ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_traffic_class = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 4518685915377940785ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_attributes = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 1579987900695249206ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_bir = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 1559973235074030135ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_bar_mask = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 15454364021163189008ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_enabled_byte_count = VL_SCOPED_RAND_RESET_I(13, __VscopeHash, 1850574033774366470ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_first_completion_byte_count = VL_SCOPED_RAND_RESET_I(13, __VscopeHash, 18428095296036693083ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_first_completion_dw = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 14215509204931888788ull);
    vlSelf->tb_top__DOT__i_bar__DOT__norm_first_lower_address = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 16392820115818806066ull);
    vlSelf->tb_top__DOT__i_bar__DOT__in_is_first = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4519978964630640394ull);
    vlSelf->tb_top__DOT__i_bar__DOT__in_is_rd = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 16695204164290664404ull);
    vlSelf->tb_top__DOT__i_bar__DOT__in_is_wr = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 7310972847972917702ull);
    vlSelf->tb_top__DOT__i_bar__DOT__wr_bar = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 17719236591828138186ull);
    vlSelf->tb_top__DOT__i_bar__DOT__wr_addr = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 4953886031739836147ull);
    vlSelf->tb_top__DOT__i_bar__DOT__wr_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 16521364630982098765ull);
    vlSelf->tb_top__DOT__i_bar__DOT__wr_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 7993883811967434411ull);
    vlSelf->tb_top__DOT__i_bar__DOT__wr_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5007307757657208397ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__rd_req_ctx, __VscopeHash, 8251038352763046038ull);
    vlSelf->tb_top__DOT__i_bar__DOT__rd_req_bar = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 1873229759907353423ull);
    vlSelf->tb_top__DOT__i_bar__DOT__rd_req_addr = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 7039103940836144943ull);
    vlSelf->tb_top__DOT__i_bar__DOT__rd_req_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 16781609293966398729ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__rd_rsp_ctx, __VscopeHash, 13261241748354092281ull);
    vlSelf->tb_top__DOT__i_bar__DOT__rd_rsp_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 14928343645545481626ull);
    vlSelf->tb_top__DOT__i_bar__DOT__rd_rsp_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8506901964692840361ull);
    vlSelf->tb_top__DOT__i_bar__DOT__wr_addr_bar0 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 3888220385984426726ull);
    vlSelf->tb_top__DOT__i_bar__DOT__rd_req_addr_bar0 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 10118447542561705131ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_wr_hit = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14987071750530931483ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_rd_hit = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 13642565424851012803ull);
    vlSelf->tb_top__DOT__i_bar__DOT__ur_request_ready = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1678583407718958362ull);
    vlSelf->tb_top__DOT__i_bar__DOT__ur_request_pending = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 746721529250601957ull);
    for (int __Vi0 = 0; __Vi0 < 256; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__ur_request_fifo[__Vi0] = VL_SCOPED_RAND_RESET_I(30, __VscopeHash, 2641706491715608331ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__ur_request_read_ptr = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 18425634244444122287ull);
    vlSelf->tb_top__DOT__i_bar__DOT__ur_request_write_ptr = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 5725271153268259351ull);
    vlSelf->tb_top__DOT__i_bar__DOT__ur_request_count = VL_SCOPED_RAND_RESET_I(9, __VscopeHash, 11336360434725756733ull);
    vlSelf->tb_top__DOT__i_bar__DOT__ur_request_dequeue = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5897347339434977214ull);
    vlSelf->tb_top__DOT__i_bar__DOT__ur_request_enqueue = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5147179950584291595ull);
    vlSelf->tb_top__DOT__i_bar__DOT__cpl_grant_active = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 7318156789130892464ull);
    vlSelf->tb_top__DOT__i_bar__DOT__cpl_grant_ur = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 852586218100470448ull);
    vlSelf->tb_top__DOT__i_bar__DOT__cpl_rr_ur = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 15799296632406399269ull);
    vlSelf->tb_top__DOT__i_bar__DOT__cpl_select_ur = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 6988097237201841934ull);
    vlSelf->tb_top__DOT__i_bar__DOT__cpl_selected_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 13710345178547841344ull);
    vlSelf->tb_top__DOT__i_bar__DOT__cpl_selected_last = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4553644441760093735ull);
    vlSelf->tb_top__DOT__i_bar__DOT__wrengine_ready = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_wr_ack = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17271872947348560473ull);
    vlSelf->tb_top__DOT__i_bar__DOT__hda_crst_falling = 0U;
    ;
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__bar0_raw_ctx, __VscopeHash, 14610067568952763635ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_raw_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 6539557169725307975ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_raw_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5922681203298322916ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__bar0_emu_ctx, __VscopeHash, 1219858546133133140ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_emu_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 8219891235491692205ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_emu_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 16607686985978314308ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_emu_busy = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 13940883529852168599ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_emu_ready = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4044801866266259774ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_emu_busy_d1 = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 2817747359966543038ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_emu_accepted = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9507176434138655670ull);
    for (int __Vi0 = 0; __Vi0 < 16; ++__Vi0) {
        VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__bar0_buf_ctx[__Vi0], __VscopeHash, 5755260116445561114ull);
    }
    for (int __Vi0 = 0; __Vi0 < 16; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__bar0_buf_data[__Vi0] = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 11207682139054631092ull);
    }
    for (int __Vi0 = 0; __Vi0 < 16; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__bar0_buf_addr[__Vi0] = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 8319130871302735265ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_buf_count = VL_SCOPED_RAND_RESET_I(5, __VscopeHash, 1015307088905129286ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_buf_rd = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 5153029898916908461ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_buf_wr = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 6540597480098517195ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_buf_full = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 12465697214527164437ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_buf_empty = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 812677986694818710ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_buf_almost_full = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 2292946139510002373ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__bar0_buf_ctx_w, __VscopeHash, 6269135954524683359ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_buf_data_w = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 5996471273050975329ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_buf_addr_w = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 16209506918006437674ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_buf_pop = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4540186713707347944ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_req_addr_from_ctx = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 6275011101821015048ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_spill_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1666758482024705777ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__bar0_spill_ctx, __VscopeHash, 13627815836949199128ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_spill_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 15154799841976368145ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_buf_write = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14453470191500949793ull);
    for (int __Vi0 = 0; __Vi0 < 7; ++__Vi0) {
        VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__bar_base_ctx[__Vi0], __VscopeHash, 12797274884874905472ull);
    }
    for (int __Vi0 = 0; __Vi0 < 7; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__bar_base_data[__Vi0] = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 11459727899665174671ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__bar_base_valid = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 11769542718768725757ull);
    for (int __Vi0 = 0; __Vi0 < 7; ++__Vi0) {
        VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__bar_rsp_ctx[__Vi0], __VscopeHash, 16188974058547500150ull);
    }
    for (int __Vi0 = 0; __Vi0 < 7; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__bar_rsp_data[__Vi0] = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 11207305771248842079ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__bar_rsp_valid = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 731102323566620405ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__bar0_base_ctx, __VscopeHash, 13694318319896257271ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_base_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 3446294472501668799ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar0_base_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4925751892757974743ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar_rsp_ready = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 8262763467872852818ull);
    for (int __Vi0 = 0; __Vi0 < 7; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data[__Vi0] = 0;
    }
    for (int __Vi0 = 0; __Vi0 < 7; ++__Vi0) {
        VL_ZERO_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[__Vi0]);
    }
    vlSelf->tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17611092729108411831ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__fundamental_reset = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 12272690359321834691ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__flr_in_process = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__pm_dstate = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__memory_space_enable = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__bus_master_enable = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__turnoff_pending = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__link_up = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__device_reset = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 13299285693723612845ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__io_enabled = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 11149746773420881637ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__dma_enabled = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8639127319454075576ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__quiesce = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 7214520789390645155ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__generation = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 11073769541763995849ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__active = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 10948711629753982543ull);
    VL_SCOPED_RAND_RESET_W(128, vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data, __VscopeHash, 12136481029597844549ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user = VL_SCOPED_RAND_RESET_I(9, __VscopeHash, 1471548353443477044ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_last = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 18191758911221334302ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_present = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 15409840407864790160ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_supported = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8998318726157690486ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__unsupported_request = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 757028488009279527ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__ur_required = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 269371317727181275ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__non_posted_request = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 815787885156797700ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__malformed_request = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5867036560420800764ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__posted_write = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4419229439443202204ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__memory_read = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1291883132477010308ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__memory_write = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4899660134512013133ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 16358995384666465268ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__header_4dw = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14028157248983767148ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__address = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 11037231395632799907ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__length_dw = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 10553239437019467575ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 11200725595556063072ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__last_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 11724125651482586493ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__requester_id = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 16100524498921714468ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tag = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 11102949603475112353ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__traffic_class = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 17584393674726585902ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__attributes = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 5205942078050064679ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__bir = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 8082769386851961944ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__bar_mask = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 2759014945134106502ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__enabled_byte_count = VL_SCOPED_RAND_RESET_I(13, __VscopeHash, 662461686530537885ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_completion_byte_count = VL_SCOPED_RAND_RESET_I(13, __VscopeHash, 5389465066311369984ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_completion_dw = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 15239025682080244055ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_lower_address = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 13407636989728098428ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__fmt = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 10004277275484593337ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type = VL_SCOPED_RAND_RESET_I(5, __VscopeHash, 10062534069526433554ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__rcb_bytes_left = VL_SCOPED_RAND_RESET_I(13, __VscopeHash, 6294247054685183622ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__rcb_dw_left = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 9567192295528427365ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__mps_dw = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 13932812775287428822ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__legal_be = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 13400987461910030621ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__legal_bar = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5609420069962930978ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__crosses_4k = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5539354474640957254ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 7027151011361561689ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 3359715167957482809ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__pcie_id = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__stall = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9377486140851851664ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tlps_in_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 15602043212891547237ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_address = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 12024391740930737944ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_length_dw = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 12076864393847767920ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_first_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 638225789106801789ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_last_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 4048452139725201871ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_requester_id = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 14751078984628697404ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_tag = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 17204046472329617980ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_traffic_class = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 18179285695847645795ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_attributes = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 13945577347013169595ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_header_4dw = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 13625783957220191562ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_bar_mask = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 1782632832970242639ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_enabled_byte_count = VL_SCOPED_RAND_RESET_I(13, __VscopeHash, 16994493130238236509ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_first_completion_byte_count = VL_SCOPED_RAND_RESET_I(13, __VscopeHash, 16229423785586675107ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_first_completion_dw = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 4246329542365532064ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx, __VscopeHash, 17754761007339143482ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_bar = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 17043103541372833682ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_addr = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 9262631915840031302ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 2520752771172232120ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx, __VscopeHash, 10966342942561099958ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 15206010298319616372ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 525364858053364565ull);
    for (int __Vi0 = 0; __Vi0 < 256; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_addr_fifo[__Vi0] = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 14548523646858582557ull);
    }
    for (int __Vi0 = 0; __Vi0 < 256; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_len_fifo[__Vi0] = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 12951927092952390255ull);
    }
    for (int __Vi0 = 0; __Vi0 < 256; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_fbe_fifo[__Vi0] = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 1615965170483315198ull);
    }
    for (int __Vi0 = 0; __Vi0 < 256; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_lbe_fifo[__Vi0] = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 3694991662334937911ull);
    }
    for (int __Vi0 = 0; __Vi0 < 256; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_id_fifo[__Vi0] = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 14805273242841230271ull);
    }
    for (int __Vi0 = 0; __Vi0 < 256; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tag_fifo[__Vi0] = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 8573479040611135591ull);
    }
    for (int __Vi0 = 0; __Vi0 < 256; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tc_fifo[__Vi0] = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 18343009262586286025ull);
    }
    for (int __Vi0 = 0; __Vi0 < 256; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_attr_fifo[__Vi0] = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 12242136453563656935ull);
    }
    for (int __Vi0 = 0; __Vi0 < 256; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_4dw_fifo[__Vi0] = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 10483289713848719499ull);
    }
    for (int __Vi0 = 0; __Vi0 < 256; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bar_fifo[__Vi0] = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 12496645275137768328ull);
    }
    for (int __Vi0 = 0; __Vi0 < 256; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bc_fifo[__Vi0] = VL_SCOPED_RAND_RESET_I(13, __VscopeHash, 6437724838154933975ull);
    }
    for (int __Vi0 = 0; __Vi0 < 256; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_cpl_fifo[__Vi0] = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 17200714662985010175ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 7819285513138569879ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 17687266675360448079ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_count = VL_SCOPED_RAND_RESET_I(9, __VscopeHash, 14589552408859965854ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__state = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 3276865816128091519ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 10033310504181904828ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_length_dw = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 1047623409924455368ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 5757154220708688933ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 15824375240031142457ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_remaining_dw = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 16792650734501542250ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_byte_count = VL_SCOPED_RAND_RESET_I(13, __VscopeHash, 15628215898098507733ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 1052444324994097029ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_last_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 12406854124855379478ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_id = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 15040006121516913509ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_tag = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 3375333296393572313ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_tc = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 4764477398230992062ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_attr = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 10042891893830733691ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_4dw = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14956194969775202322ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_bar = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 6777977506952216086ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_packet = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8061713393977722372ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1193373416423944104ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_fifo_full = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8055941948043082575ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_pop = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4883591434322110158ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_can_accept = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8025279231267483039ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_push = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17207633180256421304ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd3_enable = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__issue_dw = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9869800665846497294ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_last = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4864142483273669941ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_last_packet = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 159754463990906202ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 6347663435981569048ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first_offset = VL_SCOPED_RAND_RESET_I(2, __VscopeHash, 6515775924015022909ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__context_address = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 5526843980828492036ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_first = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 6188279682892092386ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_last = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 3837410381341015494ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_dwlen = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 13184714786623280257ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_byte_count = VL_SCOPED_RAND_RESET_I(12, __VscopeHash, 7351059514085578852ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_tc = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 8029888193220179965ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_attr = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 9452300551571009013ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_4dw = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17040713170920800479ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_tag = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 3885128951022111978ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_requester = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 14460765703017969842ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_lower_addr = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 3737248518345978092ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_data_swapped = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 1884363739708437934ull);
    VL_ZERO_RESET_W(128, vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tlast = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__first = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tvalid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17356823098888229906ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__unused_rd_rsp_4dw = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14158449898470532999ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__srst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4292488959391821060ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4927079725973654404ull);
    VL_SCOPED_RAND_RESET_W(134, vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din, __VscopeHash, 14005757747615435669ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_en = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 3786283773102732583ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_en = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 12111194580342211055ull);
    VL_SCOPED_RAND_RESET_W(134, vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout, __VscopeHash, 4254065054523016041ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__full = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 10474434409013400504ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__empty = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 2420307557110027503ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__prog_empty = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_data_count = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 7606395333725309128ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_data_count = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 15502214173104636032ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__prog_full = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 13047320522420947792ull);
    for (int __Vi0 = 0; __Vi0 < 1024; ++__Vi0) {
        VL_SCOPED_RAND_RESET_W(134, vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem[__Vi0], __VscopeHash, 4552261246111517376ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_ptr = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 7509420565499172625ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 1836831664626320755ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 6081495390900961372ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__rst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 11333062838125275937ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 10181691641038787256ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__pcie_id = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 15431100563936522723ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_ready = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17856268900428142844ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_id = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 1859469119429603590ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tag = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 5852495216098401167ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__traffic_class = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 9609458673723037843ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__attributes = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 4063989200599166349ull);
    for (int __Vi0 = 0; __Vi0 < 256; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_fifo[__Vi0] = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 18028871186493383720ull);
    }
    for (int __Vi0 = 0; __Vi0 < 256; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tag_fifo[__Vi0] = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 875566054611768428ull);
    }
    for (int __Vi0 = 0; __Vi0 < 256; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tc_fifo[__Vi0] = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 16716655315766721724ull);
    }
    for (int __Vi0 = 0; __Vi0 < 256; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__attr_fifo[__Vi0] = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 280996066427318383ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__write_ptr = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 18406945927413750248ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__read_ptr = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 9712001818662027772ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_count = VL_SCOPED_RAND_RESET_I(9, __VscopeHash, 5182982631348709598ull);
    VL_SCOPED_RAND_RESET_W(128, vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data, __VscopeHash, 4106315297482566958ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_fifo_full = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 126906578322660825ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_pop = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4294546283607602073ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_can_accept = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 11692303367160648089ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_push = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 10171753536477216083ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__rst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 7690059365784341104ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17765404443339016686ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlps_in_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 18413766410526898609ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlps_in_ready = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_bar = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_addr = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_be = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_data = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_valid = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_rd_en = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 2724691466935730594ull);
    VL_SCOPED_RAND_RESET_W(128, vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tdata, __VscopeHash, 292803266241502453ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tkeepdw = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 4788525290111385710ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tuser = VL_SCOPED_RAND_RESET_I(9, __VscopeHash, 9387931360854932329ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tvalid = 0U;
    ;
    VL_ZERO_RESET_W(128, vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tkeepdw = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlast = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__be_first = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__be_last = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__first_dw = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__addr = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__srst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14685834634794911458ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 16710715816376335393ull);
    VL_SCOPED_RAND_RESET_W(141, vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__din, __VscopeHash, 6604281099596795594ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__wr_en = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 15610702416469863339ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__rd_en = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17855736603401518623ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__dout[0] = 0U;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__dout[1] = 0U;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__dout[2] = 0U;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__dout[3] = 0U;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__dout[4] = 0U;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__full = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__empty = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__prog_empty = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__valid = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h9dcf22b6__3 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h9dcf22b6__2 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h9dcf22b6__1 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h0038da16__1 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h99f8a03b__1 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_hd1ae10b8__0 = 0;
    VL_ZERO_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h511400fb__0);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h9dcf22b6__0 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h0038da16__0 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h99f8a03b__0 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 12817172362774534661ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 6758454371385821257ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_valid = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 17765574644916611430ull);
    for (int __Vi0 = 0; __Vi0 < 7; ++__Vi0) {
        VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ctx[__Vi0], __VscopeHash, 9834274472725206175ull);
    }
    for (int __Vi0 = 0; __Vi0 < 7; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_data[__Vi0] = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 1641662856856891938ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ready = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 16479966212292955377ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 3214204615883013624ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_ctx, __VscopeHash, 7069966123613131099ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 2639979193623536804ull);
    for (int __Vi0 = 0; __Vi0 < 7; ++__Vi0) {
        for (int __Vi1 = 0; __Vi1 < 32; ++__Vi1) {
            VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem[__Vi0][__Vi1], __VscopeHash, 6858423992204020019ull);
        }
    }
    for (int __Vi0 = 0; __Vi0 < 7; ++__Vi0) {
        for (int __Vi1 = 0; __Vi1 < 32; ++__Vi1) {
            vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__data_mem[__Vi0][__Vi1] = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 17749399617485459054ull);
        }
    }
    for (int __Vi0 = 0; __Vi0 < 7; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr[__Vi0] = VL_SCOPED_RAND_RESET_I(5, __VscopeHash, 4956336040666635764ull);
    }
    for (int __Vi0 = 0; __Vi0 < 7; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr[__Vi0] = VL_SCOPED_RAND_RESET_I(5, __VscopeHash, 5648133099910066950ull);
    }
    for (int __Vi0 = 0; __Vi0 < 7; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[__Vi0] = VL_SCOPED_RAND_RESET_I(6, __VscopeHash, 803299143233705248ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rr_ptr = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 982037937091174902ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 8607285621657203718ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5184924676544917955ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__push = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 16815347888616282992ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__scan = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 11531037248243639491ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__index = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 8135600016828842229ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 14178332524652603004ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__rst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 13853070278805325401ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17751004454835934049ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_addr = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 17721809742872435443ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 15213009381949121025ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 742281854378311161ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 10625552151783057154ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx, __VscopeHash, 13108009479100878641ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_addr = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 9559050671644004218ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 13018920188005904527ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx, __VscopeHash, 1822050775297574728ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 13968794311833874934ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 16364693785524139842ull);
    for (int __Vi0 = 0; __Vi0 < 1024; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem[__Vi0] = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 1879061402119085858ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_word = VL_SCOPED_RAND_RESET_I(10, __VscopeHash, 14010960704941080223ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_word = VL_SCOPED_RAND_RESET_I(10, __VscopeHash, 2390154456369695076ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx_d1, __VscopeHash, 13252563380064055013ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_valid_d1 = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 2294618463596330379ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_data_d1 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 6233769887282796986ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__i = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 3606110009478453145ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17274473514162731682ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1607606331957084357ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__crst_falling = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 12603822790675100680ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_ctx, __VscopeHash, 647388229280192684ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 1316651541702023067ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_addr = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 5018956804487725694ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_ready = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5660591340403472251ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 6279897160485420154ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ctx, __VscopeHash, 7402471626819601648ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 1470898242124865980ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ready = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17778388190341586313ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__busy = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 13715894518118174143ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 18008222481279509927ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_addr = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 12819059297684822355ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_ack = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 93029947801653083ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8624314671615411365ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__latency_counter = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 16430190146692931472ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__target_latency = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 6186267863112937171ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx, __VscopeHash, 12186424947219401023ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 3398839208889672995ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_addr = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 1193375400513412513ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_latency = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 3419551982763982988ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 13629808592118501103ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__idle_counter = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 579937614425925685ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter = VL_SCOPED_RAND_RESET_I(19, __VscopeHash, 4286959343746150174ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_expired = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8502244976335014480ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s0 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 2212669786618932485ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s1 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 9812674158045656778ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s2 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 5758546980699966818ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s3 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 3548648969349597349ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 12977989139598971368ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_next_s0 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 4401747417608676111ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_next_s3 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 7900295193560127039ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_counter = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 12009256166162407929ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 6789807858103081770ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__addr_region = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 5976095196241499705ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adjust = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 4695228092850871503ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__is_sequential = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 7347591870852348456ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__corr_weight = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 16791881206399847415ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__correlated_latency = 0U;
    ;
    for (int __Vi0 = 0; __Vi0 < 16; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[__Vi0] = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 12267197994070530747ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 17846601996768874725ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__latency_range = 13U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_latency = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 13679687949906105547ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__base_jitter = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 4784813292563664761ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_adj = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 5130485241660675617ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 16337058534593895849ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__jittered = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 668108554917362555ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__computed_latency = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 9904913569405222238ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_range = 7U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_latency = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 172305930387999504ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_counter = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 12092512540757580383ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_target = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 15131728473859726211ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_pending = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9652697969904091117ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar1__DOT__rst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 11081802774862500893ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar1__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14979736950893149937ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar1__DOT__wr_addr = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 6277204267368199395ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar1__DOT__wr_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 9170712676778711614ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar1__DOT__wr_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 8300389338789340101ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar1__DOT__wr_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 3197464688640389955ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_req_ctx, __VscopeHash, 7082534712726065749ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_req_addr = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 3504108880010869400ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_req_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8296886989652691257ull);
    VL_ZERO_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_rsp_ctx);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_rsp_data = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_rsp_valid = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar1__DOT__intr_req = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar2__DOT__rst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4953066982640007285ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar2__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8851001158670589725ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar2__DOT__wr_addr = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 3695724982617331442ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar2__DOT__wr_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 4185846105946096074ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar2__DOT__wr_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 12215491492734819743ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar2__DOT__wr_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1496749725471319109ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_req_ctx, __VscopeHash, 11899062684195044250ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_req_addr = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 6101399736109892872ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_req_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14325407655044554082ull);
    VL_ZERO_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_rsp_ctx);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_rsp_data = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_rsp_valid = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar2__DOT__intr_req = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar3__DOT__rst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 611345893012404269ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar3__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 6079733654727763109ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar3__DOT__wr_addr = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 2056104148931861243ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar3__DOT__wr_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 17978084027637513773ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar3__DOT__wr_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 128009326595473092ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar3__DOT__wr_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1318474006087336036ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_req_ctx, __VscopeHash, 1327204415307674845ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_req_addr = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 3538630157845865143ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_req_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5051525046933078053ull);
    VL_ZERO_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_rsp_ctx);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_rsp_data = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_rsp_valid = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar3__DOT__intr_req = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar4__DOT__rst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 12929354174499460225ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar4__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4083992169182027090ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar4__DOT__wr_addr = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 4769485863098451803ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar4__DOT__wr_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 14292091248313114127ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar4__DOT__wr_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 12835314699307987322ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar4__DOT__wr_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8583202639986115438ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_req_ctx, __VscopeHash, 17193068982322286308ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_req_addr = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 5960535576661340027ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_req_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 3145712846033432280ull);
    VL_ZERO_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_rsp_ctx);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_rsp_data = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_rsp_valid = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar4__DOT__intr_req = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar5__DOT__rst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4225779787604671497ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar5__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 3039485170297358439ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar5__DOT__wr_addr = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 9327062980429305619ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar5__DOT__wr_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 2961713434484230481ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar5__DOT__wr_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 10276223614739803291ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar5__DOT__wr_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9867452615354510308ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_req_ctx, __VscopeHash, 2488568540218634340ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_req_addr = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 11047487827144459029ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_req_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8121453173957905732ull);
    VL_ZERO_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_rsp_ctx);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_rsp_data = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_rsp_valid = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar5__DOT__intr_req = 0U;
    ;
    vlSelf->__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__valid_first_edge__2__value = 0;
    vlSelf->__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__valid_last_edge__3__value = 0;
    vlSelf->__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__popcount4__4__value = 0;
    vlSelf->__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__popcount4__5__value = 0;
    vlSelf->__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__popcount4__6__value = 0;
    for (int __Vi0 = 0; __Vi0 < 2; ++__Vi0) {
        vlSelf->__VstlTriggered[__Vi0] = 0;
    }
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__cpl_select_ur__0 = 0;
    vlSelf->__VstlDidInit = 0;
    for (int __Vi0 = 0; __Vi0 < 2; ++__Vi0) {
        vlSelf->__VicoTriggered[__Vi0] = 0;
    }
    vlSelf->__VicoDidInit = 0;
    for (int __Vi0 = 0; __Vi0 < 1; ++__Vi0) {
        vlSelf->__VactTriggered[__Vi0] = 0;
    }
    for (int __Vi0 = 0; __Vi0 < 1; ++__Vi0) {
        vlSelf->__VactTriggeredAcc[__Vi0] = 0;
    }
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__cpl_select_ur__1 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_bar0__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__clk__0 = 0;
    for (int __Vi0 = 0; __Vi0 < 1; ++__Vi0) {
        vlSelf->__VnbaTriggered[__Vi0] = 0;
    }
    vlSelf->__Vi = 0;
}
