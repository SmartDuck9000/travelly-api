package main

import (
	"github.com/SmartDuck9000/travelly-api/services/feed_service/config"
	"github.com/SmartDuck9000/travelly-api/services/feed_service/server"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config.New()
	var api server.FeedServiceInterface = server.CreateServer(*conf)
	err := api.Run()

	if err != nil {
		log.Print(err.Error())
	}
}
