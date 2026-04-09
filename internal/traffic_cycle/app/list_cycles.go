package app

import (
	"context"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/traffic_cycle/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/traffic_cycle/domain/repository"
	"github.com/google/uuid"
)

type ListCyclesUseCase struct {
	repo repository.TrafficCycleRepository
}

func NewListCyclesUseCase(r repository.TrafficCycleRepository) *ListCyclesUseCase {
	return &ListCyclesUseCase{repo: r}
}

func (uc *ListCyclesUseCase) Execute(ctx context.Context, intersectionID string, limit, offset int) ([]*entities.TrafficCycle, error) {
	uid, err := uuid.Parse(intersectionID)
	if err != nil {
		return nil, errors.New("id de intersección inválido")
	}

	if limit <= 0 || limit > 100 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	return uc.repo.ListCyclesByIntersection(ctx, uid, limit, offset)
}
