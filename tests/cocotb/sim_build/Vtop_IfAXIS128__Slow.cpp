// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vtop.h for the primary calling header

#include "Vtop__pch.h"

void Vtop_IfAXIS128___ctor_var_reset(Vtop_IfAXIS128* vlSelf);

Vtop_IfAXIS128::Vtop_IfAXIS128() = default;
Vtop_IfAXIS128::~Vtop_IfAXIS128() = default;

void Vtop_IfAXIS128::ctor(Vtop__Syms* symsp, const char* namep) {
    vlSymsp = symsp;
    vlNamep = strdup(Verilated::catName(vlSymsp->name(), namep));
    // Reset structure values
    Vtop_IfAXIS128___ctor_var_reset(this);
}

void Vtop_IfAXIS128::__Vconfigure(bool first) {
    (void)first;  // Prevent unused variable warning
}

void Vtop_IfAXIS128::dtor() {
    VL_DO_DANGLING(std::free(const_cast<char*>(vlNamep)), vlNamep);
}
