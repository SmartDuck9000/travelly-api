package config_reader

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

type EnvReader struct{}

func CreateEnvReader(filename string) (*EnvReader, error) {
	var reader = &EnvReader{}
	err := godotenv.Load(filename)

	if err != nil {
		log.Print("Config file not found")
		reader = nil
	}

	return reader, err
}

func (reader EnvReader) GetString(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func (reader EnvReader) GetInt(key string, defaultVal int) int {
	valueStr := reader.GetString(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func (reader EnvReader) GetHours(key string, defaultVal int) time.Duration {
	valueStr := reader.GetString(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return time.Hour * time.Duration(value)
	}

	return time.Hour * time.Duration(defaultVal)
}
