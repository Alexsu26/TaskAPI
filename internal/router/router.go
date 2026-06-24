package router

import (
	"taskapi/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	h.RegisterHealthRoutes(r)
	h.RegisterTaskCreateRoutes(r)
	h.RegisterTasksListRoutes(r)
	h.RegisterGetTaskRoutes(r)
	h.RegisterUpdateTaskRoutes(r)
	h.RegisterDeleteTaskRoutes(r)
	h.RegisterCreateUserRoutes(r)

	return r
}
