# Review: 2026-06-26 T015 Add Basic Tests

## Scope

Reviewed learner implementation for T015 basic tests.

Files reviewed:

- `internal/auth/token_test.go`

Focus:

- Go test file structure
- behavior coverage
- success and error paths
- deterministic execution
- T015 acceptance criteria

## Findings

### High Priority

- None.

### Medium Priority

- None.

### Low Priority

- Finding: Expiration and wrong-secret parsing behavior are not covered yet.
- Reason: Current T015 only requires a first meaningful test file with success and error paths.
- Suggested fix: Keep this for a later test-expansion task after the basic testing workflow is comfortable.

## Strengths

- The test file uses the correct `_test.go` naming convention and `package auth`.
- `TestTokenManager_GenerateAndParseToken` verifies real behavior by generating a token, parsing it, and checking the `UserID`.
- `TestTokenManager_InvalidUserID` verifies invalid generation input maps to `ErrTokenInvalid`.
- `TestTokenManager_InvalidToken` verifies malformed token parsing maps to `ErrTokenInvalid`.
- Error checks use `errors.Is`, which is the right habit for sentinel errors.
- The tests do not depend on a live HTTP server or database.

## Verification

Static checks:

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

## Final Review Result

Accepted.

T015 acceptance criteria are met:

- At least one meaningful Go test file was added.
- The tests verify behavior instead of only calling functions.
- Success and error paths are covered.
- The tested behavior is tied to Stage 1 JWT/auth logic.
- `go test ./...` passes.
- Existing runtime behavior was not changed.

## Follow-Up Tasks

- T016 should complete Stage 1 documentation and include `go test ./...` in the README.
- Later testing tasks can add table-driven tests, handler tests with `httptest`, and service tests after dependency boundaries are easier to fake.
