package config

import (
	"os"
	"time"
)

var (
	Port         string
	Host         string
	ClientID     string
	ClientSecret string
	AuthentikURL string
	StartTime    time.Time
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
	ClientID = GetEnv("CLIENT_ID", "client-id")
	ClientSecret = GetEnv("CLIENT_SECRET", "client-secret")
	AuthentikURL = GetEnv("AUTHENTIK_URL", "http://localhost:9000")
	StartTime = time.Now()
}
