package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
   PUBLIC_HOST string
   PORT string
   DB_PORT string
   DB_USER string
   DB_PWD  string
   DB_NAME string
   DB_ADDR string
   JWTExpiration int64
   JWT_SECRET string
 }

var Envs = initConfig();

func initConfig() Config{
	godotenv.Load()
return Config{
	PUBLIC_HOST:getEnv("PUBLIC_HOST","http://localhost"),
	PORT:getEnv("PORT","8080"),
	DB_PORT: getEnv("DB_PORT","3306"),
	DB_USER: getEnv("DB_USER","root"),
	DB_PWD: getEnv("DB_PWD","yordanos"),
	DB_NAME: getEnv("DB_NAME","EcomGo"),
	DB_ADDR: fmt.Sprintf("%s:%s",getEnv("PUBLIC_HOST","127.0.0.1"),getEnv("DB_PORT","3306")),	
    JWTExpiration: getEnvAsInt("JWTExpiration",3600 * 24 * 7),
	JWT_SECRET: getEnv("JWT_SECRET","$2b$10$yG7Ivndj5Q7FxXHvfY1Xh.1yqFOsclCAXPYygwKopAZwgUDEn2WS6"),
}
}

func getEnv(key,fallback string) string{
if value,ok := os.LookupEnv(key); ok{
	return value
}
	return fallback
}

func getEnvAsInt(key string,fallback int64) (int64){

   if value, ok := os.LookupEnv(key); ok{
	i, err := strconv.ParseInt(value,10,64)
	if err != nil {
		return fallback
	}
	return i
   }
   return fallback
}