package middlewares

import (
	"finances/services"

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
		c.Set("userCPF", claims.CPF) //end point de mostrar dividas de usuario, comparar o cpf do parametro com o cpf do token
	}
}
