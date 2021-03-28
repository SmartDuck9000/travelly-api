package token_manager

import (
	"github.com/SmartDuck9000/travelly-api/config_reader"
	"time"
)

type TokenConfig struct {
	Method          string
	AccessKey       string
	RefreshKey      string
	AccessLifeTime  time.Duration
	RefreshLifeTime time.Duration
}

func CreateTokenConfig(reader config_reader.ConfigReader) *TokenConfig {
	return &TokenConfig{
		Method:          reader.GetString("SIGNING_METHOD", ""),
		AccessKey:       reader.GetString("ACCESS_KEY", ""),
		RefreshKey:      reader.GetString("REFRESH_KEY", ""),
		AccessLifeTime:  reader.GetHours("ACCESS_LIFETIME_HOURS", 1),
		RefreshLifeTime: reader.GetHours("REFRESH_LIFETIME_HOURS", 24),
	}
}
