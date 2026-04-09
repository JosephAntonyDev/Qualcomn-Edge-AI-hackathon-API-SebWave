package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/entities"
	domainRepo "github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/user/domain/repository"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) domainRepo.UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	query := `
		SELECT id, username, email, password_hash, role, full_name, is_active, last_login_at, created_at, updated_at, deleted_at
		FROM users
		WHERE id = $1 AND deleted_at IS NULL
	`
	return r.scanUser(r.db.QueryRowContext(ctx, query, id))
}

func (r *PostgresUserRepository) GetUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	query := `
		SELECT id, username, email, password_hash, role, full_name, is_active, last_login_at, created_at, updated_at, deleted_at
		FROM users
		WHERE username = $1 AND deleted_at IS NULL
	`
	return r.scanUser(r.db.QueryRowContext(ctx, query, username))
}

func (r *PostgresUserRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	query := `
		SELECT id, username, email, password_hash, role, full_name, is_active, last_login_at, created_at, updated_at, deleted_at
		FROM users
		WHERE email = $1 AND deleted_at IS NULL
	`
	return r.scanUser(r.db.QueryRowContext(ctx, query, email))
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, user *entities.User) error {
	query := `
		INSERT INTO users (username, email, password_hash, role, full_name, is_active)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRowContext(ctx, query,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.Role,
		user.FullName,
		user.IsActive,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresUserRepository) UpdateUser(ctx context.Context, user *entities.User) error {
	query := `
		UPDATE users 
		SET email = $1, role = $2, full_name = $3, is_active = $4, updated_at = NOW()
		WHERE id = $5 AND deleted_at IS NULL
		RETURNING updated_at
	`

	err := r.db.QueryRowContext(ctx, query,
		user.Email,
		user.Role,
		user.FullName,
		user.IsActive,
		user.ID,
	).Scan(&user.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("user not found")
		}
		return err
	}
	return nil
}

func (r *PostgresUserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	query := `
		UPDATE users 
		SET deleted_at = NOW(), is_active = FALSE 
		WHERE id = $1 AND deleted_at IS NULL
	`

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *PostgresUserRepository) GetMyProfile(ctx context.Context, userID uuid.UUID) (*entities.User, error) {
	return r.GetUserByID(ctx, userID)
}

func (r *PostgresUserRepository) scanUser(row *sql.Row) (*entities.User, error) {
	var user entities.User
	var deletedAt sql.NullTime
	var lastLoginAt sql.NullTime
	var fullName sql.NullString

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&fullName,
		&user.IsActive,
		&lastLoginAt,
		&user.CreatedAt,
		&user.UpdatedAt,
		&deletedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if fullName.Valid {
		user.FullName = fullName.String
	}
	if lastLoginAt.Valid {
		user.LastLoginAt = &lastLoginAt.Time
	}
	if deletedAt.Valid {
		user.DeletedAt = &deletedAt.Time
	}

	return &user, nil
}
