package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	"taskapi/internal/worker"
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

	// worker
	taskWorker := worker.New(log, 100)
	taskWorker.Start()

	// handler
	handler := handler.NewHandler(taskService, userService, log, taskWorker)
	r := router.SetupRouter(handler, tokenManager, log, redisClient)

	addr := ":" + cfg.Server.Port
	log.Info("server starting",
		"addr", addr,
		"database_host", cfg.Database.Host,
		"database_port", cfg.Database.Port,
		"database_name", cfg.Database.Name,
		"jwt_expiration_minutes", cfg.Auth.JWTExpirationMinutes)

	httpServer := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	serverErr := make(chan error, 1)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
	}()

	// 等待 os.Interrupt / syscall.SIGTERM
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(shutdownSignal)

	select {
	case err := <-serverErr:
		return fmt.Errorf("server error: %w", err)
	case sig := <-shutdownSignal:
		log.Info("shutdown signal received", "signal", sig.String())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown server: %w", err)
	}

	taskWorker.Stop()
	log.Info("worker stopped")
	log.Info("server shutdown")
	return nil
}
