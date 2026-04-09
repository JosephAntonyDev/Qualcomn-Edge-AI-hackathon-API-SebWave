package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/domain/entities"
	domainRepo "github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/intersection/domain/repository"
	"github.com/google/uuid"
)

type PostgresIntersectionRepository struct {
	db *sql.DB
}

func NewPostgresIntersectionRepository(db *sql.DB) domainRepo.IntersectionRepository {
	return &PostgresIntersectionRepository{db: db}
}

func (r *PostgresIntersectionRepository) Create(ctx context.Context, is *entities.Intersection) error {
	query := `
		INSERT INTO intersections (
			serial_number, name, description, latitude, longitude, node_id,
			max_congestion_pct, min_green_time_s, max_green_time_s,
			default_green_s, default_red_s, yellow_time_s,
			status, operation_mode, created_by
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
		RETURNING id, created_at, updated_at
	`
	err := r.db.QueryRowContext(ctx, query,
		is.SerialNumber, is.Name, is.Description, is.Latitude, is.Longitude, is.NodeID,
		is.MaxCongestionPct, is.MinGreenTimeS, is.MaxGreenTimeS,
		is.DefaultGreenS, is.DefaultRedS, is.YellowTimeS,
		is.Status, is.OperationMode, is.CreatedBy,
	).Scan(&is.ID, &is.CreatedAt, &is.UpdatedAt)

	return err
}

func (r *PostgresIntersectionRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Intersection, error) {
	query := `SELECT id, serial_number, name, description, latitude, longitude, node_id, max_congestion_pct, min_green_time_s, max_green_time_s, default_green_s, default_red_s, yellow_time_s, status, operation_mode, current_phase, current_density_ns, current_density_eo, firmware_version, last_heartbeat_at, created_by, created_at, updated_at, deleted_at FROM intersections WHERE id = $1 AND deleted_at IS NULL`
	return r.scanIntersection(r.db.QueryRowContext(ctx, query, id))
}

func (r *PostgresIntersectionRepository) GetBySerialNumber(ctx context.Context, serialNumber string) (*entities.Intersection, error) {
	query := `SELECT id, serial_number, name, description, latitude, longitude, node_id, max_congestion_pct, min_green_time_s, max_green_time_s, default_green_s, default_red_s, yellow_time_s, status, operation_mode, current_phase, current_density_ns, current_density_eo, firmware_version, last_heartbeat_at, created_by, created_at, updated_at, deleted_at FROM intersections WHERE serial_number = $1 AND deleted_at IS NULL`
	return r.scanIntersection(r.db.QueryRowContext(ctx, query, serialNumber))
}

func (r *PostgresIntersectionRepository) GetAll(ctx context.Context) ([]*entities.Intersection, error) {
	query := `SELECT id, serial_number, name, description, latitude, longitude, node_id, max_congestion_pct, min_green_time_s, max_green_time_s, default_green_s, default_red_s, yellow_time_s, status, operation_mode, current_phase, current_density_ns, current_density_eo, firmware_version, last_heartbeat_at, created_by, created_at, updated_at, deleted_at FROM intersections WHERE deleted_at IS NULL`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ints []*entities.Intersection
	for rows.Next() {
		var is entities.Intersection
		var desc, nID, cPhase, fv sql.NullString
		var densNS, densEO sql.NullFloat64
		var lhAt, delAt sql.NullTime
		var cBy uuid.NullUUID

		if err := rows.Scan(
			&is.ID, &is.SerialNumber, &is.Name, &desc, &is.Latitude, &is.Longitude, &nID,
			&is.MaxCongestionPct, &is.MinGreenTimeS, &is.MaxGreenTimeS, &is.DefaultGreenS, &is.DefaultRedS, &is.YellowTimeS,
			&is.Status, &is.OperationMode, &cPhase, &densNS, &densEO, &fv, &lhAt, &cBy, &is.CreatedAt, &is.UpdatedAt, &delAt,
		); err != nil {
			return nil, err
		}

		if desc.Valid {
			is.Description = &desc.String
		}
		if nID.Valid {
			is.NodeID = &nID.String
		}
		if cPhase.Valid {
			is.CurrentPhase = &cPhase.String
		}
		if densNS.Valid {
			is.CurrentDensityNS = &densNS.Float64
		}
		if densEO.Valid {
			is.CurrentDensityEO = &densEO.Float64
		}
		if fv.Valid {
			is.FirmwareVersion = &fv.String
		}
		if lhAt.Valid {
			is.LastHeartbeatAt = &lhAt.Time
		}
		if cBy.Valid {
			is.CreatedBy = &cBy.UUID
		}
		if delAt.Valid {
			is.DeletedAt = &delAt.Time
		}

		ints = append(ints, &is)
	}
	return ints, nil
}

