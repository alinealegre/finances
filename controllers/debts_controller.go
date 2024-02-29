package controllers

import (
	"strconv"
	"time"

	"finances/database"
	"finances/models"

	"github.com/gin-gonic/gin"
)

func ShowDebtsByUser(c *gin.Context) {
	db := database.GetDatabase()
	var debts []models.Debts

	userCpf := c.GetString("userCPF")

	err := db.Where("cpf = ?", userCpf).Find(&debts).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível listar dívidas: " + err.Error(),
		})
		return
	}

	type DebtInfo struct {
		Value       float64   `json:"value"`
		ExpiratedAt time.Time `json:"expirated_at"`
		CPF         string    `json:"cpf"`
	}

	var debtsInfo []DebtInfo

	for _, debt := range debts {
		debtsInfo = append(debtsInfo, DebtInfo{
			Value:       debt.Value,
			ExpiratedAt: debt.ExpiratedAt,
			CPF:         debt.CPF,
		})
	}

	c.JSON(200, debtsInfo)
}
func CreateDebt(c *gin.Context) {
	db := database.GetDatabase()

	cpf := c.Param("cpf")

	var debts models.Debts

	err := c.ShouldBindJSON(&debts)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível vincular JSON:  " + err.Error(),
		})
		return
	}
	debts.CPF = cpf

	err = db.Create(&debts).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível criar um usuário: " + err.Error(),
		})
		return
	}

	c.JSON(200, debts)
}

func DeleteDebt(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID precisa ser um número inteiro",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Debts{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível deletar: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func EditDebt(c *gin.Context) {
	db := database.GetDatabase()

	var debts models.Debts

	err := c.ShouldBindJSON(&debts)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Dívida não encontrada " + err.Error(),
		})
		return
	}

	err = db.Save(&debts).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível criar usuário: " + err.Error(),
		})
		return
	}

	c.JSON(200, debts)
}
