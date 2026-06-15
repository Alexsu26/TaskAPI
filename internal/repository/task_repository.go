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

func (r *TaskRepo) List(limit, offset int) ([]*model.Task, error) {
	rows, err := r.db.Query(
		`select id, user_id, title, description, status, created_at, updated_at
		from tasks
		order by updated_at DESC, id DESC
		limit $1 offset $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]*model.Task, 0)
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(
			&task.ID,
			&task.UserID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}
