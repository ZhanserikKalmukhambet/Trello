package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const authUserID = "userID"

func (h *Handler) authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			NewErrorResponse(ctx, http.StatusUnauthorized, "Authorization header is not set")
			return
		}

		parts := strings.Fields(authHeader)
		if len(parts) != 2 {
			NewErrorResponse(ctx, http.StatusUnauthorized, "Authorization header is incorrect")
			return
		}

		userID, err := h.srvs.VerifyToken(parts[1])
		if err != nil {
			NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		ctx.Set(authUserID, userID)
		ctx.Next()
	}
}

func getAuthUserID(ctx *gin.Context) (int, error) {
	userID, ok := ctx.MustGet(authUserID).(int)
	if !ok {
		NewErrorResponse(ctx, http.StatusUnauthorized, "Can't get user id")
		return 0, fmt.Errorf("no such user id")
	}

	return userID, nil
}
