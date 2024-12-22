package routes

import (
	"backend/internal/delivery/http"
	"backend/internal/delivery/middlewares"
	"backend/pkg/constants"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	routes := router.Group(constants.API_PREFIX)
	{
		routes.GET("/heatbeat", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "beat"})
		})

		// AUTH
		authRoutes := routes.Group("/auth")
		{
			authRoutes.POST("/register", http.RegisterUser)
			authRoutes.GET("/confirm-register/:token", http.ActivateUserRegister)
			authRoutes.POST("/resend-confirmation", http.ResendUserRegisterConfirmation)
			authRoutes.POST("/login", http.LoginUser)
		}

		routes.Use(middlewares.AuthMiddleware())

		// USERS
		userRoutes := routes.Group("/users")
		{
			userRoutes.GET("/profile", http.UserProfile)
			userRoutes.PUT("/profile", http.UpdateUserProfile)
			userRoutes.POST("/password", http.UpdateUserPassword)
		}
	}
}
