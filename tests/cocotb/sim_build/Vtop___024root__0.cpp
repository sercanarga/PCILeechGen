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
                                                    ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en) 
                                                       != (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en__0)) 
                                                      << 1U) 
                                                     | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_select_ur) 
                                                        != (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__cpl_select_ur__0)))));
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__cpl_select_ur__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_select_ur;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en;
    if (VL_UNLIKELY(((1U & (~ (IData)(vlSelfRef.__VicoDidInit)))))) {
        vlSelfRef.__VicoDidInit = 1U;
        vlSelfRef.__VicoTriggered[0U] = (1ULL | vlSelfRef.__VicoTriggered[0U]);
        vlSelfRef.__VicoTriggered[0U] = (2ULL | vlSelfRef.__VicoTriggered[0U]);
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
    // Locals
    CData/*0:0*/ tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid;
    tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid = 0;
    CData/*7:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__35__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__35__Vfuncout = 0;
    QData/*63:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__35__lba;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__35__lba = 0;
    SData/*14:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__40__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__40__Vfuncout = 0;
    QData/*63:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__40__lba;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__40__lba = 0;
    CData/*6:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__40__word_off;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__40__word_off = 0;
    CData/*7:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__40____VlefCall_0__backing_slot;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__40____VlefCall_0__backing_slot = 0;
    CData/*7:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__41__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__41__Vfuncout = 0;
    QData/*63:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__41__lba;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__41__lba = 0;
    CData/*0:0*/ __VdfgRegularize_hebeb780c_0_10;
    __VdfgRegularize_hebeb780c_0_10 = 0;
    CData/*0:0*/ __VdfgRegularize_hebeb780c_0_11;
    __VdfgRegularize_hebeb780c_0_11 = 0;
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_delivery_done 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_done;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_power2_entries 
        = (0U == ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_last) 
                  & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_last))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_power2_entries 
        = (0U == ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last) 
                  & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cid 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[0U] 
           >> 0x00000010U);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_disk_word 
        = (0x0000007fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx);
    vlSelfRef.tb_top__DOT__i_bar__DOT__lifecycle_generation 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__generation;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_ready 
        = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_data_count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_data_count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__full 
        = (0x0400U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_data_count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__wr_data_count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__count;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_wr_ack 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_ack;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_delivery_vector 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_vector;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_event_vector 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__pba_set_vector;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_id_rom_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_next 
        = (0x000000ffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_idx)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_outcome_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_outcome_tag 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_tag;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_index 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_tag) 
           - (IData)(0x00000010U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_outcome_status 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_status;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_last_packet 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_length_dw) 
           == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_base 
        = (((QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_hi)) 
            << 0x00000020U) | (QData)((IData)((0xfffff000U 
                                               & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_lo))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail_next 
        = (0x0000ffffU & (((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail)) 
                          & (- (IData)(((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail) 
                                        < (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_disk_lba 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_slba 
           + (QData)((IData)((0x0001ffffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx 
                                             >> 7U)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_almost_full 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_valid) 
           | (0x0eU <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_count)));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_completion_class 
        = ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status))
            ? ((4U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_class))
                ? ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_read_seen_miss)
                    ? 6U : 5U) : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_class))
            : 0x0cU);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_state 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_active_qid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_id_rom_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_id_rom_data_r;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_csts 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000001C;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_event_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__pba_set_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_delivery_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_req 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_req;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_len 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_len;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_done 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_done;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_req 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_req;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_done 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_done;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__dbg_state 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_done 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_done;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_rdata 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_rdata;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_hit 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_hit;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_busy 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__busy;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_error 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__error;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__smart_temp_k 
        = (0x0000ffffU & ((IData)(0x012fU) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__thermal_load)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_next_index 
        = (0x000003ffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_index)));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_full 
        = (0x10U <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_nlb64 
        = (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw1));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_base 
        = (((QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_hi)) 
            << 0x00000020U) | (QData)((IData)((0xfffff000U 
                                               & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_lo))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_valid 
        = ((0U == (0x00000fffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_lo)) 
           & ((0U != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_hi) 
              | (0U != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_lo)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_valid 
        = ((0U == (0x00000fffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_lo)) 
           & ((0U != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_hi) 
              | (0U != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_lo)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head_next 
        = (0x0000ffffU & (((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head)) 
                          & (- (IData)(((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head) 
                                        < (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_last))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode 
        = (0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[0U]);
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_intr_req 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__msi_pulse;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tuser 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tuser;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tkeepdw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tkeepdw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tdata[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tdata[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tdata[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tdata[3U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[3U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_aqa 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_asq_lo 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_asq_hi 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000002C;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_acq_lo 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_acq_hi 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000034;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_empty 
        = (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ready 
        = ((((((0x20U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[6U]) 
               << 3U) | ((0x20U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[5U]) 
                         << 2U)) | (((0x20U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[4U]) 
                                     << 1U) | (0x20U 
                                               > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[3U]))) 
            << 3U) | (((0x20U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[2U]) 
                       << 2U) | (((0x20U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[1U]) 
                                  << 1U) | (0x20U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__count[0U]))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__dout;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout1 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__dout;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout2 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__dout;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout3 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__dout;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_fifo_full 
        = (0x0100U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp2 
        = (((QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[9U])) 
            << 0x00000020U) | (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[8U])));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_pop 
        = (1U & (~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__state) 
                    | (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__req_count)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw13 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[13U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw14 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[14U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw15 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[15U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_total_bytes 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_total_dw 
           << 2U);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_byte_off 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx 
           << 2U);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_timing_info 
        = (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_saturated) 
            << 0x0000001fU) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_effective_class) 
                                << 0x0000001bU) | (0x07ffffffU 
                                                   & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_wait_cycles)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot_dbg 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_fifo_full 
        = (0x0100U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__request_count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acqs 
        = (0x00000fffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_aqa 
                          >> 0x00000010U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_last_dw 
        = ((0x00ffffffU & ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx)) 
           >= vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_total_dw);
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tlast 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tlast;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tvalid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_bar 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_bar;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s0 
           + vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s3);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__issue_dw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__state;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__empty 
        = (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_last 
        = (1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_remaining_dw));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_cc 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_format_slot 
        = (0x000000ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx) 
                          >> 7U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__full 
        = (0x0400U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw11 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[11U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asqs 
        = (0x00000fffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_aqa);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1 
        = (((QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[7U])) 
            << 0x00000020U) | (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[6U])));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_addr 
        = (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw12 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[12U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be 
        = (0x0000000fU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_packet)
                           ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_be)
                           : ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_last_be) 
                              | (- (IData)((1U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw)))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_word_cached 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_hit) 
           & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid
              [vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot] 
              & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_lba 
                 == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag
                 [vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot])));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_addr_w 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_addr
        [vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_rd];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem
        [(0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr))][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem
        [(0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr))][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem
        [(0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr))][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[3U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem
        [(0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr))][3U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[4U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem
        [(0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr))][4U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_io 
        = (0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid));
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[10U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__clk = vlSelfRef.tb_top__DOT__clk;
    __VdfgRegularize_hebeb780c_0_11 = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_error)) 
                                       & (3U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_bar = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_bar;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_word 
        = (0x0000007fU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_clear 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_error)) 
           & (2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_lba 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_base_lba 
           + (QData)((IData)((0x000001ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx) 
                                             >> 7U)))));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__rst = vlSelfRef.tb_top__DOT__rst;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_lba 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_lba;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_write 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_write;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_flush 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_flush;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_word 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_word;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_wdata 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_wdata;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_resp_dbg_state 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_state;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_resp_dbg_qid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_active_qid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_id_rom_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_csts 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_csts;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__event_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_event_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_delivery_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_req 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_req;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_len 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_len;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_done 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_done;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_req 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_req;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_done 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_done;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_dbg_state 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__dbg_state;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_done 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_done;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_rdata 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_rdata;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_hit 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_hit;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_busy 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_busy;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_error 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_error;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__smart_warning_byte 
        = (((0U != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_media_errors) 
            << 2U) | ((0x0157U <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__smart_temp_k)) 
                      << 1U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_next_addr 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_base_addr 
           + (QData)((IData)(((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_next_index) 
                              << 2U))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__rd_en 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_rd_en;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_opcode 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode;
    vlSelfRef.tb_top__DOT__i_bar__DOT__intr_req = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_intr_req;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_aqa 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_aqa;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_asq_lo 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_asq_lo;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_asq_hi 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_asq_hi;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_acq_lo 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_acq_lo;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_acq_hi 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_acq_hi;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_valid 
        = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_empty)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ready 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ready;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__ram_dout 
        = ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__read_bank_q))
            ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout0
            : ((1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__read_bank_q))
                ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout1
                : ((2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__read_bank_q))
                    ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout2
                    : vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout3)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_can_accept 
        = (1U & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_fifo_full)) 
                 | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_pop)));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_last 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acqs;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_io_progress 
        = (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status) 
            << 0x00000018U) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dir) 
                                << 0x00000016U) | (
                                                   ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_last_dw) 
                                                    << 0x00000015U) 
                                                   | (0x001fffffU 
                                                      & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tdata[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tdata[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tdata[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din[3U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tdata[3U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din[4U] 
        = (0x0000003fU & ((0x00000020U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tuser) 
                                          << 5U)) | 
                          (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tlast) 
                            << 4U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tkeepdw))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_write 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_full)) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_valid) 
              | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_valid)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_bar;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_latency 
        = (0x000000ffU & ((IData)(2U) + VL_MODDIV_III(32, 
                                                      (0x000000ffU 
                                                       & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out 
                                                          >> 8U)), (IData)(5U))));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__issue_dw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_cc 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_cc;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid_selects_ns1 
        = ((1U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid) 
           | (0xffffffffU == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid_lists_ns1 
        = ((0xffffffffU == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid) 
           | (0U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_full 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__full;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_aqa_valid 
        = ((1U <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asqs)) 
           & (1U <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acqs)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_last 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asqs;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1_page_valid 
        = ((0ULL != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1) 
           & (0U == (0x00000fffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp1_invalid 
        = ((0U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_total_dw) 
           | ((0ULL == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1) 
              | (0U != (3U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_first_span 
        = ((IData)(0x00001000U) - (0x00000fffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count 
        = (0x0001ffffU & ((IData)(1U) + (0x0000ffffU 
                                         & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw12)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first_offset 
        = (3U & (((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be))
                   ? 1U : ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be))
                            ? 2U : (- (IData)((1U & 
                                               ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be) 
                                                >> 3U)))))) 
                 & (- (IData)((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be)))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_addr_w;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tkeepdw 
        = (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[4U]);
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tdata[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tdata[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tdata[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tdata[3U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[3U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_first 
        = (1U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[4U] 
                 >> 5U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tlast 
        = (1U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[4U] 
                 >> 4U));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_val 
        = (0x0000ffffU & vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_log_last_dw 
        = ((0x03ffU < (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10 
                       >> 0x00000010U)) ? 0x000003ffU
            : (0x000007ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10 
                              >> 0x00000010U)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_dsm_ranges 
        = (0x000000ffU & ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__format_nvm_ok 
        = (IData)((((((((1U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid) 
                        & (0U == (0xffffffefU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))) 
                       & (0U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw11)) 
                      & (0U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw12)) 
                     & (0U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw13)) 
                    & (0U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw14)) 
                   & (0U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw15)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_cns 
        = (0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_slba 
        = (((QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw11)) 
            << 0x00000020U) | (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10)));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__clk 
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_read 
        = ((IData)(__VdfgRegularize_hebeb780c_0_11) 
           & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_write)) 
              & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_word_cached)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_write 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_write) 
           & (IData)(__VdfgRegularize_hebeb780c_0_11));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_req_addr_from_ctx 
        = (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[0U]);
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_wdata 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_wdata 
           & (- (IData)((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_clear))))));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_wr_hit 
        = (0x00004000U > vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr);
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__fundamental_reset 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rst;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_lba 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_lba;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_write 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_write;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_flush 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_flush;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_word 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_word;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_wdata;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_resp_dbg_opcode 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_opcode;
    vlSelfRef.tb_top__DOT__intr_req = vlSelfRef.tb_top__DOT__i_bar__DOT__intr_req;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aqa 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_aqa;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_lo 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_asq_lo;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_hi 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_asq_hi;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_lo 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_acq_lo;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_hi 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_acq_hi;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ready 
        = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_accepted 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy_d1)) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_power2_entries 
        = (0U == ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_last) 
                  & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_last))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail_next 
        = (0x0000ffffU & (((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail)) 
                          & (- (IData)(((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail) 
                                        < (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_last))))));
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
        = (0x000000ffU & ((IData)(3U) + (((IData)(0x0000000aU) 
                                          * (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket)) 
                                         >> 4U)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en 
        = (1U & vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_cc);
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_valid) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw3_value 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_data;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_entry_hi_value 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_data;
    } else {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw3_value 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw3;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_entry_hi_value 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_entry_hi;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tready 
        = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_full)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_wr_en 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_full)) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tvalid));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_queue_config_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_aqa_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_valid) 
              & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_valid)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_power2_entries 
        = (0U == ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_last) 
                  & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_last))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head_next 
        = (0x0000ffffU & (((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head)) 
                          & (- (IData)(((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head) 
                                        < (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_last))))));
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_io) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_tail 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_base 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_base;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_phase 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_phase;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_last 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_admin_queues 
            = ((vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw12 
                << 0x00000010U) | ((0x0000ff00U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head) 
                                                   << 8U)) 
                                   | (0x000000ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_tail))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_sq_head_next 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head_next;
    } else {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_tail 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_base 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_base;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_phase 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_phase;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_last 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_last;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_admin_queues 
            = ((((0x0000ff00U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head) 
                                 << 8U)) | (0x000000ffU 
                                            & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_tail))) 
                << 0x00000010U) | ((0x0000ff00U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_last) 
                                                   << 8U)) 
                                   | (0x000000ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_head_db))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_sq_head_next 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head_next;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_after_first_page 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_byte_off 
           >= vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_first_span);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_uses_list 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_total_bytes 
           > ((IData)(0x00001000U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_first_span));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp2_required 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_total_bytes 
           > vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_first_span);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_after_first 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_byte_off 
           - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_first_span);
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_total_dw 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count 
           << 7U);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count64 
        = (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_val 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_val;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_io_cmd_info 
        = ((2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))
            ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_slba)
            : vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_io_progress);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk;
    __VdfgRegularize_hebeb780c_0_10 = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_clear) 
                                       | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_write));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[0U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[0U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[0U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__din 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_wdata;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__din 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_wdata;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__din 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_wdata;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__din 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_wdata;
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
    tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_wr_hit) 
              & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_bar)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_off 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0 
           - (IData)(0x00001000U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__device_reset 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__fundamental_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_detect_now 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_write) 
              & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_flush)) 
                 & ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_word)) 
                    & (((0xebU == (0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata)) 
                        | (0xe9U == (0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata))) 
                       & (0x4eU == (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata 
                                    >> 0x00000018U)))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ready 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ready;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_pop 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_empty)) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_accepted));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_valid 
        = (0x0000003fU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_valid));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_lba_value 
        = (((QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw3_value)) 
            << 0x00000020U) | (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw2)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_list_page_base 
        = (((QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_entry_hi_value)) 
            << 0x00000020U) | (QData)((IData)((0xfffff000U 
                                               & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_entry_lo))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_list_entry_value 
        = (((QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_entry_hi_value)) 
            << 0x00000020U) | (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_entry_lo)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tready 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tready;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_en 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_wr_en;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_inc 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_wr_en) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tlast));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_resp_dbg_admin_queues 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_admin_queues;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp2_invalid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp2_required) 
           & ((0ULL == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp2) 
              | (0U != (0x00000fffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp2)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_list_index 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_after_first 
           >> 0x0000000cU);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_rd_hit) 
              & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr_bar0;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_mdts_ok 
        = (0x00002000U >= vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_total_dw);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_lba_ok 
        = ((1U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid) 
           & ((0x00000000773bd2b0ULL > vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_slba) 
              & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count64 
                 <= (0x00000000773bd2b0ULL - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_slba))));
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
        = (0xa0U & (- (IData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__is_sequential))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__base_jitter 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_latency) 
           & (- (IData)((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__is_sequential))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_cmd_info 
        = (((0x1fU == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state)) 
            | ((0x20U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state)) 
               | (0x1eU == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))))
            ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_timing_info
            : ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_io)
                ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_io_cmd_info
                : ((((0x0000ff00U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status) 
                                     << 8U)) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_cns)) 
                    << 0x00000010U) | ((0x0000ff00U 
                                        & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid 
                                           << 8U)) 
                                       | (0x000000ffU 
                                          & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_we 
        = __VdfgRegularize_hebeb780c_0_10;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_active 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_read) 
           | (IData)(__VdfgRegularize_hebeb780c_0_10));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_valid 
        = tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_valid 
        = tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_wr = 
        ((IData)(tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid) 
         & ((0x00001000U <= vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0) 
            & (0x00002000U > vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_cc_wr = 
        ((0x00000014U == (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0)) 
         & (IData)(tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_in_range 
        = (0x00004000U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_addr);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_offset 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_index 
        = VL_SHIFTR_III(32,32,32, vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_off, 2U);
    vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__device_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__active 
        = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__device_reset)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_route_now 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_armed)) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_detect_now));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_ok 
        = ((0U != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw1) 
           & ((0x00000000773bd2b0ULL > vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_lba_value) 
              & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_nlb64 
                 <= (0x00000000773bd2b0ULL - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_lba_value))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_list_entry_invalid 
        = ((0ULL == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_list_entry_value) 
           | (0U != (0x00000fffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_list_entry_value))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_list_entry_addr 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp2 
           + (QData)((IData)((vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_list_index 
                              << 3U))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_in_range 
        = (0x00004000U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_addr);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_adj 
        = (0x000000ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__base_jitter) 
                          + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_resp_dbg_cmd_info 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_cmd_info;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_wr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_wr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_cc_enable_wr 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_cc_wr) 
           & vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data);
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_cc_disable_wr 
        = ((~ vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_cc_wr));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_is_cq 
        = (1U & vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_index);
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_qid 
        = (0x0000ffffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_index 
                          >> 1U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__rst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__rst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__rst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__rst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__rst 
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__io_enabled 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__active;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__dma_enabled 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__active;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__eff_boot_armed 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_route_now) 
           | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_armed));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__eff_pin_boot 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_route_now)
            ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_lba
            : vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pin_boot_lba);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__push 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_valid) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ready));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ctx 
        = vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__unused_rd_rsp_4dw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_rsp_4dw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_enable_wr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_cc_enable_wr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_disable_wr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_cc_disable_wr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_is_cq 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_is_cq;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_qid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_qid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__srst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rst;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__srst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__rst;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__rst 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rst;
    vlSelfRef.tb_top__DOT__i_bar__DOT__io_enabled = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__io_enabled;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__quiesce 
        = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__dma_enabled)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__dma_enabled 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__dma_enabled;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_full_format) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT____VlemCall_1__backing_slot 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_format_slot;
    } else {
        vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__37__lba 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_lba;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mix 
            = (0x00000fffU & ((((((IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__37__lba) 
                                  ^ (IData)((vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__37__lba 
                                             >> 0x0cU))) 
                                 ^ (IData)((vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__37__lba 
                                            >> 0x18U))) 
                                ^ (IData)((vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__37__lba 
                                           >> 0x24U))) 
                               ^ (IData)((vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__37__lba 
                                          >> 0x30U))) 
                              ^ (0x0000000fU & (IData)(
                                                       (vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__37__lba 
                                                        >> 0x3cU)))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__boot_off 
            = (vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__37__lba 
               - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__eff_pin_boot);
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mft_off 
            = (vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__37__lba 
               - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pin_mft_lba);
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__hash_slot 
            = (0x0000007fU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mix));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT____VlemCall_1__backing_slot 
            = (0x000000ffU & ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mft_armed) 
                                & (vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__37__lba 
                                   >= vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pin_mft_lba)) 
                               & (0x0000000000000040ULL 
                                  > vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mft_off))
                               ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mft_off)
                               : ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__eff_boot_armed) 
                                    & (vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__37__lba 
                                       >= vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__eff_pin_boot)) 
                                   & (0x0000000000000040ULL 
                                      > vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__boot_off))
                                   ? ((IData)(0x40U) 
                                      + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__boot_off))
                                   : ((IData)(0x80U) 
                                      + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__hash_slot)))));
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_slot 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT____VlemCall_1__backing_slot;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__35__lba 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_lba;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mix 
        = (0x00000fffU & ((((((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__35__lba) 
                              ^ (IData)((__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__35__lba 
                                         >> 0x0cU))) 
                             ^ (IData)((__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__35__lba 
                                        >> 0x18U))) 
                            ^ (IData)((__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__35__lba 
                                       >> 0x24U))) 
                           ^ (IData)((__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__35__lba 
                                      >> 0x30U))) ^ 
                          (0x0000000fU & (IData)((__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__35__lba 
                                                  >> 0x3cU)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__boot_off 
        = (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__35__lba 
           - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__eff_pin_boot);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mft_off 
        = (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__35__lba 
           - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pin_mft_lba);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__hash_slot 
        = (0x0000007fU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mix));
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__35__Vfuncout 
        = (0x000000ffU & ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mft_armed) 
                            & (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__35__lba 
                               >= vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pin_mft_lba)) 
                           & (0x0000000000000040ULL 
                              > vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mft_off))
                           ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mft_off)
                           : ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__eff_boot_armed) 
                                & (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__35__lba 
                                   >= vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__eff_pin_boot)) 
                               & (0x0000000000000040ULL 
                                  > vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__boot_off))
                               ? ((IData)(0x40U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__boot_off))
                               : ((IData)(0x80U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__hash_slot)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_slot 
        = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__35__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__40__word_off 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_word;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__40__lba 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_lba;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__41__lba 
        = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__40__lba;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mix 
        = (0x00000fffU & ((((((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__41__lba) 
                              ^ (IData)((__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__41__lba 
                                         >> 0x0cU))) 
                             ^ (IData)((__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__41__lba 
                                        >> 0x18U))) 
                            ^ (IData)((__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__41__lba 
                                       >> 0x24U))) 
                           ^ (IData)((__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__41__lba 
                                      >> 0x30U))) ^ 
                          (0x0000000fU & (IData)((__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__41__lba 
                                                  >> 0x3cU)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__boot_off 
        = (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__41__lba 
           - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__eff_pin_boot);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mft_off 
        = (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__41__lba 
           - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pin_mft_lba);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__hash_slot 
        = (0x0000007fU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mix));
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__41__Vfuncout 
        = (0x000000ffU & ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mft_armed) 
                            & (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__41__lba 
                               >= vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pin_mft_lba)) 
                           & (0x0000000000000040ULL 
                              > vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mft_off))
                           ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mft_off)
                           : ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__eff_boot_armed) 
                                & (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__41__lba 
                                   >= vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__eff_pin_boot)) 
                               & (0x0000000000000040ULL 
                                  > vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__boot_off))
                               ? ((IData)(0x40U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__boot_off))
                               : ((IData)(0x80U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__hash_slot)))));
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_full_format) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT____VlemCall_2__backing_index 
            = (0x00007fffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx));
    } else {
        vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__38__word_off 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_word;
        vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__38__lba 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_lba;
        vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__39__lba 
            = vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__38__lba;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mix 
            = (0x00000fffU & ((((((IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__39__lba) 
                                  ^ (IData)((vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__39__lba 
                                             >> 0x0cU))) 
                                 ^ (IData)((vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__39__lba 
                                            >> 0x18U))) 
                                ^ (IData)((vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__39__lba 
                                           >> 0x24U))) 
                               ^ (IData)((vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__39__lba 
                                          >> 0x30U))) 
                              ^ (0x0000000fU & (IData)(
                                                       (vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__39__lba 
                                                        >> 0x3cU)))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__boot_off 
            = (vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__39__lba 
               - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__eff_pin_boot);
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mft_off 
            = (vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__39__lba 
               - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pin_mft_lba);
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__hash_slot 
            = (0x0000007fU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mix));
        vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__39__Vfuncout 
            = (0x000000ffU & ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mft_armed) 
                                & (vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__39__lba 
                                   >= vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pin_mft_lba)) 
                               & (0x0000000000000040ULL 
                                  > vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mft_off))
                               ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mft_off)
                               : ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__eff_boot_armed) 
                                    & (vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__39__lba 
                                       >= vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__eff_pin_boot)) 
                                   & (0x0000000000000040ULL 
                                      > vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__boot_off))
                                   ? ((IData)(0x40U) 
                                      + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__boot_off))
                                   : ((IData)(0x80U) 
                                      + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__hash_slot)))));
        vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__38____VlefCall_0__backing_slot 
            = vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__39__Vfuncout;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT____VlemCall_2__backing_index 
            = (((IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__38____VlefCall_0__backing_slot) 
                << 7U) | (IData)(vlSelfRef.__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__38__word_off));
    }
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__40____VlefCall_0__backing_slot 
        = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__41__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__40__Vfuncout 
        = (((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__40____VlefCall_0__backing_slot) 
            << 7U) | (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__40__word_off));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__data_ram_addr 
        = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__40__Vfuncout;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_ram_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT____VlemCall_2__backing_index;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__jittered 
        = ((0x0000000cU < ((IData)(3U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj)))
            ? 0x0000000cU : ((3U > ((IData)(3U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj)))
                              ? 3U : (0x000000ffU & 
                                      ((IData)(3U) 
                                       + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_start_event 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en_prev)) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en) 
              | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_enable_wr)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_stop_event 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en_prev) 
           & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en)) 
              | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_disable_wr)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__lifecycle_quiesce 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__quiesce;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__srst 
        = (1U & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__dma_enabled)) 
                 | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_enabled 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__dma_enabled;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_enabled 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__dma_enabled;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_slot_dbg 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_slot;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_slot_hit 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid
           [vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_slot] 
           & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_lba 
              == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag
              [vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_slot]));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_addr 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_clear)
            ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_ram_addr)
            : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__data_ram_addr));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__computed_latency 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__jittered;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_reset_event 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_start_event) 
           | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_stop_event));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__quiesce 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__lifecycle_quiesce;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_ready 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_enabled) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en) 
              & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_done)) 
                 & (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_all 
        = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_enabled)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_word_cached 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_slot_hit;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__dbg_info 
        = ((((0x0000ff00U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot_dbg) 
                             << 8U)) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_word) 
                                         << 1U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_write))) 
            << 0x00000010U) | ((0x0000ff00U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_slot_dbg) 
                                               << 8U)) 
                               | ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state) 
                                    << 5U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_slot_hit) 
                                              << 4U)) 
                                  | ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_hit) 
                                       << 3U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_word_cached) 
                                                 << 2U)) 
                                     | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid
                                                 [vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_slot]) 
                                         << 1U) | vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid
                                        [vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot])))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_bank_addr 
        = (0x00001fffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_addr) 
                          >> 2U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_bank 
        = (3U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_addr));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_delivery_ready 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_ready;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_ready = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_tag = 0x10U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__selected = 0xffffffffU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__scan_index = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outstanding_count = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__timeout_found = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__timeout_index = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_found = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_index = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_found = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_index = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i = 0U;
    while (VL_GTS_III(32, 0x00000010U, vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i)) {
        if ((1U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags) 
                   >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i)))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outstanding_count 
                = (0x0000001fU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outstanding_count)));
            if ((((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__timeout_found)) 
                  & (~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending) 
                        >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i)))) 
                 & (0x00fffffeU <= vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age
                    [(0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i)]))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__timeout_found = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__timeout_index 
                    = (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i);
            }
        }
        if ((1U & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_found)) 
                   & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_report_pending) 
                      >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i))))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_found = 1U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_index 
                = (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i);
        }
        if ((1U & (((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_found)) 
                    & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending) 
                       >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i))) 
                   & (~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled) 
                         >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i)))))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_found = 1U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_index 
                = (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i);
        }
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i 
            = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i);
    }
    if ((1U & ((~ (0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_report_pending))) 
               & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_all))))) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i = 0U;
        while (VL_GTS_III(32, 0x00000010U, vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i)) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__scan_index 
                = VL_MODDIVS_III(32, ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_cursor) 
                                      + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i), (IData)(0x00000010U));
            if ((VL_GTS_III(32, 0U, vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__selected) 
                 & (~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags) 
                       >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__scan_index))))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__selected 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__scan_index;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_ready = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_tag 
                    = (0x000000ffU & ((IData)(0x10U) 
                                      + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__scan_index));
            }
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i 
                = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i);
        }
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_dbg_info 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__dbg_info;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_bank_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_bank_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_bank_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_bank_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__en 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_active) 
           & (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_bank)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__we 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_we) 
           & (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_bank)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__en 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_active) 
           & (1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_bank)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__we 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_we) 
           & (1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_bank)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__en 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_active) 
           & (2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_bank)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__we 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_we) 
           & (2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_bank)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__en 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_active) 
           & (3U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_bank)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__we 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_we) 
           & (3U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_bank)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_ready 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_ready;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_tag 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_tag;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_rx_tkeepdw 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tkeepdw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_rx_tvalid 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tvalid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_last 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tlast;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_rx_tuser 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tuser;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_user 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tuser;
    vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_bar = 
        ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__io_enabled) 
         & (0U != (0x0000007fU & ((IData)(vlSymsp->TOP__tb_top__DOT__tlps_in_if.tuser) 
                                  >> 2U))));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_rx_tdata[0U] 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tdata[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_rx_tdata[1U] 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tdata[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_rx_tdata[2U] 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tdata[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_rx_tdata[3U] 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tdata[3U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[0U] 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tdata[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[1U] 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tdata[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[2U] 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tdata[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tlp_data[3U] 
        = vlSymsp->TOP__tb_top__DOT__tlps_in_if.tdata[3U];
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completion_tag 
        = (0x000000ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_rx_tdata[2U] 
                          >> 8U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_status 
        = (7U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_rx_tdata[1U] 
                 >> 0x0000000dU));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_has_data 
        = (0x25U == (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_rx_tdata[0U] 
                     >> 0x00000019U));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_bir = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__bir;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completed_index 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completion_tag) 
           - (IData)(0x00000010U));
    vlSelfRef.__VdfgRegularize_hebeb780c_0_6 = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_has_data) 
                                                & (0U 
                                                   == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_status)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_any 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_rx_tvalid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_rx_tuser) 
              & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_has_data) 
                 | (5U == (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_rx_tdata[0U] 
                           >> 0x00000019U)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_requester_id 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__requester_id;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_tag = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__tag;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_traffic_class 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__traffic_class;
    vlSelfRef.tb_top__DOT__i_bar__DOT__norm_attributes 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_normalizer__DOT__attributes;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_tag_match 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_any) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_active_tag) 
              == (0x000000ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_rx_tdata[2U] 
                                 >> 8U))));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completion_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_tag_match;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_error 
        = ((~ (IData)(vlSelfRef.__VdfgRegularize_hebeb780c_0_6)) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_tag_match));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_match 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_tag_match) 
           & (IData)(vlSelfRef.__VdfgRegularize_hebeb780c_0_6));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completion_error 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_error;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dbg_status 
        = ((((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_any) 
               << 9U) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_tag_match) 
                          << 8U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_match) 
                                    << 7U))) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_status) 
                                                 << 4U) 
                                                | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate))) 
            << 0x00000016U) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_active_tag) 
                                << 0x0000000eU) | (
                                                   ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_tag) 
                                                    << 6U) 
                                                   | (0x0000003fU 
                                                      & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_remaining)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_dbg_status 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dbg_status;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en) 
           & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__empty)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tvalid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_dec 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tvalid) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tlast));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_count_next 
        = (0x000007ffU & (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_count) 
                           + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_inc)) 
                          - (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_dec)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_rd_en 
        = ((IData)(vlSymsp->TOP__tb_top__DOT__tlps_dma_out_if.tready) 
           & (0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_count_next)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_rd_en;
}

