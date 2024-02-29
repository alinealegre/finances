package services

import (
	"finances/models"
	"math"

	"gorm.io/gorm"
)

func CalculateUserScore(db *gorm.DB) int {
	var debts []models.Debts
	db.Find(&debts)

	var totalValue float64
	totalValue = 0
	for _, debt := range debts {
		totalValue += debt.Value
	}

	var averageValue float64
	averageValue = 0
	if totalValue != 0 && len(debts) > 0 {
		averageValue = float64((int(totalValue) / len(debts)))
	}

	score := (10000 / math.Sqrt(averageValue+100))

	return int(score)
}
