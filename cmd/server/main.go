package main

import (
	"fmt"
	"log"

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
	taskHandler := handler.NewHandler(taskService)
	r := router.SetupRouter(taskHandler)

	addr := ":" + cfg.Server.Port
	if err := r.Run(addr); err != nil {
		return fmt.Errorf("start server on %s: %w", addr, err)
	}
	return nil
}
