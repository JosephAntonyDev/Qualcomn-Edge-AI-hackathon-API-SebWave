package repository

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/traffic_cycle/domain/entities"
	"github.com/google/uuid"
)

type TrafficCycleRepository interface {
	RecordCycle(ctx context.Context, cycle *entities.TrafficCycle) error
	GetCycleByID(ctx context.Context, id int64) (*entities.TrafficCycle, error)
	ListCyclesByIntersection(ctx context.Context, intersectionID uuid.UUID, limit int, offset int) ([]*entities.TrafficCycle, error)
	DeleteCycle(ctx context.Context, id int64) error
}
