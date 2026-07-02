# Review: 2026-07-02 T025 Background Worker

## Scope

Reviewed learner implementation of a minimal in-process async workflow for task-created events.

Files reviewed:

- `cmd/server/main.go`
- `internal/handler/handler.go`
- `internal/worker/worker.go`

## Result

Accepted.

## Findings

No blocking findings remain.

## Review Notes

- Project structure is appropriate for this learning task. The new worker boundary lives in `internal/worker`, startup owns worker construction, and the handler only publishes a task-created event after successful task creation.
- Gin handler behavior is preserved. `POST /tasks` still calls `taskService.Create`, handles errors through the existing error path, returns the existing task DTO with `201 Created`, and does not publish an event on failure.
- The worker now starts correctly. `cmd/server/main.go` creates `taskWorker`, calls `taskWorker.Start()`, and passes it into `handler.NewHandler`.
- The worker uses a goroutine to consume events from a channel and logs processed events with structured `task_id`, `user_id`, and `title` fields.
- Event publication is intentionally non-blocking. `PublishTaskCreated` uses `select` with a `default` case, so a full channel logs a warning instead of blocking the HTTP request path.
- Graceful shutdown is explicitly deferred to T026; the worker currently exits with the process.

## Verification

Commands run:

```text
gofmt -l cmd/server internal
go test ./...
go vet ./...
docker compose ps
docker compose exec -T postgres pg_isready -U taskapi -d taskapi
SERVER_PORT=18080 go run ./cmd/server
curl -s -X POST http://127.0.0.1:18080/users/register ...
curl -s -X POST http://127.0.0.1:18080/users/login ...
curl -s -o /tmp/taskapi-worker-create.json -w "%{http_code}" -X POST http://127.0.0.1:18080/tasks ...
```

Results:

```text
gofmt: no files listed
go test ./...: PASS
go vet ./...: PASS
postgres readiness: accepting connections
POST /tasks status: 201
worker log: "msg":"task created event processed","task_id":4,"user_id":7,"title":"worker event test"
```

## Follow-Up

- T026 should replace the current process-exit lifecycle with graceful shutdown for the HTTP server and explicit worker shutdown behavior.
