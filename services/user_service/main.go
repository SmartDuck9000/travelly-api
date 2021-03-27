package main

import (
	"github.com/SmartDuck9000/travelly-api/config_reader"
	"github.com/SmartDuck9000/travelly-api/services/user_service/config"
	"github.com/SmartDuck9000/travelly-api/services/user_service/server"
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

	conf := config.New(configReader)
	var api server.UserServiceInterface = server.CreateServer(*conf)
	err = api.Run()

	if err != nil {
		log.Print(err.Error())
	}
}
