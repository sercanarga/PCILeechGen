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
SRC = src/main.c src/vfio_device.c src/device_model.c src/behavior_static.c src/behavior_nvme.c src/behavior_hda.c src/behavior_gpu.c src/behavior_ahci.c src/behavior_xhci.c src/behavior_ethernet.c

.PHONY: all image test test-matrix test-device-model test-behavior-static test-behavior-nvme test-behavior-hda test-behavior-gpu test-behavior-storage test-server matrix fuzz fuzz-smoke test-valgrind test-msan

all: $(TARGET)

test: test-matrix test-device-model test-behavior-static test-behavior-nvme test-behavior-hda test-behavior-gpu test-behavior-storage

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

build/test_behavior_static: tests/test_behavior_static.c src/behavior_static.c src/behavior_nvme.c src/behavior_hda.c src/behavior_gpu.c src/behavior_ahci.c src/behavior_xhci.c src/behavior_ethernet.c include/device_behavior.h include/device_model.h
	mkdir -p build
	$(CC) $(CFLAGS) $(SANITIZE) \
		$$(pkg-config --cflags cmocka) -o $@ tests/test_behavior_static.c src/behavior_static.c src/behavior_nvme.c src/behavior_hda.c src/behavior_gpu.c src/behavior_ahci.c src/behavior_xhci.c src/behavior_ethernet.c \
		$$(pkg-config --libs cmocka)

test-behavior-static: build/test_behavior_static
	./build/test_behavior_static

build/test_behavior_nvme: tests/test_behavior_nvme.c src/behavior_nvme.c src/behavior_static.c src/behavior_hda.c src/behavior_gpu.c src/behavior_ahci.c src/behavior_xhci.c src/behavior_ethernet.c src/device_model.c include/device_behavior.h include/device_model.h
	mkdir -p build
	$(CC) $(CFLAGS) $(SANITIZE) \
		$$(pkg-config --cflags cmocka json-c openssl) -o $@ \
		tests/test_behavior_nvme.c src/behavior_nvme.c src/behavior_static.c src/behavior_hda.c src/behavior_gpu.c src/behavior_ahci.c src/behavior_xhci.c src/behavior_ethernet.c src/device_model.c \
		$$(pkg-config --libs cmocka json-c openssl)

test-behavior-nvme: build/test_behavior_nvme
	./build/test_behavior_nvme

build/test_behavior_hda: tests/test_behavior_hda.c src/behavior_hda.c src/behavior_ethernet.c src/behavior_static.c src/behavior_nvme.c src/behavior_gpu.c src/behavior_ahci.c src/behavior_xhci.c src/device_model.c include/device_behavior.h include/device_model.h
	mkdir -p build
	$(CC) $(CFLAGS) $(SANITIZE) \
		$$(pkg-config --cflags cmocka json-c openssl) -o $@ \
		tests/test_behavior_hda.c src/behavior_hda.c src/behavior_ethernet.c src/behavior_static.c src/behavior_nvme.c src/behavior_gpu.c src/behavior_ahci.c src/behavior_xhci.c src/device_model.c \
		$$(pkg-config --libs cmocka json-c openssl)

test-behavior-hda: build/test_behavior_hda
	./build/test_behavior_hda

build/test_behavior_gpu: tests/test_behavior_gpu.c src/behavior_gpu.c src/behavior_ethernet.c src/behavior_static.c src/behavior_nvme.c src/behavior_hda.c src/behavior_ahci.c src/behavior_xhci.c src/device_model.c include/device_behavior.h include/device_model.h
	mkdir -p build
	$(CC) $(CFLAGS) $(SANITIZE) \
		$$(pkg-config --cflags cmocka json-c openssl) -o $@ \
		tests/test_behavior_gpu.c src/behavior_gpu.c src/behavior_ethernet.c src/behavior_static.c src/behavior_nvme.c src/behavior_hda.c src/behavior_ahci.c src/behavior_xhci.c src/device_model.c \
		$$(pkg-config --libs cmocka json-c openssl)

test-behavior-gpu: build/test_behavior_gpu
	./build/test_behavior_gpu

build/test_behavior_storage: tests/test_behavior_storage.c src/behavior_ahci.c src/behavior_xhci.c src/behavior_static.c src/behavior_nvme.c src/behavior_hda.c src/behavior_gpu.c src/behavior_ethernet.c src/device_model.c include/device_behavior.h include/device_model.h
	mkdir -p build
	$(CC) $(CFLAGS) $(SANITIZE) \
		$$(pkg-config --cflags cmocka json-c openssl) -o $@ \
		tests/test_behavior_storage.c src/behavior_ahci.c src/behavior_xhci.c \
		src/behavior_static.c src/behavior_nvme.c src/behavior_hda.c src/behavior_gpu.c src/behavior_ethernet.c src/device_model.c \
		$$(pkg-config --libs cmocka json-c openssl)

test-behavior-storage: build/test_behavior_storage
	./build/test_behavior_storage

fuzz:
	@if ! command -v clang >/dev/null 2>&1; then echo 'clang is required for fuzz; install clang or use a clang toolchain' >&2; exit 2; fi
	@mkdir -p build
	clang $(CFLAGS) -fsanitize=fuzzer,address,undefined -fno-omit-frame-pointer \
		$$(pkg-config --cflags json-c openssl) -o build/fuzz_behavior \
		tests/fuzz_behavior.c src/behavior_nvme.c src/behavior_static.c src/behavior_ethernet.c \
		src/behavior_hda.c src/behavior_gpu.c src/behavior_ahci.c src/behavior_xhci.c src/device_model.c \
		$$(pkg-config --libs json-c openssl)
	./build/fuzz_behavior -runs=1000

fuzz-smoke:
	@if ! command -v clang >/dev/null 2>&1; then echo 'clang is required for fuzz-smoke; install clang or use a clang toolchain' >&2; exit 2; fi
	@mkdir -p build
	clang $(CFLAGS) -DFUZZ_SMOKE $(SANITIZE) -fno-omit-frame-pointer \
		$$(pkg-config --cflags json-c openssl) -o build/fuzz_behavior_smoke \
		tests/fuzz_behavior.c src/behavior_nvme.c src/behavior_static.c src/behavior_ethernet.c \
		src/behavior_hda.c src/behavior_gpu.c src/behavior_ahci.c src/behavior_xhci.c src/device_model.c \
		$$(pkg-config --libs json-c openssl)
	./build/fuzz_behavior_smoke

test-valgrind:
	@if ! command -v valgrind >/dev/null 2>&1; then \
		echo 'valgrind not installed; skipping optional memory check'; \
	else \
		$(MAKE) clean test && \
		valgrind --leak-check=full --error-exitcode=1 ./build/test_device_model && \
		valgrind --leak-check=full --error-exitcode=1 ./build/test_behavior_nvme && \
		valgrind --leak-check=full --error-exitcode=1 ./build/test_behavior_hda && \
		valgrind --leak-check=full --error-exitcode=1 ./build/test_behavior_storage; \
	fi

test-msan:
	@if [ "$(UNAME_S)" != "Linux" ] || ! command -v clang >/dev/null 2>&1; then \
		echo 'MSan requires Linux and clang; skipping optional memory check'; \
	else \
		$(MAKE) clean test SANITIZE='-fsanitize=memory -fno-omit-frame-pointer'; \
	fi

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
