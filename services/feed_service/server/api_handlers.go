package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (api FeedServiceAPI) getHotels(c *gin.Context) {
	orderedField := c.Query("ordered_by")
	hotels := api.db.GetHotels(orderedField)
	c.JSON(http.StatusOK, hotels)
}

func (api FeedServiceAPI) getEvents(c *gin.Context) {
	orderedField := c.Query("ordered_by")
	events := api.db.GetEvents(orderedField)
	c.JSON(http.StatusOK, events)
}

func (api FeedServiceAPI) getRestaurants(c *gin.Context) {
	orderedField := c.Query("ordered_by")
	restaurants := api.db.GetRestaurants(orderedField)
	c.JSON(http.StatusOK, restaurants)
}
