package controllers

import (
	"net/http"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/audit/app"
	"github.com/gin-gonic/gin"
)

type CreateLogController struct {
	UseCase *app.CreateLogUseCase
}

func NewCreateLogController(useCase *app.CreateLogUseCase) *CreateLogController {
	return &CreateLogController{UseCase: useCase}
}

type CreateLogRequest struct {
	UserID         string `json:"user_id" binding:"required"`
	Action         string `json:"action" binding:"required"`
	TargetResource string `json:"target_resource"`
	TargetID       string `json:"target_id"`
	Details        string `json:"details"`
}

func (c *CreateLogController) Handle(ctx *gin.Context) {
	var req CreateLogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.UseCase.Execute(ctx.Request.Context(), req.UserID, req.Action, req.TargetResource, req.TargetID, req.Details)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Audit log created successfully"})
}
