package entities

import (
	"time"

	"github.com/google/uuid"
)

type DetectionMethod string

const (
	DetectionCameraMic DetectionMethod = "camera_mic"
	DetectionINMP441   DetectionMethod = "inmp441"
	DetectionManual    DetectionMethod = "manual"
	DetectionAPI       DetectionMethod = "api"
)

type EmergencyEvent struct {
	ID                int64           `json:"id"`
	IntersectionID    uuid.UUID       `json:"intersection_id"`
	ConfidenceScore   float64         `json:"confidence_score"`
	DetectionMethod   DetectionMethod `json:"detection_method"`
	ThresholdPassed   bool            `json:"threshold_passed"`
	PersistencePassed bool            `json:"persistence_passed"`
	RadarCorrPassed   bool            `json:"radar_corr_passed"`
	CorridorActivated bool            `json:"corridor_activated"`
	ResponseTimeMs    *int            `json:"response_time_ms,omitempty"`
	DetectedAt        time.Time       `json:"detected_at"`
	ResolvedAt        *time.Time      `json:"resolved_at,omitempty"`
}
