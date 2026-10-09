## Go Development Guide - From Maven to Go

Since you're familiar with Maven (`mvn clean install`), here's how Go testing and building works:

### Installation ✅

Go 1.27.2 is now installed on your Mac:

```bash
go version
# go version go1.27.2 darwin/arm64
```

---

## Running Tests (Like `mvn clean install`)

### 1. Run ALL Tests (Unit + Integration)
```bash
cd /Users/bryancastillomartel/projects/exemplar
go test -v ./...
```

**Output:**
- Integration tests from `internal/handler`
- Unit tests from `internal/service`
- Coverage percentages for each package
- Pass/Fail for each test function

### 2. Run Unit Tests Only
```bash
go test -v ./internal/service/...
```

Runs only service tests (business logic, validation, data handling)

### 3. Run Integration Tests Only
```bash
go test -v ./internal/handler/...
```

Runs only API endpoint tests

### 4. Run with Coverage Report
```bash
go test -cover ./...
```

**Output shows:**
```
internal/handler    coverage: 71.6% of statements
internal/service    coverage: 85.4% of statements
```

### 5. Generate HTML Coverage Report
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

Opens an HTML report showing which lines are tested

### 6. Run with Race Detector (Find Concurrency Bugs)
```bash
go test -race ./...
```

---

## Building the App

### 1. Build the Binary
```bash
go build -o app ./cmd
```

Creates an executable file named `app`

### 2. Build for Multiple Platforms
```bash
mkdir -p dist

# Linux
GOOS=linux GOARCH=amd64 go build -o dist/app-linux ./cmd

# macOS
GOOS=darwin GOARCH=amd64 go build -o dist/app-macos ./cmd

# Windows
GOOS=windows GOARCH=amd64 go build -o dist/app-windows.exe ./cmd
```

---

## Running the Application

### Start the Server
```bash
./app
```

**Output:**
```
Server running on port 3000
```

### Start on Custom Port
```bash
PORT=8080 ./app
```

### Run Directly Without Building
```bash
go run ./cmd/main.go
```

---

## Using the Makefile (Easiest Way)

The project includes a Makefile with convenient commands:

```bash
# Run all tests
make test

# Run unit tests only
make test-unit

# Run integration tests only
make test-integration

# Run tests with coverage
make test-coverage

# Build the app
make build

# Run the app
make run

# Clean build artifacts
make clean

# View all available commands
make help
```

---

## Quick Comparison: Maven vs Go

| Maven | Go |
|-------|-----|
| `mvn clean` | `go clean` |
| `mvn test` | `go test ./...` |
| `mvn verify` | `go test -race ./...` |
| `mvn install` | `go build ./cmd` |
| `mvn package` | `go build -o app ./cmd` |
| `pom.xml` | `go.mod` + `go.sum` |

---

## Project Layout (Go Standards)

```
exemplar/
├── cmd/main.go                    # Entry point (like pom.xml main class)
├── internal/service/              # Internal packages (private)
│   ├── service.go                 # Business logic
│   ├── utils.go                   # Utility functions
│   ├── service_test.go            # Unit tests for service
│   └── utils_test.go              # Unit tests for utils
├── internal/handler/              # HTTP handlers
│   ├── handler.go                 # API routes & handlers
│   └── handler_test.go            # Integration tests
├── go.mod                         # Module definition (like pom.xml)
├── go.sum                         # Dependency checksums
└── Makefile                       # Build/test commands
```

---

## Test Results Summary

✅ **All 40+ Tests Passing:**
- 10 Integration Tests (API endpoints)
- 25 Unit Tests (Service & Utils)
- Coverage: Service 85.4%, Handler 71.6%

---

## Common Tasks

### Task: Add a New Feature
```bash
1. Write tests first (TDD)
2. Write implementation
3. Run: go test ./...
4. Build: go build -o app ./cmd
```

### Task: Check Code Quality
```bash
go vet ./...           # Find potential bugs
go fmt ./...           # Format code
go test -race ./...    # Check for race conditions
```

### Task: Deploy
```bash
# Build for production
GOOS=linux GOARCH=amd64 go build -o app ./cmd

# Or use make
make build-all         # Builds for Linux, macOS, Windows
```

---

## Environment Setup (Already Done ✅)

```bash
# Dependencies installed
✅ Go 1.27.2 installed via Homebrew
✅ go.mod configured with gorilla/mux dependency
✅ All packages downloaded and verified
✅ Tests passing locally
```

---

## Next Steps

1. **Explore the code:**
   ```bash
   cat internal/service/service.go    # View business logic
   cat internal/handler/handler.go    # View API handlers
   cat internal/service/service_test.go # View unit tests
   ```

2. **Run the app locally:**
   ```bash
   make run
   # Server runs on http://localhost:3000
   ```

3. **Make a test API call:**
   ```bash
   curl http://localhost:3000/health
   ```

4. **Modify code and rebuild:**
   ```bash
   # Edit a file, then:
   make test              # Verify tests pass
   make build             # Build new binary
   make run               # Run it
   ```

---

**You're all set to develop in Go! 🚀**
