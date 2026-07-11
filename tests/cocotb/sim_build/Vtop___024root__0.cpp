// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vtop.h for the primary calling header

#include "Vtop__pch.h"

VL_ATTR_COLD void Vtop___024root___eval_initial__TOP(Vtop___024root* vlSelf);
VlCoroutine Vtop___024root___eval_initial__TOP__Vtiming__0(Vtop___024root* vlSelf);
void Vtop_IfAXIS128___eval_initial__TOP__tb_top__DOT__tlps_out_if(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___eval_initial__TOP__tb_top__DOT__i_bar__DOT__tlps_ur(Vtop_IfAXIS128* vlSelf);

void Vtop___024root___eval_initial(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_initial\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    Vtop___024root___eval_initial__TOP(vlSelf);
    Vtop___024root___eval_initial__TOP__Vtiming__0(vlSelf);
    Vtop_IfAXIS128___eval_initial__TOP__tb_top__DOT__tlps_out_if((&vlSymsp->TOP__tb_top__DOT__tlps_out_if));
    Vtop_IfAXIS128___eval_initial__TOP__tb_top__DOT__tlps_out_if((&vlSymsp->TOP__tb_top__DOT__tlps_dma_out_if));
    Vtop_IfAXIS128___eval_initial__TOP__tb_top__DOT__i_bar__DOT__tlps_ur((&vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_ur));
}

VlCoroutine Vtop___024root___eval_initial__TOP__Vtiming__0(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_initial__TOP__Vtiming__0\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    while (VL_LIKELY(!vlSymsp->_vm_contextp__->gotFinish())) {
        co_await vlSelfRef.__VdlySched.delay(0x0000000000001388ULL, 
                                             nullptr, 
                                             "tb_top.sv", 
                                             7);
        vlSelfRef.tb_top__DOT__clk = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__clk)));
    }
    co_return;
}

void Vtop___024root___eval_triggers_vec__ico(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_triggers_vec__ico\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.__VicoTriggered[1U] = ((0xfffffffffffffffeULL 
                                      & vlSelfRef.__VicoTriggered[1U]) 
                                     | (IData)((IData)(vlSelfRef.__VicoFirstIteration)));
    vlSelfRef.__VicoTriggered[0U] = (QData)((IData)(
                                                    ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_select_ur) 
                                                     != (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__cpl_select_ur__0))));
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__cpl_select_ur__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_select_ur;
    if (VL_UNLIKELY(((1U & (~ (IData)(vlSelfRef.__VicoDidInit)))))) {
        vlSelfRef.__VicoDidInit = 1U;
        vlSelfRef.__VicoTriggered[0U] = (1ULL | vlSelfRef.__VicoTriggered[0U]);
    }
}

bool Vtop___024root___trigger_anySet__ico(const VlUnpacked<QData/*63:0*/, 2> &in) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___trigger_anySet__ico\n"); );
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

