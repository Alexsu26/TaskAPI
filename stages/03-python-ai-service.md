# Stage 3: Python AI Service

## Goal

Build a Python backend service for AI workflows.

## Skills To Learn

- FastAPI
- Pydantic
- SQLAlchemy or SQLModel
- pytest
- Redis, RQ, or Celery
- LLM API wrapper design
- Async task status tracking
- Docker

## Project: AI Summary Service

Core features:

- Accept text input.
- Create an AI summary task.
- Return `task_id`.
- Run summary generation asynchronously.
- Store task status and result.
- Expose result query API.

## Substages

### 3.1 FastAPI Project Setup

Deliverables:

- FastAPI app.
- Health check endpoint.
- Basic config.

### 3.2 Persistence

Deliverables:

- PostgreSQL connection.
- Summary task model.
- Repository layer.

### 3.3 AI Summary Flow

Deliverables:

- LLM client wrapper.
- Summary service.
- Result persistence.

### 3.4 Async Task Execution

Deliverables:

- Redis/RQ or Celery setup.
- Task status API.
- Worker process.

### 3.5 Tests And Docker

Deliverables:

- pytest tests.
- Dockerfile.
- Docker Compose.
- README instructions.

