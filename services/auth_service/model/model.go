package model

import (
	"fmt"
	"github.com/SmartDuck9000/travelly-api/services/auth_service/config"
	"github.com/SmartDuck9000/travelly-api/services/auth_service/db"
	"github.com/SmartDuck9000/travelly-api/token_manager"
	"strings"
)

type AuthModelInterface interface {
	Run() error
	RefreshToken(httpHeader string) (*AuthData, error)
	Register(user db.User) (*AuthData, error)
	Login(user db.User) (*AuthData, error)
}

type AuthModel struct {
	db           db.AuthDB
	tokenManager token_manager.TokenManager
}

type AuthData struct {
	UserId       int    `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func CreateAuthModel(config config.AuthModelConfig) AuthModelInterface {
	return &AuthModel{
		db:           db.CreateAuthDB(*config.DbConfig),
		tokenManager: token_manager.CreateJWTManager(*config.TokenConfig),
	}
}

func (model AuthModel) Run() error {
	return model.db.Open()
}

func (model AuthModel) RefreshToken(authHeader string) (*AuthData, error) {
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return nil, fmt.Errorf("wrong length of header")
	}

	if headerParts[0] != "Bearer" {
		return nil, fmt.Errorf("wrong header")
	}

	claims, err := model.tokenManager.ParseRefreshToken(headerParts[1])
	if err != nil {
		return nil, err
	}

	return model.getAuthData(claims.ID)
}

func (model AuthModel) Register(user db.User) (*AuthData, error) {
	err := model.db.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return model.getAuthData(user.ID)
}

func (model AuthModel) Login(user db.User) (*AuthData, error) {
	userData := model.db.GetUser(user.Email)
	if userData == nil {
		return nil, fmt.Errorf("user with this email doesn't exist")
	}

	if user.Password != userData.Password {
		return nil, fmt.Errorf("wrong password")
	}

	return model.getAuthData(userData.ID)
}

func (model AuthModel) getAuthData(userId int) (*AuthData, error) {
	accessToken, accessTokenErr := model.tokenManager.CreateAccessToken(userId)
	if accessTokenErr != nil {
		return nil, accessTokenErr
	}

	refreshToken, refreshTokenErr := model.tokenManager.CreateRefreshToken(userId)
	if refreshTokenErr != nil {
		return nil, refreshTokenErr
	}

	return &AuthData{
		UserId:       userId,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
