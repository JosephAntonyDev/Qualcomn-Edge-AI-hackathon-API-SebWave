package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string
type UserStatus string

const (
	RoleAdmin UserRole = "admin"
	RoleOperator UserRole = "operator"
	RoleViewer UserRole = "viewer"
)

type User struct {
	ID		     uuid.UUID `json:"id"`
	Username	 string    `json:"username"`
	Email		 string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role		 UserRole  `json:"role"`
	FullName     string    `json:"full_name,omitempty"`
	IsActive     bool      `json:"is_active"`
	LastLoginAt  *time.Time `json:"last_login_at,omitempty"`
    CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}