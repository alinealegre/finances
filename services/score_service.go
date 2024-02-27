package services

import (
	"finances/models"
	"math"

	"gorm.io/gorm"
)

func CalculateUserScore(db *gorm.DB) float64 {
	var debts []models.Debts
	db.Find(&debts)

	if len(debts) == 0 {
		return 0
	}

	var totalValue float64
	for _, debt := range debts {
		totalValue += float64(debt.Value)
	}

	averageValue := totalValue / float64(len(debts))
	score := 10000 / math.Sqrt(averageValue+100)

	return score
}
