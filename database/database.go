package database

import (
	"URL_shortner/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func PostgresDBInit(cfg config.Config) (*sql.DB, error) {
	strCon := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		cfg.DB_USER, cfg.DB_NAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT)
	var err error
	var db *sql.DB
	db, err = sql.Open("postgres", strCon)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Printf("Connected to database %s", cfg.DB_NAME)

	return db, nil
}
