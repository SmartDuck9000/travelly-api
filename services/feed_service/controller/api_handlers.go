package controller

import (
	"errors"
	db2 "github.com/SmartDuck9000/travelly-api/server/db"
	"github.com/SmartDuck9000/travelly-api/token_manager"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller FeedController) getHotels(c *gin.Context) {
	var filterParameters db2.HotelFilterParameters
	var hotels []db2.Hotel

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
	var filterParameters db2.EventsFilterParameters
	var events []db2.Event

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
	var filterParameters db2.RestaurantFilterParameters
	var restaurants []db2.Restaurant

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
	var filterParameters db2.TicketFilterParameters
	var tickets []db2.Ticket

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
