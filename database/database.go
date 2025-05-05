package database

import (
	"URL_shortner/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func PostgresDBInit(cfg config.Config) (*sql.DB, error) {
	fmt.Println(cfg)
	var err error
	var db *sql.DB
	db, err = sql.Open("postgres", cfg.DB_URL)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Printf("Connected to database")

	return db, nil
}
