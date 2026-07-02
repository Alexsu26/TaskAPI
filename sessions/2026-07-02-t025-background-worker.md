# Session: 2026-07-02 T025 Background Worker

## Context Read

- `AGENTS.md`
- `PROFILE.md`
- `ROADMAP.md`
- `SKILLS.md`
- `PROGRESS.md`
- `TASKS.md`
- `TASK_BACKLOG.md`
- `stages/02-go-engineering.md`
- `sessions/2026-07-01-t024-redis-rate-limit.md`

## Task

T025: Add a small background worker or async task using `goroutine` and `channel`.

Selected workflow: publish a task-created event after successful task creation and let a background worker log the event.

## Work Completed

- Clarified why the worker exists: this task is mainly for practicing async boundaries and Go concurrency, not because the current business feature requires a worker.
- Reviewed the learner's first implementation and found that the worker was created but never started.
- Re-reviewed the fixed implementation after `taskWorker.Start()` was added.
- Verified the worker is started during server startup and consumes events through a goroutine.
- Verified `RegisterTaskCreateRoutes` publishes a task-created event only after `taskService.Create` succeeds.
- Verified event publication is non-blocking through a buffered channel and `select` with a default drop path.
- Confirmed graceful worker shutdown is deferred to T026.
- Updated progress, tasks, backlog, skills, reflections, review, and session records.

## Evidence

Commands:

```text
git status --short --branch
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
docker compose ps: postgres and redis running
postgres readiness: accepting connections
POST /tasks status: 201
worker log: "msg":"task created event processed","task_id":4,"user_id":7,"title":"worker event test"
```

## Learner Performance

Strengths:

- Correctly understood the final flow: create the task synchronously, publish an event after success, return the HTTP response, and let the worker handle the event in the background.
- Kept the worker intentionally small and avoided external queues or unrelated async business features.
- Used a buffered channel plus non-blocking send to protect the HTTP path from a stuck or full worker queue.
- Fixed the review blocker by actually starting the worker goroutine.

Weak Areas:

- The first implementation missed the goroutine startup call, which meant the channel had no consumer.
- The first `Start` signature had an unused logger parameter, which made lifecycle ownership less clear.
- Worker shutdown and draining are not implemented yet and need focused practice in T026.

## Next Task

T026: Add graceful shutdown for the HTTP server and worker lifecycle using signal handling and timeout-bound context.
