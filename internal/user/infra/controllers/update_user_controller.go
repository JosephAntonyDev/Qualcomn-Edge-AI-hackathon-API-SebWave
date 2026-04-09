package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/app"
	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	useCase *app.UpdateUserUseCase
}

func NewUpdateUserController(uc *app.UpdateUserUseCase) *UpdateUserController {
	return &UpdateUserController{
		useCase: uc,
	}
}

func (ctrl *UpdateUserController) Handle(c *gin.Context) {
	requesterRoleVal, exists := c.Get("userRole")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se pudo identificar tu nivel de acceso"})
		return
	}
	requesterRole := requesterRoleVal.(string)

	requesterIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se pudo identificar tu usuario"})
		return
	}
	requesterID := requesterIDVal.(string)

	targetID := c.Param("id")

	var req app.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: verifica los campos requeridos"})
		return
	}

	user, err := ctrl.useCase.Execute(c.Request.Context(), targetID, requesterID, requesterRole, req)
	if err != nil {
		if len(err.Error()) >= 18 && err.Error()[:18] == "operación denegada" {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "id de usuario inválido" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuario actualizado exitosamente",
		"data":    user,
	})
}
