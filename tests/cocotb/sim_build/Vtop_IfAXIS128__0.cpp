// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vtop.h for the primary calling header

#include "Vtop__pch.h"

void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__tlps_in_if__0(Vtop_IfAXIS128* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+            Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__tlps_in_if__0\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tkeepdw = vlSymsp->TOP.tb_top__DOT__tlps_in_tkeepdw;
    vlSelfRef.tvalid = vlSymsp->TOP.tb_top__DOT__tlps_in_tvalid;
    vlSelfRef.tlast = vlSymsp->TOP.tb_top__DOT__tlps_in_tlast;
    vlSelfRef.tuser = vlSymsp->TOP.tb_top__DOT__tlps_in_tuser;
    vlSelfRef.tdata[0U] = vlSymsp->TOP.tb_top__DOT__tlps_in_tdata[0U];
    vlSelfRef.tdata[1U] = vlSymsp->TOP.tb_top__DOT__tlps_in_tdata[1U];
    vlSelfRef.tdata[2U] = vlSymsp->TOP.tb_top__DOT__tlps_in_tdata[2U];
    vlSelfRef.tdata[3U] = vlSymsp->TOP.tb_top__DOT__tlps_in_tdata[3U];
}

void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__tlps_in_if__1(Vtop_IfAXIS128* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+            Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__tlps_in_if__1\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.__VdfgRegularize_hebeb780c_0_3 = ((IData)(vlSelfRef.tvalid) 
                                                & (IData)(vlSymsp->TOP.tb_top__DOT__i_bar__DOT__in_is_bar));
}

void Vtop_IfAXIS128___eval_initial__TOP__tb_top__DOT__tlps_out_if(Vtop_IfAXIS128* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+            Vtop_IfAXIS128___eval_initial__TOP__tb_top__DOT__tlps_out_if\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tready = 1U;
}

void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__tlps_out_if__0(Vtop_IfAXIS128* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+            Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__tlps_out_if__0\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tdata[0U] = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_cpl.tdata[0U];
    vlSelfRef.tdata[1U] = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_cpl.tdata[1U];
    vlSelfRef.tdata[2U] = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_cpl.tdata[2U];
    vlSelfRef.tdata[3U] = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_cpl.tdata[3U];
    vlSelfRef.tkeepdw = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_cpl.tkeepdw;
    vlSelfRef.tuser = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_cpl.tuser;
    vlSelfRef.has_data = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_cpl.has_data;
    vlSelfRef.tvalid = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_cpl.tvalid;
    vlSelfRef.tlast = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_cpl.tlast;
}

void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__tlps_dma_out_if__0(Vtop_IfAXIS128* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+            Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__tlps_dma_out_if__0\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tkeepdw = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tkeepdw;
    vlSelfRef.tdata[0U] = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tdata[0U];
    vlSelfRef.tdata[1U] = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tdata[1U];
    vlSelfRef.tdata[2U] = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tdata[2U];
    vlSelfRef.tdata[3U] = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tdata[3U];
    vlSelfRef.tlast = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tlast;
    vlSelfRef.tuser = (((IData)(vlSymsp->TOP.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tlast) 
                        << 1U) | (IData)(vlSymsp->TOP.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_first));
}

void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__tlps_dma_out_if__0(Vtop_IfAXIS128* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+            Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__tlps_dma_out_if__0\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tvalid = ((IData)(vlSymsp->TOP.tb_top__DOT__i_bar__DOT__dma_enabled) 
                        & (IData)(vlSymsp->TOP.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tvalid));
    vlSelfRef.has_data = ((IData)(vlSymsp->TOP.tb_top__DOT__i_bar__DOT__dma_enabled) 
                          & (0U != (IData)(vlSymsp->TOP.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_count_next)));
}

void Vtop_IfAXIS128___nba_comb__TOP__tb_top__DOT__tlps_dma_out_if__0(Vtop_IfAXIS128* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+            Vtop_IfAXIS128___nba_comb__TOP__tb_top__DOT__tlps_dma_out_if__0\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tvalid = ((IData)(vlSymsp->TOP.tb_top__DOT__i_bar__DOT__dma_enabled) 
                        & (IData)(vlSymsp->TOP.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tvalid));
}

void Vtop_IfAXIS128___nba_comb__TOP__tb_top__DOT__tlps_dma_out_if__1(Vtop_IfAXIS128* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+            Vtop_IfAXIS128___nba_comb__TOP__tb_top__DOT__tlps_dma_out_if__1\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.has_data = ((IData)(vlSymsp->TOP.tb_top__DOT__i_bar__DOT__dma_enabled) 
                          & (0U != (IData)(vlSymsp->TOP.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_count_next)));
}

