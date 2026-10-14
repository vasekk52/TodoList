package config

import (
	"os"
	"todoapp/internal/model"

	"github.com/joho/godotenv"
)

func Load() model.Config {
	godotenv.Load()
	Config := model.Config{}
	Config.DB_HOST = getEnv("DB_HOST", "localhost")
    Config.DB_PORT = getEnv("DB_PORT", "5432")
    Config.DB_USER = getEnv("DB_USER", "postgres")
    Config.DB_PASS = getEnv("DB_PASSWORD", "postgres123")
    Config.DB_NAME = getEnv("DB_NAME", "todo_db")
	Config.ServerPort = getEnv("SERVER_PORT", "8080")
	return Config
}
func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists && value != "" {
        return value
    }
	return defaultValue
}