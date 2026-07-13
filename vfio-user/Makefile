CC ?= gcc
.DEFAULT_GOAL := all
CFLAGS = -D_GNU_SOURCE -std=c11 -Wall -Wextra -Werror -O2 -g -Iinclude -I/usr/local/include
LDFLAGS = -L/usr/local/lib -lvfio-user $$(pkg-config --libs json-c openssl) -lpthread
UNAME_S := $(shell uname -s)
SANITIZE :=
ifeq ($(UNAME_S),Linux)
SANITIZE := -fsanitize=address,undefined
endif
FUZZ_CC ?= clang
FUZZ_RUNS ?= 10000
override FUZZ_FIXTURE_ROOT := $(abspath $(CURDIR)/build/fuzz-fixtures)
FUZZ_CORPUS_TEMPLATE := $(CURDIR)/build/fuzz-corpus.XXXXXX
FUZZ_SUMMARY := $(CURDIR)/build/fuzz-summary.txt
FUZZ_LOG := $(CURDIR)/build/fuzz-ci.log
FUZZ_LSAN_SUPPRESSIONS := $(CURDIR)/tests/lsan_json_c.supp
ROOT_DIR := $(abspath ..)
FUZZ_GENERATOR := $(ROOT_DIR)/bin/pcileechgen

TARGET = build/vfio-device
SRC = src/main.c src/vfio_device.c src/device_model.c src/behavior_static.c src/behavior_nvme.c src/behavior_hda.c src/behavior_gpu.c src/behavior_ahci.c src/behavior_xhci.c src/behavior_ethernet.c

.PHONY: all image test test-matrix test-device-model test-behavior-static test-behavior-nvme test-behavior-hda test-behavior-gpu test-behavior-storage test-server matrix fuzz fuzz-ci fuzz-fixtures fuzz-fixture-root-check fuzz-contract fuzz-smoke test-valgrind test-msan

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

fuzz-contract:
	python3 -m unittest tests/test_make_contract.py tests/test_fuzz_fixture_guard.py -v

fuzz-fixture-root-check:
	@python3 tests/fuzz_fixture_guard.py --project-root "$(CURDIR)" --fixture-root "$(FUZZ_FIXTURE_ROOT)"

fuzz-fixtures: fuzz-fixture-root-check
	@mkdir -p "$(FUZZ_FIXTURE_ROOT)"
	@$(MAKE) -C "$(ROOT_DIR)" build
	@printf '%s\n' 'generating isolated NVMe fuzz fixture'
	"$(FUZZ_GENERATOR)" build --from-json "$(ROOT_DIR)/testdata/donors/nvme.json" --board ac701_ft601 --skip-vivado --force --lib-dir "$(ROOT_DIR)/lib/pcileech-fpga" --output "$(FUZZ_FIXTURE_ROOT)/nvme"
	"$(FUZZ_GENERATOR)" verify-manifest --manifest "$(FUZZ_FIXTURE_ROOT)/nvme/build_manifest.json" --output-dir "$(FUZZ_FIXTURE_ROOT)/nvme"
	@printf '%s\n' 'generating isolated Ethernet fuzz fixture'
	"$(FUZZ_GENERATOR)" build --from-json "$(ROOT_DIR)/testdata/donors/ethernet.json" --board PCIeSquirrel --skip-vivado --force --lib-dir "$(ROOT_DIR)/lib/pcileech-fpga" --output "$(FUZZ_FIXTURE_ROOT)/ethernet"
	"$(FUZZ_GENERATOR)" verify-manifest --manifest "$(FUZZ_FIXTURE_ROOT)/ethernet/build_manifest.json" --output-dir "$(FUZZ_FIXTURE_ROOT)/ethernet"

build/fuzz_behavior: tests/fuzz_behavior.c src/behavior_nvme.c src/behavior_static.c src/behavior_ethernet.c src/behavior_hda.c src/behavior_gpu.c src/behavior_ahci.c src/behavior_xhci.c src/device_model.c
	@mkdir -p build
	$(FUZZ_CC) $(CFLAGS) -fsanitize=fuzzer,address,undefined -fno-omit-frame-pointer \
		$$(pkg-config --cflags json-c openssl) -o $@ \
		tests/fuzz_behavior.c src/behavior_nvme.c src/behavior_static.c src/behavior_ethernet.c \
		src/behavior_hda.c src/behavior_gpu.c src/behavior_ahci.c src/behavior_xhci.c src/device_model.c \
		$$(pkg-config --libs json-c openssl)

fuzz-ci:
	@mkdir -p build
	@printf 'fuzz-ci runs=%s compiler=%s\n' "$(FUZZ_RUNS)" "$(FUZZ_CC)" > "$(FUZZ_SUMMARY)"
	@command -v "$(FUZZ_CC)" >/dev/null 2>&1 || { echo 'libFuzzer-capable clang is required for fuzz-ci' >&2; exit 2; }
	@printf '%s\n' 'compiler preflight passed' >> "$(FUZZ_SUMMARY)"
	@$(MAKE) fuzz-fixtures > "$(CURDIR)/build/fuzz-fixtures.log" 2>&1 || { cat "$(CURDIR)/build/fuzz-fixtures.log"; exit 1; }
	@printf '%s\n' 'isolated fixtures verified' >> "$(FUZZ_SUMMARY)"
	@$(MAKE) build/fuzz_behavior
	@corpus_root="$$(mktemp -d "$(FUZZ_CORPUS_TEMPLATE)")"; \
		test -n "$$corpus_root" && test -d "$$corpus_root" || { echo 'failed to create isolated fuzz corpus' >&2; exit 1; }; \
		cp tests/fuzz-corpus/*.bin "$$corpus_root/"; \
		test "$$(find tests/fuzz-corpus -maxdepth 1 -type f -name '*.bin' | wc -l | tr -d ' ')" = 2 || exit 1; \
		VFIO_FUZZ_FIXTURE_ROOT="$(FUZZ_FIXTURE_ROOT)" ASAN_OPTIONS=detect_leaks=1:halt_on_error=1 LSAN_OPTIONS="suppressions=$(FUZZ_LSAN_SUPPRESSIONS):print_suppressions=0" UBSAN_OPTIONS=halt_on_error=1 ./build/fuzz_behavior -runs="$(FUZZ_RUNS)" -artifact_prefix="$$corpus_root/" "$$corpus_root" > "$(FUZZ_LOG)" 2>&1 || { cat "$(FUZZ_LOG)"; exit 1; }; \
		if grep -Eq 'ERROR: (AddressSanitizer|UndefinedBehaviorSanitizer|LeakSanitizer)|runtime error:' "$(FUZZ_LOG)"; then cat "$(FUZZ_LOG)"; exit 1; fi; \
		if find "$$corpus_root" -maxdepth 1 -type f \( -name 'crash-*' -o -name 'leak-*' -o -name 'timeout-*' \) -print -quit | grep -q .; then echo 'fuzz crash artifact detected' >&2; exit 1; fi
	@cat "$(FUZZ_LOG)"
	@printf '%s\n' 'sanitizers and crash-artifact checks passed' >> "$(FUZZ_SUMMARY)"

fuzz: fuzz-ci

fuzz-smoke:
	@$(MAKE) fuzz-ci FUZZ_RUNS=100

test-valgrind:
	@if ! command -v valgrind >/dev/null 2>&1; then \
		echo 'valgrind not installed; skipping optional memory check'; \
	else \
		$(MAKE) clean test SANITIZE= && \
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
