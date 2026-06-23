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

### T006: Implement Task Creation
Status: Completed and verified on 2026-06-13.

Summary:

- Added `POST /tasks` for creating a task.
- Wired task creation through handler, service, and repository boundaries.
- Added PostgreSQL table creation and a temporary dev user for pre-auth task ownership.
- Inserted tasks with `INSERT ... RETURNING` so the response includes database-generated fields.
- Added validation for missing and whitespace-only titles.
- Preserved `/health` and kept list/detail/update/delete/auth out of scope.

### T007: Implement Task List Query
Status: Completed and verified on 2026-06-15.

Summary:

- Added `GET /tasks` for listing tasks.
- Added `limit` and `offset` query parameters with default values.
- Returned clear client errors for non-numeric and invalid pagination values.
- Queried tasks from PostgreSQL with explicit selected columns.
- Used stable ordering with `updated_at DESC, id DESC`.
- Preserved `/health` and kept detail/update/delete/auth out of scope.

### T008: Implement Task Detail, Update, And Delete
Status: Completed and verified on 2026-06-23.

Summary:

- Added `GET /tasks/:id`, `PUT /tasks/:id`, and `DELETE /tasks/:id` routes.
- Wired detail, update, and delete through handler, service, and repository boundaries.
- Used `SELECT ... WHERE id = $1` with `QueryRow` + `Scan` for detail lookup.
- Used `UPDATE ... SET ... WHERE id = $1 RETURNING ...` for update with full row return.
- Used `DELETE FROM tasks WHERE id = $1` with `RowsAffected` check for delete.
- Mapped `sql.ErrNoRows` and `RowsAffected == 0` to `ErrTaskNotFound` → HTTP 404.
- Refactored pagination: handler parses strings, service owns default values and range validation using `*int` pointers.
- Deferred minor improvements to backlog: ID validation consistency, JSON tags, status whitelist, error message leakage.

## Active Task

### T009: Add Unified Response And Error Handling

Objective:

Standardize API responses and application errors.

Learner should implement:

- a unified response wrapper for successful and error responses
- a centralized error handler that maps service-layer errors to HTTP status codes
- consistent error message format across all existing endpoints
- removal of inline error-to-status-code mapping from individual handlers

Agent may provide:

- response envelope design suggestions
- error mapping strategy review
- small isolated examples
- review after implementation

Agent should not:

- implement authentication or current-user ownership yet
- add new business endpoints

Acceptance Criteria:

- All endpoints return a consistent JSON response shape.
- All errors use a centralized mapping instead of inline `if/else` chains in each handler.
- `go test ./...` still passes.
- existing `/health`, `POST /tasks`, `GET /tasks`, `GET /tasks/:id`, `PUT /tasks/:id`, `DELETE /tasks/:id` endpoints still work.

Skills Practiced:

- error handling
- HTTP status codes
- response design

## Upcoming Tasks

### T010: Implement User Registration

Objective:

Add user registration with password hashing.

Skills Practiced:

- authentication
- password hashing
- validation
- database uniqueness

### T011: Implement User Login

Objective:

Add login with password verification.

Skills Practiced:

- authentication
- error handling
- security basics
