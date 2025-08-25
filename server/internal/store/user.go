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
	q *db.Queries
}

func (u *userStore) CreateUser(ctx context.Context, user db.User) (*models.User, error) {
	args := db.CreateUserParams{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	user, err := u.q.CreateUser(ctx, args)
	if err != nil {
		return &models.User{}, err
	}
	modelUser := &models.User{
		ID:       int(user.ID.ID()),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return modelUser, err

}

func (u *userStore) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	dbUser, err := u.q.GetUserByEmail(ctx, email)
	if err != nil {
		return &models.User{}, err
	}
	modelUser := &models.User{
		ID:       int(dbUser.ID.ID()),
		Name:     dbUser.Name,
		Email:    dbUser.Email,
		Password: dbUser.Password,
	}
	return modelUser, err

}
