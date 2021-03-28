package config

import (
	"os"
	"strconv"
	"time"
)

type FullInfoDbConfig struct {
	URL             string
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifetime time.Duration
}

type FullInfoModelConfig struct {
	Db *FullInfoDbConfig
}

type FullInfoControllerConfig struct {
	Model *FullInfoModelConfig
	Host  string
	Port  string
}

func CreateFullInfoDbConfig() *FullInfoDbConfig {
	return &FullInfoDbConfig{
		URL:             getEnv("DB_URL", ""),
		MaxIdleConn:     getIntEnv("MAX_IDLE_CONN", 10),
		MaxOpenConn:     getIntEnv("MAX_OPEN_CONN", 100),
		ConnMaxLifetime: getHoursEnv("CONN_MAX_LIFETIME", 1),
	}
}

func CreateFullInfoModelConfig() *FullInfoModelConfig {
	return &FullInfoModelConfig{
		Db: CreateFullInfoDbConfig(),
	}
}

func CreateFullInfoControllerConfig() *FullInfoControllerConfig {
	return &FullInfoControllerConfig{
		Model: CreateFullInfoModelConfig(),
		Host:  getEnv("HOST", ""),
		Port:  getEnv("PORT", ""),
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
