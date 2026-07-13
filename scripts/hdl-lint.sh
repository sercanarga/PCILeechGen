#!/usr/bin/env bash
# lint generated SV across all donor fixtures and boards
set -uo pipefail
shopt -s nullglob

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$REPO_ROOT"

require() { command -v "$1" >/dev/null 2>&1 || { echo "missing dep: $1" >&2; exit 1; }; }
require verilator
require ./bin/pcileechgen

FIXTURES=(testdata/donors/*.json)
if [ "${#FIXTURES[@]}" -eq 0 ]; then
  echo "ERROR: no donor fixtures found under testdata/donors" >&2
  exit 1
fi
LEGACY_BOARDS=(pciescreamer NeTV2_35T NeTV2_100T acorn litefury sp605_ft601)
# Board names come from the CLI so the matrix stays in sync with board.go.
# (NR>2 skips the two-line table header; bash 3.2 compat via while-read.)
BOARDS=()
while IFS= read -r line; do BOARDS+=("$line"); done < <(
  ./bin/pcileechgen boards | awk 'NR>2 && $1!="----" && $1!="" && $1!="Total:"{print $1}'
)

if [ "${#BOARDS[@]}" -eq 0 ]; then
  echo "ERROR: no boards parsed from 'pcileechgen boards'" >&2
  exit 1
fi

# whitelist of svgen output files (primitives like pcileech_fifo.sv are blackboxed)
SVGEN_PATTERN='pcileech_lifecycle_service.sv|pcileech_dma_tag_service.sv|pcileech_interrupt_service.sv|pcileech_bar_impl_device.sv|pcileech_tlps128_bar_controller.sv|pcileech_tlp_normalizer.sv|pcileech_tlp_ur_completer.sv|pcileech_bar_rsp_arbiter.sv|pcileech_tlps128_bar_rdengine.sv|pcileech_tlps128_bar_wrengine.sv|pcileech_bar_impl_none.sv|pcileech_bar_impl_zerowrite4k.sv|pcileech_bar_impl_msi.sv|tlp_latency_emulator.sv|device_config.sv|pcileech_msix_table.sv|pcileech_nvme_admin_responder.sv|pcileech_nvme_dma_bridge.sv|pcileech_bram_disk.sv|pcileech_hda_rirb_dma.sv|pcileech_hda_msi.sv'

TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT

REPORT="${HDL_LINT_REPORT:-hdl-lint-report.tsv}"
printf 'fixture\tboard\tstatus\tdetail\n' > "$REPORT"

total=0; pass=0; skip=0; fail=0; modern_multibar_pass=0

for fixture in "${FIXTURES[@]}"; do
  class="$(basename "$fixture" .json)"
  for board in "${BOARDS[@]}"; do
    total=$((total+1))
    cell="${class}×${board}"
    out="$TMP/$class/$board"
    case " ${LEGACY_BOARDS[*]} " in
      *" $board "*)
        echo "SKIP  $cell (explicit legacy board allowlist)"
        printf '%s\t%s\tSKIP\tlegacy board allowlist\n' "$class" "$board" >> "$REPORT"
        skip=$((skip+1))
        continue
        ;;
    esac
    mkdir -p "$out"

    build_log="$out/build.log"
    if ! ./bin/pcileechgen build --from-json "$fixture" --board "$board" \
          --skip-vivado --output "$out" --force >"$build_log" 2>&1; then
      # Donor BAR > board BRAM (or other benign incompatibility): skip, do not fail CI.
      if grep -Eq 'board sources not found at' "$build_log"; then
        echo "SKIP  $cell (board source unavailable — see $build_log)"
        printf '%s\t%s\tSKIP\tboard source unavailable\n' "$class" "$board" >> "$REPORT"
        skip=$((skip+1))
      elif grep -Eq 'insufficient block RAM|exceeds board BRAM' "$build_log"; then
        echo "SKIP  $cell (build incompatible — see $build_log)"
        printf '%s\t%s\tSKIP\tbuild incompatible\n' "$class" "$board" >> "$REPORT"
        skip=$((skip+1))
      else
        echo "FAIL  $cell (build failed — see $build_log)"
        printf '%s\t%s\tFAIL\tbuild failed\n' "$class" "$board" >> "$REPORT"
        cat "$build_log" >&2
        fail=$((fail+1))
      fi
      continue
    fi

    sv_files=()
    while IFS= read -r f; do sv_files+=("$f"); done < <(
      find "$out" -name '*.sv' -print \
        | grep -E "$SVGEN_PATTERN" \
        | sort
    )

    if [ "${#sv_files[@]}" -eq 0 ]; then
      echo "SKIP  $cell (no svgen SV emitted)"
      printf '%s\t%s\tSKIP\tno generated SV\n' "$class" "$board" >> "$REPORT"
      skip=$((skip+1))
      continue
    fi

    lint_log="$out/verilator.log"
    if verilator --lint-only -Wno-fatal --top-module pcileech_tlps128_bar_controller \
          +incdir+"$out/src" \
          "${sv_files[@]}" testdata/stubs/blackbox.sv \
          >"$lint_log" 2>&1; then
      echo "PASS  $cell"
      printf '%s\t%s\tPASS\tlint passed\n' "$class" "$board" >> "$REPORT"
      pass=$((pass+1))
      if [ "$cell" = "multibar×PCIeSquirrel" ]; then
        modern_multibar_pass=1
      fi
    else
      echo "FAIL  $cell (see $lint_log)"
      printf '%s\t%s\tFAIL\tverilator failed\n' "$class" "$board" >> "$REPORT"
      cat "$lint_log" >&2
      fail=$((fail+1))
    fi
  done
done

if [ "$modern_multibar_pass" -ne 1 ]; then
  echo "FAIL  mandatory modern multibar×PCIeSquirrel cell did not pass" >&2
  fail=$((fail+1))
fi
echo "HDL lint: total=$total pass=$pass skip=$skip fail=$fail"
[ "$fail" -eq 0 ]
