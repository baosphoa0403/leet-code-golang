# -----------------------------
# Go Project Makefile
# -----------------------------

APP_NAME := leet_code
MAIN := .
BINARY := ./tmp/$(APP_NAME)
GO := go

# Git & version info
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME := $(shell date '+%Y-%m-%d_%H:%M:%S')
GO_VERSION := $(shell go version | awk '{print $$3}')

# Default target
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build        - Build the binary"
	@echo "  make run          - Run the app"
	@echo "  make clean        - Remove build files"
	@echo "  make test         - Run tests"
	@echo "  make lint         - Run static check"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-run   - Run Docker container"
	@echo "  make info         - Show build info"

# -----------------------------
# Core targets
# -----------------------------

build:
	@echo "🔧 Building $(APP_NAME)..."
	@$(GO) build -ldflags "-X main.Commit=$(GIT_COMMIT) -X main.BuildTime=$(BUILD_TIME)" -o $(BINARY) $(MAIN)
	@echo "✅ Build done: $(BINARY)"

run:
	@echo "🚀 Running $(APP_NAME)..."
	@$(GO) run $(MAIN)

clean:
	@echo "🧹 Cleaning..."
	@rm -rf tmp
	@echo "✅ Clean done"

test:
	@echo "🧪 Running tests..."
	@$(GO) test -v ./...

lint:
	@echo "🔍 Running lint..."
	@golangci-lint run --fix || echo "Linting skipped (install golangci-lint?)"

info:
	@echo "📦 Build Info"
	@echo "Commit: $(GIT_COMMIT)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Go Version: $(GO_VERSION)"

# -----------------------------
# Docker
# -----------------------------

docker-build:
	@echo "🐳 Building Docker image..."
	docker build -t $(APP_NAME):latest .

docker-run:
	@echo "🐳 Running Docker container..."
	docker run --rm -it -p 8080:8080 $(APP_NAME):latest
