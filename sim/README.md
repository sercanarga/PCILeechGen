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
| `ahci_engine.sv` | SATA/AHCI had no command engine — PxCI never cleared, so storahci times out (Code 10 "I/O adapter hardware error"). Implements the slot-0 command FSM: IDENTIFY + READ/WRITE DMA over a sector store, D2H FIS, PxIS/intr, PxCI clear. |
| `xhci_ring_engine.sv` | xHCI had no Command/Event Ring engine, so usbxhci.sys hangs waiting for command completions. Implements the Command Ring walk (CRCR-latched dequeue pointer, Link TRB following, cycle-bit consumer state) answering No-Op Command (type 23) and Enable Slot Command (type 9) — plus a generic success fallback for anything else — with Command Completion Events (type 33) posted through the single-segment Event Ring (ERST fetch, producer cycle state), IMAN.IP/interrupt. |

## Integrating into firmware

These are reference modules. Wiring them into the generated SV (e.g. routing
`cfg_w1c_shadow` into the config-space shadow write path, instantiating
`nvme_store` in `pcileech_nvme_admin_responder.sv`) is the next step and must be
validated with Vivado synthesis + on-hardware testing.
