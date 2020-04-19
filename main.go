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
	err := gocron.Every(1).Day().At("4:30").Do(setupDay)
	if err != nil {

	}

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
	openFrontTime := sunrise.Add(time.Duration(6) * time.Hour)
	openBackTime := sunset.Add(time.Duration(-10) * time.Minute)
	err := gocron.Every(15).Minutes().From(&morningStartTime).Do(morningTask)
	if err != nil {
		logger.Warn("Error in morning task", err)
	}
	err = gocron.Every(15).Minutes().From(&afternoonStartTime).Do(afternoonTask)
	if err != nil {
		logger.Warn("Error in afternoon task", err)
	}
	err = gocron.Every(1).Day().At(fmt.Sprintf("%02d:%02d", openFrontTime.Hour(), openFrontTime.Minute())).Do(openFront)
	if err != nil {
		logger.Warn("Error in opening front task", err)
	}
	err = gocron.Every(1).Day().At(fmt.Sprintf("%02d:%02d", openBackTime.Hour(), openBackTime.Minute())).Do(openBack)
	if err != nil {
		logger.Warn("Error in opening back task", err)
	}
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
		frontClose()
		fmt.Println("Ending Morning Task: Closed shades")
		logger.Info("Ending Morning Task: Closed shades")
		gocron.Remove(morningTask)
		return
	}
	fmt.Println("Morning task run: noop")
}

func afternoonTask() {
	if time.Now().After(sunset.Add(time.Duration(-1) * time.Hour)) {
		fmt.Println("Ending Afternoon Task: Time Expired")
		logger.Info("Ending Afternoon Task: Time Expired")
		gocron.Remove(afternoonTask)
		return
	}

	w := currentWeather(Config.ApiKey)
	if w.Clouds < 70 {
		backClose()
		fmt.Println("Ending Afternoon Task: Closed shades")
		logger.Info("Ending Afternoon Task: Closed shades")
		gocron.Remove(afternoonTask)
		return
	}
	fmt.Println("Afternoon task run: noop")
}

func openFront() {
	fmt.Println("Opening front shades")
	logger.Info("Opening front shades")
	frontOpen()
	gocron.Remove(openFront)
}

func openBack() {
	fmt.Println("Opening back shades")
	logger.Info("Opening back shades")
	backOpen()
	gocron.Remove(openBack)
}

var logger = configureLogging()

var Config = struct {
	APPName string `default:"shaderrific"`
	ApiKey  string
	HubIp   string
}{}
