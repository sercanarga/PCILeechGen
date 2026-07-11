SIM ?= verilator
TOPLEVEL_LANG ?= verilog

FIXTURE ?= out
RTL := $(wildcard $(FIXTURE)/*.sv)
SRC_NEEDED := pcileech_bar_impl_none.sv pcileech_bar_impl_zerowrite4k.sv pcileech_tlps128_bar_wrengine.sv
RTL += $(foreach f,$(SRC_NEEDED),$(wildcard $(FIXTURE)/src/$(f)))
STUBS := cocotb_stubs.sv

VERILOG_SOURCES := $(RTL) $(STUBS) tb_top.sv
TOPLEVEL := tb_top
MODULE := test_pcie_host
SIM_BUILD := sim_build

EXTRA_ARGS += -Wno-fatal -I$(FIXTURE) -I$(FIXTURE)/src --timing

include $(shell cocotb-config --makefiles)/Makefile.sim
