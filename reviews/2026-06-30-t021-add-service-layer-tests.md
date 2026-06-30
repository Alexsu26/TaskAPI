# Review: T021 Add Service Layer Tests

## Scope

Files reviewed:

- `internal/service/task_service.go`
- `internal/service/task_service_test.go`

## Findings

### High Priority

- None.

### Medium Priority

- None.

### Low Priority

- None after follow-up fixes.

## Strengths

- Introduced a service-level repository interface without changing route, handler, repository, or database behavior.
- Used a fake repository instead of a real PostgreSQL-backed repository.
- Kept the tests in `internal/service`, which matches the behavior being tested.
- Covered one success path and validation-error paths for `TaskService.Create`.
- Follow-up fixes now assert the fake repository received the expected task and use `errors.Is` for sentinel error checks.
- Tests do not start Gin, do not require an HTTP server, and do not require PostgreSQL.

## Learning Notes

- Service unit tests should verify business rules and boundary behavior without testing SQL.
- Repository integration tests should be handled separately with real database setup and cleanup.
- Fake repositories are useful because they let tests control dependency behavior and inspect what the service attempted to pass downstream.

## Verification

Commands:

```text
gofmt -l cmd/server internal
go test ./...
go vet ./...
```

Results:

```text
gofmt: no files listed
go test ./...: PASS for all packages
go vet ./...: PASS
```

Follow-up verification after fixing low-priority findings:

```text
gofmt -l cmd/server internal
go test ./...
go vet ./...
```

Results:

```text
gofmt: no files listed
go test ./...: PASS for all packages
go vet ./...: PASS
```

## Result

T021 accepted.

## Follow-Up Tasks

- Task: T022 Add Repository Integration Tests.
- Acceptance criteria: At least one repository package has integration tests with isolated setup data and cleanup, covering one success path and one not-found or database-error path.
