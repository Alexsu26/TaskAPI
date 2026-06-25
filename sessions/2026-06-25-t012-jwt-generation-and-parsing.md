# Session: 2026-06-25 T012 JWT Generation And Parsing

## Context Read

- `AGENTS.md`
- `PROFILE.md`
- `ROADMAP.md`
- `PROGRESS.md`
- `TASKS.md`
- `TASK_BACKLOG.md`
- `SKILLS.md`
- `stages/01-go-task-api.md`
- `sessions/2026-06-24-t011-user-login.md`

## Task

T012: Implement JWT generation and parsing.

## Work Completed

- Explained JWT purpose, claims, signing, `IssuedAt`, and why JWT is needed for later authenticated APIs.
- Explained responsibility boundaries between token helpers, middleware, services, and repository/database validation.
- Reviewed the learner implementation of JWT config and `internal/auth.TokenManager`.
- First review found nonstandard env var names, missing positive user ID validation in parsing, missing expired-token distinction, and missing positive expiration validation.
- Learner fixed the config and token parsing issues.
- Second review found token generation should also reject non-positive user IDs.
- Learner fixed token generation validation.
- Reviewed login token integration.
- Third review found response envelope nesting problems in login and registration responses.
- Learner fixed the response shape so login returns `data.user` and `data.token`, and registration returns `data.user`.
- Final verification accepted T012 and updated task, progress, skill, reflection, review, and session records.

## Evidence

Commands:

```text
gofmt -l cmd/server internal
go test ./...
go vet ./...
docker compose up -d postgres
SERVER_PORT=18080 JWT_SECRET=acceptance-secret JWT_EXPIRATION_MINUTES=60 go run ./cmd/server
```

Runtime checks:

```text
GET /health -> 200
POST /users/register valid request -> 201, user DTO without PasswordHash
POST /users/login valid credentials -> 200, user DTO plus JWT token
POST /users/login wrong password -> 401
GET /tasks?limit=1&offset=0 -> 200
JWT token -> three segments, payload includes user_id, exp, and iat
```

Results:

```text
gofmt: no output
go test ./...: PASS for all packages
go vet ./...: no output
```

## Learner Performance

Strengths:

- Asked precise conceptual questions before implementing JWT.
- Preserved the existing handler/service/repository structure and added `internal/auth` as a focused boundary.
- Passed token configuration through startup rather than reading env vars from token helpers.
- Corrected token validation and response envelope issues after review.

Weaknesses:

- Initially mixed response wrapper concerns by adding an extra `data` nesting layer.
- Initially did not validate all token boundary cases such as invalid user IDs and invalid expiration values.
- Needs more practice separating cryptographic token validity from business validity such as user existence and authorization.

## Next Task

T013: Add JWT auth middleware.
