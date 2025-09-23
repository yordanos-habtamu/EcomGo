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

	// Use Railway environment variables if they exist, otherwise fallback to local
	dbHost := getEnv("RAILWAY_MYSQL_HOST", getEnv("PUBLIC_HOST", "127.0.0.1"))
	dbPort := getEnv("RAILWAY_MYSQL_PORT", getEnv("DB_PORT", "3306"))
	dbUser := getEnv("RAILWAY_MYSQL_USER", getEnv("DB_USER", "root"))
	dbPwd := getEnv("RAILWAY_MYSQL_PASSWORD", getEnv("DB_PWD", "password"))
	dbName := getEnv("RAILWAY_MYSQL_DATABASE", getEnv("DB_NAME", "EcomGo"))

	return Config{
		PUBLIC_HOST:   getEnv("PUBLIC_HOST", "127.0.0.1"),
		PORT:          getEnv("PORT", "8080"),
		DB_PORT:       dbPort,
		DB_USER:       dbUser,
		DB_PWD:        dbPwd,
		DB_NAME:       dbName,
		DB_ADDR:       fmt.Sprintf("%s:%s", dbHost, dbPort),
		JWTExpiration: getEnvAsInt("JWTExpiration", 3600*24*7),
		JWT_SECRET:    getEnv("JWT_SECRET", "$2b$10$yG7Ivndj5Q7FxXHvfY1Xh.1yqFOsclCAXPYygwKopAZwgUDEn2WS6"),
	}
}

// getEnv reads an environment variable or returns fallback
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}

// getEnvAsInt reads an environment variable as int64 or returns fallback
func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
