package server

import (
	"github.com/SmartDuck9000/travelly-api/services/auth_service/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (api AuthAPI) register(c *gin.Context) {
	var user db.User
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		err := api.db.CreateUser(&user)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		} else {
			var userID = user.ID
			c.JSON(http.StatusOK, gin.H{
				"access_token":  api.tokenManager.CreateAccessToken(userID),
				"refresh_token": api.tokenManager.CreateRefreshToken(userID),
			})
		}
	}
}

func (api AuthAPI) login(c *gin.Context) {
	var user db.User
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		userData := api.db.GetUser(user.Email)
		if userData == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User with this email doesn't exist",
			})
		} else {
			if user.Password != userData.Password {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Wrong password",
				})
			} else {
				var userID = userData.ID
				c.JSON(http.StatusOK, gin.H{
					"access_token":  api.tokenManager.CreateAccessToken(userID),
					"refresh_token": api.tokenManager.CreateRefreshToken(userID),
				})
			}
		}
	}
}
