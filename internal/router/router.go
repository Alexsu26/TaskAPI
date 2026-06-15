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

	return r
}
