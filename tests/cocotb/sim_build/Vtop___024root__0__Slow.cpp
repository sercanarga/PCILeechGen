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
    const uint64_t __VscopeHash = VL_MURMUR64_HASH(vlSelf->vlNamep);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 15127220858866356494ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__first_span = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 12447639704256422621ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_units_for_nlb__Vstatic__units = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 4830326655735459091ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__Vstatic__x = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 11545543387255737976ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__Vstatic__lba32 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 454360589040559504ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__Vstatic__range32 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 9284965525539080496ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mix = VL_SCOPED_RAND_RESET_I(12, __VscopeHash, 18234568166550879083ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__boot_off = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 7557652827098347188ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mft_off = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 3183189246054579892ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__hash_slot = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 16895896662978922891ull);
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__cpl_select_ur__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_select_ur;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en;
    vlSelfRef.__VactTriggered[0U] = (1ULL | vlSelfRef.__VactTriggered[0U]);
    vlSelfRef.__VactTriggered[0U] = (2ULL | vlSelfRef.__VactTriggered[0U]);
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
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_msix_table__DOT__clk__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__clk;
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
    const uint64_t __VscopeHash = VL_MURMUR64_HASH(vlSelf->vlNamep);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__byte_off = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 15127220858866356494ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_data_addr__Vstatic__first_span = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 12447639704256422621ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_units_for_nlb__Vstatic__units = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 4830326655735459091ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_mix_lfsr__Vstatic__x = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 11545543387255737976ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__Vstatic__lba32 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 454360589040559504ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_profile_us__Vstatic__range32 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 9284965525539080496ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mix = VL_SCOPED_RAND_RESET_I(12, __VscopeHash, 18234568166550879083ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__boot_off = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 7557652827098347188ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__mft_off = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 3183189246054579892ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__Vstatic__hash_slot = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 16895896662978922891ull);
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__hda_tlp_tx_tvalid = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__hda_tlp_tx_tdata[0U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__hda_tlp_tx_tdata[1U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__hda_tlp_tx_tdata[2U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__hda_tlp_tx_tdata[3U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__hda_tlp_tx_tkeepdw = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__hda_tlp_tx_tlast = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__hda_tlp_tx_tuser = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_base = 0x00001000U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__flash_csn = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__flash_sdi_dq0 = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__flash_wpn_dq2 = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__flash_hldn_dq3 = 1U;
    VL_READMEM_N(true, 32, 1024, 0, "identify_init.hex"s
                 ,  &(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_identify_rom)
                 , 0, ~0ULL);
    vlSelfRef.tb_top__DOT__i_bar__DOT__intr_req = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__flr_in_process = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__pm_dstate = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__memory_space_enable = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__bus_master_enable = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__turnoff_pending = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__link_up = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__generation = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__pcie_id = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__stall = 0U;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__latency_range = 0x0aU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_range = 5U;
    VL_READMEM_N(true, 32, 20, 0, "msix_table_init.hex"s
                 ,  &(vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__msix_table_ram)
                 , 0, ~0ULL);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__k = 0U;
    while (VL_GTS_III(32, 1U, vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__k)) {
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT____Vlvbound_hb093faa3__0 = 0U;
        if (VL_LIKELY(((0U >= (1U & vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__k))))) {
            vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__msix_pba[(1U 
                                                                            & vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__k)] 
                = vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT____Vlvbound_hb093faa3__0;
        }
        vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__k 
            = ((IData)(1U) + vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__k);
    }
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__msix_mode = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__function_enable = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__function_mask = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__nvme_disk_bytes[0U] = 0x77a56000U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__nvme_disk_bytes[1U] = 0x000000eeU;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__nvme_disk_bytes[2U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__nvme_disk_bytes[3U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__pcie_id = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__pcie_id_fmt = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_ready = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_ready_i = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__prog_full = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__prog_empty = 1U;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__in_is_wr_ready = 1U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__f_tkeepdw = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[6U][0U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[6U][1U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_ctx[6U][2U] = 0U;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_data[6U] = 0U;
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
                                                    ((((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en) 
                                                       != (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en__0)) 
                                                      << 1U) 
                                                     | ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_select_ur) 
                                                        != (IData)(vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__cpl_select_ur__0)))));
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__cpl_select_ur__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__cpl_select_ur;
    vlSelfRef.__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en__0 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en;
    if (VL_UNLIKELY(((1U & (~ (IData)(vlSelfRef.__VstlDidInit)))))) {
        vlSelfRef.__VstlDidInit = 1U;
        vlSelfRef.__VstlTriggered[0U] = (1ULL | vlSelfRef.__VstlTriggered[0U]);
        vlSelfRef.__VstlTriggered[0U] = (2ULL | vlSelfRef.__VstlTriggered[0U]);
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
    if ((1U & (IData)((triggers[0U] >> 1U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 1 is active: @([hybrid] tb_top.i_bar.i_fifo_nvme_dma.rd_en)\n");
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

VL_ATTR_COLD void Vtop___024root___stl_sequent__TOP__0(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___stl_sequent__TOP__0\n"); );
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
    CData/*0:0*/ __VdfgRegularize_hebeb780c_0_5;
    __VdfgRegularize_hebeb780c_0_5 = 0;
    CData/*0:0*/ __VdfgRegularize_hebeb780c_0_7;
    __VdfgRegularize_hebeb780c_0_7 = 0;
    CData/*0:0*/ __VdfgRegularize_hebeb780c_0_13;
    __VdfgRegularize_hebeb780c_0_13 = 0;
    CData/*0:0*/ __VdfgRegularize_hebeb780c_0_14;
    __VdfgRegularize_hebeb780c_0_14 = 0;
    // Body
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__int_lfsr_fb 
        = (1U & VL_REDXOR_16((0xd008U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__int_lfsr))));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__msix_rsp_bir 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_rsp_bir;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_delivery_vector 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_vector;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_delivery_done 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_done;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_msix_pba_set_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__pba_set_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_msix_pba_clear_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__pba_clear_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_event_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__pba_set_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_event_vector 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__pba_set_vector;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_msix_vector_select 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__vector_select;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vector_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vector_addr_r;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vector_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vector_data_r;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vector_masked 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vector_masked_r;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_msix_pba_vector 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__pba_vector;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__msix_rsp_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_rsp_valid;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__msix_rsp_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_rsp_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__msix_rsp_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_rsp_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__msix_rsp_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_rsp_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__msix_rsp_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_rsp_data;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tlast 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tlast;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_out 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s0 
           + vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__prng_s3);
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tvalid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_rd_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_format_slot 
        = (0x000000ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx) 
                          >> 7U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_last 
        = (1U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_remaining_dw));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_cc 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw11 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[11U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__empty 
        = (0U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asqs 
        = (0x00000fffU & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_aqa);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1 
        = (((QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[7U])) 
            << 0x00000020U) | (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[6U])));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw12 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[12U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__full 
        = (0x0400U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be 
        = (0x0000000fU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_packet)
                           ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_first_be)
                           : ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_last_be) 
                              | (- (IData)((1U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__remaining_dw)))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_io 
        = (0U != (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[10U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_data;
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
    __VdfgRegularize_hebeb780c_0_14 = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_error)) 
                                       & (3U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__clk = vlSelfRef.tb_top__DOT__clk;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_word 
        = (0x0000007fU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_rsp_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_bar 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__request_bar;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_clear 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_error)) 
           & (2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__issue_dw 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__state;
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_bar = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_bar;
    vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__wr_valid;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_addr 
        = (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr);
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_done 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_delivery_done;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_resp_dbg_state 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_state;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_resp_dbg_qid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_active_qid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_id_rom_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_csts 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_csts;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__pba_set_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_msix_pba_set_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__pba_clear_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_msix_pba_clear_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__event_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_event_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__event_vector 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_event_vector;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vector_select 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_msix_vector_select;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_msix_vector_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vector_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_msix_vector_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vector_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_msix_vector_masked 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vector_masked;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__pba_set_vector 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_msix_pba_vector;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__pba_clear_vector 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_msix_pba_vector;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[6U][0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__msix_rsp_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[6U][1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__msix_rsp_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_ctx[6U][2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__msix_rsp_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar_rsp_data[6U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__msix_rsp_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_emu_busy 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__busy;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_be 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_be;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_be 
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_write 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_full)) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_spill_valid) 
              | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_valid)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_base_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_raw_valid;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_cc 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_cc;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count 
        = (0x0001ffffU & ((IData)(1U) + (0x0000ffffU 
                                         & vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw12)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_full 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__full;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_buf_addr_w;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first_offset 
        = (3U & (((2U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be))
                   ? 1U : ((4U & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be))
                            ? 2U : (- (IData)((1U & 
                                               ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be) 
                                                >> 3U)))))) 
                 & (- (IData)((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_start_be)))))));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_data 
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_read 
        = ((IData)(__VdfgRegularize_hebeb780c_0_14) 
           & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_write)) 
              & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_word_cached)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_write 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_write) 
           & (IData)(__VdfgRegularize_hebeb780c_0_14));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__clk 
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_bar;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_wdata 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_wdata 
           & (- (IData)((1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_clear))))));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__issue_dw;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid;
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
    __VdfgRegularize_hebeb780c_0_5 = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_valid) 
                                      & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__wr_bar));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_addr;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_addr 
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vsel_base 
        = ((- (IData)((5U > (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vector_select)))) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vector_select) 
              << 2U));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__msix_vector_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_msix_vector_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__msix_vector_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_msix_vector_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__vector_masked 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_msix_vector_masked;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en 
        = (1U & vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_reg_cc);
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_total_dw 
        = (vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count 
           << 7U);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count64 
        = (QData)((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tready 
        = (1U & (~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_full)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_wr_en 
        = ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_full)) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tvalid));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__context_address 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first)
            ? (((IData)((vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr 
                         >> 2U)) << 2U) | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__packet_first_offset))
            : (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__current_addr));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_io_cmd_info 
        = ((2U == (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state))
            ? (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_slba)
            : vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_io_progress);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_val 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_val;
    __VdfgRegularize_hebeb780c_0_13 = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_clear) 
                                       | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_write));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__msix_wr_pba_select 
        = ((IData)(__VdfgRegularize_hebeb780c_0_5) 
           & ((0x00003000U <= vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr) 
              & (0x00003008U > vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__msix_wr_table_select 
        = ((IData)(__VdfgRegularize_hebeb780c_0_5) 
           & ((0x00002000U <= vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr) 
              & (0x00002050U > vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr)));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_offset 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_addr;
    __VdfgRegularize_hebeb780c_0_7 = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_wr_hit) 
                                      & (IData)(__VdfgRegularize_hebeb780c_0_5));
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
        = (((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__msix_rsp_valid) 
            << 6U) | (0x0000003fU & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar_base_valid)));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tready 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tready;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_en 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_wr_en;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_pkt_inc 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_dma_fifo_wr_en) 
           & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tlast));
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
        = __VdfgRegularize_hebeb780c_0_13;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_active 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_read) 
           | (IData)(__VdfgRegularize_hebeb780c_0_13));
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
    vlSelfRef.__VdfgRegularize_hebeb780c_0_4 = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_valid) 
                                                & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_bar));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_pba_select 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__msix_wr_pba_select;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_table_select 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__msix_wr_table_select;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_addr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr_bar0;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_req_offset 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_req_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_wr = 
        ((IData)(__VdfgRegularize_hebeb780c_0_7) & 
         ((0x00001000U <= vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0) 
          & (0x00002000U > vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0)));
    tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid 
        = ((~ ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__msix_wr_table_select) 
               | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__msix_wr_pba_select))) 
           & (IData)(__VdfgRegularize_hebeb780c_0_7));
    vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_cc_wr = 
        ((0x00000014U == (0xfffffffcU & vlSelfRef.tb_top__DOT__i_bar__DOT__wr_addr_bar0)) 
         & (IData)(__VdfgRegularize_hebeb780c_0_7));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_adj 
        = (0x000000ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__base_jitter) 
                          + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_offset)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__rd_req_ctx[2U];
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__msix_rd_table_select 
        = ((IData)(vlSelfRef.__VdfgRegularize_hebeb780c_0_4) 
           & ((0x00002000U <= vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr) 
              & (0x00002050U > vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__msix_rd_pba_select 
        = ((IData)(vlSelfRef.__VdfgRegularize_hebeb780c_0_4) 
           & ((0x00003000U <= vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr) 
              & (0x00003008U > vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_addr)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT___unused_wr_pba_select 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_pba_select;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_is_table 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_table_select) 
           & ((0x00002000U <= vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_offset) 
              & (0x00002050U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_offset)));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_in_range 
        = (0x00004000U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_addr);
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_addr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_wr 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_db_wr;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_valid 
        = tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_valid 
        = tb_top__DOT__i_bar__DOT____Vcellinp__i_bar0__wr_valid;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rst 
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj 
        = (0x000000ffU & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adjust) 
                          + (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_adj)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx[2U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[2U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_req_ctx[0U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[0U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_req_ctx[1U] 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__rd_req_ctx[1U];
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_req_ctx[2U] 
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_data 
        = vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__in_ctx 
        = vlSelfRef.tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_table_select 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__msix_rd_table_select;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_req_valid 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__msix_rd_table_select) 
           | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__msix_rd_pba_select));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_pba_select 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__msix_rd_pba_select;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__is_table 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_table_select) 
           & ((0x00002000U <= vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_req_offset) 
              & (0x00002050U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_req_offset)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__is_pba 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_pba_select) 
           & ((0x00003000U <= vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_req_offset) 
              & (0x00003008U > vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_req_offset)));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__addr_hit 
        = ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_req_valid) 
           & ((IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__is_table) 
              | (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__is_pba)));
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__msix_addr_hit 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_msix_table__DOT__addr_hit;
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
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_valid 
        = ((IData)(vlSelfRef.__VdfgRegularize_hebeb780c_0_4) 
           & ((~ (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__msix_addr_hit)) 
              & (IData)(vlSelfRef.tb_top__DOT__i_bar__DOT__bar0_rd_hit)));
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_ready 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__nvme_irq_delivery_ready;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_ready 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_ready;
    vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_tag 
        = vlSelfRef.tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_tag;
}

