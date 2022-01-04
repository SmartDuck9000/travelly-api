package controller

import (
	"errors"
	"github.com/SmartDuck9000/travelly-api/server/db"
	"github.com/SmartDuck9000/travelly-api/server/model"
	"github.com/SmartDuck9000/travelly-api/token_manager"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// auth

func (controller Controller) register(c *gin.Context) {
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

func (controller Controller) login(c *gin.Context) {
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

func (controller Controller) refreshToken(c *gin.Context) {
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

// user data

func (controller Controller) getUser(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	userId, err := strconv.Atoi(c.Param("id"))
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

func (controller Controller) getTours(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	userId, err := strconv.Atoi(c.Param("id"))
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

func (controller Controller) getCityTours(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	tourId, err := strconv.Atoi(c.Param("tour_id"))
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

func (controller Controller) getCityTourEvents(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	cityTourId, err := strconv.Atoi(c.Param("city_tour_id"))
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

func (controller Controller) getCityTourRestaurantBookings(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	cityTourId, err := strconv.Atoi(c.Param("city_tour_id"))
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

func (controller Controller) getCityTourTickets(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	cityTourId, err := strconv.Atoi(c.Param("city_tour_id"))
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

func (controller Controller) getCityTourHotel(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	cityTourId, err := strconv.Atoi(c.Param("city_tour_id"))
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

func (controller Controller) postTour(c *gin.Context) {
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

func (controller Controller) postCityTour(c *gin.Context) {
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

func (controller Controller) postCityTourEvent(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	var cityTourEvent db.CityToursEvent
	err := c.Bind(&cityTourEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = controller.model.CreateCityTourEvent(&cityTourEvent, authHeader)
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

func (controller Controller) postRestaurantBooking(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	var rb db.RestaurantBookingDTO
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

func (controller Controller) updateUser(c *gin.Context) {
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

func (controller Controller) updateTour(c *gin.Context) {
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

func (controller Controller) updateCityTour(c *gin.Context) {
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

func (controller Controller) updateRestaurantBooking(c *gin.Context) {
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

func (controller Controller) deleteUser(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	userId, err := strconv.Atoi(c.Param("id"))
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

func (controller Controller) deleteTour(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	tourId, err := strconv.Atoi(c.Param("tour_id"))
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

func (controller Controller) deleteCityTour(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header",
		})
		return
	}

	cityTourId, err := strconv.Atoi(c.Param("city_tour_id"))
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

func (controller Controller) deleteRestaurantBooking(c *gin.Context) {
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

// info

func (controller Controller) getCities(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no authorization header",
		})
		return
	}

	cities, err := controller.model.GetCities(authHeader)

	if err != nil {
		var statusCode = http.StatusBadRequest
		if errors.Is(err, token_manager.InvalidTokenError{}) {
			statusCode = http.StatusUnauthorized
		}

		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, cities)
	}
}

func (controller Controller) getHotel(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no authorization header",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
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

func (controller Controller) getEvent(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no authorization header",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
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

func (controller Controller) getRestaurant(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no authorization header",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
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

func (controller Controller) getTicket(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no authorization header",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
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

// feed

func (controller Controller) getHotels(c *gin.Context) {
	var filterParameters db.HotelFilterParameters
	var hotels []db.HotelFeedItem

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

func (controller Controller) getEvents(c *gin.Context) {
	var filterParameters db.EventsFilterParameters
	var events []db.EventFeedItem

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

func (controller Controller) getRestaurants(c *gin.Context) {
	var filterParameters db.RestaurantFilterParameters
	var restaurants []db.RestaurantFeedItem

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

func (controller Controller) getTickets(c *gin.Context) {
	var filterParameters db.TicketFilterParameters
	var tickets []db.TicketFeedItem

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
