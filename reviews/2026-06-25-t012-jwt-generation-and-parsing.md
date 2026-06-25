# Review: 2026-06-25 T012 JWT Generation And Parsing

## Result

Accepted after fixes. The implementation keeps JWT concerns in `internal/auth`, loads auth configuration through `internal/config`, returns a token from successful login, and satisfies the T012 acceptance criteria.

## Findings

1. Fixed after first review: JWT environment variable names now use conventional uppercase snake case: `JWT_SECRET` and `JWT_EXPIRATION_MINUTES`.
2. Fixed after first review: token parsing now rejects claims with non-positive `user_id`.
3. Fixed after first review: token parsing now distinguishes expired tokens with `ErrTokenExpired`.
4. Fixed after first review: JWT expiration minutes now reject invalid, zero, and negative values by falling back to the default.
5. Fixed after second review: token generation now rejects non-positive user IDs.
6. Fixed after third review: login response no longer nests `data` inside `data`, and registration response shape was restored.

## Verification

Commands run:

```text
gofmt -l cmd/server internal
go test ./...
go vet ./...
docker compose up -d postgres
SERVER_PORT=18080 JWT_SECRET=acceptance-secret JWT_EXPIRATION_MINUTES=60 go run ./cmd/server
```

Observed results:

```text
gofmt: no output
go test ./...: PASS for all packages
go vet: no output
docker compose: postgres running
GET /health: 200
POST /users/register valid request: 201, response includes user DTO without PasswordHash
POST /users/login valid credentials: 200, response includes user DTO and JWT token
POST /users/login wrong password: 401, invalid email or password
GET /tasks?limit=1&offset=0: 200
login token: three JWT segments, payload includes user_id, exp, and iat
```

## Notes

- `internal/auth.TokenManager` correctly owns JWT generation and parsing.
- `cmd/server/main.go` constructs the token manager from config and injects it into `UserService`, which keeps config and auth responsibilities separated.
- `ParseToken` validates signing method, signature, expiration, token validity, and positive user ID. Database user existence should remain a middleware or service concern in later tasks, not a token helper concern.
- Login now returns a JWT but task routes are not protected yet, which matches T012 scope. T013 should add middleware; T014 should restrict task rows to the current user.
