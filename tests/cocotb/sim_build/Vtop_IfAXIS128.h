// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design internal header
// See Vtop.h for the primary calling header

#ifndef VERILATED_VTOP_IFAXIS128_H_
#define VERILATED_VTOP_IFAXIS128_H_  // guard

#include "verilated.h"
#include "verilated_timing.h"


class Vtop__Syms;

class alignas(VL_CACHE_LINE_BYTES) Vtop_IfAXIS128 final {
  public:

    // DESIGN SPECIFIC STATE
    CData/*3:0*/ tkeepdw;
    CData/*0:0*/ tvalid;
    CData/*0:0*/ tlast;
    CData/*0:0*/ tready;
    CData/*0:0*/ has_data;
    CData/*0:0*/ __VdfgRegularize_hebeb780c_0_3;
    SData/*8:0*/ tuser;
    VlWide<4>/*127:0*/ tdata;

    // INTERNAL VARIABLES
    Vtop__Syms* vlSymsp;
    const char* vlNamep;

    // CONSTRUCTORS
    Vtop_IfAXIS128();
    ~Vtop_IfAXIS128();
    void ctor(Vtop__Syms* symsp, const char* namep);
    void dtor();
    VL_UNCOPYABLE(Vtop_IfAXIS128);

    // INTERNAL METHODS
    void __Vconfigure(bool first);
};

std::string VL_TO_STRING(const Vtop_IfAXIS128* obj);

#endif  // guard
