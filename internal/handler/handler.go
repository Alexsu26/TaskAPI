package handler

import (
	"errors"
	"net/http"
	"strconv"

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

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
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

func parseListPara(ctx *gin.Context) (limit *int, offset *int, err error) {
	limitStr, offsetStr := ctx.Query("limit"), ctx.Query("offset")
	if limitStr != "" {
		parseLimit, err := strconv.Atoi(limitStr)
		if err != nil {
			return nil, nil, err
		}
		limit = &parseLimit
	}
	if offsetStr != "" {
		parseOffset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return nil, nil, err
		}
		offset = &parseOffset
	}
	return limit, offset, nil
}

func (h *Handler) RegisterTasksListRoutes(r *gin.Engine) {
	r.GET("/tasks", func(ctx *gin.Context) {
		limit, offset, err := parseListPara(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid request parameters",
			})
			return
		}
		tasks, err := h.taskService.List(limit, offset)
		if err != nil {
			if errors.Is(err, service.ErrParaInvalid) {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to list tasks",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"tasks":  tasks,
		})
	})
}

func (h *Handler) RegisterGetTaskRoutes(r *gin.Engine) {
	r.GET("/tasks/:id", func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "request parameter id is invalid",
			})
			return
		}
		task, err := h.taskService.GetByID(id)
		if err != nil {
			if errors.Is(err, service.ErrTaskNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			} else if errors.Is(err, service.ErrParaInvalid) {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "invalid parameter",
				})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to get task",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"task":   task,
		})
	})
}

func (h *Handler) RegisterUpdateTaskRoutes(r *gin.Engine) {
	r.PUT("/tasks/:id", func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "request parameter id is invalid",
			})
			return
		}
		var req UpdateTaskRequest
		if err = ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid request body",
			})
			return
		}
		task, err := h.taskService.Update(id, req.Title, req.Description, req.Status)
		if err != nil {
			if errors.Is(err, service.ErrTaskNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": "task not found",
				})
				return
			} else if errors.Is(err, service.ErrParaInvalid) {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "invalid parameter",
				})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "update task failed",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"task":   task,
		})
	})
}

func (h *Handler) RegisterDeleteTaskRoutes(r *gin.Engine) {
	r.DELETE("/tasks/:id", func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid task id",
			})
			return
		}
		err = h.taskService.Delete(id)
		if err != nil {
			if errors.Is(err, service.ErrTaskNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": "task not found",
				})
				return
			} else if errors.Is(err, service.ErrParaInvalid) {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "invalid parameter",
				})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to delete task",
			})
			return
		}
		ctx.Status(http.StatusNoContent)
	})
}
