package config

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	Port      string
	Host      string
	DB_url    string
	StartTime time.Time
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
	port := GetEnv("POSTRGRES_PORT", "5432")
	db := GetEnv("POSTGRES_DB", "db")

	DB_url = fmt.Sprintf("postgres://%s:%s@workerbee-database:%s/%s?sslmode=disable", user, password, port, db)

	StartTime = time.Now()
}
