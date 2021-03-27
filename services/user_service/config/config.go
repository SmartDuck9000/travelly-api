package config

import (
	"github.com/SmartDuck9000/travelly-api/config_reader"
	"time"
)

type UserServiceDbConfig struct {
	URL             string
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifetime time.Duration
}

type UserServiceConfig struct {
	DB   UserServiceDbConfig
	Host string
	Port string
}

func New(reader config_reader.ConfigReader) *UserServiceConfig {
	return &UserServiceConfig{
		DB: UserServiceDbConfig{
			URL:             reader.GetString("DB_URL", ""),
			MaxIdleConn:     reader.GetInt("MAX_IDLE_CONN", 10),
			MaxOpenConn:     reader.GetInt("MAX_OPEN_CONN", 100),
			ConnMaxLifetime: reader.GetHours("CONN_MAX_LIFETIME", 1),
		},
		Host: reader.GetString("HOST", ""),
		Port: reader.GetString("PORT", ""),
	}
}
