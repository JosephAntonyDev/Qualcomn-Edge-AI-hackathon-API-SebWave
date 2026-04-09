package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/app"
	"github.com/gin-gonic/gin"
)

type RegisterUserController struct {
	useCase *app.RegisterUserUseCase
}

func NewRegisterUserController(uc *app.RegisterUserUseCase) *RegisterUserController {
	return &RegisterUserController{
		useCase: uc,
	}
}

func (ctrl *RegisterUserController) Handle(c *gin.Context) {
	var req app.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: verifica los campos requeridos"})
		return
	}

	user, err := ctrl.useCase.Execute(c.Request.Context(), req)
	if err != nil {
		if err.Error() == "el usuario ya está registrado en el sistema" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado exitosamente",
		"data":    user,
	})
}
