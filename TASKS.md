# Tasks

## Active Task

### T003: Add Configuration Management

Objective:

Add environment-based configuration for server and database settings.

Learner should implement:

- a minimal config struct in `internal/config`
- environment variable reading for server port
- sensible defaults for local development
- startup code that uses config instead of hard-coded `:8080`

Agent may provide:

- suggested config fields
- explanation of environment variables
- small isolated examples
- review after implementation

Agent should not:

- implement database connection code yet
- add Docker Compose yet

Acceptance Criteria:

- `go run ./cmd/server` still starts the server.
- `GET /health` still returns HTTP 200.
- server port can be configured through an environment variable.
- local default still works when no environment variable is set.
- startup errors are handled explicitly.
- no database, auth, or task CRUD code is added yet.

Skills Practiced:

- Go `struct`
- configuration management
- environment variables
- error handling
- HTTP

## Upcoming Tasks

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
