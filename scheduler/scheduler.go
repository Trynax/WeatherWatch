package scheduler

import (
	"fmt"
	"time"
	"os"
    "os/signal"
    "syscall"
	"github.com/trynax/WeatherWatch/api"
	"github.com/trynax/WeatherWatch/utils"
	"github.com/trynax/WeatherWatch/notify"
)




func CheckWeatherOnSchedule(timeInMins int, city string) {

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	ticker := time.NewTicker(time.Duration(timeInMins) * time.Minute)


	for {
		select{
		case <- ticker.C:
			result, err:= api.GetCurrentWeatherWithCityName(city)

			if err !=nil{
				fmt.Println("Error while caling: ", err)
			}else{
				result := utils.FormatSummary(result)
				fmt.Println(result)
				notify.CreateAndAddContents(result)
			}
		case <- sigChan:
			fmt.Println("Shuttting down....")
			return
		
	}

	}

}