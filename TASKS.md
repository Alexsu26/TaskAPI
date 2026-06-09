# Tasks

## Active Task

### T004: Add PostgreSQL With Docker Compose

Objective:

Run PostgreSQL locally and connect the Go service to it.

Learner should implement:

- `docker-compose.yml` with PostgreSQL
- database configuration usage from `internal/config`
- connection initialization in a clear package boundary
- startup code that fails clearly if the database connection cannot be initialized
- a health-related check or startup evidence that confirms the database connection works

Agent may provide:

- suggested Docker Compose structure
- explanation of PostgreSQL environment variables
- database connection package boundary suggestions
- small isolated examples
- review after implementation

Agent should not:

- implement task CRUD
- implement user/auth code
- add migrations yet unless explicitly discussed

Acceptance Criteria:

- `docker compose up` starts PostgreSQL.
- the Go service can initialize a database connection using configuration values.
- startup errors for database connection are handled explicitly.
- `go run ./cmd/server` still starts successfully when PostgreSQL is available.
- `GET /health` still returns HTTP 200.
- no auth or task CRUD code is added yet.

Skills Practiced:

- PostgreSQL
- SQL basics
- Docker Compose
- database connection handling
- error handling

## Upcoming Tasks

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
