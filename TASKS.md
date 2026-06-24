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

### T009: Add Unified Response And Error Handling
Status: Completed and verified on 2026-06-23.

Summary:

- Added a unified success/error response envelope through handler response helpers.
- Centralized service-layer error mapping in handler code.
- Preserved HTTP-layer parsing and binding errors as 400 responses while returning the same error envelope shape.
- Restored `POST /tasks` to `201 Created` while using the unified success response.
- Changed `DELETE /tasks/:id` from `204 No Content` to `200 OK` with the unified success response.
- Verified `/health`, task create/list/detail/update/delete, invalid body, invalid query, invalid ID, and not-found paths.

### T010: Implement User Registration
Status: Completed and verified on 2026-06-24.

Summary:

- Added `POST /users/register` for user registration.
- Wired registration through handler, service, and repository boundaries.
- Validated required `name`, `email`, and `password` fields, including whitespace-only values in the service layer.
- Hashed passwords with bcrypt before storing them in PostgreSQL.
- Inserted users with `INSERT ... RETURNING` so the response includes database-generated fields.
- Mapped duplicate email conflicts to a clear `409 Conflict` unified error response.
- Returned a dedicated user response DTO so `PasswordHash` is not exposed to clients.
- Verified `/health`, task listing, registration success, duplicate email, missing field, and whitespace-only validation paths.

## Active Task

### T011: Implement User Login

Objective:

Add login with password verification.

Learner should implement:

- a `POST /users/login` or similar login endpoint
- request validation for required email/password fields
- repository logic for finding a user by email
- bcrypt password verification
- clear client errors for invalid credentials

Agent may provide:

- API contract suggestions
- password verification guidance
- validation and error mapping review
- small isolated examples
- review after implementation

Agent should not:

- generate JWT yet
- protect task routes yet
- add current-user task ownership yet

Acceptance Criteria:

- A registered user can log in with email and password.
- Password verification uses bcrypt against the stored password hash.
- Wrong email or wrong password returns a clear client error without revealing which field was wrong.
- Invalid request bodies and missing fields return clear client errors using the unified error response.
- `go test ./...` still passes.
- existing `/health` and task CRUD endpoints still work.

Skills Practiced:

- authentication
- password verification
- validation
- error handling

## Upcoming Tasks

### T012: Implement JWT Generation And Parsing

Objective:

Generate JWT after login and parse JWT for later protected APIs.

Skills Practiced:

- JWT
- configuration
- security basics

### T013: Add Auth Middleware

Objective:

Protect private routes with JWT middleware.

Skills Practiced:

- middleware
- request context
- authorization
