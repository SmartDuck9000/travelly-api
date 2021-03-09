package server

import (
	"github.com/SmartDuck9000/travelly-api/services/auth_service/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (api AuthAPI) register(c *gin.Context) {
	var user db.User
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = api.db.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
	} else {
		api.returnTokens(c, user.ID)
	}
}

func (api AuthAPI) login(c *gin.Context) {
	var user db.User
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userData := api.db.GetUser(user.Email)
	if userData == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User with this email doesn't exist",
		})
		return
	}

	if user.Password != userData.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Wrong password",
		})
	} else {
		api.returnTokens(c, userData.ID)
	}
}

func (api AuthAPI) refreshToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "No authorization header",
		})
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Wrong length of header",
		})
		return
	}

	if headerParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Wrong header",
		})
		return
	}

	claims, err := api.tokenManager.ParseToken(headerParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	api.returnTokens(c, claims.ID)
}

func (api AuthAPI) returnTokens(c *gin.Context, userID int) {
	accessToken, accessTokenErr := api.tokenManager.CreateAccessToken(userID)
	if accessTokenErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": accessTokenErr.Error(),
		})
		return
	}

	refreshToken, refreshTokenErr := api.tokenManager.CreateRefreshToken(userID)
	if refreshTokenErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": refreshTokenErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