void Vtop___024root___ico_comb__TOP__1(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___ico_comb__TOP__1\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__tlps_dma_out_tvalid = vlSymsp->TOP__tb_top__DOT__tlps_dma_out_if.tvalid;
}

void Vtop___024root___ico_comb__TOP__2(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___ico_comb__TOP__2\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_en 
        = vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng.tready;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__rd_en) 
           & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__empty)));
}

void Vtop___024root___ico_comb__TOP__3(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___ico_comb__TOP__3\n"); );
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

void Vtop___024root___ico_comb__TOP__4(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___ico_comb__TOP__4\n"); );
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
void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__tlps_dma_out_if__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__tlps_in_if__1(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__tlps_dma_out_if__0(Vtop_IfAXIS128* vlSelf);
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
        Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__tlps_dma_out_if__0((&vlSymsp->TOP__tb_top__DOT__tlps_dma_out_if));
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
    if (((1ULL & vlSelfRef.__VicoTriggered[1U]) | (2ULL 
                                                   & vlSelfRef.__VicoTriggered[0U]))) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__valid 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en) 
               & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__empty)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tvalid 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__valid;
        vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_dec 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tvalid) 
               & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tlast));
        vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_count_next 
            = (0x000007ffU & (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_count) 
                               + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_inc)) 
                              - (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_dec)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_rd_en 
            = ((IData)(vlSymsp->TOP__tb_top__DOT__tlps_dma_out_if.tready) 
               & (0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_count_next)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_rd_en;
        Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__tlps_dma_out_if__0((&vlSymsp->TOP__tb_top__DOT__tlps_dma_out_if));
        vlSelfRef.tb_top__DOT__tlps_dma_out_tvalid 
            = vlSymsp->TOP__tb_top__DOT__tlps_dma_out_if.tvalid;
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
        Vtop___024root___ico_comb__TOP__4(vlSelf);
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
                                                    (((((vlSelfRef.__VdlySched.awaitingCurrentTime() 
                                                         << 5U) 
                                                        | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__clk) 
                                                            & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__clk__0))) 
                                                           << 4U)) 
                                                       | (((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__clk) 
                                                             & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__clk__0))) 
                                                            << 3U) 
                                                           | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__clk) 
                                                               & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__clk__0))) 
                                                              << 2U)) 
                                                          | ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__clk) 
                                                               & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__clk__0))) 
                                                              << 1U) 
                                                             | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__clk) 
                                                                & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__clk__0)))))) 
                                                      << 0x00000010U) 
                                                     | ((((((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__clk) 
                                                              & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__clk__0))) 
                                                             << 3U) 
                                                            | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__clk) 
                                                                & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__clk__0))) 
                                                               << 2U)) 
                                                           | ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk) 
                                                                & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk__0))) 
                                                               << 1U) 
                                                              | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__clk) 
                                                                 & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__clk__0))))) 
                                                          << 0x0000000cU) 
                                                         | ((((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__clk) 
                                                                & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__clk__0))) 
                                                               << 3U) 
                                                              | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__clk) 
                                                                  & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__clk__0))) 
                                                                 << 2U)) 
                                                             | ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__clk) 
                                                                  & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_bar0__DOT__clk__0))) 
                                                                 << 1U) 
                                                                | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__clk) 
                                                                   & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__clk__0))))) 
                                                            << 8U)) 
                                                        | (((((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__clk) 
                                                                & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__clk__0))) 
                                                               << 3U) 
                                                              | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__clk) 
                                                                  & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__clk__0))) 
                                                                 << 2U)) 
                                                             | ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__clk) 
                                                                  & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__clk__0))) 
                                                                 << 1U) 
                                                                | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__clk) 
                                                                   & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__clk__0))))) 
                                                            << 4U) 
                                                           | (((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__clk) 
                                                                 & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__clk__0))) 
                                                                << 3U) 
                                                               | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__clk) 
                                                                   & (~ (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__clk__0))) 
                                                                  << 2U)) 
                                                              | ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en) 
                                                                   != (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en__1)) 
                                                                  << 1U) 
                                                                 | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_select_ur) 
                                                                    != (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__cpl_select_ur__1)))))))));
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__cpl_select_ur__1 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_select_ur;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en__1 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en;
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
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__clk__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__clk;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__clk__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__clk;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__clk__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__clk;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__clk__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__clk;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__clk__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__clk;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__clk__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__clk;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__clk__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__clk;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__clk__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__clk;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__clk__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__clk;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__clk 
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__clk 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk;
}

void Vtop___024root___eval_act(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_act\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    if ((0x0000000000200000ULL & vlSelfRef.__VactTriggered[0U])) {
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
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__clk 
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
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__clk 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk;
    }
    if ((2ULL & vlSelfRef.__VactTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__valid 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en) 
               & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__empty)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tvalid 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__valid;
        vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_dec 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tvalid) 
               & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tlast));
        vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_count_next 
            = (0x000007ffU & (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_count) 
                               + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_inc)) 
                              - (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_dec)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_rd_en 
            = ((IData)(vlSymsp->TOP__tb_top__DOT__tlps_dma_out_if.tready) 
               & (0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_count_next)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_rd_en;
        Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__tlps_dma_out_if__0((&vlSymsp->TOP__tb_top__DOT__tlps_dma_out_if));
        vlSelfRef.tb_top__DOT__tlps_dma_out_tvalid 
            = vlSymsp->TOP__tb_top__DOT__tlps_dma_out_if.tvalid;
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
        Vtop___024root___ico_comb__TOP__4(vlSelf);
    }
}

