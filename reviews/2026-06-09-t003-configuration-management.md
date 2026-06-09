# Review: 2026-06-09 T003 Add Configuration Management

## Scope

Reviewed learner implementation for T003 configuration management.

Files reviewed:

- `cmd/server/main.go`
- `internal/config/config.go`
- `internal/router/router.go`
- `internal/handler/handler.go`

## Acceptance Check

Passed:

- `go run ./cmd/server` starts the server on the default port.
- `GET /health` returns HTTP 200 on the default port.
- `SERVER_PORT=8090 go run ./cmd/server` starts the server on port 8090.
- `GET /health` returns HTTP 200 on the configured port.
- Invalid startup configuration produces an explicit startup error through `run()` and `log.Fatalf`.
- No database connection, auth flow, or task CRUD code was added.

Commands run:

```text
gofmt -l cmd/server internal
go test ./...
go run ./cmd/server
curl -i http://localhost:8080/health
SERVER_PORT=8090 go run ./cmd/server
curl -i http://localhost:8090/health
SERVER_PORT=abc go run ./cmd/server
```

Evidence:

```text
gofmt produced no output.
go test ./... passed.
Default /health returned HTTP/1.1 200 OK with {"status":"ok"}.
SERVER_PORT=8090 /health returned HTTP/1.1 200 OK with {"status":"ok"}.
SERVER_PORT=abc exited with: server startup failed: start server on :abc: listen tcp: lookup tcp/abc: unknown port.
```

## Findings

No blocking findings remain.

Follow-up review confirmed that `DATABASE_USER`, `DATABASE_PASSWORD`, and `DATABASE_NAME` are now loaded through `envOrDefault`.

Minor note for T004:

- `DATABASE_USER` defaults to `root`, while `DATABASE_NAME` defaults to an empty string. This is acceptable for T003 because no database connection exists yet, but T004 should align these defaults with the Docker Compose PostgreSQL service.

## Strengths

- `cmd/server/main.go` now has a clean `run() error` startup path.
- Startup errors from config loading and Gin startup are explicitly wrapped.
- Gin routing remains in `internal/router`, and health handler registration remains in `internal/handler`.
- The implementation stays within T003 scope and does not jump into database/auth/CRUD.

## Status

T003 completed and verified.
