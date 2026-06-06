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

## Local Run Command
start the repo:
```bash
go run ./cmd/server
```

test local health:
```bash
curl -i http://localhost:8080/health
```