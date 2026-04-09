package entities

import (
	"time"

	"github.com/google/uuid"
)

type TrafficCycle struct {
	ID             int64     `json:"id"`
	IntersectionID uuid.UUID `json:"intersection_id"`
	OperationMode  string    `json:"operation_mode"` // Ej: 'fixed', 'adaptive', etc.

	// Tiempos
	GreenNSMs int `json:"green_ns_ms"`
	GreenEOMs int `json:"green_eo_ms"`
	YellowMs  int `json:"yellow_ms"`

	// Densidad vehicular por sensor
	RadarOccupied      *bool `json:"radar_occupied,omitempty"`
	RadarDistanceCm    *int  `json:"radar_distance_cm,omitempty"`
	RadarEnergy        *int  `json:"radar_energy,omitempty"`
	ModulinoOccupied   *bool `json:"modulino_occupied,omitempty"`
	ModulinoDistanceMm *int  `json:"modulino_distance_mm,omitempty"`
	CameraNSOccupied   *bool `json:"camera_ns_occupied,omitempty"`
	CameraEOOccupied   *bool `json:"camera_eo_occupied,omitempty"`

	// Métricas calculadas
	WaitTimeSavedMs *int     `json:"wait_time_saved_ms,omitempty"`
	CO2SavedG       *float64 `json:"co2_saved_g,omitempty"`

	// Predicción del modelo
	ModelConfidence *float64 `json:"model_confidence,omitempty"`

	StartedAt time.Time  `json:"started_at"`
	EndedAt   *time.Time `json:"ended_at,omitempty"`
}
