package service

import (
	"errors"
	"fmt"
	"strings"

	"taskapi/internal/model"
	"taskapi/internal/repository"
)

type TaskRepository interface {
	Create(task *model.Task) error
	List(userID int64, limit, offset int) ([]*model.Task, error)
	GetByID(userID, id int64) (task *model.Task, err error)
	Update(task *model.Task) error
	Delete(userID, id int64) error
}

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

var (
	ErrTitleRequired = errors.New("title is required")
	ErrParaInvalid   = errors.New("error parameters")
	ErrTaskNotFound  = errors.New("task not found")
)

func (s *TaskService) Create(userID int64, title, description string) (*model.Task, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return nil, ErrTitleRequired
	}
	if userID <= 0 {
		return nil, ErrParaInvalid
	}
	task := &model.Task{
		UserID:      userID,
		Title:       title,
		Description: description,
		Status:      "todo",
	}
	if err := s.repo.Create(task); err != nil {
		return nil, fmt.Errorf("create task error: %w", err)
	}
	return task, nil
}

func (s *TaskService) List(userID int64, limit, offset *int) ([]*model.Task, error) {
	l := 20
	o := 0
	if limit != nil {
		l = *limit
	}
	if offset != nil {
		o = *offset
	}

	if l <= 0 || l > 100 || o < 0 || userID <= 0 {
		return nil, ErrParaInvalid
	}
	tasks, err := s.repo.List(userID, l, o)
	if err != nil {
		return nil, fmt.Errorf("list tasks error :%w", err)
	}
	return tasks, nil
}

func (s *TaskService) GetByID(userID, id int64) (*model.Task, error) {
	if id <= 0 || userID <= 0 {
		return nil, ErrParaInvalid
	}
	task, err := s.repo.GetByID(userID, id)
	if err != nil {
		if errors.Is(err, repository.ErrTaskNotFound) {
			return nil, ErrTaskNotFound
		}
		return nil, fmt.Errorf("get task id error: %w", err)
	}
	return task, nil
}

func (s *TaskService) Update(userID, id int64, title, description, status string) (task *model.Task, err error) {
	if id <= 0 || userID <= 0 {
		return nil, ErrParaInvalid
	}
	titleTrim := strings.TrimSpace(title)
	if titleTrim == "" {
		return nil, ErrTitleRequired
	}
	task = &model.Task{
		ID:          id,
		UserID:      userID,
		Title:       titleTrim,
		Description: description,
		Status:      status,
	}
	err = s.repo.Update(task)
	if err != nil {
		if errors.Is(err, repository.ErrTaskNotFound) {
			return nil, ErrTaskNotFound
		}
		return nil, fmt.Errorf("update task error: %w", err)
	}
	return task, nil
}

func (s *TaskService) Delete(userID, id int64) error {
	if id <= 0 || userID <= 0 {
		return ErrParaInvalid
	}
	err := s.repo.Delete(userID, id)
	if err != nil {
		if errors.Is(err, repository.ErrTaskNotFound) {
			return ErrTaskNotFound
		}
		return fmt.Errorf("delete task error: %w", err)
	}
	return nil
}
