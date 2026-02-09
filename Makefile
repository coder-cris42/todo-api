.PHONY: help install-swagger-cli generate-swagger install-deps build run test clean
.PHONY: compose-up compose-down

help:
	@echo "Available commands:"
	@echo "  make install-deps       - Install all Go dependencies"
	@echo "  make build              - Build the application"
	@echo "  make run                - Run the application"
	@echo "  make test               - Run tests"
	@echo "  make clean              - Clean build artifacts"

install-deps:
	@echo "Installing Go dependencies..."
	go mod download
	go mod tidy

build: install-deps
	@echo "Building the application..."
	go build -o bin/todo-api ./cmd/main.go

run: build
	@echo "Running the application..."
	./bin/todo-api

test:
	@echo "Running tests..."
	go test -v ./...

clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	go clean

compose-up:
	@echo "Starting docker-compose stack..."
	docker-compose up --build

compose-down:
	@echo "Stopping docker-compose stack..."
	docker-compose down
