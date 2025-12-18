# Go API Gateway - Copilot Instructions

## Project Overview
Go API Gateway is a lightweight Go API built with **chi router** that proxies requests to external APIs (currently dummyjson.com). It's structured using **hexagonal architecture** with clear separation of concerns across handlers, services, repositories, and ports (interfaces).

**Current capabilities**: GET/POST user endpoints that forward requests to dummyjson API.

---

## Architecture & Key Files

### Layered Structure (Hexagonal Architecture)
```
main.go                          → Entry point, chi router setup
├── internal/handlers/           → HTTP request handlers (empty - in development)
├── internal/services/           → Business logic layer (empty - in development)
├── internal/ports/              → Interface definitions (users_repository.go)
├── internal/types/
│   ├── interfaces/              → Service interfaces (users_interface.go)
│   └── types/                   → Data models (users.go: ID, FirstName, LastName, Age, Email)
├── internal/routes/             → Route definitions (empty - in development)
└── internal/middleware/         → Auth/logging middleware (placeholder)

configs/config.go                → Config struct with Port, Env, DbConnectionString, LogLevel, FeatureFlags
pkg/logger/logger.go             → Structured JSON logging with LogLevels: DEBUG, INFO, WARN, ERROR, FATAL
pkg/common-utils/common_utils.go → File I/O utilities (ReadFile, ReadJsonFile)
```

### Data Flow
1. **Chi router** in `main.go` receives HTTP requests
2. **Current handlers** (getUsersHandler, addUserHandler) directly call external APIs
3. **Target refactor**: Route → Handler → Service → Repository (following interface pattern in ports/)
4. **Data models**: `Users` type defined in `internal/types/types/users.go`

---

## Development Patterns

### Logging
Use the **structured JSON logger** from `pkg/logger/`:
```go
log := logger.New("INFO")  // LogLevels: "debug", "info", "warn", "error", "fatal"
log.Info("message", map[string]interface{}{"field": value})
```

### Configuration
Load config via `configs/config.go::Load()` which reads from environment variables:
- `PORT` (default: "8080")
- `LOG_LEVEL` (default: "INFO")
- `DBCREDS` (format: "user:password@/dbname")
- `DEV` (default: "DEV")

### Type Validation
User struct uses struct tags for validation (note: "name" tag on Email field appears to be a bug):
```go
type Users struct {
	ID        int    `json:"id validate:"required"`
	FirstName string `json:"firstName"`
	Email     string `json:"name" validate:"required,min=1,max=100"`  // Should be json:"email"
}
```

---

## Router & Dependencies
- **HTTP Framework**: `github.com/go-chi/chi/v5` with CORS support
- **Middleware**: chi's built-in Logger and Recoverer
- **Current routes** (in main.go):
  - `GET /users` → proxies to https://dummyjson.com/users
  - `POST /users/add` → proxies to https://dummyjson.com/users/add

---

## Common Tasks

### Adding a New Endpoint
1. Define handler in `internal/handlers/` (currently empty)
2. Define service interface in `internal/types/interfaces/`
3. Implement service in `internal/services/`
4. Add route in `internal/routes/routes.go` or inline in main.go

### Working with Configuration
```go
cfg := config.Load()
fmt.Println(cfg.Port, cfg.LogLevel)  // Environment-driven config
```

### Refactoring Main Handlers
The inline handlers in `main.go` (getUsersHandler, addUserHandler) should be moved to `internal/handlers/users_handler.go` and refactored through the service/repository layers defined in the ports pattern.

---

## Environment & Build

**Go Version**: 1.25.4  
**Start Server**: `go run main.go` (listens on port 5200)  
**Module**: `api-gateway`

Environment files in `envs/`:
- `dev_env.json` / `dev_db_creds_dev.json` (dev configuration)
- `qa_env.json` (QA configuration)

---

## Code Quality Notes
- Logger outputs structured JSON suitable for log aggregation
- Feature flags architecture in place (Config.FeatureFlags) but unused
- Request/Response bodies are logged in plain text - consider structured logging for PII
- Panic usage in handlers should be replaced with proper error handling
