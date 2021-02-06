package server

import (
	"github.com/SmartDuck9000/travelly-api/services/auth_service/config"
	"github.com/SmartDuck9000/travelly-api/services/auth_service/db"
	"github.com/gin-gonic/gin"
)

type AuthInterface interface {
	Run() error
}

type AuthAPI struct {
	server *gin.Engine
	db     db.AuthDB
	host   string
	port   string
}

func CreateServer(conf config.AuthServiceConfig) *AuthAPI {
	var api = AuthAPI{
		server: gin.Default(),
		db:     db.CreateAuthDB(conf.DB),
		host:   conf.Host,
		port:   conf.Port,
	}

	api.server.POST("/api/auth/email_register", api.register)
	api.server.POST("/api/auth/login", api.login)

	return &api
}

func (api AuthAPI) Run() error {
	err := api.db.Open()
	if err != nil {
		return err
	}

	err = api.server.Run(api.host + ":" + api.port)

	return err
}
