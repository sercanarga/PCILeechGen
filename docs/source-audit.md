# External Source Audit

This audit records which ideas from related public projects are suitable for
PCILeechGen and which are intentionally out of scope. External repositories are
reference material only; code is not copied without a license review and a local
test that proves the adapted behavior.

## Safety Boundary

PCILeechGen supports educational and research firmware generation, donor
profiling, artifact validation, board metadata, and driver-compatibility
diagnostics. It must not add features whose purpose is anti-cheat bypass,
unauthorized memory access, activation locks, game cheating, detection evasion,
or fake activity generation for bypass claims.

## Adapted Sources

| Source | Useful idea | Decision | Local surface |
| --- | --- | --- | --- |
| `ret2c/MMIO2Verilog` | Text MMIO traces can seed BAR behavior models. | Adapt trace format support; do not copy generated Verilog. | `internal/donor/mmio`, `pcileechgen mmio-trace --format mmio2verilog` |
| `Shocka-Zulu/Bar2Verilog` | BAR binary dumps can produce static read maps. | Keep as future importer input; generate through `BARModel`, not raw pasted HDL. | `internal/firmware/barmodel` |
| `Ap3x/PCIeConfigSpace` and `Simonrak/writemask.it` | Config-space dumps and write masks are useful validation inputs. | Future work: typed importers for sanitized dump formats. | `internal/pci`, `internal/firmware/codegen` |
| `Moer2831/PCILeechFWGenerator` | Guided workflows, board checks, manifests, and validation improve UX. | Already aligned; adapt only safe UX/validation ideas. | `cmd/pcileechgen`, `internal/firmware/output` |
| `Shocka-Zulu/DevicePopulation` | Population data can guide donor class/profile prioritization. | Use only aggregate class/profile prioritization, not identity spoofing claims. | `internal/firmware/devclass` |
| `ufrisk/pcileech` and `ufrisk/pcileech-fpga` | Upstream protocol and FPGA framework behavior. | Treat as primary compatibility reference. | `lib/pcileech-fpga`, generated SV/TCL |
| `Moer2831/pcileech-nvme`, `16SalomonArs/Pcileech-DMA-NVMe-VMD`, `rick-heig/eNVMe` | NVMe register, queue, identify, and diagnostics behavior. | Adapt spec-faithful NVMe behavior only. Exclude VMD/bypass claims. | `internal/firmware/nvme`, `internal/firmware/svgen` |
| `ret2c/pcileech-rt5392`, `ekknod/pcileech-wifi`, Wi-Fi forks | Device-class BAR and interrupt examples. | Adapt class-profile patterns only when backed by public driver/spec behavior. | `internal/firmware/devclass`, `internal/firmware/barmodel` |
| `noy00y/PCILeech-Litefury`, board/TCL forks | Additional board metadata and build scripts. | Add boards only after source path, top module, part, lane count, and BRAM size validate. | `internal/board/boards.json`, `pcileechgen boards --json` |

## Excluded Sources

Repositories or documents framed around anti-cheat bypass, game cheating,
activation cracking, fake device activity for evasion, or unauthorized memory
operations are not implementation sources. They may be listed in the audit only
to explain exclusion and prevent accidental feature creep.

Examples from the supplied list that require exclusion or heavy filtering:

- `Pcileech-Intel-I226-V-FullEmu`
- `Pcileech-DMA-Unblocker`
- `VGK-DMA-BYPASS`
- `DMA-Activator*`
- `mcp_server_pcileech`
- anti-cheat detector/bypass repositories

## Acceptance Rule

Every future adaptation from this audit needs:

1. A source entry in this file.
2. A license/safety decision before implementation.
3. A Go test or HDL lint fixture proving the local behavior.
4. Generated output through PCILeechGen templates and models, not copied HDL.
