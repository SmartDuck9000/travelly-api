package controller

import (
	"github.com/SmartDuck9000/travelly-api/services/user_service/config"
	"github.com/SmartDuck9000/travelly-api/services/user_service/model"
	"github.com/gin-gonic/gin"
)

type ControllerInterface interface {
	Run() error
}

type Controller struct {
	server *gin.Engine
	model  model.ModelInterface
	host   string
	port   string
}

func CreateController(conf config.ControllerConfig) ControllerInterface {
	gin.SetMode(gin.ReleaseMode)

	var controller = Controller{
		server: gin.Default(),
		model:  model.CreateUserModel(*conf.ModelConfig),
		host:   conf.Host,
		port:   conf.Port,
	}

	controller.server.GET("/api/v2/users", controller.getUser)
	controller.server.GET("/api/v2/users/tours", controller.getTours)
	controller.server.GET("/api/v2/users/tours/city_tours", controller.getCityTours)
	controller.server.GET("/api/v2/users/tours/city_tours/events", controller.getCityTourEvents)
	controller.server.GET("/api/v2/users/tours/city_tours/restaurant_bookings", controller.getCityTourRestaurantBookings)
	controller.server.GET("/api/v2/users/tours/city_tours/tickets", controller.getCityTourTickets)
	controller.server.GET("/api/v2/users/tours/city_tours/hotels", controller.getCityTourHotel)

	controller.server.POST("/api/v2/users/tours", controller.postTour)
	controller.server.POST("/api/v2/users/city_tours", controller.postCityTour)
	controller.server.POST("/api/v2/users/city_tour_events", controller.postCityTourEvent)
	controller.server.POST("/api/v2/users/restaurant_bookings", controller.postRestaurantBooking)

	controller.server.PUT("/api/v2/users", controller.updateUser)
	controller.server.PUT("/api/v2/users/tours", controller.updateTour)
	controller.server.PUT("/api/v2/users/city_tours", controller.updateCityTour)
	controller.server.PUT("/api/v2/users/restaurant_bookings", controller.updateRestaurantBooking)

	controller.server.DELETE("/api/v2/users", controller.deleteUser)
	controller.server.DELETE("/api/v2/users/tours", controller.deleteTour)
	controller.server.DELETE("/api/v2/users/city_tours", controller.deleteCityTour)
	controller.server.DELETE("/api/v2/users/restaurant_bookings", controller.deleteRestaurantBooking)

	return &controller
}

func (controller Controller) Run() error {
	err := controller.model.Run()
	if err != nil {
		return err
	}

	err = controller.server.Run(controller.host + ":" + controller.port)

	return err
}
