package middlewares

import (
	"backend/internal/application/security"
	"backend/pkg/constants"
	"backend/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	logger := utils.NewLogger("")

	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			logger.Error("authorization is empty")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": constants.INVALID_TOKEN})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			logger.Error("wrong token format")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": constants.INVALID_TOKEN})
			ctx.Abort()
			return
		}

		token, err := security.VerifyJWT(tokenString)

		if err != nil || !token.Valid {
			logger.Error("invalid token sent")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": constants.INVALID_USER_TOKEN})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx.Set("user_id", claims["sub"])
		}

		ctx.Next()
	}
}
