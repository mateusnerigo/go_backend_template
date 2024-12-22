package http

import (
	"backend/internal/application/usecases"
	delivery_helpers "backend/internal/delivery/helpers"
	"backend/internal/domain/validations/schemas"
	"backend/pkg/constants"
	"backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserProfile(ctx *gin.Context) {
	logger := utils.NewLogger("")
	userId := delivery_helpers.GetUserIdFromContext(ctx)
	if userId == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": constants.USER_NOT_FOUND})
		return
	}

	userData, err := usecases.UserProfileUseCase(&userId)
	if err != nil {
		logger.Error("user not found")
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": constants.USER_NOT_FOUND},
		)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": userData})
}

func UpdateUserProfile(ctx *gin.Context) {
	userId := delivery_helpers.GetUserIdFromContext(ctx)
	if userId == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": constants.USER_NOT_FOUND})
		return
	}

	var userUpdateData schemas.UserUpdateSchema
	if err := ctx.ShouldBind(&userUpdateData); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": utils.ValidationMsgHandler(err, userUpdateData)},
		)
		return
	}

	err := usecases.UpdateUserDataUseCase(&userId, &userUpdateData)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": constants.USER_UPDATED})
}

func UpdateUserPassword(ctx *gin.Context) {
	userId := delivery_helpers.GetUserIdFromContext(ctx)
	if userId == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": constants.USER_NOT_FOUND})
		return
	}

	var userPasswordData schemas.UserPasswordUpdateSchema
	if err := ctx.ShouldBind(&userPasswordData); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": utils.ValidationMsgHandler(err, userPasswordData)},
		)
		return
	}

	if userPasswordData.Password != userPasswordData.PasswordConfirmation {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": constants.PASSWORDS_DONT_MATCH},
		)
		return
	}

	err := usecases.UpdateUserPasswordUseCase(&userId, &userPasswordData)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": constants.USER_PASSWORD_UPDATED})
}
