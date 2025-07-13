package models



type WeatherResponse struct{
	Location Location `json:"location"`
	Current  Current  `json:"current"`
	
}


type Location struct {
	Name string `json:"name"`
	Country string `json:"country"`
	Latitude float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Localtime string `json:"localtime"`
}

type Current struct {
	LastUpdated string `json:"last_updated"`
	TempC float64 `json:"temp_c"`
	IsDay int     `json:"is_day"`
	Condition Condition `json:"condition"`
	WindKph float64 `json:"wind_kph"`
	WindDeg int    `json:"wind_deg"`
	WindDir string `json:"wind_dir"`
	Humidity int    `json:"humidity"`
	Cloud int    `json:"cloud"`
	FeelsLikeC float32 `json:"feelslike_c"`
}

type Condition struct {
	Text string `json:"text"`
	Code int    `json:"code"`
	Icon string `json:"icon"`
}


type WeatherRequest struct {
	City string `json:"city"`
}

type WeatherError struct {
	Message string `json:"message"`
	Status int    `json:"status"`
}


type Threshold struct {
	Temp float64
	WindSpeed float64
	Humidity int
	Cloud int
}