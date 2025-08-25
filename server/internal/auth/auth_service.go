package auth

import (
	"context"

	"github.com/kasyap1234/portfolio/server/internal/models"
)

type AuthService interface {
	RegisterUser(ctx context.Context, user *models.User) (*models.User, error)
	VerifyEmail(ctx context.Context, email string) (bool, error)
	LoginUser(ctx context.Context, user *models.User) (*models.User, error)
}
