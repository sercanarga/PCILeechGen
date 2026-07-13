SIM ?= verilator
TOPLEVEL_LANG ?= verilog

COCOTB_TEST_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
FIXTURE ?= $(COCOTB_TEST_DIR)/out
FIXTURE_DIR := $(abspath $(FIXTURE))
TB_TOP ?= tb_top
COCOTB_TEST_MODULES ?= test_pcie_host
SIM_BUILD ?= $(CURDIR)/sim_build
COCOTB_RESULTS_FILE ?= $(CURDIR)/results.xml

RTL := $(wildcard $(FIXTURE_DIR)/*.sv)
SRC_NEEDED := pcileech_bar_impl_none.sv pcileech_bar_impl_zerowrite4k.sv pcileech_tlps128_bar_wrengine.sv pcileech_tlps128_cfgspace_shadow.sv
RTL += $(foreach f,$(SRC_NEEDED),$(wildcard $(FIXTURE_DIR)/src/$(f)))

VERILOG_SOURCES := $(RTL) $(COCOTB_TEST_DIR)/cocotb_stubs.sv $(COCOTB_TEST_DIR)/$(TB_TOP).sv
TOPLEVEL := tb_top
EXTRA_ARGS += -Wno-fatal -I$(FIXTURE_DIR) -I$(FIXTURE_DIR)/src --timing
override export PYTHONPATH := $(COCOTB_TEST_DIR)

ifndef COCOTB_MAKEFILES
$(error COCOTB_MAKEFILES must point to the active virtualenv's cocotb makefiles)
endif

.DEFAULT_GOAL := sim
.PHONY: prepare-config

prepare-config:
	@mkdir -p "$(SIM_BUILD)"
	@set -e; found=0; \
	for file in "$(FIXTURE_DIR)"/*.hex; do \
		[ -f "$$file" ] || continue; \
		cp "$$file" "$(SIM_BUILD)/"; \
		cp "$$file" "$(CURDIR)/"; \
		found=1; \
	done; \
	[ "$$found" -eq 1 ] || { echo "fixture contains no hex initialization files" >&2; exit 1; }

sim: prepare-config

include $(COCOTB_MAKEFILES)/Makefile.sim
