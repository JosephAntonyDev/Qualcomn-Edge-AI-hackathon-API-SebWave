package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/domain/entities"
	domainRepo "github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/emergency/domain/repository"
)

type PostgresEmergencyRepository struct {
	db *sql.DB
}

func NewPostgresEmergencyRepository(db *sql.DB) domainRepo.EmergencyRepository {
	return &PostgresEmergencyRepository{db: db}
}

func (r *PostgresEmergencyRepository) Record(ctx context.Context, em *entities.EmergencyEvent) error {
	if em.DetectionMethod == "" {
		em.DetectionMethod = entities.DetectionCameraMic
	}

	query := `
		INSERT INTO emergency_events (
			intersection_id, confidence_score, detection_method, 
			threshold_passed, persistence_passed, radar_corr_passed, 
			corridor_activated, response_time_ms
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, detected_at
	`
	err := r.db.QueryRowContext(ctx, query,
		em.IntersectionID, em.ConfidenceScore, em.DetectionMethod,
		em.ThresholdPassed, em.PersistencePassed, em.RadarCorrPassed,
		em.CorridorActivated, em.ResponseTimeMs,
	).Scan(&em.ID, &em.DetectedAt)

	return err
}

func (r *PostgresEmergencyRepository) GetByID(ctx context.Context, id int64) (*entities.EmergencyEvent, error) {
	query := `
		SELECT id, intersection_id, confidence_score, detection_method, 
		       threshold_passed, persistence_passed, radar_corr_passed, 
		       corridor_activated, response_time_ms, detected_at, resolved_at
		FROM emergency_events WHERE id = $1
	`
	var em entities.EmergencyEvent
	var responseTime sql.NullInt32
	var resolvedAt sql.NullTime

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&em.ID, &em.IntersectionID, &em.ConfidenceScore, &em.DetectionMethod,
		&em.ThresholdPassed, &em.PersistencePassed, &em.RadarCorrPassed,
		&em.CorridorActivated, &responseTime, &em.DetectedAt, &resolvedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("evento de emergencia no encontrado")
		}
		return nil, err
	}

	if responseTime.Valid {
		val := int(responseTime.Int32)
		em.ResponseTimeMs = &val
	}
	if resolvedAt.Valid {
		em.ResolvedAt = &resolvedAt.Time
	}

	return &em, nil
}

func (r *PostgresEmergencyRepository) List(ctx context.Context, filter domainRepo.EmergencyFilter) ([]*entities.EmergencyEvent, error) {
	query := `
		SELECT id, intersection_id, confidence_score, detection_method, 
		       threshold_passed, persistence_passed, radar_corr_passed, 
		       corridor_activated, response_time_ms, detected_at, resolved_at
		FROM emergency_events WHERE 1=1
	`
	var args []interface{}
	var conditions []string
	argID := 1

	if filter.IntersectionID != nil {
		conditions = append(conditions, fmt.Sprintf("intersection_id = $%d", argID))
		args = append(args, *filter.IntersectionID)
		argID++
	}
	if filter.DetectionMethod != nil {
		conditions = append(conditions, fmt.Sprintf("detection_method = $%d", argID))
		args = append(args, *filter.DetectionMethod)
		argID++
	}
	if filter.CorridorActivated != nil {
		conditions = append(conditions, fmt.Sprintf("corridor_activated = $%d", argID))
		args = append(args, *filter.CorridorActivated)
		argID++
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	query += fmt.Sprintf(" ORDER BY detected_at DESC LIMIT $%d OFFSET $%d", argID, argID+1)
	args = append(args, filter.Limit, filter.Offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*entities.EmergencyEvent
	for rows.Next() {
		var em entities.EmergencyEvent
		var responseTime sql.NullInt32
		var resolvedAt sql.NullTime

		err := rows.Scan(
			&em.ID, &em.IntersectionID, &em.ConfidenceScore, &em.DetectionMethod,
			&em.ThresholdPassed, &em.PersistencePassed, &em.RadarCorrPassed,
			&em.CorridorActivated, &responseTime, &em.DetectedAt, &resolvedAt,
		)
		if err != nil {
			return nil, err
		}

		if responseTime.Valid {
			val := int(responseTime.Int32)
			em.ResponseTimeMs = &val
		}
		if resolvedAt.Valid {
			em.ResolvedAt = &resolvedAt.Time
		}

		events = append(events, &em)
	}

	return events, nil
}

func (r *PostgresEmergencyRepository) UpdateCorridorStatus(ctx context.Context, id int64, activate bool, responseTimeMs *int) error {
	query := "UPDATE emergency_events SET corridor_activated = $1, response_time_ms = $2 WHERE id = $3"
	res, err := r.db.ExecContext(ctx, query, activate, responseTimeMs, id)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("evento de emergencia no encontrado")
	}
	return nil
}

func (r *PostgresEmergencyRepository) Resolve(ctx context.Context, id int64) error {
	query := "UPDATE emergency_events SET resolved_at = NOW() WHERE id = $1 AND resolved_at IS NULL"
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("evento no encontrado o ya resuelto")
	}
	return nil
}

func (r *PostgresEmergencyRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM emergency_events WHERE id = $1"
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("evento de emergencia no encontrado")
	}
	return nil
}
