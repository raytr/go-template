# Go Template - REST API with Gin Framework

A production-ready Go REST API template using Gin framework, PostgreSQL, and clean architecture principles.

## Features

- 🚀 **Gin Web Framework** - Fast HTTP web framework
- 🗄️ **PostgreSQL with GORM** - Database with ORM support and migrations
- 🔧 **Clean Architecture** - Organized code structure with separation of concerns
- 📝 **CRUD Operations** - Complete user management API
- 🔄 **Database Migrations** - Dual support: SQL migrations with golang-migrate and GORM auto-migration
- ⚙️ **Configuration Management** - Environment-based config with Viper
- 📚 **API Documentation** - Comprehensive API docs

## Project Structure

```
go-template/
├── cmd/                 # Main applications
│   └── app/
│       └── main.go     # Application entry point
├── internal/           # Private application code
│   ├── config/        # Configuration management
│   ├── database/      # Database connection (GORM)
│   ├── handler/       # HTTP handlers (controllers)
│   ├── migration/     # Migration runner
│   ├── model/         # Data models (entities and requests)
│   ├── repository/    # Data access layer (GORM repository)
│   └── service/       # Business logic layer
├── api/               # API documentation
├── migrations/        # Database migration files
├── .env              # Environment variables
├── Makefile          # Build and run commands
└── go.mod            # Go dependencies
```

## Prerequisites

- Go 1.21 or higher
- PostgreSQL database
- Make (optional, for using Makefile commands)

## Installation

**Install dependencies**
   ```bash
   make deps
   # or
   go mod download
   ```

## Running the Application

### Using Make (Recommended)

```bash
# Run with automatic migration check
make run

# Build the application
make build

# Run tests
make test
```

### Manual Commands

```bash
# Run the application
go run cmd/app/main.go

# Build
go build -o bin/app cmd/app/main.go

# Run built binary
./bin/app
```

## Database Migrations

The application automatically checks and applies migrations on startup. You can also manage migrations manually:

```bash
# Apply all migrations
make migrate-up

# Rollback all migrations
make migrate-down

# Create a new migration
make migrate-create NAME=add_new_table
```

## Development

### Code Formatting

```bash
make fmt
```

### Linting

```bash
make lint
```

### Testing

```bash
# Run all tests
make test

# Run with coverage
make test-coverage
```

## Architecture

This project follows clean architecture principles:

1. **Handler Layer** (`internal/handler/`)
   - HTTP request/response handling
   - Input validation
   - Route definitions

2. **Service Layer** (`internal/service/`)
   - Business logic
   - Data validation
   - Transaction orchestration

3. **Repository Layer** (`internal/repository/`)
   - Database operations using GORM
   - Query building with ORM
   - Data persistence

4. **Model Layer** (`internal/model/`)
   - `UserEntity` - GORM entity that maps to database tables
   - `UserReq` - Request structures for API endpoints
   - Response DTOs and transformations

## Technologies Used

- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [GORM](https://gorm.io/) - ORM library for Go
- [PostgreSQL](https://www.postgresql.org/) - Database
- [golang-migrate](https://github.com/golang-migrate/migrate) - Database migrations
- [Viper](https://github.com/spf13/viper) - Configuration management
