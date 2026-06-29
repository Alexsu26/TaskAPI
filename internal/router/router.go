package router

import (
	"log/slog"

	"taskapi/internal/auth"
	"taskapi/internal/handler"
	"taskapi/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.Handler, tokenManager *auth.TokenManager, log *slog.Logger) *gin.Engine {
	r := gin.New()
	r.Use(middleware.RequestID())
	r.Use(middleware.RequestLogger(log))
	r.Use(middleware.PanicRecovery(log))

	// 不需要auth
	h.RegisterHealthRoutes(r)
	h.RegisterCreateUserRoutes(r)
	h.RegisterUserLoginRoutes(r)

	// 需要auth
	taskGroup := r.Group("")
	taskGroup.Use(middleware.AuthMiddleware(tokenManager))

	h.RegisterTaskCreateRoutes(taskGroup)
	h.RegisterTasksListRoutes(taskGroup)
	h.RegisterGetTaskRoutes(taskGroup)
	h.RegisterUpdateTaskRoutes(taskGroup)
	h.RegisterDeleteTaskRoutes(taskGroup)
	return r
}
