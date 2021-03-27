package main

import (
	"github.com/SmartDuck9000/travelly-api/config_reader"
	"github.com/SmartDuck9000/travelly-api/services/user_service/config"
	"github.com/SmartDuck9000/travelly-api/services/user_service/controller"
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

	conf := config.CreateUserControllerConfig(configReader)
	var userController = controller.CreateUserController(*conf)
	err = userController.Run()

	if err != nil {
		log.Print(err.Error())
	}
}
