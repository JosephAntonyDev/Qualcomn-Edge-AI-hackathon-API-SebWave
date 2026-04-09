package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/traffic_cycle/domain/entities"
	domainRepo "github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/traffic_cycle/domain/repository"
	"github.com/google/uuid"
)

type PostgresTrafficCycleRepository struct {
	db *sql.DB
}

func NewPostgresTrafficCycleRepository(db *sql.DB) domainRepo.TrafficCycleRepository {
	return &PostgresTrafficCycleRepository{db: db}
}

func (r *PostgresTrafficCycleRepository) RecordCycle(ctx context.Context, c *entities.TrafficCycle) error {
	query := `
		INSERT INTO traffic_cycles (
			intersection_id, operation_mode, green_ns_ms, green_eo_ms, yellow_ms,
			radar_occupied, radar_distance_cm, radar_energy, modulino_occupied, modulino_distance_mm,
			camera_ns_occupied, camera_eo_occupied, wait_time_saved_ms, co2_saved_g, model_confidence, 
			started_at, ended_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, COALESCE($16, NOW()), $17
		) RETURNING id, started_at
	`
	err := r.db.QueryRowContext(ctx, query,
		c.IntersectionID, c.OperationMode, c.GreenNSMs, c.GreenEOMs, c.YellowMs,
		c.RadarOccupied, c.RadarDistanceCm, c.RadarEnergy, c.ModulinoOccupied, c.ModulinoDistanceMm,
		c.CameraNSOccupied, c.CameraEOOccupied, c.WaitTimeSavedMs, c.CO2SavedG, c.ModelConfidence,
		c.StartedAt, c.EndedAt,
	).Scan(&c.ID, &c.StartedAt)

	return err
}

func (r *PostgresTrafficCycleRepository) GetCycleByID(ctx context.Context, id int64) (*entities.TrafficCycle, error) {
	query := `
		SELECT id, intersection_id, operation_mode, green_ns_ms, green_eo_ms, yellow_ms,
		       radar_occupied, radar_distance_cm, radar_energy, modulino_occupied, modulino_distance_mm,
			   camera_ns_occupied, camera_eo_occupied, wait_time_saved_ms, co2_saved_g, model_confidence,
			   started_at, ended_at
		FROM traffic_cycles WHERE id = $1
	`
	return r.scanCycle(r.db.QueryRowContext(ctx, query, id))
}

func (r *PostgresTrafficCycleRepository) ListCyclesByIntersection(ctx context.Context, intID uuid.UUID, limit, offset int) ([]*entities.TrafficCycle, error) {
	query := `
		SELECT id, intersection_id, operation_mode, green_ns_ms, green_eo_ms, yellow_ms,
		       radar_occupied, radar_distance_cm, radar_energy, modulino_occupied, modulino_distance_mm,
			   camera_ns_occupied, camera_eo_occupied, wait_time_saved_ms, co2_saved_g, model_confidence,
			   started_at, ended_at
		FROM traffic_cycles 
		WHERE intersection_id = $1
		ORDER BY started_at DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.QueryContext(ctx, query, intID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cycles []*entities.TrafficCycle
	for rows.Next() {
		c, err := r.scanCycleRow(rows)
		if err != nil {
			return nil, err
		}
		cycles = append(cycles, c)
	}

	return cycles, nil
}

func (r *PostgresTrafficCycleRepository) DeleteCycle(ctx context.Context, id int64) error {
	query := "DELETE FROM traffic_cycles WHERE id = $1"
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("ciclo no encontrado")
	}
	return nil
}

type rowScanner interface {
	Scan(dest ...interface{}) error
}

func (r *PostgresTrafficCycleRepository) scanCycle(row rowScanner) (*entities.TrafficCycle, error) {
	var c entities.TrafficCycle
	var rOcc, mOcc, cNsOcc, cEoOcc sql.NullBool
	var rDist, rEn, mDist, wTs sql.NullInt32
	var co2, mConf sql.NullFloat64
	var endAt sql.NullTime

	err := row.Scan(
		&c.ID, &c.IntersectionID, &c.OperationMode, &c.GreenNSMs, &c.GreenEOMs, &c.YellowMs,
		&rOcc, &rDist, &rEn, &mOcc, &mDist, &cNsOcc, &cEoOcc, &wTs, &co2, &mConf,
		&c.StartedAt, &endAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("ciclo de semáforo no encontrado")
		}
		return nil, err
	}

	if rOcc.Valid {
		b := rOcc.Bool
		c.RadarOccupied = &b
	}
	if mOcc.Valid {
		b := mOcc.Bool
		c.ModulinoOccupied = &b
	}
	if cNsOcc.Valid {
		b := cNsOcc.Bool
		c.CameraNSOccupied = &b
	}
	if cEoOcc.Valid {
		b := cEoOcc.Bool
		c.CameraEOOccupied = &b
	}

	if rDist.Valid {
		i := int(rDist.Int32)
		c.RadarDistanceCm = &i
	}
	if rEn.Valid {
		i := int(rEn.Int32)
		c.RadarEnergy = &i
	}
	if mDist.Valid {
		i := int(mDist.Int32)
		c.ModulinoDistanceMm = &i
	}
	if wTs.Valid {
		i := int(wTs.Int32)
		c.WaitTimeSavedMs = &i
	}

	if co2.Valid {
		c.CO2SavedG = &co2.Float64
	}
	if mConf.Valid {
		c.ModelConfidence = &mConf.Float64
	}
	if endAt.Valid {
		c.EndedAt = &endAt.Time
	}

	return &c, nil
}

func (r *PostgresTrafficCycleRepository) scanCycleRow(row *sql.Rows) (*entities.TrafficCycle, error) {
	return r.scanCycle(row)
}
