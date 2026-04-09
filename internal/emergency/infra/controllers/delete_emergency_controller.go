package controllers

import (
	"net/http"
	"strconv"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/app"
	"github.com/gin-gonic/gin"
)

type DeleteEmergencyController struct {
	uc *app.DeleteEmergencyUseCase
}

func NewDeleteEmergencyController(uc *app.DeleteEmergencyUseCase) *DeleteEmergencyController {
	return &DeleteEmergencyController{uc: uc}
}

func (ctrl *DeleteEmergencyController) Handle(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido. Debe ser un entero."})
		return
	}

	if err := ctrl.uc.Execute(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la emergencia", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Emergencia eliminada exitosamente"})
}
