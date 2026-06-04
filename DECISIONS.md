# Technical Decisions

This file records technical decisions so future sessions do not repeat the same discussion.

## D001: Use Go As The Main Backend Language

Decision:

- Use Go for the main backend service.

Reason:

- Go is well suited for API services, concurrency, deployment, and cloud-native backend work.
- It complements Python AI services well.

Status:

- Accepted

## D002: Use Python For AI Services

Decision:

- Use Python for AI, agent workflow, text processing, and data-oriented services.

Reason:

- Python has stronger AI, data, and LLM ecosystem support.
- It allows the Go service to remain focused on core backend responsibilities.

Status:

- Accepted

## D003: Agent Acts As Coach And Reviewer

Decision:

- The agent should guide, review, explain, and debug instead of directly implementing all core code.

Reason:

- The goal is durable backend skill acquisition, not only producing a working repository.

Status:

- Accepted

## D004: Start With Gin For The First Go API

Decision:

- Use Gin for the first Go backend project.

Reason:

- Gin is simple, widely used, and suitable for learning routing, middleware, request binding, and response handling.

Alternatives:

- Echo
- Fiber
- Standard `net/http`

Status:

- Accepted

