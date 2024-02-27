package controllers

import (
	"finances/database"
	"finances/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CleanupHandler(c *gin.Context) {
	db := database.GetDatabase() // Obtenha o objeto *gorm.DB do seu arquivo database/database.go

	// Obtenha o CPF do usuário a partir do contexto, por exemplo, de um token JWT
	cpf := "12345678900" // Substitua pelo CPF real do usuário obtido do token ou de outra fonte

	// Consulte as dívidas do usuário no banco de dados
	var debts []models.Debts
	db.Where("cpf = ?", cpf).Find(&debts)

	// Prepare a resposta contendo apenas as informações necessárias sobre as dívidas
	var response []gin.H
	totalDebt := 0
	for _, debt := range debts {
		totalDebt += int(debt.Value)
		// Construa um mapa contendo as informações específicas de cada dívida
		debtInfo := gin.H{
			"value":        debt.Value,
			"expirated_at": debt.ExpiratedAt,
			"cpf":          debt.CPF,
			"company_name": debt.Company,
		}
		response = append(response, debtInfo)
	}

	// Retorne as informações das dívidas como resposta
	c.JSON(http.StatusOK, response)
}
