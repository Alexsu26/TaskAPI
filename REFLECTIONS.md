# Reflections

This file records learner performance after each completed task.

## Template

### YYYY-MM-DD: Task Title

Task:

- What was implemented or reviewed.

What went well:

- Concrete strengths shown in this task.

Weak areas:

- Concepts, implementation details, or habits that need more practice.

Next improvement:

- One or two concrete improvements for the next task.

Evidence:

- Commands run, tests passed, or review files created.

### 2026-06-25: T013 Add Auth Middleware

Task:

- Implemented and reviewed JWT auth middleware for protected task routes.

What went well:

- Added a focused `internal/middleware` package boundary for authentication.
- Used a Gin route group to protect `/tasks` routes while keeping `/health`, `/users/register`, and `/users/login` public.
- Reused the existing token manager instead of duplicating JWT parsing logic.
- Returned unified `401 Unauthorized` responses for missing, malformed, invalid, and expired tokens.
- Stored `current_user_id` in Gin context, preparing the code for T014 current-user ownership filtering.

Weak areas:

- The first implementation introduced an import cycle by making `handler` import `router` for a route registration interface.
- Needs more practice designing one-way package dependencies before adding helper interfaces.
- The context key is currently a string literal; a shared constant would reduce typo risk in T014.

Next improvement:

- In T014, focus on reading `current_user_id` from Gin context and passing it through handler, service, and repository boundaries.
- Replace the temporary hard-coded task `UserID: 1` with the authenticated user ID.

Evidence:

- `gofmt -l cmd/server internal` produced no output.
- `go test ./...` passed for all packages after the import cycle was fixed.
- `go vet ./...` produced no output.
- `docker compose up -d postgres` confirmed PostgreSQL was running.
- `SERVER_PORT=18080 JWT_SECRET=acceptance-secret JWT_EXPIRATION_MINUTES=60 go run ./cmd/server` started the service.
- `GET /health` returned 200 without authentication.
- `POST /users/register` returned 201 without authentication.
- `POST /users/login` returned 200 and returned a JWT token.
- `GET /tasks?limit=1&offset=0` without `Authorization` returned 401.
- `GET /tasks?limit=1&offset=0` with malformed `Authorization` returned 401.
- `GET /tasks?limit=1&offset=0` with an invalid token returned 401.
- `GET /tasks?limit=1&offset=0` with an expired signed token returned 401.
- `GET /tasks?limit=1&offset=0` with a valid token returned 200.
- Review record: `reviews/2026-06-25-t013-auth-middleware.md`.

### 2026-06-25: T012 Implement JWT Generation And Parsing

Task:

- Implemented and reviewed JWT configuration, token generation, token parsing, and login token response.

What went well:

- Kept JWT code in a dedicated `internal/auth` boundary instead of mixing signing logic into handlers.
- Passed JWT secret and expiration from config through startup into the token manager, which keeps configuration ownership clear.
- Used `jwt.RegisteredClaims` with `exp` and `iat`, and included `user_id` for later middleware.
- Added signing-method validation, expired-token detection, and positive user ID validation during parsing.
- Corrected response envelope nesting after review so registration and login responses stay consistent.

Weak areas:

- Needed guidance on JWT concepts, `IssuedAt`, signing, and the relationship between token helpers and business validation.
- First response-shape attempt added an extra nested `data` object and changed registration response shape unintentionally.
- Token generation initially needed a reminder to reject invalid user IDs at the generation boundary.

Next improvement:

- In T013, focus on Gin middleware flow, `Authorization: Bearer <token>` parsing, and storing the current user ID in request context.
- Keep auth middleware responsible for request authentication only; defer task ownership filtering to T014.

Evidence:

