package controller

import (
	"github.com/SmartDuck9000/travelly-api/services/auth_service/config"
	"github.com/SmartDuck9000/travelly-api/services/auth_service/model"
	"github.com/gin-gonic/gin"
)

type AuthControllerInterface interface {
	Run() error
}

type AuthController struct {
	server *gin.Engine
	model  model.AuthModelInterface
	host   string
	port   string
}

func CreateAuthController(conf config.AuthControllerConfig) AuthControllerInterface {
	gin.SetMode(gin.ReleaseMode)

	var controller = AuthController{
		server: gin.Default(),
		model:  model.CreateAuthModel(*conf.ModelConfig),
		host:   conf.Host,
		port:   conf.Port,
	}

	controller.server.GET("/api/auth/", controller.refreshToken)
	controller.server.POST("/api/auth/email_register", controller.register)
	controller.server.POST("/api/auth/login", controller.login)

	return &controller
}

func (controller AuthController) Run() error {
	err := controller.model.Run()
	if err != nil {
		return err
	}

	err = controller.server.Run(controller.host + ":" + controller.port)

	return err
}
