package controller

import (
	"errors"
	"github.com/SmartDuck9000/travelly-api/services/feed_service/db"
	"github.com/SmartDuck9000/travelly-api/token_manager"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller FeedController) getHotels(c *gin.Context) {
	var filterParameters db.HotelFilterParameters
	var hotels []db.Hotel

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	err := c.ShouldBind(&filterParameters)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hotels, err = controller.model.GetHotels(filterParameters, authHeader)
	if err != nil {
		var statusCode = http.StatusBadRequest
		if errors.Is(err, token_manager.InvalidTokenError{}) {
			statusCode = http.StatusUnauthorized
		}

		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, hotels)
	}
}

func (controller FeedController) getEvents(c *gin.Context) {
	var filterParameters db.EventsFilterParameters
	var events []db.Event

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	err := c.ShouldBind(&filterParameters)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	events, err = controller.model.GetEvents(filterParameters, authHeader)
	if err != nil {
		var statusCode = http.StatusBadRequest
		if errors.Is(err, token_manager.InvalidTokenError{}) {
			statusCode = http.StatusUnauthorized
		}

		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, events)
	}
}

func (controller FeedController) getRestaurants(c *gin.Context) {
	var filterParameters db.RestaurantFilterParameters
	var restaurants []db.Restaurant

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	err := c.ShouldBind(&filterParameters)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	restaurants, err = controller.model.GetRestaurants(filterParameters, authHeader)
	if err != nil {
		var statusCode = http.StatusBadRequest
		if errors.Is(err, token_manager.InvalidTokenError{}) {
			statusCode = http.StatusUnauthorized
		}

		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, restaurants)
	}
}

func (controller FeedController) getTickets(c *gin.Context) {
	var filterParameters db.TicketFilterParameters
	var tickets []db.Ticket

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	err := c.ShouldBind(&filterParameters)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tickets, err = controller.model.GetTickets(filterParameters, authHeader)
	if err != nil {
		var statusCode = http.StatusBadRequest
		if errors.Is(err, token_manager.InvalidTokenError{}) {
			statusCode = http.StatusUnauthorized
		}

		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, tickets)
	}
}
