package models



type WeatherResponse struct{
	Location Location `json:"location"`
	Current  Current  `json:"current"`
	WindKph float64 `json:"wind_kph"`
	WindDeg int    `json:"wind_deg"`
	WindDir string `json:"wind_dir"`
	Humidity int    `json:"humidity"`
	Cloud int    `json:"cloud"`
}


type Location struct {
	Name string `json:"name"`
	Country string `json:"country"`
	Latitude float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Localtime string `json:"localtime"`
}

type Current struct {
	lastUpdated string `json:"last_updated"`
	TempC float64 `json:"temp_c"`
	IsDay int     `json:"is_day"`
	Condition Condition `json:"condition"`
}

type Condition struct {
	Text string `json:"text"`
	Code int    `json:"code"`
	Icon string `json:"icon"`
}