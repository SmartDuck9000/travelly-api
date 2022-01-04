package controller

import (
	"github.com/SmartDuck9000/travelly-api/server/config"
	"github.com/SmartDuck9000/travelly-api/server/model"
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
		model:  model.CreateModel(*conf.ModelConfig),
		host:   conf.Host,
		port:   conf.Port,
	}

	// auth
	controller.server.GET("/api/v2/auth/", controller.refreshToken)
	controller.server.GET("/api/v2/auth/login", controller.login)
	controller.server.POST("/api/v2/auth/email_register", controller.register)

	// user data
	controller.server.GET("/api/v2/users/:id", controller.getUser)
	controller.server.GET("/api/v2/users/:id/tours", controller.getTours)
	controller.server.GET("/api/v2/users/:id/tours/:tour_id/city_tours", controller.getCityTours)
	controller.server.GET("/api/v2/users/:id/tours/:tour_id/city_tours/:city_tour_id/events", controller.getCityTourEvents)
	controller.server.GET("/api/v2/users/:id/tours/:tour_id/city_tours/:city_tour_id/restaurant_bookings", controller.getCityTourRestaurantBookings)
	controller.server.GET("/api/v2/users/:id/tours/:tour_id/city_tours/:city_tour_id/tickets", controller.getCityTourTickets)
	controller.server.GET("/api/v2/users/:id/tours/:tour_id/city_tours/:city_tour_id/hotels", controller.getCityTourHotel)

	controller.server.POST("/api/v2/users/:id/tours", controller.postTour)
	controller.server.POST("/api/v2/users/:id/tours/:tour_id/city_tours", controller.postCityTour)

	controller.server.PUT("/api/v2/users/:id", controller.updateUser)
	controller.server.PUT("/api/v2/users/:id/tours/:tour_id", controller.updateTour)
	controller.server.PUT("/api/v2/users/:id/tours/:tour_id/city_tours/:city_tour_id", controller.updateCityTour)

	controller.server.DELETE("/api/v2/users/:id", controller.deleteUser)
	controller.server.DELETE("/api/v2/users/:id/tours/:tour_id", controller.deleteTour)
	controller.server.DELETE("/api/v2/users/:id/tours/:tour_id/city_tours/:city_tour_id", controller.deleteCityTour)

	// info
	controller.server.GET("/api/v2/cities", controller.getCities)
	controller.server.GET("/api/v2/hotels/:id", controller.getHotel)
	controller.server.GET("/api/v2/events/:id", controller.getEvent)
	controller.server.GET("/api/v2/restaurants/:id", controller.getRestaurant)
	controller.server.GET("/api/v2/tickets/:id", controller.getTicket)

	controller.server.GET("/api/v2/hotels", controller.getHotels)
	controller.server.GET("/api/v2/events", controller.getEvents)
	controller.server.GET("/api/v2/restaurants", controller.getRestaurants)
	controller.server.GET("/api/v2/tickets", controller.getTickets)

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
