# Stage 2: Go Engineering

## Goal

Improve the Go Task API into a more production-like service.

## Skills To Learn

- Database migrations
- Request validation
- Structured error handling
- Structured logging
- Request ID middleware
- Redis basics
- Integration tests
- Background worker basics
- Graceful shutdown

## Substages

### 2.1 Database Migration

Deliverables:

- Add migration tool.
- Replace ad hoc schema creation with versioned migrations.

### 2.2 Validation And Error Handling

Deliverables:

- Separate request DTOs from database models.
- Add validation rules.
- Add unified application error type.

### 2.3 Logging And Observability Basics

Deliverables:

- Structured logger.
- Request ID.
- Request duration logs.
- Panic recovery.

### 2.4 Testing

Deliverables:

- Service tests.
- Repository tests.
- Handler/API tests.

### 2.5 Redis

Deliverables:

- Redis in Docker Compose.
- One practical Redis use case such as token blacklist, query cache, or simple rate limiting.

### 2.6 Background Tasks

Deliverables:

- Simple worker process.
- Task event creation and consumption.

