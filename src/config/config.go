package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Postgres PostgresConfig
	Redis    RedisConfig
	JWT      JWTConfig
	CORS     CORSConfig

	AppName     string
	Debug       bool
	Environment string
	PORT        string
}

type PostgresConfig struct {
	User     string
	Password string
	Name     string
	Port     int
	Host     string
}

type RedisConfig struct {
	Port int
	Host string
	Name string
}

type JWTConfig struct {
	PrivateKeyPath           string
	PublicKeyPath            string
	AccessTokenExpireMinutes int
	RefreshTokenExpireDays   int
}

type CORSConfig struct {
	AllowOrigins     []string
	AllowCredentials []string
	AllowMethods     []string
	AllowHeaders     []string
}

func GetConfig() Config {
	// Load .env for local development.
	// In production, environment variables can come directly
	// from Docker, Kubernetes, the OS, etc.
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println(err)
		panic("Can't Load env file..")
	}
	return Config{
		Postgres: PostgresConfig{
			User:     os.Getenv("POSTGRES__USER"),
			Password: os.Getenv("POSTGRES__PASSWORD"),
			Name:     os.Getenv("POSTGRES__NAME"),
			Port:     getIntEnv("POSTGRES__PORT"),
			Host:     os.Getenv("POSTGRES__HOST"),
		},

		Redis: RedisConfig{
			Port: getIntEnv("REDIS__PORT"),
			Host: os.Getenv("REDIS__HOST"),
			Name: os.Getenv("REDIS__NAME"),
		},

		JWT: JWTConfig{
			PrivateKeyPath:           os.Getenv("JWT__PRIVATE_KEY_PATH"),
			PublicKeyPath:            os.Getenv("JWT__PUBLIC_KEY_PATH"),
			AccessTokenExpireMinutes: getIntEnv("JWT__ACCESS_TOKEN_EXPIRE_MINUTES"),
			RefreshTokenExpireDays:   getIntEnv("JWT__REFRESH_TOEKN_EXPIRE_DAYS"),
		},

		CORS: CORSConfig{
			AllowOrigins:     getSliceEnv("CORS__ALLOW_ORIGINS"),
			AllowCredentials: getSliceEnv("CORS__ALLOW_CREDITIONALS"),
			AllowMethods:     getSliceEnv("CORS__ALLOW_METHODS"),
			AllowHeaders:     getSliceEnv("CORS__ALLOW_HEADERS"),
		},

		AppName:     os.Getenv("APP_NAME"),
		Debug:       getBoolEnv("DEBUG"),
		Environment: os.Getenv("ENVIRONMENT"),
		PORT:        os.Getenv("PORT"),
	}
}

func getIntEnv(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		panic("Invalid integer environment variable: " + key)
	}

	return value
}

func getBoolEnv(key string) bool {
	value, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		panic("Invalid boolean environment variable: " + key)
	}

	return value
}

func getSliceEnv(key string) []string {
	value := os.Getenv(key)

	value = strings.TrimSpace(value)
	value = strings.Trim(value, "[]")

	if value == "" {
		return []string{}
	}

	parts := strings.Split(value, ",")
	result := make([]string, 0, len(parts)+1)
	for _, part := range parts {
		part = strings.TrimSpace(part)
		part = strings.Trim(part, `"`)

		if part != "" {
			result = append(result, part)
		}
	}
	return result
}
