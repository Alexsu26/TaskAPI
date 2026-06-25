package middleware

import (
	"net/http"
	"strings"

	"taskapi/internal/auth"
	"taskapi/internal/handler"

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
