package routes

import (
	"finances/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")

	{
		// Rotas sem autenticação
		main.POST("/cadastro", controllers.CreateUser)
		main.POST("/entrar", controllers.Login)

		// Rotas para usuários autenticados
		// usuario := main.Group("usuario", middlewares.Auth())
		usuario := main.Group("usuario")
		{
			usuario.GET("/dividas/:cpf", controllers.ShowDebtsByUser)
			usuario.GET("/score/:cpf", controllers.CalculateUserScoreHandler)

			// Rota para logout
			usuario.POST("/sair", controllers.Logout)
		}

		// Rotas adicionais para usuários administradores
		// admin := main.Group("admin", middlewares.Auth())
		admin := main.Group("admin")
		{
			admin.POST("/adicionar/:cpf", controllers.CreateDebt)
			admin.PATCH("/atualizar-divida/:id", controllers.EditDebt)
			admin.DELETE("/deletar-divida/:id", controllers.DeleteDebt)
			admin.DELETE("/deletar/:cpf", controllers.DeleteUser)

		}
	}

	return router
}
