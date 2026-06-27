# Skills Checklist

This file tracks the backend skills that should be learned through the roadmap.

The checklist is not meant to be completed by reading only. A skill should be marked as learned only after it has been used in a project task, reviewed, and verified.

## Status Legend

- `[ ]`: not started
- `[~]`: practiced but not stable yet
- `[x]`: used in a project and reviewed

## Recent Evidence

- 2026-06-27 T017 practiced database migrations, SQL schema versioning, and local verification. Learner moved the current `users` and `tasks` schema out of the unversioned startup helper into `migrations/000001_create_users_and_tasks.up.sql` plus a matching down migration, removed the `database.RunMigrations` startup dependency, and documented the `migrate` CLI workflow. Review first found `tasks.user_id` incorrectly declared as `bigserial`; learner fixed it to `bigint`. Verification passed `go test ./...`, `go vet ./...`, `migrate up`, `migrate down -all`, and PostgreSQL schema inspection in an isolated temporary database. Database migrations and schema versioning are now `[x]` for the initial Stage 2 scope.
- 2026-06-26 T016 practiced API documentation and local verification. Learner updated `README.md` with local startup steps, current configuration defaults, public route examples, authenticated task route examples using `Authorization: Bearer <token>`, response envelope examples, and `go test ./...` instructions. Review first found non-runnable examples for PUT continuation, pagination placeholders, task ID placeholders, and one database configuration wording issue; learner fixed them. Verification ran `go test ./...`. API documentation is now `[x]` for Stage 1's current scope.
- 2026-06-26 T015 practiced Go testing and behavior verification. Learner added `internal/auth/token_test.go` using the standard `testing` package, verified JWT generation plus parsing preserves `user_id`, and covered error paths for invalid user ID and malformed token parsing with `errors.Is`. Review verified the test is deterministic, does not require a live server or database, and passes `gofmt`, `go test ./...`, and `go vet ./...`. Testing is now `[~]` because this is the first reviewed unit-test coverage, with broader handler/service/integration testing still to practice.
- 2026-06-26 T014 practiced authorization, Gin request context, and SQL ownership filtering. Learner threaded `current_user_id` from Gin handlers through task service methods into repository SQL, removed the temporary hard-coded `UserID: 1`, filtered list/detail/update/delete by `user_id`, and preserved `404 Not Found` for cross-user task access so task existence is not exposed. First review found a missing `return` after two handler auth failures and an `UPDATE` SQL typo (`userID` instead of `user_id`); learner fixed both. Verification passed `gofmt`, `go test ./...`, `go vet ./...`, and runtime checks for two-user create/list/detail/update/delete isolation plus unauthenticated rejection. Authentication is now `[x]` for Stage 1's current scope.
- 2026-06-25 T013 practiced Gin auth middleware and request context. Learner added an `internal/middleware` auth boundary, protected task routes through a Gin route group, parsed `Authorization: Bearer <token>`, validated JWTs with the existing token manager, returned unified HTTP 401 responses for missing, malformed, invalid, and expired tokens, and stored `current_user_id` in Gin context. Review first caught an import-cycle issue between `handler` and `router`; learner fixed it by moving the route registration helper type out of the router dependency path. Verification passed `gofmt`, `go test ./...`, `go vet ./...`, and runtime checks for public routes, missing/malformed/invalid/expired token rejection, and valid-token task access. Middleware is now `[x]`; authentication remains `[~]` until T014 current-user task ownership is complete.
- 2026-06-25 T012 practiced JWT generation and parsing. Learner added `JWT_SECRET` and `JWT_EXPIRATION_MINUTES` configuration, introduced an `internal/auth` token manager, generated `HS256` JWTs with `user_id`, `exp`, and `iat` claims, validated signing method, signature, expiration, and positive user ID during parsing, injected token configuration through `main` into the user service, and returned a token from successful login without exposing `PasswordHash`. Review verified `gofmt`, `go test ./...`, `go vet ./...`, and runtime checks for `/health`, registration, login token response, wrong-password login, and task listing. JWT is now `[x]`; authentication remains `[~]` until middleware and current-user authorization are complete.
- 2026-06-24 T011 practiced user login and password verification. Learner added `POST /users/login`, preserved handler/service/repository boundaries, added email lookup with `sql.ErrNoRows` mapped to a repository not-found error, verified passwords with `bcrypt.CompareHashAndPassword`, unified wrong-email and wrong-password responses as HTTP 401 without revealing which field was wrong, and kept `PasswordHash` out of the login response. Review verified `gofmt`, `go test ./...`, `go vet ./...`, and runtime checks for login success, wrong password, wrong email, missing password, whitespace-only password, `/health`, and task listing. Authentication remains `[~]` until JWT generation, middleware, and current-user task ownership are complete.
- 2026-06-24 T010 practiced the first auth workflow with user registration and password hashing. Learner added `POST /users/register`, wired handler/service/repository boundaries, hashed passwords with bcrypt, handled duplicate email with PostgreSQL unique violation code `23505` mapped to HTTP 409, returned a DTO without `PasswordHash`, and fixed a double-response bug caused by a missing `return` after an error response. Review verified `gofmt`, `go test ./...`, `go vet ./...`, and runtime checks for registration success, duplicate email, missing password, whitespace-only fields, `/health`, and task listing.
- 2026-06-23 T009 practiced unified response design and centralized HTTP error handling. Learner introduced success/error response helpers, centralized service error mapping, kept HTTP parsing/binding errors as 400 responses, preserved `POST /tasks` as 201, and verified health, CRUD success paths, invalid body, invalid query, invalid ID, and not-found paths. Review verified `gofmt`, `go vet`, `go test ./...`, and runtime curl checks all pass. Error handling is now `[x]` for current Stage 1 CRUD scope, while later auth-specific error cases will be practiced again in T010-T014.
- 2026-06-23 T008 practiced REST CRUD (detail, update, delete), route parameters with `ctx.Param`, handler/service/repository boundaries for `SELECT`/`UPDATE`/`DELETE`, `sql.ErrNoRows` handling, `RowsAffected` for delete, `UPDATE ... RETURNING`, nil pointer prevention with named return values, pagination refactoring (`*int` pointers for default-value ownership in service layer), and HTTP status code design (200, 204, 400, 404, 500). Review verified `go build`, `go vet`, and `go test ./...` all pass clean. Error handling improved but remains `[~]` — ID validation inconsistency and internal error leakage deferred to backlog (IMP-001, IMP-004).

