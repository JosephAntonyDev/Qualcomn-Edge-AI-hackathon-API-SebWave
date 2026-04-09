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
	getUserByUsernameCtrl *controllers.GetUserByUsernameController,
	getUserByIDCtrl *controllers.GetUserByIDController,
	updateUserCtrl *controllers.UpdateUserController,
	deleteUserCtrl *controllers.DeleteUserController,
	getMyProfileCtrl *controllers.GetMyProfileController,
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
			// Rutas que requieren estar autenticado
			protected.GET("/profile", getMyProfileCtrl.Handle)
			protected.GET("/search/:username", getUserByUsernameCtrl.Handle)
			protected.GET("/:id", getUserByIDCtrl.Handle)
			protected.PUT("/:id", updateUserCtrl.Handle)
			protected.DELETE("/:id", deleteUserCtrl.Handle)

			// Restringidas
			adminOnly := protected.Group("")
			adminOnly.Use(middleware.RequireRoles(entities.RoleAdmin))
			{
				adminOnly.POST("/create", createUserCtrl.Handle)
			}
		}
	}
}
