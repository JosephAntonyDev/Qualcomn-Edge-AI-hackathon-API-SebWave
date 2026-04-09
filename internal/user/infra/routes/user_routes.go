package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/middleware"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/infra/controllers"
)

func SetupUserRoutes(
	r *gin.Engine,
	createUserCtrl *controllers.CreateUserController,
	registerUserCtrl *controllers.RegisterUserController,
	loginUserCtrl *controllers.LoginUserController,
	jwtSecret string,
) {
	api := r.Group("/users")
	{
		// Públicas
		api.POST("/register", registerUserCtrl.Handle)
		api.POST("/login", loginUserCtrl.Handle)

		// Protegidas
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware(jwtSecret))
		{
			// Restringidas
			adminOnly := protected.Group("")
			adminOnly.Use(middleware.RequireRoles(entities.RoleAdmin))
			{
				adminOnly.POST("/create", createUserCtrl.Handle)
			}
		}
	}
}
