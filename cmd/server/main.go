package main

import (
	"fmt"
	"log"
	"taskapi/internal/config"
	"taskapi/internal/router"
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

	r := router.SetupRouter()
	addr := ":" + cfg.Server.Port
	if err := r.Run(addr); err != nil {
		return fmt.Errorf("start server on %s: %w", addr, err)
	}
	return nil
}
