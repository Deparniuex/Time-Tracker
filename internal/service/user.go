package service

import (
	"context"

	"example.com/tracker/internal/entity"
)

func (m *Manager) GetUsers(ctx context.Context) ([]*entity.User, error) {
	return m.Repository.GetUsers(ctx)
}
