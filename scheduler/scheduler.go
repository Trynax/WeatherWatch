package scheduler

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/trynax/WeatherWatch/alert"
	"github.com/trynax/WeatherWatch/api"
	"github.com/trynax/WeatherWatch/notify"
	"github.com/trynax/WeatherWatch/utils"
)

func CheckWeatherOnSchedule(timeInMins int, city string) {

	sigChan := make(chan os.Signal, 1)
	firstTime := true

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	ticker := time.NewTicker(time.Duration(timeInMins) * time.Minute)

	for {
		select {
		case <-ticker.C:
			result, err := api.GetCurrentWeatherWithCityName(city)

			if err != nil {
				fmt.Println("Error while calling API:", err)
			} else {

				alerts := alert.CheckAllConditions(result)

				if firstTime {
			
					summary := utils.FormatSummary(result)
					fmt.Println("\n" + summary)
					notify.CreateAndAddContents(summary)
					firstTime = false
				} else {
	
					if len(alerts) > 0 {
			
						alertMsg := utils.FormatAlertMessage(
							result.Location.Name+", "+result.Location.Country,
							result.Current.LastUpdated,
							alerts,
						)
						notify.PrintAlert(alertMsg)
						notify.CreateAndAddContents(alertMsg)
					} else {
						// No alerts - show brief OK status
						okMsg := utils.FormatNoAlertMessage(*result)
						notify.PrintStatus(okMsg)

						// Log entry for file
						logEntry := utils.FormatLogEntry(*result, alerts)
						notify.CreateAndAddContents(logEntry)
					}
				}
			}
		case <-sigChan:
			fmt.Println("Shuttting down....")
			return

		}

	}

}
