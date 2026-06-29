package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetRequestID(ctx *gin.Context) string {
	if value, exist := ctx.Get("request_id"); exist {
		if requestID, ok := value.(string); ok {
			return requestID
		}
	}

	requestID := ctx.GetHeader("X-Request-ID")
	if requestID == "" || uuid.Validate(requestID) != nil {
		requestID = uuid.NewString()
	}
	return requestID
}
