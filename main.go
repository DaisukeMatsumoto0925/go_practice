package main

import (
	"fmt"
	"go-trading/bitflyer"
	"go-trading/utils"

	"go-trading/config"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println((apiClient.GetBalance()))
}
