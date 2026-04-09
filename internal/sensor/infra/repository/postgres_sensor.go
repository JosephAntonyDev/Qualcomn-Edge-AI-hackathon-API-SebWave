package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/domain/entities"
	domainRepo "github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/sensor/domain/repository"
	"github.com/google/uuid"
)

type PostgresSensorRepository struct {
	db *sql.DB
}

func NewPostgresSensorRepository(db *sql.DB) domainRepo.SensorRepository {
	return &PostgresSensorRepository{db: db}
}

func (r *PostgresSensorRepository) RegisterSensor(ctx context.Context, s *entities.Sensor) error {
	query := `
		INSERT INTO node_sensors (
			intersection_id, sensor_type, lane_direction, connection_type, pin_assignment, is_active
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, failure_count
	`
	err := r.db.QueryRowContext(ctx, query,
		s.IntersectionID, s.SensorType, s.LaneDirection, s.ConnectionType, s.PinAssignment, s.IsActive,
	).Scan(&s.ID, &s.CreatedAt, &s.FailureCount)

	return err
}

func (r *PostgresSensorRepository) GetSensor(ctx context.Context, id uuid.UUID) (*entities.Sensor, error) {
	query := `
		SELECT id, intersection_id, sensor_type, lane_direction, connection_type, pin_assignment, is_active, failure_count, last_reading_at, created_at
		FROM node_sensors WHERE id = $1
	`
	return r.scanSensor(r.db.QueryRowContext(ctx, query, id))
}

func (r *PostgresSensorRepository) ListSensorsByIntersection(ctx context.Context, intersectionID uuid.UUID) ([]*entities.Sensor, error) {
	query := `
		SELECT id, intersection_id, sensor_type, lane_direction, connection_type, pin_assignment, is_active, failure_count, last_reading_at, created_at
		FROM node_sensors WHERE intersection_id = $1
	`
	rows, err := r.db.QueryContext(ctx, query, intersectionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sensors []*entities.Sensor
	for rows.Next() {
		s, err := r.scanSensorRow(rows)
		if err != nil {
			return nil, err
		}
		sensors = append(sensors, s)
	}
	return sensors, nil
}

func (r *PostgresSensorRepository) UpdateSensor(ctx context.Context, s *entities.Sensor) error {
	query := `
		UPDATE node_sensors SET
			lane_direction = $1, connection_type = $2, pin_assignment = $3, is_active = $4, failure_count = $5
		WHERE id = $6
	`
	res, err := r.db.ExecContext(ctx, query, s.LaneDirection, s.ConnectionType, s.PinAssignment, s.IsActive, s.FailureCount, s.ID)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("sensor no encontrado")
	}
	return nil
}

func (r *PostgresSensorRepository) DeleteSensor(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM node_sensors WHERE id = $1"
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("sensor no encontrado")
	}
	return nil
}

func (r *PostgresSensorRepository) RecordReading(ctx context.Context, read *entities.SensorReading) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Insertar lectura
	queryRead := `
		INSERT INTO sensor_readings (
			sensor_id, intersection_id, is_occupied, distance_mm, energy_pct, confidence, raw_data
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at
	`

	rawJson := []byte(nil)
	if read.RawData != nil {
		rawJson = read.RawData
	}

	err = tx.QueryRowContext(ctx, queryRead,
		read.SensorID, read.IntersectionID, read.IsOccupied, read.DistanceMM, read.EnergyPct, read.Confidence, rawJson,
	).Scan(&read.ID, &read.CreatedAt)
	if err != nil {
		return err
	}

	// Update last_reading_at for sensor
	queryUpdateSensor := `UPDATE node_sensors SET last_reading_at = NOW() WHERE id = $1`
	_, err = tx.ExecContext(ctx, queryUpdateSensor, read.SensorID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PostgresSensorRepository) GetReadings(ctx context.Context, sensorID uuid.UUID, limit int) ([]*entities.SensorReading, error) {
	query := `
		SELECT id, sensor_id, intersection_id, is_occupied, distance_mm, energy_pct, confidence, raw_data, created_at
		FROM sensor_readings
		WHERE sensor_id = $1
		ORDER BY created_at DESC LIMIT $2
	`
	rows, err := r.db.QueryContext(ctx, query, sensorID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var readings []*entities.SensorReading
	for rows.Next() {
		var rd entities.SensorReading
		var isOccupied sql.NullBool
		var distMM, enPct sql.NullInt32
		var conf sql.NullFloat64
		var rawData []byte

		if err := rows.Scan(
			&rd.ID, &rd.SensorID, &rd.IntersectionID, &isOccupied, &distMM, &enPct, &conf, &rawData, &rd.CreatedAt,
		); err != nil {
			return nil, err
		}

		if isOccupied.Valid {
			rd.IsOccupied = &isOccupied.Bool
		}
		if distMM.Valid {
			v := int(distMM.Int32)
			rd.DistanceMM = &v
		}
		if enPct.Valid {
			v := int(enPct.Int32)
			rd.EnergyPct = &v
		}
		if conf.Valid {
			rd.Confidence = &conf.Float64
		}
		if len(rawData) > 0 {
			rd.RawData = json.RawMessage(rawData)
		}

		readings = append(readings, &rd)
	}

	return readings, nil
}

type rowScanner interface {
	Scan(dest ...interface{}) error
}

func (r *PostgresSensorRepository) scanSensor(row rowScanner) (*entities.Sensor, error) {
	var s entities.Sensor
	var lDir, cType, pinAssign sql.NullString
	var lastReading sql.NullTime

	err := row.Scan(
		&s.ID, &s.IntersectionID, &s.SensorType, &lDir, &cType, &pinAssign, &s.IsActive, &s.FailureCount, &lastReading, &s.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("sensor no encontrado")
		}
		return nil, err
	}

	if lDir.Valid {
		s.LaneDirection = &lDir.String
	}
	if cType.Valid {
		s.ConnectionType = &cType.String
	}
	if pinAssign.Valid {
		s.PinAssignment = &pinAssign.String
	}
	if lastReading.Valid {
		s.LastReadingAt = &lastReading.Time
	}

	return &s, nil
}

func (r *PostgresSensorRepository) scanSensorRow(row *sql.Rows) (*entities.Sensor, error) {
	return r.scanSensor(row)
}
