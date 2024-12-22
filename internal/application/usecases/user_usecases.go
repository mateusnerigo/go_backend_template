package usecases

import (
	"backend/internal/application/security"
	"backend/internal/domain/dtos"
	"backend/internal/domain/validations/schemas"
	"backend/internal/infra/repositories"
	"backend/pkg/constants"
	"errors"
)

func UserProfileUseCase(userId *uint64) (*dtos.UserProfileDTO, error) {
	user := repositories.FindById(userId)

	if user.ID <= 0 {
		return nil, errors.New(constants.USER_NOT_FOUND)
	}

	profileData := &dtos.UserProfileDTO{
		ID:        user.ID,
		UUID:      user.UUID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		DarkMode:  user.DarkMode,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return profileData, nil
}

func UpdateUserDataUseCase(userId *uint64, updateData *schemas.UserUpdateSchema) error {
	user := repositories.FindByEmail(&updateData.Email, userId)
	if user.ID > 0 {
		return errors.New(constants.EMAIL_ALREADY_IN_USE)
	}

	user = repositories.FindByPhone(&updateData.Phone, userId)
	if user.ID > 0 {
		return errors.New(constants.PHONE_ALREADY_IN_USE)
	}

	updateData.ID = *userId
	status := repositories.UpdateUserData(updateData)

	if !status {
		return errors.New(constants.ERROR_UPDATING_USER)
	}

	return nil
}

func UpdateUserPasswordUseCase(userId *uint64, updateData *schemas.UserPasswordUpdateSchema) error {
	updateData.ID = *userId
	updateData.Password = security.HashPassword(updateData.Password)
	status := repositories.UpdateUserPassword(updateData)

	if !status {
		return errors.New(constants.ERROR_UPDATING_USER_PASSWORD)
	}

	return nil
}
