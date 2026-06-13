# Review: T006 Implement Task Creation

Date: 2026-06-13

## Scope

Reviewed files:

- `cmd/server/main.go`
- `internal/database/migration.go`
- `internal/handler/handler.go`
- `internal/router/router.go`
- `internal/service/task_service.go`
- `internal/repository/task_repository.go`

Task scope:

- Stage 1, substage 1.3 Task CRUD
- T006 only: `POST /tasks` task creation
- Out of scope: task list/detail/update/delete and authentication

## Verification

Commands run:

```text
gofmt -l cmd/server internal
go test ./...
docker compose up -d
go run ./cmd/server
curl -i http://localhost:8080/health
curl -i -X POST http://localhost:8080/tasks -H 'Content-Type: application/json' -d '{"description":"missing title"}'
curl -i -X POST http://localhost:8080/tasks -H 'Content-Type: application/json' -d '{"title":"   ","description":"blank title"}'
curl -i -X POST http://localhost:8080/tasks -H 'Content-Type: application/json' -d '{"title":"learn gin","description":"practice task creation"}'
```

Results:

- `gofmt -l cmd/server internal`: passed, no files printed.
- `go test ./...`: passed.
- `docker compose up -d`: started PostgreSQL container.
- `GET /health`: returned `HTTP/1.1 200 OK` and `{"status":"ok"}`.
- `POST /tasks` without `title`: returned `HTTP/1.1 400 Bad Request` and `{"error":"invalid request body"}`.
- `POST /tasks` with whitespace-only `title`: returned `HTTP/1.1 400 Bad Request` and `{"error":"title is required"}`.
- `POST /tasks` with valid title and description: returned `HTTP/1.1 201 Created` and a task containing `ID`, `UserID`, `Title`, `Description`, `Status`, `CreatedAt`, and `UpdatedAt`.

## Findings

### Low - temporary dev user strategy should be marked as temporary

Files:

- `internal/database/migration.go`
- `internal/service/task_service.go`

The current strategy is acceptable for T006: migration seeds user `id = 1`, and task creation uses `UserID: 1` until authentication exists.

Add a short TODO near `UserID: 1` or the seed SQL later, so it is clear this must be replaced by the authenticated user ID in T013/T014.

### Low - startup error typo

File: `cmd/server/main.go`

`creat table error` should be `create table error` or `run migrations`.

### Low - remove commented router struct

File: `internal/router/router.go`

The commented `Router` struct is not used. Remove it when cleaning up the task.

## Strengths

- `POST /tasks` is wired through `main -> router -> handler -> service -> repository`.
- `ctx.ShouldBindJSON` is used correctly for request parsing.
- Validation errors and database/internal errors are separated.
- Error responses return immediately and do not fall through to success responses.
- Successful create returns `201 Created`, which matches REST creation semantics.
- Repository uses `QueryRow` plus `RETURNING id, created_at, updated_at`, so the response includes database-generated fields.
- Title validation trims whitespace before checking required input.
- The temporary dev user strategy handles the lack of authentication without implementing auth early.
- T006 remains scoped to create-only behavior; list/detail/update/delete/auth were not added.

## Acceptance Status

T006 is completed and verified.

Acceptance criteria:

- `POST /tasks` route exists: yes.
- request body includes title and optional description: yes.
- invalid input returns a clear client error: yes.
- valid input inserts into PostgreSQL: yes, verified through HTTP.
- successful creation returns the created task/key fields: yes, includes database-generated fields.
- code follows handler/service/repository boundaries: yes.
- `go test ./...` passes: yes.
- existing `/health` endpoint still works: yes.
- list/detail/update/delete and authentication remain out of scope: yes.

## Follow-Up

- T007: implement task list query with pagination.
