package app

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/domain/repository"
	"github.com/google/uuid"
)

type GetAlertUseCase struct {
	repo repository.AlertRepository
}

func NewGetAlertUseCase(repo repository.AlertRepository) *GetAlertUseCase {
	return &GetAlertUseCase{repo: repo}
}

func (uc *GetAlertUseCase) Execute(ctx context.Context, id uuid.UUID) (*entities.Alert, error) {
	return uc.repo.GetByID(ctx, id)
}
