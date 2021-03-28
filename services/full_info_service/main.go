package main

import (
	"github.com/SmartDuck9000/travelly-api/config_reader"
	"github.com/SmartDuck9000/travelly-api/services/full_info_service/config"
	"github.com/SmartDuck9000/travelly-api/services/full_info_service/controller"
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

	conf := config.CreateFullInfoControllerConfig(configReader)
	var fullInfoController = controller.CreateFullInfoController(*conf)
	err = fullInfoController.Run()

	if err != nil {
		log.Print(err.Error())
	}
}