## Go Required Skills

- [x] `struct` and data modeling
- [ ] `interface` and dependency boundaries
- [x] error handling
- [x] Go package boundaries
- [x] responsibility separation
- [ ] goroutine
- [ ] channel
- [~] `context`
- [ ] `net/http`
- [x] Gin
- [ ] GORM or `sqlc`
- [x] PostgreSQL
- [x] database migrations
- [ ] Redis
- [x] JWT
- [~] Docker
- [~] testing
- [ ] logging
- [x] middleware
- [ ] graceful shutdown

## Python Required Skills

- [ ] FastAPI
- [ ] Pydantic
- [ ] SQLAlchemy or SQLModel
- [ ] pytest
- [ ] Celery or RQ
- [ ] `asyncio` basics
- [ ] LLM API call wrapper
- [ ] agent workflow service design
- [ ] configuration management
- [ ] Docker

## General Backend Required Skills

- [x] HTTP
- [x] REST API
- [x] SQL
- [x] schema versioning
- [ ] Redis
- [x] authentication
- [ ] logging
- [~] testing
- [~] Docker
- [ ] Linux basics
- [ ] CI/CD basics
- [ ] service-to-service communication
- [~] configuration management
- [ ] timeout handling
- [ ] retry handling
- [ ] rate limiting
- [ ] database indexes
- [x] API documentation

## Skill-To-Stage Mapping

### Stage 1: Go Task API

Primary skills:

- Go `struct`
- Go error handling
- Gin
- PostgreSQL
- GORM or `sqlc`
- JWT
- Docker
- testing
- middleware
- HTTP
- REST API
- SQL
- authentication
- API documentation

### Stage 2: Go Engineering

Primary skills:

- Go `interface`
- goroutine
- channel
- `context`
- Redis
- logging
- graceful shutdown
- database indexes
- timeout handling
- retry handling
- rate limiting
- service testing
- integration testing

### Stage 3: Python AI Service

Primary skills:

- FastAPI
- Pydantic
- SQLAlchemy or SQLModel
- pytest
- Celery or RQ
- `asyncio` basics
- LLM API call wrapper
- agent workflow service design
- Python configuration management
- Docker

### Stage 4: Go + Python Integrated System

Primary skills:

- service-to-service communication
- timeout handling
- retry handling
- Redis or queue-based async workflow
- cross-service API design
- API documentation
- CI/CD basics
- Linux basics

## How Agents Should Use This File

At the beginning of each task, identify which skills from this file the task will practice.

At the end of each task, update this file only when there is evidence that the learner practiced or demonstrated a skill.

Do not mark a skill as complete just because code exists. Mark it as complete only when the learner can explain the concept and the implementation has been reviewed or verified.
