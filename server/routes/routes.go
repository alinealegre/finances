package routes

import (
	"finances/controllers"
	"finances/server/middlewares"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")

	{
		// Rotas sem autenticação
		main.POST("/cadastro", controllers.CreateUser)
		main.POST("/entrar", controllers.Login)

		// Rotas para usuários autenticados

		usuario := main.Group("usuario", middlewares.Auth())
		{
			usuario.GET("/dividas", controllers.ShowDebtsByUser)
			usuario.GET("/score", controllers.CalculateUserScoreHandler)
			usuario.PUT("/editar/:cpf", controllers.UptadeUserInfo)

			// Rota para logout
			usuario.POST("/sair", controllers.Logout)
		}

		// Rotas adicionais para usuários administradores
		admin := main.Group("admin", middlewares.Auth())
		{
			admin.POST("/adicionar", controllers.CreateDebt)
			admin.GET("/empresa/:company", controllers.GetDebtsByCompany)
			admin.PATCH("/atualizar-divida/:id", controllers.EditDebt)
			admin.DELETE("/deletar-divida/:id", controllers.DeleteDebt)
			admin.DELETE("/deletar/:user", controllers.DeleteUser)
			admin.GET("/divida/:id", controllers.ShowDebt)

		}
	}

	return router
}
