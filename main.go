package main

import (
	"fmt"
	"github.com/jinzhu/configor"
	"time"
)

func main() {
	_ = configor.Load(&Config, "config.yml")
	fmt.Printf("%s", Config.HubIp)
	a := allSadeData(Config.HubIp)
	fmt.Println(a)
	a = allSceneData(Config.HubIp)
	fmt.Println(a)

}

func main3() {
	_ = configor.Load(&Config, "config.yml")
	ticker := time.NewTicker(15 * time.Minute)
	fmt.Println("Starting weather monitoring...")
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				currentWeather(Config.ApiKey)
			}
		}
	}()

	time.Sleep(24 * time.Hour)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")

}

func morningTrigger() {
	// start a sunrise, check weather, if !cloudy close shades, end
	// if cloudy check every 15 mins.  If !cloudy close shades, end
	// stop ticker at time, open shades
}

var logger = configureLogging()

var Config = struct {
	APPName string `default:"shaderrific"`
	ApiKey  string
	HubIp   string
}{}
