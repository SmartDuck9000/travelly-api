package config

import (
	"github.com/SmartDuck9000/travelly-api/config_reader"
	"github.com/SmartDuck9000/travelly-api/token_manager"
	"time"
)

type UserDbConfig struct {
	URL             string
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifetime time.Duration
}

type UserModelConfig struct {
	DbConfig    *UserDbConfig
	TokenConfig *token_manager.TokenConfig
}

type UserControllerConfig struct {
	ModelConfig *UserModelConfig
	Host        string
	Port        string
}

func CreateUserDbConfig(reader config_reader.ConfigReader) *UserDbConfig {
	return &UserDbConfig{
		URL:             reader.GetString("DB_URL", ""),
		MaxIdleConn:     reader.GetInt("MAX_IDLE_CONN", 10),
		MaxOpenConn:     reader.GetInt("MAX_OPEN_CONN", 100),
		ConnMaxLifetime: reader.GetHours("CONN_MAX_LIFETIME", 1),
	}
}

func CreateUserModelConfig(reader config_reader.ConfigReader) *UserModelConfig {
	return &UserModelConfig{
		DbConfig:    CreateUserDbConfig(reader),
		TokenConfig: token_manager.CreateTokenConfig(reader),
	}
}

func CreateUserControllerConfig(reader config_reader.ConfigReader) *UserControllerConfig {
	return &UserControllerConfig{
		ModelConfig: CreateUserModelConfig(reader),
		Host:        reader.GetString("HOST", ""),
		Port:        reader.GetString("PORT", ""),
	}
}
