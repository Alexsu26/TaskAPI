# Review: 2026-07-02 T026 Graceful Shutdown

## Scope

Reviewed the learner's implementation of graceful shutdown for:

- `cmd/server/main.go`
- `internal/worker/worker.go`

## Result

Approved after follow-up fixes.

## Review Notes

Initial implementation:

- Correctly replaced `r.Run(addr)` with `http.Server`.
- Correctly started `ListenAndServe` in a goroutine.
- Correctly waited for `os.Interrupt` and `SIGTERM` in the main goroutine.
- Correctly used `context.WithTimeout` and `httpServer.Shutdown(ctx)`.
- Added a worker `Stop` method, but it only closed the event channel and did not wait for the worker goroutine to finish draining queued events.

Follow-up review:

- Learner added a `done` channel so `Stop` could wait for worker completion.
- Review found that `Done` was exposed as a public field and initialized from `main`.
- Review also found that `defer taskWorker.Stop()` could close the worker channel if HTTP shutdown timed out while handlers might still publish events.

Final implementation:

- Added `worker.New(log, 100)` so worker internals are initialized inside the worker package.
- Kept `done` private.
- Used `defer close(w.done)` inside the worker goroutine to signal completion.
- Implemented `Stop` as close events, wait for `done`.
- Stopped the worker only after successful HTTP shutdown.
- Preserved existing routes, PostgreSQL startup, Redis startup, handler wiring, and task-created worker behavior.

## Verification

Commands run:

```text
gofmt -l cmd/server internal
go test ./...
go vet ./...
docker compose exec -T postgres pg_isready -U taskapi -d taskapi
docker compose exec -T redis redis-cli ping
SERVER_PORT=18080 go run ./cmd/server
curl -i --max-time 5 http://127.0.0.1:18080/health
```

Results:

```text
gofmt: no files listed
go test ./...: PASS
go vet ./...: PASS
postgres readiness: accepting connections
redis ping: PONG
GET /health: HTTP/1.1 200 OK, {"status":"ok"}
Ctrl+C logs: shutdown signal received, worker stopped, server shutdown
```

## Residual Notes

- If `ListenAndServe` returns an immediate non-`http.ErrServerClosed` error, the current code returns that error without explicitly stopping the worker. In the current `main` flow the process exits through `log.Fatalf`, so this is not blocking for T026, but a future lifecycle cleanup could handle that path explicitly.
