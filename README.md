# Backend AI Roadmap

This repository is a learning control center for building backend engineering skills with Go and Python.

The goal is not to let an agent generate a complete project. The goal is to use an agent as a coach, planner, reviewer, and debugging partner while the learner implements the core backend logic.

## Learning Direction

- Go: main backend engineering language.
- Python: AI service, agent workflow, automation, and data-processing language.
- Final target: build a multi-service backend where Go handles core business logic and Python handles AI workflows.

## Repository Map

- `AGENTS.md`: rules for how an AI agent should collaborate in this repository.
- `PROFILE.md`: learner background, preferences, weak areas, and long-term goals.
- `ROADMAP.md`: high-level learning stages.
- `SKILLS.md`: required backend skills and their learning status.
- `PROGRESS.md`: current stage, current substage, and completion state.
- `TASKS.md`: active and upcoming tasks with acceptance criteria.
- `TASK_BACKLOG.md`: long-term task pool for the full roadmap.
- `REFLECTIONS.md`: learner performance notes after each task.
- `DECISIONS.md`: technical decisions and reasons.
- `PROJECTS.md`: project descriptions and expected features.
- `stages/`: detailed breakdown for each learning stage.
- `sessions/`: per-session records for cross-session continuity.
- `reviews/`: code review notes.
- `templates/`: reusable templates for sessions, reviews, and reflections.

## Standard Session Workflow

At the beginning of each learning session, ask the agent to read:

```text
AGENTS.md
PROFILE.md
ROADMAP.md
SKILLS.md
PROGRESS.md
TASKS.md
TASK_BACKLOG.md
the relevant file under stages/
```

Then the agent should identify the current task, explain the implementation direction, and avoid directly implementing the core feature unless explicitly requested.

At the end of each session, ask the agent to update:

```text
PROGRESS.md
TASKS.md
TASK_BACKLOG.md when task status changes
SKILLS.md
REFLECTIONS.md
sessions/
reviews/ if code was reviewed
```

## Local Development

### Prerequisites
 - Go
 - Docker / Docker Compose

### Start PostgreSQL
```bash
docker compose up -d
```

### Start Server
```bash
go run ./cmd/server
```

### Configuration
| Env                    | Default        | Meaning                             |
|:------------------------:|:----------------:|:-------------------------------------:|
| SERVER_PORT            | 8080           | Port the server running             |
| DATABASE_HOST          | localhost      | PG host, localhost for docker image |
| DATABASE_PORT          | 5432           | PG port, 5432 for docker image      |
| DATABASE_USER          | taskapi        | PG user                             |
| DATABASE_PASSWORD      | taskapi        | PG password                         |
| DATABASE_NAME          | taskapi        | PG database name                       |
| JWT_SECRET             | dev-jwt-secret | secret to generate JWT              |
| JWT_EXPIRATION_MINUTES | 60             | JWT expiration time, default 60 min |

### Run Tests
in the root dict
```bash
go test ./...
```

## API Examples
### Health
```bash
curl -i http://localhost:8080/health
```

### Register
```bash
curl -i "http://localhost:8080/users/register" \
    -H "Content-Type: application/json" \
    -d '{"name": "name", "email": "example@ex.com", "password": "test"}'
```
Got Response:
```text
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Fri, 26 Jun 2026 08:23:09 GMT
Content-Length: 80

{"data":{"user":{"ID":17,"Name":"name","Email":"example@ex.com"}},"status":"ok"}
```

### Login
```bash
curl -i "http://localhost:8080/users/login" \
    -H "Content-Type: application/json" \
    -d '{"email": "example@ex.com", "password": "test"}'
```
this will return the JWT token
```text
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 26 Jun 2026 08:27:05 GMT
Content-Length: 236

{"data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNywiZXhwIjoxNzgyNDY2MDI1LCJpYXQiOjE3ODI0NjI0MjV9.gnmUdxI9jE25A6GNB89x_etS1WfNIe3K5K3ceT6ekK8","user":{"ID":17,"Name":"name","Email":"example@ex.com"}},"status":"ok"}
```

above URL DO NOT required token, below DOES  

### Create Task
```bash
curl -i "http://localhost:8080/tasks" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer your_token_here" \
    -d '{"title": "task1", "description": "just a test task"}'
```
Response:
```text
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Fri, 26 Jun 2026 08:30:58 GMT
Content-Length: 211

{"data":{"task":{"ID":8,"UserID":17,"Title":"task1","Description":"just a test task","Status":"todo","CreatedAt":"2026-06-26T16:30:58.216073+08:00","UpdatedAt":"2026-06-26T16:30:58.216073+08:00"}},"status":"ok"}
```
### List Task
default `limit=20`, `offset=0`, can be empty
```bash
curl -i "http://localhost:8080/tasks?limit=20&offset=0" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer your_token_here"
```

### Get Task
```bash
curl -i "http://localhost:8080/tasks/8" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer your_token_here"
```

### Update Task
```bash
curl -i -X PUT "http://localhost:8080/tasks/8" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer your_token_here" \
    -d '{"title": "test2", "description": "another", "status": "doing"}'
```

### Delete Task
```bash
curl -i -X DELETE "http://localhost:8080/tasks/8" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer your_token_here"
```

### Response Struct
when success, response may like this:
```json
{
  "status": "ok",
  "data": {}
}
```
when fail, response may like this:
```json
{
  "status": "error",
  "error": {
    "message": ""
  }
}
```

## Shut Down PG
after servering, use
```bash
docker compose down
```
to shutdown the PG
