package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/app"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateAlertController struct {
	uc *app.CreateAlertUseCase
}

func NewCreateAlertController(uc *app.CreateAlertUseCase) *CreateAlertController {
	return &CreateAlertController{uc: uc}
}

func (ctrl *CreateAlertController) Handle(c *gin.Context) {
	var alert entities.Alert
	if err := c.ShouldBindJSON(&alert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido", "details": err.Error()})
		return
	}

	if err := ctrl.uc.Execute(c.Request.Context(), &alert); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la alerta", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Alerta creada exitosamente", "data": alert})
}
