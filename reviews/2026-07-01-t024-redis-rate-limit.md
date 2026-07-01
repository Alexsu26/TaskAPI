# Review: 2026-07-01 T024 Redis Rate Limit

## Scope

Reviewed learner implementation of one focused Redis-backed use case: fixed-window rate limiting for `POST /users/login`.

## Result

Accepted.

## Findings

No blocking findings remain.

## Review Notes

- Project structure is appropriate for this task. The Redis-backed behavior lives in `internal/middleware/ratelimit.go`, while `cmd/server` owns Redis client creation and passes the client into `router.SetupRouter`.
- Gin usage is appropriate. The router creates a small group with `middleware.RateLimit(redisClient)` and registers only `POST /users/login` on that group, so registration and authenticated task routes are not accidentally rate limited.
- Handler/service/repository boundaries are preserved. The login handler and user service continue to own request parsing and authentication behavior; the middleware only decides whether the request may proceed.
- Redis key behavior is clear enough for the current learning task: `rate_limit:login:<client_ip>` uses `INCR`, sets a one-minute TTL with `EXPIRE` on first access, and returns `429 Too Many Requests` when the counter exceeds `5`.
- The implementation now checks both `INCR` and `EXPIRE` errors, avoiding silent limiter state where a key might remain without TTL.
- README documents that login is limited to five attempts per minute and that the sixth attempt returns `429`.

## Verification

Commands run:

```text
gofmt -l cmd/server internal
go test ./...
go vet ./...
docker compose ps
docker compose exec -T redis redis-cli DEL rate_limit:login:127.0.0.1
docker compose exec -T postgres pg_isready -U taskapi -d taskapi
SERVER_PORT=18080 go run ./cmd/server
curl POST /users/login six times with invalid credentials
docker compose exec -T redis redis-cli TTL rate_limit:login:127.0.0.1
docker compose exec -T redis redis-cli GET rate_limit:login:127.0.0.1
```

Results:

```text
gofmt: no files listed
go test ./...: PASS
go vet ./...: PASS
postgres readiness: accepting connections
login status sequence: 401, 401, 401, 401, 401, 429
sixth response body: {"error":{"message":"please try again later"},"status":"error"}
Redis TTL: 60
Redis value: 6
```

## Follow-Up

- Keep rate limiting marked as practiced but not stable yet. Future production work should consider proxy-aware client IP handling, configurability, and atomic `INCR` plus TTL behavior through Lua or another Redis-safe pattern.
