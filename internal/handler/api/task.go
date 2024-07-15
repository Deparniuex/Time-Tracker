package api

import (
	"time"

	"example.com/tracker/internal/entity"
)

type TaskStartRequest struct {
	TaskName        string    `json:"task_name" binding:"required" example:"Time-Tracker API" default:"Time-Tracker API"`
	TaskDescription string    `json:"task_description" binding:"required" example:"Write an API Server for Effective Mobile"`
	TaskEnds        time.Time `json:"task_ends" binding:"omitempty" default:"2024-07-16T23:00:00Z"`
}

type WorkLoadsRequest struct {
	StartDate time.Time `form:"start_date" binding:"required" default:"2024-04-16T23:00:00Z"`
	EndDate   time.Time `form:"end_date" binding:"required" default:"2024-06-16T23:00:00Z"`
}

type WorkLoadsResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Body    []*entity.WorkLoad `json:"body"`
}
