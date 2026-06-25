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

### T011: Implement User Login
Status: Completed and verified on 2026-06-24.

Summary:

- Added `POST /users/login` for user login.
- Wired login through handler, service, and repository boundaries.
- Added repository lookup by email with `sql.ErrNoRows` mapped to a repository-level not-found error.
- Validated required `email` and `password` fields, including whitespace-only values in the service layer.
- Verified passwords with `bcrypt.CompareHashAndPassword` against the stored password hash.
- Returned a unified `401 Unauthorized` error for wrong email and wrong password without revealing which field was wrong.
- Returned a user response DTO without `PasswordHash`.
- Verified `/health`, task listing, registration success, login success, wrong password, wrong email, missing password, and whitespace-only password paths.

### T012: Implement JWT Generation And Parsing
Status: Completed and verified on 2026-06-25.

Summary:

- Added JWT configuration values for token secret and expiration minutes.
- Added an `internal/auth` token manager for JWT generation and parsing.
- Generated signed JWTs with `HS256`, `user_id`, `exp`, and `iat` claims.
- Added parsing logic that validates signing method, signature, expiration, and positive user ID.
- Injected the token manager through startup into the user service instead of reading config directly from auth helpers.
- Updated successful login responses to include both the user DTO and JWT token without exposing `PasswordHash`.
- Verified `/health`, task listing, registration success, login success with token, and wrong-password login behavior.

### T013: Add Auth Middleware
Status: Completed and verified on 2026-06-25.

Summary:

- Added an `internal/middleware` auth middleware boundary.
- Read the `Authorization` header and parsed `Bearer <token>` values.
- Validated JWTs through the existing `auth.TokenManager`.
- Returned unified `401 Unauthorized` responses for missing, malformed, invalid, and expired tokens.
- Stored the authenticated user ID in Gin context as `current_user_id` for later use.
- Protected task routes through a Gin route group while keeping `/health`, `/users/register`, and `/users/login` public.
- Fixed the initial import-cycle issue by keeping the route registration helper type out of the router package dependency path.
- Verified `gofmt`, `go test ./...`, `go vet ./...`, and runtime auth middleware behavior.

## Active Task

### T014: Restrict Tasks To The Current User

Objective:

Ensure users can only access their own tasks.

Learner should implement:

- read the current user ID from Gin context in task handlers
- pass the current user ID from handler to task service methods
- pass the current user ID from service to repository methods
- filter task create/list/detail/update/delete operations by authenticated user
- remove the temporary hard-coded task `UserID: 1`
- return `404 Not Found` when a task exists but does not belong to the current user

Agent may provide:

- handler/service/repository boundary guidance
- SQL filtering guidance
- current-user context retrieval examples
- authorization review
- small isolated examples
- review after implementation

Agent should not:

- implement the full task ownership flow unless explicitly asked
- add role-based permission systems
- add organization/team ownership
- add tests beyond focused examples unless requested

Acceptance Criteria:

- Creating a task stores the authenticated user's ID instead of the temporary hard-coded user ID.
- Listing tasks returns only tasks owned by the authenticated user.
- Getting task detail returns the task only when it belongs to the authenticated user.
- Updating a task only updates tasks owned by the authenticated user.
- Deleting a task only deletes tasks owned by the authenticated user.
- Accessing another user's task returns `404 Not Found` instead of exposing its existence.
- Public `/health`, user registration, and user login endpoints still work.
- Protected task routes still reject unauthenticated requests.
- `go test ./...` still passes.

Skills Practiced:

- authorization
- request context
- SQL filtering
- service-layer validation

## Upcoming Tasks

### T015: Add Basic Tests

Objective:

Add tests for the most important service or handler behavior.

Skills Practiced:

- Go testing
- test data setup
- behavior verification
