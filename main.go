package main

import (
	"go-trading/app/controllers"
	"go-trading/config"
	"go-trading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}