void Vtop___024root___ico_sequent__TOP__0(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___ico_sequent__TOP__0\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_ready 
        = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__lifecycle_generation 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__generation;
    vlSelfRef.tb_top__DOT__tlps_dma_out_tvalid = vlSymsp->TOP__tb_top__DOT__tlps_dma_out_if.tvalid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__full 
        = (0x0400U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_data_count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_data_count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_last_packet 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw) 
           == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_expired 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending) 
           & (0x00010000U <= vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_id 
        = (0x0000ffffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_fifo
                          [vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_read_ptr] 
                          >> 0x0000000eU));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tag 
        = (0x000000ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_fifo
                          [vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_read_ptr] 
                          >> 6U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__traffic_class 
        = (7U & (vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_fifo
                 [vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_read_ptr] 
                 >> 3U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__attributes 
        = (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_fifo
           [vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_read_ptr]);
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[3U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[3U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[4U] 
        = (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__first) 
            << 5U) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tlast) 
                       << 4U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_valid 
        = (0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx_w[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx
        [vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_rd][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx_w[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx
        [vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_rd][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx_w[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx
        [vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_rd][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_data_w 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_data
        [vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_rd];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tvalid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tlast) 
           | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw) 
              >> 3U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_next_s0 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s0 
           ^ VL_SHIFTL_III(32,32,32, vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s0, 0x0000000bU));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_full 
        = (0x10U <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_rd_en 
        = ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state)) 
           | ((2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state)) 
              | ((7U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state)) 
                 | (((~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tkeepdw) 
                         >> 3U)) & (6U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state))) 
                    | (((~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tkeepdw) 
                            >> 1U)) & (4U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state))) 
                       | ((~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tkeepdw) 
                              >> 2U)) & (5U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state))))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_empty 
        = (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_fifo_full 
        = (0x0100U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_pop 
        = (1U & (~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__state) 
                    | (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_count)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ready 
        = ((((((0x20U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[6U]) 
               << 3U) | ((0x20U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[5U]) 
                         << 2U)) | (((0x20U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[4U]) 
                                     << 1U) | (0x20U 
                                               > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[3U]))) 
            << 3U) | (((0x20U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[2U]) 
                       << 2U) | (((0x20U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[1U]) 
                                  << 1U) | (0x20U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[0U]))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__busy 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[0U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[1U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[2U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[3U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[0U] 
        = (0x0a000000U | (0x00ffffffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[0U]));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[0U] 
        = ((0xff8fffffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[0U]) 
           | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tc_fifo
                      [vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__read_ptr]) 
              << 0x00000014U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[0U] 
        = ((0xfffbffffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[0U]) 
           | (0x00040000U & ((vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__attr_fifo
                              [vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__read_ptr] 
                              >> 2U) << 0x00000012U)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[0U] 
        = ((0xffffcfffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[0U]) 
           | (0x00003000U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__attr_fifo
                             [vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__read_ptr] 
                             << 0x0000000cU)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[1U] = 0x00002000U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[2U] 
        = (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_fifo
                    [vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__read_ptr]) 
            << 0x00000010U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tag_fifo
                                       [vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__read_ptr]) 
                               << 8U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_be = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_wr_ack 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_ack;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_fifo_full 
        = (0x0100U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_bar 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_bar;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s0 
           + vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s3);
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_bar = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_bar;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_last 
        = (1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_remaining_dw));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_addr 
        = (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr);
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_almost_full 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_valid) 
           | (0x0eU <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_count)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__clk = vlSelfRef.tb_top__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be 
        = (0x0000000fU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_packet)
                           ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_be)
                           : ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_last_be) 
                              | (- (IData)((1U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw)))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_addr_w 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_addr
        [vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_rd];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem
        [(0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr))][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem
        [(0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr))][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem
        [(0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr))][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[3U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem
        [(0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr))][3U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[4U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem
        [(0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr))][4U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[1U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_rsp_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[1U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_rsp_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[1U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_rsp_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[2U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_rsp_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[2U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_rsp_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[2U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_rsp_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[3U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_rsp_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[3U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_rsp_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[3U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_rsp_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[4U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_rsp_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[4U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_rsp_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[4U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_rsp_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[5U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_rsp_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[5U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_rsp_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[5U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_rsp_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_rsp_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_rsp_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[3U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_rsp_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[4U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_rsp_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[5U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_rsp_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__rst = vlSelfRef.tb_top__DOT__rst;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected_valid = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__scan = 0U;
    while (VL_GTS_III(32, 7U, vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__scan)) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__index 
            = (0x0000000fU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rr_ptr) 
                              + vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__scan));
        if ((7U <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__index))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__index 
                = (0x0000000fU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__index) 
                                  - (IData)(7U)));
        }
        if (((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected_valid)) 
             & (0U != ((6U >= (7U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__index)))
                        ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count
                       [(7U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__index))]
                        : 0U)))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected 
                = (7U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__index));
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected_valid = 1U;
        }
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__scan 
            = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__scan);
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__empty 
        = (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_req_addr_from_ctx 
        = (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[0U]);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx_w[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx_w[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx_w[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_data_w;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_en 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tvalid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_next_s3 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s3 
           ^ (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_next_s0 
              ^ (VL_SHIFTR_III(32,32,32, vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s3, 0x00000013U) 
                 ^ VL_SHIFTR_III(32,32,32, vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_next_s0, 8U))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_write 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_full)) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_valid) 
              | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_valid)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__rd_en 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_rd_en;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_valid 
        = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_empty)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_can_accept 
        = (1U & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_fifo_full)) 
                 | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_pop)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ready 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ready;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__busy;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__wr_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__wr_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__wr_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__wr_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__wr_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__wr_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__wr_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__wr_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__wr_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__wr_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr_ready 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_wr_ack;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_bar;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_latency 
        = (0x000000ffU & ((IData)(2U) + VL_MODDIV_III(32, 
                                                      (0x000000ffU 
                                                       & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out 
                                                          >> 8U)), (IData)(7U))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 0x0fU;
    if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
         < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 0U;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[1U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 1U;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[2U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 2U;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[3U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 3U;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[4U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 4U;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[5U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 5U;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[6U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 6U;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[7U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 7U;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[8U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 8U;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[9U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 9U;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[10U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 0x0aU;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[11U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 0x0bU;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[12U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 0x0cU;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[13U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 0x0dU;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[14U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 0x0eU;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__wr_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_bar) 
              >> 1U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__wr_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_bar) 
              >> 2U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__wr_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_bar) 
              >> 3U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__wr_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_bar) 
              >> 4U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__wr_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_bar) 
              >> 5U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__stall 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_almost_full;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_wr_hit 
        = (0x00004000U > vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first_offset 
        = (3U & (((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be))
                   ? 1U : ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be))
                            ? 2U : (- (IData)((1U & 
                                               ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be) 
                                                >> 3U)))))) 
                 & (- (IData)((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be)))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_addr_w;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__fundamental_reset 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rst;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected_valid) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_valid = 1U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_data 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__data_mem
            [((6U >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected))
               ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected)
               : 0U)][((6U >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected))
                        ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr
                       [vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected]
                        : 0U)];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_ctx[0U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem
            [((6U >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected))
               ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected)
               : 0U)][((6U >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected))
                        ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr
                       [vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected]
                        : 0U)][0U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_ctx[1U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem
            [((6U >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected))
               ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected)
               : 0U)][((6U >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected))
                        ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr
                       [vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected]
                        : 0U)][1U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_ctx[2U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem
            [((6U >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected))
               ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected)
               : 0U)][((6U >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected))
                        ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr
                       [vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected]
                        : 0U)][2U];
    } else {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_valid = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_data = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_ctx[0U] = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_ctx[1U] = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_ctx[2U] = 0U;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ready 
        = (1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ready));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ready 
        = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_accepted 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy_d1)) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_valid 
        = ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_rsp_valid) 
             << 5U) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_rsp_valid) 
                        << 4U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_rsp_valid) 
                                  << 3U))) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_rsp_valid) 
                                               << 2U) 
                                              | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_rsp_valid) 
                                                  << 1U) 
                                                 | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_valid))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_latency 
        = (0x000000ffU & ((IData)(3U) + (((IData)(0x0000000dU) 
                                          * (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket)) 
                                         >> 4U)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_rd_hit 
        = (0x00004000U > vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr);
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr_bar0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__issue_dw 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__stall)) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__state));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0;
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_wr_hit) 
              & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_bar)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__context_address 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first)
            ? (((IData)((vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr 
                         >> 2U)) << 2U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first_offset))
            : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__addr_region 
        = (7U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_addr 
                 >> 0x0000000cU));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__is_sequential 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_valid) 
           & (((vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_addr 
                >> 8U) == (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_addr 
                           >> 8U)) | ((vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_addr 
                                       >> 8U) == (0x00ffffffU 
                                                  & ((IData)(1U) 
                                                     + 
                                                     (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_addr 
                                                      >> 8U))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[0U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[0U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[0U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__device_reset 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__fundamental_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_rsp_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_rsp_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_rsp_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_rsp_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_rsp_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_pop 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_empty)) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_accepted));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_valid 
        = (0x0000003fU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_valid));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr_bar0;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__issue_dw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_word 
        = (0x000003ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_addr 
                          >> 2U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__context_address;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[1U] 
        = ((0xf0000000U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[1U]) 
           | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_attr) 
               << 0x00000019U) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_4dw) 
                                   << 0x00000018U) 
                                  | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_tag) 
                                      << 0x00000010U) 
                                     | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_id)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[1U] 
        = ((0x0fffffffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[1U]) 
           | (((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first) 
                 << 0x0000001bU) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_last) 
                                     << 0x0000001aU) 
                                    | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw) 
                                       << 0x0000000fU))) 
               | ((0x00007ff8U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_byte_count) 
                                  << 3U)) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_tc))) 
              << 0x0000001cU));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[2U] 
        = (0x00ffffffU & (((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first) 
                             << 0x0000001bU) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_last) 
                                                 << 0x0000001aU) 
                                                | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw) 
                                                   << 0x0000000fU))) 
                           | ((0x00007ff8U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_byte_count) 
                                              << 3U)) 
                              | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_tc))) 
                          >> 4U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adjust 
        = (((1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__addr_region))
             ? 2U : ((2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__addr_region))
                      ? 4U : ((3U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__addr_region))
                               ? 1U : 3U))) & (- (IData)(
                                                         (0U 
                                                          != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__addr_region)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__corr_weight 
        = (0x80U & (- (IData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__is_sequential))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__base_jitter 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_latency) 
           & (- (IData)((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__is_sequential))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[0U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[0U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[0U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[0U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[0U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[0U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[1U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[1U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[1U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[1U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[1U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[1U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[2U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[2U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[2U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[2U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[2U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[2U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[3U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[3U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[3U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[3U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[3U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[3U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[4U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[4U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[4U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[4U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[4U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[4U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[5U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[5U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[5U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[5U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[5U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[5U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[3U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[3U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[4U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[4U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[5U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[5U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__device_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__active 
        = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__device_reset)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_rsp_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_rsp_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_rsp_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_rsp_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_rsp_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_word 
        = (0x000003ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_addr 
                          >> 2U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_adj 
        = (0x000000ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__base_jitter) 
                          + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset)));
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[0U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[0U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[0U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[0U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[0U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[0U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[1U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[1U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[1U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[1U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[1U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[1U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[2U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[2U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[2U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[2U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[2U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[2U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[3U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[3U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[3U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[3U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[3U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[3U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[4U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[4U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[4U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[4U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[4U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[4U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[5U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[5U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[5U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[5U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[5U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[5U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[6U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[6U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[6U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[6U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[6U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[6U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data[3U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[3U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data[4U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[4U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data[5U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[5U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data[6U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[6U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__rst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__rst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__dma_enabled 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__active;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__io_enabled 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__active;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_data_swapped 
        = ((((0x0000ff00U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_data 
                             << 8U)) | (0x000000ffU 
                                        & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_data 
                                           >> 8U))) 
            << 0x00000010U) | ((0x0000ff00U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_data 
                                               >> 8U)) 
                               | (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_data 
                                  >> 0x00000018U)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_first 
        = (1U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[2U] 
                 >> 0x00000017U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_last 
        = (1U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[2U] 
                 >> 0x00000016U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_dwlen 
        = (0x000007ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[2U] 
                          >> 0x0000000bU));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_byte_count 
        = (0x00000fffU & ((vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[2U] 
                           << 1U) | (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[1U] 
                                     >> 0x0000001fU)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_tc 
        = (7U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[1U] 
                 >> 0x0000001cU));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_attr 
        = (7U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[1U] 
                 >> 0x00000019U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_tag 
        = (0x000000ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[1U] 
                          >> 0x00000010U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_requester 
        = (0x0000ffffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[1U]);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_lower_addr 
        = (0x0000007fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[0U]);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_4dw 
        = (1U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[1U] 
                 >> 0x00000018U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__push 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_valid) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ready));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_req_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar) 
              >> 1U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_req_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar) 
              >> 2U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_req_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar) 
              >> 3U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_req_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar) 
              >> 4U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_req_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar) 
              >> 5U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_rd_hit) 
              & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_req_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_req_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_req_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_req_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_req_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj 
        = (0x000000ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adjust) 
                          + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_adj)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ctx 
        = vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__srst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rst;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__srst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__rst;
    vlSelfRef.tb_top__DOT__i_bar__DOT__dma_enabled 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__dma_enabled;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__quiesce 
        = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__dma_enabled)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__io_enabled = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__io_enabled;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__unused_rd_rsp_4dw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_4dw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__jittered 
        = ((0x0000000fU < ((IData)(3U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj)))
            ? 0x0000000fU : ((3U > ((IData)(3U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj)))
                              ? 3U : (0x000000ffU & 
                                      ((IData)(3U) 
                                       + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__lifecycle_quiesce 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__quiesce;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__computed_latency 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__jittered;
}

void Vtop___024root___ico_sequent__TOP__1(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___ico_sequent__TOP__1\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    CData/*1:0*/ tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT____VlemCall_7__first_enabled_offset;
    CData/*0:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__onehot7__1__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__onehot7__1__Vfuncout = 0;
    CData/*6:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__onehot7__1__value;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__onehot7__1__value = 0;
    CData/*3:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_enabled_offset__7__value;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_enabled_offset__7__value = 0;
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_last 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tlast;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[0U] 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tdata[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[1U] 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tdata[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[2U] 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tdata[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[3U] 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tdata[3U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__din[0U] 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tdata[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__din[1U] 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tdata[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__din[2U] 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tdata[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__din[3U] 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tdata[3U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__din[4U] 
        = (((IData)(vlSymsp->TOP__tb_top__DOT__tlps_in_if.tuser) 
            << 4U) | (IData)(vlSymsp->TOP__tb_top__DOT__tlps_in_if.tkeepdw));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tuser;
    vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_bar = 
        ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__io_enabled) 
         & (0U != (0x0000007fU & ((IData)(vlSymsp->TOP__tb_top__DOT__tlps_in_if.tuser) 
                                  >> 2U))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__requester_id 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[1U] 
           >> 0x00000010U);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tag 
        = (0x000000ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[1U] 
                          >> 8U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__traffic_class 
        = (7U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[0U] 
                 >> 0x00000014U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__attributes 
        = ((4U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[0U] 
                  >> 0x00000010U)) | (3U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[0U] 
                                            >> 0x0000000cU)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__bir 
        = (7U & ((0x00000100U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user))
                  ? (6U | ((- (IData)((1U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                             >> 6U)))) 
                           | ((- (IData)((1U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                                >> 5U)))) 
                              | ((- (IData)((1U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                                   >> 4U)))) 
                                 | ((- (IData)((1U 
                                                & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                                   >> 3U)))) 
                                    | ((- (IData)((1U 
                                                   & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                                      >> 2U)))) 
                                       | (- (IData)(
                                                    (1U 
                                                     & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                                        >> 7U))))))))))
                  : ((0x00000080U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user))
                      ? (5U | ((- (IData)((1U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                                 >> 5U)))) 
                               | ((- (IData)((1U & 
                                              ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                               >> 4U)))) 
                                  | ((- (IData)((1U 
                                                 & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                                    >> 3U)))) 
                                     | ((- (IData)(
                                                   (1U 
                                                    & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                                       >> 2U)))) 
                                        | (- (IData)(
                                                     (1U 
                                                      & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                                         >> 6U)))))))))
                      : ((0x00000040U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user))
                          ? (4U | ((- (IData)((1U & 
                                               ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                                >> 4U)))) 
                                   | ((- (IData)((1U 
                                                  & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                                     >> 3U)))) 
                                      | ((- (IData)(
                                                    (1U 
                                                     & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                                        >> 2U)))) 
                                         | (- (IData)(
                                                      (1U 
                                                       & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                                          >> 5U))))))))
                          : ((0x00000020U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user))
                              ? (3U | ((- (IData)((1U 
                                                   & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                                      >> 3U)))) 
                                       | ((- (IData)(
                                                     (1U 
                                                      & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                                         >> 2U)))) 
                                          | (- (IData)(
                                                       (1U 
                                                        & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                                           >> 4U)))))))
                              : ((0x00000010U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user))
                                  ? (2U | ((- (IData)(
                                                      (1U 
                                                       & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                                          >> 2U)))) 
                                           | (- (IData)(
                                                        (1U 
                                                         & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                                            >> 3U))))))
                                  : ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user))
                                      ? (1U | (- (IData)(
                                                         (1U 
                                                          & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                                             >> 2U)))))
                                      : (- (IData)(
                                                   (1U 
                                                    & (~ 
                                                       ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                                                        >> 2U))))))))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__fmt 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[0U] 
           >> 0x0000001dU);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type 
        = (0x0000001fU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[0U] 
                          >> 0x00000018U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_present 
        = (1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_supported = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__unsupported_request = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__ur_required = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__non_posted_request = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__malformed_request = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__posted_write = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__memory_read = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__memory_write = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__header_4dw 
        = (1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__fmt));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__length_dw 
        = ((0U == (0x000003ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[0U]))
            ? 0x0400U : (0x000003ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[0U]));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_be 
        = (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[1U]);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__last_be 
        = (0x0000000fU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[1U] 
                          >> 4U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__bar_mask 
        = (0x0000007fU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user) 
                          >> 2U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__address 
        = ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__fmt))
            ? (((QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[2U])) 
                << 0x00000020U) | (QData)((IData)((0xfffffffcU 
                                                   & vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[3U]))))
            : ((QData)((IData)((vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[2U] 
                                >> 2U))) << 2U));
    if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__fmt))) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = 0U;
    } else if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__fmt))) {
        if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__fmt))) {
            if ((0x00000010U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = 0U;
            } else if ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind 
                    = ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))
                        ? ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))
                            ? ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))
                                ? 0U : 5U) : 5U) : 0U);
            } else if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = 0U;
            } else if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = 0U;
            } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = 0U;
            } else {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = 2U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__memory_write = 1U;
            }
        } else if ((0x00000010U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = 0U;
        } else if ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind 
                = ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))
                    ? ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))
                        ? ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))
                            ? 0U : 5U) : 5U) : 0U);
        } else if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind 
                = ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))
                    ? 0U : 7U);
        } else if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind 
                = ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))
                    ? 0U : 4U);
        } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = 0U;
        } else {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = 2U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__memory_write = 1U;
        }
    } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__fmt))) {
        if ((0x00000010U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = 0U;
        } else if ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = 0U;
        } else if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = 0U;
        } else if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = 0U;
        } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = 6U;
        } else {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = 1U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__memory_read = 1U;
        }
    } else if ((0x00000010U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = 0U;
    } else if ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = 0U;
    } else if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind 
            = ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))
                ? 0U : 7U);
    } else if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind 
            = ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))
                ? 0U : 3U);
    } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_type))) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = 6U;
    } else {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind = 1U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__memory_read = 1U;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__non_posted_request 
        = (((((3U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind)) 
              | (4U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind))) 
             | (5U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind))) 
            | (6U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind))) 
           | (7U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind)));
    __Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__onehot7__1__value 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__bar_mask;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__onehot7__1__Vfuncout 
        = ((0U != (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__onehot7__1__value)) 
           & (0U == ((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__onehot7__1__value) 
                     & ((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__onehot7__1__value) 
                        - (IData)(1U)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__legal_bar 
        = __Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__onehot7__1__Vfuncout;
    if ((1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__length_dw))) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT____VlemCond_2 
            = (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__last_be));
        vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__popcount4__4__value 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_be;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT____VlemCall_3__popcount4 
            = (7U & ((((1U & (IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__popcount4__4__value)) 
                       + (1U & ((IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__popcount4__4__value) 
                                >> 1U))) + (1U & ((IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__popcount4__4__value) 
                                                  >> 2U))) 
                     + (1U & ((IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__popcount4__4__value) 
                              >> 3U))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT____VlemCond_6 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT____VlemCall_3__popcount4;
    } else {
        if (((2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__length_dw)) 
             & (0U == (7U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__address))))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT____VlemCond_2 
                = ((0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_be)) 
                   & (0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__last_be)));
        } else {
            vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__valid_first_edge__2__value 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_be;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT____VlemCall_0__valid_first_edge 
                = ((((8U == (IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__valid_first_edge__2__value)) 
                     | (0x0cU == (IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__valid_first_edge__2__value))) 
                    | (0x0eU == (IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__valid_first_edge__2__value))) 
                   | (0x0fU == (IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__valid_first_edge__2__value)));
            vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__valid_last_edge__3__value 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__last_be;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT____VlemCall_1__valid_last_edge 
                = ((((1U == (IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__valid_last_edge__3__value)) 
                     | (3U == (IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__valid_last_edge__3__value))) 
                    | (7U == (IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__valid_last_edge__3__value))) 
                   | (0x0fU == (IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__valid_last_edge__3__value)));
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT____VlemCond_2 
                = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT____VlemCall_0__valid_first_edge) 
                   & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT____VlemCall_1__valid_last_edge));
        }
        vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__popcount4__5__value 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_be;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT____VlemCall_4__popcount4 
            = (7U & ((((1U & (IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__popcount4__5__value)) 
                       + (1U & ((IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__popcount4__5__value) 
                                >> 1U))) + (1U & ((IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__popcount4__5__value) 
                                                  >> 2U))) 
                     + (1U & ((IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__popcount4__5__value) 
                              >> 3U))));
        vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__popcount4__6__value 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__last_be;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT____VlemCall_5__popcount4 
            = (7U & ((((1U & (IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__popcount4__6__value)) 
                       + (1U & ((IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__popcount4__6__value) 
                                >> 1U))) + (1U & ((IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__popcount4__6__value) 
                                                  >> 2U))) 
                     + (1U & ((IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__popcount4__6__value) 
                              >> 3U))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT____VlemCond_6 
            = (0x00001fffU & (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT____VlemCall_4__popcount4) 
                               + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT____VlemCall_5__popcount4)) 
                              + (0x00001ffcU & (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__length_dw) 
                                                 - (IData)(2U)) 
                                                << 2U))));
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__legal_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT____VlemCond_2;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__crosses_4k 
        = (0x1000U < (0x00001fffU & ((0x00000fffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__address)) 
                                     + VL_SHIFTL_III(13,13,32, (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__length_dw), 2U))));
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_present) {
        if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__memory_read) 
             | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__memory_write))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__malformed_request 
                = (1U & (((((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__legal_bar)) 
                            | (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__legal_be))) 
                           | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__crosses_4k)) 
                          | ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__header_4dw)) 
                             & (0U != (IData)((vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__address 
                                               >> 0x20U))))) 
                         | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__memory_read) 
                            & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_last)))));
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_supported 
                = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__malformed_request)));
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__posted_write 
                = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__memory_write) 
                   & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__malformed_request)));
        } else {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__unsupported_request = 1U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__ur_required = 1U;
        }
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__enabled_byte_count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT____VlemCond_6;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_completion_byte_count 
        = ((((1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__length_dw)) 
             & (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_be))) 
            & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__memory_read))
            ? 4U : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__enabled_byte_count));
    __Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_enabled_offset__7__value 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_be;
    tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT____VlemCall_7__first_enabled_offset 
        = ((1U & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_enabled_offset__7__value))
            ? 0U : ((2U & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_enabled_offset__7__value))
                     ? 1U : ((4U & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_enabled_offset__7__value))
                              ? 2U : ((8U & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_enabled_offset__7__value))
                                       ? 3U : 0U))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_lower_address 
        = (0x0000007fU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__address) 
                          + (IData)(tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT____VlemCall_7__first_enabled_offset)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__rcb_bytes_left 
        = (0x00001fffU & ((IData)(0x0040U) - (0x003fU 
                                              & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__address))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__rcb_dw_left 
        = (0x000007ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__rcb_bytes_left) 
                          >> 2U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__mps_dw = 0x0020U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_completion_dw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__length_dw;
    if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_completion_dw) 
         > (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__rcb_dw_left))) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_completion_dw 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__rcb_dw_left;
    }
    if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_completion_dw) 
         > (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__mps_dw))) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_completion_dw 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__mps_dw;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_requester_id 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__requester_id;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_tag = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tag;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_traffic_class 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__traffic_class;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_attributes 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__attributes;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_bir = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__bir;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_unsupported_request 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__unsupported_request;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_malformed_request 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__malformed_request;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_memory_write 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__memory_write;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_request_kind 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_kind;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_first_lower_address 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_lower_address;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_header_4dw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__header_4dw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_length_dw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__length_dw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_first_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_last_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__last_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_bar_mask 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__bar_mask;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_address 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__address;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_first_completion_byte_count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_completion_byte_count;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_first_completion_dw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__first_completion_dw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_ur_required 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__ur_required;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_non_posted_request 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__non_posted_request;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_memory_read 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__memory_read;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_posted_write 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__posted_write;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_enabled_byte_count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__enabled_byte_count;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_request_supported 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_supported;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_request_present 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__request_present;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_requester_id 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__norm_requester_id;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_tag 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__norm_tag;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_traffic_class 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__norm_traffic_class;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_attributes 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__norm_attributes;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_header_4dw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__norm_header_4dw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_length_dw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__norm_length_dw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_first_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__norm_first_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_last_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__norm_last_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_bar_mask 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__norm_bar_mask;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_address 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__norm_address;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_first_completion_byte_count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__norm_first_completion_byte_count;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_first_completion_dw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__norm_first_completion_dw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_enabled_byte_count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__norm_enabled_byte_count;
    vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_first 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__norm_request_present;
    vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_rd = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_first) 
                                                   & ((IData)(vlSymsp->TOP__tb_top__DOT__tlps_in_if.tlast) 
                                                      & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_memory_read) 
                                                         & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_request_supported))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr_last) 
                                                   | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_first) 
                                                      & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr_ready) 
                                                         & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_posted_write) 
                                                            & (0U 
                                                               != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_enabled_byte_count))))));
}

void Vtop___024root___ico_sequent__TOP__2(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___ico_sequent__TOP__2\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
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

void Vtop___024root___ico_comb__TOP__0(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___ico_comb__TOP__0\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_en 
        = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng.tready;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_en) 
           & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__empty)));
}

void Vtop___024root___ico_comb__TOP__1(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___ico_comb__TOP__1\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
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
}

void Vtop___024root___ico_comb__TOP__2(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___ico_comb__TOP__2\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_pop 
        = ((IData)(vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_ur.tvalid) 
           & (IData)(vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_ur.tready));
    vlSelfRef.tb_top__DOT__tlps_out_tdata[0U] = vlSymsp->TOP__tb_top__DOT__tlps_out_if.tdata[0U];
    vlSelfRef.tb_top__DOT__tlps_out_tdata[1U] = vlSymsp->TOP__tb_top__DOT__tlps_out_if.tdata[1U];
    vlSelfRef.tb_top__DOT__tlps_out_tdata[2U] = vlSymsp->TOP__tb_top__DOT__tlps_out_if.tdata[2U];
    vlSelfRef.tb_top__DOT__tlps_out_tdata[3U] = vlSymsp->TOP__tb_top__DOT__tlps_out_if.tdata[3U];
    vlSelfRef.tb_top__DOT__tlps_out_tkeepdw = vlSymsp->TOP__tb_top__DOT__tlps_out_if.tkeepdw;
    vlSelfRef.tb_top__DOT__tlps_out_tuser = vlSymsp->TOP__tb_top__DOT__tlps_out_if.tuser;
    vlSelfRef.tb_top__DOT__tlps_out_has_data = vlSymsp->TOP__tb_top__DOT__tlps_out_if.has_data;
    vlSelfRef.tb_top__DOT__tlps_out_tvalid = vlSymsp->TOP__tb_top__DOT__tlps_out_if.tvalid;
    vlSelfRef.tb_top__DOT__tlps_out_tlast = vlSymsp->TOP__tb_top__DOT__tlps_out_if.tlast;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_can_accept 
        = (1U & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_fifo_full)) 
                 | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_pop)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_push 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_valid) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_can_accept));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_ready 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_can_accept;
    vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_ready 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_ready;
    vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_dequeue 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_ready) 
           & (0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_count)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_enqueue 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_pending) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_dequeue) 
              | (0x0100U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_count))));
}

void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__tlps_in_if__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_ur__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_cpl__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__tlps_in_if__1(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__1(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_cpl__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_ur__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__tlps_out_if__0(Vtop_IfAXIS128* vlSelf);

void Vtop___024root___eval_ico(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_ico\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    if ((1ULL & vlSelfRef.__VicoTriggered[1U])) {
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
    if ((1ULL & (vlSelfRef.__VicoTriggered[1U] | vlSelfRef.__VicoTriggered[0U]))) {
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

#ifdef VL_DEBUG
VL_ATTR_COLD void Vtop___024root___dump_triggers__ico(const VlUnpacked<QData/*63:0*/, 2> &triggers, const std::string &tag);
#endif  // VL_DEBUG

bool Vtop___024root___eval_phase__ico(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_phase__ico\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    CData/*0:0*/ __VicoExecute;
    // Body
    Vtop___024root___eval_triggers_vec__ico(vlSelf);
#ifdef VL_DEBUG
    if (VL_UNLIKELY(vlSymsp->_vm_contextp__->debug())) {
        Vtop___024root___dump_triggers__ico(vlSelfRef.__VicoTriggered, "ico"s);
    }
#endif
    __VicoExecute = Vtop___024root___trigger_anySet__ico(vlSelfRef.__VicoTriggered);
    if (__VicoExecute) {
        Vtop___024root___eval_ico(vlSelf);
    }
    return (__VicoExecute);
}

void Vtop___024root___eval_triggers_vec__act(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_triggers_vec__act\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.__VactTriggered[0U] = (QData)((IData)(
                                                    (((vlSelfRef.__VdlySched.awaitingCurrentTime() 
                                                       << 0x0000000aU) 
                                                      | ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__clk) 
                                                           & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__clk__0))) 
                                                          << 9U) 
                                                         | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__clk) 
                                                             & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_bar0__DOT__clk__0))) 
                                                            << 8U))) 
                                                     | (((((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__clk) 
                                                             & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__clk__0))) 
                                                            << 3U) 
                                                           | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__clk) 
                                                               & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__clk__0))) 
                                                              << 2U)) 
                                                          | ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__clk) 
                                                               & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__clk__0))) 
                                                              << 1U) 
                                                             | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__clk) 
                                                                & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__clk__0))))) 
                                                         << 4U) 
                                                        | (((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__clk) 
                                                              & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__clk__0))) 
                                                             << 3U) 
                                                            | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__clk) 
                                                                & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__clk__0))) 
                                                               << 2U)) 
                                                           | ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__clk) 
                                                                & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__clk__0))) 
                                                               << 1U) 
                                                              | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_select_ur) 
                                                                 != (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__cpl_select_ur__1))))))));
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
}

bool Vtop___024root___trigger_anySet__act(const VlUnpacked<QData/*63:0*/, 1> &in) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___trigger_anySet__act\n"); );
    // Locals
    IData/*31:0*/ n;
    // Body
    n = 0U;
    do {
        if (in[n]) {
            return (1U);
        }
        n = ((IData)(1U) + n);
    } while ((1U > n));
    return (0U);
}

void Vtop___024root___act_sequent__TOP__0(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___act_sequent__TOP__0\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__clk = vlSelfRef.tb_top__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__clk;
}

void Vtop___024root___eval_act(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_act\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    if ((0x0000000000000400ULL & vlSelfRef.__VactTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__clk = vlSelfRef.tb_top__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__clk;
    }
    if ((1ULL & vlSelfRef.__VactTriggered[0U])) {
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

void Vtop___024root___nba_sequent__TOP__0(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__0\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    SData/*10:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__total_length;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__total_length = 0;
    SData/*10:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__remaining_before;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__remaining_before = 0;
    SData/*10:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__completion_length;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__completion_length = 0;
    CData/*3:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__first_enable;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__first_enable = 0;
    CData/*3:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__last_enable;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__last_enable = 0;
    CData/*2:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9____VlefCall_2__popcount4;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9____VlefCall_2__popcount4 = 0;
    CData/*2:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9____VlefCall_1__popcount4;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9____VlefCall_1__popcount4 = 0;
    CData/*2:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9____VlefCall_0__popcount4;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9____VlefCall_0__popcount4 = 0;
    SData/*12:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__value;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__value = 0;
    CData/*2:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__10__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__10__Vfuncout = 0;
    CData/*3:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__10__value;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__10__value = 0;
    CData/*2:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__11__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__11__Vfuncout = 0;
    CData/*3:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__11__value;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__11__value = 0;
    CData/*2:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__12__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__12__Vfuncout = 0;
    CData/*3:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__12__value;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__12__value = 0;
    SData/*10:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__Vfuncout = 0;
    QData/*63:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__address_in;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__address_in = 0;
    SData/*10:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__remaining_in;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__remaining_in = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__rcb_dw;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__rcb_dw = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__mps_dw;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__mps_dw = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__selected_dw;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__selected_dw = 0;
    SData/*10:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__Vfuncout = 0;
    QData/*63:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__address_in;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__address_in = 0;
    SData/*10:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__remaining_in;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__remaining_in = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__rcb_dw;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__rcb_dw = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__mps_dw;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__mps_dw = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__selected_dw;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__selected_dw = 0;
    CData/*7:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr = 0;
    CData/*7:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__state;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__state = 0;
    QData/*63:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr = 0;
    SData/*10:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_length_dw;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_length_dw = 0;
    SData/*10:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw = 0;
    SData/*10:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw = 0;
    SData/*10:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_remaining_dw;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_remaining_dw = 0;
    SData/*12:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_byte_count;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_byte_count = 0;
    CData/*3:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_be;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_be = 0;
    CData/*3:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_last_be;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_last_be = 0;
    CData/*3:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw = 0;
    QData/*63:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_addr_fifo__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_addr_fifo__v0 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_addr_fifo__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_addr_fifo__v0 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_addr_fifo__v0;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_addr_fifo__v0 = 0;
    SData/*10:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_len_fifo__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_len_fifo__v0 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_len_fifo__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_len_fifo__v0 = 0;
    CData/*3:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_fbe_fifo__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_fbe_fifo__v0 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_fbe_fifo__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_fbe_fifo__v0 = 0;
    CData/*3:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_lbe_fifo__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_lbe_fifo__v0 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_lbe_fifo__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_lbe_fifo__v0 = 0;
    SData/*15:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_id_fifo__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_id_fifo__v0 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_id_fifo__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_id_fifo__v0 = 0;
    CData/*7:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tag_fifo__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tag_fifo__v0 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tag_fifo__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tag_fifo__v0 = 0;
    CData/*2:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tc_fifo__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tc_fifo__v0 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tc_fifo__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tc_fifo__v0 = 0;
    CData/*2:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_attr_fifo__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_attr_fifo__v0 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_attr_fifo__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_attr_fifo__v0 = 0;
    CData/*0:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_4dw_fifo__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_4dw_fifo__v0 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_4dw_fifo__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_4dw_fifo__v0 = 0;
    CData/*6:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bar_fifo__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bar_fifo__v0 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bar_fifo__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bar_fifo__v0 = 0;
    SData/*12:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bc_fifo__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bc_fifo__v0 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bc_fifo__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bc_fifo__v0 = 0;
    SData/*10:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_cpl_fifo__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_cpl_fifo__v0 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_cpl_fifo__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_cpl_fifo__v0 = 0;
    // Body
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_addr_fifo__v0 = 0U;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_length_dw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_length_dw;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_byte_count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_byte_count;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_remaining_dw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_remaining_dw;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__state 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__state;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_be;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_last_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_last_be;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rst) {
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_count = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[0U] = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[1U] = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[2U] = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[3U] = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__state = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr = 0ULL;
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_length_dw = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_remaining_dw = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_byte_count = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_be = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_last_be = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_id = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_tag = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_tc = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_attr = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_4dw = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_bar = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_packet = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first = 0U;
    } else {
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_pop) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr 
                = (0x000000ffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr)));
        }
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_count 
            = (0x000001ffU & ((2U == (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_push) 
                                       << 1U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_pop)))
                               ? ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_count))
                               : ((1U == (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_push) 
                                           << 1U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_pop)))
                                   ? ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_count) 
                                      - (IData)(1U))
                                   : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_count))));
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_push) {
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_addr_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_address;
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_addr_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr;
            __VdlySet__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_addr_fifo__v0 = 1U;
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_len_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_length_dw;
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_len_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr;
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_fbe_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_first_be;
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_fbe_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr;
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_lbe_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_last_be;
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_lbe_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr;
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_id_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_requester_id;
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_id_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr;
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tag_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_tag;
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tag_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr;
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tc_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_traffic_class;
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tc_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr;
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_attr_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_attributes;
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_attr_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr;
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_4dw_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_header_4dw;
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_4dw_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr;
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bar_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_bar_mask;
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bar_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr;
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bc_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_first_completion_byte_count;
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bc_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr;
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_cpl_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__norm_first_completion_dw;
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_cpl_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr;
            __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr 
                = (0x000000ffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr)));
        }
        if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_valid) 
             & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_first))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[0U] 
                = ((0x0003ffffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[0U]) 
                   | (0xfffc0000U & (0x4a000000U | 
                                     (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_tc) 
                                       << 0x00000014U) 
                                      | (0x00040000U 
                                         & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_attr) 
                                            << 0x00000010U))))));
            __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw = 0x0fU;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[0U] 
                = ((0xfffc03ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[0U]) 
                   | (0x00003000U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_attr) 
                                     << 0x0000000cU)));
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[0U] 
                = ((0xfffffc00U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[0U]) 
                   | (0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_dwlen)));
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[1U] 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_byte_count;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[2U] 
                = (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_requester) 
                    << 0x00000010U) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_tag) 
                                        << 8U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_lower_addr)));
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[3U] 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_data_swapped;
        } else {
            __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw 
                = (0x0000000fU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tvalid)
                                   ? ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_valid)
                                       ? 1U : 0U) : 
                                  ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_valid)
                                    ? (1U | VL_SHIFTL_III(4,4,32, (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw), 1U))
                                    : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw))));
            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_valid) {
                if ((1U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tvalid) 
                           | (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw))))) {
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[0U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_data_swapped;
                }
                if ((1U & (~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw) 
                              >> 1U)))) {
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[1U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_data_swapped;
                }
                if ((1U & (~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw) 
                              >> 2U)))) {
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[2U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_data_swapped;
                }
                if ((1U & (~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw) 
                              >> 3U)))) {
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[3U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_data_swapped;
                }
            }
        }
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__state) {
            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__issue_dw) {
                if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_last) {
                    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_last_packet) {
                        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__state = 0U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first = 0U;
                    } else {
                        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr 
                            = (4ULL + vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr);
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_packet = 0U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first = 1U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw 
                            = (0x000007ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw) 
                                              - (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw)));
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__last_enable 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_last_be;
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__first_enable 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_be;
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__completion_length 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw;
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__remaining_before 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw;
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__total_length 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_length_dw;
                        if ((1U == (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__total_length))) {
                            __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__10__value 
                                = __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__first_enable;
                            __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__10__Vfuncout 
                                = (7U & ((((1U & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__10__value)) 
                                           + (1U & 
                                              ((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__10__value) 
                                               >> 1U))) 
                                          + (1U & ((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__10__value) 
                                                   >> 2U))) 
                                         + (1U & ((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__10__value) 
                                                  >> 3U))));
                            __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9____VlefCall_0__popcount4 
                                = __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__10__Vfuncout;
                            __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__value 
                                = __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9____VlefCall_0__popcount4;
                        } else {
                            __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__value 
                                = ((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__completion_length) 
                                   << 2U);
                            if (((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__remaining_before) 
                                 == (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__total_length))) {
                                __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__11__value 
                                    = __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__first_enable;
                                __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__11__Vfuncout 
                                    = (7U & ((((1U 
                                                & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__11__value)) 
                                               + (1U 
                                                  & ((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__11__value) 
                                                     >> 1U))) 
                                              + (1U 
                                                 & ((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__11__value) 
                                                    >> 2U))) 
                                             + (1U 
                                                & ((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__11__value) 
                                                   >> 3U))));
                                __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9____VlefCall_1__popcount4 
                                    = __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__11__Vfuncout;
                                __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__value 
                                    = (0x00001fffU 
                                       & ((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__value) 
                                          - ((IData)(4U) 
                                             - (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9____VlefCall_1__popcount4))));
                            }
                            if (((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__completion_length) 
                                 == (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__remaining_before))) {
                                __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__12__value 
                                    = __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__last_enable;
                                __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__12__Vfuncout 
                                    = (7U & ((((1U 
                                                & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__12__value)) 
                                               + (1U 
                                                  & ((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__12__value) 
                                                     >> 1U))) 
                                              + (1U 
                                                 & ((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__12__value) 
                                                    >> 2U))) 
                                             + (1U 
                                                & ((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__12__value) 
                                                   >> 3U))));
                                __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9____VlefCall_2__popcount4 
                                    = __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__popcount4__12__Vfuncout;
                                __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__value 
                                    = (0x00001fffU 
                                       & ((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__value) 
                                          - ((IData)(4U) 
                                             - (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9____VlefCall_2__popcount4))));
                            }
                        }
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT____VlemCall_0__completed_enabled_bytes 
                            = __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__completed_enabled_bytes__9__value;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_byte_count 
                            = (0x00001fffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_byte_count) 
                                              - (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT____VlemCall_0__completed_enabled_bytes)));
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__remaining_in 
                            = (0x000007ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw) 
                                              - (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw)));
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__address_in 
                            = (4ULL + vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr);
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__rcb_dw 
                            = (IData)(((0x0000000000000040ULL 
                                        - (0x000000000000003fULL 
                                           & __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__address_in)) 
                                       >> 2U));
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__mps_dw = 0x00000020U;
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__selected_dw 
                            = __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__remaining_in;
                        if (VL_GTS_III(32, __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__selected_dw, __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__rcb_dw)) {
                            __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__selected_dw 
                                = __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__rcb_dw;
                        }
                        if (VL_LTS_III(32, 0x00000020U, __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__selected_dw)) {
                            __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__selected_dw 
                                = __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__mps_dw;
                        }
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__Vfuncout 
                            = (0x000007ffU & __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__selected_dw);
                        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw 
                            = __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__13__Vfuncout;
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__remaining_in 
                            = (0x000007ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw) 
                                              - (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw)));
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__address_in 
                            = (4ULL + vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr);
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__rcb_dw 
                            = (IData)(((0x0000000000000040ULL 
                                        - (0x000000000000003fULL 
                                           & __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__address_in)) 
                                       >> 2U));
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__mps_dw = 0x00000020U;
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__selected_dw 
                            = __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__remaining_in;
                        if (VL_GTS_III(32, __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__selected_dw, __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__rcb_dw)) {
                            __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__selected_dw 
                                = __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__rcb_dw;
                        }
                        if (VL_LTS_III(32, 0x00000020U, __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__selected_dw)) {
                            __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__selected_dw 
                                = __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__mps_dw;
                        }
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__Vfuncout 
                            = (0x000007ffU & __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__selected_dw);
                        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_remaining_dw 
                            = __Vfunc_tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__next_completion_dw__14__Vfuncout;
                    }
                } else {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr 
                        = (4ULL + vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr);
                    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_remaining_dw 
                        = (0x000007ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_remaining_dw) 
                                          - (IData)(1U)));
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first = 0U;
                }
            }
        } else if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_pop) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_addr_fifo
                [vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr];
            __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_length_dw 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_len_fifo
                [vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr];
            __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_len_fifo
                [vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr];
            __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_cpl_fifo
                [vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr];
            __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_remaining_dw 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_cpl_fifo
                [vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr];
            __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_byte_count 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bc_fifo
                [vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr];
            __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_be 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_fbe_fifo
                [vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr];
            __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_last_be 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_lbe_fifo
                [vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr];
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_id 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_id_fifo
                [vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr];
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_tag 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tag_fifo
                [vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr];
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_tc 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tc_fifo
                [vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr];
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_attr 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_attr_fifo
                [vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr];
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_4dw 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_4dw_fifo
                [vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr];
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_bar 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bar_fifo
                [vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr];
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_packet = 1U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first = 1U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__state = 1U;
        }
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__first 
        = ((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rst))) 
           && ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_valid) 
               & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_first)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tlast 
        = ((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rst))) 
           && (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_valid) 
                & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_first))
                ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_last)
                : ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_valid) 
                   & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_last))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_wr_ptr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_length_dw 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_length_dw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_rd_ptr;
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_addr_fifo__v0) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_addr_fifo[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_addr_fifo__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_addr_fifo__v0;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_len_fifo[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_len_fifo__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_len_fifo__v0;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_cpl_fifo[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_cpl_fifo__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_cpl_fifo__v0;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bc_fifo[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bc_fifo__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bc_fifo__v0;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_fbe_fifo[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_fbe_fifo__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_fbe_fifo__v0;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_lbe_fifo[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_lbe_fifo__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_lbe_fifo__v0;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_id_fifo[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_id_fifo__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_id_fifo__v0;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tag_fifo[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tag_fifo__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tag_fifo__v0;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tc_fifo[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tc_fifo__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_tc_fifo__v0;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_attr_fifo[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_attr_fifo__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_attr_fifo__v0;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_4dw_fifo[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_4dw_fifo__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_4dw_fifo__v0;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bar_fifo[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bar_fifo__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_bar_fifo__v0;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_byte_count 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_byte_count;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_remaining_dw 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_remaining_dw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__state 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__state;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_be 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_last_be 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_last_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_fifo_full 
        = (0x0100U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tvalid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tlast) 
           | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw) 
              >> 3U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_bar 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_bar;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_last 
        = (1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_remaining_dw));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_pop 
        = (1U & (~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__state) 
                    | (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_count)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_last_packet 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw) 
           == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be 
        = (0x0000000fU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_packet)
                           ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_be)
                           : ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_last_be) 
                              | (- (IData)((1U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw)))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_addr 
        = (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr);
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_bar;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_can_accept 
        = (1U & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_fifo_full)) 
                 | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_pop)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first_offset 
        = (3U & (((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be))
                   ? 1U : ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be))
                            ? 2U : (- (IData)((1U & 
                                               ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be) 
                                                >> 3U)))))) 
                 & (- (IData)((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be)))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_push 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tlps_in_valid) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_can_accept));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__context_address 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first)
            ? (((IData)((vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr 
                         >> 2U)) << 2U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first_offset))
            : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_rd_hit 
        = (0x00004000U > vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr);
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr_bar0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__context_address;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[1U] 
        = ((0xf0000000U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[1U]) 
           | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_attr) 
               << 0x00000019U) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_4dw) 
                                   << 0x00000018U) 
                                  | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_tag) 
                                      << 0x00000010U) 
                                     | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_id)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[1U] 
        = ((0x0fffffffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[1U]) 
           | (((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first) 
                 << 0x0000001bU) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_last) 
                                     << 0x0000001aU) 
                                    | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw) 
                                       << 0x0000000fU))) 
               | ((0x00007ff8U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_byte_count) 
                                  << 3U)) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_tc))) 
              << 0x0000001cU));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[2U] 
        = (0x00ffffffU & (((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first) 
                             << 0x0000001bU) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_last) 
                                                 << 0x0000001aU) 
                                                | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw) 
                                                   << 0x0000000fU))) 
                           | ((0x00007ff8U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_byte_count) 
                                              << 3U)) 
                              | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_tc))) 
                          >> 4U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_req_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_req_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_req_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_req_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_req_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[2U];
}

void Vtop___024root___nba_sequent__TOP__1(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__1\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_counter;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_counter = 0;
    CData/*3:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending = 0;
    CData/*7:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__latency_counter;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__latency_counter = 0;
    CData/*7:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__target_latency;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__target_latency = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_valid;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_valid = 0;
    VlWide<3>/*87:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx;
    VL_ZERO_W(88, __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx);
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_data;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_data = 0;
    CData/*3:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__idle_counter;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__idle_counter = 0;
    IData/*18:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_pending;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_pending = 0;
    CData/*7:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_counter;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_counter = 0;
    CData/*7:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_target;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_target = 0;
    // Body
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_pending 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_pending;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_counter 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_counter;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_target 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_target;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_counter 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_counter;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__latency_counter 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__latency_counter;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__target_latency 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__target_latency;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[0U];
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[1U];
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[2U];
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_data;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__idle_counter 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__idle_counter;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_valid;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rst) {
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_counter = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_pending = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_counter = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_target = 2U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_ack = 1U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s0 = 0xc049b907U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__latency_counter = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__target_latency = 7U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_valid = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[0U] = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[1U] = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[2U] = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_data = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_addr = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_latency = 7U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_valid = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__idle_counter = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s1 = 0xc85a4b71U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s2 = 0x644c6af7U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s3 = 0x6541532eU;
    } else {
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_counter 
            = (0x0000ffffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_counter)));
        if ((0x8000U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_counter))) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset 
                = (0x0000000fU & ((1U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out)
                                   ? ((0x0fU > (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset))
                                       ? ((IData)(1U) 
                                          + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset))
                                       : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset))
                                   : ((0U < (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset))
                                       ? ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset) 
                                          - (IData)(1U))
                                       : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset))));
            __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_counter = 0U;
        }
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_pending) {
            if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_counter) 
                 < (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_target))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_counter 
                    = (0x000000ffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_counter)));
            } else {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_ack = 1U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_pending = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_counter = 0U;
            }
        } else if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_valid) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_pending = 1U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_ack = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_counter = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_target 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_latency;
        }
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s0 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s1;
        if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_valid) 
             & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ready))) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_valid = 0U;
        }
        if ((1U & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending)) 
                   & (~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_valid) 
                         & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_ready)))))) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__idle_counter 
                = (0x0000000fU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__idle_counter)));
        } else if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_valid) 
                    & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_ready))) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__idle_counter = 0U;
        }
        if ((0x0fU <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__idle_counter))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_valid = 0U;
        }
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter 
                = (0x0007ffffU & ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter));
            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_expired) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_valid = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ctx[0U] 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[0U];
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ctx[1U] 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[1U];
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ctx[2U] 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[2U];
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_data = 0xffffffffU;
                __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__latency_counter = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter = 0U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_valid = 0U;
            } else if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__latency_counter) 
                        < (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__target_latency))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__latency_counter 
                    = (0x000000ffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__latency_counter)));
            } else if ((1U & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_valid)) 
                              | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ready)))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_valid = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ctx[0U] 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[0U];
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ctx[1U] 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[1U];
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ctx[2U] 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[2U];
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_data 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_data;
                __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__latency_counter = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter = 0U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_latency 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__target_latency;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_valid = 1U;
            }
        } else if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_valid) 
                    & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_ready))) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending = 1U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[0U] 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_ctx[0U];
            __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[1U] 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_ctx[1U];
            __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[2U] 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_ctx[2U];
            __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_data 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_data;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_addr 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_addr;
            __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__latency_counter = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__target_latency 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__computed_latency;
        }
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s1 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s2;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s2 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s3;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s3 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_next_s3;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_counter 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_counter;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_pending 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_pending;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_counter 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_counter;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_target 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_target;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__latency_counter 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__latency_counter;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__target_latency 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__target_latency;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[0U] 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[1U] 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[2U] 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_data 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__idle_counter 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__idle_counter;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_valid 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_wr_ack 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_ack;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_next_s0 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s0 
           ^ VL_SHIFTL_III(32,32,32, vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s0, 0x0000000bU));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_ready 
        = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_expired 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending) 
           & (0x00010000U <= vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__busy 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr_ready 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_wr_ack;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_valid 
        = ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_rsp_valid) 
             << 5U) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_rsp_valid) 
                        << 4U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_rsp_valid) 
                                  << 3U))) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_rsp_valid) 
                                               << 2U) 
                                              | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_rsp_valid) 
                                                  << 1U) 
                                                 | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_valid))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[0U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[0U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[0U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_valid 
        = (0x0000003fU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_valid));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[0U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[0U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[0U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[0U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[0U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[0U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[1U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[1U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[1U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[1U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[1U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[1U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[2U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[2U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[2U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[2U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[2U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[2U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[3U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[3U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[3U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[3U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[3U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[3U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[4U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[4U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[4U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[4U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[4U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[4U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[5U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[5U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[5U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[5U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[5U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[5U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[3U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[3U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[4U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[4U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[5U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[5U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[0U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[0U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[0U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[0U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[0U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[0U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[1U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[1U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[1U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[1U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[1U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[1U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[2U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[2U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[2U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[2U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[2U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[2U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[3U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[3U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[3U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[3U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[3U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[3U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[4U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[4U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[4U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[4U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[4U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[4U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[5U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[5U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[5U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[5U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[5U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[5U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[6U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[6U][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[6U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[6U][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[6U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[6U][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data[3U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[3U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data[4U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[4U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data[5U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[5U];
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data[6U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[6U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_next_s3 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s3 
           ^ (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_next_s0 
              ^ (VL_SHIFTR_III(32,32,32, vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s3, 0x00000013U) 
                 ^ VL_SHIFTR_III(32,32,32, vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_next_s0, 8U))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s0 
           + vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s3);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_latency 
        = (0x000000ffU & ((IData)(2U) + VL_MODDIV_III(32, 
                                                      (0x000000ffU 
                                                       & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out 
                                                          >> 8U)), (IData)(7U))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 0x0fU;
    if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
         < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 0U;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[1U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 1U;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[2U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 2U;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[3U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 3U;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[4U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 4U;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[5U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 5U;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[6U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 6U;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[7U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 7U;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[8U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 8U;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[9U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 9U;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[10U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 0x0aU;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[11U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 0x0bU;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[12U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 0x0cU;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[13U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 0x0dU;
    } else if (((0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out) 
                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_table[14U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket = 0x0eU;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_latency 
        = (0x000000ffU & ((IData)(3U) + (((IData)(0x0000000dU) 
                                          * (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket)) 
                                         >> 4U)));
}

void Vtop___024root___nba_sequent__TOP__2(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__2\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    CData/*7:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__write_ptr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__write_ptr = 0;
    CData/*7:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__read_ptr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__read_ptr = 0;
    SData/*15:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_fifo__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_fifo__v0 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_fifo__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_fifo__v0 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_fifo__v0;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_fifo__v0 = 0;
    CData/*7:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tag_fifo__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tag_fifo__v0 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tag_fifo__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tag_fifo__v0 = 0;
    CData/*2:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tc_fifo__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tc_fifo__v0 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tc_fifo__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tc_fifo__v0 = 0;
    CData/*2:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__attr_fifo__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__attr_fifo__v0 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__attr_fifo__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__attr_fifo__v0 = 0;
    // Body
    __Vdly__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__read_ptr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__read_ptr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__write_ptr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__write_ptr;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_fifo__v0 = 0U;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__rst) {
        __Vdly__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__read_ptr = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__write_ptr = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_count = 0U;
    } else {
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_pop) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__read_ptr 
                = (0x000000ffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__read_ptr)));
        }
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_push) {
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_id;
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__write_ptr;
            __VdlySet__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_fifo__v0 = 1U;
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tag_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tag;
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tag_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__write_ptr;
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tc_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__traffic_class;
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tc_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__write_ptr;
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__attr_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__attributes;
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__attr_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__write_ptr;
            __Vdly__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__write_ptr 
                = (0x000000ffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__write_ptr)));
        }
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_count 
            = (0x000001ffU & ((2U == (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_push) 
                                       << 1U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_pop)))
                               ? ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_count))
                               : ((1U == (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_push) 
                                           << 1U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_pop)))
                                   ? ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_count) 
                                      - (IData)(1U))
                                   : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_count))));
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__read_ptr 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__read_ptr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__write_ptr 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__write_ptr;
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_fifo__v0) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_fifo[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_fifo__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_fifo__v0;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tag_fifo[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tag_fifo__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tag_fifo__v0;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tc_fifo[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tc_fifo__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tc_fifo__v0;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__attr_fifo[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__attr_fifo__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__attr_fifo__v0;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[0U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[1U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[2U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[3U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[0U] 
        = (0x0a000000U | (0x00ffffffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[0U]));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[0U] 
        = ((0xff8fffffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[0U]) 
           | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tc_fifo
                      [vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__read_ptr]) 
              << 0x00000014U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[0U] 
        = ((0xfffbffffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[0U]) 
           | (0x00040000U & ((vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__attr_fifo
                              [vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__read_ptr] 
                              >> 2U) << 0x00000012U)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[0U] 
        = ((0xffffcfffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[0U]) 
           | (0x00003000U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__attr_fifo
                             [vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__read_ptr] 
                             << 0x0000000cU)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[1U] = 0x00002000U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__completion_data[2U] 
        = (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_fifo
                    [vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__read_ptr]) 
            << 0x00000010U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tag_fifo
                                       [vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__read_ptr]) 
                               << 8U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_fifo_full 
        = (0x0100U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_count));
}

extern const VlUnpacked<CData/*2:0*/, 128> Vtop__ConstPool__TABLE_h181a664b_0;
extern const VlUnpacked<CData/*0:0*/, 128> Vtop__ConstPool__TABLE_ha0b177da_0;
extern const VlUnpacked<CData/*0:0*/, 128> Vtop__ConstPool__TABLE_h8bdc9cb2_0;
extern const VlUnpacked<CData/*0:0*/, 128> Vtop__ConstPool__TABLE_ha1fdee70_0;

void Vtop___024root___nba_sequent__TOP__3(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__3\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    CData/*6:0*/ __Vtableidx1;
    __Vtableidx1 = 0;
    CData/*7:0*/ __Vdly__tb_top__DOT__i_bar__DOT__ur_request_write_ptr;
    __Vdly__tb_top__DOT__i_bar__DOT__ur_request_write_ptr = 0;
    CData/*7:0*/ __Vdly__tb_top__DOT__i_bar__DOT__ur_request_read_ptr;
    __Vdly__tb_top__DOT__i_bar__DOT__ur_request_read_ptr = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__bar0_spill_valid;
    __Vdly__tb_top__DOT__i_bar__DOT__bar0_spill_valid = 0;
    CData/*3:0*/ __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_rd;
    __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_rd = 0;
    CData/*4:0*/ __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_count;
    __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_count = 0;
    CData/*3:0*/ __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_wr;
    __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_wr = 0;
    IData/*29:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__ur_request_fifo__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__ur_request_fifo__v0 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__ur_request_fifo__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__ur_request_fifo__v0 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__ur_request_fifo__v0;
    __VdlySet__tb_top__DOT__i_bar__DOT__ur_request_fifo__v0 = 0;
    VlWide<3>/*87:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v0;
    VL_ZERO_W(88, __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v0);
    CData/*3:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v0 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v0;
    __VdlySet__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v0 = 0;
    IData/*31:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_data__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_data__v0 = 0;
    CData/*3:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_data__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_data__v0 = 0;
    IData/*31:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_addr__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_addr__v0 = 0;
    CData/*3:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_addr__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_addr__v0 = 0;
    VlWide<3>/*87:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v1;
    VL_ZERO_W(88, __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v1);
    CData/*3:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v1;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v1 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v1;
    __VdlySet__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v1 = 0;
    IData/*31:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_data__v1;
    __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_data__v1 = 0;
    CData/*3:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_data__v1;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_data__v1 = 0;
    IData/*31:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_addr__v1;
    __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_addr__v1 = 0;
    CData/*3:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_addr__v1;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_addr__v1 = 0;
    // Body
    __Vdly__tb_top__DOT__i_bar__DOT__ur_request_read_ptr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_read_ptr;
    __Vdly__tb_top__DOT__i_bar__DOT__ur_request_write_ptr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_write_ptr;
    __VdlySet__tb_top__DOT__i_bar__DOT__ur_request_fifo__v0 = 0U;
    __Vdly__tb_top__DOT__i_bar__DOT__bar0_spill_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_valid;
    __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_rd = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_rd;
    __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_count;
    __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_wr = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_wr;
    __VdlySet__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v0 = 0U;
    __VdlySet__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v1 = 0U;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset) {
        __Vdly__tb_top__DOT__i_bar__DOT__ur_request_read_ptr = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_count = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__ur_request_write_ptr = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_rd = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_count = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_wr = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr_last = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__bar0_spill_valid = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_ctx[0U] = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_ctx[1U] = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_ctx[2U] = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_data = 0U;
    } else {
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_dequeue) {
            __Vdly__tb_top__DOT__i_bar__DOT__ur_request_read_ptr 
                = (0x000000ffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_read_ptr)));
        }
        vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_count 
            = (0x000001ffU & ((2U == (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_enqueue) 
                                       << 1U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_dequeue)))
                               ? ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_count))
                               : ((1U == (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_enqueue) 
                                           << 1U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_dequeue)))
                                   ? ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_count) 
                                      - (IData)(1U))
                                   : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_count))));
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_enqueue) {
            __VdlyVal__tb_top__DOT__i_bar__DOT__ur_request_fifo__v0 
                = (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_requester_id) 
                    << 0x0000000eU) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_tag) 
                                        << 6U) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_traffic_class) 
                                                   << 3U) 
                                                  | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_attributes))));
            __VdlyDim0__tb_top__DOT__i_bar__DOT__ur_request_fifo__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_write_ptr;
            __VdlySet__tb_top__DOT__i_bar__DOT__ur_request_fifo__v0 = 1U;
            __Vdly__tb_top__DOT__i_bar__DOT__ur_request_write_ptr 
                = (0x000000ffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_write_ptr)));
        }
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_pop) {
            __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_rd 
                = (0x0000000fU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_rd)));
        }
        if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_write) 
             & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_pop)))) {
            __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_count 
                = (0x0000001fU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_count)));
        } else if (((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_write)) 
                    & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_pop))) {
            __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_count 
                = (0x0000001fU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_count) 
                                  - (IData)(1U)));
        }
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_write) {
            if (vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_valid) {
                __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v0[0U] 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_ctx[0U];
                __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v0[1U] 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_ctx[1U];
                __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v0[2U] 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_ctx[2U];
                __VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v0 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_wr;
                __VdlySet__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v0 = 1U;
                __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_data__v0 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_data;
                __VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_data__v0 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_wr;
                __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_addr__v0 
                    = (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_ctx[0U]);
                __VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_addr__v0 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_wr;
            } else {
                __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v1[0U] 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[0U];
                __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v1[1U] 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[1U];
                __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v1[2U] 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[2U];
                __VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v1 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_wr;
                __VdlySet__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v1 = 1U;
                __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_data__v1 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_data;
                __VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_data__v1 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_wr;
                __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_addr__v1 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_req_addr_from_ctx;
                __VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_addr__v1 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_wr;
            }
            __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_wr 
                = (0x0000000fU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_wr)));
        }
        if (vlSymsp->TOP__tb_top__DOT__tlps_in_if.tvalid) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr_last 
                = ((~ (IData)(vlSymsp->TOP__tb_top__DOT__tlps_in_if.tlast)) 
                   & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr));
        }
        if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_valid) 
             & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_full)))) {
            __Vdly__tb_top__DOT__i_bar__DOT__bar0_spill_valid = 0U;
        } else if ((((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_valid)) 
                     & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_valid)) 
                    & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_full))) {
            __Vdly__tb_top__DOT__i_bar__DOT__bar0_spill_valid = 1U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_ctx[0U] 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[0U];
            vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_ctx[1U] 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[1U];
            vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_ctx[2U] 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[2U];
            vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_data 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_data;
        }
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy_d1 
        = ((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset))) 
           && (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy));
    __Vtableidx1 = ((((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_select_ur) 
                        << 3U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_grant_ur) 
                                  << 2U)) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_selected_last) 
                                              << 1U) 
                                             | (IData)(vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_cpl.tready))) 
                     << 3U) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_selected_valid) 
                                << 2U) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_grant_active) 
                                           << 1U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset))));
    if ((1U & Vtop__ConstPool__TABLE_h181a664b_0[__Vtableidx1])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_grant_active 
            = Vtop__ConstPool__TABLE_ha0b177da_0[__Vtableidx1];
    }
    if ((2U & Vtop__ConstPool__TABLE_h181a664b_0[__Vtableidx1])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_grant_ur 
            = Vtop__ConstPool__TABLE_h8bdc9cb2_0[__Vtableidx1];
    }
    if ((4U & Vtop__ConstPool__TABLE_h181a664b_0[__Vtableidx1])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_rr_ur 
            = Vtop__ConstPool__TABLE_ha1fdee70_0[__Vtableidx1];
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_read_ptr 
        = __Vdly__tb_top__DOT__i_bar__DOT__ur_request_read_ptr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_write_ptr 
        = __Vdly__tb_top__DOT__i_bar__DOT__ur_request_write_ptr;
    if (__VdlySet__tb_top__DOT__i_bar__DOT__ur_request_fifo__v0) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_fifo[__VdlyDim0__tb_top__DOT__i_bar__DOT__ur_request_fifo__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__ur_request_fifo__v0;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_rd 
        = __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_rd;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_count 
        = __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_count;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_wr 
        = __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_wr;
    if (__VdlySet__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v0) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx[__VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v0][0U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v0[0U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx[__VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v0][1U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v0[1U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx[__VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v0][2U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v0[2U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_data[__VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_data__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_data__v0;
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_addr[__VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_addr__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_addr__v0;
    }
    if (__VdlySet__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v1) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx[__VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v1][0U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v1[0U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx[__VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v1][1U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v1[1U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx[__VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v1][2U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v1[2U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_data[__VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_data__v1] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_data__v1;
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_addr[__VdlyDim0__tb_top__DOT__i_bar__DOT__bar0_buf_addr__v1] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__bar0_buf_addr__v1;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_valid 
        = (0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__requester_id 
        = (0x0000ffffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_fifo
                          [vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_read_ptr] 
                          >> 0x0000000eU));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__tag 
        = (0x000000ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_fifo
                          [vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_read_ptr] 
                          >> 6U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__traffic_class 
        = (7U & (vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_fifo
                 [vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_read_ptr] 
                 >> 3U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__attributes 
        = (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_fifo
           [vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_read_ptr]);
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_empty 
        = (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx_w[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx
        [vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_rd][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx_w[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx
        [vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_rd][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx_w[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx
        [vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_rd][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_data_w 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_data
        [vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_rd];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_addr_w 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_addr
        [vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_rd];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_valid 
        = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_empty)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx_w[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx_w[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_ctx_w[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_data_w;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_addr_w;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_full 
        = (0x10U <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_valid 
        = __Vdly__tb_top__DOT__i_bar__DOT__bar0_spill_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__addr_region 
        = (7U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_addr 
                 >> 0x0000000cU));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_almost_full 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_valid) 
           | (0x0eU <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_count)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adjust 
        = (((1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__addr_region))
             ? 2U : ((2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__addr_region))
                      ? 4U : ((3U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__addr_region))
                               ? 1U : 3U))) & (- (IData)(
                                                         (0U 
                                                          != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__addr_region)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__stall 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_almost_full;
}

void Vtop___024root___nba_sequent__TOP__4(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__4\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    CData/*3:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__addr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__addr = 0;
    VlWide<4>/*127:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata;
    VL_ZERO_W(128, __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata);
    CData/*3:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tkeepdw;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tkeepdw = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlast;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlast = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__first_dw;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__first_dw = 0;
    // Body
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__addr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[0U];
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[1U];
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[2U];
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[3U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[3U];
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlast 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlast;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__first_dw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__first_dw;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tkeepdw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tkeepdw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_valid 
        = (((5U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state)) 
            | (6U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state))) 
           | (7U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__addr;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__rst) {
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state = 0U;
    } else if ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state))) {
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state = 1U;
    } else if ((1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state))) {
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state = 0U;
    } else if ((2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state))) {
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state = 4U;
    } else if ((4U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state))) {
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__addr 
            = ((IData)(4U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__addr);
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[0U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tdata[0U];
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[1U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tdata[1U];
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[2U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tdata[2U];
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[3U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tdata[3U];
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tkeepdw 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tkeepdw;
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlast 
            = (1U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tuser) 
                     >> 1U));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_data 
            = ((((0x0000ff00U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tdata[0U] 
                                 << 8U)) | (0x000000ffU 
                                            & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tdata[0U] 
                                               >> 8U))) 
                << 0x00000010U) | ((0x0000ff00U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tdata[0U] 
                                                   >> 8U)) 
                                   | (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tdata[0U] 
                                      >> 0x00000018U)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_be 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__first_dw)
                ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__be_first)
                : ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tkeepdw))
                    ? 0x0fU : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__be_last)));
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__first_dw = 0U;
    } else if ((5U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state))) {
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__addr 
            = ((IData)(4U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__addr);
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_data 
            = ((((0x0000ff00U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[1U] 
                                 << 8U)) | (0x000000ffU 
                                            & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[1U] 
                                               >> 8U))) 
                << 0x00000010U) | ((0x0000ff00U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[1U] 
                                                   >> 8U)) 
                                   | (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[1U] 
                                      >> 0x00000018U)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_be 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__first_dw)
                ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__be_first)
                : ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tkeepdw))
                    ? 0x0fU : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__be_last)));
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state 
            = ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tkeepdw))
                ? 6U : 1U);
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__first_dw = 0U;
    } else if ((6U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state))) {
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__addr 
            = ((IData)(4U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__addr);
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_data 
            = ((((0x0000ff00U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[2U] 
                                 << 8U)) | (0x000000ffU 
                                            & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[2U] 
                                               >> 8U))) 
                << 0x00000010U) | ((0x0000ff00U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[2U] 
                                                   >> 8U)) 
                                   | (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[2U] 
                                      >> 0x00000018U)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_be 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__first_dw)
                ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__be_first)
                : ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tkeepdw))
                    ? 0x0fU : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__be_last)));
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state 
            = ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tkeepdw))
                ? 7U : 1U);
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__first_dw = 0U;
    } else if ((7U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state))) {
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__addr 
            = ((IData)(4U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__addr);
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_data 
            = ((((0x0000ff00U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[3U] 
                                 << 8U)) | (0x000000ffU 
                                            & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[3U] 
                                               >> 8U))) 
                << 0x00000010U) | ((0x0000ff00U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[3U] 
                                                   >> 8U)) 
                                   | (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[3U] 
                                      >> 0x00000018U)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_be 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__first_dw)
                ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__be_first)
                : ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlast)
                    ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__be_last)
                    : 0x0fU));
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlast)
                ? 1U : 4U);
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__first_dw = 0U;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__addr 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[0U] 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[1U] 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[2U] 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[3U] 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tdata[3U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlast 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlast;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__first_dw 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__first_dw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tkeepdw 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tkeepdw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_rd_en 
        = ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state)) 
           | ((2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state)) 
              | ((7U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state)) 
                 | (((~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tkeepdw) 
                         >> 3U)) & (6U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state))) 
                    | (((~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tkeepdw) 
                            >> 1U)) & (4U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state))) 
                       | ((~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tkeepdw) 
                              >> 2U)) & (5U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__state))))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_be = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__wr_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_bar) 
              >> 1U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__wr_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_bar) 
              >> 2U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__wr_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_bar) 
              >> 3U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__wr_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_bar) 
              >> 4U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__wr_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_bar) 
              >> 5U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_wr_hit 
        = (0x00004000U > vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__rd_en 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_rd_en;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__wr_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__wr_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__wr_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__wr_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__wr_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__wr_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__wr_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__wr_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__wr_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__wr_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0;
    vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_wr_hit) 
              & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_bar)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid;
}

void Vtop___024root___nba_sequent__TOP__5(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__5\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    CData/*7:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v0 = 0;
    SData/*9:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v0 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v0;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v0 = 0;
    CData/*7:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v1;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v1 = 0;
    SData/*9:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v1;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v1 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v1;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v1 = 0;
    CData/*7:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v2;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v2 = 0;
    SData/*9:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v2;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v2 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v2;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v2 = 0;
    CData/*7:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v3;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v3 = 0;
    SData/*9:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v3;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v3 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v3;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v3 = 0;
    // Body
    __VdlySet__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v0 = 0U;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v1 = 0U;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v2 = 0U;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v3 = 0U;
    if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_valid) 
         & (0x00001000U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_addr))) {
        if ((0U == (3U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_addr))) {
            if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v0 
                    = (0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data);
                __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v0 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_word;
                __VdlySet__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v0 = 1U;
            }
            if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v1 
                    = (0x000000ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data 
                                      >> 8U));
                __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v1 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_word;
                __VdlySet__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v1 = 1U;
            }
            if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v2 
                    = (0x000000ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data 
                                      >> 0x10U));
                __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v2 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_word;
                __VdlySet__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v2 = 1U;
            }
            if ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v3 
                    = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data 
                       >> 0x18U);
                __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v3 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_word;
                __VdlySet__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v3 = 1U;
            }
        }
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx_d1[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx_d1[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx_d1[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_valid_d1;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_data_d1;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx_d1[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx_d1[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx_d1[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_valid_d1 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_valid) 
           & (0x00004000U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_addr));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_data_d1 
        = (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_valid) 
            & (0x00001000U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_addr))
            ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem
           [vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_word]
            : 0U);
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_req_addr_from_ctx 
        = (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[0U]);
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v0) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v0] 
            = ((0xffffff00U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem
                [__VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v0]) 
               | (IData)(__VdlyVal__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v0));
    }
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v1) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v1] 
            = ((0xffff00ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem
                [__VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v1]) 
               | ((IData)(__VdlyVal__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v1) 
                  << 8U));
    }
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v2) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v2] 
            = ((0xff00ffffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem
                [__VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v2]) 
               | ((IData)(__VdlyVal__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v2) 
                  << 0x00000010U));
    }
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v3) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v3] 
            = ((0x00ffffffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem
                [__VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v3]) 
               | ((IData)(__VdlyVal__tb_top__DOT__i_bar__DOT__i_bar0__DOT__bar_mem__v3) 
                  << 0x00000018U));
    }
}

void Vtop___024root___nba_sequent__TOP__6(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__6\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__device_reset) 
         & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__reset_d)))) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__generation 
            = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__generation);
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__lifecycle_generation 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__generation;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__reset_d 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__device_reset;
}

void Vtop___024root___nba_sequent__TOP__7(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__7\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    SData/*10:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr = 0;
    SData/*10:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_ptr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_ptr = 0;
    SData/*10:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count = 0;
    VlWide<5>/*133:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0;
    VL_ZERO_W(134, __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0);
    SData/*9:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0 = 0;
    VlWide<5>/*133:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1;
    VL_ZERO_W(134, __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1);
    SData/*9:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1 = 0;
    // Body
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count;
    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_ptr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_ptr;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0 = 0U;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1 = 0U;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__srst) {
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_ptr = 0U;
    } else {
        if ((2U != ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_en) 
                      & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__full))) 
                     << 1U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_en) 
                               & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__empty)))))) {
            if ((1U == ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_en) 
                          & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__full))) 
                         << 1U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_en) 
                                   & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__empty)))))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr 
                    = (0x000007ffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr)));
            } else if ((3U == ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_en) 
                                 & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__full))) 
                                << 1U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_en) 
                                          & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__empty)))))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr 
                    = (0x000007ffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr)));
            }
        }
        if ((2U == ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_en) 
                      & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__full))) 
                     << 1U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_en) 
                               & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__empty)))))) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count 
                = (0x000007ffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count)));
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0[0U] 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[0U];
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0[1U] 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[1U];
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0[2U] 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[2U];
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0[3U] 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[3U];
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0[4U] 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[4U];
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0 
                = (0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_ptr));
            __VdlySet__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0 = 1U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_ptr 
                = (0x000007ffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_ptr)));
        } else {
            if ((1U == ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_en) 
                          & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__full))) 
                         << 1U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_en) 
                                   & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__empty)))))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count 
                    = (0x000007ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count) 
                                      - (IData)(1U)));
            }
            if ((1U != ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_en) 
                          & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__full))) 
                         << 1U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_en) 
                                   & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__empty)))))) {
                if ((3U == ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_en) 
                              & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__full))) 
                             << 1U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_en) 
                                       & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__empty)))))) {
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1[0U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[0U];
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1[1U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[1U];
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1[2U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[2U];
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1[3U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[3U];
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1[4U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[4U];
                    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1 
                        = (0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_ptr));
                    __VdlySet__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1 = 1U;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_ptr 
                        = (0x000007ffU & ((IData)(1U) 
                                          + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_ptr)));
                }
            }
        }
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_ptr 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_ptr;
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0][0U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0[0U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0][1U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0[1U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0][2U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0[2U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0][3U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0[3U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0][4U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v0[4U];
    }
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1][0U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1[0U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1][1U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1[1U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1][2U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1[2U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1][3U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1[3U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1][4U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem__v1[4U];
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__full 
        = (0x0400U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_data_count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_data_count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__empty 
        = (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem
        [(0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr))][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem
        [(0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr))][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem
        [(0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr))][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[3U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem
        [(0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr))][3U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__dout[4U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__mem
        [(0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_ptr))][4U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_en) 
           & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__empty)));
}

void Vtop___024root___nba_sequent__TOP__8(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__8\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    CData/*4:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr__v0 = 0;
    CData/*2:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr__v0 = 0;
    CData/*4:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr__v0 = 0;
    CData/*2:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr__v0 = 0;
    CData/*5:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v0 = 0;
    CData/*2:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v0 = 0;
    VlWide<3>/*87:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem__v0;
    VL_ZERO_W(88, __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem__v0);
    CData/*4:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem__v0 = 0;
    CData/*2:0*/ __VdlyDim1__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem__v0;
    __VdlyDim1__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem__v0 = 0;
    IData/*31:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__data_mem__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__data_mem__v0 = 0;
    CData/*4:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__data_mem__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__data_mem__v0 = 0;
    CData/*2:0*/ __VdlyDim1__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__data_mem__v0;
    __VdlyDim1__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__data_mem__v0 = 0;
    CData/*4:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr__v1;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr__v1 = 0;
    CData/*2:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr__v1;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr__v1 = 0;
    CData/*4:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr__v1;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr__v1 = 0;
    CData/*2:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr__v1;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr__v1 = 0;
    CData/*5:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v1;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v1 = 0;
    CData/*2:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v1;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v1 = 0;
    CData/*5:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v2;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v2 = 0;
    CData/*2:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v2;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v2 = 0;
    CData/*5:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v3;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v3 = 0;
    CData/*2:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v3;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v3 = 0;
    // Body
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rst) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k = 0U;
        while (VL_GTS_III(32, 7U, vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h99f8a03b__0 = 0U;
            if (VL_LIKELY(((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k))))) {
                __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr__v0 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h99f8a03b__0;
                __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr__v0 
                    = (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k);
                vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr.enqueue(__VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr__v0, (IData)(__VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr__v0));
            }
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h0038da16__0 = 0U;
            if (VL_LIKELY(((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k))))) {
                __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr__v0 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h0038da16__0;
                __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr__v0 
                    = (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k);
                vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr.enqueue(__VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr__v0, (IData)(__VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr__v0));
            }
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h9dcf22b6__0 = 0U;
            if (VL_LIKELY(((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k))))) {
                __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v0 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h9dcf22b6__0;
                __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v0 
                    = (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k);
                vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count.enqueue(__VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v0, (IData)(__VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v0));
            }
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k 
                = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k);
        }
    } else {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k = 0U;
        while (VL_GTS_III(32, 7U, vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)) {
            if (((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)) 
                 && (1U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__push) 
                           >> (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k))))) {
                if ((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k))) {
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h511400fb__0[0U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ctx
                        [(7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)][0U];
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h511400fb__0[1U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ctx
                        [(7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)][1U];
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h511400fb__0[2U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ctx
                        [(7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)][2U];
                } else {
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h511400fb__0[0U] = 0U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h511400fb__0[1U] = 0U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h511400fb__0[2U] = 0U;
                }
                if (VL_LIKELY(((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k))))) {
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem__v0[0U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h511400fb__0[0U];
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem__v0[1U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h511400fb__0[1U];
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem__v0[2U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h511400fb__0[2U];
                    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem__v0 
                        = ((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k))
                            ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr
                           [(7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)]
                            : 0U);
                    __VdlyDim1__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem__v0 
                        = (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k);
                    vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem.enqueue(__VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem__v0, (IData)(__VdlyDim1__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem__v0), __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem__v0);
                }
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_hd1ae10b8__0 
                    = ((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k))
                        ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_data
                       [(7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)]
                        : 0U);
                if (VL_LIKELY(((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k))))) {
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__data_mem__v0 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_hd1ae10b8__0;
                    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__data_mem__v0 
                        = ((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k))
                            ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr
                           [(7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)]
                            : 0U);
                    __VdlyDim1__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__data_mem__v0 
                        = (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k);
                    vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__data_mem.enqueue(__VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__data_mem__v0, (IData)(__VdlyDim1__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__data_mem__v0), __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__data_mem__v0);
                }
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h99f8a03b__1 
                    = (0x0000001fU & ((IData)(1U) + 
                                      ((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k))
                                        ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr
                                       [(7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)]
                                        : 0U)));
                if (VL_LIKELY(((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k))))) {
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr__v1 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h99f8a03b__1;
                    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr__v1 
                        = (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k);
                    vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr.enqueue(__VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr__v1, (IData)(__VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr__v1));
                }
            }
            if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected_valid) 
                 & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected) 
                    == (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h0038da16__1 
                    = (0x0000001fU & ((IData)(1U) + 
                                      ((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k))
                                        ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr
                                       [(7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)]
                                        : 0U)));
                if (VL_LIKELY(((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k))))) {
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr__v1 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h0038da16__1;
                    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr__v1 
                        = (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k);
                    vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr.enqueue(__VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr__v1, (IData)(__VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr__v1));
                }
            }
            if ((2U == ((((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)) 
                          && (1U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__push) 
                                    >> (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)))) 
                         << 1U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected_valid) 
                                   & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected) 
                                      == (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)))))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h9dcf22b6__1 
                    = (0x0000003fU & ((IData)(1U) + 
                                      ((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k))
                                        ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count
                                       [(7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)]
                                        : 0U)));
                if (VL_LIKELY(((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k))))) {
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v1 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h9dcf22b6__1;
                    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v1 
                        = (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k);
                    vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count.enqueue(__VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v1, (IData)(__VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v1));
                }
            } else if ((1U == ((((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)) 
                                 && (1U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__push) 
                                           >> (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)))) 
                                << 1U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected_valid) 
                                          & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected) 
                                             == (7U 
                                                 & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)))))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h9dcf22b6__2 
                    = (0x0000003fU & (((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k))
                                        ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count
                                       [(7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)]
                                        : 0U) - (IData)(1U)));
                if (VL_LIKELY(((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k))))) {
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v2 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h9dcf22b6__2;
                    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v2 
                        = (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k);
                    vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count.enqueue(__VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v2, (IData)(__VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v2));
                }
            } else {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h9dcf22b6__3 
                    = ((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k))
                        ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count
                       [(7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k)]
                        : 0U);
                if (VL_LIKELY(((6U >= (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k))))) {
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v3 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT____Vlvbound_h9dcf22b6__3;
                    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v3 
                        = (7U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k);
                    vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count.enqueue(__VdlyVal__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v3, (IData)(__VdlyDim0__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count__v3));
                }
            }
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k 
                = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__k);
        }
    }
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rst) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rr_ptr = 0U;
    } else if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected_valid) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rr_ptr 
            = ((6U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected))
                ? 0U : (7U & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected))));
    }
    vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr.commit(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__wr_ptr);
    vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__data_mem.commit(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__data_mem);
    vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem.commit(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem);
    vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr.commit(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr);
    vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count.commit(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ready 
        = ((((((0x20U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[6U]) 
               << 3U) | ((0x20U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[5U]) 
                         << 2U)) | (((0x20U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[4U]) 
                                     << 1U) | (0x20U 
                                               > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[3U]))) 
            << 3U) | (((0x20U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[2U]) 
                       << 2U) | (((0x20U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[1U]) 
                                  << 1U) | (0x20U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[0U]))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected_valid = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__scan = 0U;
    while (VL_GTS_III(32, 7U, vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__scan)) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__index 
            = (0x0000000fU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rr_ptr) 
                              + vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__scan));
        if ((7U <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__index))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__index 
                = (0x0000000fU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__index) 
                                  - (IData)(7U)));
        }
        if (((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected_valid)) 
             & (0U != ((6U >= (7U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__index)))
                        ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count
                       [(7U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__index))]
                        : 0U)))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected 
                = (7U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__index));
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected_valid = 1U;
        }
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__scan 
            = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__scan);
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ready 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ready;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected_valid) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_valid = 1U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_data 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__data_mem
            [((6U >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected))
               ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected)
               : 0U)][((6U >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected))
                        ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr
                       [vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected]
                        : 0U)];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_ctx[0U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem
            [((6U >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected))
               ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected)
               : 0U)][((6U >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected))
                        ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr
                       [vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected]
                        : 0U)][0U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_ctx[1U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem
            [((6U >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected))
               ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected)
               : 0U)][((6U >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected))
                        ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr
                       [vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected]
                        : 0U)][1U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_ctx[2U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__ctx_mem
            [((6U >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected))
               ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected)
               : 0U)][((6U >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected))
                        ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rd_ptr
                       [vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__selected]
                        : 0U)][2U];
    } else {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_valid = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_data = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_ctx[0U] = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_ctx[1U] = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_ctx[2U] = 0U;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ready 
        = (1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ready));
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_rsp_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_rsp_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_rsp_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_rsp_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_rsp_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__out_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_rsp_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_rsp_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_rsp_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_rsp_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_rsp_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_data_swapped 
        = ((((0x0000ff00U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_data 
                             << 8U)) | (0x000000ffU 
                                        & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_data 
                                           >> 8U))) 
            << 0x00000010U) | ((0x0000ff00U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_data 
                                               >> 8U)) 
                               | (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_data 
                                  >> 0x00000018U)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_first 
        = (1U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[2U] 
                 >> 0x00000017U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_last 
        = (1U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[2U] 
                 >> 0x00000016U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_dwlen 
        = (0x000007ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[2U] 
                          >> 0x0000000bU));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_byte_count 
        = (0x00000fffU & ((vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[2U] 
                           << 1U) | (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[1U] 
                                     >> 0x0000001fU)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_tc 
        = (7U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[1U] 
                 >> 0x0000001cU));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_attr 
        = (7U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[1U] 
                 >> 0x00000019U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_tag 
        = (0x000000ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[1U] 
                          >> 0x00000010U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_requester 
        = (0x0000ffffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[1U]);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_lower_addr 
        = (0x0000007fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[0U]);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_4dw 
        = (1U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_ctx[1U] 
                 >> 0x00000018U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__unused_rd_rsp_4dw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_4dw;
}

void Vtop___024root___nba_sequent__TOP__9(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__9\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__busy;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ctx 
        = vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ready 
        = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy)));
}

void Vtop___024root___nba_comb__TOP__0(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_comb__TOP__0\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr_last) 
                                                   | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_first) 
                                                      & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr_ready) 
                                                         & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_posted_write) 
                                                            & (0U 
                                                               != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_enabled_byte_count))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__is_sequential 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_valid) 
           & (((vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_addr 
                >> 8U) == (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_addr 
                           >> 8U)) | ((vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_addr 
                                       >> 8U) == (0x00ffffffU 
                                                  & ((IData)(1U) 
                                                     + 
                                                     (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_addr 
                                                      >> 8U))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_accepted 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy_d1)) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlps_in_valid 
        = ((IData)(vlSymsp->TOP__tb_top__DOT__tlps_in_if.__VdfgRegularize_hebeb780c_0_3) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr) 
              & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_request_supported) 
                 | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr_last))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__corr_weight 
        = (0x80U & (- (IData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__is_sequential))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__base_jitter 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_latency) 
           & (- (IData)((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__is_sequential))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_pop 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_empty)) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_accepted));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__wr_en 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlps_in_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_adj 
        = (0x000000ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__base_jitter) 
                          + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj 
        = (0x000000ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adjust) 
                          + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_adj)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__jittered 
        = ((0x0000000fU < ((IData)(3U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj)))
            ? 0x0000000fU : ((3U > ((IData)(3U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj)))
                              ? 3U : (0x000000ffU & 
                                      ((IData)(3U) 
                                       + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__computed_latency 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__jittered;
}

void Vtop___024root___nba_comb__TOP__1(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_comb__TOP__1\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__issue_dw 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__stall)) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__state));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__issue_dw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_req_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar) 
              >> 1U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_req_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar) 
              >> 2U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_req_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar) 
              >> 3U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_req_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar) 
              >> 4U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_req_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar) 
              >> 5U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_rd_hit) 
              & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar)));
}

void Vtop___024root___nba_sequent__TOP__10(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__10\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_word 
        = (0x000003ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_addr 
                          >> 2U));
}

void Vtop___024root___nba_sequent__TOP__11(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__11\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr_bar0;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_en 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tvalid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[3U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[3U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[4U] 
        = (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__first) 
            << 5U) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tlast) 
                       << 4U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_word 
        = (0x000003ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_addr 
                          >> 2U));
}

void Vtop___024root___nba_comb__TOP__2(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_comb__TOP__2\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_write 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_full)) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_valid) 
              | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_valid)));
}

void Vtop___024root___nba_comb__TOP__5(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_comb__TOP__5\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__push 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_valid) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ready));
}

void Vtop_IfAXIS128___nba_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__0(Vtop_IfAXIS128* vlSelf);

void Vtop___024root___eval_nba(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_nba\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    if ((8ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__0(vlSelf);
    }
    if ((0x0000000000000200ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__1(vlSelf);
    }
    if ((0x0000000000000020ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__2(vlSelf);
        Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_ur__0((&vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_ur));
    }
    if ((2ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__3(vlSelf);
    }
    if ((0x0000000000000040ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__4(vlSelf);
    }
    if ((0x0000000000000100ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__5(vlSelf);
    }
    if ((4ULL & vlSelfRef.__VnbaTriggered[0U])) {
        if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__device_reset) 
             & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__reset_d)))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__generation 
                = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__generation);
        }
        vlSelfRef.tb_top__DOT__i_bar__DOT__lifecycle_generation 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__generation;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__reset_d 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__device_reset;
    }
    if ((0x0000000000000010ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__7(vlSelf);
        Vtop_IfAXIS128___nba_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__0((&vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng));
    }
    if ((0x0000000000000080ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__8(vlSelf);
    }
    if ((0x0000000000000200ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__busy;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ctx 
            = vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_data 
            = vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data;
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ready 
            = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy)));
    }
    if ((0x0000000000000202ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr_last) 
               | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_first) 
                  & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr_ready) 
                     & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_posted_write) 
                        & (0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_enabled_byte_count))))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__is_sequential 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_valid) 
               & (((vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_addr 
                    >> 8U) == (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_addr 
                               >> 8U)) | ((vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_addr 
                                           >> 8U) == 
                                          (0x00ffffffU 
                                           & ((IData)(1U) 
                                              + (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_addr 
                                                 >> 8U))))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_accepted 
            = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy_d1)) 
               & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlps_in_valid 
            = ((IData)(vlSymsp->TOP__tb_top__DOT__tlps_in_if.__VdfgRegularize_hebeb780c_0_3) 
               & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr) 
                  & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_request_supported) 
                     | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr_last))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__corr_weight 
            = (0x80U & (- (IData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__is_sequential))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__base_jitter 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_latency) 
               & (- (IData)((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__is_sequential))))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_pop 
            = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_empty)) 
               & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_accepted));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__wr_en 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlps_in_valid;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_adj 
            = (0x000000ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__base_jitter) 
                              + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj 
            = (0x000000ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adjust) 
                              + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_adj)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__jittered 
            = ((0x0000000fU < ((IData)(3U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj)))
                ? 0x0000000fU : ((3U > ((IData)(3U) 
                                        + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj)))
                                  ? 3U : (0x000000ffU 
                                          & ((IData)(3U) 
                                             + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj)))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__computed_latency 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__jittered;
    }
    if ((0x000000000000000aULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__issue_dw 
            = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__stall)) 
               & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__state));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_valid 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__issue_dw;
        vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_valid;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar1__DOT__rd_req_valid 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
               & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar) 
                  >> 1U));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar2__DOT__rd_req_valid 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
               & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar) 
                  >> 2U));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar3__DOT__rd_req_valid 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
               & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar) 
                  >> 3U));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar4__DOT__rd_req_valid 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
               & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar) 
                  >> 4U));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar5__DOT__rd_req_valid 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
               & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar) 
                  >> 5U));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_valid 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
               & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_rd_hit) 
                  & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar)));
    }
    if ((0x0000000000000040ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_valid 
            = vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_be;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_addr 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_word 
            = (0x000003ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_addr 
                              >> 2U));
    }
    if ((8ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[0U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[1U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[2U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[2U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_addr 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr_bar0;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_en 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tvalid;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[0U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[0U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[1U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[1U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[2U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[2U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[3U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tdata[3U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__din[4U] 
            = (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__first) 
                << 5U) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tlast) 
                           << 4U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__tkeepdw)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_word 
            = (0x000003ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_addr 
                              >> 2U));
    }
    if ((0x0000000000000102ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_write 
            = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_full)) 
               & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_valid) 
                  | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_valid)));
    }
    if ((0x0000000000000032ULL & vlSelfRef.__VnbaTriggered[0U])) {
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
    if ((0x0000000000000280ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__push 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_valid) 
               & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ready));
    }
}

void Vtop___024root___timing_resume(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___timing_resume\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    if ((0x0000000000000400ULL & vlSelfRef.__VactTriggered[0U])) {
        vlSelfRef.__VdlySched.resume();
    }
}

void Vtop___024root___trigger_orInto__act_vec_vec(VlUnpacked<QData/*63:0*/, 1> &out, const VlUnpacked<QData/*63:0*/, 1> &in) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___trigger_orInto__act_vec_vec\n"); );
    // Locals
    IData/*31:0*/ n;
    // Body
    n = 0U;
    do {
        out[n] = (out[n] | in[n]);
        n = ((IData)(1U) + n);
    } while ((0U >= n));
}

#ifdef VL_DEBUG
VL_ATTR_COLD void Vtop___024root___dump_triggers__act(const VlUnpacked<QData/*63:0*/, 1> &triggers, const std::string &tag);
#endif  // VL_DEBUG

bool Vtop___024root___eval_phase__act(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_phase__act\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    CData/*0:0*/ __VactExecute;
    // Body
    Vtop___024root___eval_triggers_vec__act(vlSelf);
    Vtop___024root___trigger_orInto__act_vec_vec(vlSelfRef.__VactTriggered, vlSelfRef.__VactTriggeredAcc);
#ifdef VL_DEBUG
    if (VL_UNLIKELY(vlSymsp->_vm_contextp__->debug())) {
        Vtop___024root___dump_triggers__act(vlSelfRef.__VactTriggered, "act"s);
    }
#endif
    Vtop___024root___trigger_orInto__act_vec_vec(vlSelfRef.__VnbaTriggered, vlSelfRef.__VactTriggered);
    __VactExecute = Vtop___024root___trigger_anySet__act(vlSelfRef.__VactTriggered);
    if (__VactExecute) {
        vlSelfRef.__VactTriggeredAcc.fill(0ULL);
        Vtop___024root___timing_resume(vlSelf);
        Vtop___024root___eval_act(vlSelf);
    }
    return (__VactExecute);
}

bool Vtop___024root___eval_phase__inact(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_phase__inact\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    CData/*0:0*/ __VinactExecute;
    // Body
    __VinactExecute = vlSelfRef.__VdlySched.awaitingZeroDelay();
    if (__VinactExecute) {
        VL_FATAL_MT("tb_top.sv", 4, "", "ZERODLY: Design Verilated with '--no-sched-zero-delay', but #0 delay executed at runtime");
    }
    return (__VinactExecute);
}

void Vtop___024root___trigger_clear__act(VlUnpacked<QData/*63:0*/, 1> &out) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___trigger_clear__act\n"); );
    // Locals
    IData/*31:0*/ n;
    // Body
    n = 0U;
    do {
        out[n] = 0ULL;
        n = ((IData)(1U) + n);
    } while ((1U > n));
}

bool Vtop___024root___eval_phase__nba(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_phase__nba\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    CData/*0:0*/ __VnbaExecute;
    // Body
    __VnbaExecute = Vtop___024root___trigger_anySet__act(vlSelfRef.__VnbaTriggered);
    if (__VnbaExecute) {
        Vtop___024root___eval_nba(vlSelf);
        Vtop___024root___trigger_clear__act(vlSelfRef.__VnbaTriggered);
    }
    return (__VnbaExecute);
}

void Vtop___024root___eval(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    IData/*31:0*/ __VicoIterCount;
    IData/*31:0*/ __VnbaIterCount;
    // Body
    __VicoIterCount = 0U;
    vlSelfRef.__VicoFirstIteration = 1U;
    do {
        if (VL_UNLIKELY(((0x00002710U < __VicoIterCount)))) {
#ifdef VL_DEBUG
            Vtop___024root___dump_triggers__ico(vlSelfRef.__VicoTriggered, "ico"s);
#endif
            VL_FATAL_MT("tb_top.sv", 4, "", "DIDNOTCONVERGE: Input combinational region did not converge after '--converge-limit' of 10000 tries");
        }
        __VicoIterCount = ((IData)(1U) + __VicoIterCount);
        vlSelfRef.__VicoPhaseResult = Vtop___024root___eval_phase__ico(vlSelf);
        vlSelfRef.__VicoFirstIteration = 0U;
    } while (vlSelfRef.__VicoPhaseResult);
    __VnbaIterCount = 0U;
    do {
        if (VL_UNLIKELY(((0x00002710U < __VnbaIterCount)))) {
#ifdef VL_DEBUG
            Vtop___024root___dump_triggers__act(vlSelfRef.__VnbaTriggered, "nba"s);
#endif
            VL_FATAL_MT("tb_top.sv", 4, "", "DIDNOTCONVERGE: NBA region did not converge after '--converge-limit' of 10000 tries");
        }
        __VnbaIterCount = ((IData)(1U) + __VnbaIterCount);
        vlSelfRef.__VinactIterCount = 0U;
        do {
            if (VL_UNLIKELY(((0x00002710U < vlSelfRef.__VinactIterCount)))) {
                VL_FATAL_MT("tb_top.sv", 4, "", "DIDNOTCONVERGE: Inactive region did not converge after '--converge-limit' of 10000 tries");
            }
            vlSelfRef.__VinactIterCount = ((IData)(1U) 
                                           + vlSelfRef.__VinactIterCount);
            vlSelfRef.__VactIterCount = 0U;
            do {
                if (VL_UNLIKELY(((0x00002710U < vlSelfRef.__VactIterCount)))) {
#ifdef VL_DEBUG
                    Vtop___024root___dump_triggers__act(vlSelfRef.__VactTriggered, "act"s);
#endif
                    VL_FATAL_MT("tb_top.sv", 4, "", "DIDNOTCONVERGE: Active region did not converge after '--converge-limit' of 10000 tries");
                }
                vlSelfRef.__VactIterCount = ((IData)(1U) 
                                             + vlSelfRef.__VactIterCount);
                vlSelfRef.__VactPhaseResult = Vtop___024root___eval_phase__act(vlSelf);
            } while (vlSelfRef.__VactPhaseResult);
            vlSelfRef.__VinactPhaseResult = Vtop___024root___eval_phase__inact(vlSelf);
        } while (vlSelfRef.__VinactPhaseResult);
        vlSelfRef.__VnbaPhaseResult = Vtop___024root___eval_phase__nba(vlSelf);
    } while (vlSelfRef.__VnbaPhaseResult);
}

#ifdef VL_DEBUG
void Vtop___024root___eval_debug_assertions(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_debug_assertions\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
}
#endif  // VL_DEBUG
