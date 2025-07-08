package utils

import (
	"fmt"
	"github.com/trynax/WeatherWatch/models"
)

const (
	TempEmoji     = "🌡️ "
	HumidityEmoji = "💧 "
	WindEmoji     = "🌬️ "
	OkEmoji       = "✅ "
	AlertEmoji    = "⚠️ "
	CloudEmoji    = "☁️ "
	LocationEmoji="📍 "
	TimeEmoji="🕓 "
)
func WeatherCurrentResponse(resp *models.WeatherResponse) string {
	result := fmt.Sprintf("%s%s, %s\n", LocationEmoji, resp.Location.Name, resp.Location.Country)
	result += fmt.Sprintf("%sLocal Time: %s\n", TimeEmoji, resp.Location.Localtime)
	result += fmt.Sprintf("%sTemperature: %.1f°C (Feels like: %.1f°C)\n", TempEmoji, resp.Current.TempC, resp.Current.FeelsLikeC)
	result += fmt.Sprintf("%sCondition: %s\n", CloudEmoji, resp.Current.Condition.Text)
	result += fmt.Sprintf("%sHumidity: %d%%\n", HumidityEmoji, resp.Current.Humidity)
	result += fmt.Sprintf("%sWind: %.1f kph (%s)\n", WindEmoji, resp.Current.WindKph, resp.Current.WindDir)
	result += fmt.Sprintf("%sCloud Cover: %d%%\n",CloudEmoji, resp.Current.Cloud)

	return result
}