void Vtop___024root___nba_sequent__TOP__0(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__0\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    IData/*31:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0 = 0;
    SData/*12:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0 = 0;
    // Body
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0 = 0U;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__en) {
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__we) {
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__din;
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__addr;
            __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0 = 1U;
        }
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__dout 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram
            [vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__addr];
    }
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__dout;
}

void Vtop___024root___nba_sequent__TOP__1(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__1\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    IData/*31:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0 = 0;
    SData/*12:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0 = 0;
    // Body
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0 = 0U;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__en) {
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__we) {
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__din;
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__addr;
            __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0 = 1U;
        }
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__dout 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram
            [vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__addr];
    }
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout1 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__dout;
}

void Vtop___024root___nba_sequent__TOP__2(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__2\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    IData/*31:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0 = 0;
    SData/*12:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0 = 0;
    // Body
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0 = 0U;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__en) {
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__we) {
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__din;
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__addr;
            __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0 = 1U;
        }
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__dout 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram
            [vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__addr];
    }
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout2 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__dout;
}

void Vtop___024root___nba_sequent__TOP__3(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__3\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    IData/*31:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0 = 0;
    SData/*12:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0 = 0;
    // Body
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0 = 0U;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__en) {
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__we) {
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__din;
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__addr;
            __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0 = 1U;
        }
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__dout 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram
            [vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__addr];
    }
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout3 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__dout;
}

void Vtop___024root___nba_sequent__TOP__4(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__4\n"); );
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__issue_dw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__state;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__issue_dw;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_valid;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr_bar0;
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

void Vtop___024root___nba_sequent__TOP__5(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__5\n"); );
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
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_valid;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter;
    __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rst) {
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_counter = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_pending = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_counter = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_target = 2U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_ack = 1U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s0 = 0x343bd444U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__latency_counter = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__target_latency = 6U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_valid = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[0U] = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[1U] = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_ctx[2U] = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__stored_data = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_addr = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_latency = 6U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prev_valid = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__idle_counter = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s1 = 0x477c3009U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s2 = 0xe9782444U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s3 = 0x38c48c4dU;
    } else {
        __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_counter 
            = (0x0000ffffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_counter)));
        if ((0xc000U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_counter))) {
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_valid 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_wr_ack 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_ack;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_next_s0 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s0 
           ^ VL_SHIFTL_III(32,32,32, vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s0, 0x0000000bU));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_ready 
        = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_expired 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending) 
           & (0x00010000U <= vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__timeout_counter));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__busy 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__request_pending;
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
                                                          >> 8U)), (IData)(5U))));
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
        = (0x000000ffU & ((IData)(3U) + (((IData)(0x0000000aU) 
                                          * (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_bucket)) 
                                         >> 4U)));
}

void Vtop___024root___nba_sequent__TOP__6(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__6\n"); );
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

void Vtop___024root___nba_sequent__TOP__7(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__7\n"); );
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
    __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_count;
    __Vdly__tb_top__DOT__i_bar__DOT__ur_request_write_ptr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_write_ptr;
    __VdlySet__tb_top__DOT__i_bar__DOT__ur_request_fifo__v0 = 0U;
    __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_rd = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_rd;
    __Vdly__tb_top__DOT__i_bar__DOT__bar0_spill_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_valid;
    __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_wr = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_wr;
    __VdlySet__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v0 = 0U;
    __VdlySet__tb_top__DOT__i_bar__DOT__bar0_buf_ctx__v1 = 0U;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset) {
        __Vdly__tb_top__DOT__i_bar__DOT__ur_request_read_ptr = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_count = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_count = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__ur_request_write_ptr = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_rd = 0U;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_id_rom_data_r 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_identify_rom
        [(0x000003ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_id_rom_addr) 
                         >> 2U))];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy_d1 
        = ((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset))) 
           && (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_count 
        = ((1U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__device_reset) 
                  | (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__dma_enabled))))
            ? 0U : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_count_next));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_count 
        = __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_count;
    vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_write_ptr 
        = __Vdly__tb_top__DOT__i_bar__DOT__ur_request_write_ptr;
    if (__VdlySet__tb_top__DOT__i_bar__DOT__ur_request_fifo__v0) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__ur_request_fifo[__VdlyDim0__tb_top__DOT__i_bar__DOT__ur_request_fifo__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__ur_request_fifo__v0;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_rd 
        = __Vdly__tb_top__DOT__i_bar__DOT__bar0_buf_rd;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_empty 
        = (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_count));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_id_rom_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_id_rom_data_r;
    vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr_last) 
                                                   | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_first) 
                                                      & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr_ready) 
                                                         & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_posted_write) 
                                                            & (0U 
                                                               != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_enabled_byte_count))))));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlps_in_valid 
        = ((IData)(vlSymsp->TOP__tb_top__DOT__tlps_in_if.__VdfgRegularize_hebeb780c_0_3) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr) 
              & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__norm_request_supported) 
                 | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr_last))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__addr_region 
        = (7U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_addr 
                 >> 0x0000000cU));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_almost_full 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_valid) 
           | (0x0eU <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_count)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__wr_en 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__tlps_in_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adjust 
        = (((1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__addr_region))
             ? 2U : ((2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__addr_region))
                      ? 4U : ((3U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__addr_region))
                               ? 1U : 3U))) & (- (IData)(
                                                         (0U 
                                                          != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__addr_region)))));
}

void Vtop___024root___nba_sequent__TOP__8(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__8\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    CData/*3:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_transition_cnt;
    __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_transition_cnt = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_rdy_pending;
    __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_rdy_pending = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014;
    __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014 = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024;
    __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024 = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028;
    __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028 = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030;
    __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030 = 0;
    // Body
    __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024;
    __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028;
    __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030;
    __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_transition_cnt 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_transition_cnt;
    __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_rdy_pending 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_rdy_pending;
    __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_valid_d1;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx_d1[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx_d1[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx_d1[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_data_d1;
    if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_valid) 
         & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_in_range))) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_valid_d1 = 1U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_valid 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_valid;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx_d1[0U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[0U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx_d1[1U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[1U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx_d1[2U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[2U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[0U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[0U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[1U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[1U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[2U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[2U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_data 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_data;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_data_d1 
            = (((((((((0U == (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset)) 
                      | (4U == (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset))) 
                     | (8U == (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset))) 
                    | (0x0000000cU == (0xfffffffcU 
                                       & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset))) 
                   | (0x00000010U == (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset))) 
                  | (0x00000014U == (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset))) 
                 | (0x0000001cU == (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset))) 
                | (0x00000020U == (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset)))
                ? ((0U == (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset))
                    ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000000
                    : ((4U == (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset))
                        ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000004
                        : ((8U == (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset))
                            ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000008
                            : ((0x0000000cU == (0xfffffffcU 
                                                & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset))
                                ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000000C
                                : ((0x00000010U == 
                                    (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset))
                                    ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000010
                                    : ((0x00000014U 
                                        == (0xfffffffcU 
                                            & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset))
                                        ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014
                                        : ((0x0000001cU 
                                            == (0xfffffffcU 
                                                & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset))
                                            ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000001C
                                            : vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000020)))))))
                : ((0x00000024U == (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset))
                    ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024
                    : ((0x00000028U == (0xfffffffcU 
                                        & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset))
                        ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028
                        : ((0x0000002cU == (0xfffffffcU 
                                            & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset))
                            ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000002C
                            : ((0x00000030U == (0xfffffffcU 
                                                & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset))
                                ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030
                                : ((0x00000034U == 
                                    (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset))
                                    ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000034
                                    : 0U))))));
    } else {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_valid_d1 = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_valid 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_valid;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx_d1[0U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[0U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx_d1[1U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[1U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx_d1[2U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[2U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[0U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[0U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[1U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[1U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[2U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[2U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_data 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_data;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_data_d1 = 0U;
    }
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rst) {
        __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024 = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028 = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030 = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014 = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000000 = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000004 = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000008 = 0x00010400U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000000C = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000010 = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000020 = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000002C = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000034 = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_transition_cnt = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_rdy_pending = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000001C = 0U;
    } else {
        if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_valid) 
             & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_in_range))) {
            if ((0x00000014U != (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_offset))) {
                if ((0x00000020U != (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_offset))) {
                    if ((0x00000024U == (0xfffffffcU 
                                         & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_offset))) {
                        if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024 
                                = ((0xffffff00U & __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024) 
                                   | (0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data));
                        }
                        if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024 
                                = ((0xffff00ffU & __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024) 
                                   | (((0xf0U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024 
                                                 >> 8U)) 
                                       | (0x0fU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data 
                                                   >> 8U))) 
                                      << 8U));
                        }
                        if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024 
                                = ((0xff00ffffU & __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024) 
                                   | (0x00ff0000U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data));
                        }
                        if ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024 
                                = ((0x00ffffffU & __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024) 
                                   | (((0xf0U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024 
                                                 >> 0x18U)) 
                                       | (0x0fU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data 
                                                   >> 0x18U))) 
                                      << 0x00000018U));
                        }
                    }
                    if ((0x00000024U != (0xfffffffcU 
                                         & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_offset))) {
                        if ((0x00000028U == (0xfffffffcU 
                                             & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_offset))) {
                            if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028 
                                    = ((0xffffff00U 
                                        & __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028) 
                                       | (0x000000ffU 
                                          & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028));
                            }
                            if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028 
                                    = ((0xffff00ffU 
                                        & __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028) 
                                       | (((0x0fU & 
                                            (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028 
                                             >> 8U)) 
                                           | (0xf0U 
                                              & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data 
                                                 >> 8U))) 
                                          << 8U));
                            }
                            if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028 
                                    = ((0xff00ffffU 
                                        & __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028) 
                                       | (0x00ff0000U 
                                          & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data));
                            }
                            if ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028 
                                    = ((0x00ffffffU 
                                        & __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028) 
                                       | (0xff000000U 
                                          & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data));
                            }
                        }
                        if ((0x00000028U != (0xfffffffcU 
                                             & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_offset))) {
                            if ((0x0000002cU != (0xfffffffcU 
                                                 & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_offset))) {
                                if ((0x00000030U == 
                                     (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_offset))) {
                                    if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                                        __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030 
                                            = ((0xffffff00U 
                                                & __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030) 
                                               | (0x000000ffU 
                                                  & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030));
                                    }
                                    if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                                        __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030 
                                            = ((0xffff00ffU 
                                                & __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030) 
                                               | (((0x0fU 
                                                    & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030 
                                                       >> 8U)) 
                                                   | (0xf0U 
                                                      & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data 
                                                         >> 8U))) 
                                                  << 8U));
                                    }
                                    if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                                        __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030 
                                            = ((0xff00ffffU 
                                                & __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030) 
                                               | (0x00ff0000U 
                                                  & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data));
                                    }
                                    if ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                                        __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030 
                                            = ((0x00ffffffU 
                                                & __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030) 
                                               | (0xff000000U 
                                                  & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data));
                                    }
                                }
                                if ((0x00000030U != 
                                     (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_offset))) {
                                    if ((0x00000034U 
                                         == (0xfffffffcU 
                                             & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_offset))) {
                                        if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000034 
                                                = (
                                                   (0xffffff00U 
                                                    & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000034) 
                                                   | (0x000000ffU 
                                                      & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data));
                                        }
                                        if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000034 
                                                = (
                                                   (0xffff00ffU 
                                                    & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000034) 
                                                   | (0x0000ff00U 
                                                      & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data));
                                        }
                                        if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000034 
                                                = (
                                                   (0xff00ffffU 
                                                    & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000034) 
                                                   | (0x00ff0000U 
                                                      & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data));
                                        }
                                        if ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000034 
                                                = (
                                                   (0x00ffffffU 
                                                    & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000034) 
                                                   | (0xff000000U 
                                                      & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data));
                                        }
                                    }
                                }
                            }
                            if ((0x0000002cU == (0xfffffffcU 
                                                 & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_offset))) {
                                if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000002C 
                                        = ((0xffffff00U 
                                            & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000002C) 
                                           | (0x000000ffU 
                                              & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data));
                                }
                                if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000002C 
                                        = ((0xffff00ffU 
                                            & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000002C) 
                                           | (0x0000ff00U 
                                              & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data));
                                }
                                if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000002C 
                                        = ((0xff00ffffU 
                                            & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000002C) 
                                           | (0x00ff0000U 
                                              & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data));
                                }
                                if ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000002C 
                                        = ((0x00ffffffU 
                                            & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000002C) 
                                           | (0xff000000U 
                                              & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data));
                                }
                            }
                        }
                    }
                }
                if ((0x00000020U == (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_offset))) {
                    if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000020 
                            = ((0xffffff00U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000020) 
                               | (0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data));
                    }
                    if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000020 
                            = ((0xffff00ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000020) 
                               | (0x0000ff00U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data));
                    }
                    if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000020 
                            = ((0xff00ffffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000020) 
                               | (0x00ff0000U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data));
                    }
                    if ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000020 
                            = ((0x00ffffffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000020) 
                               | (0xff000000U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data));
                    }
                }
            }
            if ((0x00000014U == (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_offset))) {
                if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014 
                        = ((0xffffff00U & __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014) 
                           | ((0x0eU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014) 
                              | (0xf1U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data)));
                }
                if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014 
                        = ((0xffff00ffU & __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014) 
                           | (0x0000ff00U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data));
                }
                if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014 
                        = ((0xff00ffffU & __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014) 
                           | (0x00ff0000U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data));
                }
                if ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014 
                        = ((0x00ffffffU & __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014) 
                           | (0xff000000U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014));
                }
            }
        }
        if ((1U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014 
                   & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__cc_en_prev))))) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_rdy_pending = 1U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_transition_cnt = 0U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000001C 
                = (0xfffffffeU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000001C);
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000001C 
                = (0xfffffff3U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000001C);
        } else if (((~ vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014) 
                    & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__cc_en_prev))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000001C 
                = (0xfffffffeU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000001C);
            __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_rdy_pending = 0U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000001C 
                = (0xfffffff3U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000001C);
        }
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_rdy_pending) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_transition_cnt 
                = (0x0000000fU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_transition_cnt)));
            if ((4U <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_transition_cnt))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000001C 
                    = (1U | vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000001C);
                __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_rdy_pending = 0U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000001C 
                    = (0xfffffff3U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000001C);
            }
        }
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_req_addr_from_ctx 
        = (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[0U]);
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_aqa 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_asq_lo 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_acq_lo 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_aqa 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_aqa;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_asq_lo 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_asq_lo;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_acq_lo 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_acq_lo;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_transition_cnt 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_transition_cnt;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_rdy_pending 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_rdy_pending;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_asq_hi 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000002C;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_acq_hi 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000034;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_csts 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000001C;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__cc_en_prev 
        = ((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rst))) 
           && (1U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_asq_hi 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_asq_hi;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_acq_hi 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_acq_hi;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_csts 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_csts;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_cc 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_cc 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_cc;
}

void Vtop___024root___nba_sequent__TOP__9(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__9\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_valid;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_valid = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__msi_pulse;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__msi_pulse = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__scan_vector;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__scan_vector = 0;
    CData/*1:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__query_state;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__query_state = 0;
    // Body
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__scan_vector 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__scan_vector;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__query_state 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__query_state;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_valid;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__msi_pulse 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__msi_pulse;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__pba_set_valid = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__pba_clear_valid = 0U;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__rst) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__pending = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__vector_select = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_valid = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__msi_pulse = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__scan_vector = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__query_state = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__pba_vector = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_vector = 0U;
    } else if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__quiesce) {
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_valid = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__msi_pulse = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__query_state = 0U;
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__event_valid) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__pending = 1U;
        }
    } else {
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__event_valid) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__pending = 1U;
        }
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_valid) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_valid = 0U;
            if ((1U & (~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__event_valid) 
                          & (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_vector)))))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT____Vlvbound_h25bb7b76__0 = 0U;
                if ((0U >= (1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_vector)))) {
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__pending 
                        = (((~ ((IData)(1U) << (1U 
                                                & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_vector)))) 
                            & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__pending)) 
                           | (1U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT____Vlvbound_h25bb7b76__0) 
                                    << (1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_vector)))));
                }
            }
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__pba_clear_valid = 1U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__pba_vector 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_vector;
        } else if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__msi_pulse) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__msi_pulse = 0U;
            if ((1U & (~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__event_valid) 
                          & (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_vector)))))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT____Vlvbound_h25bb7b76__1 = 0U;
                if ((0U >= (1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_vector)))) {
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__pending 
                        = (((~ ((IData)(1U) << (1U 
                                                & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_vector)))) 
                            & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__pending)) 
                           | (1U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT____Vlvbound_h25bb7b76__1) 
                                    << (1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_vector)))));
                }
            }
        } else if ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__query_state))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__vector_select 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__scan_vector;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__query_state = 1U;
        } else if ((1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__query_state))) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__query_state = 2U;
        } else if ((2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__query_state))) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__scan_vector 
                = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__scan_vector) 
                   & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__scan_vector)));
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__query_state = 0U;
        } else {
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__query_state = 0U;
        }
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__scan_vector 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__scan_vector;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__query_state 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__query_state;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_valid 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__msi_pulse 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__msi_pulse;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_delivery_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_intr_req 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__msi_pulse;
    vlSelfRef.tb_top__DOT__i_bar__DOT__intr_req = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_intr_req;
    vlSelfRef.tb_top__DOT__intr_req = vlSelfRef.tb_top__DOT__i_bar__DOT__intr_req;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_delivery_vector 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_vector;
}

void Vtop___024root___nba_sequent__TOP__10(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__10\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    CData/*0:0*/ tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid;
    tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid = 0;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__addr;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__addr 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_valid;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_be = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_addr;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__i_fifo_141_141_clk1_bar_wr__DOT__rd_en 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_rd_en;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_be;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_val 
        = (0x0000ffffU & vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_wr_hit 
        = (0x00004000U > vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr);
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr;
    tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_wr_hit) 
              & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_bar)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_off 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0 
           - (IData)(0x00001000U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_valid 
        = tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_valid 
        = tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_wr = 
        ((IData)(tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid) 
         & ((0x00001000U <= vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0) 
            & (0x00002000U > vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_cc_wr = 
        ((0x00000014U == (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0)) 
         & (IData)(tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_in_range 
        = (0x00004000U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_addr);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_offset 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_index 
        = VL_SHIFTR_III(32,32,32, vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_off, 2U);
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_cc_enable_wr 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_cc_wr) 
           & vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data);
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_cc_disable_wr 
        = ((~ vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_cc_wr));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_is_cq 
        = (1U & vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_index);
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_qid 
        = (0x0000ffffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_index 
                          >> 1U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_enable_wr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_cc_enable_wr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_disable_wr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_cc_disable_wr;
}

void Vtop___024root___nba_sequent__TOP__11(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__11\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    SData/*10:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr = 0;
    SData/*10:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_ptr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_ptr = 0;
    SData/*10:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count;
    __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count = 0;
    VlWide<5>/*133:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0;
    VL_ZERO_W(134, __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0);
    SData/*9:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0 = 0;
    VlWide<5>/*133:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1;
    VL_ZERO_W(134, __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1);
    SData/*9:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1 = 0;
    // Body
    __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count;
    __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_ptr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_ptr;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0 = 0U;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1 = 0U;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__srst) {
        __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_ptr = 0U;
    } else {
        if ((2U != ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_en) 
                      & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__full))) 
                     << 1U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en) 
                               & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__empty)))))) {
            if ((1U == ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_en) 
                          & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__full))) 
                         << 1U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en) 
                                   & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__empty)))))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr 
                    = (0x000007ffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr)));
            } else if ((3U == ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_en) 
                                 & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__full))) 
                                << 1U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en) 
                                          & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__empty)))))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr 
                    = (0x000007ffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr)));
            }
        }
        if ((2U == ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_en) 
                      & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__full))) 
                     << 1U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en) 
                               & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__empty)))))) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count 
                = (0x000007ffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count)));
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0[0U] 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din[0U];
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0[1U] 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din[1U];
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0[2U] 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din[2U];
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0[3U] 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din[3U];
            __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0[4U] 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din[4U];
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0 
                = (0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_ptr));
            __VdlySet__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0 = 1U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_ptr 
                = (0x000007ffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_ptr)));
        } else {
            if ((1U == ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_en) 
                          & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__full))) 
                         << 1U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en) 
                                   & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__empty)))))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count 
                    = (0x000007ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count) 
                                      - (IData)(1U)));
            }
            if ((1U != ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_en) 
                          & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__full))) 
                         << 1U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en) 
                                   & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__empty)))))) {
                if ((3U == ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_en) 
                              & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__full))) 
                             << 1U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en) 
                                       & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__empty)))))) {
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1[0U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din[0U];
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1[1U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din[1U];
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1[2U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din[2U];
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1[3U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din[3U];
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1[4U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din[4U];
                    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1 
                        = (0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_ptr));
                    __VdlySet__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1 = 1U;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_ptr 
                        = (0x000007ffU & ((IData)(1U) 
                                          + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_ptr)));
                }
            }
        }
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_ptr 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_ptr;
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0][0U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0[0U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0][1U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0[1U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0][2U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0[2U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0][3U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0[3U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0][4U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v0[4U];
    }
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1][0U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1[0U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1][1U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1[1U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1][2U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1[2U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1][3U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1[3U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1][4U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem__v1[4U];
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_data_count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_data_count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__empty 
        = (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__full 
        = (0x0400U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem
        [(0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr))][0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem
        [(0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr))][1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem
        [(0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr))][2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[3U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem
        [(0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr))][3U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[4U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem
        [(0x000003ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr))][4U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_full 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__full;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tkeepdw 
        = (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[4U]);
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tdata[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tdata[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tdata[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tdata[3U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[3U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_first 
        = (1U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[4U] 
                 >> 5U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tlast 
        = (1U & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout[4U] 
                 >> 4U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tready 
        = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_full)));
}

void Vtop___024root___nba_sequent__TOP__12(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__12\n"); );
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

void Vtop___024root___nba_sequent__TOP__13(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__13\n"); );
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

void Vtop___024root___nba_sequent__TOP__14(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__14\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    QData/*63:0*/ __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__43__addr;
    __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__43__addr = 0;
    CData/*7:0*/ __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__43__tag;
    __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__43__tag = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__swap32__46__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__swap32__46__Vfuncout = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__swap32__46__value;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__swap32__46__value = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__swap32__47__value;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__swap32__47__value = 0;
    QData/*63:0*/ __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__48__addr;
    __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__48__addr = 0;
    CData/*3:0*/ __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__48__be;
    __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__48__be = 0;
    QData/*63:0*/ __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__51__addr;
    __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__51__addr = 0;
    CData/*3:0*/ __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__51__be;
    __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__51__be = 0;
    QData/*63:0*/ __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__54__addr;
    __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__54__addr = 0;
    CData/*7:0*/ __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__54__tag;
    __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__54__tag = 0;
    QData/*63:0*/ __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__57__addr;
    __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__57__addr = 0;
    CData/*7:0*/ __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__57__tag;
    __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__57__tag = 0;
    CData/*3:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tlast;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tlast = 0;
    CData/*7:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_active_tag;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_active_tag = 0;
    SData/*9:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_remaining;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_remaining = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_payload;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_payload = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_data_hold;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_data_hold = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending = 0;
    QData/*63:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_addr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_addr = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_data;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_data = 0;
    CData/*3:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_be;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_be = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending = 0;
    QData/*63:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_addr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_addr = 0;
    SData/*9:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_len;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_len = 0;
    // Body
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_payload 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_payload;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_data_hold 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_data_hold;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_addr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_data;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_be;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_addr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_len 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_len;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_remaining 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_remaining;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tlast 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tlast;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_active_tag 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_active_tag;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rst) {
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_done = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_valid = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_done = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_data = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[0U] = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[1U] = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[2U] = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[3U] = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tkeepdw = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tlast = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tuser = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_active_tag = 0x10U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_base_addr = 0ULL;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_remaining = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_index = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_payload = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_data_hold = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_addr = 0ULL;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_data = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_be = 0x0fU;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_addr = 0ULL;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_len = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_valid = 0U;
    } else {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_done = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_valid = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_done = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_valid = 0U;
        if ((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_enabled)))) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending = 0U;
        }
        if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_enabled) 
             & (0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate)))) {
            if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_valid) 
                 & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending)))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending = 1U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_addr 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_addr;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_data 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_data;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_be 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_be;
            }
            if ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_req) 
                  & (0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_len))) 
                 & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending)))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending = 1U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_addr 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_addr;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_len 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_len;
            }
        }
        if ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate))) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate = 0U;
        } else if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate))) {
            if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate))) {
                if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid = 0U;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate = 0U;
                } else {
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_data 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_payload;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_valid = 1U;
                    if ((1U & ((1U >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_remaining)) 
                               | (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_enabled))))) {
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_done = 1U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_remaining = 0U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate = 0U;
                    } else {
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_remaining 
                            = (0x000003ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_remaining) 
                                              - (IData)(1U)));
                        __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__43__tag 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_tag;
                        __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__43__addr 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_next_addr;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_index 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_next_index;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_active_tag 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_tag;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_valid = 1U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid = 1U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tlast = 1U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tuser = 1U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[0U] 
                            = (IData)((0x0000000020000001ULL 
                                       | ((QData)((IData)(
                                                          (0x0fU 
                                                           | ((IData)(__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__43__tag) 
                                                              << 8U)))) 
                                          << 0x00000020U)));
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[1U] 
                            = (IData)(((0x0000000020000001ULL 
                                        | ((QData)((IData)(
                                                           (0x0fU 
                                                            | ((IData)(__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__43__tag) 
                                                               << 8U)))) 
                                           << 0x00000020U)) 
                                       >> 0x00000020U));
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[2U] 
                            = (IData)((((QData)((IData)(
                                                        ((IData)(
                                                                 (__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__43__addr 
                                                                  >> 2U)) 
                                                         << 2U))) 
                                        << 0x00000020U) 
                                       | (QData)((IData)(
                                                         (__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__43__addr 
                                                          >> 0x20U)))));
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[3U] 
                            = (IData)(((((QData)((IData)(
                                                         ((IData)(
                                                                  (__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__43__addr 
                                                                   >> 2U)) 
                                                          << 2U))) 
                                         << 0x00000020U) 
                                        | (QData)((IData)(
                                                          (__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__43__addr 
                                                           >> 0x20U)))) 
                                       >> 0x00000020U));
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tkeepdw = 0x0fU;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate = 4U;
                    }
                }
            } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate))) {
                if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_match) {
                    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__swap32__46__value 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_rx_tdata[3U];
                    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__swap32__46__Vfuncout 
                        = ((((0x0000ff00U & (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__swap32__46__value 
                                             << 8U)) 
                             | (0x000000ffU & (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__swap32__46__value 
                                               >> 8U))) 
                            << 0x00000010U) | ((0x0000ff00U 
                                                & (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__swap32__46__value 
                                                   >> 8U)) 
                                               | (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__swap32__46__value 
                                                  >> 0x18U)));
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_payload 
                        = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__swap32__46__Vfuncout;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate = 6U;
                } else if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_error) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_payload = 0U;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_remaining = 0U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_done = 1U;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate = 0U;
                } else if ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_outcome_valid) 
                             & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_outcome_tag) 
                                == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_active_tag))) 
                            & (0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_outcome_status)))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_payload = 0U;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_remaining = 0U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_done = 1U;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate = 0U;
                }
            } else if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid) 
                        & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tready))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate = 5U;
            }
        } else if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate))) {
            if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_done = 1U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate = 0U;
            } else if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid) 
                        & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tready))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate = 3U;
            }
        } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate))) {
            if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid) 
                 & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tready))) {
                if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tlast) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid = 0U;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate = 3U;
                } else {
                    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__swap32__47__value 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_data_hold;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT____VlemCall_0__swap32 
                        = ((((0x0000ff00U & (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__swap32__47__value 
                                             << 8U)) 
                             | (0x000000ffU & (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__swap32__47__value 
                                               >> 8U))) 
                            << 0x00000010U) | ((0x0000ff00U 
                                                & (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__swap32__47__value 
                                                   >> 8U)) 
                                               | (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__swap32__47__value 
                                                  >> 0x18U)));
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[0U] 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT____VlemCall_0__swap32;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[1U] = 0U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[2U] = 0U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[3U] = 0U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tkeepdw = 1U;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tlast = 1U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tuser = 0U;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid = 1U;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate = 2U;
                }
            }
        } else {
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tlast = 0U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tuser = 0U;
            if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_enabled) 
                 & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending))) {
                __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__48__be 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_be;
                __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__48__addr 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_addr;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_data_hold 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_data;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tuser = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[0U] = 0x60000001U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[1U] 
                    = (IData)((((QData)((IData)((__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__48__addr 
                                                 >> 0x20U))) 
                                << 0x00000020U) | (QData)((IData)(__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__48__be))));
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[2U] 
                    = (IData)(((((QData)((IData)((__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__48__addr 
                                                  >> 0x20U))) 
                                 << 0x00000020U) | (QData)((IData)(__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__48__be))) 
                               >> 0x00000020U));
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[3U] 
                    = ((IData)((__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__48__addr 
                                >> 2U)) << 2U);
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tkeepdw = 0x0fU;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tlast = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate = 1U;
            } else if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_enabled) 
                        & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_valid))) {
                __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__51__be 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_be;
                __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__51__addr 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_addr;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_data_hold 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_data;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tuser = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[0U] = 0x60000001U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[1U] 
                    = (IData)((((QData)((IData)((__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__51__addr 
                                                 >> 0x20U))) 
                                << 0x00000020U) | (QData)((IData)(__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__51__be))));
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[2U] 
                    = (IData)(((((QData)((IData)((__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__51__addr 
                                                  >> 0x20U))) 
                                 << 0x00000020U) | (QData)((IData)(__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__51__be))) 
                               >> 0x00000020U));
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[3U] 
                    = ((IData)((__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mwr1__51__addr 
                                >> 2U)) << 2U);
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tkeepdw = 0x0fU;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tlast = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate = 1U;
            } else if ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_enabled) 
                         & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending)) 
                        & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_ready))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_base_addr 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_addr;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_remaining 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_len;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_index = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_active_tag 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_tag;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_valid = 1U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid = 1U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tlast = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tuser = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tkeepdw = 0x0fU;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate = 4U;
                __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__54__tag 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_tag;
                __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__54__addr 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_addr;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[0U] 
                    = (IData)((0x0000000020000001ULL 
                               | ((QData)((IData)((0x0fU 
                                                   | ((IData)(__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__54__tag) 
                                                      << 8U)))) 
                                  << 0x00000020U)));
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[1U] 
                    = (IData)(((0x0000000020000001ULL 
                                | ((QData)((IData)(
                                                   (0x0fU 
                                                    | ((IData)(__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__54__tag) 
                                                       << 8U)))) 
                                   << 0x00000020U)) 
                               >> 0x00000020U));
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[2U] 
                    = (IData)((((QData)((IData)(((IData)(
                                                         (__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__54__addr 
                                                          >> 2U)) 
                                                 << 2U))) 
                                << 0x00000020U) | (QData)((IData)(
                                                                  (__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__54__addr 
                                                                   >> 0x20U)))));
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[3U] 
                    = (IData)(((((QData)((IData)(((IData)(
                                                          (__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__54__addr 
                                                           >> 2U)) 
                                                  << 2U))) 
                                 << 0x00000020U) | (QData)((IData)(
                                                                   (__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__54__addr 
                                                                    >> 0x20U)))) 
                               >> 0x00000020U));
            } else if (((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_enabled) 
                          & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_req)) 
                         & (0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_len))) 
                        & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_ready))) {
                __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__57__tag 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_tag;
                __Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__57__addr 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_addr;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_base_addr 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_addr;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_remaining 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_len;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_index = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_active_tag 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_tag;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_valid = 1U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid = 1U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tlast = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tuser = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[0U] 
                    = (IData)((0x0000000020000001ULL 
                               | ((QData)((IData)((0x0fU 
                                                   | ((IData)(__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__57__tag) 
                                                      << 8U)))) 
                                  << 0x00000020U)));
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[1U] 
                    = (IData)(((0x0000000020000001ULL 
                                | ((QData)((IData)(
                                                   (0x0fU 
                                                    | ((IData)(__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__57__tag) 
                                                       << 8U)))) 
                                   << 0x00000020U)) 
                               >> 0x00000020U));
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[2U] 
                    = (IData)((((QData)((IData)(((IData)(
                                                         (__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__57__addr 
                                                          >> 2U)) 
                                                 << 2U))) 
                                << 0x00000020U) | (QData)((IData)(
                                                                  (__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__57__addr 
                                                                   >> 0x20U)))));
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[3U] 
                    = (IData)(((((QData)((IData)(((IData)(
                                                          (__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__57__addr 
                                                           >> 2U)) 
                                                  << 2U))) 
                                 << 0x00000020U) | (QData)((IData)(
                                                                   (__Vtask_tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__load_mrd1__57__addr 
                                                                    >> 0x20U)))) 
                               >> 0x00000020U));
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tkeepdw = 0x0fU;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate = 4U;
            } else if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_enabled) 
                        & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_req))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_done = 1U;
            }
        }
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_payload 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_payload;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_data_hold 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_data_hold;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_addr 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_data 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_be 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_addr 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_len 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_len;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_remaining 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_remaining;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tlast 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tlast;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_active_tag 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_active_tag;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_done 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_done;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_done 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_done;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_next_index 
        = (0x000003ffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_index)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tdata[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tdata[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tdata[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tdata[3U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata[3U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tkeepdw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tkeepdw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tuser 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tuser;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tlast 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tlast;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tvalid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_tag_match 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_any) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_active_tag) 
              == (0x000000ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_rx_tdata[2U] 
                                 >> 8U))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_next_addr 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_base_addr 
           + (QData)((IData)(((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_next_index) 
                              << 2U))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tdata[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tdata[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tdata[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din[3U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tdata[3U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din[4U] 
        = (0x0000003fU & ((0x00000020U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tuser) 
                                          << 5U)) | 
                          (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tlast) 
                            << 4U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tkeepdw))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_error 
        = ((~ (IData)(vlSelfRef.__VdfgRegularize_hebeb780c_0_6)) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_tag_match));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_match 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_tag_match) 
           & (IData)(vlSelfRef.__VdfgRegularize_hebeb780c_0_6));
}

void Vtop___024root___nba_sequent__TOP__15(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__15\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_valid;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_valid = 0;
    CData/*1:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_status;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_status = 0;
    CData/*3:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v0 = 0;
    CData/*3:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status__v0 = 0;
    IData/*23:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v1;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v1 = 0;
    CData/*3:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v1;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v1 = 0;
    CData/*3:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v2;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v2 = 0;
    CData/*3:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v3;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v3 = 0;
    CData/*3:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v4;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v4 = 0;
    CData/*1:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status__v1;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status__v1 = 0;
    CData/*3:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status__v1;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status__v1 = 0;
    CData/*3:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v5;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v5 = 0;
    CData/*3:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status__v2;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status__v2 = 0;
    CData/*3:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status__v3;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status__v3 = 0;
    CData/*3:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v6;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v6 = 0;
    // Body
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_valid;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_status 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_status;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__rst) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_cursor = 0U;
    } else if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_valid) 
                & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_ready))) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_cursor 
            = (VL_LTES_III(32, 0x0000000fU, vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__selected)
                ? 0U : (0x0000000fU & ((IData)(1U) 
                                       + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__selected)));
    }
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__rst) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_report_pending = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_valid = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_tag = 0x10U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_status = 3U;
        while (VL_GTS_III(32, 0x00000010U, vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i)) {
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v0 
                = (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i);
            vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age.enqueue(0U, (IData)(__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v0));
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status__v0 
                = (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i);
            vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status.enqueue(0U, (IData)(__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status__v0));
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i 
                = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i);
        }
    } else {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i = 0U;
        while (VL_GTS_III(32, 0x00000010U, vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i)) {
            if ((1U & (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags) 
                        >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i)) 
                       & (~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending) 
                             >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i)))))) {
                __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v1 
                    = (0x00ffffffU & ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age
                                      [(0x0000000fU 
                                        & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i)]));
                __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v1 
                    = (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i);
                vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age.enqueue(__VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v1, (IData)(__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v1));
            }
            if (((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_all) 
                   & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags) 
                      >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i))) 
                  & (~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled) 
                        >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i)))) 
                 & (~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending) 
                       >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i))))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled 
                    = ((IData)(__Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled) 
                       | (0x0000ffffU & ((IData)(1U) 
                                         << (0x0000000fU 
                                             & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i))));
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_report_pending 
                    = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_report_pending) 
                       | (0x0000ffffU & ((IData)(1U) 
                                         << (0x0000000fU 
                                             & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i))));
            }
            if ((1U & (((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags) 
                          & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled)) 
                         & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported)) 
                        & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending)) 
                       >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i)))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags 
                    = ((~ ((IData)(1U) << (0x0000000fU 
                                           & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i))) 
                       & (IData)(__Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags));
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled 
                    = ((~ ((IData)(1U) << (0x0000000fU 
                                           & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i))) 
                       & (IData)(__Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled));
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported 
                    = ((~ ((IData)(1U) << (0x0000000fU 
                                           & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i))) 
                       & (IData)(__Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported));
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending 
                    = ((~ ((IData)(1U) << (0x0000000fU 
                                           & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i))) 
                       & (IData)(__Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending));
                __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v2 
                    = (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i);
                vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age.enqueue(0U, (IData)(__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v2));
            }
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i 
                = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i);
        }
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_valid) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_valid = 0U;
            if ((VL_LTES_III(32, 0U, vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_index) 
                 & VL_GTS_III(32, 0x00000010U, vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_index))) {
                if ((3U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_status))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported 
                        = ((IData)(__Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported) 
                           | (0x0000ffffU & ((IData)(1U) 
                                             << (0x0000000fU 
                                                 & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_index))));
                } else {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags 
                        = ((~ ((IData)(1U) << (0x0000000fU 
                                               & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_index))) 
                           & (IData)(__Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags));
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending 
                        = ((~ ((IData)(1U) << (0x0000000fU 
                                               & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_index))) 
                           & (IData)(__Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending));
                    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v3 
                        = (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_index);
                    vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age.enqueue(0U, (IData)(__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v3));
                }
            }
        } else if (((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_valid)) 
                    & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_found))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_report_pending 
                = ((~ ((IData)(1U) << (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_index))) 
                   & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_report_pending));
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_valid = 1U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_tag 
                = (0x000000ffU & ((IData)(0x10U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_index)));
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_status = 3U;
        } else if (((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_valid)) 
                    & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_found))) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_valid = 1U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_tag 
                = (0x000000ffU & ((IData)(0x10U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_index)));
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_status 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status
                [vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_index];
        }
        if ((((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completion_valid) 
                & (0x10U <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completion_tag))) 
               & VL_GTS_III(32, 0x00000010U, vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completed_index)) 
              & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags) 
                 >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completed_index))) 
             & (~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending) 
                   >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completed_index))))) {
            if ((1U & (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled) 
                        & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported)) 
                       >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completed_index)))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags 
                    = ((~ ((IData)(1U) << (0x0000000fU 
                                           & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completed_index))) 
                       & (IData)(__Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags));
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled 
                    = ((~ ((IData)(1U) << (0x0000000fU 
                                           & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completed_index))) 
                       & (IData)(__Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled));
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported 
                    = ((~ ((IData)(1U) << (0x0000000fU 
                                           & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completed_index))) 
                       & (IData)(__Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported));
                __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v4 
                    = (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completed_index);
                vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age.enqueue(0U, (IData)(__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v4));
            } else {
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending 
                    = ((IData)(__Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending) 
                       | (0x0000ffffU & ((IData)(1U) 
                                         << (0x0000000fU 
                                             & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completed_index))));
                __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status__v1 
                    = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completion_error)
                        ? 1U : 0U);
                __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status__v1 
                    = (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completed_index);
                vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status.enqueue(__VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status__v1, (IData)(__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status__v1));
            }
        } else if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__timeout_found) {
            if ((1U & (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled) 
                        & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported)) 
                       >> (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__timeout_index)))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags 
                    = ((~ ((IData)(1U) << (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__timeout_index))) 
                       & (IData)(__Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags));
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled 
                    = ((~ ((IData)(1U) << (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__timeout_index))) 
                       & (IData)(__Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled));
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported 
                    = ((~ ((IData)(1U) << (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__timeout_index))) 
                       & (IData)(__Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported));
                __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v5 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__timeout_index;
                vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age.enqueue(0U, (IData)(__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v5));
            } else {
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending 
                    = ((IData)(__Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending) 
                       | (0x0000ffffU & ((IData)(1U) 
                                         << (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__timeout_index))));
                __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status__v2 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__timeout_index;
                vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status.enqueue(2U, (IData)(__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status__v2));
            }
        }
        if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_valid) 
             & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_ready))) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags 
                = ((IData)(__Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags) 
                   | (0x0000ffffU & ((IData)(1U) << 
                                     (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__selected))));
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled 
                = ((~ ((IData)(1U) << (0x0000000fU 
                                       & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__selected))) 
                   & (IData)(__Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled));
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported 
                = ((~ ((IData)(1U) << (0x0000000fU 
                                       & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__selected))) 
                   & (IData)(__Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported));
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_report_pending 
                = ((~ ((IData)(1U) << (0x0000000fU 
                                       & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__selected))) 
                   & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_report_pending));
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending 
                = ((~ ((IData)(1U) << (0x0000000fU 
                                       & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__selected))) 
                   & (IData)(__Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending));
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status__v3 
                = (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__selected);
            vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status.enqueue(0U, (IData)(__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status__v3));
            __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v6 
                = (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__selected);
            vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age.enqueue(0U, (IData)(__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age__v6));
        }
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported;
    vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status.commit(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_valid 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_status 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_status;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending;
    vlSelfRef.__VdlyCommitQueuetb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age.commit(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_outcome_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_outcome_status 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_status;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_outcome_tag 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_tag;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_index 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_tag) 
           - (IData)(0x00000010U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_ready = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_tag = 0x10U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__selected = 0xffffffffU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__scan_index = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outstanding_count = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__timeout_found = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__timeout_index = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_found = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_index = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_found = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_index = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i = 0U;
    while (VL_GTS_III(32, 0x00000010U, vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i)) {
        if ((1U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags) 
                   >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i)))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outstanding_count 
                = (0x0000001fU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outstanding_count)));
            if ((((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__timeout_found)) 
                  & (~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending) 
                        >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i)))) 
                 & (0x00fffffeU <= vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age
                    [(0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i)]))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__timeout_found = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__timeout_index 
                    = (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i);
            }
        }
        if ((1U & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_found)) 
                   & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_report_pending) 
                      >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i))))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_found = 1U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_index 
                = (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i);
        }
        if ((1U & (((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_found)) 
                    & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending) 
                       >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i))) 
                   & (~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled) 
                         >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i)))))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_found = 1U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_index 
                = (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i);
        }
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i 
            = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i);
    }
    if ((1U & ((~ (0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_report_pending))) 
               & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_all))))) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i = 0U;
        while (VL_GTS_III(32, 0x00000010U, vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i)) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__scan_index 
                = VL_MODDIVS_III(32, ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_cursor) 
                                      + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i), (IData)(0x00000010U));
            if ((VL_GTS_III(32, 0U, vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__selected) 
                 & (~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags) 
                       >> (0x0000000fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__scan_index))))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__selected 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__scan_index;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_ready = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_tag 
                    = (0x000000ffU & ((IData)(0x10U) 
                                      + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__scan_index));
            }
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i 
                = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i);
        }
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_ready 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_ready;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_tag 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_tag;
}

void Vtop___024root___nba_sequent__TOP__16(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__16\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15__Vfuncout = 0;
    CData/*3:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15__cls;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15__cls = 0;
    IData/*16:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15__lbas;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15__lbas = 0;
    CData/*7:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15__dsm_ranges;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15__dsm_ranges = 0;
    SData/*15:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15__jitter;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15__jitter = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15____VlefCall_0__timing_profile_us;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15____VlefCall_0__timing_profile_us = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__Vfuncout = 0;
    CData/*3:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__cls;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__cls = 0;
    IData/*16:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__lbas;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__lbas = 0;
    CData/*7:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__dsm_ranges;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__dsm_ranges = 0;
    SData/*15:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__jitter;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__jitter = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16____VlefCall_4__timing_min_u32;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16____VlefCall_4__timing_min_u32 = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16____VlefCall_3__timing_min_u32;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16____VlefCall_3__timing_min_u32 = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16____VlefCall_2__timing_min_u32;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16____VlefCall_2__timing_min_u32 = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16____VlefCall_1__timing_min_u32;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16____VlefCall_1__timing_min_u32 = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16____VlefCall_0__timing_min_u32;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16____VlefCall_0__timing_min_u32 = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__17__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__17__Vfuncout = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__17__a;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__17__a = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__18__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__18__Vfuncout = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__18__a;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__18__a = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__19__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__19__Vfuncout = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__19__a;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__19__a = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__20__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__20__Vfuncout = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__20__a;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__20__a = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__21__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__21__Vfuncout = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__21__a;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__21__a = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_us_to_cycles__22__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_us_to_cycles__22__Vfuncout = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_us_to_cycles__22__us;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_us_to_cycles__22__us = 0;
    QData/*63:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__23__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__23__Vfuncout = 0;
    IData/*23:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__23__dw_index;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__23__dw_index = 0;
    QData/*63:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__24__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__24__Vfuncout = 0;
    IData/*23:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__24__dw_index;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__24__dw_index = 0;
    QData/*63:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__25__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__25__Vfuncout = 0;
    IData/*23:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__25__dw_index;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__25__dw_index = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__Vfuncout = 0;
    SData/*10:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index = 0;
    QData/*63:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__27__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__27__Vfuncout = 0;
    IData/*23:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__27__dw_index;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__27__dw_index = 0;
    QData/*63:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__28__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__28__Vfuncout = 0;
    IData/*23:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__28__dw_index;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__28__dw_index = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__Vfuncout = 0;
    SData/*12:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__base;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__base = 0;
    SData/*10:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data = 0;
    SData/*15:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__30__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__30__Vfuncout = 0;
    SData/*15:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__30__current;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__30__current = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__30__mix;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__30__mix = 0;
    SData/*15:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__31__Vfuncout;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__31__Vfuncout = 0;
    SData/*15:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__31__current;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__31__current = 0;
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__31__mix;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__31__mix = 0;
    IData/*16:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_units_for_nlb__32__nlb_count;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_units_for_nlb__32__nlb_count = 0;
    IData/*16:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_units_for_nlb__33__nlb_count;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_units_for_nlb__33__nlb_count = 0;
    IData/*16:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_units_for_nlb__34__nlb_count;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_units_for_nlb__34__nlb_count = 0;
    CData/*5:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0;
    CData/*5:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_tail;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_tail = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_phase;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_phase = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_created;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_created = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_created;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_created = 0;
    QData/*63:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_base;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_base = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_last;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_last = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_tail;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_tail = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_phase;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_phase = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_vector;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_vector = 0;
    CData/*3:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe_dw_idx;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe_dw_idx = 0;
    CData/*1:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_result;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_result = 0;
    SData/*14:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 0;
    QData/*63:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_count;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_count = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_sqid;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_sqid = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_cid;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_cid = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_status;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_status = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_peloc;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_peloc = 0;
    QData/*63:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_lba;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_lba = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_nsid;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_nsid = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_cmd_active;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_cmd_active = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_lfsr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_lfsr = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_jitter_sample;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_jitter_sample = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_elapsed_cycles;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_elapsed_cycles = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_target_cycles;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_target_cycles = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_wait_cycles;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_wait_cycles = 0;
    SData/*10:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt = 0;
    SData/*12:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_offset;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_offset = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_registered;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_registered = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_cid;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_cid = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_pending;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_pending = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_info;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_info = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__num_queues_granted;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__num_queues_granted = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_arbitration;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_arbitration = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_power_mgmt;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_power_mgmt = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_temp_threshold;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_temp_threshold = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_write_cache;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_write_cache = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_async_event_cfg;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_async_event_cfg = 0;
    CData/*1:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dir;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dir = 0;
    IData/*23:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_data;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_data = 0;
    QData/*63:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_addr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_addr = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_valid;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_valid = 0;
    IData/*19:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_index;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_index = 0;
    QData/*63:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_base;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_base = 0;
    CData/*7:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_idx;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_idx = 0;
    CData/*7:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_last;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_last = 0;
    CData/*1:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw_idx;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw_idx = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw1;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw1 = 0;
    QData/*63:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_read;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_read = 0;
    QData/*63:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_written;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_written = 0;
    QData/*63:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_read_cmds;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_read_cmds = 0;
    QData/*63:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_write_cmds;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_write_cmds = 0;
    QData/*63:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_flush_cmds;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_flush_cmds = 0;
    QData/*63:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_write_zero_cmds;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_write_zero_cmds = 0;
    QData/*63:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_dataset_cmds;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_dataset_cmds = 0;
    QData/*63:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_cmds_completed;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_cmds_completed = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_error_log_entries;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_error_log_entries = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_unsafe_shutdowns;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_unsafe_shutdowns = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__power_on_hours;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__power_on_hours = 0;
    SData/*11:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__hour_second_count;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__hour_second_count = 0;
    IData/*26:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__second_tick;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__second_tick = 0;
    CData/*7:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__thermal_load;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__thermal_load = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__warning_temp_time;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__warning_temp_time = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__critical_temp_time;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__critical_temp_time = 0;
    IData/*31:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v0 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v0;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v0 = 0;
    IData/*31:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v2;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v2 = 0;
    IData/*31:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v3;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v3 = 0;
    IData/*31:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe__v0 = 0;
    CData/*3:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe__v0 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe__v0;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe__v0 = 0;
    IData/*31:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v4;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v4 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v4;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v4 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v5;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v5 = 0;
    IData/*31:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v6;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v6 = 0;
    IData/*31:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v7;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v7 = 0;
    // Body
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_created 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_created;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_created 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_created;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_base 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_base;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_vector 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_vector;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe_dw_idx 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe_dw_idx;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_result 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_result;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_count;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_sqid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_sqid;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_cid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_cid;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_status 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_status;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_peloc 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_peloc;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_lba 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_lba;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_nsid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_nsid;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_cmd_active 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_cmd_active;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_lfsr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_lfsr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_jitter_sample 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_jitter_sample;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_elapsed_cycles 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_elapsed_cycles;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_target_cycles 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_target_cycles;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_offset 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_offset;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_registered 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_registered;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_cid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_cid;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_pending 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_pending;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_info 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_info;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__num_queues_granted 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__num_queues_granted;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_arbitration 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_arbitration;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_power_mgmt 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_power_mgmt;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_temp_threshold 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_temp_threshold;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_write_cache 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_write_cache;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_async_event_cfg 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_async_event_cfg;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_data;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_addr;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_valid;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_index 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_index;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_base 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_base;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_last 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_last;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw_idx 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw_idx;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_read 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_read;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_written 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_written;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_read_cmds 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_read_cmds;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_write_cmds 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_write_cmds;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_flush_cmds 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_flush_cmds;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_write_zero_cmds 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_write_zero_cmds;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_dataset_cmds 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_dataset_cmds;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_cmds_completed 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_cmds_completed;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_error_log_entries 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_error_log_entries;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_unsafe_shutdowns 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_unsafe_shutdowns;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__power_on_hours 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__power_on_hours;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__hour_second_count 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__hour_second_count;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__second_tick 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__second_tick;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__warning_temp_time 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__warning_temp_time;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__critical_temp_time 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__critical_temp_time;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_idx 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_idx;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_phase 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_phase;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_phase 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_phase;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__thermal_load 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__thermal_load;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v0 = 0U;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v4 = 0U;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v5 = 0U;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_last 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_last;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_tail 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_tail;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_tail 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_tail;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw1 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw1;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_wait_cycles 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_wait_cycles;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dir 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dir;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe__v0 = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__msix_trigger = 0U;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__rst) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_hi = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_lo = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_hi = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_lo = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_aqa = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_tail = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_head_db = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_phase = 1U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_created = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_created = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_base = 0ULL;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_base = 0ULL;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_last = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_tail = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_head_db = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_phase = 1U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_vector = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_req = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_addr = 0ULL;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_len = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_req = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_addr = 0ULL;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_data = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_be = 0x0fU;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_valid = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_done = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_valid = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_write = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_flush = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_lba = 0ULL;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_word = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_wdata = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__pba_set_valid = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__pba_set_vector = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe_dw_idx = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_result = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_count = 0ULL;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_sqid = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_cid = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_status = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_peloc = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_lba = 0ULL;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_nsid = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_class = 1U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_effective_class = 1U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_done_for_cmd = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_cmd_active = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_saturated = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_read_seen_hit = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_read_seen_miss = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_lfsr = 0xace1U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_jitter_sample = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_elapsed_cycles = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_target_cycles = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_wait_cycles = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_offset = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_addr = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_registered = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_cid = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_pending = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_info = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__num_queues_granted = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_arbitration = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_power_mgmt = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_temp_threshold = 0x00000157U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_write_cache = 1U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_async_event_cfg = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dir = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_slba = 0ULL;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_total_dw = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_data = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_addr = 0ULL;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_entry_lo = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_entry_hi = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_valid = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_index = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_base = 0ULL;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_idx = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_last = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw_idx = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw0 = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw1 = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw2 = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw3 = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_read = 0x0000000001ee7cfdULL;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_written = 0x0000000000a8e922ULL;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_read_cmds = 0x000000006c2b5758ULL;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_write_cmds = 0x00000000093cbfdcULL;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_flush_cmds = 0ULL;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_write_zero_cmds = 0ULL;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_dataset_cmds = 0ULL;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_cmds_completed = 0x000000000000525dULL;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_error_log_entries = 4U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_unsafe_shutdowns = 0x00000017U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__power_on_hours = 0x000016e1U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__hour_second_count = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__second_tick = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__thermal_load = 8U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__warning_temp_time = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__critical_temp_time = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__power_cycle_count = 0x00000345U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_media_errors = 2U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_irq_coalescing = 0U;
    } else {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_req = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_req = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_valid = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_valid = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_done = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__pba_set_valid = 0U;
        if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_cmd_active) 
             & (0xffffffffU != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_elapsed_cycles))) {
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_elapsed_cycles 
                = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_elapsed_cycles);
        }
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en) {
            if ((0x0773593fU <= vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__second_tick)) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__thermal_load 
                    = ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__thermal_load))
                        ? 0U : (0x000000ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__thermal_load) 
                                               - (IData)(1U))));
                if (((0x0157U <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__smart_temp_k)) 
                     & (0xffffffffU != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__warning_temp_time))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__warning_temp_time 
                        = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__warning_temp_time);
                }
                if (((0x0166U <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__smart_temp_k)) 
                     & (0xffffffffU != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__critical_temp_time))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__critical_temp_time 
                        = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__critical_temp_time);
                }
                if ((0x0e0fU <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__hour_second_count))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__power_on_hours 
                        = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__power_on_hours);
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__hour_second_count = 0U;
                } else {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__hour_second_count 
                        = (0x00000fffU & ((IData)(1U) 
                                          + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__hour_second_count)));
                }
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__second_tick = 0U;
            } else {
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__second_tick 
                    = (0x07ffffffU & ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__second_tick));
            }
        }
        if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en) 
             & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_wr))) {
            if ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_qid))) {
                if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_is_cq) {
                    if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_val) 
                         <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_last))) {
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_head_db 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_val;
                    } else if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_power2_entries) {
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_head_db 
                            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_val) 
                               & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_last));
                    }
                } else if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_val) 
                            <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_last))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_tail 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_val;
                } else if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_power2_entries) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_tail 
                        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_val) 
                           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_last));
                }
            } else if ((1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_qid))) {
                if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_is_cq) {
                    if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_val) 
                         <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last))) {
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_head_db 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_val;
                    } else if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_power2_entries) {
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_head_db 
                            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_val) 
                               & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last));
                    }
                } else if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_val) 
                            <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_last))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_tail 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_val;
                } else if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_power2_entries) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_tail 
                        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_val) 
                           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_last));
                }
            }
        }
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_reset_event) {
            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_start_event) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_hi 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_hi;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_lo 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_lo;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_hi 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_hi;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_lo 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_lo;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_aqa 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aqa;
            } else {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_hi = 0U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_lo = 0U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_hi = 0U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_lo = 0U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_aqa = 0U;
            }
            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_stop_event) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_unsafe_shutdowns 
                    = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_unsafe_shutdowns);
            }
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_tail = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail = 0U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_head_db = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_phase = 1U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_created = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_created = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_base = 0ULL;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_base = 0ULL;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_last = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_tail = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail = 0U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_head_db = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_phase = 1U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_vector = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe_dw_idx = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_result = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_count = 0ULL;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_sqid = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_cid = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_status = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_peloc = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_lba = 0ULL;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_nsid = 0U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_class = 1U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_effective_class = 1U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_done_for_cmd = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_cmd_active = 0U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_saturated = 0U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_read_seen_hit = 0U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_read_seen_miss = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_jitter_sample = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_elapsed_cycles = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_target_cycles = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_wait_cycles = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_offset = 0U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_addr = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_registered = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_cid = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_pending = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_info = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__num_queues_granted = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dir = 0U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_slba = 0ULL;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_total_dw = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_data = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_addr = 0ULL;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_entry_lo = 0U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_entry_hi = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_valid = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_index = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_base = 0ULL;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_idx = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_last = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw_idx = 0U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw0 = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw1 = 0U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw2 = 0U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw3 = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__warning_temp_time = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__critical_temp_time = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_arbitration = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_power_mgmt = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_temp_threshold = 0x00000157U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_write_cache = 1U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_async_event_cfg = 0U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_irq_coalescing = 0U;
        } else if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_enabled) {
            if (((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en)) 
                 & (0x14U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state)))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0U;
            } else if ((0x00000020U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                if ((0x00000010U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0U;
                } else if ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0U;
                } else if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0U;
                } else if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                    if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0U;
                    } else {
                        if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail) 
                             >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_last))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_phase 
                                = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_phase)));
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail = 0U;
                        } else {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail 
                                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail_next;
                        }
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_registered = 0U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_pending = 0U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__pba_set_valid = 1U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__pba_set_vector = 0U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0U;
                    }
                } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_req = 1U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_valid = 1U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_addr 
                        = ((vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_base 
                            + ((QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail)) 
                               << 4U)) + ((QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx)) 
                                          << 2U));
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_be = 0x0fU;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x0aU;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_data 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe
                        [vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx];
                    if ((3U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx))) {
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx = 0U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 0x22U;
                    } else {
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx 
                            = (3U & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx)));
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 0x21U;
                    }
                } else if ((vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_target_cycles 
                            > vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_elapsed_cycles)) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_wait_cycles 
                        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_target_cycles 
                           - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_elapsed_cycles);
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_saturated = 0U;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x1eU;
                } else {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_wait_cycles = 0U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_saturated = 1U;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                }
            } else if ((0x00000010U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                if ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                    if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                        if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                            if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                                __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15__jitter 
                                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_jitter_sample;
                                __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15__dsm_ranges 
                                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_dsm_ranges;
                                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_done_for_cmd = 1U;
                                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_effective_class 
                                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_completion_class;
                                __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15__lbas 
                                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x20U;
                                __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15__cls 
                                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_completion_class;
                                __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__jitter 
                                    = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15__jitter;
                                __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__dsm_ranges 
                                    = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15__dsm_ranges;
                                __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__lbas 
                                    = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15__lbas;
                                __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__cls 
                                    = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15__cls;
                                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__Vstatic__lba32 
                                    = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__lbas;
                                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__Vstatic__range32 
                                    = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__dsm_ranges;
                                if ((8U & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__cls))) {
                                    if ((4U & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__cls))) {
                                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__Vfuncout 
                                            = ((2U 
                                                & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__cls))
                                                ? ((IData)(0x0000000cU) 
                                                   + 
                                                   (3U 
                                                    & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__jitter)))
                                                : (
                                                   (1U 
                                                    & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__cls))
                                                    ? 
                                                   ((IData)(0x0000000cU) 
                                                    + 
                                                    (3U 
                                                     & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__jitter)))
                                                    : 
                                                   ((IData)(0x0000000cU) 
                                                    + 
                                                    (3U 
                                                     & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__jitter)))));
                                    } else if ((2U 
                                                & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__cls))) {
                                        if ((1U & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__cls))) {
                                            __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__Vfuncout 
                                                = ((IData)(0x00001f40U) 
                                                   + 
                                                   (0x00000fffU 
                                                    & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__jitter)));
                                        } else {
                                            __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__17__a 
                                                = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__Vstatic__range32 
                                                   << 3U);
                                            __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__17__Vfuncout 
                                                = (
                                                   (0x000003e8U 
                                                    > __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__17__a)
                                                    ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__17__a
                                                    : 0x000003e8U);
                                            __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16____VlefCall_4__timing_min_u32 
                                                = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__17__Vfuncout;
                                            __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__Vfuncout 
                                                = ((IData)(0x000000faU) 
                                                   + 
                                                   (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16____VlefCall_4__timing_min_u32 
                                                    + 
                                                    (0x0000007fU 
                                                     & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__jitter))));
                                        }
                                    } else if ((1U 
                                                & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__cls))) {
                                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__Vfuncout 
                                            = ((IData)(0x000000dcU) 
                                               + (0x0000003fU 
                                                  & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__jitter)));
                                    } else {
                                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__18__a 
                                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__Vstatic__lba32;
                                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__18__Vfuncout 
                                            = ((0x00000080U 
                                                > __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__18__a)
                                                ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__18__a
                                                : 0x00000080U);
                                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16____VlefCall_3__timing_min_u32 
                                            = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__18__Vfuncout;
                                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__Vfuncout 
                                            = ((IData)(0x000000b4U) 
                                               + (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16____VlefCall_3__timing_min_u32 
                                                  + 
                                                  (0x0000003fU 
                                                   & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__jitter))));
                                    }
                                } else if ((4U & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__cls))) {
                                    if ((2U & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__cls))) {
                                        if ((1U & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__cls))) {
                                            __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__19__a 
                                                = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__Vstatic__lba32 
                                                   + 
                                                   (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__Vstatic__lba32 
                                                    << 1U));
                                            __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__19__Vfuncout 
                                                = (
                                                   (0x00000080U 
                                                    > __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__19__a)
                                                    ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__19__a
                                                    : 0x00000080U);
                                            __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16____VlefCall_2__timing_min_u32 
                                                = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__19__Vfuncout;
                                            __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__Vfuncout 
                                                = ((IData)(0x00000096U) 
                                                   + 
                                                   (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16____VlefCall_2__timing_min_u32 
                                                    + 
                                                    (0x0000003fU 
                                                     & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__jitter))));
                                        } else {
                                            __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__20__a 
                                                = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__Vstatic__lba32 
                                                   << 1U);
                                            __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__20__Vfuncout 
                                                = (
                                                   (0x00000040U 
                                                    > __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__20__a)
                                                    ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__20__a
                                                    : 0x00000040U);
                                            __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16____VlefCall_1__timing_min_u32 
                                                = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__20__Vfuncout;
                                            __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__Vfuncout 
                                                = ((IData)(0x00000073U) 
                                                   + 
                                                   (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16____VlefCall_1__timing_min_u32 
                                                    + 
                                                    (0x0000001fU 
                                                     & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__jitter))));
                                        }
                                    } else if ((1U 
                                                & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__cls))) {
                                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__21__a 
                                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__Vstatic__lba32;
                                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__21__Vfuncout 
                                            = ((0x00000020U 
                                                > __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__21__a)
                                                ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__21__a
                                                : 0x00000020U);
                                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16____VlefCall_0__timing_min_u32 
                                            = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_min_u32__21__Vfuncout;
                                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__Vfuncout 
                                            = ((IData)(0x00000046U) 
                                               + (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16____VlefCall_0__timing_min_u32 
                                                  + 
                                                  (0x0000000fU 
                                                   & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__jitter))));
                                    } else {
                                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__Vfuncout 
                                            = ((IData)(0x0000000cU) 
                                               + (3U 
                                                  & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__jitter)));
                                    }
                                } else {
                                    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__Vfuncout 
                                        = ((2U & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__cls))
                                            ? ((1U 
                                                & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__cls))
                                                ? ((IData)(0x0000001cU) 
                                                   + 
                                                   (7U 
                                                    & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__jitter)))
                                                : ((IData)(0x00000023U) 
                                                   + 
                                                   (7U 
                                                    & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__jitter))))
                                            : ((IData)(0x0000000cU) 
                                               + (3U 
                                                  & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__jitter))));
                                }
                                __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15____VlefCall_0__timing_profile_us 
                                    = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__16__Vfuncout;
                                __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_us_to_cycles__22__us 
                                    = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15____VlefCall_0__timing_profile_us;
                                __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_us_to_cycles__22__Vfuncout 
                                    = ((VL_SHIFTL_III(32,32,32, __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_us_to_cycles__22__us, 7U) 
                                        - VL_SHIFTL_III(32,32,32, __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_us_to_cycles__22__us, 1U)) 
                                       - __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_us_to_cycles__22__us);
                                __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15__Vfuncout 
                                    = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_us_to_cycles__22__Vfuncout;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_target_cycles 
                                    = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_cycles__15__Vfuncout;
                            } else if ((1U >= vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_wait_cycles)) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_wait_cycles = 0U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            } else {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_wait_cycles 
                                    = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_wait_cycles 
                                       - (IData)(1U));
                            }
                        } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_done) {
                                if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_error) {
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 4U;
                                }
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            }
                        } else if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_done) {
                            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_error) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 4U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            } else if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_idx) 
                                        == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_last))) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            } else {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_idx 
                                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_next;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw_idx = 0U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx 
                                    = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_next) 
                                       << 2U);
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 0x1aU;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x16U;
                            }
                        }
                    } else if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                        if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_valid) {
                                if ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw_idx))) {
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw0 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_data;
                                } else if ((1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw_idx))) {
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw1 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_data;
                                } else if ((2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw_idx))) {
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw2 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_data;
                                } else {
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw3 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_data;
                                }
                            }
                            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_done) {
                                if ((3U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw_idx))) {
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw_idx 
                                        = (3U & ((IData)(1U) 
                                                 + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw_idx)));
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx 
                                        = (0x00ffffffU 
                                           & ((IData)(1U) 
                                              + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx));
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 0x1aU;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x16U;
                                } else if ((0U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw1)) {
                                    if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_idx) 
                                         == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_last))) {
                                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                                    } else {
                                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_idx 
                                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_next;
                                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw_idx = 0U;
                                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx 
                                            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_next) 
                                               << 2U);
                                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 0x1aU;
                                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x16U;
                                    }
                                } else if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_ok) {
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_valid = 1U;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_write = 1U;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_flush = 1U;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_lba 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_lba_value;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_word = 0U;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_wdata 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw1;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x1cU;
                                } else {
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 0x0080U;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                                }
                            }
                        } else {
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_req = 1U;
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_addr 
                                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_addr;
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_len = 1U;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x1bU;
                        }
                    } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_req = 1U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_valid = 1U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_addr 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_addr;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_data 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_data;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_be = 0x0fU;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 0x10U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x0aU;
                    } else {
                        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_valid) {
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_entry_hi 
                                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_data;
                        }
                        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_done) {
                            if ((0x00000200U < vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_list_index)) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 2U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            } else if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_list_entry_invalid) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 0x0013U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            } else {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_valid = 1U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_index 
                                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_list_index;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_base 
                                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_list_page_base;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_addr 
                                    = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_list_page_base 
                                       + (QData)((IData)(
                                                         (0x00000fffU 
                                                          & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_after_first))));
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state 
                                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state;
                            }
                        }
                    }
                } else if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                    if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                        if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_valid) {
                                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_entry_lo 
                                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_data;
                            }
                            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_done) {
                                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_req = 1U;
                                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_addr 
                                    = (4ULL + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_list_entry_addr);
                                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_len = 1U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x18U;
                            }
                        } else if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp1_invalid) 
                                    | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp2_invalid))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 0x0013U;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                        } else if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_after_first_page) {
                            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_uses_list) {
                                if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_valid) 
                                     & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_index 
                                        == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_list_index))) {
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_addr 
                                        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_base 
                                           + (QData)((IData)(
                                                             (0x00000fffU 
                                                              & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_after_first))));
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state;
                                } else {
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_req = 1U;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_addr 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_list_entry_addr;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_len = 1U;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x17U;
                                }
                            } else {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_addr 
                                    = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp2 
                                       + (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_after_first)));
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state 
                                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state;
                            }
                        } else {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_addr 
                                = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1 
                                   + (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_byte_off)));
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state 
                                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state;
                        }
                    } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__23__dw_index 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off 
                            = (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__23__dw_index 
                               << 2U);
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_req = 1U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_valid = 1U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__first_span 
                            = ((IData)(0x00001000U) 
                               - (0x00000fffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1)));
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_be = 0x0fU;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x0aU;
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__23__Vfuncout 
                            = ((vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off 
                                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__first_span)
                                ? (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1 
                                   + (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off)))
                                : (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp2 
                                   + (QData)((IData)(
                                                     (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off 
                                                      - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__first_span)))));
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_addr 
                            = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__23__Vfuncout;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_data 
                            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid_selects_ns1)
                                ? ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt))
                                    ? 0x00001002U : 
                                   ((1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt))
                                     ? 0xddcc9ed1U : 
                                    ((2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt))
                                      ? 0xc24308efU
                                      : ((3U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt))
                                          ? 0x25f2bb9dU
                                          : ((4U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt))
                                              ? 0x11ba97ceU
                                              : 0U)))))
                                : 0U);
                        if ((0x03ffU <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 7U;
                        } else {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt 
                                = (0x000007ffU & ((IData)(1U) 
                                                  + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt)));
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 0x15U;
                        }
                    } else {
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_tail = 0U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head = 0U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail = 0U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_head_db = 0U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_phase = 1U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_created = 0U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_created = 0U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_tail = 0U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head = 0U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail = 0U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_head_db = 0U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_phase = 1U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_registered = 0U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_cid = 0U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_pending = 0U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_info = 0U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0U;
                    }
                } else if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                    if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_done) {
                            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_error) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 4U;
                            }
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                        }
                    } else {
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_valid = 1U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_write = 0U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_flush = 1U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_lba = 0ULL;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_word = 0U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x13U;
                    }
                } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_valid = 1U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_write = 1U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_flush = 0U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_lba 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_disk_lba;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_word 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_disk_word;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_wdata = 0U;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x0dU;
                } else if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_last_dw) 
                            | (0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status)))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state 
                        = (((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status)) 
                            & ((1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dir)) 
                               | (2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dir))))
                            ? 0x12U : 7U);
                } else {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx 
                        = (0x00ffffffU & ((IData)(1U) 
                                          + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx));
                    if ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dir))) {
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x0eU;
                    } else if ((2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dir))) {
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x11U;
                    } else {
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 0x0bU;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x16U;
                    }
                }
            } else if ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                    if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                        if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_done) {
                                if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_hit) {
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_read_seen_hit = 1U;
                                } else {
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_read_seen_miss = 1U;
                                }
                                if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_error) {
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 4U;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                                } else {
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_data 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_rdata;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 0x19U;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x16U;
                                }
                            }
                        } else {
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_valid = 1U;
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_write = 0U;
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_flush = 0U;
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_lba 
                                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_disk_lba;
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_word 
                                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_disk_word;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x0fU;
                        }
                    } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_done) {
                            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_error) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 4U;
                            }
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x10U;
                        }
                    } else {
                        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_done) {
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_valid = 1U;
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_write = 1U;
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_flush = 0U;
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_lba 
                                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_disk_lba;
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_word 
                                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_disk_word;
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_wdata 
                                = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_valid)
                                    ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_data
                                    : vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_data);
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x0dU;
                        }
                        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_valid) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_data 
                                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_data;
                        }
                    }
                } else if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                    if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_req = 1U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_addr 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_addr;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_len = 1U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x0cU;
                    } else if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_done) {
                        if ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state))) {
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_done = 1U;
                        }
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state;
                    }
                } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_cmds_completed 
                        = (1ULL + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_cmds_completed);
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_cmd_active = 0U;
                    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_io) {
                        if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail) 
                             >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_phase 
                                = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_phase)));
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail = 0U;
                        } else {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail 
                                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail_next;
                        }
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head_next;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__pba_set_vector 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_vector;
                    } else {
                        if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail) 
                             >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_last))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_phase 
                                = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_phase)));
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail = 0U;
                        } else {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail 
                                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail_next;
                        }
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head_next;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__pba_set_vector = 0U;
                    }
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__pba_set_valid = 1U;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0U;
                } else {
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_req = 1U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_valid = 1U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_addr 
                        = ((vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_base 
                            + ((QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_tail)) 
                               << 4U)) + ((QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx)) 
                                          << 2U));
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_be = 0x0fU;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x0aU;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_data 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe
                        [vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx];
                    if ((3U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx))) {
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx = 0U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 9U;
                    } else {
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx 
                            = (3U & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx)));
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 8U;
                    }
                }
            } else if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                    if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                        if ((0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_error_log_entries 
                                = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_error_log_entries);
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_count 
                                = (QData)((IData)(((IData)(1U) 
                                                   + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_error_log_entries)));
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_sqid 
                                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_cid 
                                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cid;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_status 
                                = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status) 
                                   << 1U);
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_peloc = 0U;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_nsid 
                                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_lba 
                                = ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid))
                                    ? 0ULL : vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_slba);
                        }
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx = 0U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 8U;
                        __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v0 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_result;
                        __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v0 = 1U;
                        __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v2 
                            = (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid) 
                                << 0x00000010U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_sq_head_next));
                        __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v3 
                            = (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status) 
                                << 0x00000011U) | (
                                                   ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_phase) 
                                                    << 0x00000010U) 
                                                   | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cid)));
                    } else {
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__24__dw_index 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off 
                            = (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__24__dw_index 
                               << 2U);
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_req = 1U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_valid = 1U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__first_span 
                            = ((IData)(0x00001000U) 
                               - (0x00000fffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1)));
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_data = 0U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_be = 0x0fU;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x0aU;
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__24__Vfuncout 
                            = ((vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off 
                                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__first_span)
                                ? (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1 
                                   + (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off)))
                                : (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp2 
                                   + (QData)((IData)(
                                                     (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off 
                                                      - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__first_span)))));
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_addr 
                            = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__24__Vfuncout;
                        if ((0x03ffU <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 7U;
                        } else {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt 
                                = (0x000007ffU & ((IData)(1U) 
                                                  + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt)));
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 6U;
                        }
                    }
                } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__25__dw_index 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off 
                        = (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__25__dw_index 
                           << 2U);
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_req = 1U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_valid = 1U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__first_span 
                        = ((IData)(0x00001000U) - (0x00000fffU 
                                                   & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1)));
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_be = 0x0fU;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x0aU;
                    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__25__Vfuncout 
                        = ((vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off 
                            < vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__first_span)
                            ? (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1 
                               + (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off)))
                            : (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp2 
                               + (QData)((IData)((vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off 
                                                  - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__first_span)))));
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_addr 
                        = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__25__Vfuncout;
                    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt;
                    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__Vfuncout 
                        = ((0U == (0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))
                            ? ((0U == (0x000000ffU 
                                       & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index)))
                                ? 0x01010101U : 0U)
                            : ((1U == (0x000000ffU 
                                       & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))
                                ? (((((((((0U == (0x000000ffU 
                                                  & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))) 
                                          | (1U == 
                                             (0x000000ffU 
                                              & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index)))) 
                                         | (2U == (0x000000ffU 
                                                   & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index)))) 
                                        | (3U == (0x000000ffU 
                                                  & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index)))) 
                                       | (4U == (0x000000ffU 
                                                 & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index)))) 
                                      | (5U == (0x000000ffU 
                                                & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index)))) 
                                     | (6U == (0x000000ffU 
                                               & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index)))) 
                                    | (7U == (0x000000ffU 
                                              & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))))
                                    ? ((0U == (0x000000ffU 
                                               & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index)))
                                        ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_count)
                                        : ((1U == (0x000000ffU 
                                                   & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index)))
                                            ? (IData)(
                                                      (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_count 
                                                       >> 0x20U))
                                            : ((2U 
                                                == 
                                                (0x000000ffU 
                                                 & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index)))
                                                ? (
                                                   ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_cid) 
                                                    << 0x00000010U) 
                                                   | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_sqid))
                                                : (
                                                   (3U 
                                                    == 
                                                    (0x000000ffU 
                                                     & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index)))
                                                    ? 
                                                   (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_peloc) 
                                                     << 0x00000010U) 
                                                    | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_status))
                                                    : 
                                                   ((4U 
                                                     == 
                                                     (0x000000ffU 
                                                      & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index)))
                                                     ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_lba)
                                                     : 
                                                    ((5U 
                                                      == 
                                                      (0x000000ffU 
                                                       & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index)))
                                                      ? (IData)(
                                                                (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_lba 
                                                                 >> 0x20U))
                                                      : 
                                                     ((6U 
                                                       == 
                                                       (0x000000ffU 
                                                        & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index)))
                                                       ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_nsid
                                                       : 0U)))))))
                                    : 0U) : ((2U == 
                                              (0x000000ffU 
                                               & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))
                                              ? ((0x00000080U 
                                                  & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                  ? 0U
                                                  : 
                                                 ((0x00000040U 
                                                   & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                   ? 0U
                                                   : 
                                                  ((0x00000020U 
                                                    & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                    ? 
                                                   ((0x00000010U 
                                                     & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                     ? 
                                                    ((8U 
                                                      & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                      ? 0U
                                                      : 
                                                     ((4U 
                                                       & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                       ? 0U
                                                       : 
                                                      ((2U 
                                                        & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                        ? 
                                                       ((1U 
                                                         & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                         ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__smart_temp_k)
                                                         : 
                                                        (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__smart_temp_k) 
                                                          << 0x00000010U) 
                                                         | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__smart_temp_k)))
                                                        : 
                                                       ((1U 
                                                         & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                         ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__critical_temp_time
                                                         : vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__warning_temp_time))))
                                                     : 
                                                    ((8U 
                                                      & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                      ? 
                                                     ((4U 
                                                       & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                       ? 
                                                      ((2U 
                                                        & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                        ? 0U
                                                        : 
                                                       ((1U 
                                                         & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                         ? 0U
                                                         : vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_error_log_entries))
                                                       : 
                                                      ((2U 
                                                        & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                        ? 0U
                                                        : 
                                                       ((1U 
                                                         & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                         ? 0U
                                                         : vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_media_errors)))
                                                      : 
                                                     ((4U 
                                                       & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                       ? 
                                                      ((2U 
                                                        & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                        ? 0U
                                                        : 
                                                       ((1U 
                                                         & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                         ? 0U
                                                         : vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_unsafe_shutdowns))
                                                       : 
                                                      ((2U 
                                                        & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                        ? 0U
                                                        : 
                                                       ((1U 
                                                         & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                         ? 0U
                                                         : vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__power_on_hours)))))
                                                    : 
                                                   ((0x00000010U 
                                                     & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                     ? 
                                                    ((8U 
                                                      & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                      ? 
                                                     ((4U 
                                                       & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                       ? 
                                                      ((2U 
                                                        & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                        ? 0U
                                                        : 
                                                       ((1U 
                                                         & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                         ? 0U
                                                         : vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__power_cycle_count))
                                                       : 
                                                      ((2U 
                                                        & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                        ? 0U
                                                        : 
                                                       ((1U 
                                                         & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                         ? (IData)(
                                                                   (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_cmds_completed 
                                                                    >> 0x20U))
                                                         : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_cmds_completed))))
                                                      : 
                                                     ((4U 
                                                       & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                       ? 
                                                      ((2U 
                                                        & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                        ? 0U
                                                        : 
                                                       ((1U 
                                                         & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                         ? (IData)(
                                                                   (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_write_cmds 
                                                                    >> 0x20U))
                                                         : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_write_cmds)))
                                                       : 
                                                      ((2U 
                                                        & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                        ? 0U
                                                        : 
                                                       ((1U 
                                                         & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                         ? (IData)(
                                                                   (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_read_cmds 
                                                                    >> 0x20U))
                                                         : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_read_cmds)))))
                                                     : 
                                                    ((8U 
                                                      & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                      ? 
                                                     ((4U 
                                                       & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                       ? 
                                                      ((2U 
                                                        & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                        ? 0U
                                                        : 
                                                       ((1U 
                                                         & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                         ? (IData)(
                                                                   (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_written 
                                                                    >> 0x20U))
                                                         : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_written)))
                                                       : 
                                                      ((2U 
                                                        & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                        ? 0U
                                                        : 
                                                       ((1U 
                                                         & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                         ? (IData)(
                                                                   (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_read 
                                                                    >> 0x20U))
                                                         : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_read))))
                                                      : 
                                                     ((4U 
                                                       & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                       ? 0U
                                                       : 
                                                      ((2U 
                                                        & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                        ? 0U
                                                        : 
                                                       ((1U 
                                                         & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index))
                                                         ? 0x0000000aU
                                                         : 
                                                        (0x64000000U 
                                                         | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__smart_temp_k) 
                                                             << 8U) 
                                                            | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__smart_warning_byte)))))))))))
                                              : ((3U 
                                                  == 
                                                  (0x000000ffU 
                                                   & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))
                                                  ? 
                                                 ((0U 
                                                   == 
                                                   (0x000000ffU 
                                                    & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index)))
                                                   ? 1U
                                                   : 
                                                  ((2U 
                                                    == 
                                                    (0x000000ffU 
                                                     & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index)))
                                                    ? 0x51324235U
                                                    : 
                                                   ((3U 
                                                     == 
                                                     (0x000000ffU 
                                                      & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__dw_index)))
                                                     ? 0x37415847U
                                                     : 0U)))
                                                  : 0U))));
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_data 
                        = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__log_page_word__26__Vfuncout;
                    if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt) 
                         >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_log_last_dw))) {
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 7U;
                    } else {
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt 
                            = (0x000007ffU & ((IData)(1U) 
                                              + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt)));
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 5U;
                    }
                } else {
                    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__27__dw_index 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off 
                        = (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__27__dw_index 
                           << 2U);
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_req = 1U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_valid = 1U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__first_span 
                        = ((IData)(0x00001000U) - (0x00000fffU 
                                                   & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1)));
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_be = 0x0fU;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x0aU;
                    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__27__Vfuncout 
                        = ((vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off 
                            < vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__first_span)
                            ? (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1 
                               + (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off)))
                            : (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp2 
                               + (QData)((IData)((vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off 
                                                  - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__first_span)))));
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_addr 
                        = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__27__Vfuncout;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_data 
                        = (((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt)) 
                            & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid_lists_ns1))
                            ? 1U : 0U);
                    if ((0x03ffU <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt))) {
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 7U;
                    } else {
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt 
                            = (0x000007ffU & ((IData)(1U) 
                                              + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt)));
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 4U;
                    }
                }
            } else if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                    if ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt))) {
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_addr 
                            = (0x00001fffU & ((IData)(4U) 
                                              + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_offset)));
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt = 1U;
                    } else {
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__28__dw_index 
                            = (0x000007ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt) 
                                              - (IData)(1U)));
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off 
                            = (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__28__dw_index 
                               << 2U);
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_req = 1U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_valid = 1U;
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__first_span 
                            = ((IData)(0x00001000U) 
                               - (0x00000fffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1)));
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_be = 0x0fU;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x0aU;
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__28__Vfuncout 
                            = ((vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off 
                                < vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__first_span)
                                ? (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1 
                                   + (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off)))
                                : (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp2 
                                   + (QData)((IData)(
                                                     (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off 
                                                      - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__first_span)))));
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_addr 
                            = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__28__Vfuncout;
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_data;
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index 
                            = (0x000007ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt) 
                                              - (IData)(1U)));
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__base 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_offset;
                        __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__Vfuncout 
                            = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data;
                        if ((0U == (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__base))) {
                            __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__Vfuncout 
                                = ((0x00000400U & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                    ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data
                                    : ((0x00000200U 
                                        & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                        ? ((0x00000100U 
                                            & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                            ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data
                                            : ((0x00000080U 
                                                & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data
                                                : (
                                                   (0x00000040U 
                                                    & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                    ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data
                                                    : 
                                                   ((0x00000020U 
                                                     & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                     ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data
                                                     : 
                                                    ((0x00000010U 
                                                      & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                      ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data
                                                      : 
                                                     ((8U 
                                                       & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                       ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data
                                                       : 
                                                      ((4U 
                                                        & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                        ? 0U
                                                        : 
                                                       ((2U 
                                                         & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                         ? 0U
                                                         : 
                                                        ((1U 
                                                          & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                          ? 0U
                                                          : 0x00000320U)))))))))
                                        : ((0x00000100U 
                                            & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                            ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data
                                            : ((0x00000080U 
                                                & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                ? (
                                                   (0x00000040U 
                                                    & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                    ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data
                                                    : 
                                                   ((0x00000020U 
                                                     & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                     ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data
                                                     : 
                                                    ((0x00000010U 
                                                      & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                      ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data
                                                      : 
                                                     ((8U 
                                                       & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                       ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data
                                                       : 
                                                      ((4U 
                                                        & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                        ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data
                                                        : 
                                                       ((2U 
                                                         & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                         ? 
                                                        ((1U 
                                                          & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                          ? 0x00000100U
                                                          : 0x0000000cU)
                                                         : 
                                                        ((1U 
                                                          & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                          ? 1U
                                                          : 0x00004466U)))))))
                                                : (
                                                   (0x00000040U 
                                                    & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                    ? 
                                                   ((0x00000020U 
                                                     & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                     ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data
                                                     : 
                                                    ((0x00000010U 
                                                      & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                      ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data
                                                      : 
                                                     ((8U 
                                                       & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                       ? 
                                                      ((4U 
                                                        & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                        ? 
                                                       ((2U 
                                                         & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                         ? __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data
                                                         : 0U)
                                                        : 0U)
                                                       : 
                                                      ((4U 
                                                        & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                        ? 
                                                       ((2U 
                                                         & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                         ? 
                                                        ((1U 
                                                          & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                          ? 0x000000eeU
                                                          : 0x77a56000U)
                                                         : __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data)
                                                        : 
                                                       ((2U 
                                                         & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                         ? 
                                                        ((1U 
                                                          & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                          ? 0x00000166U
                                                          : 0x01570000U)
                                                         : 
                                                        ((1U 
                                                          & (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                          ? 2U
                                                          : 0x07030002U))))))
                                                    : __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__rom_data)))));
                        } else if ((0x1000U == (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__base))) {
                            __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__Vfuncout 
                                = ((((0U == (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index)) 
                                     || (2U == (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))) 
                                    || (4U == (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index)))
                                    ? 0x773bd2b0U : 
                                   ((((1U == (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index)) 
                                      || (3U == (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))) 
                                     || (5U == (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index)))
                                     ? 0U : ((6U == (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                              ? 0U : 
                                             ((7U == (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                               ? 0U
                                               : ((0x0020U 
                                                   == (IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__dw_index))
                                                   ? 0x00090000U
                                                   : 0U)))));
                        }
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_data 
                            = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_data_word__29__Vfuncout;
                        if ((0x0400U <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 7U;
                        } else {
                            if ((0x03ffU > (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt))) {
                                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_addr 
                                    = (0x00001fffU 
                                       & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_offset) 
                                          + (0x00001ffcU 
                                             & (((IData)(1U) 
                                                 + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt)) 
                                                << 2U))));
                            }
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 3U;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt 
                                = (0x000007ffU & ((IData)(1U) 
                                                  + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt)));
                        }
                    }
                } else {
                    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__30__mix 
                        = (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid) 
                            << 0x00000018U) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode) 
                                                << 0x00000010U) 
                                               | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cid)));
                    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__30__current 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_lfsr;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_result = 0U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_read_seen_hit = 0U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_read_seen_miss = 0U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__Vstatic__x 
                        = (0x0000ffffU & (((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__30__current) 
                                           ^ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__30__mix) 
                                          ^ (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__30__mix 
                                             >> 0x10U)));
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 0U;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_class = 1U;
                    if ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__Vstatic__x))) {
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__Vstatic__x = 0xace1U;
                    }
                    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__30__Vfuncout 
                        = ((0x0000fffeU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__Vstatic__x) 
                                           << 1U)) 
                           | (1U & VL_REDXOR_16((0xb400U 
                                                 & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__Vstatic__x)))));
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_jitter_sample 
                        = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__30__Vfuncout;
                    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_io) {
                        if ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_flush_cmds 
                                = (1ULL + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_flush_cmds);
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_class = 9U;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_valid = 0U;
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_valid = 1U;
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_write = 0U;
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_flush = 1U;
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_lba = 0ULL;
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_word = 0U;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x13U;
                        } else if ((1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_class = 7U;
                            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_mdts_ok) {
                                if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_lba_ok) {
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_write_cmds 
                                        = (1ULL + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_write_cmds);
                                    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_units_for_nlb__32__nlb_count 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__thermal_load 
                                        = ((0xffU == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__thermal_load))
                                            ? 0x000000ffU
                                            : (0x000000ffU 
                                               & ((IData)(1U) 
                                                  + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__thermal_load))));
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_units_for_nlb__Vstatic__units 
                                        = VL_SHIFTR_III(32,32,32, 
                                                        ((IData)(0x000003ffU) 
                                                         + __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_units_for_nlb__32__nlb_count), 0x0000000aU);
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dir = 1U;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_slba 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_slba;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_total_dw 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_total_dw;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx = 0U;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_valid = 0U;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT____VlemCall_0__data_units_for_nlb 
                                        = (QData)((IData)(
                                                          (0x0000ffffU 
                                                           & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_units_for_nlb__Vstatic__units)));
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 0x0bU;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x16U;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_written 
                                        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_written 
                                           + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT____VlemCall_0__data_units_for_nlb);
                                } else {
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 0x0080U;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                                }
                            } else {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 2U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            }
                        } else if ((2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_class = 4U;
                            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_mdts_ok) {
                                if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_lba_ok) {
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_read_cmds 
                                        = (1ULL + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_read_cmds);
                                    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_units_for_nlb__33__nlb_count 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__thermal_load 
                                        = ((0xffU == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__thermal_load))
                                            ? 0x000000ffU
                                            : (0x000000ffU 
                                               & ((IData)(1U) 
                                                  + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__thermal_load))));
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_units_for_nlb__Vstatic__units 
                                        = VL_SHIFTR_III(32,32,32, 
                                                        ((IData)(0x000003ffU) 
                                                         + __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_units_for_nlb__33__nlb_count), 0x0000000aU);
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dir = 0U;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_slba 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_slba;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_total_dw 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_total_dw;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx = 0U;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_valid = 0U;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT____VlemCall_1__data_units_for_nlb 
                                        = (QData)((IData)(
                                                          (0x0000ffffU 
                                                           & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_units_for_nlb__Vstatic__units)));
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x0eU;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_read 
                                        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_read 
                                           + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT____VlemCall_1__data_units_for_nlb);
                                } else {
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 0x0080U;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                                }
                            } else {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 2U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            }
                        } else if ((8U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_class = 8U;
                            if (((((0U != (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw12 
                                           >> 0x10U)) 
                                   | (0U != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw13)) 
                                  | (0U != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw14)) 
                                 | (0U != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw15))) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 2U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            } else if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_mdts_ok) {
                                if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_lba_ok) {
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_write_zero_cmds 
                                        = (1ULL + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_write_zero_cmds);
                                    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_units_for_nlb__34__nlb_count 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_units_for_nlb__Vstatic__units 
                                        = VL_SHIFTR_III(32,32,32, 
                                                        ((IData)(0x000003ffU) 
                                                         + __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_units_for_nlb__34__nlb_count), 0x0000000aU);
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dir = 2U;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_slba 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_slba;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_total_dw 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_total_dw;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx = 0U;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_valid = 0U;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT____VlemCall_2__data_units_for_nlb 
                                        = (QData)((IData)(
                                                          (0x0000ffffU 
                                                           & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_units_for_nlb__Vstatic__units)));
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_valid = 1U;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_write = 1U;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_flush = 1U;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_lba 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_slba;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_word = 0U;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_wdata 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x13U;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_written 
                                        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_written 
                                           + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT____VlemCall_2__data_units_for_nlb);
                                } else {
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 0x0080U;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                                }
                            } else {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 2U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            }
                        } else if ((9U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_dataset_cmds 
                                = (1ULL + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_dataset_cmds);
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_class = 0x0aU;
                            if ((1U != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid)) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 0x000bU;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            } else if (((0U != (0xfffffffbU 
                                                & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw11)) 
                                        | (0U != (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10 
                                                  >> 8U)))) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 2U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            } else if ((4U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw11)) {
                                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_total_dw 
                                    = (0x00ffffffU 
                                       & VL_SHIFTL_III(24,24,32, 
                                                       ((IData)(1U) 
                                                        + 
                                                        (0x000000ffU 
                                                         & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10)), 2U));
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx = 0U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_idx = 0U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_last 
                                    = (0x000000ffU 
                                       & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10);
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw_idx = 0U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_valid = 0U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 0x1aU;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x16U;
                            } else {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            }
                        } else {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 1U;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                        }
                    } else if ((0x00000080U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                        if ((0x00000040U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 1U;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                        } else if ((0x00000020U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 1U;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                        } else if ((0x00000010U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 1U;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                        } else if ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 1U;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                        } else if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 1U;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                        } else if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 1U;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                        } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 1U;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                        } else {
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_class = 0x0bU;
                            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__format_nvm_ok) {
                                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_valid = 1U;
                                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_write = 1U;
                                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_flush = 1U;
                                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_lba = 0ULL;
                                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_word = 0U;
                                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_wdata = 0xffffffffU;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x1dU;
                            } else {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 0x010aU;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            }
                        }
                    } else if ((0x00000040U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 1U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                    } else if ((0x00000020U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 1U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                    } else if ((0x00000010U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 1U;
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                    } else if ((8U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                        if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                            if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 1U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 1U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            } else if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_registered) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 2U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            } else {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_cid 
                                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cid;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_registered = 1U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head 
                                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head_next;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_cmd_active = 0U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0U;
                            }
                        } else if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                            if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 1U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            } else {
                                if (((((((((1U == (0x000000ffU 
                                                   & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10)) 
                                           | (2U == 
                                              (0x000000ffU 
                                               & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))) 
                                          | (4U == 
                                             (0x000000ffU 
                                              & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))) 
                                         | (6U == (0x000000ffU 
                                                   & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))) 
                                        | (7U == (0x000000ffU 
                                                  & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))) 
                                       | (8U == (0x000000ffU 
                                                 & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))) 
                                      | ((9U == (0x000000ffU 
                                                 & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10)) 
                                         || (0x0aU 
                                             == (0x000000ffU 
                                                 & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10)))) 
                                     | (0x0bU == (0x000000ffU 
                                                  & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10)))) {
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_result 
                                        = ((1U == (0x000000ffU 
                                                   & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))
                                            ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_arbitration
                                            : ((2U 
                                                == 
                                                (0x000000ffU 
                                                 & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))
                                                ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_power_mgmt
                                                : (
                                                   (4U 
                                                    == 
                                                    (0x000000ffU 
                                                     & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))
                                                    ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_temp_threshold
                                                    : 
                                                   ((6U 
                                                     == 
                                                     (0x000000ffU 
                                                      & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))
                                                     ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_write_cache
                                                     : 
                                                    ((7U 
                                                      == 
                                                      (0x000000ffU 
                                                       & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))
                                                      ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__num_queues_granted
                                                      : 
                                                     ((8U 
                                                       == 
                                                       (0x000000ffU 
                                                        & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))
                                                       ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_irq_coalescing
                                                       : 
                                                      (((9U 
                                                         == 
                                                         (0x000000ffU 
                                                          & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10)) 
                                                        || (0x0aU 
                                                            == 
                                                            (0x000000ffU 
                                                             & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10)))
                                                        ? 0U
                                                        : vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_async_event_cfg)))))));
                                } else {
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 2U;
                                }
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            }
                        } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                            if ((1U == (0x000000ffU 
                                        & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_arbitration 
                                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw11;
                            } else if ((2U == (0x000000ffU 
                                               & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_power_mgmt 
                                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw11;
                            } else if ((4U == (0x000000ffU 
                                               & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_temp_threshold 
                                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw11;
                            } else if ((6U == (0x000000ffU 
                                               & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_write_cache 
                                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw11;
                            } else if ((7U == (0x000000ffU 
                                               & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__num_queues_granted = 0U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_result = 0U;
                            } else if ((1U & (~ (((8U 
                                                   == 
                                                   (0x000000ffU 
                                                    & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10)) 
                                                  || (9U 
                                                      == 
                                                      (0x000000ffU 
                                                       & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))) 
                                                 || (0x0aU 
                                                     == 
                                                     (0x000000ffU 
                                                      & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10)))))) {
                                if ((0x0bU == (0x000000ffU 
                                               & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))) {
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_async_event_cfg 
                                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw11;
                                } else {
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 2U;
                                }
                            }
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                        } else {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                        }
                    } else if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                        if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                            if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 1U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            } else {
                                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_class = 2U;
                                if (((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_cns)) 
                                     || (0x11U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_cns)))) {
                                    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid_selects_ns1) {
                                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_offset = 0x1000U;
                                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_addr = 0x1000U;
                                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt = 0U;
                                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 3U;
                                    } else {
                                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 0x000bU;
                                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                                    }
                                } else if ((1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_cns))) {
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_offset = 0U;
                                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_addr = 0U;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt = 0U;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 3U;
                                } else if (((2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_cns)) 
                                            || (0x10U 
                                                == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_cns)))) {
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt = 0U;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 4U;
                                } else if ((3U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_cns))) {
                                    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid_selects_ns1) {
                                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt = 0U;
                                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x15U;
                                    } else {
                                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 0x000bU;
                                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                                    }
                                } else {
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 2U;
                                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                                }
                            }
                        } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                            if (((((1U == (0x0000ffffU 
                                           & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10)) 
                                   & (0U != (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10 
                                             >> 0x10U))) 
                                  & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw11) 
                                 & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1_page_valid))) {
                                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_base 
                                    = (0xfffffffffffff000ULL 
                                       & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1);
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last 
                                    = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10 
                                       >> 0x10U);
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail = 0U;
                                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_head_db = 0U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_phase = 1U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_vector 
                                    = ((1U <= (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw11 
                                               >> 0x10U))
                                        ? 0U : (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw11 
                                                >> 0x10U));
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_created = 1U;
                            } else {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 1U;
                            }
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                        } else {
                            if ((1U == (0x0000ffffU 
                                        & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_created = 0U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail = 0U;
                                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_head_db = 0U;
                            }
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                        }
                    } else if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                        if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 1U;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                        } else {
                            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_class = 3U;
                            if (((((0U == (0x000000ffU 
                                           & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10)) 
                                   || (1U == (0x000000ffU 
                                              & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))) 
                                  || (2U == (0x000000ffU 
                                             & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))) 
                                 || (3U == (0x000000ffU 
                                            & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10)))) {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt = 0U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 5U;
                            } else {
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 2U;
                                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                            }
                        }
                    } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode))) {
                        if ((IData)((((((1U == (0x0000ffffU 
                                                & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10)) 
                                        & (0x00010001U 
                                           == (0xffff0001U 
                                               & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw11))) 
                                       & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_created)) 
                                      & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1_page_valid)) 
                                     & (0U != (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10 
                                               >> 0x10U))))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_base 
                                = (0xfffffffffffff000ULL 
                                   & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1);
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_last 
                                = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10 
                                   >> 0x10U);
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_tail = 0U;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head = 0U;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_created = 1U;
                        } else {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = 1U;
                        }
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                    } else {
                        if ((1U == (0x0000ffffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))) {
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_created = 0U;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_tail = 0U;
                            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head = 0U;
                        }
                        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 7U;
                    }
                    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__31__mix 
                        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_slba) 
                           ^ (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid) 
                               << 0x00000010U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cid)));
                    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__31__current 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_lfsr;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__Vstatic__x 
                        = (0x0000ffffU & (((IData)(__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__31__current) 
                                           ^ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__31__mix) 
                                          ^ (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__31__mix 
                                             >> 0x10U)));
                    if ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__Vstatic__x))) {
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__Vstatic__x = 0xace1U;
                    }
                    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__31__Vfuncout 
                        = ((0x0000fffeU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__Vstatic__x) 
                                           << 1U)) 
                           | (1U & VL_REDXOR_16((0xb400U 
                                                 & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__Vstatic__x)))));
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_lfsr 
                        = __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__31__Vfuncout;
                }
            } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))) {
                if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_valid) {
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe__v0 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_data;
                    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe__v0 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe_dw_idx;
                    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe__v0 = 1U;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe_dw_idx 
                        = (0x0000000fU & ((IData)(1U) 
                                          + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe_dw_idx)));
                }
                if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_done) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 2U;
                }
            } else if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_registered) 
                        & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_pending))) {
                __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v4 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_info;
                __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v4 = 1U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x21U;
                __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v5 = 1U;
                __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v6 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head;
                __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v7 
                    = (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_phase) 
                        << 0x00000010U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_cid));
            } else if ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_registered) 
                         & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_pending))) 
                        & (0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__smart_warning_byte)))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_pending = 1U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_info = 2U;
            } else if (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_valid) 
                        & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_ready))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_req = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_valid = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_addr = 0ULL;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_data = 0U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_be = 0x0fU;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0x0aU;
            } else if ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_tail) 
                         != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head)) 
                        & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_queue_config_valid))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid = 0U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_req = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_addr 
                    = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_base 
                       + VL_SHIFTL_QQI(64,64,32, (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head)), 6U));
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_len = 0x0010U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe_dw_idx = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_cmd_active = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_done_for_cmd = 0U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_class = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_read_seen_hit = 0U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_read_seen_miss = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_elapsed_cycles = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 1U;
            } else if ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_created) 
                         & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_created)) 
                        & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_tail) 
                           != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head)))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_req = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_addr 
                    = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_base 
                       + VL_SHIFTL_QQI(64,64,32, (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head)), 6U));
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_len = 0x0010U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe_dw_idx = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_cmd_active = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_done_for_cmd = 0U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_class = 1U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_read_seen_hit = 0U;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_read_seen_miss = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_elapsed_cycles = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 1U;
            }
        } else {
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_cmd_active = 0U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_done_for_cmd = 0U;
        }
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en_prev 
        = ((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__rst))) 
           && (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_created 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_created;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_created 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_created;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_base 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_base;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_vector 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_vector;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe_dw_idx 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe_dw_idx;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_result 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_result;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_count 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_count;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_sqid 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_sqid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_cid 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_cid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_status 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_status;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_peloc 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_peloc;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_lba 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_lba;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_nsid 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_nsid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_cmd_active 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_cmd_active;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_lfsr 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_lfsr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_jitter_sample 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_jitter_sample;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_elapsed_cycles 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_elapsed_cycles;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_target_cycles 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_target_cycles;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_offset 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_offset;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_registered 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_registered;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_cid 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_cid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_pending 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_pending;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_info 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_info;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__num_queues_granted 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__num_queues_granted;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_arbitration 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_arbitration;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_power_mgmt 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_power_mgmt;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_temp_threshold 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_temp_threshold;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_write_cache 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_write_cache;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_async_event_cfg 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_async_event_cfg;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_data 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_addr 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_valid 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_index 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_index;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_base 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_base;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_last 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_last;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw_idx 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw_idx;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_read 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_read;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_written 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_written;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_read_cmds 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_read_cmds;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_write_cmds 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_write_cmds;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_flush_cmds 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_flush_cmds;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_write_zero_cmds 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_write_zero_cmds;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_dataset_cmds 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_dataset_cmds;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_cmds_completed 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_cmds_completed;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_error_log_entries 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_error_log_entries;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_unsafe_shutdowns 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_unsafe_shutdowns;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__power_on_hours 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__power_on_hours;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__hour_second_count 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__hour_second_count;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__second_tick 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__second_tick;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__warning_temp_time 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__warning_temp_time;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__critical_temp_time 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__critical_temp_time;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_idx 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_idx;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_phase 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_phase;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_phase 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_phase;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__thermal_load 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__thermal_load;
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v0) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe[0U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v0;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe[1U] = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe[2U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v2;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe[3U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v3;
    }
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v4) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe[0U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v4;
    }
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v5) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe[1U] = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe[2U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v6;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe[3U] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe__v7;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_last 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_last;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_tail 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_tail;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_tail 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_tail;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw1 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw1;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_wait_cycles 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_wait_cycles;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dir 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dir;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx;
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe__v0) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe__v0;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_base 
        = (((QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_hi)) 
            << 0x00000020U) | (QData)((IData)((0xfffff000U 
                                               & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_lo))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_valid 
        = ((0U == (0x00000fffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_lo)) 
           & ((0U != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_hi) 
              | (0U != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_lo)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_base 
        = (((QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_hi)) 
            << 0x00000020U) | (QData)((IData)((0xfffff000U 
                                               & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_lo))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_valid 
        = ((0U == (0x00000fffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_lo)) 
           & ((0U != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_hi) 
              | (0U != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_lo)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acqs 
        = (0x00000fffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_aqa 
                          >> 0x00000010U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asqs 
        = (0x00000fffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_aqa);
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_event_vector 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__pba_set_vector;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_id_rom_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_next 
        = (0x000000ffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_idx)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_req 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_req;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_len 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_len;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_req 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_req;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_event_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__pba_set_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__smart_temp_k 
        = (0x0000ffffU & ((IData)(0x012fU) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__thermal_load)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_power2_entries 
        = (0U == ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_last) 
                  & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_last))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_power2_entries 
        = (0U == ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last) 
                  & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail_next 
        = (0x0000ffffU & (((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail)) 
                          & (- (IData)(((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail) 
                                        < (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_delivery_done 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_done;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_nlb64 
        = (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw1));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_timing_info 
        = (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_saturated) 
            << 0x0000001fU) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_effective_class) 
                                << 0x0000001bU) | (0x07ffffffU 
                                                   & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_wait_cycles)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head_next 
        = (0x0000ffffU & (((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head)) 
                          & (- (IData)(((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head) 
                                        < (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_last))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_completion_class 
        = ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status))
            ? ((4U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_class))
                ? ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_read_seen_miss)
                    ? 6U : 5U) : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_class))
            : 0x0cU);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_state 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_total_bytes 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_total_dw 
           << 2U);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_active_qid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_io 
        = (0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_disk_word 
        = (0x0000007fU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_disk_lba 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_slba 
           + (QData)((IData)((0x0001ffffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx 
                                             >> 7U)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_byte_off 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx 
           << 2U);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_last_dw 
        = ((0x00ffffffU & ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx)) 
           >= vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_total_dw);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cid 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[0U] 
           >> 0x00000010U);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode 
        = (0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[0U]);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp2 
        = (((QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[9U])) 
            << 0x00000020U) | (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[8U])));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw13 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[13U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw14 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[14U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw15 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[15U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw11 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[11U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1 
        = (((QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[7U])) 
            << 0x00000020U) | (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[6U])));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw12 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[12U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[10U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_lba 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_lba;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_write 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_write;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_flush 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_flush;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_word 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_word;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_wdata 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_wdata;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_last 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acqs;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_aqa_valid 
        = ((1U <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asqs)) 
           & (1U <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acqs)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_last 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asqs;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_req 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_req;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_len 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_len;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_req 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_req;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__event_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_event_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_resp_dbg_state 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_state;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_resp_dbg_qid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_active_qid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_io_progress 
        = (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status) 
            << 0x00000018U) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dir) 
                                << 0x00000016U) | (
                                                   ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_last_dw) 
                                                    << 0x00000015U) 
                                                   | (0x001fffffU 
                                                      & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_opcode 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid_selects_ns1 
        = ((1U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid) 
           | (0xffffffffU == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid_lists_ns1 
        = ((0xffffffffU == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid) 
           | (0U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1_page_valid 
        = ((0ULL != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1) 
           & (0U == (0x00000fffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp1_invalid 
        = ((0U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_total_dw) 
           | ((0ULL == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1) 
              | (0U != (3U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_first_span 
        = ((IData)(0x00001000U) - (0x00000fffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count 
        = (0x0001ffffU & ((IData)(1U) + (0x0000ffffU 
                                         & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw12)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_log_last_dw 
        = ((0x03ffU < (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10 
                       >> 0x00000010U)) ? 0x000003ffU
            : (0x000007ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10 
                              >> 0x00000010U)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_dsm_ranges 
        = (0x000000ffU & ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__format_nvm_ok 
        = (IData)((((((((1U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid) 
                        & (0U == (0xffffffefU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10))) 
                       & (0U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw11)) 
                      & (0U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw12)) 
                     & (0U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw13)) 
                    & (0U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw14)) 
                   & (0U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw15)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_cns 
        = (0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_slba 
        = (((QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw11)) 
            << 0x00000020U) | (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_power2_entries 
        = (0U == ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_last) 
                  & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_last))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail_next 
        = (0x0000ffffU & (((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail)) 
                          & (- (IData)(((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail) 
                                        < (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_last))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_queue_config_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_aqa_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_valid) 
              & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_valid)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_power2_entries 
        = (0U == ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_last) 
                  & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_last))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head_next 
        = (0x0000ffffU & (((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head)) 
                          & (- (IData)(((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head) 
                                        < (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_last))))));
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_io) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_tail 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_base 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_base;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_phase 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_phase;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_last 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_admin_queues 
            = ((vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw12 
                << 0x00000010U) | ((0x0000ff00U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head) 
                                                   << 8U)) 
                                   | (0x000000ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_tail))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_sq_head_next 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head_next;
    } else {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_tail 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_base 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_base;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_phase 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_phase;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_last 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_last;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_admin_queues 
            = ((((0x0000ff00U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head) 
                                 << 8U)) | (0x000000ffU 
                                            & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_tail))) 
                << 0x00000010U) | ((0x0000ff00U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_last) 
                                                   << 8U)) 
                                   | (0x000000ffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_head_db))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_sq_head_next 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head_next;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__smart_warning_byte 
        = (((0U != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_media_errors) 
            << 2U) | ((0x0157U <= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__smart_temp_k)) 
                      << 1U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_resp_dbg_opcode 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_opcode;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_after_first_page 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_byte_off 
           >= vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_first_span);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_uses_list 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_total_bytes 
           > ((IData)(0x00001000U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_first_span));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp2_required 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_total_bytes 
           > vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_first_span);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_after_first 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_byte_off 
           - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_first_span);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_total_dw 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count 
           << 7U);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count64 
        = (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_io_cmd_info 
        = ((2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))
            ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_slba)
            : vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_io_progress);
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_resp_dbg_admin_queues 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_admin_queues;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp2_invalid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp2_required) 
           & ((0ULL == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp2) 
              | (0U != (0x00000fffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp2)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_list_index 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_after_first 
           >> 0x0000000cU);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_mdts_ok 
        = (0x00002000U >= vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_total_dw);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_lba_ok 
        = ((1U == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid) 
           & ((0x00000000773bd2b0ULL > vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_slba) 
              & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count64 
                 <= (0x00000000773bd2b0ULL - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_slba))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_cmd_info 
        = (((0x1fU == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state)) 
            | ((0x20U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state)) 
               | (0x1eU == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))))
            ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_timing_info
            : ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_io)
                ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_io_cmd_info
                : ((((0x0000ff00U & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status) 
                                     << 8U)) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_cns)) 
                    << 0x00000010U) | ((0x0000ff00U 
                                        & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid 
                                           << 8U)) 
                                       | (0x000000ffU 
                                          & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_list_entry_addr 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp2 
           + (QData)((IData)((vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_list_index 
                              << 3U))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_resp_dbg_cmd_info 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_cmd_info;
}
