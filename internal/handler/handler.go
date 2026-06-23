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

func handlerServiceError(ctx *gin.Context, err error) {
	if errors.Is(err, service.ErrParaInvalid) {
		ctx.JSON(http.StatusBadRequest, FailResp(map[string]any{"message": "invalid parameter"}))
		return
	}
	if errors.Is(err, service.ErrTitleRequired) {
		ctx.JSON(http.StatusBadRequest, FailResp(map[string]any{"message": "missing task title"}))
		return
	}
	if errors.Is(err, service.ErrTaskNotFound) {
		ctx.JSON(http.StatusNotFound, FailResp(map[string]any{"message": "task not found"}))
		return
	}
	ctx.JSON(http.StatusInternalServerError, FailResp(map[string]any{"message": "failed to execute"}))
}

func handlerCommonError(ctx *gin.Context, errCode int, message string) {
	ctx.JSON(errCode, FailResp(map[string]any{"message": message}))
}

func handlerSuccessResp(ctx *gin.Context, statusCode int, data map[string]any) {
	ctx.JSON(statusCode, SuccessResp(data))
}

func (h *Handler) RegisterHealthRoutes(r *gin.Engine) {
	r.GET("/health", func(ctx *gin.Context) {
		handlerSuccessResp(ctx, http.StatusOK, nil)
	})
}

func (h *Handler) RegisterTaskCreateRoutes(r *gin.Engine) {
	r.POST("/tasks", func(ctx *gin.Context) {
		var req CreateTaskRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			handlerCommonError(ctx, http.StatusBadRequest, "invalid request body")
			return
		}
		task, err := h.taskService.Create(req.Title, req.Description)
		if err != nil {
			handlerServiceError(ctx, err)
			return

		}
		handlerSuccessResp(ctx, http.StatusCreated, map[string]any{"task": task})
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
			handlerCommonError(ctx, http.StatusBadRequest, "invalid request parameters")
			return
		}
		tasks, err := h.taskService.List(limit, offset)
		if err != nil {
			handlerServiceError(ctx, err)
			return
		}
		handlerSuccessResp(ctx, http.StatusOK, map[string]any{"tasks": tasks})
	})
}

func (h *Handler) RegisterGetTaskRoutes(r *gin.Engine) {
	r.GET("/tasks/:id", func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			handlerCommonError(ctx, http.StatusBadRequest, "request parameter id is invalid")
			return
		}
		task, err := h.taskService.GetByID(id)
		if err != nil {
			handlerServiceError(ctx, err)
			return
		}
		handlerSuccessResp(ctx, http.StatusOK, map[string]any{"task": task})
	})
}

func (h *Handler) RegisterUpdateTaskRoutes(r *gin.Engine) {
	r.PUT("/tasks/:id", func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			handlerCommonError(ctx, http.StatusBadRequest, "request parameter id is invalid")
			return
		}
		var req UpdateTaskRequest
		if err = ctx.ShouldBindJSON(&req); err != nil {
			handlerCommonError(ctx, http.StatusBadRequest, "invalid request body")
			return
		}
		task, err := h.taskService.Update(id, req.Title, req.Description, req.Status)
		if err != nil {
			handlerServiceError(ctx, err)
			return
		}
		handlerSuccessResp(ctx, http.StatusOK, map[string]any{"task": task})
	})
}

func (h *Handler) RegisterDeleteTaskRoutes(r *gin.Engine) {
	r.DELETE("/tasks/:id", func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			handlerCommonError(ctx, http.StatusBadRequest, "invalid task id")
			return
		}
		err = h.taskService.Delete(id)
		if err != nil {
			handlerServiceError(ctx, err)
			return
		}
		handlerSuccessResp(ctx, http.StatusOK, nil)
	})
}
