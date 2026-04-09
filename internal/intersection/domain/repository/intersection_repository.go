package repository

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/domain/entities"
	"github.com/google/uuid"
)

type IntersectionRepository interface {
	Create(ctx context.Context, intersection *entities.Intersection) error
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Intersection, error)
	GetBySerialNumber(ctx context.Context, serialNumber string) (*entities.Intersection, error)
	GetAll(ctx context.Context) ([]*entities.Intersection, error)
	Update(ctx context.Context, intersection *entities.Intersection) error
	Delete(ctx context.Context, id uuid.UUID) error
	UpdateState(ctx context.Context, id uuid.UUID, status entities.NodeStatus, mode entities.OperationMode, phase *string, densityNS, densityEO *float64) error
	Heartbeat(ctx context.Context, serialNumber string, firmwareVersion *string) error
}
