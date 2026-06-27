package config

import (
	
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	JWTSecret string
	Dsn       string
}

func LoadEnv() *Config {
	
	_ = godotenv.Load(".env")

	return &Config{
		Port:      os.Getenv("PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		Dsn:       os.Getenv("DSN"),
	}
}
