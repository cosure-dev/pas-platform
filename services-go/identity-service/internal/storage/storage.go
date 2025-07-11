package storage

import (
	"context"
	"github.com/pas-platform/identity-service/internal/domain"
)

type Storage interface {
	CreateTenantAndUser(ctx context.Context, req domain.RegistrationRequest) (userID, tenantID string, err error)
}