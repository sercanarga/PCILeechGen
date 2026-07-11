// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vtop.h for the primary calling header

#include "Vtop__pch.h"

void Vtop___024root___nba_sequent__TOP__17(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__17\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    IData/*31:0*/ __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__lba_count_to_dw__42__lba_count;
    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__lba_count_to_dw__42__lba_count = 0;
    CData/*0:0*/ __VdfgRegularize_hebeb780c_0_10;
    __VdfgRegularize_hebeb780c_0_10 = 0;
    CData/*0:0*/ __VdfgRegularize_hebeb780c_0_11;
    __VdfgRegularize_hebeb780c_0_11 = 0;
    CData/*7:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_spc;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_spc = 0;
    QData/*63:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_mft_lcn;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_mft_lcn = 0;
    CData/*2:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_error;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_error = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_clear;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_clear = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_full_format;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_full_format = 0;
    CData/*0:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_write;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_write = 0;
    IData/*31:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_wdata;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_wdata = 0;
    QData/*63:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_lba;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_lba = 0;
    CData/*7:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx = 0;
    SData/*15:0*/ __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_total_dw;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_total_dw = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v0 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v0;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v0 = 0;
    QData/*63:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag__v0;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag__v0 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag__v0;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag__v0 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v1;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v1 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v1;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v1 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag__v1;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag__v1 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v2;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v2 = 0;
    CData/*0:0*/ __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v2;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v2 = 0;
    QData/*63:0*/ __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag__v2;
    __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag__v2 = 0;
    CData/*7:0*/ __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag__v2;
    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag__v2 = 0;
    // Body
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_clear 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_clear;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_total_dw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_total_dw;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_wdata 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_wdata;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_write 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_write;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_full_format 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_full_format;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_lba 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_lba;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v0 = 0U;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v1 = 0U;
    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v2 = 0U;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_error 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_error;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_spc 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_spc;
    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_mft_lcn 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_mft_lcn;
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__rst) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__reset_idx = 0U;
        while (VL_GTS_III(32, 0x00000100U, vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__reset_idx)) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid[(0x000000ffU 
                                                                                & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__reset_idx)] = 0U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag[(0x000000ffU 
                                                                                & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__reset_idx)] = 0ULL;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__reset_idx 
                = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__reset_idx);
        }
    }
    if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__rst) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__read_bank_q = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_done = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_rdata = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_hit = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__busy = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__error = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_error = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_clear = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_full_format = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_write = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_wdata = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_lba = 0ULL;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_word = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_hit = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_base_lba = 0ULL;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_total_dw = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_armed = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_spc = 0U;
        __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_mft_lcn = 0ULL;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mft_armed = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pin_mft_lba = 0ULL;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pin_boot_lba = 0ULL;
    } else {
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_read) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__read_bank_q 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_bank;
        }
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_done = 0U;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_hit = 0U;
        if ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state))) {
            if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state = 0U;
            } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state))) {
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state = 0U;
            } else {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__busy = 1U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state = 1U;
            }
        } else if ((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state))) {
            if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__busy = 1U;
                if (((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_error)) 
                     & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_write))) {
                    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v0 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot;
                    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v0 = 1U;
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag__v0 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_lba;
                    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag__v0 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot;
                }
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state 
                    = ((((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_error)) 
                         & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_write))) 
                        & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_word_cached))
                        ? 4U : 1U);
            } else {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__busy = 1U;
                if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_full_format) {
                    if ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_word))) {
                        __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v1 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_format_slot;
                        __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v1 = 1U;
                        __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag__v1 
                            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_format_slot;
                    }
                } else {
                    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v2 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_slot;
                    __VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v2 = 1U;
                    __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag__v2 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_lba;
                    __VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag__v2 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_slot;
                }
                if (((0x0000ffffU & ((IData)(1U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx))) 
                     >= (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_total_dw))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state = 1U;
                } else {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx 
                        = (0x0000ffffU & ((IData)(1U) 
                                          + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx)));
                }
            }
        } else if ((1U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_done = 1U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_hit 
                = (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_word_cached) 
                    | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_write)) 
                   | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_clear));
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__busy = 0U;
            __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state = 0U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__error 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_error;
            if ((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_error)))) {
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_rdata 
                    = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_clear)
                        ? 0U : ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_write)
                                 ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_wdata
                                 : ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_word_cached)
                                     ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__ram_dout
                                     : 0U)));
            }
        } else {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__busy = 0U;
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__error = 0U;
            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_valid) {
                if (((0ULL == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_lba) 
                     & (0xffffffffU == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata))) {
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT____VlemCall_3__lba_count_to_dw = 0x8000U;
                } else {
                    __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__lba_count_to_dw__42__lba_count 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT____VlemCall_3__lba_count_to_dw 
                        = ((0U == __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__lba_count_to_dw__42__lba_count)
                            ? 0x0080U : ((0x00000100U 
                                          <= __Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__lba_count_to_dw__42__lba_count)
                                          ? 0x8000U
                                          : (0x00007f80U 
                                             & (__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__lba_count_to_dw__42__lba_count 
                                                << 7U))));
                }
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__busy = 1U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_write 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_write;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_wdata 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_lba 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_lba;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_slot;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_word 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_word;
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_hit 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_slot_hit;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_clear 
                    = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_flush) 
                       & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_write));
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_full_format 
                    = ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_flush) 
                         & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_write)) 
                        & (0ULL == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_lba)) 
                       & (0xffffffffU == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata));
                vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_base_lba 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_lba;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_total_dw 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT____VlemCall_3__lba_count_to_dw;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_error = 0U;
                __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state 
                    = (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_flush) 
                        & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_write))
                        ? 2U : ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_flush)
                                 ? 1U : (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_write) 
                                          | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_word_cached))
                                          ? 3U : 1U)));
            }
        }
        if ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_valid) 
              & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_write)) 
             & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_flush)))) {
            if ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_word))) {
                if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_route_now) {
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_armed = 1U;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_spc = 0U;
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_mft_lcn = 0ULL;
                    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pin_boot_lba 
                        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_lba;
                }
            } else if ((vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pin_boot_lba 
                        == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_lba)) {
                if ((1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_word))) {
                    if ((0x00534654U != (0x00ffffffU 
                                         & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata))) {
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_armed = 0U;
                    }
                } else if ((3U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_word))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_spc 
                        = (0x000000ffU & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata 
                                          >> 8U));
                } else if ((0x0cU == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_word))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_mft_lcn 
                        = ((0xffffffff00000000ULL & __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_mft_lcn) 
                           | (IData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata)));
                } else if ((0x0dU == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_word))) {
                    __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_mft_lcn 
                        = ((0x00000000ffffffffULL & __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_mft_lcn) 
                           | ((QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata)) 
                              << 0x00000020U));
                } else if ((0x7fU == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_word))) {
                    if ((0xaa55U == (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata 
                                     >> 0x10U))) {
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pin_mft_lba 
                            = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pin_boot_lba 
                               + (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_mft_lcn 
                                  * (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_spc))));
                        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mft_armed = 1U;
                    }
                }
            }
        }
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_clear 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_clear;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_total_dw 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_total_dw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_wdata 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_wdata;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_write 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_write;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_full_format 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_full_format;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_lba 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_lba;
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v0) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag__v0] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag__v0;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v0] = 1U;
    }
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v1) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag__v1] = 0ULL;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v1] = 0U;
    }
    if (__VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v2) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag__v2] 
            = __VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag__v2;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid[__VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid__v2] = 1U;
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_error 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_error;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_spc 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_spc;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_mft_lcn 
        = __Vdly__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_mft_lcn;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot_dbg 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_word_cached 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_hit) 
           & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid
              [vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot] 
              & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_lba 
                 == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag
                 [vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot])));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__dbg_state 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state;
    __VdfgRegularize_hebeb780c_0_11 = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_error)) 
                                       & (3U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_clear 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_error)) 
           & (2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_format_slot 
        = (0x000000ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx) 
                          >> 7U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_word 
        = (0x0000007fU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_lba 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_base_lba 
           + (QData)((IData)((0x000001ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx) 
                                             >> 7U)))));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_dbg_state 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__dbg_state;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_read 
        = ((IData)(__VdfgRegularize_hebeb780c_0_11) 
           & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_write)) 
              & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_word_cached)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_write 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_write) 
           & (IData)(__VdfgRegularize_hebeb780c_0_11));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_wdata 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_wdata 
           & (- (IData)((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_clear))))));
    __VdfgRegularize_hebeb780c_0_10 = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_clear) 
                                       | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_write));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__din 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_wdata;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__din 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_wdata;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__din 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_wdata;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__din 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_wdata;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_we 
        = __VdfgRegularize_hebeb780c_0_10;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_active 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_read) 
           | (IData)(__VdfgRegularize_hebeb780c_0_10));
}

