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
bars=0
if [ -f "$found/resource" ]; then
    bars="$(grep -vc '^0* 0* 0*$' "$found/resource")"
fi
status=pass
detail=driver-bound
case "$case_name" in
    nvme|sata|xhci)
        if [ "$driver" = none ]; then
            status=fail
            detail=required-driver-not-bound
        fi
        ;;
    audio|ethernet|wifi|gpu|thunderbolt)
        if [ "$driver" = none ]; then
            status=skip
            detail=optional-driver-not-bound
        fi
        ;;
    generic|multibar)
        detail=enumeration
        ;;
esac
printf '{"event":"result","case":"%s","status":"%s","bdf":"%s","vendor":"%s","device":"%s","class":"%s","driver":"%s","bars":%s,"detail":"%s"}\n' \
    "$case_name" "$status" "$bdf" "$vendor" "$device" "$class" "$driver" "$bars" "$detail"
[ "$status" != fail ]
