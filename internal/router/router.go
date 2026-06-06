package router

import (
	"taskapi/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	handler.RegisterHealthRoutes(r)

	return r
}
