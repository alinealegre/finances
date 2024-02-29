package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	CPF       string         `json:"cpf" gorm:"primaryKey"`
	Name      string         `json:"name"`
	BirthDate time.Time      `json:"birth_at"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Score     int            `json:"score"`
	Type      bool           `json:"type"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}
