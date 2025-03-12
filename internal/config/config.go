package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	Port string
	DBConn string
	JWTSecret string
}


func LoadConfig() *Config{
    if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

   return &Config{
	Port: getEnv("PORT","8080"),
	DBConn: getEnv("DATABASE_URL","postgres://go_user:securepassword@localhost:5432/go_server?sslmode=disable"),
	JWTSecret: getEnv("JWT_SECRET","default-secret-for-dev-only"),
   }
}

func getEnv(key string , fallback string) string{
	if value, exists := os.LookupEnv(key); exists{
		return value
	}
	return fallback
}


