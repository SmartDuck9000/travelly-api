package main

import (
	"github.com/SmartDuck9000/travelly-api/config_reader"
	"github.com/SmartDuck9000/travelly-api/services/auth_service/config"
	"github.com/SmartDuck9000/travelly-api/services/auth_service/controller"
	"log"
)

func main() {
	var err error
	var configReader config_reader.ConfigReader
	configReader, err = config_reader.CreateEnvReader(".env")

	if err != nil {
		log.Print(err.Error())
		return
	}

	conf := config.CreateAuthControllerConfig(configReader)
	var authController = controller.CreateAuthController(*conf)
	err = authController.Run()

	if err != nil {
		log.Print(err.Error())
	}
}
