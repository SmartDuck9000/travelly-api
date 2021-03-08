package server

import (
	"github.com/SmartDuck9000/travelly-api/services/auth_service/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenManager interface {
	CreateAccessToken(id int) (string, error)
	CreateRefreshToken(id int) (string, error)
}

type JWTManager struct {
	method          string
	accessKey       string
	refreshKey      string
	accessLifetime  time.Duration
	refreshLifetime time.Duration
}

type AuthClaims struct {
	jwt.StandardClaims
	ID int
}

func CreateJWTManager(conf config.TokenConfig) *JWTManager {
	return &JWTManager{
		method:          conf.Method,
		accessKey:       conf.AccessKey,
		refreshKey:      conf.RefreshKey,
		accessLifetime:  conf.AccessLifeTime,
		refreshLifetime: conf.RefreshLifeTime,
	}
}

func (m JWTManager) CreateAccessToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod(m.method), &AuthClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(m.accessLifetime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		ID: id,
	})

	return token.SignedString(m.accessKey)
}

func (m JWTManager) CreateRefreshToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod(m.method), &AuthClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(m.accessLifetime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		ID: id,
	})

	return token.SignedString(m.refreshKey)
}
