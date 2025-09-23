package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PUBLIC_HOST   string
	PORT          string
	DB_PORT       string
	DB_USER       string
	DB_PWD        string
	DB_NAME       string
	DB_ADDR       string
	JWTExpiration int64
	JWT_SECRET    string
}

var Envs = initConfig()

func initConfig() Config {
	// Load .env locally if exists
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using environment variables")
	}

	dbHost := getEnv("MYSQLHOST" // Railway injects MYSQLHOST
	dbPort := getEnv("MYSQLPORT")      // Railway injects MYSQLPORT

	return Config{
		PUBLIC_HOST:   getEnv("PUBLIC_HOST", ),
		PORT:          getEnv("PORT"),
		DB_PORT:       dbPort,
		DB_USER:       getEnv("MYSQLUSER"),
		DB_PWD:        getEnv("MYSQLPASSWORD"),
		DB_NAME:       getEnv("MYSQLDATABASE"),
		DB_ADDR:       fmt.Sprintf("%s:%s", dbHost, dbPort),
		JWTExpiration: getEnvAsInt("JWTExpiration", 3600*24*7),
		JWT_SECRET:    getEnv("JWT_SECRET", "$2b$10$yG7Ivndj5Q7FxXHvfY1Xh.1yqFOsclCAXPYygwKopAZwgUDEn2WS6"),
	}
}

// getEnv reads an environment variable or returns fallback
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// getEnvAsInt reads an environment variable as int64 or returns fallback
func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
