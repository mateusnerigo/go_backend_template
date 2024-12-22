package models

import (
	"time"

	"gorm.io/gorm"
)

type TimestampedRegister struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
