package app

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/domain/repository"
)

type UpdateAlertUseCase struct {
	repo repository.AlertRepository
}

func NewUpdateAlertUseCase(repo repository.AlertRepository) *UpdateAlertUseCase {
	return &UpdateAlertUseCase{repo: repo}
}

func (uc *UpdateAlertUseCase) Execute(ctx context.Context, alert *entities.Alert) error {
	return uc.repo.Update(ctx, alert)
}
