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

## Active Task

### T008: Implement Task Detail, Update, And Delete

Objective:

Implement task detail lookup, update, and deletion.

Learner should implement:

- route parameters for task ID
- handler entrypoints for detail, update, and delete
- service/repository boundaries needed for lookup, update, and delete
- database queries for `SELECT`, `UPDATE`, and `DELETE`
- clear behavior for missing tasks and invalid IDs
- status code choices for successful update and delete

Agent may provide:

- route and package boundary suggestions
- request/response contract review
- repository query examples
- SQL `SELECT`, `UPDATE`, and `DELETE` explanation
- small isolated examples
- review after implementation

Agent should not:

- implement the full handler/service/repository flow unless explicitly requested
- implement authentication or current-user ownership yet
- implement unified error handling yet

Acceptance Criteria:

- `GET /tasks/:id` route exists.
- `PUT` or `PATCH /tasks/:id` route exists for updating a task.
- `DELETE /tasks/:id` route exists.
- invalid task IDs return a clear client error.
- missing tasks return a clear not-found response.
- valid detail request reads one task from PostgreSQL.
- valid update request persists changes in PostgreSQL.
- valid delete request removes the task from PostgreSQL or marks it deleted, depending on the chosen design.
- code follows handler/service/repository package boundaries.
- `go test ./...` still passes.
- existing `/health` endpoint still works.
- authentication remains out of scope.

Skills Practiced:

- REST API
- handler/service/repository boundaries
- SQL update/delete
- error handling
- status code design
- validation

## Upcoming Tasks

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
