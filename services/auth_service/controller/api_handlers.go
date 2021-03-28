package controller

import (
	"github.com/SmartDuck9000/travelly-api/services/auth_service/db"
	"github.com/SmartDuck9000/travelly-api/services/auth_service/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller AuthController) register(c *gin.Context) {
	var user db.User
	var authData *model.AuthData

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	authData, err = controller.model.Register(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, *authData)
	}
}

func (controller AuthController) login(c *gin.Context) {
	var user db.User
	var authData *model.AuthData

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	authData, err = controller.model.Login(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, *authData)
	}
}

func (controller AuthController) refreshToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	authData, err := controller.model.RefreshToken(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, *authData)
	}
}
