.PHONY: build install clean run test help

# Build variables
BINARY_NAME=sg
BUILD_DIR=./bin
INSTALL_DIR=/usr/local/bin

# Go variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build the application
build:
	@echo "🔨 Building sicktiergit..."
	@mkdir -p $(BUILD_DIR)
	@$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v
	@echo "✓ Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Install the application
install: build
	@echo "📦 Installing sicktiergit..."
	@cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)/
	@chmod +x $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "✓ Installed to $(INSTALL_DIR)/$(BINARY_NAME)"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning..."
	@$(GOCLEAN)
	@rm -rf $(BUILD_DIR)
	@echo "✓ Clean complete"

# Run the application
run: build
	@$(BUILD_DIR)/$(BINARY_NAME)

# Run tests
test:
	@echo "🧪 Running tests..."
	@$(GOTEST) -v ./...

# Download dependencies
deps:
	@echo "📥 Downloading dependencies..."
	@$(GOMOD) download
	@$(GOMOD) tidy
	@echo "✓ Dependencies ready"

# Development build with live reload
dev:
	@echo "🔧 Running in development mode..."
	@$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v && $(BUILD_DIR)/$(BINARY_NAME) $(ARGS)

# Show help
help:
	@echo "sg (sicktiergit) - Makefile commands"
	@echo ""
	@echo "Usage: make [command]"
	@echo ""
	@echo "Commands:"
	@echo "  build    - Build the application"
	@echo "  install  - Install to system (requires sudo)"
	@echo "  clean    - Remove build artifacts"
	@echo "  run      - Build and run the application"
	@echo "  test     - Run tests"
	@echo "  deps     - Download dependencies"
	@echo "  dev      - Development build and run"
	@echo "  help     - Show this help message"
	@echo ""
	@echo "Examples:"
	@echo "  make build"
	@echo "  make dev ARGS='init my-repo'"
	@echo "  sudo make install"
