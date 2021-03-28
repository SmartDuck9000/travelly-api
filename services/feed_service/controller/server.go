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

	controller.server.GET("/controller/feed/hotels", controller.getHotels)
	controller.server.GET("/controller/feed/events", controller.getEvents)
	controller.server.GET("/controller/feed/restaurants", controller.getRestaurants)

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
