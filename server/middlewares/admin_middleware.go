package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func AdminPermission(c *gin.Context) bool {
// 	userType := c.GetString("userType")
// 	fmt.Println("User type:", userType)
// 	return userType == "admin"
// }

// func SendAdminPermissionError(c *gin.Context) {
// 	c.JSON(400, gin.H{
// 		"error": "Usuário não tem permissão de acesso",
// 	})
// }

func AdminPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		userType := c.GetString("userType")
		if userType != "admin" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
