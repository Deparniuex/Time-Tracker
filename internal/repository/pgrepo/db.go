package pgrepo

import "database/sql"

type Postgres struct {
	DB *sql.DB
}

const (
	usersTable = "users"
)

func New(db *sql.DB) *Postgres {
	return &Postgres{
		DB: db,
	}
}

func (p *Postgres) Close() {
	p.DB.Close()
}
