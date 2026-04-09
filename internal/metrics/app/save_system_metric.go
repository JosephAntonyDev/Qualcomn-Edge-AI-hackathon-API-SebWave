package app

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/domain/repository"
)

type SaveSystemMetricUseCase struct {
	repo repository.MetricsRepository
}

func NewSaveSystemMetricUseCase(repo repository.MetricsRepository) *SaveSystemMetricUseCase {
	return &SaveSystemMetricUseCase{repo: repo}
}

func (uc *SaveSystemMetricUseCase) Execute(ctx context.Context, metric *entities.SystemDailyMetric) error {
	return uc.repo.SaveSystemDailyMetric(ctx, metric)
}
