package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/domain/entities"
	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/metrics/domain/repository"
)

type PostgresMetricsRepository struct {
	db *sql.DB
}

func NewPostgresMetricsRepository(db *sql.DB) repository.MetricsRepository {
	return &PostgresMetricsRepository{db: db}
}

func (r *PostgresMetricsRepository) SaveDailyMetric(ctx context.Context, m *entities.DailyMetric) error {
	query := `
		INSERT INTO daily_metrics (
			intersection_id, metric_date, total_cycles, adaptive_cycles, emergency_events, 
			total_vehicles, avg_density_pct, peak_density_pct, peak_density_at, 
			total_wait_saved_ms, total_co2_saved_g, uptime_pct
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		ON CONFLICT (intersection_id, metric_date) DO UPDATE SET
			total_cycles = EXCLUDED.total_cycles,
			adaptive_cycles = EXCLUDED.adaptive_cycles,
			emergency_events = EXCLUDED.emergency_events,
			total_vehicles = EXCLUDED.total_vehicles,
			avg_density_pct = EXCLUDED.avg_density_pct,
			peak_density_pct = EXCLUDED.peak_density_pct,
			peak_density_at = EXCLUDED.peak_density_at,
			total_wait_saved_ms = EXCLUDED.total_wait_saved_ms,
			total_co2_saved_g = EXCLUDED.total_co2_saved_g,
			uptime_pct = EXCLUDED.uptime_pct
		RETURNING id, created_at
	`
	return r.db.QueryRowContext(ctx, query,
		m.IntersectionID, m.MetricDate, m.TotalCycles, m.AdaptiveCycles, m.EmergencyEvents,
		m.TotalVehicles, m.AvgDensityPct, m.PeakDensityPct, m.PeakDensityAt,
		m.TotalWaitSavedMs, m.TotalCO2SavedG, m.UptimePct,
	).Scan(&m.ID, &m.CreatedAt)
}

func (r *PostgresMetricsRepository) ListDailyMetrics(ctx context.Context, filter repository.MetricsFilter) ([]*entities.DailyMetric, error) {
	query := `
		SELECT id, intersection_id, metric_date, total_cycles, adaptive_cycles, emergency_events,
		       total_vehicles, avg_density_pct, peak_density_pct, peak_density_at,
		       total_wait_saved_ms, total_co2_saved_g, uptime_pct, created_at
		FROM daily_metrics WHERE 1=1
	`
	var args []interface{}
	var conditions []string
	argID := 1

	if filter.IntersectionID != nil {
		conditions = append(conditions, fmt.Sprintf("intersection_id = $%d", argID))
		args = append(args, *filter.IntersectionID)
		argID++
	}
	if filter.StartDate != nil {
		conditions = append(conditions, fmt.Sprintf("metric_date >= $%d", argID))
		args = append(args, *filter.StartDate)
		argID++
	}
	if filter.EndDate != nil {
		conditions = append(conditions, fmt.Sprintf("metric_date <= $%d", argID))
		args = append(args, *filter.EndDate)
		argID++
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	query += fmt.Sprintf(" ORDER BY metric_date DESC, intersection_id ASC LIMIT $%d OFFSET $%d", argID, argID+1)
	args = append(args, filter.Limit, filter.Offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var metrics []*entities.DailyMetric
	for rows.Next() {
		var m entities.DailyMetric
		var avgDensity, peakDensity, uptime sql.NullFloat64
		var peakAt sql.NullTime

		err := rows.Scan(
			&m.ID, &m.IntersectionID, &m.MetricDate, &m.TotalCycles, &m.AdaptiveCycles, &m.EmergencyEvents,
			&m.TotalVehicles, &avgDensity, &peakDensity, &peakAt,
			&m.TotalWaitSavedMs, &m.TotalCO2SavedG, &uptime, &m.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		if avgDensity.Valid {
			m.AvgDensityPct = &avgDensity.Float64
		}
		if peakDensity.Valid {
			m.PeakDensityPct = &peakDensity.Float64
		}
		if peakAt.Valid {
			m.PeakDensityAt = &peakAt.Time
		}
		if uptime.Valid {
			m.UptimePct = &uptime.Float64
		}

		metrics = append(metrics, &m)
	}

	return metrics, nil
}

func (r *PostgresMetricsRepository) SaveSystemDailyMetric(ctx context.Context, m *entities.SystemDailyMetric) error {
	query := `
		INSERT INTO system_daily_metrics (
			metric_date, active_nodes, total_vehicles, total_emergencies,
			total_wait_saved_ms, total_co2_saved_g
		) VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (metric_date) DO UPDATE SET
			active_nodes = EXCLUDED.active_nodes,
			total_vehicles = EXCLUDED.total_vehicles,
			total_emergencies = EXCLUDED.total_emergencies,
			total_wait_saved_ms = EXCLUDED.total_wait_saved_ms,
			total_co2_saved_g = EXCLUDED.total_co2_saved_g
		RETURNING created_at
	`
	return r.db.QueryRowContext(ctx, query,
		m.MetricDate, m.ActiveNodes, m.TotalVehicles, m.TotalEmergencies,
		m.TotalWaitSavedMs, m.TotalCO2SavedG,
	).Scan(&m.CreatedAt)
}

func (r *PostgresMetricsRepository) ListSystemDailyMetrics(ctx context.Context, filter repository.MetricsFilter) ([]*entities.SystemDailyMetric, error) {
	query := `
		SELECT metric_date, active_nodes, total_vehicles, total_emergencies,
		       total_wait_saved_ms, total_co2_saved_g, created_at
		FROM system_daily_metrics WHERE 1=1
	`
	var args []interface{}
	var conditions []string
	argID := 1

	if filter.StartDate != nil {
		conditions = append(conditions, fmt.Sprintf("metric_date >= $%d", argID))
		args = append(args, *filter.StartDate)
		argID++
	}
	if filter.EndDate != nil {
		conditions = append(conditions, fmt.Sprintf("metric_date <= $%d", argID))
		args = append(args, *filter.EndDate)
		argID++
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	query += fmt.Sprintf(" ORDER BY metric_date DESC LIMIT $%d OFFSET $%d", argID, argID+1)
	args = append(args, filter.Limit, filter.Offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var metrics []*entities.SystemDailyMetric
	for rows.Next() {
		var m entities.SystemDailyMetric
		err := rows.Scan(
			&m.MetricDate, &m.ActiveNodes, &m.TotalVehicles, &m.TotalEmergencies,
			&m.TotalWaitSavedMs, &m.TotalCO2SavedG, &m.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, &m)
	}

	return metrics, nil
}
