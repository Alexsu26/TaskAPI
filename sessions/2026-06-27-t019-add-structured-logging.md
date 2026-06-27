# Session: 2026-06-27 T019 Add Structured Logging

## Context Read

- `AGENTS.md`
- `PROFILE.md`
- `ROADMAP.md`
- `SKILLS.md`
- `PROGRESS.md`
- `TASKS.md`
- `TASK_BACKLOG.md`
- `stages/02-go-engineering.md`
- `sessions/2026-06-27-t018-dto-model-response-boundaries.md`

## Task

T019: Add structured logging.

## Work Completed

- Explained that logging is server-side observability output rather than API input/output.
- Broke the task into small implementation steps: logger package, startup log, router injection, Gin request logger middleware, and 500 error logging.
- Reviewed the learner's implementation.
- First review found that the 500 error log used the full request URL and could include query string values.
- Re-reviewed after the learner changed the field to `ctx.Request.URL.Path`.
- Verified format, tests, vet, and runtime logging behavior.
- Accepted T019.
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
curl -s -i http://localhost:18080/tasks
curl -s -i -X POST http://localhost:18080/users/login -H 'Content-Type: application/json' -d '{"email":"nobody@example.com","password":"bad"}'
rg -n "gin\\.Default|gin\\.New|RequestLogger|Recovery|Request\\.URL|RawQuery|Authorization|Request\\.Body|JWTSecret|Database\\.Password|PasswordHash|password_hash|log\\.(Info|Error)" cmd internal
```

Results:

```text
gofmt: no files listed
go test ./...: PASS for all packages
go vet ./...: PASS
runtime /health: 200 OK
runtime unauthenticated /tasks: 401 not login
runtime failed login: 401 invalid email or password
runtime logs: JSON startup/request logs emitted with method, path, status, duration_ms, and client_ip
sensitive log scan: no Authorization, request body, JWT secret, DB password, password, password hash, or query string logging in the new log statements
```

## Learner Performance

Strengths:

- Chose a small standard-library logging approach.
- Kept logging in a focused package and middleware boundary.
- Correctly used Gin middleware to capture request status and duration.
- Preserved route behavior and response bodies.
- Fixed the security review finding narrowly by switching from full URL to path-only logging.

Weaknesses:

- Needed extra clarification to distinguish logs from API request/response data.
- Initial error logging did not account for query-string sensitivity.

## Next Task

T020: Add request ID and panic recovery.
