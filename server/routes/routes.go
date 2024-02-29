package routes

import (
	"finances/controllers"
	"finances/server/middlewares"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")

	{
		// Rotas públicas
		main.POST("/cadastro", controllers.CreateUser)
		main.POST("/entrar", controllers.Login)
		main.POST("/sair", controllers.Logout)
	}

	// Rotas para usuários autenticados
	usuario := main.Group("usuario", middlewares.Auth())
	{
		usuario.GET("/dividas", controllers.ShowDebtsByUser)
		usuario.GET("/score/:cpf", middlewares.CheckCPF(), controllers.CalculateUserScoreHandler)

		// Rotas adicionais para usuários administradores
		admin := main.Group("admin", middlewares.AdminRequired())
		{
			admin.POST("/adicionar/:cpf", controllers.CreateDebt)
			admin.PATCH("/atualizar-divida/:id", controllers.EditDebt)
			admin.DELETE("/deletar-divida/:id", controllers.DeleteDebt)
			admin.DELETE("/deletar/:cpf", controllers.DeleteUser)

		}
	}

	return router
}
