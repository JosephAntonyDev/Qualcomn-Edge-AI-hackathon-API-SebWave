package app

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/domain/repository"
)

type DeactivateCorridorUseCase struct {
	repo repository.EmergencyRepository
}

func NewDeactivateCorridorUseCase(repo repository.EmergencyRepository) *DeactivateCorridorUseCase {
	return &DeactivateCorridorUseCase{repo: repo}
}

func (uc *DeactivateCorridorUseCase) Execute(ctx context.Context, id int64) error {
	// También podríamos guardar un responseTime final si es requerido, pero por simplificación le pasamos nil
	err := uc.repo.UpdateCorridorStatus(ctx, id, false, nil)
	if err != nil {
		return err
	}
	return uc.repo.Resolve(ctx, id)
}
