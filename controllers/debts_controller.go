package controllers

import (
	"strconv"

	"finances/database"
	"finances/models"

	"github.com/gin-gonic/gin"
)

func ShowDebtsByUser(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.Debts

	userCpf := c.GetString("userCPF")

	err := db.Where("cpf = ?", userCpf).Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível listar dívidas: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func CreateDebt(c *gin.Context) {
	db := database.GetDatabase()

	userType := c.GetString("userType")
	if userType != "admin" {
		c.JSON(400, gin.H{
			"error": "Usuario nao é admin",
		})
		return
	}

	cpf := c.Param("cpf")

	var p models.Debts

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível vincular JSON:  " + err.Error(),
		})
		return
	}
	p.CPF = cpf

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível criar um usuário: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
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

	var p models.Debts

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Dívida não encontrada " + err.Error(),
		})
		return
	}

	err = db.Save(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível criar usuário: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}
