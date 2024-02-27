package controllers

import (
	"finances/database"
	"finances/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db := database.GetDatabase()

	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "não foi possível fazer o bind do JSON: " + err.Error(),
		})
		return
	}

	// Verificar as credenciais do usuário no banco de dados
	dbError := db.Where(&user).First(&user).Error
	if dbError != nil {
		c.JSON(401, gin.H{
			"error": "credenciais inválidas",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Usuário logado com sucesso",
		"isAdmin": user.Type == "admin",
	})
}
