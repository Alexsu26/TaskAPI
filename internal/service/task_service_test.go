package service

import (
	"errors"
	"testing"

	"taskapi/internal/model"
)

type fakeTaskRepo struct {
	createdTask *model.Task
	createErr   error
}

func (f *fakeTaskRepo) Create(task *model.Task) error {
	f.createdTask = task
	return f.createErr
}

func (f *fakeTaskRepo) List(userID int64, limit, offset int) ([]*model.Task, error) {
	panic("not implement")
}

func (f *fakeTaskRepo) GetByID(userID, id int64) (task *model.Task, err error) {
	panic("not implement")
}

func (f *fakeTaskRepo) Update(task *model.Task) error {
	panic("not implement")
}

func (f *fakeTaskRepo) Delete(userID, id int64) error {
	panic("not implement")
}

func TestTaskService_EmptyTitle(t *testing.T) {
	taskService := NewTaskService(&fakeTaskRepo{})
	_, err := taskService.Create(11111, "", "")
	if !errors.Is(err, ErrTitleRequired) {
		t.Fatalf("expected ErrTitleRequired, got %v", err)
	}
}

func TestTaskService_InvalidUserID(t *testing.T) {
	taskService := NewTaskService(&fakeTaskRepo{})
	_, err := taskService.Create(0, "test", "")
	if !errors.Is(err, ErrParaInvalid) {
		t.Fatalf("expected ErrParaInvalid, got %v", err)
	}
}

func TestTaskService_SuccessCreate(t *testing.T) {
	fakeRepo := &fakeTaskRepo{}
	taskService := NewTaskService(fakeRepo)

	task, err := taskService.Create(1111, "   test   ", "test")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if fakeRepo.createdTask == nil {
		t.Fatalf("expected repo Create to receive task, got nil")
	}

	if fakeRepo.createdTask.UserID != 1111 {
		t.Fatalf("expected UserID 1111, got %d", fakeRepo.createdTask.UserID)
	}

	if fakeRepo.createdTask.Title != "test" {
		t.Fatalf("expected Title is 'test', got %q", fakeRepo.createdTask.Title)
	}

	if fakeRepo.createdTask.Status != "todo" {
		t.Fatalf("expected Status is 'todo', got %q", fakeRepo.createdTask.Status)
	}

	if fakeRepo.createdTask.Description != "test" {
		t.Fatalf("expected Description is 'test', got %q", fakeRepo.createdTask.Description)
	}

	if task != fakeRepo.createdTask {
		t.Fatalf("expected service to return the same task to repo")
	}
}
