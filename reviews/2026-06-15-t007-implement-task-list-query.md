# Review: T007 Implement Task List Query

Date: 2026-06-15

## Scope

Reviewed files:

- `internal/handler/handler.go`
- `internal/router/router.go`
- `internal/service/task_service.go`
- `internal/repository/task_repository.go`

Task scope:

- Stage 1, substage 1.3 Task CRUD
- T007 only: `GET /tasks` task list query with basic pagination
- Out of scope: task detail/update/delete and authentication

## Verification

Commands run:

```text
gofmt -l cmd/server internal
go test ./...
docker compose ps
SERVER_PORT=18080 go run ./cmd/server
curl -s -i http://localhost:18080/health
curl -s -i 'http://localhost:18080/tasks'
curl -s -i 'http://localhost:18080/tasks?limit=10&offset=0'
curl -s -i 'http://localhost:18080/tasks?limit=-1&offset=0'
curl -s -i 'http://localhost:18080/tasks?limit=abc&offset=0'
curl -s 'http://localhost:18080/tasks?limit=1&offset=1'
```

Results:

- `gofmt -l cmd/server internal`: passed, no files printed.
- `go test ./...`: passed.
- PostgreSQL container was running through Docker Compose.
- Temporary server started from current source on `:18080`.
- `GET /health`: returned `HTTP/1.1 200 OK` and `{"status":"ok"}`.
- `GET /tasks`: returned `HTTP/1.1 200 OK` and task data from PostgreSQL.
- `GET /tasks?limit=10&offset=0`: returned `HTTP/1.1 200 OK`.
- `GET /tasks?limit=1&offset=1`: returned one later page result.
- `GET /tasks?limit=-1&offset=0`: returned `HTTP/1.1 400 Bad Request`.
- `GET /tasks?limit=abc&offset=0`: returned `HTTP/1.1 400 Bad Request`.

## Findings

### Low - pagination error name and message are too vague

Files:

- `internal/service/task_service.go`
- `internal/handler/handler.go`

`ParaInvalid` and the response `error parameters` work functionally, but they do not clearly describe the problem. A later cleanup should rename this to something like `ErrInvalidPagination` and return a message such as `invalid pagination parameters`.

### Low - no upper bound for limit yet

File:

- `internal/service/task_service.go`

The current validation rejects `limit <= 0`, which satisfies the task. In a later hardening step, add a maximum limit such as `100` to avoid accidentally returning too many rows.

## Strengths

- `GET /tasks` is registered through the existing router and handler pattern.
- Query parsing uses defaults and allows `limit` and `offset` to be provided independently.
- Handler returns immediately after invalid input and internal errors.
- Service owns pagination validation instead of pushing all validation into the repository.
- Repository uses `Query` for multi-row reads, closes rows, scans explicit columns, and checks `rows.Err`.
- SQL uses stable ordering with `updated_at DESC, id DESC`.
- Implementation stays inside T007 scope and does not add detail/update/delete or authentication.

## Acceptance Status

T007 is completed and verified.

Acceptance criteria:

- `GET /tasks` route exists: yes.
- query supports basic pagination: yes, `limit` and `offset`.
- invalid pagination input returns a clear client error: yes for status code and behavior; wording should be improved later.
- valid request reads tasks from PostgreSQL: yes, verified through HTTP.
- results are returned in a stable order: yes, `updated_at DESC, id DESC`.
- code follows handler/service/repository package boundaries: yes.
- `go test ./...` passes: yes.
- existing `/health` endpoint still works: yes.
- detail/update/delete and authentication remain out of scope: yes.

## Follow-Up

- T008: implement task detail, update, and delete.
- Later cleanup: improve pagination error naming/message and consider a maximum `limit`.
