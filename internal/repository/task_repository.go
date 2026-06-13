package repository

import (
	"database/sql"

	"taskapi/internal/model"
)

type TaskRepo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) Create(task *model.Task) error {
	err := r.db.QueryRow(
		`insert into tasks (user_id, title, description, status) values ($1, $2, $3, $4)
		returning id, created_at, updated_at`, task.UserID, task.Title, task.Description, task.Status).
		Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
	return err
}