void Vtop___024root___nba_sequent__TOP__18(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__18\n"); );
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

void Vtop___024root___nba_sequent__TOP__19(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__19\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__busy;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ready 
        = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ready 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ready;
}

void Vtop___024root___nba_comb__TOP__0(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_comb__TOP__0\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__corr_weight 
        = (0xa0U & (- (IData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__is_sequential))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__base_jitter 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_latency) 
           & (- (IData)((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__is_sequential))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_pop 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_empty)) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_accepted));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_adj 
        = (0x000000ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__base_jitter) 
                          + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj 
        = (0x000000ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adjust) 
                          + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_adj)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__jittered 
        = ((0x0000000cU < ((IData)(3U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj)))
            ? 0x0000000cU : ((3U > ((IData)(3U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj)))
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_write 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_full)) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_valid) 
              | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_valid)));
}

void Vtop___024root___nba_sequent__TOP__20(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__20\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_in_range 
        = (0x00004000U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_addr);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_rd_hit) 
              & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar)));
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
}

void Vtop___024root___nba_comb__TOP__2(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_comb__TOP__2\n"); );
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
}

void Vtop___024root___nba_sequent__TOP__21(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__21\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tready 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tready;
}

void Vtop___024root___nba_comb__TOP__6(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_comb__TOP__6\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_wr_en 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_full)) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tvalid));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_en 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_wr_en;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_inc 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_wr_en) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tlast));
}

