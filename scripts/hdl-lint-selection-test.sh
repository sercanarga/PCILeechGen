#!/usr/bin/env bash
# verify HDL lint selects generated SystemVerilog from the build manifest only
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SELECTOR="$REPO_ROOT/scripts/manifest_sv_files.py"

command -v python3 >/dev/null 2>&1 || { echo "python3 missing"; exit 1; }
[ -f "$SELECTOR" ] || { echo "manifest SV selector missing: $SELECTOR"; exit 1; }

TMP_ROOT=/tmp
TMP="$(mktemp -d "$TMP_ROOT/pcileech-hdl-selection.XXXXXX")" || {
  echo "cannot create HDL selection-test temporary directory" >&2
  exit 1
}
case "$TMP" in
  "$TMP_ROOT"/pcileech-hdl-selection.*) ;;
  *)
    echo "refusing unexpected HDL selection-test temporary directory: $TMP" >&2
    exit 1
    ;;
esac
if [ ! -d "$TMP" ] || [ -L "$TMP" ]; then
  echo "HDL selection-test temporary directory is unsafe: $TMP" >&2
  exit 1
fi
cleanup_tmp() {
  if [ ! -d "$TMP" ] || [ -L "$TMP" ]; then
    return 0
  fi
  rm -rf -- "$TMP"
}
trap cleanup_tmp EXIT
OUT="$TMP/output"
mkdir -p "$OUT/src"
OUT="$(cd "$OUT" && pwd -P)"

printf 'module device_config; endmodule\n' > "$OUT/device_config.sv"
printf 'module pcileech_tlps128_bar_controller; endmodule\n' \
  > "$OUT/pcileech_tlps128_bar_controller.sv"
printf 'module board_top; endmodule\n' > "$OUT/src/board_top.sv"
printf 'module stale_unmanifested; endmodule\n' > "$OUT/stale_unmanifested.sv"
printf 'not HDL\n' > "$OUT/notes.txt"

cat > "$OUT/build_manifest.json" <<'EOF'
{
  "files": [
    {"name": "src/board_top.sv"},
    {"name": "pcileech_tlps128_bar_controller.sv"},
    {"name": "notes.txt"},
    {"name": "device_config.sv"}
  ]
}
EOF

actual="$(python3 "$SELECTOR" "$OUT/build_manifest.json" "$OUT")"
expected="$(printf '%s\n' \
  "$OUT/device_config.sv" \
  "$OUT/pcileech_tlps128_bar_controller.sv")"

if [ "$actual" != "$expected" ]; then
  echo "HDL lint selection mismatch" >&2
  printf 'expected:\n%s\nactual:\n%s\n' "$expected" "$actual" >&2
  exit 1
fi

case "$actual" in
  *src/board_top.sv*|*stale_unmanifested.sv*)
    echo "HDL lint selected a nested or unmanifested SV file" >&2
    exit 1
    ;;
esac

echo "HDL lint manifest selection passed"
