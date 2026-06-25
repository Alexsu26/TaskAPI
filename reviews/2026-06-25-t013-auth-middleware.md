# Review: 2026-06-25 T013 Auth Middleware

## Scope

Reviewed learner implementation for `T013: Add Auth Middleware`.

Focus:

- project structure
- Gin middleware usage
- T013 acceptance criteria

## Result

Accepted after fix.

The first review found a package import cycle. The learner fixed the dependency direction, and the task now compiles and passes runtime acceptance checks.

## Findings Fixed

### P0: `handler` and `router` import each other

Evidence:

```text
go test ./...
```

Result:

```text
imports taskapi/internal/handler from main.go
imports taskapi/internal/router from handler.go
imports taskapi/internal/handler from router.go: import cycle not allowed
```

Cause:

- `internal/router/router.go` imports `taskapi/internal/handler`.
- `internal/handler/handler.go` imports `taskapi/internal/router` only to use `router.RouteRegister`.

Why this matters:

- Go packages must form a one-way dependency graph.
- The current code cannot build, so T013 acceptance criteria are not met yet.

Fix direction:

- Moved the route-registration helper type out of the router dependency path so `handler` no longer imports `router`.
- The package dependency graph is now acyclic.

## Positive Notes

- `internal/middleware` is a reasonable package boundary for auth middleware.
- The middleware reads `Authorization`, checks the `Bearer` prefix, trims the token, calls `TokenManager.ParseToken`, aborts failed requests with `401`, and sets `current_user_id` on success.
- Public routes and protected task routes are separated in router setup.
- No SQL task-ownership filtering was added, which correctly keeps T014 out of scope.

## Final Verification

```text
gofmt -l cmd/server internal
```

Result: no output.

```text
go test ./...
```

Result: passed for all packages.

```text
go vet ./...
```

Result: no output.

Runtime checks:

```text
GET /health -> 200
POST /users/register -> 201
POST /users/login -> 200 with JWT token
GET /tasks?limit=1&offset=0 without Authorization -> 401
GET /tasks?limit=1&offset=0 with malformed Authorization -> 401
GET /tasks?limit=1&offset=0 with invalid token -> 401
GET /tasks?limit=1&offset=0 with expired signed token -> 401
GET /tasks?limit=1&offset=0 with valid token -> 200
```

## Next Step

Start T014: read `current_user_id` from Gin context and restrict task create/list/detail/update/delete operations to the authenticated user.
