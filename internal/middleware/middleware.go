package middleware

import (
	"log/slog"
	"net/http"
	"strings"
	"time"

	"taskapi/internal/auth"
	"taskapi/internal/handler"
	"taskapi/internal/helper"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenManager *auth.TokenManager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1. 读取 Authorization header Authorization: Bearer eyJhbGciOi...
		authHeader := ctx.GetHeader("Authorization")

		// 2. 检查 Bearer 格式
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, handler.FailResp(map[string]any{"message": "not login"}))
			ctx.Abort()
			return
		}
		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, handler.FailResp(map[string]any{"message": "not login"}))
			ctx.Abort()
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		token = strings.TrimSpace(token)

		if token == "" {
			ctx.JSON(http.StatusUnauthorized, handler.FailResp(map[string]any{"message": "not login"}))
			ctx.Abort()
			return
		}

		// 3. 调用 tokenManager.ParseToken(token)
		claims, err := tokenManager.ParseToken(token)
		// 4. 失败就返回 401 并 Abort
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, handler.FailResp(map[string]any{"message": "not login"}))
			ctx.Abort()
			return
		}

		// 5. 成功就 ctx.Set("current_user_id", claims.UserID)
		ctx.Set("current_user_id", claims.UserID)
		// 6. ctx.Next()
		ctx.Next()
	}
}

func RequestLogger(log *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		duration := time.Since(start)

		log.Info("http request",
			"method", ctx.Request.Method,
			"path", ctx.Request.URL.Path,
			"request_id", helper.GetRequestID(ctx),
			"status", ctx.Writer.Status(),
			"duration_ms", duration.Milliseconds(),
			"client_ip", ctx.ClientIP())
	}
}

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := helper.GetRequestID(ctx)
		ctx.Set("request_id", requestID)
		ctx.Header("X-Request-ID", requestID)
		ctx.Next()
	}
}

func PanicRecovery(log *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if recovered := recover(); recovered != nil {
				log.Error("panic recovered",
					"request_id", helper.GetRequestID(ctx),
					"method", ctx.Request.Method,
					"path", ctx.Request.URL.Path,
					"panic", recovered)
				ctx.JSON(http.StatusInternalServerError, handler.FailResp(map[string]any{
					"message": "failed to execute",
				}))
				ctx.Abort()
			}
		}()

		ctx.Next()
	}
}
