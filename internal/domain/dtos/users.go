package dtos

import "time"

type UserProfileDTO struct {
	ID                uint64         `json:"id"`
	UUID              string         `json:"uuid"`
	FirstName         string         `json:"first_name"`
	LastName          string         `json:"last_name"`
	Email             string         `json:"email"`
	Phone             string         `json:"phone"`
	DarkMode          bool           `json:"dark_mode"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
}
