package server

import (
	"github.com/SmartDuck9000/travelly-api/services/feed_service/config"
	"github.com/SmartDuck9000/travelly-api/services/feed_service/db"
	"github.com/gin-gonic/gin"
)

type UserServiceInterface interface {
	Run() error
}

type FeedServiceAPI struct {
	server *gin.Engine
	db     db.FeedDB
	host   string
	port   string
}

func CreateServer(conf config.FeedServiceConfig) *FeedServiceAPI {
	var api = FeedServiceAPI{
		server: gin.Default(),
		db:     db.CreateFeedServiceDB(conf.DB),
		host:   conf.Host,
		port:   conf.Port,
	}

	api.server.GET("/api/feed/hotels", api.getHotels)
	api.server.GET("/api/feed/events", api.getEvents)
	api.server.GET("/api/feed/restaurants", api.getRestaurants)

	return &api
}

func (api FeedServiceAPI) Run() error {
	err := api.db.Open()
	if err != nil {
		return err
	}

	err = api.server.Run(api.host + ":" + api.port)

	return err
}
