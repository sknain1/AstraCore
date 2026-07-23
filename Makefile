APP_NAME := AstraCore

.PHONY: help install doctor build test run clean fmt

help:
	@echo "=================================="
	@echo "        $(APP_NAME)"
	@echo "=================================="
	@echo ""
	@echo "Available Commands:"
	@echo "  make help      - Show help"
	@echo "  make install   - Run installer"
	@echo "  make doctor    - Check system"
	@echo "  make build     - Build Go and Rust"
	@echo "  make test      - Run tests"
	@echo "  make run       - Run Go application"
	@echo "  make clean     - Clean build files"
	@echo "  make fmt       - Format source code"

install:
	./install.sh

doctor:
	./doctor.sh

build:
	@echo "===> Building Go..."
	cd broker-go && go build ./...

	@echo "===> Building Rust..."
	cd engine-rust && cargo build

	@echo ""
	@echo "✅ Build Complete"

test:
	@echo "===> Testing Go..."
	cd broker-go && go test ./...

	@echo "===> Testing Rust..."
	cd engine-rust && cargo test

	@echo ""
	@echo "✅ Tests Complete"

run:
	cd broker-go && go run ./cmd

clean:
	@echo "Cleaning..."

	cd broker-go && go clean

	cd engine-rust && cargo clean

	@echo "✅ Clean Complete"

fmt:
	@echo "Formatting Go..."
	cd broker-go && gofmt -w .

	@echo "Formatting Rust..."
	cd engine-rust && cargo fmt

	@echo "✅ Format Complete"
