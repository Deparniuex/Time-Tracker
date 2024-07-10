package pgrepo

import (
	"context"
	"fmt"

	"example.com/tracker/internal/entity"
	"github.com/sirupsen/logrus"
)

func (p *Postgres) GetUsers(ctx context.Context) ([]*entity.User, error) {
	query := fmt.Sprintf(`SELECT * FROM %s`, usersTable)
	rows, err := p.DB.Query(query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	var users []*entity.User
	for rows.Next() {
		var user entity.User
		rows.Scan(&user.ID,
			&user.Surname,
			&user.First_name,
			&user.Patronymic,
			&user.Address,
			&user.PassportSerie,
			&user.PassportNumber,
		)
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, err
}
