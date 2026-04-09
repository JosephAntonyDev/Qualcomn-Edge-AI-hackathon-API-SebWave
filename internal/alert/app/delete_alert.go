package app

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/domain/repository"
	"github.com/google/uuid"
)

type DeleteAlertUseCase struct {
	repo repository.AlertRepository
}

func NewDeleteAlertUseCase(repo repository.AlertRepository) *DeleteAlertUseCase {
	return &DeleteAlertUseCase{repo: repo}
}

func (uc *DeleteAlertUseCase) Execute(ctx context.Context, id uuid.UUID) error {
	return uc.repo.Delete(ctx, id)
}
