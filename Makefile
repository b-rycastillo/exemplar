.PHONY: build run test test-unit test-integration test-coverage clean help deploy dev

# Variables
BINARY_NAME=exemplar
GO_VERSION=1.21
PORT?=3000

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build the application
	@echo "Building $(BINARY_NAME)..."
	@go build -v -o $(BINARY_NAME) ./cmd
	@echo "✅ Build complete"

run: build ## Run the application
	@echo "Starting server on port $(PORT)..."
	@PORT=$(PORT) ./$(BINARY_NAME)

dev: ## Run with hot reload (requires air)
	@echo "Starting development server..."
	@air

test: ## Run all tests
	@echo "Running all tests..."
	@go test -v ./...
	@echo "✅ All tests passed"

test-unit: ## Run unit tests only
	@echo "Running unit tests..."
	@go test -v ./internal/service/...
	@echo "✅ Unit tests passed"

test-integration: ## Run integration tests only
	@echo "Running integration tests..."
	@go test -v ./internal/handler/...
	@echo "✅ Integration tests passed"

test-coverage: ## Generate coverage report
	@echo "Generating coverage report..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "✅ Coverage report generated: coverage.html"

test-race: ## Run tests with race detector
	@echo "Running tests with race detector..."
	@go test -race ./...
	@echo "✅ No race conditions detected"

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@go clean
	@rm -f $(BINARY_NAME)
	@rm -f coverage.out coverage.html
	@rm -rf dist/
	@echo "✅ Clean complete"

format: ## Format code
	@echo "Formatting code..."
	@go fmt ./...
	@echo "✅ Code formatted"

lint: ## Run linter
	@echo "Running linter..."
	@golangci-lint run ./...
	@echo "✅ Linting complete"

vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...
	@echo "✅ Vetting complete"

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "✅ Dependencies downloaded"

build-all: ## Build for all platforms
	@echo "Building for all platforms..."
	@mkdir -p dist
	@GOOS=linux GOARCH=amd64 go build -v -o dist/$(BINARY_NAME)-linux-amd64 ./cmd
	@GOOS=darwin GOARCH=amd64 go build -v -o dist/$(BINARY_NAME)-darwin-amd64 ./cmd
	@GOOS=windows GOARCH=amd64 go build -v -o dist/$(BINARY_NAME)-windows-amd64.exe ./cmd
	@echo "✅ Multi-platform build complete"

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@docker build -t exemplar:latest .
	@echo "✅ Docker image built"

docker-run: docker-build ## Run Docker image
	@echo "Running Docker container..."
	@docker run -p 3000:3000 exemplar:latest

check: format vet test ## Run all checks
	@echo "✅ All checks passed"
