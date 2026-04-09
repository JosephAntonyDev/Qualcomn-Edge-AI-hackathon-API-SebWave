package app

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/domain/repository"
)

type ListDailyMetricsUseCase struct {
	repo repository.MetricsRepository
}

func NewListDailyMetricsUseCase(repo repository.MetricsRepository) *ListDailyMetricsUseCase {
	return &ListDailyMetricsUseCase{repo: repo}
}

func (uc *ListDailyMetricsUseCase) Execute(ctx context.Context, filter repository.MetricsFilter) ([]*entities.DailyMetric, error) {
	if filter.Limit == 0 {
		filter.Limit = 50
	}
	return uc.repo.ListDailyMetrics(ctx, filter)
}
