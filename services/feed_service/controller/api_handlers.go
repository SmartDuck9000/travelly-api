package controller

import (
	"github.com/SmartDuck9000/travelly-api/services/feed_service/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller FeedController) getHotels(c *gin.Context) {
	var feedParameters db.FeedParameters
	var hotels []db.Hotel

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	err := c.Bind(&feedParameters)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hotels, err = controller.model.GetHotels(feedParameters, authHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, hotels)
	}
}

func (controller FeedController) getEvents(c *gin.Context) {
	var feedParameters db.FeedParameters
	var events []db.Event

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	err := c.Bind(&feedParameters)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	events, err = controller.model.GetEvents(feedParameters, authHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, events)
	}
}

func (controller FeedController) getRestaurants(c *gin.Context) {
	var feedParameters db.FeedParameters
	var restaurants []db.Restaurant

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	err := c.Bind(&feedParameters)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	restaurants, err = controller.model.GetRestaurants(feedParameters, authHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, restaurants)
	}
}
