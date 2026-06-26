package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrUserIDInvalid = errors.New("user id not found")

type RouteRegister interface {
	GET(string, ...gin.HandlerFunc) gin.IRoutes
	POST(string, ...gin.HandlerFunc) gin.IRoutes
	PUT(string, ...gin.HandlerFunc) gin.IRoutes
	DELETE(string, ...gin.HandlerFunc) gin.IRoutes
}

func getContextUserID(ctx *gin.Context) (int64, error) {
	userIDValue, exists := ctx.Get("current_user_id")
	if !exists {
		return 0, ErrUserIDInvalid
	}

	userID, ok := userIDValue.(int64)
	if !ok || userID <= 0 {
		return 0, ErrUserIDInvalid
	}
	return userID, nil
}
