package controllers

import (
	"finances/database"
	"finances/models"
	"finances/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db := database.GetDatabase()

	if _, exists := c.Get("tokenClaims"); exists {
		c.JSON(400, gin.H{
			"error": "Usuário já está autenticado",
		})
		return
	}

	var p models.Login
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível vincular JSON: " + err.Error(),
		})
		return
	}

	var user models.User
	dbError := db.Where("email = ?", p.Email).First(&user).Error
	if dbError != nil {
		c.JSON(400, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}

	if user.Password != services.SHA256Encoder(p.Password) {
		c.JSON(401, gin.H{
			"error": "Credenciais inválidas",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.CPF, user.Type, user.Email)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})

}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logout realizado com sucesso"})
}
