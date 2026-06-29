package handler

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"taskapi/internal/helper"
	"taskapi/internal/model"
	"taskapi/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	taskService *service.TaskService
	userService *service.UserService
	log         *slog.Logger
}

func NewHandler(taskService *service.TaskService, userService *service.UserService, log *slog.Logger) *Handler {
	return &Handler{taskService: taskService, userService: userService, log: log}
}

func (h *Handler) handlerServiceError(ctx *gin.Context, err error) {
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
	if errors.Is(err, service.ErrEmailAlreadyExists) {
		ctx.JSON(http.StatusConflict, FailResp(map[string]any{"message": "email already exists"}))
		return
	}
	if errors.Is(err, service.ErrParaMiss) {
		ctx.JSON(http.StatusBadRequest, FailResp(map[string]any{"message": "missing parameter"}))
		return
	}
	if errors.Is(err, service.ErrInvalidCredentials) {
		ctx.JSON(http.StatusUnauthorized, FailResp(map[string]any{"message": "invalid email or password"}))
		return
	}
	if errors.Is(err, service.ErrPasswordInvalid) {
		ctx.JSON(http.StatusUnauthorized, FailResp(map[string]any{"message": "invalid email or password"}))
		return
	}
	if errors.Is(err, service.ErrTokenInvalid) {
		ctx.JSON(http.StatusBadRequest, FailResp(map[string]any{"message": "auth failed"}))
		return
	}
	h.log.Error("internal server error",
		"method", ctx.Request.Method,
		"path", ctx.Request.URL.Path,
		"request_id", helper.GetRequestID(ctx),
		"error", err.Error())
	ctx.JSON(http.StatusInternalServerError, FailResp(map[string]any{"message": "failed to execute"}))
}

func handlerCommonError(ctx *gin.Context, errCode int, message string) {
	ctx.JSON(errCode, FailResp(map[string]any{"message": message}))
}

func handlerSuccessResp(ctx *gin.Context, statusCode int, data map[string]any) {
	ctx.JSON(statusCode, SuccessResp(data))
}

func toTaskResponse(task *model.Task) TaskResponse {
	return TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

func toTasksResponse(tasks []*model.Task) []TaskResponse {
	resp := make([]TaskResponse, 0, len(tasks))
	for _, task := range tasks {
		resp = append(resp, toTaskResponse(task))
	}
	return resp
}

func toUserResponse(user *model.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func (h *Handler) RegisterHealthRoutes(r *gin.Engine) {
	r.GET("/health", func(ctx *gin.Context) {
		handlerSuccessResp(ctx, http.StatusOK, nil)
	})
}

func (h *Handler) RegisterTaskCreateRoutes(r RouteRegister) {
	r.POST("/tasks", func(ctx *gin.Context) {
		var req CreateTaskRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			handlerCommonError(ctx, http.StatusBadRequest, "invalid request body")
			return
		}
		userID, err := getContextUserID(ctx)
		if err != nil {
			handlerCommonError(ctx, http.StatusUnauthorized, "not login")
			return
		}
		task, err := h.taskService.Create(userID, req.Title, req.Description)
		if err != nil {
			h.handlerServiceError(ctx, err)
			return

		}
		taskResp := toTaskResponse(task)
		handlerSuccessResp(ctx, http.StatusCreated, map[string]any{"task": taskResp})
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

func (h *Handler) RegisterTasksListRoutes(r RouteRegister) {
	r.GET("/tasks", func(ctx *gin.Context) {
		limit, offset, err := parseListPara(ctx)
		if err != nil {
			handlerCommonError(ctx, http.StatusBadRequest, "invalid request parameters")
			return
		}
		userID, err := getContextUserID(ctx)
		if err != nil {
			handlerCommonError(ctx, http.StatusUnauthorized, "not login")
			return
		}
		tasks, err := h.taskService.List(userID, limit, offset)
		if err != nil {
			h.handlerServiceError(ctx, err)
			return
		}
		tasksResp := toTasksResponse(tasks)
		handlerSuccessResp(ctx, http.StatusOK, map[string]any{"tasks": tasksResp})
	})
}

func (h *Handler) RegisterGetTaskRoutes(r RouteRegister) {
	r.GET("/tasks/:id", func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			handlerCommonError(ctx, http.StatusBadRequest, "request parameter id is invalid")
			return
		}
		userID, err := getContextUserID(ctx)
		if err != nil {
			handlerCommonError(ctx, http.StatusUnauthorized, "not login")
			return
		}
		task, err := h.taskService.GetByID(userID, id)
		if err != nil {
			h.handlerServiceError(ctx, err)
			return
		}
		taskResp := toTaskResponse(task)
		handlerSuccessResp(ctx, http.StatusOK, map[string]any{"task": taskResp})
	})
}

func (h *Handler) RegisterUpdateTaskRoutes(r RouteRegister) {
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
		userID, err := getContextUserID(ctx)
		if err != nil {
			handlerCommonError(ctx, http.StatusUnauthorized, "not login")
			return
		}
		task, err := h.taskService.Update(userID, id, req.Title, req.Description, req.Status)
		if err != nil {
			h.handlerServiceError(ctx, err)
			return
		}
		taskResp := toTaskResponse(task)
		handlerSuccessResp(ctx, http.StatusOK, map[string]any{"task": taskResp})
	})
}

func (h *Handler) RegisterDeleteTaskRoutes(r RouteRegister) {
	r.DELETE("/tasks/:id", func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			handlerCommonError(ctx, http.StatusBadRequest, "invalid task id")
			return
		}
		userID, err := getContextUserID(ctx)
		if err != nil {
			handlerCommonError(ctx, http.StatusUnauthorized, "not login")
			return
		}
		err = h.taskService.Delete(userID, id)
		if err != nil {
			h.handlerServiceError(ctx, err)
			return
		}
		handlerSuccessResp(ctx, http.StatusOK, nil)
	})
}

func (h *Handler) RegisterCreateUserRoutes(r *gin.Engine) {
	r.POST("/users/register", func(ctx *gin.Context) {
		var req CreateUserRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			handlerCommonError(ctx, http.StatusBadRequest, "invalid request body")
			return
		}
		user, err := h.userService.Create(req.Name, req.Email, req.Password)
		if err != nil {
			h.handlerServiceError(ctx, err)
			return
		}

		resp := toUserResponse(user)
		handlerSuccessResp(ctx, http.StatusCreated, map[string]any{"user": resp})
	})
}

func (h *Handler) RegisterUserLoginRoutes(r *gin.Engine) {
	r.POST("/users/login", func(ctx *gin.Context) {
		var req UserLoginRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			handlerCommonError(ctx, http.StatusBadRequest, "invalid request body")
			return
		}
		user, token, err := h.userService.Login(req.Email, req.Password)
		if err != nil {
			h.handlerServiceError(ctx, err)
			return
		}
		resp := &UserLoginResp{
			User:  toUserResponse(user),
			Token: token,
		}
		handlerSuccessResp(ctx, http.StatusOK, map[string]any{"user": resp.User, "token": resp.Token})
	})
}
