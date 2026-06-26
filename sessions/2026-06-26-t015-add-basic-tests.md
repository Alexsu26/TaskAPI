# Session: 2026-06-26 T015 Add Basic Tests

## Context Read

- `AGENTS.md`
- `PROFILE.md`
- `ROADMAP.md`
- `PROGRESS.md`
- `TASKS.md`
- `TASK_BACKLOG.md`
- `SKILLS.md`
- `stages/01-go-task-api.md`
- `sessions/2026-06-26-t014-restrict-tasks-to-current-user.md`

## Task

T015: Add basic tests.

## Work Completed

- Explained the basic structure of Go `_test.go` files and `TestXxx(t *testing.T)` functions.
- Explained why tests use `t.Fatalf` instead of `fmt.Printf`.
- Explained why `%v` is used for test output and `%w` is used only for `fmt.Errorf` error wrapping.
- Reviewed learner-added `internal/auth/token_test.go`.
- Verified the tests cover JWT token generation/parsing success and invalid-token/invalid-user error paths.
- Ran task verification commands.
- Accepted T015.
- Updated progress, task, backlog, skill, reflection, review, and session records.

## Evidence

Commands:

```text
gofmt -l cmd/server internal
go test ./...
go vet ./...
```

Results:

```text
gofmt: no output
go test ./...: PASS for all packages
go vet: no output
```

## Learner Performance

Strengths:

- Added a focused test file under the correct package.
- Used the standard `testing` package and simple deterministic test setup.
- Covered one meaningful success path and two meaningful error paths.
- Used `errors.Is` for sentinel error checks.

Weaknesses:

- Testing concepts are still new, especially deciding when to use `Fatalf`, `Errorf`, `%v`, and `%w`.
- Broader service/handler tests and test doubles still need later practice.

## Next Task

T016: Complete Stage 1 documentation.
