# Review: T008 Task Detail, Update, And Delete

## Scope

Files reviewed:

- `internal/repository/task_repository.go`
- `internal/service/task_service.go`
- `internal/handler/handler.go`
- `internal/router/router.go`
- `internal/model/task.go`

## Findings

### High Priority (all fixed during review)

- Finding: Nil pointer dereference panic in `GetById` — named return value `task *model.Task` started as `nil`, then `&task.ID` was used in `Scan`, causing a panic on nil dereference.
- Reason: `Scan(&task.ID)` must access `task.ID` to take its address, but `task` had no allocated `model.Task` instance.
- Suggested fix: Allocate `task = &model.Task{}` before `Scan`, or use a local `var task model.Task` and return `&task`.

- Finding: `Scan` column count mismatch in `GetById` — `SELECT` had 7 columns but `Scan` had 6 arguments (missing `&task.Status`).
- Reason: Copy-paste or oversight when building the Scan argument list.
- Suggested fix: Add the missing `&task.Status` argument.

- Finding: Compilation error in Update handler — response body referenced `task` variable that was not in scope because service `Update` returned only `error`.
- Reason: Handler expected a task return value but service signature did not provide one.
- Suggested fix: Either return `(*model.Task, error)` from service Update (recommended, since repo already uses `RETURNING`), or omit task from the response.

- Finding: `UpdateTaskRequest.Description` had JSON tag `json:"title"` instead of `json:"description"` — a copy-paste error.
- Reason: Description field would never deserialize, silently overwriting the database value with an empty string.
- Suggested fix: Correct the JSON tag to `json:"description"`.

- Finding: Delete handler used `strconv.ParseInt("id", ...)` with a hardcoded string literal instead of `ctx.Param("id")`.
- Reason: Route parameter was never read; `ParseInt("id")` always errors, so Delete always returned 400.
- Suggested fix: Use `ctx.Param("id")`.

- Finding: Update handler stored the un-trimmed `title` instead of the trimmed version that passed validation.
- Reason: Validation used `strings.TrimSpace(title)` but the struct was built with the original `title`.
- Suggested fix: Use the trimmed value when constructing the `model.Task`.

- Finding: Update and Delete routes were not registered in `router.go`.
- Reason: Handler methods existed but were never wired to the Gin engine.
- Suggested fix: Call `h.RegisterUpdateTaskRoutes(r)` and `h.RegisterDeleteTaskRoutes(r)` in `SetupRouter`.

### Medium Priority (deferred to backlog)

- Finding: Inconsistent negative/zero ID handling across operations — `Delete` validates `id <= 0` and returns 400, but `GetByID` and `Update` do not validate and return 404.
- Reason: Missing validation in two of three service methods.
- Suggested fix: Standardize — either validate `id <= 0` in all three service methods, or let all three fall through to 404.
- Deferred: See TASK_BACKLOG.md improvement items.

- Finding: Delete handler used `err.Error()` in the 500 response, potentially leaking internal error details to the client.
- Reason: Inconsistent with other handlers that use generic messages.
- Suggested fix: Replace with a generic message like `"failed to delete task"`.
- Deferred: See TASK_BACKLOG.md improvement items.

### Low Priority (deferred to backlog)

- Finding: `model.Task` struct has no JSON tags, so API returns Go-style field names (`ID`, `UserID`) instead of REST conventions (`id`, `user_id`).
- Reason: JSON serialization uses struct field names by default.
- Suggested fix: Add `json:"..."` tags to all Task fields.
- Deferred: Will be addressed in T009 or T018 (DTO refactor).

- Finding: Update does not validate `status` field against a whitelist of allowed values.
- Reason: No status constraint exists yet.
- Suggested fix: Add allowed-status validation in service layer.
- Deferred: See TASK_BACKLOG.md improvement items.

## Strengths

- Three-layer error mapping is clean: repository defines `ErrTaskNotFound`, service translates it, handler maps to HTTP 404. Each layer has a distinct responsibility.
- Correctly used `sql.ErrNoRows` handling in `GetByID` and `Update` to distinguish "not found" from real database errors.
- Used `RowsAffected()` on Delete to detect when zero rows were deleted, returning a proper 404 instead of silently succeeding.
- Pagination refactor during this task was well-executed: `parseListPara` returns `*int` to distinguish "not provided" from "explicitly zero", and default values and range validation moved to the service layer.
- Learner asked insightful architectural questions about handler/service/repository responsibility boundaries and where query parameter parsing should live.

## Learning Notes

- Named return values with pointer types start as `nil` — must allocate before use.
- `database/sql` `Scan` argument count must exactly match the `SELECT` column list.
- `ctx.Param("id")` returns a string that must be parsed — always verify the value, not just the literal parameter name.
- Route registration in `SetupRouter` is a recurring omission — verify all handlers are wired before testing.
- `204 No Content` should not include a response body; use `ctx.Status(http.StatusNoContent)`.

## Follow-Up Tasks

- Task: Standardize negative/zero ID validation across all CRUD operations.
- Acceptance criteria: All three operations return the same HTTP status for invalid IDs.
- Task: Add JSON tags to model or introduce DTO layer.
- Acceptance criteria: API response uses snake_case field names.
- Task: Add status field validation.
- Acceptance criteria: Invalid status values are rejected with 400.
