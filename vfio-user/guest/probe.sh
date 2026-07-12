#!/bin/sh
set -eu

argument() {
    tr ' ' '\n' </proc/cmdline | grep "^$1=" | cut -d= -f2
}

case_name="$(argument vfio_case)"
vendor="$(argument vfio_vendor)"
device="$(argument vfio_device)"
found=

for path in /sys/bus/pci/devices/*; do
    [ -f "$path/vendor" ] || continue
    actual_vendor="$(cat "$path/vendor")"
    actual_device="$(cat "$path/device")"
    if [ "$actual_vendor" = "0x$vendor" ] && [ "$actual_device" = "0x$device" ]; then
        found="$path"
        break
    fi
done

if [ -z "$found" ]; then
    printf '{"event":"result","case":"%s","status":"fail","detail":"device-not-found"}\n' "$case_name"
    exit 1
fi

bdf="$(basename "$found")"
class="$(cat "$found/class")"
driver=none
if [ -L "$found/driver" ]; then
    driver="$(basename "$(readlink "$found/driver")")"
fi
bars="$(grep -vc '^0* 0* 0*$' "$found/resource" || true)"
printf '{"event":"result","case":"%s","status":"pass","bdf":"%s","vendor":"%s","device":"%s","class":"%s","driver":"%s","bars":%s}\n' \
    "$case_name" "$bdf" "$vendor" "$device" "$class" "$driver" "$bars"
