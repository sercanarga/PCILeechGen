.PHONY: build test lint clean install fixtures hdl-lint early-access

BINARY_NAME=pcileechgen
BUILD_DIR=bin
GO=go
GOFLAGS=-trimpath
LDFLAGS=-s -w

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
	rm -rf $(BUILD_DIR)
	rm -f hdl-lint-report.tsv
	rm -f coverage.out coverage.html
	rm -f PCILeechGen-EarlyAccess-*.zip

# generate synthetic donor fixture jsons
fixtures: build
	$(BUILD_DIR)/$(BINARY_NAME) fixtures --out testdata/donors

# lint generated sv with verilator
hdl-lint: build
	@command -v verilator >/dev/null 2>&1 || { echo "verilator not installed; see Makefile"; exit 1; }
	./scripts/hdl-lint.sh

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
EARLYACCESS_NAME  ?= PCILeechGen-EarlyAccess
EARLYACCESS_STAMP ?= $(shell date +%m%d%y)
EARLYACCESS_ZIP   ?= $(EARLYACCESS_NAME)-$(EARLYACCESS_STAMP).zip

early-access:
	@rm -f "$(EARLYACCESS_ZIP)"
	@zip -r "$(EARLYACCESS_ZIP)" . \
		-x '.git/*' 'analysis/*' 'bin/*' 'dist/*' 'pcileech_datastore/*' \
		   '*.zip' '.DS_Store' '*/.DS_Store' '.idea/*' '.vscode/*' \
		   'coverage.out' 'coverage.html' '*.swp' '*.swo' '*~' \
		> /dev/null
	@echo "$(EARLYACCESS_ZIP)"
