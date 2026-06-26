# Session: 2026-06-26 T014 Restrict Tasks To The Current User

## Context Read

- `AGENTS.md`
- `PROFILE.md`
- `ROADMAP.md`
- `PROGRESS.md`
- `TASKS.md`
- `TASK_BACKLOG.md`
- `SKILLS.md`
- `stages/01-go-task-api.md`
- `sessions/2026-06-25-t013-auth-middleware.md`

## Task

T014: Restrict tasks to the current user.

## Work Completed

- Reviewed learner implementation of current-user task ownership.
- First review found two blockers:
  - missing `return` after failed current-user lookup in create/list handlers
  - update SQL used `userID` instead of the `user_id` database column
- Learner fixed both blockers.
- Re-ran static verification.
- Ran runtime acceptance checks with two registered users and separate JWTs.
- Accepted T014.
- Updated progress, task, backlog, skill, reflection, review, and session records.

## Evidence

Commands:

```text
gofmt -l cmd/server internal
go test ./...
go vet ./...
docker compose up -d postgres
SERVER_PORT=18080 JWT_SECRET=review-secret JWT_EXPIRATION_MINUTES=60 go run ./cmd/server
```

Runtime checks:

```text
GET /health -> 200
POST /users/register A -> 201
POST /users/register B -> 201
POST /tasks with user A token -> 201, task.UserID = user A
GET /tasks with user A token -> 200
GET /tasks with user B token -> 200, tasks=[]
GET /tasks/:id with user A token -> 200
GET /tasks/:id with user B token -> 404
PUT /tasks/:id with user B token -> 404
PUT /tasks/:id with user A token -> 200
DELETE /tasks/:id with user B token -> 404
DELETE /tasks/:id with user A token -> 200
GET /tasks without token -> 401
```

Results:

```text
gofmt: no output
go test ./...: PASS for all packages
go vet ./...: no output
```

## Learner Performance

Strengths:

- Correctly used the `current_user_id` set by auth middleware.
- Passed the authenticated user ID through handler, service, and repository boundaries.
- Removed the pre-auth hard-coded task owner.
- Added SQL ownership filters for list/detail/update/delete.
- Preserved 404 responses for cross-user access to avoid revealing task existence.
- Fixed review findings quickly and locally without broad rewrites.

Weaknesses:

- Initially missed `return` after two Gin error responses, which could have caused double-response behavior.
- SQL field naming typo was only caught by runtime verification, not static checks.
- Needs more practice using runtime acceptance tests for database-backed authorization paths.

## Next Task

T015: Add basic tests.
