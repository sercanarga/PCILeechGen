# Demo Donor Profiles

This source tree includes extra offline donor fixtures for quick generator tests.
They are not collected from real hardware.

## Profiles

- `NICv2.json` - Intel I225-V style Ethernet controller, class `0x020000`
- `RTL8125.json` - Realtek RTL8125 style Ethernet controller, class `0x020000`
- `RealtekRTL8125.json` - explicit Realtek RTL8125 fixture alias
- `IntelI210.json` - Intel I210 style Ethernet controller, class `0x020000`
- `IntelI219.json` - Intel I219 style Ethernet controller, class `0x020000`
- `IntelI225.json` - Intel I225 style Ethernet controller, class `0x020000`
- `DiskTest.json` - NVMe test controller with named model/serial metadata
- `NVMEv2.json` - NVMe v2 test controller with named model/serial metadata

## Generate

```text
pcileechgen fixtures --demo-profiles --out testdata/donors
```

## Build Example

```text
pcileechgen build --from-json testdata/donors/RTL8125.json --board CaptainDMA_75T --skip-vivado
```

## VFIO Matrix

```text
python vfio-user/matrix.py --all --include-demo
python vfio-user/matrix.py --case rtl8125
python vfio-user/matrix.py --case nvmev2
```
