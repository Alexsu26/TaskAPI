# Tasks

## Completed Tasks

### T004: Add PostgreSQL With Docker Compose
Status: Completed and verified on 2026-06-11.

Summary:

- Added PostgreSQL through Docker Compose.
- Added a database package boundary for PostgreSQL initialization.
- Reused `internal/config` database values for connection setup.
- Preserved the existing Gin `/health` endpoint.

## Active Task

### T005: Design User And Task Models

Objective:

Define initial user and task models and understand how they map to database tables.

Learner should implement:

- Go structs for initial user and task models
- field choices that map clearly to future PostgreSQL tables
- basic comments or notes explaining key fields when useful
- no repository, handler, auth, or CRUD implementation yet

Agent may provide:

- model boundary suggestions
- SQL table design explanation
- examples of field naming and timestamp choices
- small isolated examples
- review after implementation

Agent should not:

- implement model code wholesale unless explicitly requested
- implement database migrations yet
- implement task CRUD or authentication flows

Acceptance Criteria:

- user and task model structs exist in the appropriate package.
- model fields are enough for upcoming registration and task CRUD work.
- field names and types are understandable and consistent.
- the design can be mapped to SQL tables.
- `go test ./...` still passes.
- no handler, service, repository, auth, or CRUD code is added yet.

Skills Practiced:

- Go `struct`
- database modeling
- SQL schema basics

## Upcoming Tasks

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

### T010: Implement User Registration

Objective:

Add user registration with password hashing.

Skills Practiced:

- authentication
- password hashing
- validation
- database uniqueness
