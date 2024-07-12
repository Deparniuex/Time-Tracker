package service

import (
	"context"

	"example.com/tracker/internal/entity"
)

func (m *Manager) CreateUser(user *entity.User) error {
	user, err := m.ExternalApi.GetUsersInfo(user)
	if err != nil {
		return err
	}
	return m.Repository.CreateUser(user)
}

func (m *Manager) GetUsers(ctx context.Context) ([]*entity.User, error) {
	return m.Repository.GetUsers(ctx)
}

func (m *Manager) DeleteUser(ctx context.Context, userID int64) error {
	return m.Repository.DeleteUser(ctx, userID)
}

func (m *Manager) UpdateUser(user *entity.User) error {
	return m.Repository.UpdateUser(user)
}
