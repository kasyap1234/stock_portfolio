package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/kasyap1234/portfolio/server/internal/db"
	"github.com/kasyap1234/portfolio/server/internal/models"
)

type AuthStore interface {
	RegisterUser(ctx context.Context, user *db.User) (*models.User, error)
	LoginUser(ctx context.Context, user *db.User) (*models.User, error)
}

type authStore struct {
	q *db.Queries
	userStore UserStore
}



func(a*authStore)RegisterUser(ctx context.Context,user*db.User)(*models.User,error){
    modelUser :=&models.User{
ID : uuid.New(),

	}
	user,err :=a.userStore.CreateUser(ctx,modelUser)
}