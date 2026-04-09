package controllers

import (
	"net/http"
	"strconv"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/traffic_cycle/app"
	"github.com/gin-gonic/gin"
)

type GetCycleController struct {
	useCase *app.GetCycleUseCase
}

func NewGetCycleController(uc *app.GetCycleUseCase) *GetCycleController {
	return &GetCycleController{useCase: uc}
}

func (ctrl *GetCycleController) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	res, err := ctrl.useCase.Execute(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
