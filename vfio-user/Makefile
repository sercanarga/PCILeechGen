CC ?= gcc
CFLAGS = -Wall -O2 -g -I/usr/local/include
LDFLAGS = -L/usr/local/lib -lvfio-user -ljson-c -lpthread

TARGET = nvme_vfio_server
SRC = nvme_vfio_server.c

.PHONY: test test-matrix

test: test-matrix

test-matrix:
	python3 -m unittest tests/test_matrix.py -v

$(TARGET): $(SRC)
	$(CC) $(CFLAGS) -o $(TARGET) $(SRC) $(LDFLAGS)

clean:
	rm -f $(TARGET)

run: $(TARGET)
	./$(TARGET) /tmp/nvme.sock & sleep 1 && \
	qemu-system-x86_64 \
		-M q35 -m 256 -kernel /boot/vmlinuz \
		-initrd /guest/initramfs.cpio.gz \
		-append "console=ttyS0 quiet" \
		-nographic -no-reboot \
		-device vfio-user-pci,socket=/tmp/nvme.sock \
		-serial mon:stdio
