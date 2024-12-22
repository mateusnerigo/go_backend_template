package repositories

import (
	"backend/internal/domain/models"
	"backend/internal/domain/validations/schemas"
	"backend/internal/infra/database"
)

func CreateUser(userData *models.User) bool {
	db, _ := database.Client()

	result := db.Create(&userData)
	return result.Error == nil
}

func FindById(userId *uint64) models.User {
	db, _ := database.Client()

	user := models.User{}
	db.Where("id = ?", &userId).Limit(1).Find(&user)
	return user
}

func FindByEmail(email *string, userId *uint64) models.User {
	db, _ := database.Client()

	user := models.User{}

	query := db.Where("email = ?", *email)

	if userId != nil {
		query.Where("id != ?", *userId)
	}

	query.Limit(1).Find(&user)
	return user
}

func FindByPhone(phone *string, userId *uint64) models.User {
	db, _ := database.Client()

	user := models.User{}

	query := db.Where("phone = ?", *phone)

	if userId != nil {
		query.Where("id != ?", *userId)
	}

	query.Limit(1).Find(&user)
	return user
}

func UpdateVerifiedUserStatus(userId *uint64) bool {
	db, _ := database.Client()

	result := db.Model(&models.User{ID: *userId}).Updates(&models.User{
		Verified:          true,
		VerificationToken: "",
	})
	return result.Error == nil
}

func UpdateUserData(userData *schemas.UserUpdateSchema) bool {
	db, _ := database.Client()

	result := db.Model(&models.User{ID: userData.ID}).Updates(*userData)
	return result.Error == nil
}

func UpdateUserPassword(userData *schemas.UserPasswordUpdateSchema) bool {
	db, _ := database.Client()

	result := db.Model(&models.User{ID: userData.ID}).Updates(*userData)
	return result.Error == nil
}

func UpdateUserVerificationToken(userId *uint64, verificationToken *string) bool {
	db, _ := database.Client()

	result := db.Model(&models.User{ID: *userId}).Updates(&models.User{
		VerificationToken: *verificationToken,
	})
	return result.Error == nil
}
