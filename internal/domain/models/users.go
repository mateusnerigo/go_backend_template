package models

type User struct {
	ID                uint64  `json:"id" gorm:"primaryKey,autoIncrement"`
	UUID              string  `json:"uuid" gorm:"type:uuid;default:gen_random_uuid()"`
	FirstName         string  `json:"first_name"`
	LastName          string  `json:"last_name"`
	Email             string  `json:"email" gorm:"uniqueIndex"`
	Phone             string  `json:"phone" gorm:"uniqueIndex"`
	Password          string  `json:"password"`
	PasswordResetHash *string `json:"password_reset_hash"`
	Verified          bool    `json:"verified"`
	VerificationToken string  `json:"verification_token"`
	DarkMode          bool    `json:"dark_mode"`
	TimestampedRegister
}
