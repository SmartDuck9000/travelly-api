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
		c.String(http.StatusBadRequest, err.Error())
	} else {
		userData := api.db.CreateUser(user)
		if userData == nil {
			c.String(http.StatusUnauthorized, "User with this email already exist")
		} else {
			c.JSON(http.StatusOK, *userData)
		}
	}
}

func (api AuthAPI) login(c *gin.Context) {
	var user db.User
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		userData := api.db.GetUser(user.Email)
		if userData == nil {
			c.String(http.StatusNotFound, "User with this email doesn't exist")
		} else {
			if user.Password != userData.Password {
				c.String(http.StatusUnauthorized, "Wrong password")
			} else {
				c.JSON(http.StatusOK, *userData)
			}
		}
	}
}
