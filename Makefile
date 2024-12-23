# Description: Makefile for the project

# Variables
.DEFAULT_GOAL := help

# Targets
.PHONY: all
all: ## Install dependencies
	@echo "Installing dependencies..."
	@go mod download

.PHONY: clean
clean: ## Clean the project
	@echo "Cleaning..."
	@rm -rf bin/

.PHONY: test
test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

.PHONY: build
build: ## Build the project
	@echo "Building..."
	@go build -o bin/ ./...

.PHONY: run
run: ## Run the project
	@echo "Running..."
	@go run main.go

.PHONY: coverage
coverage: ## Run coverage
	@echo "Running coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

.PHONY: help
help: ## Show this help
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-10s %s\n", $$1, $$2}' $(MAKEFILE_LIST)


