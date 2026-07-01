# Session: 2026-07-01 T024 Redis Rate Limit

## Context Read

- `AGENTS.md`
- `PROFILE.md`
- `ROADMAP.md`
- `SKILLS.md`
- `PROGRESS.md`
- `TASKS.md`
- `TASK_BACKLOG.md`
- `stages/02-go-engineering.md`
- `sessions/2026-06-30-t023-add-redis.md`

## Task

T024: Implement one focused Redis-backed use case.

Selected use case: login rate limiting.

## Work Completed

- Explained how Redis `INCR`, `Result`, `count == 1`, `EXPIRE`, and `count > limit` work for a fixed-window limiter.
- Reviewed the first implementation and identified that `Expire` errors were ignored.
- Recommended keeping the use case narrow and avoiding shared registration/login limiter semantics unless explicitly intended.
- Re-reviewed the learner's final implementation.
- Verified the Redis client is created in startup, passed through router wiring, and used by a Gin middleware mounted only on `POST /users/login`.
- Verified the limiter returns `429 Too Many Requests` on the sixth login attempt within one minute.
- Updated progress, tasks, backlog, skills, reflections, review, and session records.

## Evidence

Commands:

```text
git status --short --branch
gofmt -l cmd/server internal
go test ./...
go vet ./...
docker compose ps
docker compose exec -T redis redis-cli DEL rate_limit:login:127.0.0.1
docker compose exec -T postgres pg_isready -U taskapi -d taskapi
SERVER_PORT=18080 go run ./cmd/server
curl -s -o /tmp/taskapi-rate-N.json -w "%{http_code}" -X POST http://127.0.0.1:18080/users/login ...
docker compose exec -T redis redis-cli TTL rate_limit:login:127.0.0.1
docker compose exec -T redis redis-cli GET rate_limit:login:127.0.0.1
```

Results:

```text
gofmt: no files listed
go test ./...: PASS
go vet ./...: PASS
docker compose ps: postgres and redis running
postgres readiness: accepting connections
login status sequence: 401, 401, 401, 401, 401, 429
sixth response body: {"error":{"message":"please try again later"},"status":"error"}
Redis TTL: 60
Redis value: 6
```

## Learner Performance

Strengths:

- Correctly understood Redis `INCR` as incrementing a key and returning the incremented count.
- Kept the Redis use case focused on login rate limiting instead of combining cache, blacklist, and async work.
- Injected the existing Redis client through startup and router wiring instead of creating new clients in request paths.
- Used Gin route grouping to apply rate limiting only where intended.
- Fixed the ignored `EXPIRE` error after review feedback.

Weak Areas:

- The first implementation ignored the `Expire` result, which could leave a limiter key without TTL if Redis accepted `INCR` but failed expiration.
- README documentation is enough for this task but could become more copy-paste runnable in future documentation tasks.
- Current limiter is intentionally simple; proxy-aware IP handling and Redis atomicity patterns are still future learning topics.

## Next Task

T025: Add a small background worker or async task using goroutine/channel.
