#!/bin/bash

# Quick Start Script for Exemplar Demo App
# This script shows you how to run tests and the app

echo "🚀 Exemplar Go App - Quick Start Guide"
echo "======================================"
echo ""

cd "$(dirname "$0")" || exit

# Show current status
echo "✅ Status Check:"
echo ""

echo "1️⃣ Go Version:"
go version
echo ""

echo "2️⃣ Project Structure:"
ls -la | grep -E "^d" | awk '{print "  " $9}' | grep -v "^\s*$"
echo ""

echo "3️⃣ Available Commands:"
echo ""
echo "  📝 TESTING:"
echo "     make test              # Run all tests (unit + integration)"
echo "     make test-unit         # Run unit tests only"
echo "     make test-integration  # Run integration tests only"
echo "     make test-coverage     # Generate coverage report"
echo ""

echo "  🔨 BUILDING & RUNNING:"
echo "     make build             # Build the app"
echo "     make run               # Run the app on port 3000"
echo "     make clean             # Clean build artifacts"
echo ""

echo "  🧹 CODE QUALITY:"
echo "     go fmt ./...           # Format code"
echo "     go vet ./...           # Check for bugs"
echo "     go test -race ./...    # Check for race conditions"
echo ""

echo "4️⃣ Quick Test Run:"
echo ""
if make test 2>/dev/null | grep -q "PASS"; then
    echo "  ✅ Tests: PASSING"
else
    echo "  ❌ Tests: Check output above"
fi

echo ""
echo "=================================="
echo "👉 Get Started:"
echo ""
echo "   1. Run all tests: make test"
echo "   2. Build the app: make build"
echo "   3. Start server:  make run"
echo "   4. Test endpoint: curl http://localhost:3000/health"
echo ""
echo "📚 Full Guide: See GO_GUIDE.md or README.md"
echo ""
