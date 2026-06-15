package service

import (
	"errors"
	"fmt"
	"strings"

	"taskapi/internal/model"
	"taskapi/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepo
}

func NewTaskService(repo *repository.TaskRepo) *TaskService {
	return &TaskService{repo: repo}
}

var (
	ErrTitleRequired = errors.New("title is required")
	ParaInvalid      = errors.New("error parameters")
)

func (s *TaskService) Create(title, description string) (*model.Task, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return nil, ErrTitleRequired
	}

	task := &model.Task{
		UserID:      1, // 暂时固定user_id
		Title:       title,
		Description: description,
		Status:      "todo",
	}
	if err := s.repo.Create(task); err != nil {
		return nil, fmt.Errorf("create task error: %w", err)
	}
	return task, nil
}

func (s *TaskService) List(limit, offset int) ([]*model.Task, error) {
	if limit <= 0 || offset < 0 {
		return nil, ParaInvalid
	}
	tasks, err := s.repo.List(limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list tasks error :%w", err)
	}
	return tasks, nil
}
