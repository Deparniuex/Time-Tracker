package pgrepo

import (
	"errors"
	"fmt"
	"strings"

	"example.com/tracker/internal/entity"
	"example.com/tracker/pkg/util"
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

func (p *Postgres) GetUsers(pagination *util.Pagination, filters map[string]string) ([]*entity.User, *util.Metadata, error) {
	query := fmt.Sprintf(`SELECT * FROM %s`, usersTable)

	var conditions []string
	var args []interface{}

	paramIndex := 1

	for field, value := range filters {
		if value != "" {
			conditions = append(conditions, fmt.Sprintf(`%s ILIKE $%d`, field, paramIndex))
			args = append(args, value)
			paramIndex++
		}
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")

	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", paramIndex, paramIndex+1)
	args = append(args, pagination.PageSize, pagination.Offset())
	rows, err := p.DB.Query(query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, nil, err
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
		return nil, nil, err
	}

	totalRecords, err := p.countRecords(conditions, args)

	if err != nil {
		return nil, nil, err
	}

	metadata := pagination.CalculateMetadata(totalRecords)
	return users, &metadata, nil
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
		return err
	}

	rowsSum, _ := tag.RowsAffected()

	if rowsSum == 0 {
		return errors.New("user doesn't exist")
	}

	return nil
}

func (p *Postgres) DeleteUser(userID int64) error {
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

func (p *Postgres) countRecords(conditions []string, args []interface{}) (int, error) {
	countQuery := "SELECT COUNT(*) FROM users"

	if len(conditions) > 0 {
		countQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	var totalItems int
	err := p.DB.QueryRow(countQuery, args[:len(args)-2]...).Scan(&totalItems)

	if err != nil {
		return 0, err
	}

	return totalItems, nil
}
