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