void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__0(Vtop_IfAXIS128* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+            Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__0\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tdata[0U] = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[0U];
    vlSelfRef.tdata[1U] = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[1U];
    vlSelfRef.tdata[2U] = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[2U];
    vlSelfRef.tdata[3U] = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[3U];
    vlSelfRef.tkeepdw = (0x0000000fU & vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[4U]);
    vlSelfRef.tlast = (1U & (vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[4U] 
                             >> 4U));
    vlSelfRef.tuser = (((IData)(vlSelfRef.tlast) << 1U) 
                       | (1U & (vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[4U] 
                                >> 5U)));
}

void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__0(Vtop_IfAXIS128* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+            Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__0\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tready = ((IData)(vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_cpl.tready) 
                        & (~ (IData)(vlSymsp->TOP.tb_top__DOT__i_bar__DOT__cpl_select_ur)));
}

void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__1(Vtop_IfAXIS128* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+            Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__1\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tvalid = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__valid;
    vlSelfRef.has_data = vlSelfRef.tvalid;
}

void Vtop_IfAXIS128___nba_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__0(Vtop_IfAXIS128* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+            Vtop_IfAXIS128___nba_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__0\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tdata[0U] = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[0U];
    vlSelfRef.tdata[1U] = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[1U];
    vlSelfRef.tdata[2U] = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[2U];
    vlSelfRef.tdata[3U] = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[3U];
    vlSelfRef.tkeepdw = (0x0000000fU & vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[4U]);
    vlSelfRef.tlast = (1U & (vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[4U] 
                             >> 4U));
    vlSelfRef.tvalid = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__valid;
    vlSelfRef.tuser = (((IData)(vlSelfRef.tlast) << 1U) 
                       | (1U & (vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[4U] 
                                >> 5U)));
    vlSelfRef.has_data = vlSelfRef.tvalid;
}

void Vtop_IfAXIS128___eval_initial__TOP__tb_top__DOT__i_bar__DOT__tlps_ur(Vtop_IfAXIS128* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+            Vtop_IfAXIS128___eval_initial__TOP__tb_top__DOT__i_bar__DOT__tlps_ur\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tlast = 1U;
    vlSelfRef.tkeepdw = 7U;
    vlSelfRef.tuser = 3U;
}

void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_ur__0(Vtop_IfAXIS128* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+            Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_ur__0\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.has_data = (0U != (IData)(vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_count));
    vlSelfRef.tvalid = (0U != (IData)(vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_count));
    vlSelfRef.tdata[0U] = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[0U];
    vlSelfRef.tdata[1U] = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[1U];
    vlSelfRef.tdata[2U] = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[2U];
    vlSelfRef.tdata[3U] = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[3U];
}

void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_ur__0(Vtop_IfAXIS128* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+            Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_ur__0\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tready = ((IData)(vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_cpl.tready) 
                        & (IData)(vlSymsp->TOP.tb_top__DOT__i_bar__DOT__cpl_select_ur));
}

void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_cpl__0(Vtop_IfAXIS128* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+            Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_cpl__0\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tready = vlSymsp->TOP__tb_top__DOT__tlps_out_if.tready;
}

void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_cpl__0(Vtop_IfAXIS128* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+            Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_cpl__0\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    if (vlSymsp->TOP.tb_top__DOT__i_bar__DOT__cpl_select_ur) {
        vlSelfRef.tdata[0U] = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_ur.tdata[0U];
        vlSelfRef.tdata[1U] = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_ur.tdata[1U];
        vlSelfRef.tdata[2U] = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_ur.tdata[2U];
        vlSelfRef.tdata[3U] = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_ur.tdata[3U];
        vlSelfRef.tkeepdw = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_ur.tkeepdw;
        vlSelfRef.tuser = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_ur.tuser;
        vlSelfRef.has_data = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_ur.has_data;
    } else {
        vlSelfRef.tdata[0U] = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng.tdata[0U];
        vlSelfRef.tdata[1U] = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng.tdata[1U];
        vlSelfRef.tdata[2U] = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng.tdata[2U];
        vlSelfRef.tdata[3U] = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng.tdata[3U];
        vlSelfRef.tkeepdw = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng.tkeepdw;
        vlSelfRef.tuser = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng.tuser;
        vlSelfRef.has_data = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng.has_data;
    }
    vlSelfRef.tvalid = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__cpl_selected_valid;
    vlSelfRef.tlast = vlSymsp->TOP.tb_top__DOT__i_bar__DOT__cpl_selected_last;
}

std::string VL_TO_STRING(const Vtop_IfAXIS128* obj) {
    VL_DEBUG_IF(VL_DBG_MSGF("+            Vtop_IfAXIS128::VL_TO_STRING\n"); );
    // Body
    return (obj ? obj->vlNamep : "null");
}
