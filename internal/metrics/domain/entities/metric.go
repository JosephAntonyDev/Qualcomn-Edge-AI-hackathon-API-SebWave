package entities

import (
	"time"

	"github.com/google/uuid"
)

type DailyMetric struct {
	ID               int64      `json:"id"`
	IntersectionID   uuid.UUID  `json:"intersection_id"`
	MetricDate       time.Time  `json:"metric_date"` // Solo fecha
	TotalCycles      int        `json:"total_cycles"`
	AdaptiveCycles   int        `json:"adaptive_cycles"`
	EmergencyEvents  int        `json:"emergency_events"`
	TotalVehicles    int        `json:"total_vehicles"`
	AvgDensityPct    *float64   `json:"avg_density_pct,omitempty"`
	PeakDensityPct   *float64   `json:"peak_density_pct,omitempty"`
	PeakDensityAt    *time.Time `json:"peak_density_at,omitempty"`
	TotalWaitSavedMs int64      `json:"total_wait_saved_ms"`
	TotalCO2SavedG   float64    `json:"total_co2_saved_g"`
	UptimePct        *float64   `json:"uptime_pct,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
}

type SystemDailyMetric struct {
	MetricDate       time.Time `json:"metric_date"` // PK
	ActiveNodes      int       `json:"active_nodes"`
	TotalVehicles    int       `json:"total_vehicles"`
	TotalEmergencies int       `json:"total_emergencies"`
	TotalWaitSavedMs int64     `json:"total_wait_saved_ms"`
	TotalCO2SavedG   float64   `json:"total_co2_saved_g"`
	CreatedAt        time.Time `json:"created_at"`
}
