package controllers

import (
	"finances/database"
	"finances/models"
	"finances/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Criar o usuário
func CreateUser(c *gin.Context) {
	db := database.GetDatabase()

	var p models.User

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível processar os dados enviados " + err.Error(),
		})
		return
	}

	var User = make(map[string]models.User)
	if _, ok := User[p.CPF]; ok {
		c.JSON(http.StatusConflict, gin.H{"error": "Usuário já existe"})
		return
	}

	p.Type = "usuario"
	if strings.HasSuffix(p.Email, "@finances.com") {
		p.Type = "admin"
	}

	p.Password = services.SHA256Encoder(p.Password)

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível criar usuário: " + err.Error(),
		})
		return
	}

	c.Status(201)
}

func DeleteUser(c *gin.Context) {
	db := database.GetDatabase()

	userID := c.Param("id")

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir usuário"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário excluído com sucesso"})
}
