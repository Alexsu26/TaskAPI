# Task Backlog

This file is the long-term task pool for the full Go + Python backend learning roadmap.

`TASKS.md` should contain the active task and the next few near-term tasks. This file should contain the broader route. Do not over-expand every item here in advance. When a backlog item becomes active, move or copy it into `TASKS.md` and expand it with detailed acceptance criteria.

## Usage Rules

- Keep one active task in `TASKS.md`.
- Use this file to choose the next task after the active task is completed.
- Split a backlog task only when it is about to become active.
- Adjust future tasks based on actual learning progress and review results.
- Do not mark a task done here unless there is evidence in `PROGRESS.md`, `REFLECTIONS.md`, or `sessions/`.

## Status Legend

- `[ ]`: not started
- `[~]`: in progress
- `[x]`: completed and verified

## Stage 1: Go Task API

### T001: Initialize Go Gin Project

Status: `[x]`

Objective:

- Create a minimal Go Gin backend service with a `/health` endpoint.

Skills:

- Gin
- HTTP
- REST API
- Go project structure

### T002: Add Basic Project Structure

Status: `[~]`

Objective:

- Introduce initial package boundaries for config, router, handler, service, repository, and model.

Skills:

- Go package boundaries
- `struct`
- `interface`
- responsibility separation

### T003: Add Configuration Management

Status: `[ ]`

Objective:

- Add environment-based configuration for server and database settings.

Skills:

- configuration management
- environment variables
- error handling

### T004: Add PostgreSQL With Docker Compose

Status: `[ ]`

Objective:

- Run PostgreSQL locally and connect the Go service to it.

Skills:

- PostgreSQL
- SQL
- Docker
- database connection handling

### T005: Design User And Task Models

Status: `[ ]`

Objective:

- Define initial user and task models and understand how they map to database tables.

Skills:

- Go `struct`
- database modeling
- SQL schema basics

### T006: Implement Task Creation

Status: `[ ]`

Objective:

- Implement the API flow for creating a task.

Skills:

- REST API
- handler/service/repository boundaries
- database insert
- validation

### T007: Implement Task List Query

Status: `[ ]`

Objective:

- Implement task listing with basic pagination.

Skills:

- REST API
- SQL query basics
- pagination
- response design

### T008: Implement Task Detail, Update, And Delete

Status: `[ ]`

Objective:

- Implement task detail lookup, update, and deletion.

Skills:

- REST API
- SQL update/delete
- error handling
- status code design

### T009: Add Unified Response And Error Handling

Status: `[ ]`

Objective:

- Standardize API responses and application errors.

Skills:

- error handling
- HTTP status codes
- response design

### T010: Implement User Registration

Status: `[ ]`

Objective:

- Add user registration with password hashing.

Skills:

- authentication
- password hashing
- validation
- database uniqueness

### T011: Implement User Login

Status: `[ ]`

Objective:

- Add login with password verification.

Skills:

- authentication
- error handling
- security basics

### T012: Implement JWT Generation And Parsing

Status: `[ ]`

Objective:

- Generate JWT after login and parse JWT for later protected APIs.

Skills:

- JWT
- configuration
- security basics

### T013: Add Auth Middleware

Status: `[ ]`

Objective:

- Protect private routes with JWT middleware.

Skills:

- middleware
- request context
- authorization

### T014: Restrict Tasks To The Current User

Status: `[ ]`

Objective:

- Ensure users can only access their own tasks.

Skills:

- authorization
- SQL filtering
- service-layer validation

### T015: Add Basic Tests

Status: `[ ]`

Objective:

- Add tests for the most important service or handler behavior.

Skills:

- Go testing
- test data setup
- behavior verification

### T016: Complete Stage 1 Documentation

Status: `[ ]`

Objective:

- Add startup instructions and API examples.

Skills:

- API documentation
- README writing
- local verification

## Stage 2: Go Engineering

### T017: Add Database Migrations

Status: `[ ]`

Objective:

- Replace ad hoc schema setup with versioned migrations.

Skills:

- migration workflow
- SQL
- schema versioning

### T018: Refactor DTO, Model, And Response Boundaries

Status: `[ ]`

Objective:

- Separate request DTOs, response DTOs, and database models.

Skills:

- responsibility separation
- data modeling
- validation

### T019: Add Structured Logging

Status: `[ ]`

Objective:

- Add structured logs for server startup, requests, and errors.

Skills:

- logging
- observability basics

### T020: Add Request ID And Panic Recovery

