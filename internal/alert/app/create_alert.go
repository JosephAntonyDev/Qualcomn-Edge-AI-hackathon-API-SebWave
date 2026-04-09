package app

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/domain/repository"
)

type CreateAlertUseCase struct {
	repo repository.AlertRepository
}

func NewCreateAlertUseCase(repo repository.AlertRepository) *CreateAlertUseCase {
	return &CreateAlertUseCase{repo: repo}
}

func (uc *CreateAlertUseCase) Execute(ctx context.Context, alert *entities.Alert) error {
	// default status is active
	alert.Status = entities.StatusActive
	return uc.repo.Create(ctx, alert)
}
