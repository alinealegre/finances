package models

import (
	"time"

	"gorm.io/gorm"
)

type Debts struct {
	Id          int            `json:"id" gorm:"primaryKey"`
	Company     string         `json:"company"`
	Value       float64        `json:"value"`
	ExpiratedAt time.Time      `json:"expirated_at"`
	CPF         string         `json:"cpf"`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted"`
}
