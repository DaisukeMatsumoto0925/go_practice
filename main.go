package main

import (
	"go-trading/utils"
	"log"

	"go-trading/config"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
}
