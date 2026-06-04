# Stage 1: Go Task API

## Goal

Build a complete Go backend API service.

## Skills To Learn

- Go project structure
- Gin
- PostgreSQL
- GORM or database access basics
- JWT
- Password hashing
- Error handling
- Docker
- Basic tests

## Substages

### 1.1 Project Initialization

Deliverables:

- Go module
- Basic directory structure
- Gin server
- `/health` endpoint

Acceptance criteria:

- `go run ./cmd/server` starts successfully.
- `GET /health` returns HTTP 200.

### 1.2 Database Setup

Deliverables:

- PostgreSQL in Docker Compose
- Database connection
- Initial user and task models

Acceptance criteria:

- `docker compose up` starts PostgreSQL.
- Go service can connect to the database.

### 1.3 Task CRUD

Deliverables:

- Create task
- List tasks
- Get task detail
- Update task
- Delete task

Acceptance criteria:

- Task data is persisted in PostgreSQL.
- Errors are returned consistently.

### 1.4 User Auth

Deliverables:

- Register
- Login
- Password hashing
- JWT generation
- Auth middleware

Acceptance criteria:

- A user can register and log in.
- Protected APIs reject unauthenticated requests.

### 1.5 Testing And Documentation

Deliverables:

- Basic handler or service tests
- API examples
- README startup instructions

Acceptance criteria:

- Tests can be run locally.
- A new developer can start the service from README instructions.

