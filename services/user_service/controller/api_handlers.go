package controller

import (
	"errors"
	"github.com/SmartDuck9000/travelly-api/services/user_service/db"
	"github.com/SmartDuck9000/travelly-api/services/user_service/model"
	"github.com/SmartDuck9000/travelly-api/token_manager"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (controller UserController) getUser(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		var user *db.UserData
		user, err = controller.model.GetUser(userId, authHeader)

		if err != nil {
			var statusCode = http.StatusNotFound
			if errors.Is(err, token_manager.InvalidTokenError{}) {
				statusCode = http.StatusUnauthorized
			}

			c.JSON(statusCode, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, *user)
		}
	}
}

func (controller UserController) getTours(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		var tours []db.TourData
		tours, err = controller.model.GetTours(userId, authHeader)

		if err != nil {
			var statusCode = http.StatusNotFound
			if errors.Is(err, token_manager.InvalidTokenError{}) {
				statusCode = http.StatusUnauthorized
			}

			c.JSON(statusCode, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, tours)
		}
	}
}

func (controller UserController) getCityTours(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	tourId, err := strconv.Atoi(c.Query("tour_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		var cityTours []db.CityTourData
		cityTours, err = controller.model.GetCityTours(tourId, authHeader)

		if err != nil {
			var statusCode = http.StatusNotFound
			if errors.Is(err, token_manager.InvalidTokenError{}) {
				statusCode = http.StatusUnauthorized
			}

			c.JSON(statusCode, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, cityTours)
		}
	}
}

func (controller UserController) getCityTourEvents(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	cityTourId, err := strconv.Atoi(c.Query("city_tour_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		var events []db.EventData
		events, err = controller.model.GetCityTourEvents(cityTourId, authHeader)

		if err != nil {
			var statusCode = http.StatusNotFound
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
}

func (controller UserController) getCityTourRestaurantBookings(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	cityTourId, err := strconv.Atoi(c.Query("city_tour_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		var restaurantBookings []db.RestaurantBookingData
		restaurantBookings, err = controller.model.GetCityTourRestaurantBookings(cityTourId, authHeader)

		if err != nil {
			var statusCode = http.StatusNotFound
			if errors.Is(err, token_manager.InvalidTokenError{}) {
				statusCode = http.StatusUnauthorized
			}

			c.JSON(statusCode, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, restaurantBookings)
		}
	}
}

func (controller UserController) getCityTourTickets(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	cityTourId, err := strconv.Atoi(c.Query("city_tour_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		var tickets *db.CityTourTicketData
		tickets, err = controller.model.GetCityTourTickets(cityTourId, authHeader)

		if err != nil {
			var statusCode = http.StatusNotFound
			if errors.Is(err, token_manager.InvalidTokenError{}) {
				statusCode = http.StatusUnauthorized
			}

			c.JSON(statusCode, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, *tickets)
		}
	}
}

func (controller UserController) getCityTourHotel(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	cityTourId, err := strconv.Atoi(c.Query("city_tour_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		var hotel *db.HotelData
		hotel, err = controller.model.GetCityTourHotel(cityTourId, authHeader)

		if err != nil {
			var statusCode = http.StatusNotFound
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

func (controller UserController) postTour(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	var tour db.Tour
	err := c.Bind(&tour)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = controller.model.CreateTour(&tour, authHeader)
	if err != nil {
		var statusCode = http.StatusBadRequest
		if errors.Is(err, token_manager.InvalidTokenError{}) {
			statusCode = http.StatusUnauthorized
		}

		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func (controller UserController) postCityTour(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	var cityTour db.CityTour
	err := c.Bind(&cityTour)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = controller.model.CreateCityTour(&cityTour, authHeader)
	if err != nil {
		var statusCode = http.StatusBadRequest
		if errors.Is(err, token_manager.InvalidTokenError{}) {
			statusCode = http.StatusUnauthorized
		}

		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func (controller UserController) postRestaurantBooking(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	var rb db.RestaurantBooking
	err := c.Bind(&rb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = controller.model.CreateRestaurantBooking(&rb, authHeader)
	if err != nil {
		var statusCode = http.StatusBadRequest
		if errors.Is(err, token_manager.InvalidTokenError{}) {
			statusCode = http.StatusUnauthorized
		}

		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func (controller UserController) updateUser(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	var user model.UpdateUserData
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = controller.model.UpdateUser(&user, authHeader)
	if err != nil {
		var statusCode = http.StatusBadRequest
		if errors.Is(err, token_manager.InvalidTokenError{}) {
			statusCode = http.StatusUnauthorized
		}

		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func (controller UserController) updateTour(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	var tour db.Tour
	err := c.Bind(&tour)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = controller.model.UpdateTour(&tour, authHeader)
	if err != nil {
		var statusCode = http.StatusBadRequest
		if errors.Is(err, token_manager.InvalidTokenError{}) {
			statusCode = http.StatusUnauthorized
		}

		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func (controller UserController) updateCityTour(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	var cityTour db.CityTour
	err := c.Bind(&cityTour)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = controller.model.UpdateCityTour(&cityTour, authHeader)
	if err != nil {
		var statusCode = http.StatusBadRequest
		if errors.Is(err, token_manager.InvalidTokenError{}) {
			statusCode = http.StatusUnauthorized
		}

		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func (controller UserController) updateRestaurantBooking(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	var rb db.RestaurantBooking
	err := c.Bind(&rb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = controller.model.UpdateRestaurantBooking(&rb, authHeader)
	if err != nil {
		var statusCode = http.StatusBadRequest
		if errors.Is(err, token_manager.InvalidTokenError{}) {
			statusCode = http.StatusUnauthorized
		}

		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func (controller UserController) deleteUser(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = controller.model.DeleteUser(userId, authHeader)
	if err != nil {
		var statusCode = http.StatusBadRequest
		if errors.Is(err, token_manager.InvalidTokenError{}) {
			statusCode = http.StatusUnauthorized
		}

		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func (controller UserController) deleteTour(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	tourId, err := strconv.Atoi(c.Query("tour_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = controller.model.DeleteTour(tourId, authHeader)
	if err != nil {
		var statusCode = http.StatusBadRequest
		if errors.Is(err, token_manager.InvalidTokenError{}) {
			statusCode = http.StatusUnauthorized
		}

		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func (controller UserController) deleteCityTour(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	cityTourId, err := strconv.Atoi(c.Query("city_tour_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = controller.model.DeleteCityTour(cityTourId, authHeader)
	if err != nil {
		var statusCode = http.StatusBadRequest
		if errors.Is(err, token_manager.InvalidTokenError{}) {
			statusCode = http.StatusUnauthorized
		}

		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func (controller UserController) deleteRestaurantBooking(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	rbId, err := strconv.Atoi(c.Query("restaurant_booking_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = controller.model.DeleteRestaurantBooking(rbId, authHeader)
	if err != nil {
		var statusCode = http.StatusBadRequest
		if errors.Is(err, token_manager.InvalidTokenError{}) {
			statusCode = http.StatusUnauthorized
		}

		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}
