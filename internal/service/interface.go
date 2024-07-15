package service

import (
	"time"

	"example.com/tracker/internal/entity"
	"example.com/tracker/pkg/util"
)

type Service interface {
	CreateUser(user *entity.User) error
	GetUsers(pagination *util.Pagination, filters map[string]string) ([]*entity.User, *util.Metadata, error)
	DeleteUser(userID int64) error
	UpdateUser(user *entity.User) error

	StartTimer(task *entity.Task) error
	EndTimer(taskID int64) error
	GetWorkLoads(userID int64, startDate, endDate time.Time) ([]*entity.WorkLoad, error)
}
