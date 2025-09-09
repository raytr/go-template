.PHONY: help build run test clean deps migrate-up migrate-down migrate-create dev

# Variables
APP_NAME=go-template
MAIN_PATH=cmd/app/main.go
BUILD_DIR=bin
MIGRATIONS_DIR=migrations
DB_URL=postgresql://myadmin:pswd@127.0.0.1:5432/simple-db?sslmode=disable

# Help command
help:
	@echo "Available commands:"
	@echo "  make build       - Build the application"
	@echo "  make run         - Check migrations and run the application"
	@echo "  make dev         - Run the application in development mode with hot reload"
	@echo "  make test        - Run tests"
	@echo "  make clean       - Clean build artifacts"
	@echo "  make deps        - Download dependencies"
	@echo "  make migrate-up  - Run all migrations"
	@echo "  make migrate-down - Rollback all migrations"
	@echo "  make migrate-create NAME=<name> - Create a new migration"

# Download dependencies
deps:
	go mod download
	go mod tidy

# Build the application
build: deps
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)
	@echo "Build complete: $(BUILD_DIR)/$(APP_NAME)"

# Run the application (checks and applies migrations automatically)
run:
	@echo "Starting $(APP_NAME) with migration check..."
	@go run $(MAIN_PATH)

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -cover ./...
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html
	@echo "Clean complete"

# Run database migrations up
migrate-up:
	@echo "Running migrations up..."
	@if command -v migrate > /dev/null; then \
		migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up; \
	else \
		echo "golang-migrate is not installed. Please install it first:"; \
		echo "brew install golang-migrate"; \
	fi

# Run database migrations down
migrate-down:
	@echo "Running migrations down..."
	@if command -v migrate > /dev/null; then \
		migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down; \
	else \
		echo "golang-migrate is not installed. Please install it first:"; \
		echo "brew install golang-migrate"; \
	fi

# Create a new migration
migrate-create:
	@if [ "$(NAME)" = "" ]; then \
		echo "Please provide a migration name: make migrate-create NAME=<name>"; \
		exit 1; \
	fi
	@if command -v migrate > /dev/null; then \
		migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME); \
		echo "Migration created: $(MIGRATIONS_DIR)"; \
	else \
		echo "golang-migrate is not installed. Please install it first:"; \
		echo "brew install golang-migrate"; \
	fi

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...
	@echo "Code formatted"

# Run linters
lint:
	@echo "Running linters..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint is not installed. Installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
		golangci-lint run; \
	fi
