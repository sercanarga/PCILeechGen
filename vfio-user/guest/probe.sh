#!/bin/sh
set -eu

argument() {
    tr ' ' '\n' </proc/cmdline | grep "^$1=" | cut -d= -f2
}

case_name="$(argument vfio_case)"
vendor="$(argument vfio_vendor)"
device="$(argument vfio_device)"
rebind="$(argument vfio_rebind 2>/dev/null || true)"
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

if [ "$rebind" = 1 ] && [ "$driver" != none ]; then
    driver_path="/sys/bus/pci/drivers/$driver"
    if [ ! -w "$driver_path/unbind" ] || [ ! -w "$driver_path/bind" ]; then
        printf '{"event":"result","case":"%s","status":"fail","bdf":"%s","vendor":"%s","device":"%s","class":"%s","driver":"%s","detail":"rebind-unavailable"}\n' \
            "$case_name" "$(basename "$found")" "$vendor" "$device" "$class" "$driver"
        exit 1
    fi
    printf '%s' "$(basename "$found")" >"$driver_path/unbind"
    if [ -w "$found/reset" ]; then
        printf 1 >"$found/reset"
    else
        printf '{"event":"result","case":"%s","status":"fail","bdf":"%s","vendor":"%s","device":"%s","class":"%s","driver":"%s","detail":"reset-unavailable"}\n' \
            "$case_name" "$(basename "$found")" "$vendor" "$device" "$class" "$driver"
        exit 1
    fi
    printf '%s' "$(basename "$found")" >"$driver_path/bind"
    driver=none
    if [ -L "$found/driver" ]; then
        driver="$(basename "$(readlink "$found/driver")")"
    fi
    if [ "$driver" = none ]; then
        printf '{"event":"result","case":"%s","status":"fail","bdf":"%s","vendor":"%s","device":"%s","class":"%s","driver":"none","detail":"rebind-driver-not-bound"}\n' \
            "$case_name" "$(basename "$found")" "$vendor" "$device" "$class"
        exit 1
    fi
fi
bars=0
if [ -f "$found/resource" ]; then
    while read -r start end flags; do
        if [ "$start" != 0 ] || [ "$end" != 0 ]; then
            bars=$((bars + 1))
        fi
    done <"$found/resource"
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
if [ "$rebind" = 1 ] && [ "$status" = pass ]; then
    detail=rebind-reset-driver-bound
fi
printf '{"event":"result","case":"%s","status":"%s","bdf":"%s","vendor":"%s","device":"%s","class":"%s","driver":"%s","bars":%s,"detail":"%s"}\n' \
    "$case_name" "$status" "$bdf" "$vendor" "$device" "$class" "$driver" "$bars" "$detail"
[ "$status" != fail ]
