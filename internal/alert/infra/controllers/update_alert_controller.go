package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/app"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/domain/entities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateAlertController struct {
	uc *app.UpdateAlertUseCase
}

func NewUpdateAlertController(uc *app.UpdateAlertUseCase) *UpdateAlertController {
	return &UpdateAlertController{uc: uc}
}

func (ctrl *UpdateAlertController) Handle(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var alert entities.Alert
	if err := c.ShouldBindJSON(&alert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido", "details": err.Error()})
		return
	}
	alert.ID = id

	if err := ctrl.uc.Execute(c.Request.Context(), &alert); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la alerta", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alerta actualizada exitosamente", "data": alert})
}
