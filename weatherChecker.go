package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func checkWeather() {
	url := "https://api.weather.gov/gridpoints/MPX/102,68/forecast/hourly"
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
    var bodyBytes []byte
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		logger.Info(bodyString)
		fmt.Println(bodyString)
	}

	var a weather
	err = json.Unmarshal(bodyBytes, &a)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(a.Context, a.Properties.Periods)
}


