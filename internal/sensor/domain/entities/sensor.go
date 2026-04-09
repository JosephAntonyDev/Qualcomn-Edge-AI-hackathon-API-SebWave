package entities

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type SensorType string

const (
	TypeRadarLD2410c  SensorType = "radar_ld2410c"
	TypeModulinoDist  SensorType = "modulino_distance"
	TypeCameraBrio101 SensorType = "camera_brio101"
	TypeMicINMP441    SensorType = "mic_inmp441"
	TypeVoltage       SensorType = "voltage"
	TypeSupercap      SensorType = "supercap"
)

type Sensor struct {
	ID             uuid.UUID  `json:"id"`
	IntersectionID uuid.UUID  `json:"intersection_id"`
	SensorType     SensorType `json:"sensor_type"`
	LaneDirection  *string    `json:"lane_direction,omitempty"`
	ConnectionType *string    `json:"connection_type,omitempty"`
	PinAssignment  *string    `json:"pin_assignment,omitempty"`
	IsActive       bool       `json:"is_active"`
	FailureCount   int        `json:"failure_count"`
	LastReadingAt  *time.Time `json:"last_reading_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
}

type SensorReading struct {
	ID             int64           `json:"id"`
	SensorID       uuid.UUID       `json:"sensor_id"`
	IntersectionID uuid.UUID       `json:"intersection_id"`
	IsOccupied     *bool           `json:"is_occupied,omitempty"`
	DistanceMM     *int            `json:"distance_mm,omitempty"`
	EnergyPct      *int            `json:"energy_pct,omitempty"`
	Confidence     *float64        `json:"confidence,omitempty"`
	RawData        json.RawMessage `json:"raw_data,omitempty"`
	CreatedAt      time.Time       `json:"created_at"`
}
