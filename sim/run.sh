#!/usr/bin/env bash
set -uo pipefail
cd "$(dirname "$0")"
mkdir -p build

command -v iverilog >/dev/null 2>&1 || { echo "iverilog not installed (apt-get install iverilog)"; exit 127; }

if [ "$#" -gt 0 ]; then
  tbs=()
  for n in "$@"; do tbs+=("${n%_tb.sv}_tb.sv"); done
else
  tbs=(*_tb.sv)
fi

fail=0
for tb in "${tbs[@]}"; do
  base="${tb%_tb.sv}"
  dut="${base}.sv"
  if [ ! -f "$tb" ]; then echo "[MISS] testbench $tb not found"; fail=1; continue; fi
  if [ ! -f "$dut" ]; then echo "[MISS] DUT $dut not found"; fail=1; continue; fi
  if ! iverilog -g2012 -Wall -o "build/${base}.vvp" "$dut" "$tb" 2>"build/${base}.log"; then
    echo "[FAIL] $base (compile)"; cat "build/${base}.log"; fail=1; continue
  fi
  out="$(vvp "build/${base}.vvp" 2>&1)"
  if echo "$out" | grep -q "ALL TESTS PASSED"; then
    echo "[PASS] $base"
  else
    echo "[FAIL] $base"; echo "$out"; fail=1
  fi
done

if [ "$fail" -eq 0 ]; then echo "== sim: all modules passed =="; else echo "== sim: FAILURES =="; fi
exit $fail