Status: `[ ]`

Objective:

- Add request tracing basics and safer panic handling.

Skills:

- middleware
- logging
- error handling

### T021: Add Service Layer Tests

Status: `[ ]`

Objective:

- Test business logic independently from HTTP handlers.

Skills:

- testing
- dependency boundaries
- `interface`

### T022: Add Repository Integration Tests

Status: `[ ]`

Objective:

- Test database behavior against a real or containerized database.

Skills:

- integration testing
- PostgreSQL
- test isolation

### T023: Add Redis

Status: `[ ]`

Objective:

- Add Redis to the local environment and connect the Go service to it.

Skills:

- Redis
- Docker
- configuration

### T024: Implement One Redis Use Case

Status: `[ ]`

Objective:

- Implement token blacklist, query cache, simple rate limiting, or another focused Redis use case.

Skills:

- Redis
- cache design
- rate limiting or token invalidation

### T025: Add Background Worker Or Async Task

Status: `[ ]`

Objective:

- Add a small asynchronous workflow using goroutine/channel first or Redis-backed queue later.

Skills:

- goroutine
- channel
- async workflow

### T026: Add Graceful Shutdown

Status: `[ ]`

Objective:

- Shut down HTTP server and dependencies cleanly.

Skills:

- `context`
- signal handling
- graceful shutdown

## Stage 3: Python AI Service

### T027: Initialize FastAPI AI Service

Status: `[ ]`

Objective:

- Create a minimal FastAPI service with health check.

Skills:

- FastAPI
- Python project structure
- HTTP

### T028: Add Pydantic Settings And Schemas

Status: `[ ]`

Objective:

- Add typed configuration and request/response schemas.

Skills:

- Pydantic
- configuration management
- validation

### T029: Add PostgreSQL And Task Model

Status: `[ ]`

Objective:

- Persist AI task requests and results.

Skills:

- SQLAlchemy or SQLModel
- PostgreSQL
- repository pattern

### T030: Implement AI Summary Task Creation

Status: `[ ]`

Objective:

- Create an API that accepts text and returns a task id.

Skills:

- FastAPI
- REST API
- task modeling

### T031: Add LLM API Call Wrapper

Status: `[ ]`

Objective:

- Wrap LLM calls behind a service interface.

Skills:

- LLM API call wrapper
- error handling
- timeout handling

### T032: Add Celery Or RQ Async Task Execution

Status: `[ ]`

Objective:

- Execute summary generation asynchronously.

Skills:

- Celery or RQ
- Redis
- async processing

### T033: Add Task Status And Result Query

Status: `[ ]`

Objective:

- Query AI task status and result.

Skills:

- REST API
- database query
- status modeling

### T034: Add Pytest And Docker

Status: `[ ]`

Objective:

- Add tests and containerized local startup.

Skills:

- pytest
- Docker
- local verification

## Stage 4: Go + Python Integrated System

### T035: Design Go/Python Service Contract

Status: `[ ]`

Objective:

- Define how Go and Python services communicate.

Skills:

- service-to-service communication
- API contract design
- error contract design

### T036: Let Go Call Python Summary API

Status: `[ ]`

Objective:

- Add direct HTTP communication from Go to Python.

Skills:

- `net/http`
- timeout handling
- service integration

### T037: Implement AI Task State Synchronization

Status: `[ ]`

Objective:

- Keep Go-side business state aligned with Python-side AI task state.

Skills:

- state modeling
- service boundaries
- error handling

### T038: Upgrade To Async Queue Communication

Status: `[ ]`

Objective:

- Move from direct synchronous calls to queue-based async workflow.

Skills:

- Redis or queue workflow
- retry handling
- distributed task design

### T039: Implement Document Summary And Keyword Extraction

Status: `[ ]`

Objective:

- Add document-oriented AI analysis.

Skills:

- agent workflow service design
- LLM API usage
- result modeling

### T040: Implement Simple Document Q&A

Status: `[ ]`

Objective:

- Add a simple question-answer flow over submitted document content.

Skills:

- retrieval basics
- service API design
- Python AI workflow

### T041: Add Cross-Service Timeout, Retry, And Error Handling

Status: `[ ]`

Objective:

- Make service communication more robust.

Skills:

- timeout handling
- retry handling
- error contract design

### T042: Final Documentation And Architecture Diagram

Status: `[ ]`

Objective:

- Write final README, API examples, and architecture notes.

Skills:

- API documentation
- architecture explanation
- project packaging
