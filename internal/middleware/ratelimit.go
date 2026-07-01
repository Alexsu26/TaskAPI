package middleware

import (
	"net/http"
	"time"

	"taskapi/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var (
	RatelimitPrefixKey string = "rate_limit:login:"
	limit              int64  = 5
)

func RateLimit(redisClient *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := RatelimitPrefixKey + ctx.ClientIP()

		count, err := redisClient.Incr(ctx, key).Result()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, handler.FailResp(map[string]any{"message": "redis set key error"}))
			ctx.Abort()
			return
		}

		if count == 1 {
			if err = redisClient.Expire(ctx, key, time.Minute).Err(); err != nil {
				ctx.JSON(http.StatusInternalServerError, handler.FailResp(map[string]any{"message": "redis expire key error"}))
				ctx.Abort()
				return
			}
		}

		if count > limit {
			ctx.JSON(http.StatusTooManyRequests, handler.FailResp(map[string]any{"message": "please try again later"}))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
