package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/kasyap1234/portfolio/server/internal/db"
	"github.com/kasyap1234/portfolio/server/internal/models"
	"github.com/kasyap1234/portfolio/server/pkg/security"
)

type AuthStore interface {
	RegisterUser(ctx context.Context, user *db.User) (*models.User, error)
	LoginUser(ctx context.Context, user *db.User) (*models.User, error)
}

type authStore struct {
	q         *db.Queries
	userStore UserStore
}

func (a *authStore) RegisterUser(ctx context.Context, user *models.User) (*models.User, error) {
	if user.Email == "" || user.Password == "" {
		return nil, errors.New("user email or password cant be empty")
	}
	_, err := a.userStore.GetUserByEmail(ctx, user.Email)
	if err == nil {
		return nil, errors.New("user already exists")
	}
	if err != sql.ErrNoRows {
		return nil, err
	}
	hashedPassword, err := security.HashPassword(user.Password)
	user.Password = hashedPassword

	createdUser, err := a.userStore.CreateUser(ctx, user)
	if err != nil {
		return &models.User{}, err
	}
	return createdUser, nil
}
