# Review: T019 Add Structured Logging

## Scope

Files reviewed:

- `cmd/server/main.go`
- `internal/logger/logger.go`
- `internal/router/router.go`
- `internal/middleware/middleware.go`
- `internal/handler/handler.go`

## Findings

### High Priority

- None.

### Medium Priority

- Initial finding: the 500 error log used `ctx.Request.URL`, which could include query string values.
- Reason: logs should avoid recording potentially sensitive query parameters, especially as the API grows.
- Resolution: changed the field to `ctx.Request.URL.Path`.

### Low Priority

- None blocking. Startup logs now include `database_name`, which is useful and non-sensitive.

## Strengths

- Added a small `internal/logger` package using Go's standard `log/slog` JSON handler.
- Kept logging as infrastructure and middleware instead of spreading request logging through each handler.
- Used `gin.New()` with custom request logging and `gin.Recovery()` correctly.
- Request logs include method, path, status, duration, and client IP.
- Startup logs include address and key non-sensitive configuration.
- Internal 500 errors are logged for operators while client responses remain generic.
- Sensitive values such as JWT tokens, Authorization headers, passwords, password hashes, JWT secrets, database passwords, and full request bodies are not logged.

## Learning Notes

- Logging is server-side observability output, not a user-facing API.
- Request logging belongs in middleware because it applies across routes.
- `ctx.Next()` lets middleware run code before and after the handler, which is necessary for duration and final status logging.
- Structured logs should be designed with security boundaries in mind.

## Follow-Up Tasks

- Task: T020 Add Request ID And Panic Recovery.
- Acceptance criteria: every request has a traceable request ID, logs include the request ID, panic recovery returns a safe response, and existing API behavior remains stable.
