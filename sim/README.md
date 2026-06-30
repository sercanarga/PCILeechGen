# sim/ — behavioral RTL unit simulations

Self-contained Icarus Verilog testbenches for the behavioral modules that
address the emulation-gap audit (config-space W1C, MSI-X PBA, NVMe backing
store). These prove the *logic* of each module in simulation; they are not the
full PCILeech datapath (that needs the `pcileech-fpga` submodule + Vivado).

## Run

    make sim                 # all modules
    ./sim/run.sh             # all modules
    ./sim/run.sh cfg_w1c_shadow   # one module

Requires `iverilog` (`apt-get install iverilog`). Each `<name>_tb.sv` is
compiled with its DUT `<name>.sv`; a module passes only when its testbench
prints `ALL TESTS PASSED`. The harness exits non-zero on any failure and the
red path is verified (a wrong DUT makes the testbench report `FAIL`).

## Modules

| DUT | Audit gap it closes |
|-----|---------------------|
| `cfg_w1c_shadow.sv` | config-space registers are write-1-to-clear (shadow BRAM stored writes verbatim → detectable). Driven by `pcileech_cfgspace_w1cmask.coe`. |
| `msix_pba.sv` | MSI-X PBA was a static zero array; pending bits now set on masked requests and deliver on unmask. |
| `nvme_store.sv` | NVMe IO reads returned zeros and writes were discarded (Event 11); this is the BRAM sector cache the responder writes through. |

## Integrating into firmware

These are reference modules. Wiring them into the generated SV (e.g. routing
`cfg_w1c_shadow` into the config-space shadow write path, instantiating
`nvme_store` in `pcileech_nvme_admin_responder.sv`) is the next step and must be
validated with Vivado synthesis + on-hardware testing.
