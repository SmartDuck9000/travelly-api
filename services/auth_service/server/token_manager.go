package server

import (
	"fmt"
	"github.com/SmartDuck9000/travelly-api/services/auth_service/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenManager interface {
	CreateAccessToken(id int) (string, error)
	CreateRefreshToken(id int) (string, error)
	ParseToken(tokenString string) (*AuthClaims, error)
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

	return token.SignedString([]byte(m.accessKey))
}

func (m JWTManager) CreateRefreshToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod(m.method), &AuthClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(m.accessLifetime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		ID: id,
	})

	return token.SignedString([]byte(m.refreshKey))
}

func (m JWTManager) ParseToken(tokenString string) (*AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
		}

		return m.refreshKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(AuthClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
