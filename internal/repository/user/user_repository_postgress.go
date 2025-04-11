package repository

import (
	"URL_shortner/internal/model"
	"database/sql"
)

type postgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *postgresUserRepository {
	return &postgresUserRepository{db: db}
}

func (pr *postgresUserRepository) NewUser(user *model.User) error {
	_, err := pr.db.Exec("INSERT INTO users(email, password) VALUES ($1, $2)", user.Email, user.Password)

	if err != nil {
		return err
	}
	return nil
}

func (pr *postgresUserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User

	query := "SELECT id, email, password  FROM users WHERE email = $1"
	err := pr.db.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
