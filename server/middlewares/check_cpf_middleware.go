package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckCPF() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInUserCPF := c.GetString("userCPF")
		requestedUserCPF := c.Param("cpf")

		if loggedInUserCPF != requestedUserCPF {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
