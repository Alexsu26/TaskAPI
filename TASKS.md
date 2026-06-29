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

### T014: Restrict Tasks To The Current User
Status: Completed and verified on 2026-06-26.

Summary:

- Read the authenticated user ID from Gin context in task handlers.
- Passed the current user ID through handler, service, and repository task methods.
- Removed the temporary hard-coded task `UserID: 1`.
- Stored newly created tasks under the authenticated user's ID.
- Filtered task list, detail, update, and delete operations by `user_id`.
- Preserved `404 Not Found` for cross-user task detail, update, and delete attempts so task existence is not exposed.
- Fixed review findings for missing handler `return` statements after auth-context failure and an `UPDATE` SQL column typo.
- Verified `gofmt`, `go test ./...`, `go vet ./...`, and two-user runtime ownership checks.

### T015: Add Basic Tests
Status: Completed and verified on 2026-06-26.

Summary:

- Added focused unit tests for the JWT token manager under `internal/auth`.
- Verified token generation and parsing preserves the expected `user_id`.
- Verified invalid user IDs return `ErrTokenInvalid`.
- Verified malformed token parsing returns `ErrTokenInvalid`.
- Kept the test deterministic with the standard Go `testing` package and no external service dependency.
- Verified `gofmt`, `go test ./...`, and `go vet ./...`.

### T016: Complete Stage 1 Documentation
Status: Completed and verified on 2026-06-26.

Summary:

- Updated `README.md` with local development startup steps.
- Documented current environment variables and useful defaults from `internal/config`.
- Added API examples for `/health`, registration, login, and authenticated task CRUD routes.
- Documented JWT usage through the `Authorization: Bearer <token>` header for protected task routes.
- Included the `go test ./...` command for running tests.
- Fixed review findings so documented pagination, task ID, and PUT examples are runnable and consistent with current Gin routes.
- Verified `go test ./...`.

### T017: Add Database Migrations
Status: Completed and verified on 2026-06-27.

Summary:

- Added versioned SQL migration files under `migrations/`.
- Captured the current `users` and `tasks` schema in `000001_create_users_and_tasks.up.sql`.
- Added a matching `down` migration that drops dependent `tasks` before `users`.
- Removed startup-time dependence on the previous unversioned `database.RunMigrations` helper.
- Documented the `migrate -path migrations ... up` workflow in `README.md`.
- Fixed review finding where `tasks.user_id` was incorrectly declared as `bigserial` instead of `bigint`.
- Verified `go test ./...`, `go vet ./...`, `migrate up`, `migrate down -all`, and the generated PostgreSQL table structure in an isolated temporary database.

### T018: Refactor DTO, Model, And Response Boundaries
Status: Completed and verified on 2026-06-27.

Summary:

- Moved handler request and response DTO definitions into `internal/handler/dto.go`.
- Added explicit user and task response DTOs with JSON tags for stable API field names.
- Converted task create, list, detail, and update responses from `model.Task` to `TaskResponse`.
- Converted user registration and login responses from internal user models to `UserResponse`.
- Kept `PasswordHash`, `UserID`, and model-only fields out of public API responses.
- Preserved existing routes, status codes, handler/service/repository responsibilities, and the unified response envelope.
- Fixed review findings for list response slice construction and a `UpdatedAt` field-name compile error.
- Verified `gofmt -l cmd/server internal`, `go test ./...`, `go vet ./...`, and runtime response shapes against a migrated temporary database.

### T019: Add Structured Logging
Status: Completed and verified on 2026-06-27.

Summary:

- Added a small `internal/logger` package using Go's standard `log/slog` JSON handler.
- Created the structured logger during server startup and injected it into handler and router boundaries.
- Logged server startup with address, database host, database port, database name, and JWT expiration minutes while avoiding secrets.
- Replaced `gin.Default()` with `gin.New()` plus custom request logging and `gin.Recovery()`.
- Added request logs with method, path, status, duration, and client IP.
- Added internal 500 error logs while preserving generic client-facing error responses.
- Fixed review finding by logging `ctx.Request.URL.Path` instead of the full URL to avoid query string leakage.
- Verified `gofmt -l cmd/server internal`, `go test ./...`, `go vet ./...`, and runtime logs for health, unauthenticated task access, and failed login.

### T020: Add Request ID And Panic Recovery
Status: Completed and verified on 2026-06-29.

Summary:

- Added request ID middleware that stores `request_id` in Gin context for later middleware and handlers.
- Returned the request ID through the `X-Request-ID` response header.
- Accepted incoming request IDs only when they are valid UUIDs and generated a new UUID for missing or unsafe values.
- Added request IDs to HTTP request logs and internal 500 error logs.
- Replaced `gin.Recovery()` with a custom panic recovery middleware.
- Logged panic events with request ID, method, path, and panic value while avoiding request bodies, tokens, passwords, query strings, and other sensitive values.
- Preserved the existing generic unified 500 response body for panic responses.
- Verified `gofmt -l cmd/server internal`, `go test ./...`, `go vet ./...`, and runtime request ID behavior on `/health`.

## Active Task

### T021: Add Service Layer Tests

Objective:

Add focused tests for service-layer business rules so important behavior can be checked without starting Gin or PostgreSQL.

Learner should implement:

- choose one small service package target first
- add tests for success and error paths in that service
- use fake or stub repository dependencies instead of a real database where practical
- verify service errors with `errors.Is` when sentinel errors are expected
- keep tests deterministic and runnable through `go test ./...`
- avoid changing production behavior just to make tests pass

Agent may provide:

- guidance on choosing the first narrow service target
- test case suggestions
- examples of small fake repository implementations
- review of test boundaries and error assertions
- verification commands and runtime checks

Agent should not:

- rewrite service logic for the learner
- introduce a large mocking framework before it is needed
- convert this into handler or database integration testing
- add broad flaky tests that depend on external services

Acceptance Criteria:

- At least one service package has focused unit tests.
- Tests cover at least one success path and one validation or business-error path.
- Tests do not require a running HTTP server.
- Tests do not require a running PostgreSQL instance unless explicitly justified.
- Error assertions use stable sentinel errors or stable behavior, not fragile string matching.
- `go test ./...` passes.

Skills Practiced:

- Go testing
- dependency boundaries
- fake implementations
- error handling
