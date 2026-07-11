// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vtop.h for the primary calling header

#include "Vtop__pch.h"

VL_ATTR_COLD void Vtop_IfAXIS128___ctor_var_reset(Vtop_IfAXIS128* vlSelf) {
    VL_DEBUG_IF(VL_DBG_MSGF("+            Vtop_IfAXIS128___ctor_var_reset\n"); );
    Vtop__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    auto& vlSelfRef = std::ref(*vlSelf).get();
    // Body
    const uint64_t __VscopeHash = VL_MURMUR64_HASH(vlSelf->vlNamep);
    VL_SCOPED_RAND_RESET_W(128, vlSelf->tdata, __VscopeHash, 17793882294932938261ull);
    vlSelf->tkeepdw = VL_SCOPED_RAND_RESET_I(4, __VscopeHash, 16065406661757007371ull);
    vlSelf->tvalid = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 11580290079671979871ull);
    vlSelf->tlast = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 9634498217090403279ull);
    vlSelf->tuser = VL_SCOPED_RAND_RESET_I(9, __VscopeHash, 16187945839469799350ull);
    vlSelf->tready = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 6606742779220347250ull);
    vlSelf->has_data = VL_SCOPED_RAND_RESET_I(1, __VscopeHash, 7658414166164095890ull);
    vlSelf->__VdfgRegularize_hebeb780c_0_3 = 0;
}
