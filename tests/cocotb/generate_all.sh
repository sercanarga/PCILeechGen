#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/../.."
GEN=./bin/pcileechgen
[ -x "$GEN" ] || { echo "build first: go build -o $GEN ./cmd/pcileechgen"; exit 1; }
for name in nvme audio xhci multibar ethernet wifi sata gpu thunderbolt generic; do
  fixture="testdata/donors/$name.json"
  board="PCIeSquirrel"
  [ "$name" = "nvme" ] && board="ac701_ft601"
  echo -n "  $name ($board)... "
  rm -rf "tests/cocotb/out_$name"
  "$GEN" build --from-json "$fixture" --board "$board" --skip-vivado \
    --output "tests/cocotb/out_$name" --force >/dev/null 2>&1 \
    && echo "OK" || echo "FAIL"
done
# default out = nvme (backwards compat)
rm -rf tests/cocotb/out
cp -r tests/cocotb/out_nvme tests/cocotb/out
echo "  default -> out_nvme copied to out/"
