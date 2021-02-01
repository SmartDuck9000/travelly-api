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
	db     db.TravellyDb
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
