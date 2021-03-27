package config

import (
	"github.com/SmartDuck9000/travelly-api/config_reader"
	"github.com/SmartDuck9000/travelly-api/token_manager"
	"time"
)

type AuthDbConfig struct {
	URL             string
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifetime time.Duration
}

type AuthModelConfig struct {
	DbConfig    *AuthDbConfig
	TokenConfig *token_manager.TokenConfig
}

type AuthControllerConfig struct {
	ModelConfig *AuthModelConfig
	Host        string
	Port        string
}

func CreateAuthDbConfig(reader config_reader.ConfigReader) *AuthDbConfig {
	return &AuthDbConfig{
		URL:             reader.GetString("DB_URL", ""),
		MaxIdleConn:     reader.GetInt("MAX_IDLE_CONN", 10),
		MaxOpenConn:     reader.GetInt("MAX_OPEN_CONN", 100),
		ConnMaxLifetime: reader.GetHours("CONN_MAX_LIFETIME", 1),
	}
}

func CreateAuthModelConfig(reader config_reader.ConfigReader) *AuthModelConfig {
	return &AuthModelConfig{
		DbConfig:    CreateAuthDbConfig(reader),
		TokenConfig: token_manager.CreateTokenConfig(reader),
	}
}

func CreateAuthControllerConfig(reader config_reader.ConfigReader) *AuthControllerConfig {
	return &AuthControllerConfig{
		ModelConfig: CreateAuthModelConfig(reader),
		Host:        reader.GetString("HOST", ""),
		Port:        reader.GetString("PORT", ""),
	}
}
