package repository

import (
	"context"

	"example.com/tracker/internal/entity"
)

type Repository interface {
	GetUsers(ctx context.Context) ([]*entity.User, error)
}
