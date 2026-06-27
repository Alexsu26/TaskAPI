# Review: T017 Add Database Migrations

Date: 2026-06-27

## Scope

Reviewed the learner's T017 implementation for database migration structure, schema accuracy, startup behavior, README workflow, and verification against the acceptance criteria.

## Result

Approved after one fix.

## Findings

### Fixed: `tasks.user_id` used `bigserial`

The first migration version declared `tasks.user_id` as `bigserial`. That changed the current schema semantics because `user_id` is a foreign key to an existing user, not an independently auto-incrementing value.

Expected:

```sql
user_id bigint not null references users(id)
```

The learner fixed this in `migrations/000001_create_users_and_tasks.up.sql`.

## Accepted Changes

- Added versioned SQL migration files under `migrations/`.
- Captured the current `users` and `tasks` schema in the first up migration.
- Added a down migration that drops `tasks` before `users`.
- Removed the startup call to the old unversioned `database.RunMigrations` helper.
- Deleted `internal/database/migration.go`.
- Updated `README.md` with the `migrate` CLI workflow.
- Preserved existing Gin route wiring and handler/service/repository behavior.

## Verification

Commands run:

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

- `go test ./...` passed.
- `go vet ./...` passed.
- `migrate up` applied version `1/u create_users_and_tasks`.
- PostgreSQL schema inspection confirmed `tasks.user_id` is `bigint` and references `users(id)`.
- `schema_migrations` reported `version=1` and `dirty=false`.
- `migrate down -all` succeeded.
- The server started successfully against a migrated temporary database.
- Runtime checks returned `200 OK` for `/health`, `201 Created` for registration, `200 OK` for login, and `201 Created` for authenticated task creation.
- The temporary review database was removed.

## Notes For Next Task

T018 should focus on separating request DTOs, response DTOs, and database models. In particular, task responses currently expose Go struct field names such as `ID` and `UserID`, which should become an intentional API contract.
