package main

import (
	"github.com/SmartDuck9000/travelly-api/services/full_info_service/config"
	"github.com/SmartDuck9000/travelly-api/services/full_info_service/controller"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config.CreateFullInfoControllerConfig()
	var fullInfoController = controller.CreateFullInfoController(*conf)
	err := fullInfoController.Run()

	if err != nil {
		log.Print(err.Error())
	}
}
