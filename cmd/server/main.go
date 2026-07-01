package main

import (
	"fmt"
	"log"
	"time"

	"taskapi/internal/auth"
	"taskapi/internal/cache"
	"taskapi/internal/config"
	"taskapi/internal/database"
	"taskapi/internal/handler"
	"taskapi/internal/logger"
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
	log := logger.New()
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	// init db
	db, err := database.NewPostgresDB(&cfg)
	if err != nil {
		return fmt.Errorf("init postgres db: %w", err)
	}
	defer db.Close()

	// init redis
	redisClient, err := cache.NewRedisClient(&cfg)
	if err != nil {
		return fmt.Errorf("init redis failed: %w", err)
	}
	defer redisClient.Close()

	taskRepo := repository.NewTaskRepo(db)
	taskService := service.NewTaskService(taskRepo)

	tokenManager := auth.NewTokenManager(
		cfg.Auth.JWTSecret,
		time.Duration(cfg.Auth.JWTExpirationMinutes)*time.Minute)
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo, tokenManager)

	// handler
	handler := handler.NewHandler(taskService, userService, log)
	r := router.SetupRouter(handler, tokenManager, log, redisClient)

	addr := ":" + cfg.Server.Port
	log.Info("server starting",
		"addr", addr,
		"database_host", cfg.Database.Host,
		"database_port", cfg.Database.Port,
		"database_name", cfg.Database.Name,
		"jwt_expiration_minutes", cfg.Auth.JWTExpirationMinutes)
	if err := r.Run(addr); err != nil {
		return fmt.Errorf("start server on %s: %w", addr, err)
	}
	return nil
}