void Vtop___024root___nba_sequent__TOP__22(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__22\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completion_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_tag_match;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completion_error 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_error;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_done 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_done;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_done 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_done;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_data;
}

void Vtop___024root___nba_comb__TOP__7(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_comb__TOP__7\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
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

void Vtop___024root___nba_sequent__TOP__23(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__23\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_hi 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_asq_hi;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_lo 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_asq_lo;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_hi 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_acq_hi;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_lo 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_acq_lo;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aqa 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_aqa;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en 
        = (1U & vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_cc);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ctx 
        = vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data;
}

void Vtop___024root___nba_sequent__TOP__24(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__24\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_wr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_wr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_qid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_qid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_is_cq 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_is_cq;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_val 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_val;
}

void Vtop___024root___nba_sequent__TOP__25(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__25\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_id_rom_data;
}

void Vtop___024root___nba_sequent__TOP__26(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__26\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_delivery_valid;
}

void Vtop___024root___nba_sequent__TOP__27(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_sequent__TOP__27\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_wdata;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_detect_now 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_write) 
              & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_flush)) 
                 & ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_word)) 
                    & (((0xebU == (0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata)) 
                        | (0xe9U == (0x000000ffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata))) 
                       & (0x4eU == (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata 
                                    >> 0x00000018U)))))));
}

void Vtop___024root___nba_comb__TOP__8(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_comb__TOP__8\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__ram_dout 
        = ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__read_bank_q))
            ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout0
            : ((1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__read_bank_q))
                ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout1
                : ((2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__read_bank_q))
                    ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout2
                    : vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout3)));
}

void Vtop___024root___nba_comb__TOP__9(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_comb__TOP__9\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__push 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_valid) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ready));
}

