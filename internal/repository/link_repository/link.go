package repository

import (
	"database/sql"
)

type PostgresLinkRepository struct {
	Db *sql.DB
}

func NewPostgresLinkRepository(db *sql.DB) *PostgresLinkRepository {
	return &PostgresLinkRepository{db}
}
