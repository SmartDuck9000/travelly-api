package config

import (
	"github.com/SmartDuck9000/travelly-api/config_reader"
	"github.com/SmartDuck9000/travelly-api/token_manager"
	"time"
)

type FeedDBConfig struct {
	URL             string
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifetime time.Duration
}

type FeedModelConfig struct {
	DbConfig    *FeedDBConfig
	TokenConfig *token_manager.TokenConfig
}

type FeedControllerConfig struct {
	ModelConfig *FeedModelConfig
	Host        string
	Port        string
}

func CreateFeedDbConfig(reader config_reader.ConfigReader) *FeedDBConfig {
	return &FeedDBConfig{
		URL:             reader.GetString("DB_URL", ""),
		MaxIdleConn:     reader.GetInt("MAX_IDLE_CONN", 10),
		MaxOpenConn:     reader.GetInt("MAX_OPEN_CONN", 100),
		ConnMaxLifetime: reader.GetHours("CONN_MAX_LIFETIME", 1),
	}
}

func CreateFeedModelConfig(reader config_reader.ConfigReader) *FeedModelConfig {
	return &FeedModelConfig{
		DbConfig:    CreateFeedDbConfig(reader),
		TokenConfig: token_manager.CreateTokenConfig(reader),
	}
}

func CreateFeedControllerConfig(reader config_reader.ConfigReader) *FeedControllerConfig {
	return &FeedControllerConfig{
		ModelConfig: CreateFeedModelConfig(reader),
		Host:        reader.GetString("HOST", ""),
		Port:        reader.GetString("PORT", ""),
	}
}
