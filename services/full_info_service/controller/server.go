package controller

import (
	"github.com/SmartDuck9000/travelly-api/services/full_info_service/config"
	"github.com/SmartDuck9000/travelly-api/services/full_info_service/model"
	"github.com/gin-gonic/gin"
)

type FullInfoControllerInterface interface {
	Run() error
}

type FullInfoController struct {
	server *gin.Engine
	model  model.FullInfoModelInterface
	host   string
	port   string
}

func CreateFullInfoController(conf config.FullInfoControllerConfig) FullInfoControllerInterface {
	gin.SetMode(gin.ReleaseMode)

	var controller = FullInfoController{
		server: gin.Default(),
		model:  model.CreateFullInfoModel(*conf.Model),
		host:   conf.Host,
		port:   conf.Port,
	}

	controller.server.GET("/api/v2/info/cities", controller.getCities)
	controller.server.GET("/api/v2/info/hotels", controller.getHotel)
	controller.server.GET("/api/v2/info/events", controller.getEvent)
	controller.server.GET("/api/v2/info/restaurants", controller.getRestaurant)
	controller.server.GET("/api/v2/info/tickets", controller.getTicket)

	return &controller
}

func (controller FullInfoController) Run() error {
	err := controller.model.Run()
	if err != nil {
		return err
	}

	err = controller.server.Run(controller.host + ":" + controller.port)

	return err
}
