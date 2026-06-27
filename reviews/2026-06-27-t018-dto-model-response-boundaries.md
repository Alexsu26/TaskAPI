# Review: T018 Refactor DTO, Model, And Response Boundaries

Date: 2026-06-27

## Scope

Reviewed the learner's T018 implementation for project structure, Gin handler usage, DTO boundaries, response field names, sensitive-field exposure, and acceptance criteria.

## Result

Approved after two fixes.

## Findings

### Fixed: task list response added zero-value items

The first `toTasksResponse` implementation created a slice with `len(tasks)` and then appended converted tasks. That would have returned zero-value task objects before the real tasks.

The learner fixed this by creating the slice with length `0` and capacity `len(tasks)` before appending.

### Fixed: `UpdatedAt` mapper field mismatch

After renaming the response DTO field to `UpdatedAt`, the mapper still assigned `UpdateAt`, which caused compilation to fail.

The learner fixed the mapper to assign `UpdatedAt: task.UpdatedAt`.

## Accepted Changes

- Moved handler DTO definitions into `internal/handler/dto.go`.
- Added `UserResponse` and `TaskResponse` with JSON tags.
- Converted task create, list, detail, and update responses to `TaskResponse`.
- Converted registration and login responses to `UserResponse`.
- Preserved the existing unified response envelope.
- Preserved handler, service, repository, and model responsibilities.
- Avoided exposing `PasswordHash` and task ownership internals in API responses.

## Verification

Commands run:

```text
gofmt -l cmd/server internal
go test ./...
go vet ./...
docker exec taskapi-database dropdb --if-exists -U taskapi taskapi_t018_review
docker exec taskapi-database createdb -U taskapi taskapi_t018_review
migrate -path migrations -database "postgres://taskapi:taskapi@localhost:5432/taskapi_t018_review?sslmode=disable" up
SERVER_PORT=18080 DATABASE_NAME=taskapi_t018_review go run ./cmd/server
curl -s -i "http://localhost:18080/users/register" ...
curl -s "http://localhost:18080/users/login" ...
curl -s -i "http://localhost:18080/tasks" ...
curl -s "http://localhost:18080/tasks?limit=20&offset=0" ...
docker exec taskapi-database dropdb --if-exists -U taskapi taskapi_t018_review
```

Results:

- `gofmt -l cmd/server internal` produced no output.
- `go test ./...` passed for all packages.
- `go vet ./...` passed.
- Runtime registration returned `user.id`, `user.name`, and `user.email` without `PasswordHash`.
- Runtime login returned `user.id`, `user.name`, `user.email`, and `token`.
- Runtime task creation returned `id`, `title`, `description`, `status`, `created_at`, and `updated_at`.
- Runtime task listing returned exactly the created task with the same response DTO shape and no zero-value extra items.

## Notes For Next Task

README API examples were synchronized before publishing so they use the new DTO field names.
