package store

import (
	"context"

	"github.com/kasyap1234/portfolio/server/internal/models"
)

type AuthStore interface {
	CreateUser(ctx context.Context, user *models.User) error
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	
}

type authStore struct {
	
}
