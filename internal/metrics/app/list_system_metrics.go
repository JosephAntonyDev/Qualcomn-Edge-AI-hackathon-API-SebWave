package app

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/domain/repository"
)

type ListSystemMetricsUseCase struct {
	repo repository.MetricsRepository
}

func NewListSystemMetricsUseCase(repo repository.MetricsRepository) *ListSystemMetricsUseCase {
	return &ListSystemMetricsUseCase{repo: repo}
}

func (uc *ListSystemMetricsUseCase) Execute(ctx context.Context, filter repository.MetricsFilter) ([]*entities.SystemDailyMetric, error) {
	if filter.Limit == 0 {
		filter.Limit = 50
	}
	return uc.repo.ListSystemDailyMetrics(ctx, filter)
}
