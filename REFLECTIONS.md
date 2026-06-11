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
