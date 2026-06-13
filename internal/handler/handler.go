package handler

import (
	"errors"
	"net/http"

	"taskapi/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	taskService *service.TaskService
}

func NewHandler(taskService *service.TaskService) *Handler {
	return &Handler{taskService: taskService}
}

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

func (h *Handler) RegisterHealthRoutes(r *gin.Engine) {
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
}

func (h *Handler) RegisterTaskCreateRoutes(r *gin.Engine) {
	r.POST("/tasks", func(ctx *gin.Context) {
		var req CreateTaskRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid request body",
			})
			return
		}
		task, err := h.taskService.Create(req.Title, req.Description)
		if err != nil {
			if errors.Is(err, service.ErrTitleRequired) {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to create task",
			})
			return

		}
		ctx.JSON(http.StatusCreated, gin.H{
			"status": "ok",
			"task":   task,
		})
	})
}
