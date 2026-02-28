# PCILeechGen

[![Go Report Card](https://goreportcard.com/badge/github.com/sercanarga/pcileechgen)](https://goreportcard.com/report/github.com/sercanarga/pcileechgen)
[![License: CC0-1.0](https://img.shields.io/badge/License-CC0_1.0-lightgrey.svg)](https://github.com/sercanarga/PCILeechGen/blob/main/LICENSE)
[![Go](https://img.shields.io/github/go-mod/go-version/sercanarga/PCILeechGen)](https://go.dev/)

Firmware generator for PCILeech-compatible FPGA boards. Reads a real PCI/PCIe device over VFIO and produces ready-to-build Vivado projects that replicate the donor device's identity on an FPGA DMA card.

## What it does

1. **Scan** the host for PCI devices
2. **Read** a donor device's full config space, BARs, capabilities, and serial number
3. **Generate** Xilinx COE files, patched SystemVerilog sources, and Vivado TCL scripts
4. **Build** a bitstream (requires Vivado)

## Features

| Donor Property | Emulated |
|---|---:|
| Vendor / Device / Revision ID | ✅ |
| Subsystem Vendor / Device ID | ✅ |
| Class Code (base, sub-class, interface) | ✅ |
| Device Serial Number (64-bit DSN) | ✅ |
| BAR0 Layout (type, size, 32/64-bit) | ✅ |
| Link Speed / Width (clamped to board) | ✅ |
| Config Space (full 4KB shadow + scrubbing) | ✅ |
| Write Mask (per-register) | ✅ |
| Power Management (D-state) | ✅ |

## Requirements

- Go 1.25+
- Linux with IOMMU enabled and VFIO support
- A donor PCIe device
- Xilinx Vivado 2023.2+ (optional, for synthesis)

## Install

```bash
git clone --recurse-submodules https://github.com/sercanarga/PCILeechGen.git
cd PCILeechGen
make build
```

## Usage

### `scan` — list PCI devices

```bash
sudo ./bin/pcileechgen scan
```

### `check` — verify donor compatibility

```bash
sudo ./bin/pcileechgen check --bdf 0000:03:00.0
```

| Flag | Description |
|---|---|
| `--bdf` | Device BDF address (required) |

### `build` — generate firmware

Full build (collect + generate + synthesize):
```bash
sudo ./bin/pcileechgen build --bdf 0000:03:00.0 --board CaptainDMA_100T
```

Generate artifacts only (no Vivado):
```bash
sudo ./bin/pcileechgen build --bdf 0000:03:00.0 --board CaptainDMA_100T --skip-vivado
```

Offline build from saved JSON:
```bash
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

### `validate` — verify generated artifacts

```bash
./bin/pcileechgen validate --json device_context.json --output-dir pcileech_datastore/
```

| Flag | Default | Description |
|---|---|---|
| `--json` | | Path to device_context.json (required) |
| `--output-dir` | `.` | Path to firmware output directory |

### `boards` — list supported boards

```bash
./bin/pcileechgen boards
```

## Supported boards

- **CaptainDMA**
  - `CaptainDMA_M2_x1` — XC7A35T-325, x1
  - `CaptainDMA_M2_x4` — XC7A35T-325, x4
  - `CaptainDMA_35T` — XC7A35T-484, x1
  - `CaptainDMA_75T` — XC7A75T-484, x1
  - `CaptainDMA_100T` — XC7A100T-484, x1
- **Screamer**
  - `ScreamerM2` — XC7A35T-325, x1
  - `pciescreamer` — XC7A35T-484, x1
- **NeTV2**
  - `NeTV2_35T` — XC7A35T-484, x1
  - `NeTV2_100T` — XC7A100T-484, x1
- **Other**
  - `PCIeSquirrel` — XC7A35T-484, x1
  - `EnigmaX1` — XC7A75T-484, x1
  - `ZDMA` — XC7A100T-484, x4
  - `GBOX` — XC7A35T-484, x1
  - `ac701_ft601` — XC7A200T-676, x4
  - `acorn` — XC7A200T-484, x4
  - `litefury` — XC7A100T-484, x4
  - `sp605_ft601` — XC6SLX45T-484, x1

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