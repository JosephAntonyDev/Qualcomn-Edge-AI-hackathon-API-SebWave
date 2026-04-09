package repository

import (
	"context"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/domain/entities"
	"github.com/google/uuid"
)

type SensorRepository interface {
	RegisterSensor(ctx context.Context, sensor *entities.Sensor) error
	GetSensor(ctx context.Context, id uuid.UUID) (*entities.Sensor, error)
	ListSensorsByIntersection(ctx context.Context, intersectionID uuid.UUID) ([]*entities.Sensor, error)
	UpdateSensor(ctx context.Context, sensor *entities.Sensor) error
	DeleteSensor(ctx context.Context, id uuid.UUID) error

	RecordReading(ctx context.Context, reading *entities.SensorReading) error
	GetReadings(ctx context.Context, sensorID uuid.UUID, limit int) ([]*entities.SensorReading, error)
}
