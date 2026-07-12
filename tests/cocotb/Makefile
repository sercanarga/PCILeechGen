SIM ?= verilator
TOPLEVEL_LANG ?= verilog

FIXTURE ?= out
TB_TOP ?= tb_top
MODULE ?= test_pcie_host
RTL := $(wildcard $(FIXTURE)/*.sv)
SRC_NEEDED := pcileech_bar_impl_none.sv pcileech_bar_impl_zerowrite4k.sv pcileech_tlps128_bar_wrengine.sv
RTL += $(foreach f,$(SRC_NEEDED),$(wildcard $(FIXTURE)/src/$(f)))
STUBS := cocotb_stubs.sv

VERILOG_SOURCES := $(RTL) $(STUBS) $(TB_TOP).sv
TOPLEVEL := tb_top
SIM_BUILD := sim_build

EXTRA_ARGS += -Wno-fatal -I$(FIXTURE) -I$(FIXTURE)/src --timing

ifdef TIMEOUT_CYCLES
$(shell perl -pi -e 's/\.TIMEOUT_CYCLES\s*\x28\s*24'"'"'h[0-9A-Fa-f]+\s*\x29/.TIMEOUT_CYCLES \x28 24'"'"'h$(TIMEOUT_CYCLES) \x29/' $(FIXTURE)/pcileech_nvme_dma_bridge.sv 2>/dev/null)
endif

include $(shell cocotb-config --makefiles)/Makefile.sim
