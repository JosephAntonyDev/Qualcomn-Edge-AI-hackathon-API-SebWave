package entities

import (
	"time"

	"github.com/google/uuid"
)

type NodeStatus string
type OperationMode string

const (
	StatusConnected    NodeStatus = "connected"
	StatusDegraded     NodeStatus = "degraded"
	StatusDisconnected NodeStatus = "disconnected"
	StatusFailed       NodeStatus = "failed"
	StatusMaintenance  NodeStatus = "maintenance"
)

const (
	ModeFixed      OperationMode = "fixed"
	ModeAdaptive   OperationMode = "adaptive"
	ModeEmergency  OperationMode = "emergency"
	ModeAmberBlink OperationMode = "amber_blink"
)

type Intersection struct {
	ID               uuid.UUID     `json:"id"`
	SerialNumber     string        `json:"serial_number"`
	Name             string        `json:"name"`
	Description      *string       `json:"description,omitempty"`
	Latitude         float64       `json:"latitude"`
	Longitude        float64       `json:"longitude"`
	NodeID           *string       `json:"node_id,omitempty"`
	MaxCongestionPct float64       `json:"max_congestion_pct"`
	MinGreenTimeS    int           `json:"min_green_time_s"`
	MaxGreenTimeS    int           `json:"max_green_time_s"`
	DefaultGreenS    int           `json:"default_green_s"`
	DefaultRedS      int           `json:"default_red_s"`
	YellowTimeS      int           `json:"yellow_time_s"`
	Status           NodeStatus    `json:"status"`
	OperationMode    OperationMode `json:"operation_mode"`
	CurrentPhase     *string       `json:"current_phase,omitempty"`
	CurrentDensityNS *float64      `json:"current_density_ns,omitempty"`
	CurrentDensityEO *float64      `json:"current_density_eo,omitempty"`
	FirmwareVersion  *string       `json:"firmware_version,omitempty"`
	LastHeartbeatAt  *time.Time    `json:"last_heartbeat_at,omitempty"`
	CreatedBy        *uuid.UUID    `json:"created_by,omitempty"`
	CreatedAt        time.Time     `json:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at"`
	DeletedAt        *time.Time    `json:"deleted_at,omitempty"`
}
