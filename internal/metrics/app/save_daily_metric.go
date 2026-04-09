package app

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/domain/repository"
)

type SaveDailyMetricUseCase struct {
	repo repository.MetricsRepository
}

func NewSaveDailyMetricUseCase(repo repository.MetricsRepository) *SaveDailyMetricUseCase {
	return &SaveDailyMetricUseCase{repo: repo}
}

func (uc *SaveDailyMetricUseCase) Execute(ctx context.Context, metric *entities.DailyMetric) error {
	return uc.repo.SaveDailyMetric(ctx, metric)
}
