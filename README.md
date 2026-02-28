# PCILeechGen

[![Go Report Card](https://goreportcard.com/badge/github.com/sercanarga/pcileechgen)](https://goreportcard.com/report/github.com/sercanarga/pcileechgen)
[![License: CC0-1.0](https://img.shields.io/badge/License-CC0_1.0-lightgrey.svg)](https://github.com/sercanarga/PCILeechGen/blob/main/LICENSE)
[![Go](https://img.shields.io/github/go-mod/go-version/sercanarga/PCILeechGen)](https://go.dev/)

Firmware generator for PCILeech-compatible FPGA boards. Reads a real PCI/PCIe device over VFIO and produces ready-to-build Vivado projects that replicate the donor device's identity on an FPGA DMA card.

## Features

- [x] Vendor / Device / Revision ID
- [x] Subsystem Vendor / Device ID
- [x] Class Code (base, sub-class, interface)
- [x] Device Serial Number (64-bit DSN)
- [x] BAR0 Layout (type, size, 32/64-bit)
- [x] Link Speed / Width (clamped to board)
- [x] Config Space (full 4KB shadow + scrubbing)
- [x] Write Mask (per-register)
- [x] Power Management (D-state)

## Supported Boards

| Board | FPGA | Lanes | Form Factor |
|---|---|---|---|
| [CaptainDMA_M2_x1](https://github.com/ufrisk/pcileech-fpga/tree/master/CaptainDMA) | XC7A35T-325 | x1 | M.2 |
| [CaptainDMA_M2_x4](https://github.com/ufrisk/pcileech-fpga/tree/master/CaptainDMA) | XC7A35T-325 | x4 | M.2 |
| [CaptainDMA_35T](https://github.com/ufrisk/pcileech-fpga/tree/master/CaptainDMA) | XC7A35T-484 | x1 | PCIe |
| [CaptainDMA_75T](https://github.com/ufrisk/pcileech-fpga/tree/master/CaptainDMA) | XC7A75T-484 | x1 | PCIe |
| [CaptainDMA_100T](https://github.com/ufrisk/pcileech-fpga/tree/master/CaptainDMA) | XC7A100T-484 | x1 | PCIe |
| [ScreamerM2](https://github.com/ufrisk/pcileech-fpga/tree/master/ScreamerM2) | XC7A35T-325 | x1 | M.2 |
| [pciescreamer](https://github.com/ufrisk/pcileech-fpga/tree/master/pciescreamer) | XC7A35T-484 | x1 | PCIe |
| [NeTV2_35T](https://github.com/ufrisk/pcileech-fpga/tree/master/NeTV2) | XC7A35T-484 | x1 | M.2 |
| [NeTV2_100T](https://github.com/ufrisk/pcileech-fpga/tree/master/NeTV2) | XC7A100T-484 | x1 | M.2 |
| [PCIeSquirrel](https://github.com/ufrisk/pcileech-fpga/tree/master/PCIeSquirrel) | XC7A35T-484 | x1 | PCIe |
| [EnigmaX1](https://github.com/ufrisk/pcileech-fpga/tree/master/EnigmaX1) | XC7A75T-484 | x1 | M.2 |
| [ZDMA](https://github.com/ufrisk/pcileech-fpga/tree/master/ZDMA) | XC7A100T-484 | x4 | PCIe |
| [GBOX](https://github.com/ufrisk/pcileech-fpga/tree/master/GBOX) | XC7A35T-484 | x1 | Mini PCIe |
| [ac701_ft601](https://github.com/ufrisk/pcileech-fpga/tree/master/ac701_ft601) | XC7A200T-676 | x4 | Dev Board |
| [acorn](https://github.com/ufrisk/pcileech-fpga/tree/master/acorn_ft2232h) | XC7A200T-484 | x4 | M.2 |
| [litefury](https://github.com/ufrisk/pcileech-fpga/tree/master/ZDMA) | XC7A100T-484 | x4 | M.2 |
| [sp605_ft601](https://github.com/ufrisk/pcileech-fpga/tree/master/sp605_ft601) | XC6SLX45T-484 | x1 | Dev Board |

## Quick Start

```bash
# install
git clone --recurse-submodules https://github.com/sercanarga/PCILeechGen.git
cd PCILeechGen && make build

# scan devices
sudo ./bin/pcileechgen scan

# build firmware
sudo ./bin/pcileechgen build --bdf 0000:03:00.0 --board CaptainDMA_100T
```

**Requirements:** Go 1.25+, Linux with IOMMU/VFIO, Vivado 2023.2+ (for synthesis)

## Commands

### `scan`
List PCI devices with VFIO compatibility status.
```bash
sudo ./bin/pcileechgen scan
```
```
0000:00:00.0 Host bridge [0600]: Intel Corporation Xeon E3-1200 v6/7th Gen Core Processor Host Bridge [8086:591f]
0000:00:17.0 SATA controller [0106]: Intel Corporation 200 Series PCH SATA controller [8086:a282]
0000:01:00.0 VGA compatible controller [0300]: NVIDIA Corporation GP106 [GeForce GTX 1060 3GB] [10de:1c02] [WARN] group(3)
0000:02:00.0 Audio device [0403]: Creative Labs CA0132 Sound Core3D [1102:0012] [OK] vfio

Total: 16 devices
```

### `check`
Verify a device is suitable as a donor.
```bash
sudo ./bin/pcileechgen check --bdf 0000:02:00.0
```
```
Checking device 0000:02:00.0...

[OK] Device found: 1102:0012 Audio device
[OK] Config space readable: 4096 bytes
[OK] IOMMU is enabled
[OK] VFIO modules loaded
[OK] IOMMU group: 9
[OK] Already bound to vfio-pci

Capabilities (3):
  [01] Power Management at offset 0x50
  [05] MSI at offset 0x60
  [10] PCI Express at offset 0x70

Extended Capabilities (4):
  [0001] Advanced Error Reporting at offset 0x100
  [0003] Device Serial Number at offset 0x150
  [0010] Single Root I/O Virtualization at offset 0x180
  [0002] VC at offset 0x260

BARs:
  BAR0: Memory32 at 0xfe800000, size 16 KiB

--- Board Compatibility ---
Donor Link: 2.5 GT/s x1
Donor DSN:  0x0123456789abcdef

Compatible boards:
  PCIeSquirrel           xc7a35tfgg484-2 x1 (exact match)
  CaptainDMA_100T        xc7a100tfgg484-2 x1 (exact match)
  ...

Total: 17 boards

--- Check complete ---
```

### `build`
Generate firmware and optionally run Vivado synthesis.

```bash
# full build
sudo ./bin/pcileechgen build --bdf 0000:03:00.0 --board CaptainDMA_100T

# artifacts only (no Vivado)
sudo ./bin/pcileechgen build --bdf 0000:03:00.0 --board CaptainDMA_100T --skip-vivado

# offline build from saved JSON
sudo ./bin/pcileechgen build --from-json device_context.json --board CaptainDMA_100T --skip-vivado
```

| Flag | Default | Description |
|---|---|---|
| `--bdf` | | Donor device BDF address |
| `--board` | | Target FPGA board (required) |
| `--from-json` | | Load donor data from JSON (offline build) |
| `--output` | `pcileech_datastore` | Output directory |
| `--lib-dir` | `lib/pcileech-fpga` | Path to pcileech-fpga library |
| `--skip-vivado` | `false` | Only generate artifacts, skip synthesis |
| `--vivado-path` | auto-detect | Path to Vivado installation |
| `--jobs` | `4` | Parallel Vivado jobs |
| `--timeout` | `3600` | Vivado timeout (seconds) |

### `validate`
Verify generated artifacts match the donor device context.
```bash
./bin/pcileechgen validate --json device_context.json --output-dir pcileech_datastore/
```

### `boards`
List all supported FPGA boards.
```bash
./bin/pcileechgen boards
```
```
NAME              FPGA PART          PCIe  TOP MODULE
----              ---------          ----  ----------
PCIeSquirrel      xc7a35tfgg484-2    x1    pcileech_squirrel_top
ScreamerM2        xc7a35tcsg325-2    x1    pcileech_screamer_m2_top
CaptainDMA_100T   xc7a100tfgg484-2   x1    pcileech_100t484_x1_top
ZDMA              xc7a100tfgg484-2   x4    pcileech_tbx4_100t_top
...

Total: 17 boards
```

## Output

```
pcileech_datastore/
  device_context.json
  pcileech_cfgspace.coe
  pcileech_cfgspace_writemask.coe
  pcileech_bar_zero4k.coe
  vivado_generate_project.tcl
  vivado_build.tcl
  src/                          # patched SV sources
  *.bin                         # bitstream (after Vivado build)
```

## Development

```bash
make test
make test-coverage
make lint
```

## Special Thanks

- [pcileech-fpga](https://github.com/ufrisk/pcileech-fpga) by Ulf Frisk - the FPGA framework this project builds upon
- [CaptainDMA](https://captaindma.com) - For best FPGA DMA hardware

## License

[Creative Commons Zero v1.0 Universal](https://github.com/sercanarga/PCILeechGen/blob/main/LICENSE)

## Legal Notice

This tool is provided for **educational and research purposes only**. The authors do not condone or encourage the use of this software for cheating, circumventing anti-cheat systems, or any other activity that violates terms of service of any software or platform. Users are solely responsible for ensuring their use of this tool complies with all applicable laws and agreements.