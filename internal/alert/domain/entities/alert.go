package entities

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type AlertType string
type AlertSeverity string
type AlertStatus string

const (
	TypeSirenDetected     AlertType = "siren_detected"
	TypeHighCongestion    AlertType = "high_congestion"
	TypeNodeDisconnected  AlertType = "node_disconnected"
	TypeSensorFailure     AlertType = "sensor_failure"
	TypePowerFailure      AlertType = "power_failure"
	TypeModelAnomaly      AlertType = "model_anomaly"
	TypeWatchdogTriggered AlertType = "watchdog_triggered"
	TypeManual            AlertType = "manual"
)

const (
	SeverityCritical AlertSeverity = "critical"
	SeverityWarning  AlertSeverity = "warning"
	SeverityInfo     AlertSeverity = "info"
)

const (
	StatusActive       AlertStatus = "active"
	StatusResolved     AlertStatus = "resolved"
	StatusAcknowledged AlertStatus = "acknowledged"
)

type Alert struct {
	ID             uuid.UUID       `json:"id"`
	IntersectionID *uuid.UUID      `json:"intersection_id,omitempty"`
	Type           AlertType       `json:"type"`
	Severity       AlertSeverity   `json:"severity"`
	Status         AlertStatus     `json:"status"`
	Title          string          `json:"title"`
	Description    *string         `json:"description,omitempty"`
	Metadata       json.RawMessage `json:"metadata,omitempty"`
	ResolvedBy     *uuid.UUID      `json:"resolved_by,omitempty"`
	ResolvedAt     *time.Time      `json:"resolved_at,omitempty"`
	ResolutionNote *string         `json:"resolution_note,omitempty"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}
