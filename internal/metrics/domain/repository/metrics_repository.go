package repository

import (
	"context"
	"time"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/domain/entities"
	"github.com/google/uuid"
)

type MetricsFilter struct {
	IntersectionID *uuid.UUID
	StartDate      *time.Time
	EndDate        *time.Time
	Limit          int
	Offset         int
}

type MetricsRepository interface {
	SaveDailyMetric(ctx context.Context, metric *entities.DailyMetric) error
	ListDailyMetrics(ctx context.Context, filter MetricsFilter) ([]*entities.DailyMetric, error)

	SaveSystemDailyMetric(ctx context.Context, metric *entities.SystemDailyMetric) error
	ListSystemDailyMetrics(ctx context.Context, filter MetricsFilter) ([]*entities.SystemDailyMetric, error)
}
