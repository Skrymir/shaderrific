package main

import (
	"encoding/json"
	"time"
)

type weather struct {
	Context    []interface{} `json:"@context"`
	Type       string        `json:"type"`
	Geometry   Geometry      `json:"geometry"`
	Properties Properties    `json:"properties"`
}
type Geometries struct {
	Type        string    `json:"type"`
	Coordinates []json.RawMessage `json:"coordinates"`
}
type Geometry struct {
	Type       string       `json:"type"`
	Geometries []Geometries `json:"geometries"`
}
type Elevation struct {
	Value    float64 `json:"value"`
	UnitCode string  `json:"unitCode"`
}
type Periods struct {
	Number           int         `json:"number"`
	Name             string      `json:"name"`
	StartTime        string      `json:"startTime"`
	EndTime          string      `json:"endTime"`
	IsDaytime        bool        `json:"isDaytime"`
	Temperature      int         `json:"temperature"`
	TemperatureUnit  string      `json:"temperatureUnit"`
	TemperatureTrend interface{} `json:"temperatureTrend"`
	WindSpeed        string      `json:"windSpeed"`
	WindDirection    string      `json:"windDirection"`
	Icon             string      `json:"icon"`
	ShortForecast    string      `json:"shortForecast"`
	DetailedForecast string      `json:"detailedForecast"`
}
type Properties struct {
	Updated           time.Time `json:"updated"`
	Units             string    `json:"units"`
	ForecastGenerator string    `json:"forecastGenerator"`
	GeneratedAt       time.Time `json:"generatedAt"`
	UpdateTime        time.Time `json:"updateTime"`
	ValidTimes        string `json:"validTimes"`
	Elevation         Elevation `json:"elevation"`
	Periods           []Periods `json:"periods"`
}
