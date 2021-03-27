package config

import (
	"github.com/SmartDuck9000/travelly-api/config_reader"
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

func New(reader config_reader.ConfigReader) *AuthServiceConfig {
	return &AuthServiceConfig{
		DB: AuthDBConfig{
			URL:             reader.GetString("DB_URL", ""),
			MaxIdleConn:     reader.GetInt("MAX_IDLE_CONN", 10),
			MaxOpenConn:     reader.GetInt("MAX_OPEN_CONN", 100),
			ConnMaxLifetime: reader.GetHours("CONN_MAX_LIFETIME", 1),
		},
		Token: TokenConfig{
			Method:          reader.GetString("SIGNING_METHOD", ""),
			AccessKey:       reader.GetString("ACCESS_KEY", ""),
			RefreshKey:      reader.GetString("REFRESH_KEY", ""),
			AccessLifeTime:  reader.GetHours("ACCESS_LIFETIME_HOURS", 1),
			RefreshLifeTime: reader.GetHours("REFRESH_LIFETIME_HOURS", 24),
		},
		Host: reader.GetString("HOST", ""),
		Port: reader.GetString("PORT", ""),
	}
}
