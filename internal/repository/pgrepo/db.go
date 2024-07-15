package pgrepo

import (
	"database/sql"
	"errors"
)

type Postgres struct {
	DB *sql.DB
}

var (
	ErrRecordNotFound = errors.New("no records were found")
)

const (
	usersTable = "users"
	tasksTable = "tasks"
)

func New(db *sql.DB) *Postgres {
	return &Postgres{
		DB: db,
	}
}

func (p *Postgres) Close() {
	p.DB.Close()
}
