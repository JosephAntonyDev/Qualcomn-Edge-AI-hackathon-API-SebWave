package controllers

import (
	"net/http"
	"strconv"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/traffic_cycle/app"
	"github.com/gin-gonic/gin"
)

type DeleteCycleController struct {
	useCase *app.DeleteCycleUseCase
}

func NewDeleteCycleController(uc *app.DeleteCycleUseCase) *DeleteCycleController {
	return &DeleteCycleController{useCase: uc}
}

func (ctrl *DeleteCycleController) Handle(c *gin.Context) {
	role, _ := c.Get("userRole")

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = ctrl.useCase.Execute(c.Request.Context(), id, role.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ciclo de tráfico eliminado"})
}
