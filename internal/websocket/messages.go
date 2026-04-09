package websocket

import (
	"encoding/json"
)

type MessageType string

const (
	// Desde Arduino → Backend
	MsgSensorData   MessageType = "sensor_data"
	MsgEmergencyOn  MessageType = "emergency_on"
	MsgEmergencyOff MessageType = "emergency_off"
	MsgPhaseChange  MessageType = "phase_change"

	// Desde Backend → App móvil
	MsgSensorUpdate       MessageType = "sensor_update"
	MsgEmergencyAlert     MessageType = "emergency_alert"
	MsgIntersectionStatus MessageType = "intersection_status"

	// Desde App móvil → Backend
	MsgTriggerEmergency MessageType = "trigger_emergency"
	MsgAdjustCycle      MessageType = "adjust_cycle"
)

// Estructura genérica del mensaje WebSocket
type WSMessage struct {
	Type           MessageType     `json:"type"`
	IntersectionID string          `json:"intersection_id,omitempty"`
	Payload        json.RawMessage `json:"payload"`
}

// Ejemplos de Payloads (se pueden expandir en el futuro)
type SensorDataPayload struct {
	DensityNS  *float64 `json:"density_ns,omitempty"`
	DensityEO  *float64 `json:"density_eo,omitempty"`
	IsOccupied bool     `json:"is_occupied"`
	DistanceMm int      `json:"distance_mm,omitempty"`
}

type PhaseChangePayload struct {
	PhaseName string `json:"phase_name"`
	Duration  int    `json:"duration_s"` // O duración en ms
}
