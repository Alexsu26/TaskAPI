package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"taskapi/internal/model"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func openTestDB(t *testing.T) *sql.DB {
	t.Helper()

	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		t.Skip("TEST_DATABASE_URL is not set")
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		t.Fatalf("db init error: %v", err)
	}
	if err = db.Ping(); err != nil {
		t.Fatalf("db ping error: %v", err)
	}

	t.Cleanup(func() {
		db.Close()
	})
	return db
}

func createTestUser(t *testing.T, db *sql.DB) int64 {
	t.Helper()

	email := fmt.Sprintf("repo-test-%d@example.com", time.Now().UnixNano())

	var userID int64
	err := db.QueryRow(
		`insert into users (email, name, password_hash) values ($1, $2, $3) returning id`,
		email,
		"repo-test-user",
		"test-password").Scan(&userID)
	if err != nil {
		t.Fatalf("insert test users error: %v", err)
	}
	t.Cleanup(func() {
		if _, err := db.Exec(`delete from tasks where user_id = $1`, userID); err != nil {
			t.Errorf("delete test task error: %v", err)
		}
		if _, err := db.Exec(`delete from users where id = $1`, userID); err != nil {
			t.Errorf("delete test user error: %v", err)
		}
	})
	return userID
}

func TestTaskRepo_SuccessCreate(t *testing.T) {
	db := openTestDB(t)
	userID := createTestUser(t, db)
	taskRepo := NewTaskRepo(db)
	task := &model.Task{
		UserID:      userID,
		Title:       "test",
		Description: "test",
		Status:      "todo",
	}
	err := taskRepo.Create(task)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if task.ID == 0 {
		t.Fatalf("task should be back by DB")
	}
	createdTask, err := taskRepo.GetByID(userID, task.ID)
	if err != nil {
		t.Fatalf("expected got task, got %v", err)
	}
	if createdTask.UserID != userID || createdTask.Title != "test" || createdTask.Description != "test" || createdTask.Status != "todo" {
		t.Fatalf("geted task not match")
	}
}

func TestTaskRepo_NotFound(t *testing.T) {
	db := openTestDB(t)
	userID := createTestUser(t, db)
	taskRepo := NewTaskRepo(db)
	task, err := taskRepo.GetByID(userID, 5555)
	if !errors.Is(err, ErrTaskNotFound) {
		t.Fatalf("expected ErrTaskNotFound, got %v", err)
	}
	if task != nil {
		t.Fatalf("expected nil task, got %v", task)
	}
}
