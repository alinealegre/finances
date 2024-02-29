package controllers

import (
	"finances/database"
	"finances/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CalculateUserScoreHandler(c *gin.Context) {
	db := database.GetDatabase()

	score := services.CalculateUserScore(db)

	c.JSON(http.StatusOK, gin.H{"score": score})
}
