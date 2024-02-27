package migrations

import (
	"finances/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Debts{})
	db.AutoMigrate(models.User{})
}
