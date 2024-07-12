package pgrepo

import (
	"context"
	"errors"
	"fmt"

	"example.com/tracker/internal/entity"
	"github.com/sirupsen/logrus"
)

func (p *Postgres) CreateUser(user *entity.User) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			surname,
			first_name,
			patronymic,
			address,
			passport_serie,
			passport_number
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`, usersTable)
	err := p.DB.QueryRow(query, user.Surname, user.First_name,
		user.Patronymic, user.Address,
		user.PassportSerie, user.PassportNumber).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

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
		logrus.Error(err)
		return nil, err
	}
	return users, err
}

func (p *Postgres) UpdateUser(user *entity.User) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			surname = COALESCE($1, surname),
			first_name = COALESCE($2, first_name),
			patronymic = COALESCE($3, patronymic),
			address = COALESCE($4, address),
			passport_serie = COALESCE($5, passport_serie),
			passport_number = COALESCE($6, passport_number)
		WHERE
			id = $7
	`, usersTable)

	tag, err := p.DB.Exec(query, user.Surname, user.First_name, user.Patronymic, user.Address, user.PassportSerie, user.PassportNumber, user.ID)
	if err != nil {
		logrus.Error(err)
		return err
	}
	rowsSum, _ := tag.RowsAffected()
	if rowsSum == 0 {
		return errors.New("user doesn't exist")
	}
	return nil
}

func (p *Postgres) DeleteUser(ctx context.Context, userID int64) error {
	query := fmt.Sprintf(
		`DELETE FROM %s
		 WHERE
		 	id = $1`, usersTable)
	tag, err := p.DB.Exec(query, userID)
	if err != nil {
		logrus.Error(err)
		return err
	}
	rowsSum, _ := tag.RowsAffected()
	if rowsSum == 0 {
		return ErrRecordNotFound
	}
	return nil
}
