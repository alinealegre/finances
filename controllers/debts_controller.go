package controllers

import (
	"net/http"
	"strconv"

	"finances/database"
	"finances/models"

	"github.com/gin-gonic/gin"
)

// eu acho que essa função quebraria o banco de dados, a menos que seja atrelada a um só cpf.
func ShowDebtsByUser(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.Debts
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func GetDebtsByCompany(c *gin.Context) {
	db := database.GetDatabase()

	company := c.Param("company")

	var debts []models.Debts
	if err := db.Where("company = ?", company).Find(&debts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar dívidas"})
		return
	}

	c.JSON(http.StatusOK, debts)
}

func ShowDebt(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID precisa ser um número inteiro",
		})
		return
	}

	db := database.GetDatabase()
	var p models.Debts
	err = db.First(&p, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func CreateDebt(c *gin.Context) {
	user := c.Request.Header.Get("Authorization")
	println(user)
	db := database.GetDatabase()

	var p models.Debts

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

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
