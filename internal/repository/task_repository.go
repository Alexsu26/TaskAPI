package repository

import (
	"database/sql"
	"errors"

	"taskapi/internal/model"
)

type TaskRepo struct {
	db *sql.DB
}

var ErrTaskNotFound = errors.New("task not found")

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

func (r *TaskRepo) GetByID(id int64) (task *model.Task, err error) {
	task = &model.Task{}
	err = r.db.QueryRow(
		`select id, user_id, title, description, status, created_at, updated_at
		from tasks
		where id = $1`, id).Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}
	return task, nil
}

func (r *TaskRepo) Update(task *model.Task) error {
	err := r.db.QueryRow(
		`update tasks set title = $2, description = $3, status = $4, updated_at = now()
		where id = $1
		returning id, user_id, title, description, status, created_at, updated_at`, task.ID, task.Title, task.Description, task.Status).
		Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrTaskNotFound
		}
		return err
	}
	return nil
}

func (r *TaskRepo) Delete(id int64) error {
	result, err := r.db.Exec(
		`delete from tasks where id = $1`, id)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return ErrTaskNotFound
	}
	return nil
}
