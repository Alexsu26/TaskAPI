package main

import (
	"fmt"
	"log"
	"time"

	"taskapi/internal/auth"
	"taskapi/internal/config"
	"taskapi/internal/database"
	"taskapi/internal/handler"
	"taskapi/internal/repository"
	"taskapi/internal/router"
	"taskapi/internal/service"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("server startup failed: %v", err)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	db, err := database.NewPostgresDB(&cfg)
	if err != nil {
		return fmt.Errorf("init postgres db: %w", err)
	}
	err = database.RunMigrations(db)
	if err != nil {
		return fmt.Errorf("create table error: %w", err)
	}
	taskRepo := repository.NewTaskRepo(db)
	taskService := service.NewTaskService(taskRepo)

	tokenManager := auth.NewTokenManager(
		cfg.Auth.JWTSecret,
		time.Duration(cfg.Auth.JWTExpirationMinutes)*time.Minute)
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo, tokenManager)
	handler := handler.NewHandler(taskService, userService)
	r := router.SetupRouter(handler, tokenManager)

	addr := ":" + cfg.Server.Port
	if err := r.Run(addr); err != nil {
		return fmt.Errorf("start server on %s: %w", addr, err)
	}
	return nil
}
