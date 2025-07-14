# Makefile for GoEdge: Go Language Mastery Project

# Variables
APP_NAME := goedge
CMD_DIR := ./cmd/goedge
BIN_DIR := bin
BUILD_FLAGS := -ldflags="-s -w"
COVERAGE_FILE := coverage.out

# Colors for output
RED := \033[0;31m
GREEN := \033[0;32m
YELLOW := \033[0;33m
BLUE := \033[0;34m
PURPLE := \033[0;35m
CYAN := \033[0;36m
WHITE := \033[0;37m
BOLD := \033[1m
RESET := \033[0m

.PHONY: help pointers functions methods interfaces goroutines channels all clean build run install uninstall fmt vet test test-coverage test-race deps dev ci pre-commit docker-build docker-run benchmark profile security lint

# Default target
help:
	@echo "$(BOLD)$(CYAN)🐹 GoEdge: Go Language Mastery Project - Available Commands:$(RESET)"
	@echo "$(BOLD)===============================================================$(RESET)"
	@echo ""
	@echo "$(BOLD)$(YELLOW)📚 Learning Commands:$(RESET)"
	@echo "  $(GREEN)make pointers$(RESET)    - Run pointer examples"
	@echo "  $(GREEN)make functions$(RESET)   - Run function examples"
	@echo "  $(GREEN)make methods$(RESET)     - Run method examples"
	@echo "  $(GREEN)make interfaces$(RESET)  - Run interface examples"
	@echo "  $(GREEN)make goroutines$(RESET)  - Run goroutine examples"
	@echo "  $(GREEN)make channels$(RESET)    - Run channel examples"
	@echo "  $(GREEN)make all$(RESET)         - Run all examples"
	@echo ""
	@echo "$(BOLD)$(YELLOW)🔨 Build Commands:$(RESET)"
	@echo "  $(GREEN)make build$(RESET)       - Build the project"
	@echo "  $(GREEN)make run$(RESET)         - Build and run the project"
	@echo "  $(GREEN)make install$(RESET)     - Install binary to GOPATH/bin"
	@echo "  $(GREEN)make uninstall$(RESET)   - Remove binary from GOPATH/bin"
	@echo "  $(GREEN)make clean$(RESET)       - Clean build artifacts"
	@echo ""
	@echo "$(BOLD)$(YELLOW)🧪 Testing Commands:$(RESET)"
	@echo "  $(GREEN)make test$(RESET)        - Run tests"
	@echo "  $(GREEN)make test-coverage$(RESET) - Run tests with coverage"
	@echo "  $(GREEN)make test-race$(RESET)   - Run tests with race detection"
	@echo "  $(GREEN)make benchmark$(RESET)   - Run benchmarks"
	@echo ""
	@echo "$(BOLD)$(YELLOW)🔍 Quality Commands:$(RESET)"
	@echo "  $(GREEN)make fmt$(RESET)         - Format code"
	@echo "  $(GREEN)make vet$(RESET)         - Vet code"
	@echo "  $(GREEN)make lint$(RESET)        - Run linter (golangci-lint)"
	@echo "  $(GREEN)make security$(RESET)    - Run security scan (gosec)"
	@echo ""
	@echo "$(BOLD)$(YELLOW)🚀 Development Commands:$(RESET)"
	@echo "  $(GREEN)make deps$(RESET)        - Install/update dependencies"
	@echo "  $(GREEN)make dev$(RESET)         - Full development cycle"
	@echo "  $(GREEN)make ci$(RESET)          - Run CI pipeline locally"
	@echo "  $(GREEN)make pre-commit$(RESET)  - Run pre-commit checks"
	@echo ""
	@echo "$(BOLD)$(YELLOW)🐳 Docker Commands:$(RESET)"
	@echo "  $(GREEN)make docker-build$(RESET) - Build Docker image"
	@echo "  $(GREEN)make docker-run$(RESET)  - Run Docker container"
	@echo ""
	@echo "$(BOLD)Example: $(CYAN)make pointers$(RESET)"

# Learning Commands
pointers:
	@echo "$(BOLD)$(BLUE)🔗 Running pointer examples...$(RESET)"
	@go run $(CMD_DIR) pointers

functions:
	@echo "$(BOLD)$(BLUE)⚡ Running function examples...$(RESET)"
	@go run $(CMD_DIR) functions

methods:
	@echo "$(BOLD)$(BLUE)📋 Running method examples...$(RESET)"
	@go run $(CMD_DIR) methods

interfaces:
	@echo "$(BOLD)$(BLUE)🔌 Running interface examples...$(RESET)"
	@go run $(CMD_DIR) interfaces

goroutines:
	@echo "$(BOLD)$(BLUE)🏃 Running goroutine examples...$(RESET)"
	@go run $(CMD_DIR) goroutines

channels:
	@echo "$(BOLD)$(BLUE)📡 Running channel examples...$(RESET)"
	@go run $(CMD_DIR) channels

all:
	@echo "$(BOLD)$(BLUE)🌟 Running all examples...$(RESET)"
	@go run $(CMD_DIR) all

# Build Commands
build:
	@echo "$(BOLD)$(YELLOW)🔨 Building project...$(RESET)"
	@mkdir -p $(BIN_DIR)
	@go build $(BUILD_FLAGS) -o $(BIN_DIR)/$(APP_NAME) $(CMD_DIR)
	@echo "$(GREEN)✅ Build completed: $(BIN_DIR)/$(APP_NAME)$(RESET)"

install: build
	@echo "$(BOLD)$(YELLOW)📦 Installing binary...$(RESET)"
	@go install $(CMD_DIR)
	@echo "$(GREEN)✅ Installed to GOPATH/bin/$(APP_NAME)$(RESET)"

uninstall:
	@echo "$(BOLD)$(YELLOW)🗑️ Uninstalling binary...$(RESET)"
	@rm -f $(shell go env GOPATH)/bin/$(APP_NAME)
	@echo "$(GREEN)✅ Uninstalled$(RESET)"

clean:
	@echo "$(BOLD)$(YELLOW)🧹 Cleaning build artifacts...$(RESET)"
	@rm -rf $(BIN_DIR)
	@rm -f $(COVERAGE_FILE)
	@go clean -cache -testcache -modcache
	@echo "$(GREEN)✅ Clean completed$(RESET)"

run: build
	@echo "$(BOLD)$(BLUE)🚀 Running built binary...$(RESET)"
	@./$(BIN_DIR)/$(APP_NAME)

# Testing Commands
test:
	@echo "$(BOLD)$(PURPLE)🧪 Running tests...$(RESET)"
	@go test -v ./...

test-coverage:
	@echo "$(BOLD)$(PURPLE)📊 Running tests with coverage...$(RESET)"
	@go test -v -race -coverprofile=$(COVERAGE_FILE) ./...
	@go tool cover -html=$(COVERAGE_FILE) -o coverage.html
	@echo "$(GREEN)✅ Coverage report generated: coverage.html$(RESET)"

test-race:
	@echo "$(BOLD)$(PURPLE)🏁 Running tests with race detection...$(RESET)"
	@go test -v -race ./...

benchmark:
	@echo "$(BOLD)$(PURPLE)⚡ Running benchmarks...$(RESET)"
	@go test -bench=. -benchmem ./...

# Quality Commands
fmt:
	@echo "$(BOLD)$(CYAN)🎨 Formatting code...$(RESET)"
	@go fmt ./...
	@echo "$(GREEN)✅ Code formatted$(RESET)"

vet:
	@echo "$(BOLD)$(CYAN)🔍 Vetting code...$(RESET)"
	@go vet ./...
	@echo "$(GREEN)✅ Code vetted$(RESET)"

lint:
	@echo "$(BOLD)$(CYAN)🔎 Running linter...$(RESET)"
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "$(YELLOW)⚠️  golangci-lint not found. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest$(RESET)"; \
	fi

security:
	@echo "$(BOLD)$(RED)🔒 Running security scan...$(RESET)"
	@if command -v gosec >/dev/null 2>&1; then \
		gosec ./...; \
	else \
		echo "$(YELLOW)⚠️  gosec not found. Install with: go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest$(RESET)"; \
	fi

# Development Commands
deps:
	@echo "$(BOLD)$(BLUE)📦 Installing/updating dependencies...$(RESET)"
	@go mod tidy
	@go mod download
	@go mod verify
	@echo "$(GREEN)✅ Dependencies updated$(RESET)"

dev: deps fmt vet build
	@echo "$(BOLD)$(GREEN)🔧 Development build completed$(RESET)"

ci: deps fmt vet lint security test-race build
	@echo "$(BOLD)$(GREEN)🚀 CI pipeline completed successfully$(RESET)"

pre-commit: fmt vet test
	@echo "$(BOLD)$(GREEN)✅ Pre-commit checks passed$(RESET)"

# Docker Commands
docker-build:
	@echo "$(BOLD)$(BLUE)🐳 Building Docker image...$(RESET)"
	@docker build -t $(APP_NAME):latest .
	@echo "$(GREEN)✅ Docker image built: $(APP_NAME):latest$(RESET)"

docker-run: docker-build
	@echo "$(BOLD)$(BLUE)🐳 Running Docker container...$(RESET)"
	@docker run --rm -it $(APP_NAME):latest

# Profiling Commands
profile:
	@echo "$(BOLD)$(PURPLE)📈 Running CPU profiling...$(RESET)"
	@go test -cpuprofile=cpu.prof -memprofile=mem.prof -bench=. ./...
	@echo "$(GREEN)✅ Profiles generated: cpu.prof, mem.prof$(RESET)"
	@echo "$(CYAN)View with: go tool pprof cpu.prof$(RESET)"

# Check if required tools are installed
check-tools:
	@echo "$(BOLD)$(CYAN)🔧 Checking required tools...$(RESET)"
	@command -v go >/dev/null 2>&1 || { echo "$(RED)❌ Go is not installed$(RESET)"; exit 1; }
	@command -v git >/dev/null 2>&1 || { echo "$(RED)❌ Git is not installed$(RESET)"; exit 1; }
	@command -v make >/dev/null 2>&1 || { echo "$(RED)❌ Make is not installed$(RESET)"; exit 1; }
	@echo "$(GREEN)✅ All required tools are installed$(RESET)"

# Version information
version:
	@echo "$(BOLD)$(CYAN)📋 Version Information:$(RESET)"
	@echo "Go version: $(shell go version)"
	@echo "Git commit: $(shell git rev-parse --short HEAD 2>/dev/null || echo 'unknown')"
	@echo "Build date: $(shell date -u '+%Y-%m-%d %H:%M:%S UTC')"