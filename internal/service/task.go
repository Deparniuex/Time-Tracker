package service

import (
	"time"

	"example.com/tracker/internal/entity"
)

func (m *Manager) StartTimer(task *entity.Task) error {
	return m.Repository.StartTimer(task)
}

func (m *Manager) EndTimer(taskID int64) error {
	return m.Repository.EndTimer(taskID)
}

func (m *Manager) GetWorkLoads(userID int64, startDate, endDate time.Time) ([]*entity.WorkLoad, error) {
	return m.Repository.GetWorkLoads(userID, startDate, endDate)
}
