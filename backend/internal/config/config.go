package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	Environment     string
	DBHost          string
	DBPort          string
	DBUser          string
	DBPassword      string
	DBName          string
	JWTSecret       string
	MinioEndpoint   string
	MinioPublicURL  string
	MinioAccessKey  string
	MinioSecretKey  string
	MinioBucket     string
	MinioUseSSL     bool
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		Port:           getEnv("PORT", "8080"),
		Environment:    getEnv("APP_ENV", "development"),
		DBHost:         getEnv("DB_HOST", "127.0.0.1"),
		DBPort:         getEnv("DB_PORT", "3306"),
		DBUser:         getEnv("DB_USER", "alumni_user"),
		DBPassword:     getEnv("DB_PASSWORD", "alumni_secret"),
		DBName:         getEnv("DB_NAME", "alumni_db"),
		JWTSecret:      getEnv("JWT_SECRET", "super-secret-jwt-key-for-development-32chars"),
		MinioEndpoint:  getEnv("MINIO_ENDPOINT", "127.0.0.1:9000"),
		MinioPublicURL: getEnv("MINIO_PUBLIC_URL", "http://127.0.0.1:9000"),
		MinioAccessKey: getEnv("MINIO_ROOT_USER", "minioadmin"),
		MinioSecretKey: getEnv("MINIO_ROOT_PASSWORD", "minioadmin123"),
		MinioBucket:    getEnv("MINIO_BUCKET", "alumni-media"),
		MinioUseSSL:    getEnv("MINIO_USE_SSL", "false") == "true",
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return fallback
}
