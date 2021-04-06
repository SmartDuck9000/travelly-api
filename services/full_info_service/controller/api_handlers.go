package controller

import (
	"errors"
	"github.com/SmartDuck9000/travelly-api/services/full_info_service/db"
	"github.com/SmartDuck9000/travelly-api/token_manager"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (controller FullInfoController) getHotel(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no authorization header",
		})
		return
	}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		var hotel *db.Hotel
		hotel, err = controller.model.GetHotel(id, authHeader)

		if err != nil {
			var statusCode = http.StatusBadRequest
			if errors.Is(err, token_manager.InvalidTokenError{}) {
				statusCode = http.StatusUnauthorized
			}

			c.JSON(statusCode, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, *hotel)
		}
	}
}

func (controller FullInfoController) getEvent(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no authorization header",
		})
		return
	}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		var event *db.Event
		event, err = controller.model.GetEvent(id, authHeader)

		if err != nil {
			var statusCode = http.StatusBadRequest
			if errors.Is(err, token_manager.InvalidTokenError{}) {
				statusCode = http.StatusUnauthorized
			}

			c.JSON(statusCode, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, *event)
		}
	}
}

func (controller FullInfoController) getRestaurant(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no authorization header",
		})
		return
	}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		var restaurant *db.Restaurant
		restaurant, err = controller.model.GetRestaurant(id, authHeader)

		if err != nil {
			var statusCode = http.StatusBadRequest
			if errors.Is(err, token_manager.InvalidTokenError{}) {
				statusCode = http.StatusUnauthorized
			}

			c.JSON(statusCode, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, *restaurant)
		}
	}
}

func (controller FullInfoController) getTicket(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no authorization header",
		})
		return
	}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		var ticket *db.Ticket
		ticket, err = controller.model.GetTicket(id, authHeader)

		if err != nil {
			var statusCode = http.StatusBadRequest
			if errors.Is(err, token_manager.InvalidTokenError{}) {
				statusCode = http.StatusUnauthorized
			}

			c.JSON(statusCode, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, *ticket)
		}
	}
}
