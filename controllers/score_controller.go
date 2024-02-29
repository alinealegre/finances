package controllers

import (
	"finances/database"
	"finances/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler para calcular o score do usuário
func CalculateUserScoreHandler(c *gin.Context) {
	db := database.GetDatabase() // Obtenha o objeto *gorm.DB do seu arquivo database/database.go
	cpf := c.Param("cpf")

	// Chame a função CalculateUserScore do pacote services para calcular o score do usuário
	score := services.CalculateUserScore(db, cpf)

	// Retorne o score do usuário como resposta
	c.JSON(http.StatusOK, gin.H{"score": score})
}
