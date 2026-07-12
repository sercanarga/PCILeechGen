FROM ubuntu:22.04

RUN apt-get update && apt-get install -y \
    build-essential cmake git meson ninja-build pkg-config \
    libjson-c-dev libcmocka-dev libssl-dev python3 \
    qemu-system-x86 qemu-utils \
    linux-image-generic busybox-static cpio gzip \
    && rm -rf /var/lib/apt/lists/*

# Build libvfio-user
RUN git clone --depth 1 https://github.com/nutanix/libvfio-user.git /tmp/libvfio-user && \
    cd /tmp/libvfio-user && \
    meson setup build --prefix=/usr/local && \
    cd build && ninja && meson install && \
    rm -rf /tmp/libvfio-user

COPY vfio-user /build
COPY tests/cocotb/out_generic /artifacts/generic
RUN cd /build && make all

# Minimal guest initramfs
RUN mkdir -p /guest && \
    echo '#!/bin/sh' > /guest/init && \
    echo 'mount -t proc proc /proc' >> /guest/init && \
    echo 'mount -t sysfs sys /sys' >> /guest/init && \
    echo 'mount -t devtmpfs dev /dev' >> /guest/init && \
    echo 'echo "=== PCI devices ==="' >> /guest/init && \
    echo 'lspci -vv 2>/dev/null || cat /proc/bus/pci/devices' >> /guest/init && \
    echo 'echo "=== NVMe ==="' >> /guest/init && \
    echo 'sleep 2 && nvme list 2>&1 || echo "nvme not found"' >> /guest/init && \
    echo 'echo "=== dmesg nvme ==="' >> /guest/init && \
    echo 'dmesg | grep -i nvme' >> /guest/init && \
    echo 'echo "=== dmesg pci ==="' >> /guest/init && \
    echo 'dmesg | grep -i pci' >> /guest/init && \
    echo 'poweroff -f' >> /guest/init && \
    chmod +x /guest/init && \
    cd /guest && \
    find . | cpio -o -H newc | gzip > /guest/initramfs.cpio.gz

WORKDIR /build
