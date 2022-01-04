package config

import (
	"github.com/SmartDuck9000/travelly-api/config_reader"
	"github.com/SmartDuck9000/travelly-api/token_manager"
	"time"
)

type DbConfig struct {
	URL             string
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifetime time.Duration
}

type ModelConfig struct {
	DbConfig    *DbConfig
	TokenConfig *token_manager.TokenConfig
}

type ControllerConfig struct {
	ModelConfig *ModelConfig
	Host        string
	Port        string
}

func CreateDbConfig(reader config_reader.ConfigReader) *DbConfig {
	return &DbConfig{
		URL:             reader.GetString("DB_URL", ""),
		MaxIdleConn:     reader.GetInt("MAX_IDLE_CONN", 10),
		MaxOpenConn:     reader.GetInt("MAX_OPEN_CONN", 100),
		ConnMaxLifetime: reader.GetHours("CONN_MAX_LIFETIME", 1),
	}
}

func CreateModelConfig(reader config_reader.ConfigReader) *ModelConfig {
	return &ModelConfig{
		DbConfig:    CreateDbConfig(reader),
		TokenConfig: token_manager.CreateTokenConfig(reader),
	}
}

func CreateControllerConfig(reader config_reader.ConfigReader) *ControllerConfig {
	return &ControllerConfig{
		ModelConfig: CreateModelConfig(reader),
		Host:        reader.GetString("HOST", ""),
		Port:        reader.GetString("PORT", ""),
	}
}
