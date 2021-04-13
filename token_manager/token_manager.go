package token_manager

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type TokenManager interface {
	CreateAccessToken(id int) (string, error)
	CreateRefreshToken(id int) (string, error)
	ParseRefreshToken(tokenString string) (*AuthClaims, error)
	ParseAccessToken(tokenString string) (*AuthClaims, error)
	ValidateToken(authHeader string) error
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

func CreateJWTManager(conf TokenConfig) TokenManager {
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
			ExpiresAt: time.Now().Add(m.refreshLifetime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		ID: id,
	})

	return token.SignedString([]byte(m.refreshKey))
}

func (m JWTManager) ParseRefreshToken(tokenString string) (*AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
		}

		return []byte(m.refreshKey), nil
	})

	if err != nil {
		return nil, InvalidTokenError{}
	}

	claims, ok := token.Claims.(*AuthClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, InvalidTokenError{}
}

func (m JWTManager) ParseAccessToken(tokenString string) (*AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
		}

		return []byte(m.accessKey), nil
	})

	if err != nil {
		return nil, InvalidTokenError{}
	}

	claims, ok := token.Claims.(*AuthClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, InvalidTokenError{}
}

func (m JWTManager) ValidateToken(authHeader string) error {
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return fmt.Errorf("wrong length of header")
	}

	if headerParts[0] != "Bearer" {
		return fmt.Errorf("wrong header")
	}

	_, err := m.ParseAccessToken(headerParts[1])
	return err
}