void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__tlps_in_if__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_ur__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_cpl__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__tlps_dma_out_if__0(Vtop_IfAXIS128* vlSelf);
void Vtop___024root___ico_sequent__TOP__1(Vtop___024root* vlSelf);
void Vtop_IfAXIS128___ico_sequent__TOP__tb_top__DOT__tlps_in_if__1(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__tlps_dma_out_if__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_rdeng__1(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_cpl__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__i_bar__DOT__tlps_ur__0(Vtop_IfAXIS128* vlSelf);
void Vtop_IfAXIS128___ico_comb__TOP__tb_top__DOT__tlps_out_if__0(Vtop_IfAXIS128* vlSelf);
void Vtop___024root___ico_comb__TOP__4(Vtop___024root* vlSelf);

VL_ATTR_COLD void Vtop___024root___eval_stl(Vtop___024root* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vtop___024root___eval_stl\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    if ((1ULL & vlSelfRef.__VstlTriggered[1U])) {
        Vtop___024root___stl_sequent__TOP__0(vlSelf);
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
    if (((1ULL & vlSelfRef.__VstlTriggered[1U]) | (2ULL 
                                                   & vlSelfRef.__VstlTriggered[0U]))) {
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
        vlSelfRef.tb_top__DOT__tlps_dma_out_has_data 
            = vlSymsp->TOP__tb_top__DOT__tlps_dma_out_if.has_data;
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
        Vtop___024root___ico_comb__TOP__4(vlSelf);
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
    if ((1U & (IData)((triggers[0U] >> 1U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 1 is active: @([hybrid] tb_top.i_bar.i_fifo_nvme_dma.rd_en)\n");
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
        VL_DBG_MSGS("         '" + tag + "' region trigger index 1 is active: @([hybrid] tb_top.i_bar.i_fifo_nvme_dma.rd_en)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 2U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 2 is active: @(posedge tb_top.i_bar.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 3U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 3 is active: @(posedge tb_top.i_bar.i_lifecycle_service.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 4U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 4 is active: @(posedge tb_top.i_bar.i_pcileech_tlps128_bar_rdengine.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 5U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 5 is active: @(posedge tb_top.i_bar.i_pcileech_tlps128_bar_rdengine.i_fifo_134_134_clk1_bar_rdrsp.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 6U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 6 is active: @(posedge tb_top.i_bar.i_tlp_ur_completer.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 7U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 7 is active: @(posedge tb_top.i_bar.i_pcileech_tlps128_bar_wrengine.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 8U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 8 is active: @(posedge tb_top.i_bar.i_bar_rsp_arbiter.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 9U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 9 is active: @(posedge tb_top.i_bar.i_bar0.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 0x0000000aU)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 10 is active: @(posedge tb_top.i_bar.i_latency_emu.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 0x0000000bU)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 11 is active: @(posedge tb_top.i_bar.i_msix_table.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 0x0000000cU)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 12 is active: @(posedge tb_top.i_bar.i_nvme_interrupt_service.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 0x0000000dU)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 13 is active: @(posedge tb_top.i_bar.i_nvme_responder.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 0x0000000eU)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 14 is active: @(posedge tb_top.i_bar.i_nvme_bram_disk.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 0x0000000fU)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 15 is active: @(posedge tb_top.i_bar.i_nvme_bram_disk.i_bank0.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 0x00000010U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 16 is active: @(posedge tb_top.i_bar.i_nvme_bram_disk.i_bank1.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 0x00000011U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 17 is active: @(posedge tb_top.i_bar.i_nvme_bram_disk.i_bank2.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 0x00000012U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 18 is active: @(posedge tb_top.i_bar.i_nvme_bram_disk.i_bank3.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 0x00000013U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 19 is active: @(posedge tb_top.i_bar.i_nvme_dma_bridge.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 0x00000014U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 20 is active: @(posedge tb_top.i_bar.i_nvme_dma_bridge.i_dma_tag_service.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 0x00000015U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 21 is active: @(posedge tb_top.i_bar.i_fifo_nvme_dma.clk)\n");
    }
    if ((1U & (IData)((triggers[0U] >> 0x00000016U)))) {
        VL_DBG_MSGS("         '" + tag + "' region trigger index 22 is active: @([true] __VdlySched.awaitingCurrentTime())\n");
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
    VL_SCOPED_RAND_RESET_W(128, vlSelf->tb_top__DOT__tlps_dma_out_tdata, __VscopeHash, 18122047865958817599ull);
    vlSelf->tb_top__DOT__tlps_dma_out_tvalid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 15495265200251862157ull);
    vlSelf->tb_top__DOT__tlps_dma_out_tlast = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 2938307091131388608ull);
    vlSelf->tb_top__DOT__tlps_dma_out_tuser = VL_SCOPED_RAND_RESET_I(9, __VscopeHash, 6522693139142016794ull);
    vlSelf->tb_top__DOT__tlps_dma_out_has_data = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5453980050515980065ull);
    vlSelf->tb_top__DOT__intr_req = 0U;
    ;
    for (int __Vi0 = 0; __Vi0 < 65536; ++__Vi0) {
        vlSelf->tb_top__DOT__host_mem[__Vi0] = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 16517195849062437807ull);
    }
    vlSelf->tb_top__DOT__i = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 1997728582646804290ull);
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
    vlSelf->tb_top__DOT__i_bar__DOT__flash_csn = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__flash_sdi_dq0 = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__flash_sdo_dq1 = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__flash_wpn_dq2 = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__flash_hldn_dq3 = 1U;
    ;
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
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__msix_rsp_ctx, __VscopeHash, 12463354395394348521ull);
    vlSelf->tb_top__DOT__i_bar__DOT__msix_rsp_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 9847730738462496190ull);
    vlSelf->tb_top__DOT__i_bar__DOT__msix_rsp_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8104239154539885829ull);
    vlSelf->tb_top__DOT__i_bar__DOT__msix_rsp_bir = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 10423528821533746648ull);
    vlSelf->tb_top__DOT__i_bar__DOT__msix_addr_hit = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 18400451510105826802ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_irq_event_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17688627952047220367ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_irq_event_vector = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 661008322736147881ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_irq_delivery_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5811604096792863080ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_irq_delivery_ready = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 838750229013223507ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_irq_delivery_done = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4088123199095083130ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_irq_delivery_vector = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 13532806894951480951ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_msix_vector_select = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 9856277132071595002ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_msix_vector_addr = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 10236682997022503647ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_msix_vector_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 9507648944282221936ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_msix_vector_masked = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 3377913723773419287ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_msix_pba_set_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 6683294011089662800ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_msix_pba_clear_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1240708866981026848ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_msix_pba_vector = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 6443425819900082688ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_disk_req_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17385482206379280421ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_disk_req_write = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 13686870057960947136ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_disk_req_flush = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 16449770333013339368ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_disk_req_lba = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 7824899345012185035ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_disk_req_word = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 15108186095161267298ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_disk_req_wdata = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 14317122041420780686ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_disk_req_done = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 3797195181841800576ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_disk_req_rdata = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 1274763274516444096ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_disk_req_hit = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 15526777518896115112ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_disk_busy = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14844692569173539152ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_disk_error = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9248995632233969051ull);
    vlSelf->tb_top__DOT__i_bar__DOT__bar_rsp_ready = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 8262763467872852818ull);
    for (int __Vi0 = 0; __Vi0 < 7; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_data[__Vi0] = 0;
    }
    for (int __Vi0 = 0; __Vi0 < 7; ++__Vi0) {
        VL_ZERO_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT____Vcellinp__i_bar_rsp_arbiter__in_ctx[__Vi0]);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_reg_cc = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 8669844873967991473ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_reg_aqa = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 4255730267622908821ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_reg_asq_lo = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 7814789356713367470ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_reg_asq_hi = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 14910415385081507797ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_reg_acq_lo = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 1614603551273329652ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_reg_acq_hi = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 10612180795465512570ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_reg_csts = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 13777467421369302469ull);
    vlSelf->tb_top__DOT__i_bar__DOT__msix_wr_table_select = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17682700028824613331ull);
    vlSelf->tb_top__DOT__i_bar__DOT__msix_wr_pba_select = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9890188302410712259ull);
    vlSelf->tb_top__DOT__i_bar__DOT__msix_rd_table_select = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 15726465486584053653ull);
    vlSelf->tb_top__DOT__i_bar__DOT__msix_rd_pba_select = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17361523538893476878ull);
    vlSelf->tb_top__DOT__i_bar__DOT__hda_tlp_tx_tvalid = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__hda_tlp_tx_tdata[0] = 0U;
    vlSelf->tb_top__DOT__i_bar__DOT__hda_tlp_tx_tdata[1] = 0U;
    vlSelf->tb_top__DOT__i_bar__DOT__hda_tlp_tx_tdata[2] = 0U;
    vlSelf->tb_top__DOT__i_bar__DOT__hda_tlp_tx_tdata[3] = 0U;
    vlSelf->tb_top__DOT__i_bar__DOT__hda_tlp_tx_tkeepdw = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__hda_tlp_tx_tlast = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__hda_tlp_tx_tuser = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_db_base = 4096U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_db_wr = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8504530893448073785ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_db_off = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 2004215433664124087ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_db_index = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 222379687219757492ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_db_is_cq = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 543102015414429648ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_db_qid = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 8826218301199552916ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_db_val = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 9821607675350912451ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_cc_wr = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17200313886737258641ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_cc_enable_wr = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1172635831064329125ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_cc_disable_wr = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1946123944556769094ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_rd_req = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17335277159205502901ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_rd_addr = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 7807756770767847121ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_rd_len = VL_SCOPED_RAND_RESET_I(10, __VscopeHash, 3228253617734322677ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_rd_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4065181595744171215ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_rd_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 3343843425581243675ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_rd_done = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17256227196231334351ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_wr_req = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8439558904481131297ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_wr_addr = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 15772965027442381779ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_wr_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 12965553839608483385ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_wr_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 9061778806073118593ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_wr_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5269719601486514318ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_wr_done = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 11211501475045917557ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_resp_dbg_state = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 8254079737238731615ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_resp_dbg_qid = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 16911352432425939078ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_resp_dbg_opcode = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 4588255769111018506ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_resp_dbg_admin_queues = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 2313452200819230240ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_resp_dbg_cmd_info = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 14793746835700777233ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_disk_dbg_state = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 2586817737155681821ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_disk_dbg_info = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 15534914912897260018ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_dbg_status = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 4339269120561427968ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_id_rom_addr = VL_SCOPED_RAND_RESET_I(13, __VscopeHash, 13136942424550301575ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_id_rom_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 12821244732960813322ull);
    VL_SCOPED_RAND_RESET_W(128, vlSelf->tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tdata, __VscopeHash, 12092384518282411973ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tkeepdw = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 8619408969991890920ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tvalid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17604692481445826098ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tlast = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 15134035120668268988ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tuser = VL_SCOPED_RAND_RESET_I(9, __VscopeHash, 10531215428634904838ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_tlp_tx_tready = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9941742549655877152ull);
    for (int __Vi0 = 0; __Vi0 < 1024; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__nvme_identify_rom[__Vi0] = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 3738422675862756986ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_id_rom_data_r = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 434349248051731132ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_fifo_full = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 15677094171268968514ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_fifo_first = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 3631535721124108748ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tlast = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 10457833728012173701ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tkeepdw = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 12701381819835912231ull);
    VL_SCOPED_RAND_RESET_W(128, vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tdata, __VscopeHash, 11971425328835425545ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_fifo_tvalid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 6850828800181328525ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_pkt_count = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 17180855208130914202ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_fifo_wr_en = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1303482720268667857ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_pkt_inc = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 6954633512149926255ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_pkt_dec = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 6475212916644598779ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_pkt_count_next = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 14652485815434128265ull);
    vlSelf->tb_top__DOT__i_bar__DOT__nvme_dma_fifo_rd_en = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 3917361852569546798ull);
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
    vlSelf->tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__stall = 0U;
    ;
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
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_cc = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 537818270758218058ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_aqa = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 16540471005240628938ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_asq_lo = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 5260314836426649685ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_asq_hi = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 3799613631311469199ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_acq_lo = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 9119311859301735186ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_acq_hi = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 10588010832992302910ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__nvme_csts = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 1654347338749415431ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_in_range = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5258355185944826704ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_in_range = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 6481555245557248856ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__wr_offset = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 8894269562469001849ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_offset = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 5167266672372387941ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000000 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 3728280704661799873ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000004 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 10606760971621188864ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000008 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 7629256279538641629ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000000C = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 2653296426762923664ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000010 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 1036152268255986236ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000014 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 12778201109914495245ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000001C = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 9726047732956085603ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000020 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 11050875964898839624ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000024 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 4696042073158797087ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000028 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 12430764041029538689ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x0000002C = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 12429197679661041942ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000030 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 10797317055431677608ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__reg_0x00000034 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 1959209147949253626ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__cc_en_prev = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 2529469566037103627ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_transition_cnt = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 14817181103505343280ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__csts_rdy_pending = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 115596810741523246ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__int_state = VL_SCOPED_RAND_RESET_I(2, __VscopeHash, 12236585396928239405ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__int_counter = VL_SCOPED_RAND_RESET_I(6, __VscopeHash, 4144307544523187727ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__int_cooldown_target = VL_SCOPED_RAND_RESET_I(6, __VscopeHash, 4133062662255989378ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__int_trigger = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14513379356873557507ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__int_lfsr = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 14554662465883794061ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__int_lfsr_fb = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17399709505875394386ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_ctx_d1, __VscopeHash, 13252563380064055013ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_req_valid_d1 = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 2294618463596330379ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_bar0__DOT__rd_data_d1 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 6233769887282796986ull);
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
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__latency_range = 10U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__cdf_latency = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 13679687949906105547ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__base_jitter = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 4784813292563664761ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__thermal_adj = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 5130485241660675617ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__region_adj = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 16337058534593895849ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__jittered = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 668108554917362555ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__computed_latency = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 9904913569405222238ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_range = 5U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_latency = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 172305930387999504ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_counter = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 12092512540757580383ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_target = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 15131728473859726211ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__wr_pending = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9652697969904091117ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT____Vlvbound_hef7352ea__0 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT____Vlvbound_h832e0042__0 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT____Vlvbound_hb093faa3__1 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT____Vlvbound_hc930d0d9__0 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT____Vlvbound_hc930ceef__0 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT____Vlvbound_hc930cced__0 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT____Vlvbound_hc930cca3__0 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT____Vlvbound_h358caa98__0 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT____Vlvbound_hb093faa3__0 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 10714292850726071365ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8808102547344819598ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_addr = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 3882971745317870584ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 10984969892556434942ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 7485636704085039229ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 362186790198744220ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_table_select = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17503774055487159434ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_pba_select = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14628014310011108859ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_req_ctx, __VscopeHash, 15968587240082671478ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_req_addr = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 16573270785597565693ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_req_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14953011439997972253ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_table_select = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4258345268775015408ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_pba_select = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 2761540327017474863ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_rsp_ctx, __VscopeHash, 14547611974000837117ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_rsp_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 12254422864492229988ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_rsp_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 2861993909741861345ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_rsp_bir = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 2526597843891184969ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vector_select = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 5956922888445565759ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vector_addr = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 4528918622621943677ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vector_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 1082680238013458618ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vector_masked = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17731628628565539279ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__pba_set_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 13625444318062560326ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__pba_set_vector = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 15889065089389942707ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__pba_clear_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 7479934707158792322ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__pba_clear_vector = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 7335096778001864273ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__addr_hit = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14737184720329138448ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_offset = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 388040489796313723ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_req_offset = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 1137526064865447638ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__is_table = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9645403369122545257ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__is_pba = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9259028350301390875ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__wr_is_table = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5471309418936009568ull);
    for (int __Vi0 = 0; __Vi0 < 20; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__msix_table_ram[__Vi0] = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 12794542344518410649ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vsel_base = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 7013639219395393273ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vector_addr_r = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 6893772530875828727ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vector_data_r = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 15094503730205756576ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__vector_masked_r = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1078215459866979558ull);
    for (int __Vi0 = 0; __Vi0 < 1; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__msix_pba[__Vi0] = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 17531787056671840687ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT___unused_wr_pba_select = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5215739509313397282ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__k = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 6071653874038180977ull);
    VL_SCOPED_RAND_RESET_W(88, vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_ctx_d1, __VscopeHash, 1340616194571009227ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_valid_d1 = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 2902022125541753297ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_data_d1 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 7357270752447784441ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_msix_table__DOT__rd_bir_d1 = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 9158284888136392163ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT____Vlvbound_h4dcaf4d4__2 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT____Vlvbound_h4dcaf4d4__0 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT____Vlvbound_hcff0ff53__1 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT____Vlvbound_hcff0ff53__0 = 0;
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17135168760282839683ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__rst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1078034902560137759ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__quiesce = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 77940560647245032ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__msix_mode = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__function_enable = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__function_mask = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__event_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1553756471864106507ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__event_vector = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 2698291796422966977ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__vector_select = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 7837982979936732659ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__vector_masked = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4798752618667351041ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_ready = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4432179167943985311ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_done = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 15295485977494558309ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 10395701656123662475ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__delivery_vector = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 15972990669175782957ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__msi_pulse = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 10033736600489223546ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__pba_set_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1411990150823318363ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__pba_clear_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17412141283388963183ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__pba_vector = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 12276967738276480005ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__pending = VL_SCOPED_RAND_RESET_I(5, __VscopeHash, 13022122045640304134ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__scan_vector = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 6372604847862577589ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__query_state = VL_SCOPED_RAND_RESET_I(2, __VscopeHash, 12267754228754880573ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__rst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8374208448887162100ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5560784942313036014ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_enabled = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14404255755672383870ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 10623084226353941484ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_enable_wr = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 6902221941774773720ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_disable_wr = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1798205919226776188ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_lo = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 7681392513854142921ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_hi = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 1756565525549977343ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_lo = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 2813305201716432972ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_hi = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 9930486539667673995ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aqa = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 11090345888576087584ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_wr = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 16638152667140687155ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_is_cq = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1819263591339614061ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_qid = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 17887870853548896529ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__doorbell_val = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 5587121051865306521ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__msix_vector_addr = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 2245381962701128110ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__msix_vector_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 3472092375219160269ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14785045549005177054ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_ready = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9832973957327330614ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__irq_delivery_done = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5441715316006424380ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_req = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 13486882024116681238ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_addr = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 3745095178185415763ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_len = VL_SCOPED_RAND_RESET_I(10, __VscopeHash, 13354601471876704443ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 12555881710017272916ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 6184048928439378873ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_rd_done = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 16001929720575649705ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_req = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8708021096122777360ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_addr = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 12364054469121580236ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 9784236779952739ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 5427052919730509944ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 18158234797755876898ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dma_wr_done = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17407146212815224744ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17926882486245186886ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_write = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17461742719828719250ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_flush = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8934708032412785420ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_lba = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 6768129618623085999ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_word = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 14369139016507635723ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_wdata = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 15002375389826713240ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_done = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 3368175946235867945ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_rdata = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 1582557154988748999ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_req_hit = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 12059313931742403574ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_busy = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 11375905833768409770ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__disk_error = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1972223823329281856ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__msix_trigger = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8179608880052825837ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__pba_set_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8804492106927250775ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__pba_set_vector = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 11915766877443325934ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_addr = VL_SCOPED_RAND_RESET_I(13, __VscopeHash, 12411736476511199770ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 6059571878506435091ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_state = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 10816117282904560663ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_active_qid = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 1677237692620454721ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_opcode = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 8910288007487193915ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_admin_queues = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 5799735656171329806ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_cmd_info = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 1508349451478433234ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__state = VL_SCOPED_RAND_RESET_I(6, __VscopeHash, 465114042609270620ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__return_state = VL_SCOPED_RAND_RESET_I(6, __VscopeHash, 2561287033495045621ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_qid = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 12280459626764343289ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_tail = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 17112345083572860595ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 16507726772307119801ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 8331461347878179971ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_head_db = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 7272936262323904658ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_phase = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 18063619916538258042ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_aqa = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 13492688802147304689ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_lo = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 9940932311227855551ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_hi = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 15877128871871626740ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_lo = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 7748191779726455826ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_hi = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 16063721466896340168ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asqs = VL_SCOPED_RAND_RESET_I(12, __VscopeHash, 165616618744694736ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acqs = VL_SCOPED_RAND_RESET_I(12, __VscopeHash, 10851753693264790526ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_last = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 5310254562248290181ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_last = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 7635368295444718539ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_power2_entries = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 15284622815773927486ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_power2_entries = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9158846077882679310ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sq_head_next = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 7998262889808146332ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cq_tail_next = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 806825026813289369ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__asq_base = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 2719813128290882049ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__acq_base = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 15517063428096119273ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_aqa_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 979849564060048032ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_asq_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9341177064847087163ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_acq_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8330837996895064503ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__admin_queue_config_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 6739487895490633306ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_created = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 16001038117852965834ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_created = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5843317124488295038ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_base = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 8030816494796343267ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_base = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 15919566762509118411ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_last = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 9047349582666677211ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_last = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 2894922645661536101ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_tail = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 13522302615858195427ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 10658020517528188092ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 2869903690912635105ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_head_db = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 411841125435597847ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_phase = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 10448496377068236112ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_vector = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 15519517625736625472ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_power2_entries = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 3564309076490825844ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_power2_entries = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17803675984986054309ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_sq_head_next = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 7796364383660866948ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_cq_tail_next = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 7037865179292901841ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_io = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 12325722384043390622ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_sq_head_next = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 15867295938227881358ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_tail = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 11763021199574736400ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_last = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 11289963252163578892ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_base = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 14790276190143197892ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__active_cq_phase = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 7692849561073473967ull);
    for (int __Vi0 = 0; __Vi0 < 16; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe[__Vi0] = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 2301130455592482459ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__sqe_dw_idx = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 5527990213519523514ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_opcode = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 10927969968037077853ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cid = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 839111285927283174ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 7208637401361848724ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1 = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 13313866864956318014ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp2 = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 3058459308224998704ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_prp1_page_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 497243762214223429ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw10 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 11796429673288302416ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw11 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 17661498587380889516ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw12 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 4462530418554068802ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw13 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 2508373185105621475ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw14 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 15752284852992766436ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_cdw15 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 890149375104307236ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__identify_cns = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 8059259680631336307ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid_selects_ns1 = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 16050230746574759093ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nsid_lists_ns1 = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 15679832917527958282ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_slba = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 10102883829411898420ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count = VL_SCOPED_RAND_RESET_I(17, __VscopeHash, 12772156178019238404ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_nlb_count64 = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 11865471337102966225ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_lba_ok = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5418594534065307600ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__format_nvm_ok = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9224555531650111670ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_total_dw = VL_SCOPED_RAND_RESET_I(24, __VscopeHash, 5296199930928450428ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_mdts_ok = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17294100648414811515ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__nvme_disk_bytes[0] = 2007326720U;
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__nvme_disk_bytes[1] = 238U;
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__nvme_disk_bytes[2] = 0U;
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__nvme_disk_bytes[3] = 0U;
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cmd_log_last_dw = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 16871152109500304671ull);
    for (int __Vi0 = 0; __Vi0 < 4; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe[__Vi0] = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 2739667269691183508ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_dw_idx = VL_SCOPED_RAND_RESET_I(2, __VscopeHash, 8489516850196360861ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_result = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 6733420207533437609ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cqe_status = VL_SCOPED_RAND_RESET_I(15, __VscopeHash, 15093076557093243519ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_count = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 12347052164612580862ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_sqid = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 13631201996265457081ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_cid = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 16376725168925930447ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_status = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 9287844130897155474ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_peloc = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 12005570767705431284ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_lba = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 15665767876949129568ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__last_error_nsid = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 13171084827986263406ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_class = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 5584884225906574261ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_effective_class = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 4803296625123852736ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_done_for_cmd = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 3439907587488337622ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_cmd_active = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 10262565242009001437ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_saturated = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 13399062764671326873ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_read_seen_hit = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4322842419516282382ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_read_seen_miss = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9556214986734030454ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_lfsr = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 13110957379626352005ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_jitter_sample = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 4336797432742959217ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_elapsed_cycles = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 13741219262301897667ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_target_cycles = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 15717459303928253821ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_wait_cycles = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 10784713064643503159ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__data_dw_cnt = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 16848140211632952408ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__id_rom_offset = VL_SCOPED_RAND_RESET_I(13, __VscopeHash, 255795410288500411ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_read = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 1373395701851440600ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_data_units_written = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 15855354239639435643ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_read_cmds = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 8730490244505465852ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_host_write_cmds = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 531964622184635178ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_flush_cmds = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 17551883811551678696ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_write_zero_cmds = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 15325079549621454246ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_dataset_cmds = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 10248870159551999315ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_cmds_completed = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 5612840164490035223ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_media_errors = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 9008505725966244687ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_error_log_entries = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 3315081245690217326ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__stat_unsafe_shutdowns = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 4470891771038904538ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__power_cycle_count = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 15799520210331832641ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__power_on_hours = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 2051671091205411799ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__hour_second_count = VL_SCOPED_RAND_RESET_I(12, __VscopeHash, 17624197013591529808ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__second_tick = VL_SCOPED_RAND_RESET_I(27, __VscopeHash, 12458744952260341397ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__thermal_load = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 11574179674762602752ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__warning_temp_time = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 13265479164289295062ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__critical_temp_time = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 4138833931657090704ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__smart_temp_k = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 13506070747487392945ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__smart_warning_byte = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 4572840681325428987ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_en_prev = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1894513899615135725ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_start_event = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 11385316546079206897ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_stop_event = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 7796315705765349009ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__cc_reset_event = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9098345735113651224ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_registered = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1286242396688593326ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_cid = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 3131760298963241831ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_pending = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 13878557536716447900ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__aer_event_info = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 9276310299021396519ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__num_queues_granted = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 9949152215423390831ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_arbitration = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 14008356428496782564ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_power_mgmt = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 14023131558029722041ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_temp_threshold = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 9448404086557917607ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_write_cache = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 13673331703865510013ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_irq_coalescing = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 2306898509595915215ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__feat_async_event_cfg = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 3939250470406150635ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dir = VL_SCOPED_RAND_RESET_I(2, __VscopeHash, 1537999104526396597ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_slba = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 7963260711226820803ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_total_dw = VL_SCOPED_RAND_RESET_I(24, __VscopeHash, 11718172227581120539ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_dw_idx = VL_SCOPED_RAND_RESET_I(24, __VscopeHash, 10353578128377988388ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 8389016980508514691ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_host_addr = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 13971803847978466703ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_entry_lo = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 118186217471594504ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_entry_hi = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 14453897373023252228ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 7286376904705045256ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_index = VL_SCOPED_RAND_RESET_I(20, __VscopeHash, 3029879208278998064ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_cached_base = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 10793058729533341895ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_idx = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 6262988383229460317ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_last = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 9729963209516651969ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw_idx = VL_SCOPED_RAND_RESET_I(2, __VscopeHash, 3894523172989459736ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw0 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 7433947725824182808ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw1 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 3650111357847631182ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw2 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 14588289386520474205ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw3 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 1308789823545202874ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_byte_off = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 14900977069532916858ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_first_span = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 10013622227731504342ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_after_first = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 6765703876781788988ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_total_bytes = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 12995357972985532679ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_after_first_page = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 16679317562092589649ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp2_required = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 2973899264651788323ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_uses_list = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 15367215554468483932ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_list_index = VL_SCOPED_RAND_RESET_I(20, __VscopeHash, 9623634257898809762ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp_list_entry_addr = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 13248607072263041664ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_entry_hi_value = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 17147137676091181588ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_list_entry_value = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 11153631477344225924ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_list_page_base = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 17932145111647982472ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp1_invalid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9152075572119521461ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_prp2_invalid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4380591145494998068ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__prp_list_entry_invalid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 12670828864468708669ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_dw3_value = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 6921902964422812411ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_lba_value = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 2035852607084484549ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_nlb64 = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 4509141265089762897ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_next = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 4273560898933493879ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_dsm_ranges = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 15746254246937721904ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dsm_range_ok = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 12239892563179147593ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_disk_lba = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 3509414882370490678ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_disk_word = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 16153920796680088646ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__io_last_dw = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 16744321075865711590ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__timing_completion_class = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 12954408813865674336ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_timing_info = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 14395849889299167256ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_io_progress = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 13837625544970136620ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__dbg_io_cmd_info = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 5100059953931712534ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__rst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 3341303188730569055ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 18277038309590163484ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9036773487047670990ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_write = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 3234258617600923069ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_flush = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 3548161580347184934ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_lba = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 17413338283142126395ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_word = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 15259809950988641484ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_wdata = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 17475314502882528553ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_done = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17288770066388217453ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_rdata = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 14796224491701818433ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_hit = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 2427147489146923740ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__busy = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4174351227264527715ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__error = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 16183667995848011764ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__dbg_state = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 12203873651467606565ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__dbg_info = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 8007141473245573234ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__state = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 8823083449891196142ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_error = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 6723816595435378336ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_clear = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14739752274854773786ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_full_format = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 16067978702899397107ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_write = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 13861388940585663776ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_wdata = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 9612388101445024285ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_lba = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 15177340815966041669ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 306323348515647415ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_word = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 7043701140755420853ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_hit = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9572714715667584981ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_base_lba = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 16201814223889902211ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_idx = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 2973215957408115872ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_total_dw = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 888552457569005937ull);
    for (int __Vi0 = 0; __Vi0 < 256; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_valid[__Vi0] = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9388298659949071775ull);
    }
    for (int __Vi0 = 0; __Vi0 < 256; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__cache_tag[__Vi0] = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 3291295853686468069ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__read_bank_q = VL_SCOPED_RAND_RESET_I(2, __VscopeHash, 12148908909716803763ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_armed = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 723375322760487058ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pin_boot_lba = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 4920597372221959123ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_spc = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 2193365334405859462ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__fmt_mft_lcn = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 17170127458521002610ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mft_armed = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5532439689741339975ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pin_mft_lba = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 11482356386352864945ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_detect_now = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1242352086655103492ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__boot_route_now = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 16279095307103468258ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__eff_boot_armed = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 10612045486269503047ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__eff_pin_boot = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 3908322378496606174ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_slot = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 15326734780097246130ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_slot_hit = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 12000560666385330370ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_word_cached = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 2306077595758449971ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_word_cached = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 16996477406790702487ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_lba = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 9955990344382062813ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_word = VL_SCOPED_RAND_RESET_I(7, __VscopeHash, 4816204177450693183ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_format_slot = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 4833169102432741378ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_slot = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 17794458401082519764ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clear_ram_addr = VL_SCOPED_RAND_RESET_I(15, __VscopeHash, 13389296616777390612ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__data_ram_addr = VL_SCOPED_RAND_RESET_I(15, __VscopeHash, 4697494446305588279ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_clear = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4665537733168184912ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_write = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 11293513402443123493ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_read = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 16261868675650540881ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_addr = VL_SCOPED_RAND_RESET_I(15, __VscopeHash, 3467468378276940222ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_bank = VL_SCOPED_RAND_RESET_I(2, __VscopeHash, 6927978140623206180ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_bank_addr = VL_SCOPED_RAND_RESET_I(13, __VscopeHash, 9087382969578679658ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_wdata = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 4234037505201927643ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_active = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 7646389022622982668ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__mem_we = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 11054590964939750910ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout0 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 10759023317139370813ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout1 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 4622934112223692458ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout2 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 8159109638006392968ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__bank_dout3 = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 8834126392371102896ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__ram_dout = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 671758813941307940ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__req_slot_dbg = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 6849300772666651439ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__pending_slot_dbg = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 6905525257644714220ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__reset_idx = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 15874223824218302380ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9435525972392720863ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__en = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 2295401110517612946ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__we = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 7416201998297326321ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__addr = VL_SCOPED_RAND_RESET_I(13, __VscopeHash, 6884159826883572442ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__din = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 4769746575275404792ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__dout = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 8428631030601675158ull);
    for (int __Vi0 = 0; __Vi0 < 8192; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__ram[__Vi0] = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 12616848131435677411ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5684477455675164377ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__en = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 11623742735992643638ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__we = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 13376976656469927283ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__addr = VL_SCOPED_RAND_RESET_I(13, __VscopeHash, 12145273674165444837ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__din = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 5811894222761303892ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__dout = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 4584477787289608078ull);
    for (int __Vi0 = 0; __Vi0 < 8192; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__ram[__Vi0] = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 3467148404343957862ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5128668758808369320ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__en = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8278234277616891470ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__we = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14958860282362220927ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__addr = VL_SCOPED_RAND_RESET_I(13, __VscopeHash, 6826139120566897869ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__din = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 5191175547989875574ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__dout = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 4685503635374215584ull);
    for (int __Vi0 = 0; __Vi0 < 8192; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__ram[__Vi0] = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 6559173199219006707ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 7391359395549155517ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__en = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4289010669773329253ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__we = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 15987155550188960440ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__addr = VL_SCOPED_RAND_RESET_I(13, __VscopeHash, 15536313347216083328ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__din = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 3453680650078691129ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__dout = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 18280008520333699871ull);
    for (int __Vi0 = 0; __Vi0 < 8192; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__ram[__Vi0] = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 14958491761007876328ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5065511962873266433ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17047543068013298066ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_enabled = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 10468737782606301755ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__pcie_id = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_req = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14172237584157322434ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_addr = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 9066083318963358223ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 8167154892676587543ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 12926465030993167480ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14270411284323286478ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_wr_done = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 13757848006665587683ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_req = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5150888322713771741ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_addr = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 17780605958548574911ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_len = VL_SCOPED_RAND_RESET_I(10, __VscopeHash, 5106753163422898946ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1031173711025874490ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 9381118403240508187ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dma_rd_done = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17047726853654964711ull);
    VL_SCOPED_RAND_RESET_W(128, vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tdata, __VscopeHash, 9596072572516365972ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tkeepdw = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 16562049720940437383ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tvalid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 10225291282110205696ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tlast = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17037200881335428049ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tuser = VL_SCOPED_RAND_RESET_I(9, __VscopeHash, 8437267024938468732ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_tx_tready = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 16737517907139796155ull);
    VL_SCOPED_RAND_RESET_W(128, vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_rx_tdata, __VscopeHash, 8328418754127983138ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_rx_tkeepdw = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 15359263091120735396ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_rx_tvalid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17802745848788511368ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tlp_rx_tuser = VL_SCOPED_RAND_RESET_I(9, __VscopeHash, 5125261424318089553ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__dbg_status = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 2316277919653084131ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__bstate = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 5476815249279951387ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_active_tag = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 2089561522587252736ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_base_addr = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 15025168060322819130ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_remaining = VL_SCOPED_RAND_RESET_I(10, __VscopeHash, 14263740235800596426ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_index = VL_SCOPED_RAND_RESET_I(10, __VscopeHash, 11611128502814405179ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_payload = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 16807109545955877571ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_data_hold = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 11550741331184148647ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1950792775154454430ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_addr = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 406304063189286973ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_data = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 14845232013747772281ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__wr_pending_be = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 468246657181919418ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1489650789747668881ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_addr = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 7975540911828024387ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_pending_len = VL_SCOPED_RAND_RESET_I(10, __VscopeHash, 17616039095801407961ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__pcie_id_fmt = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_next_index = VL_SCOPED_RAND_RESET_I(10, __VscopeHash, 2158507048213473325ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__rd_next_addr = VL_SCOPED_RAND_RESET_Q(64, __VscopeHash, 14201108219183945237ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_status = VL_SCOPED_RAND_RESET_I(3, __VscopeHash, 10758577085228257386ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_has_data = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 5099484183554124387ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_any = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17763983885041689657ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_tag_match = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 17425060473454555249ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_match = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 6303064321899624759ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__cpld_error = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 2486437531116568541ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_ready = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 362145795831397087ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_tag = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 12887357711988223760ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_alloc_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14654437735105827869ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_outcome_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8960563299421136328ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_outcome_tag = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 10442718784117243010ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__tag_outcome_status = VL_SCOPED_RAND_RESET_I(2, __VscopeHash, 6795631350168929059ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 1261553614682090214ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__rst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 11521105872274912487ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 562852816998158419ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_ready = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 15345041616884453352ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_tag = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 9203543538570456689ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completion_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 3223803945681922177ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completion_tag = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 15503122455432773668ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completion_error = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14047500899859136954ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_all = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 12297570812183084562ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 7275460823458231932ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_ready = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_tag = VL_SCOPED_RAND_RESET_I(8, __VscopeHash, 4263687449864094471ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_status = VL_SCOPED_RAND_RESET_I(2, __VscopeHash, 10326114130294293428ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__active_tags = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 12460480487303202034ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outstanding_count = VL_SCOPED_RAND_RESET_I(5, __VscopeHash, 16378783125112790505ull);
    for (int __Vi0 = 0; __Vi0 < 16; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__age[__Vi0] = VL_SCOPED_RAND_RESET_I(24, __VscopeHash, 15482334842705627693ull);
    }
    for (int __Vi0 = 0; __Vi0 < 16; ++__Vi0) {
        vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_status[__Vi0] = VL_SCOPED_RAND_RESET_I(2, __VscopeHash, 13512408634877381039ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancelled = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 3370004220070818474ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_reported = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 9817090862806711763ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_report_pending = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 2421166069438320201ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_pending = VL_SCOPED_RAND_RESET_I(16, __VscopeHash, 6722716355842008333ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__alloc_cursor = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 3223315666424923072ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__i = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 1534052697666785966ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__scan_index = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 1330426652893481076ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__selected = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 10233278274709384336ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__completed_index = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 4118121593814849169ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_index = VL_SCOPED_RAND_RESET_I(32, __VscopeHash, 3411670507543219435ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__timeout_found = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 15018804532170549959ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__timeout_index = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 16666708952262149927ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_found = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 553001307185020595ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__cancel_index = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 12192658328049509863ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_found = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 16671110082614852944ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__terminal_index = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 2838401703237076849ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__outcome_ready_i = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__srst = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14645292825222881219ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__clk = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 13470934246409434762ull);
    VL_SCOPED_RAND_RESET_W(134, vlSelf->tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__din, __VscopeHash, 14560660056759004756ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_en = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 4042211417518746790ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 14711987746339772087ull);
    VL_SCOPED_RAND_RESET_W(134, vlSelf->tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__dout, __VscopeHash, 17641860040881840165ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__full = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 2557251811126991167ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__empty = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 6852169048173512173ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__prog_empty = 1U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_data_count = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 12646118953469036355ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_data_count = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 11215497766907885804ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__prog_full = 0U;
    ;
    vlSelf->tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__valid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 8916805621325296189ull);
    for (int __Vi0 = 0; __Vi0 < 1024; ++__Vi0) {
        VL_SCOPED_RAND_RESET_W(134, vlSelf->tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__mem[__Vi0], __VscopeHash, 8621837382512221566ull);
    }
    vlSelf->tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__wr_ptr = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 9462263626768782445ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_ptr = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 7939601189962645479ull);
    vlSelf->tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__count = VL_SCOPED_RAND_RESET_I(11, __VscopeHash, 14028814322490572084ull);
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
    vlSelf->__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__37__lba = 0;
    vlSelf->__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__38__lba = 0;
    vlSelf->__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__38__word_off = 0;
    vlSelf->__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_index__38____VlefCall_0__backing_slot = 0;
    vlSelf->__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__39__Vfuncout = 0;
    vlSelf->__Vfunc_tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__backing_slot__39__lba = 0;
    vlSelf->__VdfgRegularize_hebeb780c_0_4 = 0;
    vlSelf->__VdfgRegularize_hebeb780c_0_9 = 0;
    for (int __Vi0 = 0; __Vi0 < 2; ++__Vi0) {
        vlSelf->__VstlTriggered[__Vi0] = 0;
    }
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__cpl_select_ur__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en__0 = 0;
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
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__rd_en__1 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_lifecycle_service__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_rdengine__DOT__i_fifo_134_134_clk1_bar_rdrsp__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_tlp_ur_completer__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_pcileech_tlps128_bar_wrengine__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_bar_rsp_arbiter__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_bar0__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_latency_emu__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_msix_table__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_interrupt_service__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_responder__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank0__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank1__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank2__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_bram_disk__DOT__i_bank3__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_nvme_dma_bridge__DOT__i_dma_tag_service__DOT__clk__0 = 0;
    vlSelf->__Vtrigprevexpr___TOP__tb_top__DOT__i_bar__DOT__i_fifo_nvme_dma__DOT__clk__0 = 0;
    for (int __Vi0 = 0; __Vi0 < 1; ++__Vi0) {
        vlSelf->__VnbaTriggered[__Vi0] = 0;
    }
    vlSelf->__Vi = 0;
}
