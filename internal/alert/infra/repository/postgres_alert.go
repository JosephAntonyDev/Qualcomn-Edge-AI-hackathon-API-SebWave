package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/domain/entities"
	domainRepo "github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/alert/domain/repository"
	"github.com/google/uuid"
)

type PostgresAlertRepository struct {
	db *sql.DB
}

func NewPostgresAlertRepository(db *sql.DB) domainRepo.AlertRepository {
	return &PostgresAlertRepository{db: db}
}

func (r *PostgresAlertRepository) Create(ctx context.Context, al *entities.Alert) error {
	query := `
		INSERT INTO alerts (
			intersection_id, type, severity, status, title, description, metadata
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at
	`

	rawJson := []byte(nil)
	if al.Metadata != nil {
		rawJson = al.Metadata
	}

	err := r.db.QueryRowContext(ctx, query,
		al.IntersectionID, al.Type, al.Severity, al.Status, al.Title, al.Description, rawJson,
	).Scan(&al.ID, &al.CreatedAt, &al.UpdatedAt)

	return err
}

func (r *PostgresAlertRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Alert, error) {
	query := `
		SELECT id, intersection_id, type, severity, status, title, description, metadata, resolved_by, resolved_at, resolution_note, created_at, updated_at
		FROM alerts WHERE id = $1
	`
	return r.scanAlert(r.db.QueryRowContext(ctx, query, id))
}

func (r *PostgresAlertRepository) List(ctx context.Context, filter domainRepo.AlertFilter) ([]*entities.Alert, error) {
	query := `
		SELECT id, intersection_id, type, severity, status, title, description, metadata, resolved_by, resolved_at, resolution_note, created_at, updated_at
		FROM alerts WHERE 1=1
	`
	var args []interface{}
	var conditions []string
	argID := 1

	if filter.IntersectionID != nil {
		conditions = append(conditions, fmt.Sprintf("intersection_id = $%d", argID))
		args = append(args, *filter.IntersectionID)
		argID++
	}
	if filter.Type != nil {
		conditions = append(conditions, fmt.Sprintf("type = $%d", argID))
		args = append(args, *filter.Type)
		argID++
	}
	if filter.Severity != nil {
		conditions = append(conditions, fmt.Sprintf("severity = $%d", argID))
		args = append(args, *filter.Severity)
		argID++
	}
	if filter.Status != nil {
		conditions = append(conditions, fmt.Sprintf("status = $%d", argID))
		args = append(args, *filter.Status)
		argID++
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", argID, argID+1)
	args = append(args, filter.Limit, filter.Offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alerts []*entities.Alert
	for rows.Next() {
		al, err := r.scanAlertRow(rows)
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, al)
	}

	return alerts, nil
}

func (r *PostgresAlertRepository) Update(ctx context.Context, al *entities.Alert) error {
	query := `
		UPDATE alerts SET
			type = $1, severity = $2, title = $3, description = $4, metadata = $5, updated_at = NOW()
		WHERE id = $6
		RETURNING updated_at
	`
	rawJson := []byte(nil)
	if al.Metadata != nil {
		rawJson = al.Metadata
	}

	err := r.db.QueryRowContext(ctx, query,
		al.Type, al.Severity, al.Title, al.Description, rawJson, al.ID,
	).Scan(&al.UpdatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return errors.New("alerta no encontrada")
	}
	return err
}

func (r *PostgresAlertRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM alerts WHERE id = $1"
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("alerta no encontrada")
	}
	return nil
}

func (r *PostgresAlertRepository) ChangeStatus(ctx context.Context, id uuid.UUID, status entities.AlertStatus, resolvedBy *uuid.UUID, note *string) error {
	var query string
	var err error
	var res sql.Result

	if status == entities.StatusResolved {
		query = `UPDATE alerts SET status = $1, resolved_by = $2, resolution_note = $3, resolved_at = NOW(), updated_at = NOW() WHERE id = $4`
		res, err = r.db.ExecContext(ctx, query, status, resolvedBy, note, id)
	} else if status == entities.StatusAcknowledged {
		query = `UPDATE alerts SET status = $1, updated_at = NOW() WHERE id = $2`
		res, err = r.db.ExecContext(ctx, query, status, id)
	} else {
		// Active (Re-activate)
		query = `UPDATE alerts SET status = $1, resolved_by = NULL, resolution_note = NULL, resolved_at = NULL, updated_at = NOW() WHERE id = $2`
		res, err = r.db.ExecContext(ctx, query, status, id)
	}

	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("alerta no encontrada")
	}
	return nil
}

type rowScanner interface {
	Scan(dest ...interface{}) error
}

func (r *PostgresAlertRepository) scanAlert(row rowScanner) (*entities.Alert, error) {
	var al entities.Alert
	var intID, resBy uuid.NullUUID
	var desc, resNote sql.NullString
	var resAt sql.NullTime
	var meta []byte

	err := row.Scan(
		&al.ID, &intID, &al.Type, &al.Severity, &al.Status, &al.Title, &desc, &meta, &resBy, &resAt, &resNote, &al.CreatedAt, &al.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("alerta no encontrada")
		}
		return nil, err
	}

	if intID.Valid {
		al.IntersectionID = &intID.UUID
	}
	if resBy.Valid {
		al.ResolvedBy = &resBy.UUID
	}
	if desc.Valid {
		al.Description = &desc.String
	}
	if resNote.Valid {
		al.ResolutionNote = &resNote.String
	}
	if resAt.Valid {
		al.ResolvedAt = &resAt.Time
	}
	if len(meta) > 0 {
		al.Metadata = json.RawMessage(meta)
	}

	return &al, nil
}

func (r *PostgresAlertRepository) scanAlertRow(row *sql.Rows) (*entities.Alert, error) {
	return r.scanAlert(row)
}
