# Session: 2026-06-27 T017 Add Database Migrations

## Context Read

- `AGENTS.md`
- `PROFILE.md`
- `ROADMAP.md`
- `SKILLS.md`
- `PROGRESS.md`
- `TASKS.md`
- `TASK_BACKLOG.md`
- `stages/02-go-engineering.md`
- `sessions/2026-06-26-t016-complete-stage-1-documentation.md`

## Task

T017: Add database migrations.

## Work Completed

- Explained what database migration means and how it differs from ORM behavior.
- Recommended a small migration workflow using `golang-migrate` CLI and SQL files.
- Reviewed the learner's implementation.
- First review found `tasks.user_id` declared as `bigserial`, which changed the intended schema semantics.
- Re-reviewed after the learner fixed `user_id` to `bigint`.
- Verified the migration workflow in an isolated temporary PostgreSQL database.
- Accepted T017.
- Updated progress, task, backlog, skill, reflection, review, and session records.

## Evidence

Commands:

```text
go test ./...
go vet ./...
docker exec taskapi-database dropdb --if-exists -U taskapi taskapi_migration_review
docker exec taskapi-database createdb -U taskapi taskapi_migration_review
migrate -path migrations -database "postgres://taskapi:taskapi@localhost:5432/taskapi_migration_review?sslmode=disable" up
docker exec taskapi-database psql -U taskapi -d taskapi_migration_review -c "\d users"
docker exec taskapi-database psql -U taskapi -d taskapi_migration_review -c "\d tasks"
docker exec taskapi-database psql -U taskapi -d taskapi_migration_review -c "select version, dirty from schema_migrations;"
migrate -path migrations -database "postgres://taskapi:taskapi@localhost:5432/taskapi_migration_review?sslmode=disable" down -all
SERVER_PORT=18080 DATABASE_NAME=taskapi_migration_review go run ./cmd/server
curl -s -i http://localhost:18080/health
curl -s -i "http://localhost:18080/users/register" ...
curl -s "http://localhost:18080/users/login" ...
curl -s -i "http://localhost:18080/tasks" ...
docker exec taskapi-database dropdb --if-exists -U taskapi taskapi_migration_review
```

Results:

```text
go test ./...: PASS for all packages
go vet ./...: PASS
migrate up: 1/u create_users_and_tasks
schema_migrations: version=1 dirty=false
migrate down -all: 1/d create_users_and_tasks
runtime migrated DB checks: health 200, register 201, login 200, create task 201
```

## Learner Performance

Strengths:

- Understood the distinction between schema migration and ORM after clarification.
- Moved schema setup out of application startup and into versioned SQL files.
- Put migrations under the repository-level `migrations/` directory and aligned README commands with that location.
- Fixed the schema-type review finding quickly.

Weaknesses:

- Initial migration changed a foreign-key column from `bigint` to `bigserial`, showing that schema migration needs exact comparison with the current database contract.
- Migration documentation should mention required tooling clearly when a new external CLI is introduced.

## Next Task

T018: Refactor DTO, model, and response boundaries.
