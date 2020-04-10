package main

import (
	"fmt"
	owm "github.com/briandowns/openweathermap"
	"time"
)

func currentWeather(apiKey string) weatherData {
	w, err := owm.NewCurrent("F", "EN", apiKey)
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

	var uv *owm.UV
	uv, err = owm.NewUV(apiKey)
	if err != nil {
		logger.Fatal(err)
	}

	if err = uv.Current(coord); err != nil {
		logger.Fatal(err)
	}

	wd := weatherData{
		Sunrise:     time.Unix(int64(w.Sys.Sunrise), 0),
		Sunset:      time.Unix(int64(w.Sys.Sunset), 0),
		Temperature: float32(w.Main.Temp),
		Conditions:  w.Weather[0].Main,
		UVIndex:     float32(uv.Value),
		Clouds:      w.Clouds.All,
	}

	fmt.Printf("%+v\n", wd)
	logger.Infof("%+v\n", wd)
	logger.Infof("%+v\n", w)

	return wd
}

type weatherData struct {
	Sunrise     time.Time
	Sunset      time.Time
	Temperature float32
	Conditions  string
	UVIndex     float32
	Clouds      int
}
