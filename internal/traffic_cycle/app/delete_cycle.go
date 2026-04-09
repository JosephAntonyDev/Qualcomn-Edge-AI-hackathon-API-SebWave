package app

import (
	"context"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/traffic_cycle/domain/repository"
	userEntities "github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/entities"
)

type DeleteCycleUseCase struct {
	repo repository.TrafficCycleRepository
}

func NewDeleteCycleUseCase(r repository.TrafficCycleRepository) *DeleteCycleUseCase {
	return &DeleteCycleUseCase{repo: r}
}

func (uc *DeleteCycleUseCase) Execute(ctx context.Context, id int64, requesterRole string) error {
	if requesterRole != string(userEntities.RoleAdmin) {
		return errors.New("operación denegada")
	}

	if id <= 0 {
		return errors.New("id de ciclo inválido")
	}

	return uc.repo.DeleteCycle(ctx, id)
}
