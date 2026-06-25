# Skills Checklist

This file tracks the backend skills that should be learned through the roadmap.

The checklist is not meant to be completed by reading only. A skill should be marked as learned only after it has been used in a project task, reviewed, and verified.

## Status Legend

- `[ ]`: not started
- `[~]`: practiced but not stable yet
- `[x]`: used in a project and reviewed

## Recent Evidence

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
- [ ] `context`
- [ ] `net/http`
- [x] Gin
- [ ] GORM or `sqlc`
- [x] PostgreSQL
- [ ] Redis
- [ ] JWT
- [~] Docker
- [ ] testing
- [ ] logging
- [ ] middleware
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
- [ ] Redis
- [~] authentication
- [ ] logging
- [ ] testing
- [~] Docker
- [ ] Linux basics
- [ ] CI/CD basics
- [ ] service-to-service communication
- [~] configuration management
- [ ] timeout handling
- [ ] retry handling
- [ ] rate limiting
- [ ] database indexes
- [~] API documentation

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
