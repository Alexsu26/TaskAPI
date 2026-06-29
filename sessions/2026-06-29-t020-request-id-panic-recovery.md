# Session: 2026-06-29 T020 Add Request ID And Panic Recovery

## Context Read

- `AGENTS.md`
- `PROFILE.md`
- `ROADMAP.md`
- `SKILLS.md`
- `PROGRESS.md`
- `TASKS.md`
- `TASK_BACKLOG.md`
- `stages/02-go-engineering.md`
- `sessions/2026-06-27-t019-add-structured-logging.md`

## Task

T020: Add request ID and panic recovery.

## Work Completed

- Explained the request ID middleware boundary and middleware ordering.
- Explained why panic recovery uses `defer`, `recover()`, and `ctx.Next()`.
- Reviewed the learner's implementation against project structure, Gin usage, and acceptance criteria.
- First review found unsafe raw `X-Request-ID` acceptance.
- Re-reviewed after the learner validated request IDs as UUIDs and generated replacements for unsafe values.
- Verified static checks and runtime request ID behavior.
- Accepted T020.
- Updated progress, task, backlog, skill, reflection, review, and session records.

## Evidence

Commands:

```text
git status --short --branch
gofmt -l cmd/server internal
go test ./...
go vet ./...
SERVER_PORT=18080 go run ./cmd/server
curl -s -i http://localhost:18080/health
curl -s -i -H 'X-Request-ID: 11111111-1111-4111-8111-111111111111' http://localhost:18080/health
curl -s -i -H 'X-Request-ID: bad id with spaces' http://localhost:18080/health
```

Results:

```text
gofmt: no files listed
go test ./...: PASS for all packages
go vet ./...: PASS
runtime /health without request ID: generated X-Request-ID UUID
runtime /health with valid UUID request ID: preserved X-Request-ID
runtime /health with unsafe request ID: generated replacement X-Request-ID UUID
runtime /health body: {"status":"ok"}
runtime request logs: emitted request_id
```

## Learner Performance

Strengths:

- Kept request tracing in middleware and helper boundaries.
- Correctly used Gin context to share request ID across middleware and handlers.
- Replaced default Gin recovery with custom recovery to include request-level log context.
- Preserved existing response envelope and avoided logging request bodies or secrets.
- Fixed the unsafe header review finding with a narrow validation change.

Weaknesses:

- Needed additional explanation of the `defer` and `recover()` control flow.
- Initially treated incoming request ID as trusted input.

## Next Task

T021: Add service layer tests.
