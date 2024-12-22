package usecases

import (
	"backend/internal/application/security"
	"backend/internal/domain/models"
	"backend/internal/domain/validations/schemas"
	"backend/internal/infra/notifications"
	"backend/internal/infra/repositories"
	"backend/pkg/constants"
	"backend/pkg/utils"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func generateEmailVerificationToken(email string) (string, error) {
	emailVerificationTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,
		"exp": time.Now().Add((time.Hour * 24) * 7).Unix(),
	})

	return emailVerificationTokenClaims.SignedString([]byte(security.GetJWTSecret()))
}

func sendVerificationEmail(email string, verificationToken string) error {
	sender := notifications.NewSMTPEmailSender()

	return sender.Send(
		email,
		"Register verification",
		fmt.Sprintf(
			"Finish your register by clicking the link below \r\n\r\nhttp://localhost.com:8000/api/v1/confirm-register/%v",
			verificationToken,
		),
	)
}

func RegisterUserUseCase(newUserData *schemas.UserRegisterSchema) error {
	logger := utils.NewLogger("")

	user := repositories.FindByEmail(&newUserData.Email, nil)
	if user.ID > 0 {
		return errors.New(constants.EMAIL_ALREADY_IN_USE)
	}

	user = repositories.FindByPhone(&newUserData.Phone, nil)
	if user.ID > 0 {
		return errors.New(constants.PHONE_ALREADY_IN_USE)
	}

	token, err := generateEmailVerificationToken(newUserData.Email)
	if err != nil {
		logger.Error("error generating confirmation token")
		return errors.New(constants.ERROR_CREATING_USER)
	}

	status := repositories.CreateUser(&models.User{
		FirstName:         newUserData.FirstName,
		LastName:          newUserData.LastName,
		Phone:             newUserData.Phone,
		Email:             newUserData.Email,
		Password:          security.HashPassword(newUserData.Password),
		VerificationToken: token,
	})

	if !status {
		logger.Error("error creating user")
		return errors.New(constants.ERROR_CREATING_USER)
	}

	emailError := sendVerificationEmail(newUserData.Email, token)
	if emailError != nil {
		logger.Error(fmt.Sprintf("error sending verification email: %v", emailError.Error()))
		return errors.New(constants.ERROR_SENDING_VERIFICATION_EMAIL)
	}

	return nil
}

func ConfirmUserRegistrationUseCase(email *string) error {
	user := repositories.FindByEmail(email, nil)
	if user.ID == 0 {
		return errors.New(constants.USER_NOT_FOUND)
	}

	if user.Verified {
		return nil
	}

	updateStatus := repositories.UpdateVerifiedUserStatus(&user.ID)
	if !updateStatus {
		return errors.New(constants.ERROR_UPDATING_USER_VERIFICATION)
	}

	return nil
}

func ResendUserRegisterConfirmationUseCase(resendData *schemas.UserResendVerificationSchema) error {
	logger := utils.NewLogger("")

	// search by email given
	user := repositories.FindByEmail(&resendData.Email, nil)

	// search by phone given if it founds no user and if phone has sent
	if (user.ID == 0) && (resendData.Phone != "") {
		user = repositories.FindByPhone(&resendData.Phone, nil)
	}

	if user.ID == 0 {
		logger.Warn("user not found by email or phone")
		return errors.New(constants.USER_NOT_FOUND)
	}

	if user.Verified {
		logger.Warn("user already verified")
		return errors.New(constants.USER_ALREADY_VERIFIED)
	}

	token, err := security.VerifyJWT(user.VerificationToken)

	if err != nil || !token.Valid {
		logger.Error("invalid confirmation token required to resend")

		token, err := generateEmailVerificationToken(user.Email)
		if err != nil {
			logger.Error("error generating new confirmation token to resend")
			return errors.New(constants.ERROR_GENERATING_NEW_TOKEN)
		}

		updateStatus := repositories.UpdateUserVerificationToken(&user.ID, &token)
		if !updateStatus {
			logger.Error("error updating new confirmation token to resend")
			return errors.New(constants.ERROR_UPDATING_USER_VERIFICATION)
		}

		// updates the token to be send
		user.VerificationToken = token
	}

	emailError := sendVerificationEmail(user.Email, user.VerificationToken)
	if emailError != nil {
		logger.Error(fmt.Sprintf("error resending verification email: %v", emailError.Error()))
		return errors.New(constants.ERROR_RESENDING_VERIFICATION_EMAIL)
	}

	return nil
}

func LoginUseCase(loginData *schemas.UserLoginSchema) (string, error) {
	logger := utils.NewLogger("")

	// search by email given
	user := repositories.FindByEmail(&loginData.Email, nil)

	// search by phone given if it founds no user and if phone has sent
	if (user.ID == 0) && (loginData.Phone != "") {
		user = repositories.FindByPhone(&loginData.Phone, nil)
	}

	// if no user found or if passwords doesn`t match
	if user.ID == 0 || !security.CheckPasswordHash(loginData.Password, user.Password) {
		logger.Warn("password doesnt check")
		return "", errors.New(constants.INVALID_CREDENTIALS)
	}

	if !user.Verified {
		logger.Warn("user not verified")
		return "", errors.New(constants.USER_NOT_VERIFIED)
	}

	logger.Log(fmt.Sprintf("user logged in. user_id: %v, type: %v", user.ID, reflect.TypeOf(user.ID)))

	// generates token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(security.GetJWTSecret()))

}
