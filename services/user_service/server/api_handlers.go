package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (api UserServiceAPI) getUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		user := api.db.GetUser(userId)
		c.JSON(http.StatusOK, *user)
	}
}

func (api UserServiceAPI) getTours(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		tours := api.db.GetTours(userId)
		c.JSON(http.StatusOK, tours)
	}
}

func (api UserServiceAPI) getCityTours(c *gin.Context) {
	tourId, err := strconv.Atoi(c.Query("tour_id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		cityTours := api.db.GetCityTours(tourId)
		c.JSON(http.StatusOK, cityTours)
	}
}

func (api UserServiceAPI) getCityTourEvents(c *gin.Context) {
	cityTourId, err := strconv.Atoi(c.Query("city_tour_id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		events := api.db.GetEvents(cityTourId)
		c.JSON(http.StatusOK, events)
	}
}

func (api UserServiceAPI) getCityTourRestaurantBookings(c *gin.Context) {
	cityTourId, err := strconv.Atoi(c.Query("city_tour_id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		restaurantBoolings := api.db.GetRestaurantBookings(cityTourId)
		c.JSON(http.StatusOK, restaurantBoolings)
	}
}

func (api UserServiceAPI) getCityTourTickets(c *gin.Context) {
	cityTourId, err := strconv.Atoi(c.Query("city_tour_id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		tickets := api.db.GetTickets(cityTourId)
		c.JSON(http.StatusOK, tickets)
	}
}

func (api UserServiceAPI) getCityTourHotel(c *gin.Context) {
	cityTourId, err := strconv.Atoi(c.Query("city_tour_id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		hotel := api.db.GetHotel(cityTourId)
		c.JSON(http.StatusOK, hotel)
	}
}
