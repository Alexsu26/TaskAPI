# Session: 2026-06-23 T008 Task Detail, Update, And Delete

## Context Read

- `AGENTS.md`
- `PROFILE.md`
- `ROADMAP.md`
- `PROGRESS.md`
- `TASKS.md`
- `TASK_BACKLOG.md`
- `SKILLS.md`
- `stages/01-go-task-api.md`

## Task

T008: Implement task detail (`GET /tasks/:id`), update (`PUT /tasks/:id`), and delete (`DELETE /tasks/:id`) with handler/service/repository boundaries.

## Work Completed

- Discussed architecture question: where to handle `limit`/`offset` query parameter parsing. Concluded that string parsing belongs in handler (HTTP concern), but default values and range validation belong in service (business policy). Learner refactored pagination to use `*int` pointers distinguishing "not provided" from "explicitly zero".
- Discussed layering concern: whether three-layer separation adds value when errors must flow through all layers. Concluded that each layer has a distinct responsibility — repo handles SQL errors, service translates to domain errors, handler maps to HTTP status codes.
- Implemented `GET /tasks/:id` (detail) through all three layers.
- Implemented `PUT /tasks/:id` (update) through all three layers with `UPDATE ... RETURNING`.
- Implemented `DELETE /tasks/:id` (delete) through all three layers with `RowsAffected` check.
- Debugged nil pointer panic in `GetById` caused by named return value `task *model.Task` being nil during `Scan`.
- Debugged 404 "page not found" caused by routes not being registered in `SetupRouter`.
- Debugged compilation error where Update handler referenced undefined `task` variable.
- Fixed multiple bugs found across two review rounds: Scan column count mismatch, JSON tag typo, hardcoded `"id"` string instead of `ctx.Param("id")`, un-trimmed title stored in update, missing route registrations.
- Final build, vet, and test all pass clean.

## Evidence

Commands:

```text
go build ./...
go vet ./...
go test ./...
```

Results:

```text
?   taskapi/cmd/server    [no test files]
?   taskapi/internal/config    [no test files]
?   taskapi/internal/database    [no test files]
?   taskapi/internal/handler    [no test files]
?   taskapi/internal/model    [no test files]
?   taskapi/internal/repository    [no test files]
?   taskapi/internal/router    [no test files]
?   taskapi/internal/service    [no test files]
ALL CHECKS PASSED
```

## Learner Performance

Strengths:

- Asked strong architectural questions about layer responsibilities, showing growing understanding of separation of concerns.
- Pagination refactored correctly after discussion — pointer-based approach to distinguish "not provided" vs "explicitly zero" is idiomatic.
- Error mapping chain across three layers is clean and consistent once bugs were fixed.
- Responded well to iterative review feedback, fixing all critical and warning-level issues.

Weaknesses:

- Route registration in `SetupRouter` was missed twice (once for GetByID, once for Update/Delete). Should build a habit of checking router wiring before testing.
- Copy-paste errors (JSON tag, hardcoded string) suggest need for more careful review of own code before requesting review.
- Nil pointer panic from named return value shows Go pointer semantics still need reinforcement.

## Next Task

T009: Add Unified Response And Error Handling — standardize API response format and application error handling.
