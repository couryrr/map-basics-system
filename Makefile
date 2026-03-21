# Makefile for Boids project
-include .env

APP_NAME := map-basics
MAIN_FILE := .
BUILD_DIR := bin
ASSETS_DIR := assets

WINDOWS_DEPLOY_DIR  := $(WINDOWS_BASE_DEPLOY)/$(APP_NAME)
WINDOWS_DEPLOY_PATH := $(WINDOWS_DEPLOY_DIR)/$(APP_NAME).exe

# Default target
.PHONY: all
all: windows

# Build for Windows and deploy
.PHONY: windows
windows:
	@echo "Building Windows executable..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 \
		go build -ldflags "-s -w" -o $(BUILD_DIR)/$(APP_NAME).exe $(MAIN_FILE)
	@echo "Copying to Windows deployment location..."
	@mkdir -p $(WINDOWS_DEPLOY_DIR)
	cp $(BUILD_DIR)/$(APP_NAME).exe $(WINDOWS_DEPLOY_PATH)
	cp -r $(ASSETS_DIR) $(WINDOWS_DEPLOY_DIR)
	@echo "✅ Windows build complete: $(WINDOWS_DEPLOY_PATH)"

# Build only (no deploy)
.PHONY: build
build:
	@echo "Building Windows executable..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 \
		go build -ldflags "-s -w" -o $(BUILD_DIR)/$(APP_NAME).exe $(MAIN_FILE)
	@echo "✅ Build complete: $(BUILD_DIR)/$(APP_NAME).exe"

# Deploy existing binary
.PHONY: deploy
deploy:
	@echo "Deploying to Windows..."
	@mkdir -p $(WINDOWS_DEPLOY_DIR)
	cp $(BUILD_DIR)/$(APP_NAME).exe $(WINDOWS_DEPLOY_PATH)
	cp -r $(ASSETS_DIR) $(WINDOWS_DEPLOY_DIR)
	@echo "✅ Deployed to: $(WINDOWS_DEPLOY_PATH)"

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build directory..."
	rm -rf $(BUILD_DIR)
	@echo "✅ Clean complete"

# Build for Linux (for local testing)
.PHONY: linux
linux:
	@echo "Building Linux executable..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	@echo "✅ Linux build complete: $(BUILD_DIR)/$(APP_NAME)"

# Help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  windows  - Build for Windows and deploy (default)"
	@echo "  build    - Build for Windows only"
	@echo "  deploy   - Deploy existing binary to Windows"
	@echo "  linux    - Build for Linux"
	@echo "  clean    - Clean build directory"
	@echo "  help     - Show this help"
