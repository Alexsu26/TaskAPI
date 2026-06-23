# Review: 2026-06-23 T009 Unified Response And Error Handling

## Result

Accepted.

## Scope Reviewed

- `internal/handler/handler.go`
- `internal/handler/response.go`

## Findings

No blocking issues remain.

## Review Notes

- Success responses now use a common response helper and consistent envelope.
- Error responses now use a common error envelope for both service-layer errors and HTTP-layer parsing/binding errors.
- Service-layer sentinel errors are mapped centrally:
  - invalid parameters -> 400
  - missing title -> 400
  - task not found -> 404
  - unknown/internal errors -> 500 with a generic message
- HTTP-layer errors remain in the handler layer and return 400 without being confused with service errors.
- `POST /tasks` correctly returns `201 Created`.
- `DELETE /tasks/:id` now returns `200 OK` with the unified success response; this is an acceptable consistency tradeoff for T009.

## Verification

Commands:

```text
gofmt -l cmd/server internal
go test ./...
go vet ./...
docker compose up -d postgres
SERVER_PORT=18080 go run ./cmd/server
```

Runtime checks:

```text
GET /health -> 200
POST /tasks with invalid body -> 400
GET /tasks?limit=abc -> 400
GET /tasks/nope -> 400
GET /tasks/999999999 -> 404
POST /tasks with valid title -> 201
GET /tasks?limit=1&offset=0 -> 200
PUT /tasks/:id with invalid body -> 400
PUT /tasks/:id with valid body -> 200
DELETE /tasks/:id -> 200
```

## Follow-Up

- T010 should reuse the unified response/error pattern for user registration.
- Consider renaming `handlerCommonError` later if a clearer local naming convention emerges.
