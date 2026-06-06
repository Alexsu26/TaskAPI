# Skills Checklist

This file tracks the backend skills that should be learned through the roadmap.

The checklist is not meant to be completed by reading only. A skill should be marked as learned only after it has been used in a project task, reviewed, and verified.

## Status Legend

- `[ ]`: not started
- `[~]`: practiced but not stable yet
- `[x]`: used in a project and reviewed

## Go Required Skills

- [ ] `struct` and data modeling
- [ ] `interface` and dependency boundaries
- [ ] error handling
- [x] Go package boundaries
- [~] responsibility separation
- [ ] goroutine
- [ ] channel
- [ ] `context`
- [ ] `net/http`
- [x] Gin
- [ ] GORM or `sqlc`
- [ ] PostgreSQL
- [ ] Redis
- [ ] JWT
- [ ] Docker
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
- [~] REST API
- [ ] SQL
- [ ] Redis
- [ ] authentication
- [ ] logging
- [ ] testing
- [ ] Docker
- [ ] Linux basics
- [ ] CI/CD basics
- [ ] service-to-service communication
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
