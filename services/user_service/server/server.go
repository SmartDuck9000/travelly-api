package server

import (
	"github.com/SmartDuck9000/travelly-api/services/user_service/config"
	"github.com/SmartDuck9000/travelly-api/services/user_service/db"
	"github.com/gin-gonic/gin"
)

type UserServiceInterface interface {
	Run() error
}

type UserServiceAPI struct {
	server *gin.Engine
	db     db.UserProfileDB
	host   string
	port   string
}

func CreateServer(conf config.UserServiceConfig) *UserServiceAPI {
	var api = UserServiceAPI{
		server: gin.Default(),
		db:     db.CreateUserServiceDb(conf.DB),
		host:   conf.Host,
		port:   conf.Port,
	}

	api.server.GET("/api/users", api.getUser)
	api.server.GET("/api/users/tours", api.getTours)
	api.server.GET("/api/users/tours/city_tours", api.getCityTours)
	api.server.GET("/api/users/tours/city_tours/events", api.getCityTourEvents)
	api.server.GET("/api/users/tours/city_tours/restaurant_bookings", api.getCityTourRestaurantBookings)
	api.server.GET("/api/users/tours/city_tours/tickets", api.getCityTourTickets)
	api.server.GET("/api/users/tours/city_tours/hotels", api.getCityTourHotel)

	api.server.POST("/api/users/tours", api.postTour)
	api.server.POST("/api/users/city_tours", api.postCityTour)
	api.server.POST("/api/users/restaurant_bookings", api.postRestaurantBooking)

	api.server.PUT("/api/users", api.updateUser)
	api.server.PUT("/api/users/tours", api.updateTour)
	api.server.PUT("/api/users/city_tours", api.updateCityTour)
	api.server.PUT("/api/users/restaurant_bookings", api.updateRestaurantBooking)

	api.server.DELETE("/api/users", api.deleteUser)
	api.server.DELETE("/api/users/tours", api.deleteTour)
	api.server.DELETE("/api/users/city_tours", api.deleteCityTour)
	api.server.DELETE("/api/users/restaurant_bookings", api.deleteRestaurantBooking)

	return &api
}

func (api UserServiceAPI) Run() error {
	err := api.db.Open()
	if err != nil {
		return err
	}

	err = api.server.Run(api.host + ":" + api.port)

	return err
}
