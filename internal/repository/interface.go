package repository

import (
	"context"

	"example.com/tracker/internal/entity"
)

type Repository interface {
	CreateUser(user *entity.User) error
	GetUsers(ctx context.Context) ([]*entity.User, error)
	DeleteUser(ctx context.Context, userID int64) error
	UpdateUser(user *entity.User) error
}

type ExternalAPI interface {
	GetUsersInfo(user *entity.User) (*entity.User, error)
}
