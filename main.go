package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/jinzhu/configor"
	"time"
)

var sunrise time.Time
var sunset time.Time

func main() {
	_ = configor.Load(&Config, "config.yml")
	setupDay()
	gocron.Every(1).Day().At("4:30").Do(setupDay)

	//morningStartTime := time.Now().Add(time.Duration(1) * time.Minute)

	<-gocron.Start()

}

func setupDay() {
	logger.Info("Setting up for the day")
	w := currentWeather(Config.ApiKey)
	sunrise = w.Sunrise
	sunset = w.Sunset
	morningStartTime := sunrise.Add(time.Duration(1) * time.Hour)
	afternoonStartTime := sunset.Add(time.Duration(-6) * time.Hour)
	openFrontTime := sunrise.Add(time.Duration(4) * time.Hour)
	openBackTime := sunset.Add(time.Duration(-30) * time.Minute)
	gocron.Every(15).Minutes().From(&morningStartTime).Do(morningTask)
	gocron.Every(15).Minutes().From(&afternoonStartTime).Do(afternoonTask)

	gocron.Every(1).Day().At(fmt.Sprintf("%02d:%02d", openFrontTime.Hour(), openFrontTime.Minute())).Do(openFront)
	gocron.Every(1).Day().At(fmt.Sprintf("%02d:%02d", openBackTime.Hour(), openBackTime.Minute())).Do(openBack)
}

func morningTask() {
	if time.Now().After(sunrise.Add(time.Duration(3) * time.Hour)) {
		fmt.Println("Ending Morning Task")
		logger.Info("Ending Morning Task: Time Expired")
		gocron.Remove(morningTask)
		return
	}

	w := currentWeather(Config.ApiKey)
	if w.Clouds < 70 {
		allSceneData(Config.HubIp)
		logger.Info("Ending Morning Task: Closed shades")
		gocron.Remove(morningTask)
		return
	}
}

func afternoonTask() {
	if time.Now().After(sunset.Add(time.Duration(1) * time.Hour)) {
		fmt.Println("Ending  Task")
		logger.Info("Ending Afternoon Task: Time Expired")
		gocron.Remove(afternoonTask)
	}

	w := currentWeather(Config.ApiKey)
	if w.Clouds < 70 {
		allSceneData(Config.HubIp)
		logger.Info("Ending Afternoon Task: Closed shades")
		gocron.Remove(afternoonTask)
		return
	}
}

func openFront() {
	logger.Info("Opening front shades")
	allSceneData(Config.HubIp)
	gocron.Remove(openFront)
}

func openBack() {
	logger.Info("Opening back shades")
	allSceneData(Config.HubIp)
	gocron.Remove(openBack)
}

var logger = configureLogging()

var Config = struct {
	APPName string `default:"shaderrific"`
	ApiKey  string
	HubIp   string
}{}
