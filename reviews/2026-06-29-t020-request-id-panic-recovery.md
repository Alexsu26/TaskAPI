# Review: 2026-06-29 T020 Add Request ID And Panic Recovery

## Scope

Reviewed learner implementation for request ID middleware, request ID propagation into logs, and custom panic recovery.

## Result

Accepted after one fix.

## Findings

### Fixed

- Incoming `X-Request-ID` was initially accepted as raw client input. This did not satisfy the "safe incoming header" acceptance criterion because unsafe or malformed values could enter structured logs. The learner fixed this by accepting UUID-shaped values only and generating a new UUID otherwise.

## What Was Verified

- Request ID middleware stores `request_id` in Gin context.
- Response includes `X-Request-ID`.
- Request logs include `request_id`.
- Internal 500 error logs include `request_id`.
- Custom panic recovery replaces `gin.Recovery()` and logs request ID, method, path, and panic value.
- Panic response uses the existing generic unified 500 shape.
- Logging fields avoid request body, Authorization header, tokens, passwords, password hashes, query strings, JWT secret, and database password.
- Normal `/health` response shape remains unchanged.

## Commands

```text
gofmt -l cmd/server internal
go test ./...
go vet ./...
SERVER_PORT=18080 go run ./cmd/server
curl -s -i http://localhost:18080/health
curl -s -i -H 'X-Request-ID: 11111111-1111-4111-8111-111111111111' http://localhost:18080/health
curl -s -i -H 'X-Request-ID: bad id with spaces' http://localhost:18080/health
```

## Evidence

```text
gofmt: no files listed
go test ./...: PASS for all packages
go vet ./...: PASS
/health without X-Request-ID: 200 OK with generated X-Request-ID UUID
/health with valid UUID X-Request-ID: 200 OK with same X-Request-ID
/health with unsafe X-Request-ID: 200 OK with replacement UUID
/health body: {"status":"ok"}
request logs: JSON logs include request_id for each request
```

## Notes For Next Task

T021 should focus on narrow service-layer tests. Prefer deterministic unit tests with small fake repository implementations before moving to integration tests.
