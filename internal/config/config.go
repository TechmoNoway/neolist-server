package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DBHost string
	DBName string
	DBPort string
	DBUser string
	DBPass string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Panicln("No .env file found")
	}

	cfg := &Config{
		Port:   getEnv("PORT", ""),
		DBHost: getEnv("DB_HOST", ""),
		DBName: getEnv("DB_NAME", ""),
		DBPort: getEnv("DB_PORT", ""),
		DBUser: getEnv("DB_USER", ""),
		DBPass: getEnv("DB_PASSWORD", ""),
	}

	log.Println("Config loaded")
	return cfg
}

func getEnv(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}
