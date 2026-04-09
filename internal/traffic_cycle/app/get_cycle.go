package app

import (
	"context"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/traffic_cycle/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/traffic_cycle/domain/repository"
)

type GetCycleUseCase struct {
	repo repository.TrafficCycleRepository
}

func NewGetCycleUseCase(r repository.TrafficCycleRepository) *GetCycleUseCase {
	return &GetCycleUseCase{repo: r}
}

func (uc *GetCycleUseCase) Execute(ctx context.Context, id int64) (*entities.TrafficCycle, error) {
	if id <= 0 {
		return nil, errors.New("id de ciclo inválido")
	}
	return uc.repo.GetCycleByID(ctx, id)
}