void Vtop___024root___nba_comb__TOP__10(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_comb__TOP__10\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
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

void Vtop___024root___nba_comb__TOP__11(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_comb__TOP__11\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_ok 
        = ((0U != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw1) 
           & ((0x00000000773bd2b0ULL > vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_lba_value) 
              & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_nlb64 
                 <= (0x00000000773bd2b0ULL - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_lba_value))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_list_entry_invalid 
        = ((0ULL == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_list_entry_value) 
           | (0U != (0x00000fffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_list_entry_value))));
}

void Vtop___024root___nba_comb__TOP__12(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_comb__TOP__12\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_ready 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_enabled) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en) 
              & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_done)) 
                 & (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state)))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_delivery_ready 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_ready;
}

void Vtop___024root___nba_comb__TOP__13(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_comb__TOP__13\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_start_event 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en_prev)) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en) 
              | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_enable_wr)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_stop_event 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en_prev) 
           & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en)) 
              | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_disable_wr)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_reset_event 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_start_event) 
           | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_stop_event));
}

void Vtop___024root___nba_comb__TOP__14(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___nba_comb__TOP__14\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
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
    // Body
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_route_now 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_armed)) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_detect_now));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__eff_boot_armed 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_route_now) 
           | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_armed));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__eff_pin_boot 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_route_now)
            ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_lba
            : vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pin_boot_lba);
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
}

void Vtop___024root___nba_sequent__TOP__4(Vtop___024root* vlSelf);
void Vtop___024root___nba_sequent__TOP__5(Vtop___024root* vlSelf);
void Vtop___024root___nba_sequent__TOP__6(Vtop___024root* vlSelf);
void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_ur__0(Vtop_IfAXIS128* vlSelf);
void Vtop___024root___nba_sequent__TOP__7(Vtop___024root* vlSelf);
void Vtop___024root___nba_sequent__TOP__8(Vtop___024root* vlSelf);
void Vtop___024root___nba_sequent__TOP__9(Vtop___024root* vlSelf);
void Vtop___024root___nba_sequent__TOP__10(Vtop___024root* vlSelf);
void Vtop___024root___nba_sequent__TOP__11(Vtop___024root* vlSelf);
void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__tlps_dma_out_if__0(Vtop_IfAXIS128* vlSelf);
void Vtop___024root___nba_sequent__TOP__12(Vtop___024root* vlSelf);
void Vtop_IfAXIS128___nba_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__0(Vtop_IfAXIS128* vlSelf);
void Vtop___024root___nba_sequent__TOP__14(Vtop___024root* vlSelf);
void Vtop___024root___nba_sequent__TOP__15(Vtop___024root* vlSelf);
void Vtop___024root___nba_sequent__TOP__16(Vtop___024root* vlSelf);
void Vtop_IfAXIS128___nba_comb__TOP__tb_top__DOT__tlps_dma_out_if__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_cpl__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_ur__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__tlps_out_if__0(Vtop_IfAXIS128* vlSelf);
void Vtop___024root___ico_comb__TOP__4(Vtop___024root* vlSelf);
void Vtop_IfAXIS128___nba_comb__TOP__tb_top__DOT__tlps_dma_out_if__1(Vtop_IfAXIS128* vlSelf);

