package config

import (
	"os"
	"strconv"
	"time"
)

type AuthDBConfig struct {
	URL             string
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifetime time.Duration
}

type TokenConfig struct {
	Method          string
	AccessKey       string
	RefreshKey      string
	AccessLifeTime  time.Duration
	RefreshLifeTime time.Duration
}

type AuthServiceConfig struct {
	DB    AuthDBConfig
	Token TokenConfig
	Host  string
	Port  string
}

// New returns a new Config struct
func New() *AuthServiceConfig {
	return &AuthServiceConfig{
		DB: AuthDBConfig{
			URL:             getEnv("DB_URL", ""),
			MaxIdleConn:     getIntEnv("MAX_IDLE_CONN", 10),
			MaxOpenConn:     getIntEnv("MAX_OPEN_CONN", 100),
			ConnMaxLifetime: getHoursEnv("CONN_MAX_LIFETIME", 1),
		},
		Token: TokenConfig{
			Method:          getEnv("SIGNING_METHOD", ""),
			AccessKey:       getEnv("ACCESS_KEY", ""),
			RefreshKey:      getEnv("REFRESH_KEY", ""),
			AccessLifeTime:  getHoursEnv("ACCESS_LIFETIME_HOURS", 1),
			RefreshLifeTime: getHoursEnv("REFRESH_LIFETIME_HOURS", 24),
		},
		Host: getEnv("HOST", ""),
		Port: getEnv("PORT", ""),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getIntEnv(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into time.Hour or return a default value
func getHoursEnv(name string, defaultVal int) time.Duration {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return time.Hour * time.Duration(value)
	}

	return time.Hour * time.Duration(defaultVal)
}
