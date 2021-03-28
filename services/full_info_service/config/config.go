package config

import (
	"github.com/SmartDuck9000/travelly-api/config_reader"
	"github.com/SmartDuck9000/travelly-api/token_manager"
	"time"
)

type FullInfoDbConfig struct {
	URL             string
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifetime time.Duration
}

type FullInfoModelConfig struct {
	Db          *FullInfoDbConfig
	TokenConfig *token_manager.TokenConfig
}

type FullInfoControllerConfig struct {
	Model *FullInfoModelConfig
	Host  string
	Port  string
}

func CreateFullInfoDbConfig(reader config_reader.ConfigReader) *FullInfoDbConfig {
	return &FullInfoDbConfig{
		URL:             reader.GetString("DB_URL", ""),
		MaxIdleConn:     reader.GetInt("MAX_IDLE_CONN", 10),
		MaxOpenConn:     reader.GetInt("MAX_OPEN_CONN", 100),
		ConnMaxLifetime: reader.GetHours("CONN_MAX_LIFETIME", 1),
	}
}

func CreateFullInfoModelConfig(reader config_reader.ConfigReader) *FullInfoModelConfig {
	return &FullInfoModelConfig{
		Db:          CreateFullInfoDbConfig(reader),
		TokenConfig: token_manager.CreateTokenConfig(reader),
	}
}

func CreateFullInfoControllerConfig(reader config_reader.ConfigReader) *FullInfoControllerConfig {
	return &FullInfoControllerConfig{
		Model: CreateFullInfoModelConfig(reader),
		Host:  reader.GetString("HOST", ""),
		Port:  reader.GetString("PORT", ""),
	}
}
