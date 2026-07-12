FROM ubuntu:22.04

ARG QEMU_VERSION=10.1.1

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential ca-certificates curl git libcmocka-dev libglib2.0-dev \
    libjson-c-dev libpixman-1-dev libssl-dev linux-image-generic meson \
    ninja-build pkg-config python3 python3-pip python3-venv xz-utils \
    busybox-static cpio gzip \
    && rm -rf /var/lib/apt/lists/*

RUN git clone --depth 1 https://github.com/nutanix/libvfio-user.git /tmp/libvfio-user \
    && meson setup /tmp/libvfio-user/build /tmp/libvfio-user --prefix=/usr/local \
    && ninja -C /tmp/libvfio-user/build install \
    && rm -rf /tmp/libvfio-user

RUN pip3 install --no-cache-dir "meson>=1.8,<1.9" tomli \
    && curl -fsSL "https://download.qemu.org/qemu-${QEMU_VERSION}.tar.xz" \
       | tar -xJ -C /tmp \
    && case "$(dpkg --print-architecture)" in \
         arm64) target=aarch64-softmmu ;; \
         amd64) target=x86_64-softmmu ;; \
         *) exit 1 ;; \
       esac \
    && cd "/tmp/qemu-${QEMU_VERSION}" \
    && ./configure --prefix=/usr/local --target-list="$target" \
         --disable-docs --disable-gtk --disable-sdl --disable-tools \
    && ninja -C build -j"$(nproc)" install \
    && rm -rf "/tmp/qemu-${QEMU_VERSION}"

COPY vfio-user /build
RUN make -C /build all

COPY vfio-user/guest /guest-source
RUN set -eu; \
    root=/guest/root; \
    kernel="$(basename "$(ls /boot/vmlinuz-* | sort | tail -1)")"; \
    version="${kernel#vmlinuz-}"; \
    mkdir -p "$root/bin" "$root/sbin" "$root/proc" "$root/sys" "$root/dev" "$root/run" "$root/lib/modules/$version"; \
    cp /bin/busybox "$root/bin/busybox"; \
    for applet in sh mount sleep poweroff cat grep cut tr basename dirname readlink; do ln -s busybox "$root/bin/$applet"; done; \
    ln -s ../bin/busybox "$root/sbin/modprobe"; \
    cp /guest-source/init "$root/init"; \
    cp /guest-source/probe.sh "$root/bin/vfio-probe"; \
    chmod +x "$root/init" "$root/bin/vfio-probe"; \
    cp "/lib/modules/$version/modules."* "$root/lib/modules/$version/"; \
    for module in nvme ahci xhci_pci snd_hda_intel e1000e iwlwifi nouveau thunderbolt; do \
      modprobe --show-depends -S "$version" "$module" 2>/dev/null \
        | awk '/^insmod / {print $2}' \
        | while read -r path; do mkdir -p "$root$(dirname "$path")"; cp "$path" "$root$path"; done; \
    done; \
    cd "$root"; \
    find . -print0 | cpio --null -o -H newc | gzip -9 > /guest/initramfs.cpio.gz; \
    cp "/boot/$kernel" /guest/kernel; \
    case "$(dpkg --print-architecture)" in arm64) echo aarch64 > /guest/arch ;; amd64) echo x86_64 > /guest/arch ;; esac

WORKDIR /build
