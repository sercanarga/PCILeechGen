.PHONY: build test lint clean install fixtures hdl-lint early-access cocotb-bootstrap cocotb-test cocotb-test-stable cocotb-test-experimental cocotb-clean

override BINARY_NAME := pcileechgen
override BUILD_DIR := bin
GO=go
GOFLAGS=-trimpath
LDFLAGS=-s -w

COCOTB_PYTHON ?= python3.13
COCOTB_VENV := bin/cocotb-venv
COCOTB_OUTPUT_ROOT := tests/cocotb/out_matrix
COCOTB_STABLE_CASES := nvme generic audio xhci wifi sata gpu
COCOTB_EXPERIMENTAL_CASES := ethernet
COCOTB_CASES ?= $(COCOTB_STABLE_CASES)

VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS += -X github.com/sercanarga/pcileechgen/internal/version.Version=$(VERSION)

# Build the binary
build:
	$(GO) build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/pcileechgen/

# Run all tests
test:
	$(GO) test -race -count=1 ./...

# Run tests with coverage
test-coverage:
	$(GO) test -race -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html

# Run linter
lint:
	golangci-lint run ./...

# Clean build artifacts
clean:
	python3 tools/cleanup_generated.py --repo-root "$(CURDIR)" --top-level

# generate synthetic donor fixture jsons
fixtures: build
	$(BUILD_DIR)/$(BINARY_NAME) fixtures --out testdata/donors

# lint generated sv with verilator
hdl-lint: build
	@command -v verilator >/dev/null 2>&1 || { echo "verilator not installed; see Makefile"; exit 1; }
	./scripts/hdl-lint.sh

cocotb-bootstrap:
	@command -v "$(COCOTB_PYTHON)" >/dev/null 2>&1 || { echo "$(COCOTB_PYTHON) is required (Python <= 3.13)" >&2; exit 1; }
	$(COCOTB_PYTHON) tests/cocotb/run_matrix.py --check-python --check-venv "$(COCOTB_VENV)"
	@test -x "$(COCOTB_VENV)/bin/python" || $(COCOTB_PYTHON) -m venv "$(COCOTB_VENV)"
	"$(COCOTB_VENV)/bin/python" tests/cocotb/run_matrix.py --check-python
	"$(COCOTB_VENV)/bin/python" -m pip install --requirement tests/cocotb/requirements.txt

cocotb-test: build cocotb-bootstrap
	"$(COCOTB_VENV)/bin/python" tests/cocotb/test_runner.py
	"$(COCOTB_VENV)/bin/python" tests/cocotb/run_matrix.py \
		--generator "$(BUILD_DIR)/$(BINARY_NAME)" \
		--output-root "$(COCOTB_OUTPUT_ROOT)" \
		$(foreach case,$(COCOTB_CASES),--case $(case))

cocotb-test-stable:
	$(MAKE) cocotb-test COCOTB_CASES="$(COCOTB_STABLE_CASES)"

cocotb-test-experimental:
	$(MAKE) cocotb-test COCOTB_CASES="$(COCOTB_EXPERIMENTAL_CASES)"

cocotb-clean:
	$(COCOTB_PYTHON) tests/cocotb/run_matrix.py --clean --output-root "$(COCOTB_OUTPUT_ROOT)"

# Install binary to GOPATH/bin
install:
	$(GO) install $(GOFLAGS) -ldflags "$(LDFLAGS)" ./cmd/pcileechgen/

# Format code
fmt:
	$(GO) fmt ./...

# Vet code
vet:
	$(GO) vet ./...

# Release build
release-build:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(OUTPUT) ./cmd/pcileechgen/

# Run all checks (vet + lint + test)
check: vet lint test

# Early access
override EARLYACCESS_NAME := PCILeechGen-EarlyAccess
override EARLYACCESS_STAMP := $(shell date +%m%d%y)
override EARLYACCESS_ZIP := $(EARLYACCESS_NAME)-$(EARLYACCESS_STAMP).zip
EARLYACCESS_REQUIRED = \
	Makefile go.mod cmd/pcileechgen/main.go \
	internal/firmware/svgen/templates/ethernet_dma_engine.sv.tmpl \
	internal/firmware/svgen/templates/bar_controller.sv.tmpl \
	internal/firmware/svgen/templates/bar_impl_device.sv.tmpl \
	internal/firmware/output/writer.go \
	vfio-user/src/behavior_ethernet.c tests/cocotb/test_eth_init.py

early-access:
	@command -v zip >/dev/null 2>&1 || { echo "zip is required" >&2; exit 1; }
	@command -v unzip >/dev/null 2>&1 || { echo "unzip is required" >&2; exit 1; }
	@python3 tools/cleanup_generated.py --repo-root "$(CURDIR)" --early-access-archives
	@zip -r "$(EARLYACCESS_ZIP)" . \
		-x '.git/*' 'analysis/*' 'bin/*' 'dist/*' 'pcileech_datastore/*' \
		   'vfio-user/build' 'vfio-user/build/*' 'tests/cocotb/out*' 'tests/cocotb/out*/**' \
		   'tests/cocotb/sim_build' 'tests/cocotb/sim_build/' 'tests/cocotb/sim_build/**' \
		   '**/__pycache__' '**/__pycache__/' '**/__pycache__/**' \
		   '*.pyc' '*.pyo' '*.dSYM' '*.dSYM/**' '*.o' '*.so' \
		   '*.zip' '.DS_Store' '*/.DS_Store' '.idea/*' '.vscode/*' \
		   'coverage.out' 'coverage.html' 'hdl-lint-report.tsv' '*.swp' '*.swo' '*~' \
		> /dev/null
	@unzip -t "$(EARLYACCESS_ZIP)" >/dev/null
	@for file in $(EARLYACCESS_REQUIRED); do \
		unzip -Z1 "$(EARLYACCESS_ZIP)" | grep -Fxq "$$file" || { \
			echo "early-access archive is missing $$file" >&2; exit 1; \
		}; \
	done
	@echo "$(EARLYACCESS_ZIP)"
