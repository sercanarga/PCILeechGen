# Build Diagnostics

## Better build failures

Build errors are classified before they reach the CLI. Vivado output now writes
summary files next to the generated TCL script, for example:

```text
vivado_generate_project_summary.txt
vivado_build_summary.txt
build_summary.txt
```

Common causes are called out directly:

- Vivado RAM/OOM or terminated process
- license failure
- corrupted or incomplete IP cache
- timing failure
- donor BAR larger than board BRAM
- generic subprocess `exit status 1`

## Structured diagnosis

Run this after a failed or suspicious build:

```text
pcileechgen diagnose --json testdata/donors/nvme.json --output-dir pcileech_datastore --board CaptainDMA_75T
```

The command checks:

- config-space presence
- BAR sizes and preferred class BAR
- MSI-X table and PBA placement
- NVMe class code, identity and namespace metadata
- Ethernet class code and MSI-X vector count
- saved build and Vivado summaries

## VFIO matrix

Default matrix:

```text
python vfio-user/matrix.py --all
```

Default matrix plus synthetic demo profiles:

```text
python vfio-user/matrix.py --all --include-demo
```

Single demo profile:

```text
python vfio-user/matrix.py --case rtl8125
python vfio-user/matrix.py --case nvmev2
```
