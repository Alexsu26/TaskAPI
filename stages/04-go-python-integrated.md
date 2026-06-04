# Stage 4: Go + Python Integrated System

## Goal

Build a multi-service backend where Go handles core business logic and Python handles AI workflows.

## System Shape

```text
API Client
  |
  v
Go Main Backend
  |
  +-- PostgreSQL
  +-- Redis
  +-- Python AI Service
```

## Go Responsibilities

- User system
- Authentication
- Document/task records
- Authorization
- API orchestration
- Calling Python service
- Saving business state

## Python Responsibilities

- Text processing
- LLM calls
- Summary generation
- Keyword extraction
- Embedding
- Simple retrieval or Q&A

## Substages

### 4.1 Multi-Service Local Setup

Deliverables:

- Go service.
- Python service.
- PostgreSQL.
- Redis if needed.
- Docker Compose startup.

### 4.2 HTTP-Based Service Communication

Deliverables:

- Go calls Python through HTTP.
- Timeout handling.
- Error handling.

### 4.3 AI Task Workflow

Deliverables:

- Go creates AI task.
- Python processes the task.
- Go exposes task status and result APIs.

### 4.4 Async Upgrade

Deliverables:

- Replace direct synchronous flow with Redis Stream, queue, or worker-based async flow.

### 4.5 Document Q&A

Deliverables:

- Basic document ingestion.
- Summary and keyword extraction.
- Simple question-answer endpoint.

### 4.6 Final Packaging

Deliverables:

- Full README.
- Architecture diagram.
- API examples.
- Tests for critical flows.

