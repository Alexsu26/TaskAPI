# Review: T004 Add PostgreSQL With Docker Compose

## Scope

Files or modules reviewed:

- `docker-compose.yml`
- `internal/config/config.go`
- `internal/database/database.go`
- `cmd/server/main.go`
- `go.mod`
- `go.sum`

## Findings

### High Priority

- None.

### Medium Priority

- None.

### Low Priority

- Finding: `cmd/server/main.go` initializes `*sql.DB` but does not retain or close it.
- Reason: For T004, startup `Ping` evidence is enough. For future repository work, the application will need explicit ownership of the database handle so repositories can use it and shutdown can close it.
- Suggested fix: In a later task, keep the database handle in the startup path and pass it into repository constructors. Add `defer db.Close()` once the application owns the handle for the process lifetime.

## Strengths

- `docker-compose.yml` is minimal and scoped to PostgreSQL.
- Compose credentials match `internal/config` defaults.
- `internal/database` is a clear package boundary for connection setup.
- Startup errors are wrapped with useful context.
- Existing Gin router and `/health` handler structure remain unchanged.
- No auth, task CRUD, or migration code was added early.

## Learning Notes

- `sql.Open` creates a database handle; `Ping` is what proves the connection can currently be established.
- `*sql.DB` is a concurrency-safe connection pool and is normally long-lived.
- A liveness endpoint like `/health` can return 200 even when readiness checks are handled separately later.

## Follow-Up Tasks

- Task: In T005, design user and task model structs that map cleanly to future SQL tables.
- Acceptance criteria: model fields are understandable, scoped to upcoming auth/task CRUD needs, and `go test ./...` still passes without adding handlers or repositories.
