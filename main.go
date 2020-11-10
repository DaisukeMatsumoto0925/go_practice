package main

import (
	"go-trading/app/controllers"
	"go-trading/config"
	"go-trading/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())
}
