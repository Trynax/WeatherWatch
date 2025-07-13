package alert

import (
	"fmt"

	"github.com/trynax/WeatherWatch/models"
)

var defaultThreshold = models.Threshold{
	Temp:      28.0,
	WindSpeed: 22.3,
	Humidity:  90,
	Cloud:     55,
}

func CheckIfTempIsHigher(temp float64) bool {
	return temp >= defaultThreshold.Temp
}

func CheckIfWindSpeedIsHigher(windSpeed float64) bool {
	return windSpeed >= defaultThreshold.WindSpeed
}

func CheckIfHumidityIsHigher(humidity int) bool {
	return humidity >= defaultThreshold.Humidity
}

func CheckIfCloudCoverIsHigher(cloud int) bool {
	return cloud >= defaultThreshold.Cloud
}

func CheckAllConditions(weather *models.WeatherResponse) []string {
	var alerts []string

	if CheckIfTempIsHigher(weather.Current.TempC) {
		alerts = append(alerts, fmt.Sprintf("🌡️ High Temperature: %.1f°C (threshold: %.1f°C)",
			weather.Current.TempC, defaultThreshold.Temp))
	}

	if CheckIfWindSpeedIsHigher(weather.Current.WindKph) {
		alerts = append(alerts, fmt.Sprintf("🌬️ High Wind Speed: %.1f kph (threshold: %.1f kph)",
			weather.Current.WindKph, defaultThreshold.WindSpeed))
	}

	if CheckIfHumidityIsHigher(weather.Current.Humidity) {
		alerts = append(alerts, fmt.Sprintf("💧 High Humidity: %d%% (threshold: %d%%)",
			weather.Current.Humidity, defaultThreshold.Humidity))
	}

	if CheckIfCloudCoverIsHigher(weather.Current.Cloud) {
		alerts = append(alerts, fmt.Sprintf("☁️ High Cloud Cover: %d%% (threshold: %d%%)",
			weather.Current.Cloud, defaultThreshold.Cloud))
	}

	return alerts
}


func HasAlerts(weather *models.WeatherResponse) bool {
	return len(CheckAllConditions(weather)) > 0
}
