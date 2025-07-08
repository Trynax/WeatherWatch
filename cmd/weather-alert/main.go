package main 


import (
	"fmt"
	"github.com/trynax/WeatherWatch/api"
	"github.com/trynax/WeatherWatch/utils"
)



func main(){


	fmt.Println("Weather Alert Service is running...")

	result, err := api.GetCurrentWeatherWithCityName("Lagos")

	if err !=nil{
		fmt.Println("Error: ", err)
	}



	fmt.Println(utils.WeatherCurrentResponse(result))
}