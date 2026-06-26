# Review: 2026-06-26 T014 Restrict Tasks To The Current User

## Scope

Reviewed learner implementation for T014 current-user task ownership.

Focus:

- project structure
- Gin usage
- handler/service/repository boundaries
- SQL ownership filtering
- T014 acceptance criteria

## First Review Findings

### P0: Update SQL used the wrong column name

File:

- `internal/repository/task_repository.go`

Finding:

- `UPDATE` filtered with `userID` instead of the PostgreSQL column `user_id`.
- Runtime result: both cross-user update and owner update returned 500.

Required fix:

- Use `where id = $1 and user_id = $5`.

### P1: Missing `return` after auth context failure in create/list handlers

File:

- `internal/handler/handler.go`

Finding:

- `RegisterTaskCreateRoutes` and `RegisterTasksListRoutes` returned a 401 response when `current_user_id` was missing, but continued executing service calls.
- This could cause double responses or incorrect follow-up errors.

Required fix:

- Add `return` after the 401 response.

## Final Review Result

Accepted.

The learner fixed both blockers:

- create/list handlers now return immediately after failed current-user lookup
- update SQL now filters by `user_id`

## What Went Well

- Preserved the existing package boundaries.
- Kept authentication extraction in handler code and task ownership enforcement in service/repository flow.
- Passed `userID` explicitly through handler, service, and repository methods.
- Removed the temporary hard-coded `UserID: 1`.
- Filtered list/detail/update/delete by authenticated `user_id`.
- Preserved 404 behavior for cross-user task access, avoiding task-existence leakage.
- Kept public routes public and protected task routes behind middleware.

## Verification

Static checks:

```text
gofmt -l cmd/server internal
go test ./...
go vet ./...
```

Results:

```text
gofmt: no output
go test ./...: PASS for all packages
go vet: no output
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

## Follow-Up

- T015 should add focused tests around service or handler behavior.
- Existing deferred cleanup remains: task response DTO/json tags and status whitelist can stay in later backlog items.
