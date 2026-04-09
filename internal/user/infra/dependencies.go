package infra

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/app"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/infra/adapters"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/infra/controllers"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/infra/repository"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/infra/routes"
)

func SetupDependencies(r *gin.Engine, db *sql.DB, jwtSecret string) {
	userRepo := repository.NewPostgresUserRepository(db)

	hasher := adapters.NewBcrypt()
	tokenManager := adapters.NewJWTManager(jwtSecret)

	createUserUseCase := app.NewCreateUserUseCase(userRepo, hasher)
	registerUserUseCase := app.NewRegisterUserUseCase(userRepo, hasher)
	loginUserUseCase := app.NewLoginUserUseCase(userRepo, hasher, tokenManager)

	createUserCtrl := controllers.NewCreateUserController(createUserUseCase)
	registerUserCtrl := controllers.NewRegisterUserController(registerUserUseCase)
	loginUserCtrl := controllers.NewLoginUserController(loginUserUseCase)

	routes.SetupUserRoutes(r, createUserCtrl, registerUserCtrl, loginUserCtrl, jwtSecret)
}
