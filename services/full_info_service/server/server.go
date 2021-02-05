package server

import (
	"github.com/SmartDuck9000/travelly-api/services/full_info_service/config"
	"github.com/SmartDuck9000/travelly-api/services/full_info_service/db"
	"github.com/gin-gonic/gin"
)

type FullInfoInterface interface {
	Run() error
}

type FullInfoAPI struct {
	server *gin.Engine
	db     db.FullInfoDb
	host   string
	port   string
}

func CreateServer(conf config.FullInfoConfig) *FullInfoAPI {
	var api = FullInfoAPI{
		server: gin.Default(),
		db:     db.CreateFullInfoDB(conf.DB),
		host:   conf.Host,
		port:   conf.Port,
	}

	api.server.GET("/api/info/hotels", api.getHotel)
	api.server.GET("/api/info/events", api.getEvent)
	api.server.GET("/api/info/restaurants", api.getRestaurant)

	return &api
}

func (api FullInfoAPI) Run() error {
	err := api.db.Open()
	if err != nil {
		return err
	}

	err = api.server.Run(api.host + ":" + api.port)

	return err
}