func (r *PostgresIntersectionRepository) Update(ctx context.Context, is *entities.Intersection) error {
	query := `
		UPDATE intersections SET
			name = $1, description = $2, latitude = $3, longitude = $4, node_id = $5,
			max_congestion_pct = $6, min_green_time_s = $7, max_green_time_s = $8,
			default_green_s = $9, default_red_s = $10, yellow_time_s = $11, updated_at = NOW()
		WHERE id = $12 AND deleted_at IS NULL
		RETURNING updated_at
	`
	err := r.db.QueryRowContext(ctx, query,
		is.Name, is.Description, is.Latitude, is.Longitude, is.NodeID,
		is.MaxCongestionPct, is.MinGreenTimeS, is.MaxGreenTimeS,
		is.DefaultGreenS, is.DefaultRedS, is.YellowTimeS, is.ID,
	).Scan(&is.UpdatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return errors.New("intersección no encontrada")
	}
	return err
}

func (r *PostgresIntersectionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE intersections SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("intersección no encontrada")
	}
	return nil
}

func (r *PostgresIntersectionRepository) UpdateState(ctx context.Context, id uuid.UUID, status entities.NodeStatus, mode entities.OperationMode, phase *string, densityNS, densityEO *float64) error {
	query := `
		UPDATE intersections SET
			status = $1, operation_mode = $2, current_phase = $3,
			current_density_ns = $4, current_density_eo = $5, updated_at = NOW()
		WHERE id = $6 AND deleted_at IS NULL
	`
	res, err := r.db.ExecContext(ctx, query, status, mode, phase, densityNS, densityEO, id)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("intersección no encontrada")
	}
	return nil
}

func (r *PostgresIntersectionRepository) Heartbeat(ctx context.Context, serialNumber string, firmwareVersion *string) error {
	query := `
		UPDATE intersections SET
			last_heartbeat_at = NOW(), firmware_version = COALESCE($1, firmware_version), status = 'connected'
		WHERE serial_number = $2 AND deleted_at IS NULL
	`
	res, err := r.db.ExecContext(ctx, query, firmwareVersion, serialNumber)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("intersección no encontrada")
	}
	return nil
}

func (r *PostgresIntersectionRepository) scanIntersection(row *sql.Row) (*entities.Intersection, error) {
	var is entities.Intersection
	var desc, nID, cPhase, fv sql.NullString
	var densNS, densEO sql.NullFloat64
	var lhAt, delAt sql.NullTime
	var cBy uuid.NullUUID

	err := row.Scan(
		&is.ID, &is.SerialNumber, &is.Name, &desc, &is.Latitude, &is.Longitude, &nID,
		&is.MaxCongestionPct, &is.MinGreenTimeS, &is.MaxGreenTimeS, &is.DefaultGreenS, &is.DefaultRedS, &is.YellowTimeS,
		&is.Status, &is.OperationMode, &cPhase, &densNS, &densEO, &fv, &lhAt, &cBy, &is.CreatedAt, &is.UpdatedAt, &delAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("intersección no encontrada")
		}
		return nil, err
	}

	if desc.Valid {
		is.Description = &desc.String
	}
	if nID.Valid {
		is.NodeID = &nID.String
	}
	if cPhase.Valid {
		is.CurrentPhase = &cPhase.String
	}
	if densNS.Valid {
		is.CurrentDensityNS = &densNS.Float64
	}
	if densEO.Valid {
		is.CurrentDensityEO = &densEO.Float64
	}
	if fv.Valid {
		is.FirmwareVersion = &fv.String
	}
	if lhAt.Valid {
		is.LastHeartbeatAt = &lhAt.Time
	}
	if cBy.Valid {
		is.CreatedBy = &cBy.UUID
	}
	if delAt.Valid {
		is.DeletedAt = &delAt.Time
	}

	return &is, nil
}