- `gofmt -l cmd/server internal` produced no output.
- `go test ./...` passed for all packages.
- `go vet ./...` produced no output.
- `docker compose up -d postgres` confirmed PostgreSQL was running.
- `SERVER_PORT=18080 JWT_SECRET=acceptance-secret JWT_EXPIRATION_MINUTES=60 go run ./cmd/server` started the service.
- `GET /health` returned 200.
- `POST /users/register` returned 201 with a user DTO and no `PasswordHash`.
- `POST /users/login` returned 200 with a user DTO and JWT token.
- The login token had three JWT segments and decoded payload fields `user_id`, `exp`, and `iat`.
- `POST /users/login` with wrong password returned 401.
- `GET /tasks?limit=1&offset=0` returned 200.
- Review record: `reviews/2026-06-25-t012-jwt-generation-and-parsing.md`.

### 2026-06-05: T001 Initialize Go Gin Project

Task:

- Implemented and reviewed a minimal Go Gin service with a `/health` endpoint and local run documentation.

What went well:

- Placed the Go module at the repository root after identifying the nested module issue.
- Created `cmd/server/main.go` and used Gin to register a minimal health check route.
- Added README commands for starting the server and checking `/health`.

Weak areas:

- Initial server code registered the route but did not call `Run`, so the program exited immediately.
- README documentation was added after review rather than in the first implementation pass.
- `r.Run(":8080")` startup errors are not handled yet; this is acceptable for T001 but should be practiced later.

Next improvement:

- In T002, focus on package boundaries and keeping `main.go` small without adding unrelated features.
- Practice checking acceptance criteria before asking for review.

Evidence:

- `gofmt -l cmd/server/main.go` produced no output.
- `go test ./...` passed with `?    taskapi/cmd/server    [no test files]`.
- `go run ./cmd/server` started the Gin server on `:8080`.
- `curl -i http://localhost:8080/health` returned `HTTP/1.1 200 OK` and `{"status":"ok"}`.
- Review record: `reviews/2026-06-05-t001-initialize-go-gin-project.md`.

### 2026-06-06: T002 Add Basic Project Structure

Task:

- Refactored the minimal Gin service into basic internal package boundaries.
- Moved route setup out of `cmd/server/main.go` into `internal/router` and `internal/handler`.
- Added durable placeholder packages for `config`, `model`, `repository`, and `service` using minimal `doc.go` files.

What went well:

- Kept the task scope narrow and did not add database, auth, or task CRUD early.
- Used `internal/router.SetupRouter` to keep `main.go` focused on program startup.
- Renamed the health route function to `RegisterHealthRoutes`, which better matches its responsibility.
- Responded to review feedback by making empty package-boundary directories trackable and visible to Go tooling.

Weak areas:

- The first T002 review missed that empty directories are not durable Go packages.
- `r.Run(":8080")` still ignores the returned error; this should be addressed in T003 when configuration is introduced.

Next improvement:

- In T003, practice defining a small config struct and using environment variables without over-expanding into database setup.
- Handle startup errors explicitly before moving on to database work.

Evidence:

- `gofmt -l cmd/server internal` produced no output.
- `go test ./...` passed for `cmd/server` and all current `internal` packages.
- `go run ./cmd/server` started the Gin server on `:8080`.
- `curl -i http://localhost:8080/health` returned `HTTP/1.1 200 OK` and `{"status":"ok"}`.
- Review record: `reviews/2026-06-06-t002-add-basic-project-structure.md`.

### 2026-06-09: T003 Add Configuration Management

Task:

- Added environment-based configuration for server and database settings.
- Replaced hard-coded `:8080` startup with configuration-driven server port selection.
- Added explicit startup error handling around configuration loading and Gin startup.

What went well:

- Kept `cmd/server/main.go` focused on startup by moving configuration into `internal/config`.
- Preserved the existing `internal/router` and `internal/handler` boundaries.
- Avoided adding database connection, auth, or task CRUD code before the relevant tasks.
- Responded to review feedback by completing all defined database configuration fields.

Weak areas:

- The first implementation defined database config fields before loading all of them.
- Config defaults should be chosen deliberately; the next task should align database defaults with the Docker Compose service.

Next improvement:

- In T004, practice connecting PostgreSQL without leaking database setup into handlers.
- Keep startup failure paths explicit when introducing database initialization.

Evidence:

- `gofmt -l cmd/server internal` produced no output.
- `go test ./...` passed for all current packages.
- `go run ./cmd/server` started the Gin server on `:8080`.
- `curl -i http://localhost:8080/health` returned `HTTP/1.1 200 OK` and `{"status":"ok"}`.
- `SERVER_PORT=8090 go run ./cmd/server` started the Gin server on `:8090`.
- `curl -i http://localhost:8090/health` returned `HTTP/1.1 200 OK` and `{"status":"ok"}`.
- `SERVER_PORT=abc go run ./cmd/server` produced an explicit startup error.
- Review record: `reviews/2026-06-09-t003-configuration-management.md`.

### 2026-06-11: T004 Add PostgreSQL With Docker Compose

Task:

- Added PostgreSQL through Docker Compose and connected the Go service to it during startup.
- Added a database package boundary using `database/sql` and pgx.
- Kept the existing Gin router and `/health` endpoint behavior unchanged.

What went well:

- Chose Compose database credentials that match the application configuration defaults.
- Kept database initialization in `internal/database` instead of mixing it into handlers.
- Wrapped startup failures clearly through `init postgres db` and `ping postgres db`.
- Stayed within T004 scope and did not add task CRUD, auth, or migrations early.

Weak areas:

- The initialized `*sql.DB` is currently not retained or closed by the application. This is acceptable for T004 startup validation, but future repository work needs an explicit ownership pattern for the database handle.
- There is no health endpoint database check yet. For this task, startup `Ping` evidence is enough, but later health design should distinguish service liveness from database readiness.

Next improvement:

- In T005, focus on simple model fields that map cleanly to SQL tables without starting CRUD implementation early.
- Before T006, decide how the database handle should be passed into repositories.

Evidence:

- `gofmt -l cmd/server internal` produced no output.
- `go test ./...` passed for all current packages.
- `docker compose config` parsed successfully.
- `docker compose up -d` started the PostgreSQL container.
- `docker compose exec -T postgres pg_isready -U taskapi -d taskapi` reported accepting connections.
- `go run ./cmd/server` started the Gin server on `:8080` while PostgreSQL was available.
- `curl -i http://localhost:8080/health` returned `HTTP/1.1 200 OK` and `{"status":"ok"}`.
- `DATABASE_PORT=15432 go run ./cmd/server` produced an explicit database startup failure.
- Review record: `reviews/2026-06-11-t004-postgresql-docker-compose.md`.

### 2026-06-11: T005 Design User And Task Models

Task:

- Added initial `User` and `Task` model structs under `internal/model`.
- Designed fields for upcoming user registration and task CRUD work.
- Kept the implementation limited to models without adding handlers, services, repositories, migrations, auth, or CRUD code.

What went well:

- Used the existing `internal/model` package boundary correctly.
- Included core model fields such as IDs, timestamps, `Email`, `PasswordHash`, task ownership, title, description, and status.
- Corrected `Task.UserID` to match the `User.ID` type, making the future SQL foreign key relationship clearer.
- Added `User.Name`, which is reasonable for future registration and display needs.

Weak areas:

- The first pass used different types for `User.ID` and `Task.UserID`, which would have made future SQL relationships awkward.
- Status is currently a free-form string; later tasks should constrain allowed values through validation or constants.

Next improvement:

- In T006, implement task creation while keeping handler, service, and repository responsibilities separate.
- Decide how the application should retain and pass the existing `*sql.DB` handle into repository code.

Evidence:

- `gofmt -l cmd/server internal` produced no output.
- `go test ./...` passed for all current packages.
- Review record: `reviews/2026-06-11-t005-design-user-and-task-models.md`.

### 2026-06-13: T006 Implement Task Creation

Task:

- Implemented and reviewed `POST /tasks` for task creation.
- Wired request handling through handler, service, and repository boundaries.
- Added PostgreSQL table setup and a temporary dev user strategy for pre-auth task ownership.

What went well:

- Preserved the existing `/health` endpoint while adding the create route.
- Passed the `*sql.DB` dependency from startup into repository code instead of using globals.
- Used Gin JSON binding for request parsing and separated validation errors from internal/database errors.
- Used `INSERT ... RETURNING` so the API response includes database-generated fields.
- Kept list/detail/update/delete and authentication out of scope.

Weak areas:

- Initial handler code had syntax and control-flow issues around `if err` and missing `return`.
- The relationship between `Exec`, `QueryRow`, and PostgreSQL `RETURNING` needed explanation.
- Pre-auth ownership required an explicit temporary strategy to avoid foreign-key failures.

Next improvement:

- In T007, focus on SQL read queries, stable ordering, and pagination input validation.
- Continue checking handler control flow carefully after every `ctx.JSON` error response.

Evidence:

- `gofmt -l cmd/server internal` produced no output.
- `go test ./...` passed for all current packages.
- `docker compose up -d` started the PostgreSQL container.
- `go run ./cmd/server` started the Gin server on `:8080`.
- `curl -i http://localhost:8080/health` returned `HTTP/1.1 200 OK` and `{"status":"ok"}`.
- `POST /tasks` without `title` returned `HTTP/1.1 400 Bad Request`.
- `POST /tasks` with whitespace-only `title` returned `HTTP/1.1 400 Bad Request`.
- `POST /tasks` with a valid title returned `HTTP/1.1 201 Created` and included `ID`, `CreatedAt`, and `UpdatedAt`.
- Review record: `reviews/2026-06-13-t006-implement-task-creation.md`.

### 2026-06-15: T007 Implement Task List Query

Task:

- Implemented and reviewed `GET /tasks` with `limit` and `offset` pagination.
- Wired the list flow through handler, service, repository, and PostgreSQL.

What went well:

- Preserved the existing handler/service/repository boundaries.
- Used `db.Query`, `rows.Next`, `rows.Scan`, `rows.Err`, and `rows.Close` correctly for multi-row SQL reads.
- Avoided `select *` and scanned explicit columns into `model.Task`.
- Added stable ordering with `updated_at DESC, id DESC`.
- Returned 400 for invalid pagination input and preserved `/health`.

Weak areas:

- The first implementation attempts confused `Exec`, `Query`, and result scanning.
- Go variable scope around `:=` inside `if` blocks needed clarification.
- Error naming and response text such as `ParaInvalid` and `error parameters` should be made clearer in a later cleanup.

Next improvement:

- In T008, focus on detail/update/delete status codes, especially distinguishing invalid ID, not found, and internal database errors.
- Continue keeping successful and error handler control flow explicit with `return` after each error response.

Evidence:

- `gofmt -l cmd/server internal` produced no output.
- `go test ./...` passed for all current packages.
- `SERVER_PORT=18080 go run ./cmd/server` started the current source.
- `GET /health` returned `HTTP/1.1 200 OK`.
- `GET /tasks` returned `HTTP/1.1 200 OK` with task data from PostgreSQL.
- `GET /tasks?limit=10&offset=0` returned `HTTP/1.1 200 OK`.
- `GET /tasks?limit=1&offset=1` returned one later page result, showing offset behavior.
- `GET /tasks?limit=-1&offset=0` returned `HTTP/1.1 400 Bad Request`.
- `GET /tasks?limit=abc&offset=0` returned `HTTP/1.1 400 Bad Request`.
- Review record: `reviews/2026-06-15-t007-implement-task-list-query.md`.

### 2026-06-23: T008 Implement Task Detail, Update, And Delete

Task:

- Implemented and reviewed `GET /tasks/:id`, `PUT /tasks/:id`, and `DELETE /tasks/:id`.
- Wired detail, update, and delete through handler, service, and repository boundaries.
- Refactored pagination to move default values and range validation from handler to service.

What went well:

- Asked strong architectural questions about layer responsibilities — specifically whether query parameter parsing belongs in handler or service, and whether three-layer separation adds value when errors flow through all layers.
- Refactored pagination correctly using `*int` pointers to distinguish "not provided" from "explicitly zero", moving business policy (defaults, limits) to the service layer.
- Three-layer error mapping is clean: repository `ErrTaskNotFound` → service translation → handler HTTP 404.
- Correctly handled `sql.ErrNoRows`, `RowsAffected`, and `UPDATE ... RETURNING` patterns.
- Responded to iterative review feedback and fixed all critical bugs across two rounds.

Weak areas:

- Route registration in `SetupRouter` was missed twice — once for `GetByID`, once for `Update`/`Delete`. Need to build a habit of verifying router wiring before testing.
- Copy-paste errors (JSON tag `json:"title"` for Description field, hardcoded `"id"` string instead of `ctx.Param("id")`) suggest need for more self-review before requesting external review.
- Nil pointer panic from named return value `task *model.Task` shows Go pointer semantics still need reinforcement — named return pointers start as `nil` and must be allocated before use.

Next improvement:

- In T009, focus on removing inline error-to-status-code mapping from individual handlers and centralizing response/error handling.
- Before requesting review, self-check: all routes registered in `SetupRouter`, all `Scan` arguments match `SELECT` columns, all JSON tags correct, all `ctx.Param` calls use actual parameter.

Evidence:

- `go build ./...` succeeded.
- `go vet ./...` succeeded.
- `go test ./...` passed for all packages.
- Review record: `reviews/2026-06-23-t008-task-crud-detail-update-delete.md`.

### 2026-06-23: T009 Add Unified Response And Error Handling

Task:

- Implemented and reviewed unified success/error response helpers.
- Centralized service-layer error mapping in handler code.
- Updated existing `/health` and task CRUD endpoints to use the unified response shape.

What went well:

- Correctly identified that HTTP status code constants such as `http.StatusCreated` can be passed as `int` values into response helpers.
- Restored `POST /tasks` to `201 Created` after the first review caught the accidental downgrade to `200 OK`.
- Separated HTTP parsing/binding errors from service-layer sentinel errors after review feedback.
- Final runtime checks showed a consistent error envelope for invalid body, invalid query, invalid ID, and not-found cases.

Weak areas:

- First implementation passed Gin binding and `strconv` parsing errors into the service error mapper, which caused client errors to become 500s.
- Second implementation fixed status codes but left some HTTP-layer errors in the old `{"error": "..."}` response shape.
- This task shows why response consistency needs runtime checks, not only compilation and unit-style command checks.

Next improvement:

- In T010, design user registration errors before implementing: invalid body, missing fields, duplicate email, hashing/storage failure.
- Keep the distinction clear: handler handles HTTP parsing and response shape; service handles business validation; repository handles SQL.

Evidence:

- `gofmt -l cmd/server internal` produced no output.
- `go test ./...` passed for all packages.
- `go vet ./...` succeeded.
- `docker compose up -d postgres` confirmed PostgreSQL was running.
- `SERVER_PORT=18080 go run ./cmd/server` started the current source.
- `GET /health` returned `HTTP/1.1 200 OK` with unified success response.
- Invalid `POST /tasks` body returned `HTTP/1.1 400 Bad Request` with unified error response.
- `GET /tasks?limit=abc` returned `HTTP/1.1 400 Bad Request` with unified error response.
- `GET /tasks/nope` returned `HTTP/1.1 400 Bad Request` with unified error response.
- `GET /tasks/999999999` returned `HTTP/1.1 404 Not Found` with unified error response.
- Valid `POST /tasks` returned `HTTP/1.1 201 Created` with unified success response.
- Valid list, update, and delete requests returned unified success responses.
- Review record: `reviews/2026-06-23-t009-unified-response-error-handling.md`.

### 2026-06-24: T010 Implement User Registration

Task:

- Implemented and reviewed `POST /users/register` with password hashing.
- Wired registration through handler, service, repository, and PostgreSQL.
- Added duplicate email handling and unified error responses for invalid registration requests.

What went well:

- Preserved the existing handler/service/repository structure while adding user registration.
- Used bcrypt for password hashing before persistence.
- Mapped PostgreSQL unique email conflicts to a domain error and then to HTTP 409.
- Responded to review feedback by replacing direct `model.User` response output with a DTO that does not expose `PasswordHash`.
- Fixed the missing `return` after writing an error response, eliminating the double-write bug.

Weak areas:

- The first successful registration response exposed `PasswordHash`, showing that response DTO boundaries need attention when models contain sensitive fields.
- The first `ErrParaMiss` mapping forgot to return after `ctx.JSON`, causing two JSON error bodies to be written for one request.
- Missing/invalid request behavior still needs runtime checks, not only compile-time checks.

Next improvement:

- In T011, focus on password verification with bcrypt and avoid leaking whether email or password was wrong.
- Continue self-checking every handler error branch for `return` after writing a response.

Evidence:

- `gofmt -l cmd/server internal` produced no output.
- `go test ./...` passed for all packages.
- `go vet ./...` succeeded.
- `docker compose up -d postgres` confirmed PostgreSQL was running.
- `SERVER_PORT=18080 go run ./cmd/server` started the current source.
- `GET /health` returned `HTTP/1.1 200 OK`.
- Valid `POST /users/register` returned `HTTP/1.1 201 Created` without `PasswordHash`.
- Duplicate email registration returned `HTTP/1.1 409 Conflict`.
- Missing password returned `HTTP/1.1 400 Bad Request`.
- Whitespace-only name returned `HTTP/1.1 400 Bad Request` with a single unified error response.
- `GET /tasks?limit=1&offset=0` returned `HTTP/1.1 200 OK`.
- Review record: `reviews/2026-06-24-t010-user-registration.md`.

### 2026-06-24: T011 Implement User Login

Task:

- Implemented and reviewed `POST /users/login` with password verification.
- Added repository lookup by email and service-layer credential validation.
- Kept JWT generation, route protection, and task ownership changes out of scope for the next tasks.

What went well:

- Preserved the existing handler/service/repository boundaries while adding login.
- Correctly added `FindByEmail` in the repository and mapped `sql.ErrNoRows` to a repository-level not-found error.
- After review, replaced incorrect bcrypt hash string comparison with `bcrypt.CompareHashAndPassword`.
- After review, mapped wrong email and wrong password to the same `401 Unauthorized` response so the API does not reveal which credential was wrong.
- Login success returns a user DTO without `PasswordHash`.

Weak areas:

- The first implementation re-hashed the submitted password and compared hash strings. bcrypt hashes include salt, so the same password does not produce the same hash string each time.
- The first error mapping returned `missing parameter` for invalid credentials, which made authentication failures look like request-shape errors.
- One handler error branch initially missed `return` after writing `ctx.JSON`, repeating a control-flow issue seen in T010.

Next improvement:

- In T012, focus on JWT claim design, config loading for token secret/expiration, and parsing errors.
- Continue self-checking auth code for three things before review: no sensitive fields in responses, same response for wrong email/wrong password, and every handler error response returns immediately.

Evidence:

- `gofmt -l cmd/server internal` produced no output.
- `go test ./...` passed for all packages.
- `go vet ./...` succeeded.
- `docker compose up -d postgres` confirmed PostgreSQL was running.
- `SERVER_PORT=18080 go run ./cmd/server` started the current source.
- `GET /health` returned `HTTP/1.1 200 OK`.
- Valid `POST /users/register` returned `HTTP/1.1 201 Created`.
- Valid `POST /users/login` returned `HTTP/1.1 200 OK` without `PasswordHash`.
- Wrong password returned `HTTP/1.1 401 Unauthorized` with `invalid email or password`.
- Wrong email returned `HTTP/1.1 401 Unauthorized` with `invalid email or password`.
- Missing password returned `HTTP/1.1 400 Bad Request`.
- Whitespace-only password returned `HTTP/1.1 400 Bad Request`.
- `GET /tasks?limit=1&offset=0` returned `HTTP/1.1 200 OK`.
- Review record: `reviews/2026-06-24-t011-user-login.md`.
