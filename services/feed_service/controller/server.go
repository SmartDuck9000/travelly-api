package controller

import (
	"github.com/SmartDuck9000/travelly-api/services/feed_service/config"
	"github.com/SmartDuck9000/travelly-api/services/feed_service/model"
	"github.com/gin-gonic/gin"
)

type FeedControllerInterface interface {
	Run() error
}

type FeedController struct {
	server *gin.Engine
	model  model.FeedModelInterface
	host   string
	port   string
}

func CreateFeedController(conf config.FeedControllerConfig) FeedControllerInterface {
	var controller = FeedController{
		server: gin.Default(),
		model:  model.CreateFeedModel(*conf.ModelConfig),
		host:   conf.Host,
		port:   conf.Port,
	}

	controller.server.GET("/api/feed/hotels", controller.getHotels)
	controller.server.GET("/api/feed/events", controller.getEvents)
	controller.server.GET("/api/feed/restaurants", controller.getRestaurants)
	controller.server.GET("/api/feed/tickets", controller.getTickets)

	return &controller
}

func (controller FeedController) Run() error {
	err := controller.model.Run()
	if err != nil {
		return err
	}

	err = controller.server.Run(controller.host + ":" + controller.port)

	return err
}
