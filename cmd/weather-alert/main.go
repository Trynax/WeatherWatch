package main 


import (
	"fmt"
	"github.com/trynax/WeatherWatch/api"
)



func main(){


	fmt.Println("Weather Alert Service is running...")

	fmt.Println(api.GetCurrrentWeather("Lagos"))
}