package repository

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/domain/entities"
	"github.com/google/uuid"
)

type EmergencyFilter struct {
	IntersectionID    *uuid.UUID
	DetectionMethod   *entities.DetectionMethod
	CorridorActivated *bool
	Limit             int
	Offset            int
}

type EmergencyRepository interface {
	Record(ctx context.Context, em *entities.EmergencyEvent) error
	GetByID(ctx context.Context, id int64) (*entities.EmergencyEvent, error)
	List(ctx context.Context, filter EmergencyFilter) ([]*entities.EmergencyEvent, error)
	UpdateCorridorStatus(ctx context.Context, id int64, activate bool, responseTimeMs *int) error
	Resolve(ctx context.Context, id int64) error
	Delete(ctx context.Context, id int64) error
}
