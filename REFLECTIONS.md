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
