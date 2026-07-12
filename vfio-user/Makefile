CC ?= gcc
CFLAGS = -Wall -O2 -g -I/usr/local/include
LDFLAGS = -L/usr/local/lib -lvfio-user -ljson-c -lpthread
UNAME_S := $(shell uname -s)
SANITIZE :=
ifeq ($(UNAME_S),Linux)
SANITIZE := -fsanitize=address,undefined
endif

TARGET = nvme_vfio_server
SRC = nvme_vfio_server.c

.PHONY: test test-matrix test-device-model

test: test-matrix test-device-model

test-matrix:
	python3 -m unittest tests/test_matrix.py -v

build/test_device_model: tests/test_device_model.c src/device_model.c include/device_model.h
	mkdir -p build
	$(CC) -std=c11 -Wall -Wextra -Werror -g $(SANITIZE) \
		-Iinclude $$(pkg-config --cflags json-c openssl cmocka) \
		-o $@ tests/test_device_model.c src/device_model.c \
		$$(pkg-config --libs json-c openssl cmocka)

test-device-model: build/test_device_model
	./build/test_device_model

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
