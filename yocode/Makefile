.PHONY: all build clean run test

# Directories
DIST_DIR := dist
BIN_NAME := yocode
GO_DIR := go
TS_DIR := ts
ASSETS_DIR := $(GO_DIR)/editor/assets

# Default target
all: build

# Build everything
build: prepare ts-build go-build

# Prepare build directories
prepare:
	@echo "Preparing build directories..."
	@mkdir -p $(DIST_DIR)
	@mkdir -p $(ASSETS_DIR)

# Build TypeScript UI
ts-build: prepare
	@echo "Building TypeScript UI..."
	@cd $(TS_DIR) && npm install && npm run build
	@cp -r $(TS_DIR)/dist/*.js $(ASSETS_DIR)/
	@cp -r $(TS_DIR)/src/index.html $(ASSETS_DIR)/
	@echo "TypeScript UI built and copied to assets directory"

# Build the Go binary
go-build: prepare
	@echo "Building Go binary..."
	@cd $(GO_DIR) && CGO_ENABLED=1 go build -o ../$(DIST_DIR)/$(BIN_NAME) ./cmd/yocode
	@echo "Go binary built as '$(DIST_DIR)/$(BIN_NAME)'"

# Run the application
run: build
	@echo "Running Yocode..."
	@cd $(DIST_DIR) && ./$(BIN_NAME) ../test/main.go

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf $(DIST_DIR)
	@rm -rf $(ASSETS_DIR)
	@cd $(TS_DIR) && npm run clean || true
	@echo "Clean completed"

# Install the application to /usr/local/bin
install: build
	@echo "Installing Yocode..."
	@cp $(DIST_DIR)/$(BIN_NAME) /usr/local/bin/
	@echo "Yocode installed to /usr/local/bin/$(BIN_NAME)"
