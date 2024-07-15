package entity

import "time"

type Task struct {
	ID          int64     `db:"id"`
	UserID      int64     `db:"user_id"`
	EndAt       time.Time `db:"task_ends"`
	Name        string    `db:"task_name"`
	Description string    `db:"task_description"`
}

type WorkLoad struct {
	TaskID   int
	TaskName string
	Total    float64
}
