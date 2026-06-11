package database

import (
	"database/sql"
	"fmt"
	"taskapi/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {
	conn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	db, err := sql.Open("pgx", conn)
	if err != nil {
		return nil, fmt.Errorf("open postgres db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping postgres db: %w", err)
	}
	return db, nil
}
