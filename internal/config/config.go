package config

import "os"

type Config struct{
	Port string
	DBConn string
	JWTSecret string
}


func LoadConfig() Config{
   return Config{
	Port: getEnv("PORT","8080"),
	DBConn: getEnv("DB_CONN","postgres://go_user:securepassword@localhost:5432/go_server?sslmode=disable"),
	JWTSecret: getEnv("JWT_SECRET","default-secret-for-dev-only"),
   }
}

func getEnv(key string , fallback string) string{
	if value, exists := os.LookupEnv(key); exists{
		return value
	}
	return fallback
}


