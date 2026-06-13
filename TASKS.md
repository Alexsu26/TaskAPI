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

## Active Task

### T007: Implement Task List Query

Objective:

Implement task listing with basic pagination.

Learner should implement:

- request query parameters for pagination
- handler entrypoint for listing tasks
- service/repository boundary needed for list query
- database query for fetching tasks
- stable ordering for list results
- clear behavior for invalid pagination input

Agent may provide:

- route and package boundary suggestions
- request/query contract review
- repository query examples
- SQL `SELECT`, `LIMIT`, and `OFFSET` explanation
- small isolated examples
- review after implementation

Agent should not:

- implement the full handler/service/repository flow unless explicitly requested
- implement detail/update/delete endpoints yet
- implement authentication or current-user ownership yet

Acceptance Criteria:

- `GET /tasks` route exists.
- query supports basic pagination, such as `limit` and `offset` or `page` and `page_size`.
- invalid pagination input returns a clear client error.
- valid request reads tasks from PostgreSQL.
- results are returned in a stable order.
- code follows handler/service/repository package boundaries.
- `go test ./...` still passes.
- existing `/health` endpoint still works.
- detail/update/delete and authentication remain out of scope.

Skills Practiced:

- REST API
- handler/service/repository boundaries
- SQL query basics
- pagination
- response design
- validation

## Upcoming Tasks

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
