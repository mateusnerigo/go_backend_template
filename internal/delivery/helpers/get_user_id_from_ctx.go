package delivery_helpers

import (
	"backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

func GetUserIdFromContext(ctx *gin.Context) uint64 {
	userId, exists := ctx.Get("user_id")
	logger := utils.NewLogger("")

	if !exists {
		logger.Error("user_id not found in token")
		return 0
	}

	return uint64(userId.(float64))

}
