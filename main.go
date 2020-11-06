package main

import (
	"fmt"
	"go-trading/bitflyer"
	"go-trading/config"
	"go-trading/utils"
	"time"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	// fmt.Println(apiClient.GetBalance())
	ticker, _ := apiClient.GetTicker("BTC_JPY")
	fmt.Println(ticker.DateTime())
	fmt.Println(ticker.TruncateDateTime(time.Hour))
}
