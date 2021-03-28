package config_reader

import "time"

type ConfigReader interface {
	GetString(key string, defaultVal string) string
	GetInt(key string, defaultVal int) int
	GetHours(key string, defaultVal int) time.Duration
}
