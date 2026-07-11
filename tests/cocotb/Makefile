SIM ?= verilator
TOPLEVEL_LANG ?= verilog

OUT_DIR := out
RTL := $(wildcard $(OUT_DIR)/*.sv) $(wildcard $(OUT_DIR)/src/pcileech_bar_impl_none.sv) $(wildcard $(OUT_DIR)/src/pcileech_bar_impl_zerowrite4k.sv) $(wildcard $(OUT_DIR)/src/pcileech_tlps128_bar_wrengine.sv)
STUBS := cocotb_stubs.sv

VERILOG_SOURCES := $(RTL) $(STUBS) tb_top.sv
TOPLEVEL := tb_top
MODULE := test_pcie_host
SIM_BUILD := sim_build

EXTRA_ARGS += -Wno-fatal -I$(OUT_DIR) -I$(OUT_DIR)/src
EXTRA_ARGS += --timing

include $(shell cocotb-config --makefiles)/Makefile.sim
