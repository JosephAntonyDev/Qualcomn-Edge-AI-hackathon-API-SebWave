package controllers

import (
	"net/http"
	"strconv"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/audit/app"
	"github.com/gin-gonic/gin"
)

type GetLogsController struct {
	UseCase *app.GetLogsUseCase
}

func NewGetLogsController(useCase *app.GetLogsUseCase) *GetLogsController {
	return &GetLogsController{UseCase: useCase}
}

func (c *GetLogsController) Handle(ctx *gin.Context) {
	limitStr := ctx.DefaultQuery("limit", "50")
	offsetStr := ctx.DefaultQuery("offset", "0")

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	logs, err := c.UseCase.Execute(ctx.Request.Context(), limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"logs": logs})
}
