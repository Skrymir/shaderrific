package main

import (
	"fmt"
	owm "github.com/briandowns/openweathermap"
	"github.com/jinzhu/configor"
)

func main() {
	_ = configor.Load(&Config, "config.yml")
	w1(Config.ApiKey)
}

var logger = configureLogging()

func w1(apiKey string) {
	w, err := owm.NewCurrent("F", "EN", apiKey) // (internal - OpenWeatherMap reference for kelvin) with English output
	if err != nil {
		logger.Fatal(err)
	}

	coord := &owm.Coordinates{
		Latitude:  44.9041,
		Longitude: -93.4561,
	}

	if err = w.CurrentByCoordinates(coord); err != nil {
		logger.Fatal(err)
	}

	fmt.Printf("%+v\n", w)

	uv, err := owm.NewUV(apiKey)
	if err != nil {
		logger.Fatal(err)
	}

	if err = uv.Current(coord); err != nil {
		logger.Fatal(err)
	}

	fmt.Printf("%+v\n", uv)

}

var Config = struct {
	APPName string `default:"twcPoller"`
	ApiKey  string
}{}
