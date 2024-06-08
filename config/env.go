package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	Host                   string
	Port                   int64
	DBUser                 string
	DBPassword             string
	DBName                 string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

// Gets the env by key or fallbacks
func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return ""
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		Host:                   getEnv("DB_HOST"),
		Port:                   getEnvAsInt("DB_PORT", 5432),
		DBUser:                 getEnv("DB_USER"),
		DBPassword:             getEnv("DB_PASSWORD"),
		DBName:                 getEnv("DB_NAME"),
		JWTSecret:              getEnv("JWT_SECRET"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
	}
}

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
