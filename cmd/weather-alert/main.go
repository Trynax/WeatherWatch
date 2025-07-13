package main

import (
    "fmt"
    

    "github.com/trynax/WeatherWatch/scheduler"


)


func main(){




    fmt.Println("Weather Alert Service is running...")
    fmt.Println("Press Ctrl + C to stop")

	scheduler.CheckWeatherOnSchedule(1, "Lagos")
}