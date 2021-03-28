package controller

import (
	"github.com/SmartDuck9000/travelly-api/services/user_service/config"
	"github.com/SmartDuck9000/travelly-api/services/user_service/model"
	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	Run() error
}

type UserController struct {
	server *gin.Engine
	model  model.UserModelInterface
	host   string
	port   string
}

func CreateUserController(conf config.UserControllerConfig) UserControllerInterface {
	gin.SetMode(gin.ReleaseMode)

	var controller = UserController{
		server: gin.Default(),
		model:  model.CreateUserModel(*conf.ModelConfig),
		host:   conf.Host,
		port:   conf.Port,
	}

	controller.server.GET("/controller/users", controller.getUser)
	controller.server.GET("/controller/users/tours", controller.getTours)
	controller.server.GET("/controller/users/tours/city_tours", controller.getCityTours)
	controller.server.GET("/controller/users/tours/city_tours/events", controller.getCityTourEvents)
	controller.server.GET("/controller/users/tours/city_tours/restaurant_bookings", controller.getCityTourRestaurantBookings)
	controller.server.GET("/controller/users/tours/city_tours/tickets", controller.getCityTourTickets)
	controller.server.GET("/controller/users/tours/city_tours/hotels", controller.getCityTourHotel)

	controller.server.POST("/controller/users/tours", controller.postTour)
	controller.server.POST("/controller/users/city_tours", controller.postCityTour)
	controller.server.POST("/controller/users/restaurant_bookings", controller.postRestaurantBooking)

	controller.server.PUT("/controller/users", controller.updateUser)
	controller.server.PUT("/controller/users/tours", controller.updateTour)
	controller.server.PUT("/controller/users/city_tours", controller.updateCityTour)
	controller.server.PUT("/controller/users/restaurant_bookings", controller.updateRestaurantBooking)

	controller.server.DELETE("/controller/users", controller.deleteUser)
	controller.server.DELETE("/controller/users/tours", controller.deleteTour)
	controller.server.DELETE("/controller/users/city_tours", controller.deleteCityTour)
	controller.server.DELETE("/controller/users/restaurant_bookings", controller.deleteRestaurantBooking)

	return &controller
}

func (controller UserController) Run() error {
	err := controller.model.Run()
	if err != nil {
		return err
	}

	err = controller.server.Run(controller.host + ":" + controller.port)

	return err
}
