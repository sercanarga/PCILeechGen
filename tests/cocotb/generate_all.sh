#!/usr/bin/env bash
set -euo pipefail

usage() {
  echo "usage: $0 --generator PATH --output-root ABSOLUTE_PATH" >&2
  exit 2
}

generator=""
output_root=""
while [ "$#" -gt 0 ]; do
  case "$1" in
    --generator)
      [ "$#" -ge 2 ] || usage
      generator=$2
      shift 2
      ;;
    --output-root)
      [ "$#" -ge 2 ] || usage
      output_root=$2
      shift 2
      ;;
    *)
      usage
      ;;
  esac
done

[ -n "$generator" ] || usage
[ -n "$output_root" ] || usage
[ -x "$generator" ] || { echo "generator is not executable: $generator" >&2; exit 2; }

case "$output_root" in
  /*) ;;
  *) echo "output root must be absolute: $output_root" >&2; exit 2 ;;
esac
[ "$output_root" != "/" ] || { echo "refusing output root /" >&2; exit 2; }

output_parent=$(dirname "$output_root")
[ -d "$output_parent" ] || { echo "output parent does not exist: $output_parent" >&2; exit 2; }
output_parent=$(cd "$output_parent" && pwd -P)
output_root="$output_parent/$(basename "$output_root")"

if [ -e "$output_root" ]; then
  [ ! -L "$output_root" ] || { echo "output root must not be a symlink: $output_root" >&2; exit 2; }
  [ -d "$output_root" ] || { echo "output root is not a directory: $output_root" >&2; exit 2; }
  if [ -n "$(find "$output_root" -mindepth 1 -maxdepth 1 -print -quit)" ]; then
    echo "output root must be empty: $output_root" >&2
    exit 2
  fi
else
  mkdir "$output_root"
fi

script_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)
repo_root=$(cd "$script_dir/../.." && pwd -P)

for name in nvme audio xhci multibar ethernet wifi sata gpu thunderbolt generic; do
  fixture="$repo_root/testdata/donors/$name.json"
  board="PCIeSquirrel"
  [ "$name" = "nvme" ] && board="ac701_ft601"
  output="$output_root/$name"
  [ ! -e "$output" ] || { echo "fixture output already exists: $output" >&2; exit 2; }

  echo "generating $name ($board)"
  # The committed matrix deliberately includes real donor BARs larger than
  # the compact PCIeSquirrel BRAM. These are behavioral RTL fixtures, so use
  # the same explicit forced-generation policy as the VFIO fixture matrix.
  "$generator" build --from-json "$fixture" --board "$board" --skip-vivado --force \
    --output "$output"
  "$generator" verify-manifest --manifest "$output/build_manifest.json" \
    --output-dir "$output"
done
