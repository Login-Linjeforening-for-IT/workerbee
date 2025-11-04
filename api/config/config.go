package config

import (
	"fmt"
	"os"
	"time"
)

var (
	Port                     string
	Host                     string
	DB_url                   string
	DO_URL                   string
	DO_access_key_id         string
	DO_secret_access_key     string
	StartTime                time.Time
	AllowedRequestsPerMinute int
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Init() {
	Port = GetEnv("PORT", "8080")
	Host = GetEnv("HOST", "0.0.0.0")

	user := GetEnv("POSTGRES_USER", "admin")
	password := GetEnv("POSTGRES_PASSWORD", "admin")
	port := GetEnv("POSTGRES_PORT", "5432")
	db_name := GetEnv("POSTGRES_DB", "db")
	db_host := GetEnv("POSTGRES_HOST", "localhost")

	DB_url = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, db_host, port, db_name)
	DO_URL = GetEnv("DO_URL", "")
	DO_access_key_id = GetEnv("DO_ACCESS_KEY_ID", "")
	DO_secret_access_key = GetEnv("DO_SECRET_ACCESS_KEY", "")
	StartTime = time.Now()
	RateLimitRoofStr := GetEnv("ALLOWED_PROTECTED_REQUESTS", "25")
	_, err := fmt.Sscanf(RateLimitRoofStr, "%d", &AllowedRequestsPerMinute)
	if err != nil || AllowedRequestsPerMinute <= 0 {
		AllowedRequestsPerMinute = 25
	}
}
