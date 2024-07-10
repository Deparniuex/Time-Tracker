package service

import (
	"context"

	"example.com/tracker/internal/entity"
)

type Service interface {
	GetUsers(ctx context.Context) ([]*entity.User, error)
}
