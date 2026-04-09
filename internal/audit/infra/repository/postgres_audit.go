package repository

import (
	"context"
	"database/sql"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/audit/domain/entities"
)

type PostgresAuditRepository struct {
	DB *sql.DB
}

func NewPostgresAuditRepository(db *sql.DB) *PostgresAuditRepository {
	return &PostgresAuditRepository{DB: db}
}

func (r *PostgresAuditRepository) CreateLog(ctx context.Context, log *entities.AuditLog) error {
	query := `
		INSERT INTO audit_log (user_id, action, target_resource, target_id, details, timestamp)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.DB.ExecContext(ctx, query, log.UserID, log.Action, log.TargetResource, log.TargetID, log.Details, log.Timestamp)
	return err
}

func (r *PostgresAuditRepository) GetLogs(ctx context.Context, limit int, offset int) ([]entities.AuditLog, error) {
	query := `
		SELECT id, user_id, action, target_resource, target_id, details, timestamp
		FROM audit_log
		ORDER BY timestamp DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := r.DB.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []entities.AuditLog
	for rows.Next() {
		var log entities.AuditLog
		if err := rows.Scan(
			&log.ID,
			&log.UserID,
			&log.Action,
			&log.TargetResource,
			&log.TargetID,
			&log.Details,
			&log.Timestamp,
		); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}
