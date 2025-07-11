package storage

import (
	"context"
	"database/sql"
	"bytes"
	"github.com/pas-platform/identity-service/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(db *sql.DB) *PostgresStore {
	return &PostgresStore{db: db}
}

func (s *PostgresStore) CreateTenantAndUser(ctx context.Context, req domain.RegistrationRequest) (string, string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return "", "", err
	}
	defer tx.Rollback()

	var tenantID string
	err = tx.QueryRowContext(ctx, "INSERT INTO tenants (name) VALUES ($1) RETURNING id", req.TenantName).Scan(&tenantID)
	if err != nil {
		return "", "", err
	}

	var userID string
	err = tx.QueryRowContext(ctx,
		`INSERT INTO users (tenant_id, email, password_hash, first_name, last_name, role)
         VALUES ($1, $2, $3, $4, $5, 'admin') RETURNING id`,
		tenantID, req.Email, string(passwordHash), req.FirstName, req.LastName,
	).Scan(&userID)
	if err != nil {
		// Check for unique constraint violation on email
		if bytes.Contains([]byte(err.Error()), []byte("users_email_key")) {
			return "", "", &ErrDuplicateEmail{Email: req.Email}
		}
		return "", "", err
	}

	return userID, tenantID, tx.Commit()
}

// Custom error for handling specific database constraints
type ErrDuplicateEmail struct {
	Email string
}

func (e *ErrDuplicateEmail) Error() string {
	return "user with this email already exists"
}