package controllers

import (
	"net/http"
	"strconv"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/app"
	"github.com/gin-gonic/gin"
)

type DeactivateCorridorController struct {
	uc *app.DeactivateCorridorUseCase
}

func NewDeactivateCorridorController(uc *app.DeactivateCorridorUseCase) *DeactivateCorridorController {
	return &DeactivateCorridorController{uc: uc}
}

func (ctrl *DeactivateCorridorController) Handle(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido."})
		return
	}

	if err := ctrl.uc.Execute(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al desactivar el corredor verde o ya estaba resuelto", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Corredor verde desactivado exitosamente (evento resuelto)"})
}
