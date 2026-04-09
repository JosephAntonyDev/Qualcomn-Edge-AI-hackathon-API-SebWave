package app

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/domain/repository"
	"github.com/google/uuid"
)

type ChangeStatusAlertUseCase struct {
	repo repository.AlertRepository
}

func NewChangeStatusAlertUseCase(repo repository.AlertRepository) *ChangeStatusAlertUseCase {
	return &ChangeStatusAlertUseCase{repo: repo}
}

func (uc *ChangeStatusAlertUseCase) Execute(ctx context.Context, id uuid.UUID, status entities.AlertStatus, resolvedBy *uuid.UUID, note *string) error {
	return uc.repo.ChangeStatus(ctx, id, status, resolvedBy, note)
}
