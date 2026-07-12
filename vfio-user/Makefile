CC ?= gcc
.DEFAULT_GOAL := all
CFLAGS = -D_GNU_SOURCE -std=c11 -Wall -Wextra -Werror -O2 -g -Iinclude -I/usr/local/include
LDFLAGS = -L/usr/local/lib -lvfio-user $$(pkg-config --libs json-c openssl) -lpthread
UNAME_S := $(shell uname -s)
SANITIZE :=
ifeq ($(UNAME_S),Linux)
SANITIZE := -fsanitize=address,undefined
endif

TARGET = build/vfio-device
SRC = src/main.c src/vfio_device.c src/device_model.c src/behavior_static.c src/behavior_nvme.c src/behavior_hda.c src/behavior_ahci.c src/behavior_xhci.c src/behavior_ethernet.c

.PHONY: all image test test-matrix test-device-model test-behavior-static test-behavior-nvme test-behavior-hda test-behavior-storage test-server matrix

all: $(TARGET)

test: test-matrix test-device-model test-behavior-static test-behavior-nvme test-behavior-hda test-behavior-storage

test-matrix:
	python3 -m unittest tests/test_matrix.py -v

build/test_device_model: tests/test_device_model.c src/device_model.c include/device_model.h
	mkdir -p build
	$(CC) $(CFLAGS) $(SANITIZE) \
		$$(pkg-config --cflags json-c openssl cmocka) \
		-o $@ tests/test_device_model.c src/device_model.c \
		$$(pkg-config --libs json-c openssl cmocka)

test-device-model: build/test_device_model
	./build/test_device_model

build/test_behavior_static: tests/test_behavior_static.c src/behavior_static.c src/behavior_nvme.c src/behavior_hda.c src/behavior_ahci.c src/behavior_xhci.c src/behavior_ethernet.c include/device_behavior.h include/device_model.h
	mkdir -p build
	$(CC) $(CFLAGS) $(SANITIZE) \
		$$(pkg-config --cflags cmocka) -o $@ tests/test_behavior_static.c src/behavior_static.c src/behavior_nvme.c src/behavior_hda.c src/behavior_ahci.c src/behavior_xhci.c src/behavior_ethernet.c \
		$$(pkg-config --libs cmocka)

test-behavior-static: build/test_behavior_static
	./build/test_behavior_static

build/test_behavior_nvme: tests/test_behavior_nvme.c src/behavior_nvme.c src/behavior_static.c src/behavior_hda.c src/behavior_ahci.c src/behavior_xhci.c src/behavior_ethernet.c src/device_model.c include/device_behavior.h include/device_model.h
	mkdir -p build
	$(CC) $(CFLAGS) $(SANITIZE) \
		$$(pkg-config --cflags cmocka json-c openssl) -o $@ \
		tests/test_behavior_nvme.c src/behavior_nvme.c src/behavior_static.c src/behavior_hda.c src/behavior_ahci.c src/behavior_xhci.c src/behavior_ethernet.c src/device_model.c \
		$$(pkg-config --libs cmocka json-c openssl)

test-behavior-nvme: build/test_behavior_nvme
	./build/test_behavior_nvme

build/test_behavior_hda: tests/test_behavior_hda.c src/behavior_hda.c src/behavior_ethernet.c src/behavior_static.c src/behavior_nvme.c src/behavior_ahci.c src/behavior_xhci.c src/device_model.c include/device_behavior.h include/device_model.h
	mkdir -p build
	$(CC) $(CFLAGS) $(SANITIZE) \
		$$(pkg-config --cflags cmocka json-c openssl) -o $@ \
		tests/test_behavior_hda.c src/behavior_hda.c src/behavior_ethernet.c src/behavior_static.c src/behavior_nvme.c src/behavior_ahci.c src/behavior_xhci.c src/device_model.c \
		$$(pkg-config --libs cmocka json-c openssl)

test-behavior-hda: build/test_behavior_hda
	./build/test_behavior_hda

build/test_behavior_storage: tests/test_behavior_storage.c src/behavior_ahci.c src/behavior_xhci.c src/behavior_static.c src/behavior_nvme.c src/behavior_hda.c src/behavior_ethernet.c src/device_model.c include/device_behavior.h include/device_model.h
	mkdir -p build
	$(CC) $(CFLAGS) $(SANITIZE) \
		$$(pkg-config --cflags cmocka json-c openssl) -o $@ \
		tests/test_behavior_storage.c src/behavior_ahci.c src/behavior_xhci.c \
		src/behavior_static.c src/behavior_nvme.c src/behavior_hda.c src/behavior_ethernet.c src/device_model.c \
		$$(pkg-config --libs cmocka json-c openssl)

test-behavior-storage: build/test_behavior_storage
	./build/test_behavior_storage

image:
	docker build -t pcileech-vfio-user -f Dockerfile ..

test-server: image
	docker run --rm \
		-e VFIO_DEVICE_BIN=/build/build/vfio-device \
		-e VFIO_ARTIFACTS=/artifacts/generic \
		pcileech-vfio-user python3 -m unittest tests/test_server_cli.py -v

matrix: test
	python3 matrix.py --all --work-dir build/matrix

$(TARGET): $(SRC)
	mkdir -p build
	$(CC) $(CFLAGS) -o $(TARGET) $(SRC) $(LDFLAGS)

clean:
	rm -rf build

run: $(TARGET)
	./$(TARGET) --artifacts ../tests/cocotb/out_generic --socket /tmp/pcileech-vfio.sock
