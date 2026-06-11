# Tasks

## Completed Tasks

### T004: Add PostgreSQL With Docker Compose
Status: Completed and verified on 2026-06-11.

Summary:

- Added PostgreSQL through Docker Compose.
- Added a database package boundary for PostgreSQL initialization.
- Reused `internal/config` database values for connection setup.
- Preserved the existing Gin `/health` endpoint.

### T005: Design User And Task Models
Status: Completed and verified on 2026-06-11.

Summary:

- Added initial `User` and `Task` structs under `internal/model`.
- Included fields needed for upcoming user registration and task CRUD.
- Aligned `Task.UserID` with `User.ID` so the models map cleanly to future SQL table relationships.
- Preserved existing Gin router and handler behavior without adding CRUD, auth, repository, service, or migration code.

## Active Task

### T006: Implement Task Creation

Objective:

Implement the API flow for creating a task.

Learner should implement:

- request shape for creating a task
- handler entrypoint for task creation
- service/repository boundary needed for task creation
- database insert for a new task
- basic input validation for required fields
- clear error behavior for invalid input and database failures

Agent may provide:

- route and package boundary suggestions
- request/response contract review
- repository interface examples
- SQL insert explanation
- small isolated examples
- review after implementation

Agent should not:

- implement the full handler/service/repository flow unless explicitly requested
- implement list/detail/update/delete endpoints yet
- implement authentication or current-user ownership yet

Acceptance Criteria:

- `POST /tasks` route exists.
- request body includes at least a task title and optional description.
- invalid input returns a clear client error.
- valid input inserts a task into PostgreSQL.
- successful creation returns the created task or its key fields.
- code follows handler/service/repository package boundaries.
- `go test ./...` still passes.
- existing `/health` endpoint still works.
- list/detail/update/delete and authentication remain out of scope.

Skills Practiced:

- REST API
- handler/service/repository boundaries
- database insert
- validation

## Upcoming Tasks

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
