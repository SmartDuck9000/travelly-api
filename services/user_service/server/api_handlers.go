package server

import (
	"github.com/SmartDuck9000/travelly-api/services/user_service/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (api UserServiceAPI) getUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		user := api.db.GetUser(userId)
		c.JSON(http.StatusOK, *user)
	}
}

func (api UserServiceAPI) getTours(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		tours := api.db.GetTours(userId)
		c.JSON(http.StatusOK, tours)
	}
}

func (api UserServiceAPI) getCityTours(c *gin.Context) {
	tourId, err := strconv.Atoi(c.Query("tour_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		cityTours := api.db.GetCityTours(tourId)
		c.JSON(http.StatusOK, cityTours)
	}
}

func (api UserServiceAPI) getCityTourEvents(c *gin.Context) {
	cityTourId, err := strconv.Atoi(c.Query("city_tour_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		events := api.db.GetEvents(cityTourId)
		c.JSON(http.StatusOK, events)
	}
}

func (api UserServiceAPI) getCityTourRestaurantBookings(c *gin.Context) {
	cityTourId, err := strconv.Atoi(c.Query("city_tour_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		restaurantBookings := api.db.GetRestaurantBookings(cityTourId)
		c.JSON(http.StatusOK, restaurantBookings)
	}
}

func (api UserServiceAPI) getCityTourTickets(c *gin.Context) {
	cityTourId, err := strconv.Atoi(c.Query("city_tour_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		tickets := api.db.GetTickets(cityTourId)
		c.JSON(http.StatusOK, tickets)
	}
}

func (api UserServiceAPI) getCityTourHotel(c *gin.Context) {
	cityTourId, err := strconv.Atoi(c.Query("city_tour_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		hotel := api.db.GetHotel(cityTourId)
		c.JSON(http.StatusOK, hotel)
	}
}

func (api UserServiceAPI) postTour(c *gin.Context) {
	var tour db.TourEntity
	err := c.Bind(&tour)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !validateTokens(c) {
		return
	}

	err = api.db.CreateTour(&tour)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func (api UserServiceAPI) postCityTour(c *gin.Context) {
	var cityTour db.CityTourEntity
	err := c.Bind(&cityTour)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !validateTokens(c) {
		return
	}

	err = api.db.CreateCityTour(&cityTour)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func (api UserServiceAPI) postRestaurantBooking(c *gin.Context) {
	var rb db.RestaurantBookingEntity
	err := c.Bind(&rb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !validateTokens(c) {
		return
	}

	err = api.db.CreateRestaurantBooking(&rb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func (api UserServiceAPI) updateUser(c *gin.Context) {
	var user db.UserEntity
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !validateTokens(c) {
		return
	}

	err = api.db.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func (api UserServiceAPI) updateTour(c *gin.Context) {
	var tour db.TourEntity
	err := c.Bind(&tour)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !validateTokens(c) {
		return
	}

	err = api.db.UpdateTour(&tour)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func (api UserServiceAPI) updateCityTour(c *gin.Context) {
	var cityTour db.CityTourEntity
	err := c.Bind(&cityTour)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !validateTokens(c) {
		return
	}

	err = api.db.UpdateCityTour(&cityTour)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func (api UserServiceAPI) updateRestaurantBooking(c *gin.Context) {
	var rb db.RestaurantBookingEntity
	err := c.Bind(&rb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !validateTokens(c) {
		return
	}

	err = api.db.UpdateRestaurantBooking(&rb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func (api UserServiceAPI) deleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !validateTokens(c) {
		return
	}

	err = api.db.DeleteUser(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func (api UserServiceAPI) deleteTour(c *gin.Context) {
	tourId, err := strconv.Atoi(c.Query("tour_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !validateTokens(c) {
		return
	}

	err = api.db.DeleteTour(tourId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func (api UserServiceAPI) deleteCityTour(c *gin.Context) {
	cityTourId, err := strconv.Atoi(c.Query("city_tour_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !validateTokens(c) {
		return
	}

	err = api.db.DeleteCityTour(cityTourId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func (api UserServiceAPI) deleteRestaurantBooking(c *gin.Context) {
	rbId, err := strconv.Atoi(c.Query("restaurant_booking_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !validateTokens(c) {
		return
	}

	err = api.db.DeleteRestaurantBooking(rbId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}
