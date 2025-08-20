package store

import (
	"context"

	"github.com/kasyap1234/portfolio/server/internal/db"
	"github.com/kasyap1234/portfolio/server/internal/models"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *models.User) (db.User, error)
	GetUserByEmail(ctx context.Context, email string) (db.User, error)
}

type userStore struct {
	db *db.Queries
}

func (u *userStore) CreateUser(ctx context.Context, user *models.User) (db.User, error) {
	args := db.CreateUserParams{
		ID:       int32(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return u.db.CreateUser(ctx, args)
}

func (u *userStore) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	return u.db.GetUserByEmail(ctx, email)
}
