# Review: T005 Design User And Task Models

## Scope

Files or modules reviewed:

- `internal/model/user.go`
- `internal/model/task.go`

## Findings

### High Priority

- None.

### Medium Priority

- None.

### Low Priority

- Finding: `Task.Status` is currently a free-form `string`.
- Reason: This is acceptable for T005 model design, but later API input should not allow arbitrary statuses.
- Suggested fix: In a later validation task, constrain allowed values such as `todo`, `doing`, and `done` through request validation or model constants.

## Strengths

- Model structs are placed in the existing `internal/model` package.
- `User` includes fields needed for upcoming registration and login: `ID`, `Email`, `Name`, `PasswordHash`, `CreatedAt`, and `UpdatedAt`.
- `Task` includes fields needed for upcoming task CRUD: `ID`, `UserID`, `Title`, `Description`, `Status`, `CreatedAt`, and `UpdatedAt`.
- `Task.UserID` now matches `User.ID`, so the future SQL relationship can map cleanly from `tasks.user_id` to `users.id`.
- Existing Gin router and `/health` handler code were not changed.
- No repository, service, auth, CRUD, or migration code was added early.

## Learning Notes

- Related model IDs should use consistent types so future SQL joins and foreign keys are straightforward.
- A model struct should describe persisted data. Request validation and API response shaping can be added separately when the related endpoint is built.
- A plain string status is a reasonable early model choice, but the accepted values should be enforced before exposing write APIs broadly.

## Verification

Commands:

```text
gofmt -l cmd/server internal
go test ./...
```

Results:

```text
gofmt produced no output.
go test ./... passed for all current packages.
```

## Follow-Up Tasks

- Task: T006 Implement Task Creation.
- Acceptance criteria: add `POST /tasks` while keeping handler, service, and repository boundaries clear and preserving `/health`.
