package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(r *gin.Engine) {
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
}
