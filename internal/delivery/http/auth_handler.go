package http

import (
	"backend/internal/application/security"
	"backend/internal/application/usecases"
	"backend/internal/domain/validations/schemas"
	"backend/pkg/constants"
	"backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RegisterUser(ctx *gin.Context) {
	var userRegisterData schemas.UserRegisterSchema

	if err := ctx.ShouldBind(&userRegisterData); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": utils.ValidationMsgHandler(err, userRegisterData)},
		)
		return
	}

	if userRegisterData.Password != userRegisterData.PasswordConfirmation {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": constants.PASSWORDS_DONT_MATCH},
		)
		return
	}

	err := usecases.RegisterUserUseCase(&userRegisterData)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": constants.USER_CREATED})
}

func ActivateUserRegister(ctx *gin.Context) {
	logger := utils.NewLogger("")
	confirmationToken := ctx.Param("token")

	if confirmationToken == "" {
		logger.Error("confirmation token is empty")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": constants.INVALID_CONFIRMATION_TOKEN})
		return
	}

	token, err := security.VerifyJWT(confirmationToken)

	if err != nil || !token.Valid {
		logger.Error("invalid confirmation token sent")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": constants.INVALID_CONFIRMATION_TOKEN})
		return
	}

	userEmail := ""
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userEmail = claims["sub"].(string)
	}

	if userEmail == "" {
		logger.Error("user email not found in confirmation token")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": constants.INVALID_CONFIRMATION_TOKEN})
		return
	}

	err = usecases.ConfirmUserRegistrationUseCase(&userEmail)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": constants.REGISTRATION_CONFIRMED})
}

func ResendUserRegisterConfirmation(ctx *gin.Context) {
	var resendData schemas.UserResendVerificationSchema

	if err := ctx.ShouldBind(&resendData); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": utils.ValidationMsgHandler(err, resendData)},
		)
		return
	}

	err := usecases.ResendUserRegisterConfirmationUseCase(&resendData)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": constants.USER_CONFIRMATION_RESENT})
}

func LoginUser(ctx *gin.Context) {
	var userLoginData schemas.UserLoginSchema

	if err := ctx.ShouldBind(&userLoginData); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": utils.ValidationMsgHandler(err, userLoginData)},
		)
		return
	}

	token, err := usecases.LoginUseCase(&userLoginData)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
