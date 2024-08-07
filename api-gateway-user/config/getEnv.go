package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	GATEWAY_USER_SERVICE_PORT string
	POLL_SERVICE_PORT    string
	DB_HOST              string
	DB_PORT              int
	DB_USER              string
	DB_PASSWORD          string
	DB_NAME              string
	SENDER_EMAIL         string
	APP_PASSWORD         string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.GATEWAY_USER_SERVICE_PORT = cast.ToString(coalesce("user_service", ":8081"))
	config.POLL_SERVICE_PORT = cast.ToString(coalesce("POLL_SERVICE_PORT", ":50051"))
	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToInt(coalesce("DB_PORT", 5432))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "root"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "delivery_auth"))
	config.SENDER_EMAIL = cast.ToString(coalesce("SENDER_EMAIL", "email@example.com"))
	config.APP_PASSWORD = cast.ToString(coalesce("APP_PASSWORD", "your_app_password_here"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
