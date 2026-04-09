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
	getUserByUsernameUseCase := app.NewGetUserByUsernameUseCase(userRepo)
	getUserByIDUseCase := app.NewGetUserByIDUseCase(userRepo)
	updateUserUseCase := app.NewUpdateUserUseCase(userRepo)
	deleteUserUseCase := app.NewDeleteUserUseCase(userRepo)
	getMyProfileUseCase := app.NewGetMyProfileUseCase(userRepo)

	createUserCtrl := controllers.NewCreateUserController(createUserUseCase)
	registerUserCtrl := controllers.NewRegisterUserController(registerUserUseCase)
	loginUserCtrl := controllers.NewLoginUserController(loginUserUseCase)
	getUserByUsernameCtrl := controllers.NewGetUserByUsernameController(getUserByUsernameUseCase)
	getUserByIDCtrl := controllers.NewGetUserByIDController(getUserByIDUseCase)
	updateUserCtrl := controllers.NewUpdateUserController(updateUserUseCase)
	deleteUserCtrl := controllers.NewDeleteUserController(deleteUserUseCase)
	getMyProfileCtrl := controllers.NewGetMyProfileController(getMyProfileUseCase)

	routes.SetupUserRoutes(
		r,
		createUserCtrl,
		registerUserCtrl,
		loginUserCtrl,
		getUserByUsernameCtrl,
		getUserByIDCtrl,
		updateUserCtrl,
		deleteUserCtrl,
		getMyProfileCtrl,
		jwtSecret,
	)
}
