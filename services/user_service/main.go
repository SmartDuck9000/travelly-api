package main

import (
	"github.com/SmartDuck9000/travelly-api/services/user_service/config"
	"github.com/SmartDuck9000/travelly-api/services/user_service/db"
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
	var userDb db.TravellyDb = db.CreateUserServiceDb(conf.DB)

	err := userDb.Open()
	if err != nil {
		log.Print(err.Error())
	}
}
