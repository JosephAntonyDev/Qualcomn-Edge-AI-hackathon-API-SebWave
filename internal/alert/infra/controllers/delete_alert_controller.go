package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/app"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteAlertController struct {
	uc *app.DeleteAlertUseCase
}

func NewDeleteAlertController(uc *app.DeleteAlertUseCase) *DeleteAlertController {
	return &DeleteAlertController{uc: uc}
}

func (ctrl *DeleteAlertController) Handle(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := ctrl.uc.Execute(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la alerta", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alerta eliminada exitosamente"})
}
