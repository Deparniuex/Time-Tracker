package service

import (
	"example.com/tracker/internal/entity"
	"example.com/tracker/pkg/util"
)

func (m *Manager) CreateUser(user *entity.User) error {
	user, err := m.ExternalApi.GetUsersInfo(user)
	if err != nil {
		return err
	}
	return m.Repository.CreateUser(user)
}

func (m *Manager) GetUsers(pagination *util.Pagination, filters map[string]string) ([]*entity.User, *util.Metadata, error) {
	return m.Repository.GetUsers(pagination, filters)
}

func (m *Manager) DeleteUser(userID int64) error {
	return m.Repository.DeleteUser(userID)
}

func (m *Manager) UpdateUser(user *entity.User) error {
	return m.Repository.UpdateUser(user)
}
