package database

import "database/sql"

func RunMigrations(db *sql.DB) error {
	_, err := db.Exec(`
		create table if not exists users (
			id		BIGSERIAL PRIMARY KEY,
			email	TEXT NOT NULL UNIQUE,
			name	TEXT NOT NULL,
			password_hash	TEXT NOT NULL,
			created_at	TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at	TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);

		create table if not exists tasks (
			id		BIGSERIAL PRIMARY KEY,
			user_id	BIGINT NOT NULL REFERENCES users(id),
			title	TEXT NOT NULL,
			description	TEXT NOT NULL DEFAULT '',
			status	TEXT NOT NULL DEFAULT 'todo',
			created_at	TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at	TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);

		insert into users (id, email, name, password_hash) values 
		(1, 'test@example.com', 'dev user', 'dev')
		on conflict (id) do nothing;
	`)
	return err
}