void Vtop___024root___eval_nba(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_nba\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Locals
    IData/*31:0*/ __Vinline__nba_sequent__TOP__0___VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0;
    __Vinline__nba_sequent__TOP__0___VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0 = 0;
    SData/*12:0*/ __Vinline__nba_sequent__TOP__0___VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0;
    __Vinline__nba_sequent__TOP__0___VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0 = 0;
    CData/*0:0*/ __Vinline__nba_sequent__TOP__0___VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0;
    __Vinline__nba_sequent__TOP__0___VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0 = 0;
    IData/*31:0*/ __Vinline__nba_sequent__TOP__1___VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0;
    __Vinline__nba_sequent__TOP__1___VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0 = 0;
    SData/*12:0*/ __Vinline__nba_sequent__TOP__1___VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0;
    __Vinline__nba_sequent__TOP__1___VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0 = 0;
    CData/*0:0*/ __Vinline__nba_sequent__TOP__1___VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0;
    __Vinline__nba_sequent__TOP__1___VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0 = 0;
    IData/*31:0*/ __Vinline__nba_sequent__TOP__2___VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0;
    __Vinline__nba_sequent__TOP__2___VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0 = 0;
    SData/*12:0*/ __Vinline__nba_sequent__TOP__2___VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0;
    __Vinline__nba_sequent__TOP__2___VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0 = 0;
    CData/*0:0*/ __Vinline__nba_sequent__TOP__2___VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0;
    __Vinline__nba_sequent__TOP__2___VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0 = 0;
    IData/*31:0*/ __Vinline__nba_sequent__TOP__3___VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0;
    __Vinline__nba_sequent__TOP__3___VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0 = 0;
    SData/*12:0*/ __Vinline__nba_sequent__TOP__3___VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0;
    __Vinline__nba_sequent__TOP__3___VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0 = 0;
    CData/*0:0*/ __Vinline__nba_sequent__TOP__3___VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0;
    __Vinline__nba_sequent__TOP__3___VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0 = 0;
    // Body
    if ((0x0000000000004000ULL & vlSelfRef.__VnbaTriggered[0U])) {
        __Vinline__nba_sequent__TOP__0___VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0 = 0U;
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__en) {
            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__we) {
                __Vinline__nba_sequent__TOP__0___VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__din;
                __Vinline__nba_sequent__TOP__0___VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__addr;
                __Vinline__nba_sequent__TOP__0___VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0 = 1U;
            }
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__dout 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram
                [vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__addr];
        }
        if (__Vinline__nba_sequent__TOP__0___VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram[__Vinline__nba_sequent__TOP__0___VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0] 
                = __Vinline__nba_sequent__TOP__0___VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram__v0;
        }
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout0 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__dout;
    }
    if ((0x0000000000008000ULL & vlSelfRef.__VnbaTriggered[0U])) {
        __Vinline__nba_sequent__TOP__1___VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0 = 0U;
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__en) {
            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__we) {
                __Vinline__nba_sequent__TOP__1___VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__din;
                __Vinline__nba_sequent__TOP__1___VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__addr;
                __Vinline__nba_sequent__TOP__1___VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0 = 1U;
            }
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__dout 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram
                [vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__addr];
        }
        if (__Vinline__nba_sequent__TOP__1___VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram[__Vinline__nba_sequent__TOP__1___VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0] 
                = __Vinline__nba_sequent__TOP__1___VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram__v0;
        }
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout1 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__dout;
    }
    if ((0x0000000000010000ULL & vlSelfRef.__VnbaTriggered[0U])) {
        __Vinline__nba_sequent__TOP__2___VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0 = 0U;
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__en) {
            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__we) {
                __Vinline__nba_sequent__TOP__2___VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__din;
                __Vinline__nba_sequent__TOP__2___VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__addr;
                __Vinline__nba_sequent__TOP__2___VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0 = 1U;
            }
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__dout 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram
                [vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__addr];
        }
        if (__Vinline__nba_sequent__TOP__2___VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram[__Vinline__nba_sequent__TOP__2___VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0] 
                = __Vinline__nba_sequent__TOP__2___VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram__v0;
        }
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout2 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__dout;
    }
    if ((0x0000000000020000ULL & vlSelfRef.__VnbaTriggered[0U])) {
        __Vinline__nba_sequent__TOP__3___VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0 = 0U;
        if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__en) {
            if (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__we) {
                __Vinline__nba_sequent__TOP__3___VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__din;
                __Vinline__nba_sequent__TOP__3___VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0 
                    = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__addr;
                __Vinline__nba_sequent__TOP__3___VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0 = 1U;
            }
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__dout 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram
                [vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__addr];
        }
        if (__Vinline__nba_sequent__TOP__3___VdlySet__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram[__Vinline__nba_sequent__TOP__3___VdlyDim0__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0] 
                = __Vinline__nba_sequent__TOP__3___VdlyVal__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram__v0;
        }
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout3 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__dout;
    }
    if ((0x0000000000000010ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__4(vlSelf);
    }
    if ((0x0000000000000400ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__5(vlSelf);
    }
    if ((0x0000000000000040ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__6(vlSelf);
        Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_ur__0((&vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_ur));
    }
    if ((4ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__7(vlSelf);
    }
    if ((0x0000000000000200ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__8(vlSelf);
    }
    if ((0x0000000000000800ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__9(vlSelf);
    }
    if ((0x0000000000000080ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__10(vlSelf);
    }
    if ((0x0000000000100000ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__11(vlSelf);
        Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__tlps_dma_out_if__0((&vlSymsp->TOP__tb_top__DOT__tlps_dma_out_if));
    }
    if ((0x0000000000000020ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__12(vlSelf);
        Vtop_IfAXIS128___nba_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__0((&vlSymsp->TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng));
    }
    if ((8ULL & vlSelfRef.__VnbaTriggered[0U])) {
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
    if ((0x0000000000040000ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__14(vlSelf);
    }
    if ((0x0000000000080000ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__15(vlSelf);
    }
    if ((0x0000000000001000ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__16(vlSelf);
    }
    if ((0x0000000000002000ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__17(vlSelf);
    }
    if ((0x0000000000000100ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_sequent__TOP__18(vlSelf);
    }
    if ((0x0000000000000400ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__busy;
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ready 
            = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__rsp_ready 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_ready;
    }
    if ((0x0000000000000404ULL & vlSelfRef.__VnbaTriggered[0U])) {
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
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__corr_weight 
            = (0xa0U & (- (IData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__is_sequential))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__base_jitter 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_latency) 
               & (- (IData)((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__is_sequential))))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_pop 
            = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_empty)) 
               & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_accepted));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_adj 
            = (0x000000ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__base_jitter) 
                              + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj 
            = (0x000000ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adjust) 
                              + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_adj)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__jittered 
            = ((0x0000000cU < ((IData)(3U) + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj)))
                ? 0x0000000cU : ((3U > ((IData)(3U) 
                                        + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj)))
                                  ? 3U : (0x000000ffU 
                                          & ((IData)(3U) 
                                             + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj)))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__computed_latency 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__jittered;
    }
    if ((0x0000000000000204ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_write 
            = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_full)) 
               & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_valid) 
                  | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_valid)));
    }
    if ((0x0000000000000010ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[0U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[1U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[2U] 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[2U];
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_addr;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_in_range 
            = (0x00004000U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_addr);
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_valid 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
               & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_rd_hit) 
                  & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar)));
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
    }
    if ((0x0000000000100002ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__valid 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en) 
               & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__empty)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tvalid 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__valid;
        vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_dec 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tvalid) 
               & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tlast));
        Vtop_IfAXIS128___nba_comb__TOP__tb_top__DOT__tlps_dma_out_if__0((&vlSymsp->TOP__tb_top__DOT__tlps_dma_out_if));
        vlSelfRef.tb_top__DOT__tlps_dma_out_tvalid 
            = vlSymsp->TOP__tb_top__DOT__tlps_dma_out_if.tvalid;
    }
    if ((0x0000000000000064ULL & vlSelfRef.__VnbaTriggered[0U])) {
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
    if ((0x0000000000100000ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tready 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tready;
    }
    if ((0x0000000000140000ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_wr_en 
            = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_full)) 
               & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tvalid));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_en 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_wr_en;
        vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_inc 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_wr_en) 
               & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tlast));
    }
    if ((0x0000000000040000ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completion_valid 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_tag_match;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completion_error 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_error;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_valid 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_valid;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_done 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_done;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_done 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_wr_done;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_valid 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_valid;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_data 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_data;
    }
    if ((0x00000000000c0000ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dbg_status 
            = ((((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_any) 
                   << 9U) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_tag_match) 
                              << 8U) | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_match) 
                                        << 7U))) | 
                 (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_status) 
                   << 4U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate))) 
                << 0x00000016U) | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_active_tag) 
                                    << 0x0000000eU) 
                                   | (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_tag) 
                                       << 6U) | (0x0000003fU 
                                                 & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_remaining)))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_dbg_status 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dbg_status;
    }
    if ((0x0000000000000200ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_hi 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_asq_hi;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_lo 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_asq_lo;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_hi 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_acq_hi;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_lo 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_acq_lo;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aqa 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_aqa;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en 
            = (1U & vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_cc);
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ctx 
            = vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_data 
            = vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data;
    }
    if ((0x0000000000000080ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_wr 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_wr;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_qid 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_qid;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_is_cq 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_is_cq;
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_val 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_val;
    }
    if ((4ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_data 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_id_rom_data;
    }
    if ((0x0000000000000800ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_valid 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_delivery_valid;
    }
    if ((0x0000000000001000ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_disk_req_wdata;
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
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_detect_now 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_valid) 
               & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_write) 
                  & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_flush)) 
                     & ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_word)) 
                        & (((0xebU == (0x000000ffU 
                                       & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata)) 
                            | (0xe9U == (0x000000ffU 
                                         & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata))) 
                           & (0x4eU == (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata 
                                        >> 0x00000018U)))))));
    }
    if ((0x000000000003e000ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__ram_dout 
            = ((0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__read_bank_q))
                ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout0
                : ((1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__read_bank_q))
                    ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout1
                    : ((2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__read_bank_q))
                        ? vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout2
                        : vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout3)));
    }
    if ((0x0000000000000300ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__push 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_valid) 
               & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ready));
    }
    if ((0x0000000000140006ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_count_next 
            = (0x000007ffU & (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_count) 
                               + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_inc)) 
                              - (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_dec)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_rd_en 
            = ((IData)(vlSymsp->TOP__tb_top__DOT__tlps_dma_out_if.tready) 
               & (0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_count_next)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_rd_en;
        Vtop_IfAXIS128___nba_comb__TOP__tb_top__DOT__tlps_dma_out_if__1((&vlSymsp->TOP__tb_top__DOT__tlps_dma_out_if));
    }
    if ((0x0000000000041000ULL & vlSelfRef.__VnbaTriggered[0U])) {
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
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_ok 
            = ((0U != vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw1) 
               & ((0x00000000773bd2b0ULL > vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_lba_value) 
                  & (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_nlb64 
                     <= (0x00000000773bd2b0ULL - vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_lba_value))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_list_entry_invalid 
            = ((0ULL == vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_list_entry_value) 
               | (0U != (0x00000fffU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_list_entry_value))));
    }
    if ((0x0000000000001200ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_ready 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_enabled) 
               & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en) 
                  & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_done)) 
                     & (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state)))));
        vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_delivery_ready 
            = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_ready;
    }
    if ((0x0000000000001280ULL & vlSelfRef.__VnbaTriggered[0U])) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_start_event 
            = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en_prev)) 
               & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en) 
                  | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_enable_wr)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_stop_event 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en_prev) 
               & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en)) 
                  | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_disable_wr)));
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_reset_event 
            = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_start_event) 
               | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_stop_event));
    }
    if ((0x0000000000003000ULL & vlSelfRef.__VnbaTriggered[0U])) {
        Vtop___024root___nba_comb__TOP__14(vlSelf);
    }
}

void Vtop___024root___timing_resume(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___timing_resume\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    if ((0x0000000000200000ULL & vlSelfRef.__VactTriggered[0U])) {
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

void Vtop___024root___eval_triggers_vec__act(Vtop___024root* vlSelf);
#ifdef VL_DEBUG
VL_ATTR_COLD void Vtop___024root___dump_triggers__act(const VlUnpacked<QData/*63:0*/, 1> &triggers, const std::string &tag);
#endif  // VL_DEBUG
bool Vtop___024root___trigger_anySet__act(const VlUnpacked<QData/*63:0*/, 1> &in);
void Vtop___024root___eval_act(Vtop___024root* vlSelf);

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

#ifdef VL_DEBUG
VL_ATTR_COLD void Vtop___024root___dump_triggers__ico(const VlUnpacked<QData/*63:0*/, 2> &triggers, const std::string &tag);
#endif  // VL_DEBUG
bool Vtop___024root___eval_phase__ico(Vtop___024root* vlSelf);

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
