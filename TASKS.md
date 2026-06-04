# Tasks

## Active Task

### T001: Initialize Go Gin Project

Objective:

Create a minimal Go Gin backend service with a health check endpoint.

Learner should implement:

- Go module initialization
- `cmd/server/main.go`
- Gin router
- `/health` endpoint
- local run command

Agent may provide:

- suggested directory structure
- explanation of Gin routing
- small isolated examples
- review after implementation

Agent should not:

- write the entire final implementation unless explicitly requested

Acceptance Criteria:

- `go run ./cmd/server` starts the server.
- `GET /health` returns HTTP 200.
- The response body includes a simple status field.
- `README.md` includes the local run command.

Skills Practiced:

- Gin
- HTTP
- REST API
- Go `struct`
- basic Go project structure

## Upcoming Tasks

### T002: Add Basic Project Structure

Objective:

Introduce initial backend package boundaries.

Expected areas:

- `internal/config`
- `internal/handler`
- `internal/router`
- `internal/service`
- `internal/repository`
- `internal/model`

Skills Practiced:

- Go package boundaries
- `struct`
- `interface`
- responsibility separation

### T003: Add PostgreSQL With Docker Compose

Objective:

Run PostgreSQL locally and connect the Go service to it.

Expected areas:

- `docker-compose.yml`
- database configuration
- connection initialization
- health check for database connectivity

Skills Practiced:

- PostgreSQL
- SQL
- Docker
- configuration management

### T004: Add Task CRUD

Objective:

Implement task creation, listing, updating, and deletion.

Skills Practiced:

- REST API
- database modeling
- GORM or `sqlc`
- error handling
- testing basics

### T005: Add User Registration And Login

Objective:

Implement user registration, password hashing, login, and JWT generation.

Skills Practiced:

- authentication
- JWT
- password hashing
- error handling
- security basics

### T006: Add Auth Middleware

Objective:

Protect task APIs so users can only access their own data.

Skills Practiced:

- middleware
- JWT validation
- authorization
- request context

### T007: Add Tests And Documentation

Objective:

Add basic tests and complete local run documentation.

Skills Practiced:

- testing
- API documentation
- Docker
- README writing
