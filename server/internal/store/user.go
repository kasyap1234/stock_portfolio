package store

import (
	"context"

	"github.com/kasyap1234/portfolio/server/internal/db"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *db.User) (db.User, error)
	GetUserByUsername(ctx context.Context, username string) (db.User, error)
}

type userStore struct {
	db *db.Queries
}

func (u *userStore) CreateUser(ctx context.Context, user *db.User) (db.User, error) {
	args := db.CreateUserParams{
		ID:       int32(user.ID),
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
	}
	return u.db.CreateUser(ctx, args)
}

func (u *userStore) FindUserByUsername(ctx context.Context, username string) (db.User, error) {
	return u.db.FindUserByUsername(ctx, username)
}
