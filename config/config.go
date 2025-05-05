package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB_URL     string `env:"DB_URL"`
	LOCAL_PORT string `env:"LOCAL_PORT" envDefault:"8080"`
	JWT_SECRET string `env:"JWT_SICRET" envDefault:""`
}

func LoadConfig() Config {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal(err)
	}

	cfg := Config{
		DB_URL:     os.Getenv("DB_URL"),
		LOCAL_PORT: os.Getenv("LOCAL_PORT"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}

	return cfg
}
