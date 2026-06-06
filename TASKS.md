# Tasks

## Active Task

### T002: Add Basic Project Structure

Objective:

Introduce initial backend package boundaries without adding database or business logic yet.

Learner should implement:

- `internal/config`
- `internal/handler`
- `internal/router`
- `internal/service`
- `internal/repository`
- `internal/model`
- move route setup out of `cmd/server/main.go` only as much as needed

Agent may provide:

- suggested directory structure
- explanation of Go package boundaries
- small isolated examples
- review after implementation

Agent should not:

- implement the full package refactor unless explicitly requested

Acceptance Criteria:

- `go run ./cmd/server` still starts the server.
- `GET /health` still returns HTTP 200.
- route registration is separated from `main.go` into a small router or handler package.
- package names are simple and match their directory responsibilities.
- no database, auth, or task CRUD code is added yet.

Skills Practiced:

- Go `struct`
- Go package boundaries
- responsibility separation
- Gin
- HTTP

## Upcoming Tasks

### T003: Add Configuration Management

Objective:

Add environment-based configuration for server and database settings.

Skills Practiced:

- configuration management
- environment variables
- error handling

### T004: Add PostgreSQL With Docker Compose

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

### T005: Design User And Task Models

Objective:

Define initial user and task models and understand how they map to database tables.

Skills Practiced:

- Go `struct`
- database modeling
- SQL schema basics

### T006: Implement Task Creation

Objective:

Implement the API flow for creating a task.

Skills Practiced:

- REST API
- handler/service/repository boundaries
- database insert
- validation

### T007: Implement Task List Query

Objective:

Implement task listing with basic pagination.

Skills Practiced:

- REST API
- SQL query basics
- pagination
- response design

### T008: Implement Task Detail, Update, And Delete

Objective:

Implement task detail lookup, update, and deletion.

Skills Practiced:

- REST API
- SQL update/delete
- error handling
- status code design

### T009: Add Unified Response And Error Handling

Objective:

Standardize API responses and application errors.

Skills Practiced:

- error handling
- HTTP status codes
- response design
