package middlewares

import (
	"finances/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatus(401)
		}

		jwtService := services.NewJWTService()
		claims, err := jwtService.ValidateToken(token)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}

		c.Set("userType", claims.Type)
		c.Set("userCPF", claims.CPF)
		c.Set("userEmail", claims.Email)
	}
}

func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		userEmail := c.GetString("userEmail")
		if !strings.HasSuffix(userEmail, "@finances.com") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Acesso restrito a administradores"})
			return
		}
		c.Next()
	}
}
