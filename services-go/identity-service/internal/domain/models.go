package domain

import "time"

type Tenant struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID           string    `json:"id"`
	TenantID     string    `json:"tenant_id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` 
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type RegistrationRequest struct {
	TenantName string `json:"tenantName" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=8"`
	FirstName  string `json:"firstName" binding:"required"`
	LastName   string `json:"lastName" binding:"required"`
}