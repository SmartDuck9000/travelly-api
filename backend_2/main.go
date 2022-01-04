package backend_2

import (
	"github.com/SmartDuck9000/travelly-api/config_reader"
	"github.com/SmartDuck9000/travelly-api/server/config"
	"github.com/SmartDuck9000/travelly-api/server/controller"
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

	conf := config.CreateControllerConfig(configReader)
	var mainController = controller.CreateController(*conf)
	err = mainController.Run()

	if err != nil {
		log.Print(err.Error())
	}
}
