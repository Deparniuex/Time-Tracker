package pgrepo

import (
	"errors"
	"fmt"
	"time"

	"example.com/tracker/internal/entity"
)

func (p *Postgres) StartTimer(task *entity.Task) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			user_id,
			task_starts,
			task_name,
			task_description,
			task_ends
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
		`, tasksTable)
	err := p.DB.QueryRow(query, task.UserID, time.Now(), task.Name, task.Description, task.EndAt).Scan(&task.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) EndTimer(taskID int64) error {
	query := fmt.Sprintf(`
		UPDATE %s
   			SET task_ends = $1
    	WHERE id = $2;
	`, tasksTable)
	tag, err := p.DB.Exec(query, time.Now(), taskID)
	if err != nil {
		return err
	}
	rowsSum, _ := tag.RowsAffected()
	if rowsSum == 0 {
		return errors.New("task doesn't exist")
	}
	return nil
}

func (p *Postgres) GetWorkLoads(userID int64, startDate, endDate time.Time) ([]*entity.WorkLoad, error) {
	query := fmt.Sprintf(`
    	SELECT t.id, t.task_name, SUM(EXTRACT(EPOCH FROM (t.task_ends - t.task_starts))) AS total
    	FROM %s t
    	WHERE t.user_id = $1 AND t.task_starts >= $2 AND t.task_ends <= $3
    	GROUP BY t.id, task_name
    	ORDER BY total DESC;
    `, tasksTable)
	rows, err := p.DB.Query(query, userID, startDate, endDate)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workLogs []*entity.WorkLoad
	for rows.Next() {
		var load entity.WorkLoad
		if err := rows.Scan(&load.TaskID, &load.TaskName, &load.Total); err != nil {
			return nil, err
		}
		workLogs = append(workLogs, &load)
	}

	return workLogs, nil
}
