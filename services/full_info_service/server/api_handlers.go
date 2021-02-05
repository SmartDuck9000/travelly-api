package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (api FullInfoAPI) getHotel(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		user := api.db.GetHotel(id)
		c.JSON(http.StatusOK, *user)
	}
}

func (api FullInfoAPI) getEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		user := api.db.GetEvent(id)
		c.JSON(http.StatusOK, *user)
	}
}

func (api FullInfoAPI) getRestaurant(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		user := api.db.GetRestaurant(id)
		c.JSON(http.StatusOK, *user)
	}
}
