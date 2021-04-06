package main

import (
	"github.com/SmartDuck9000/travelly-api/config_reader"
	"github.com/SmartDuck9000/travelly-api/services/feed_service/config"
	"github.com/SmartDuck9000/travelly-api/services/feed_service/controller"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	var err error
	var configReader config_reader.ConfigReader
	configReader, err = config_reader.CreateEnvReader(".env")

	if err != nil {
		log.Print(err.Error())
		return
	}

	conf := config.CreateFeedControllerConfig(configReader)
	var api = controller.CreateFeedController(*conf)
	err = api.Run()

	if err != nil {
		log.Print(err.Error())
	}
}
