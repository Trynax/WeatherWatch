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
func FormatSummary(weather *models.WeatherResponse) string {
	return fmt.Sprintf(
		`📍 %s, %s
🕓 Local Time: %s
🌡️ Temperature: %.1f°C (Feels like: %.1f°C)
☁️ Condition: %s
💧 Humidity: %d%%
🌬️ Wind: %.1f kph (%s)
☁️ Cloud Cover: %d%%`,
		weather.Location.Name,
		weather.Location.Country,
		weather.Location.Localtime,
		weather.Current.TempC,
		weather.Current.FeelsLikeC, // Optional feels-like estimate
		weather.Current.Condition.Text,
		weather.Current.Humidity,
		weather.Current.WindKph,
		weather.Current.WindDir,
		weather.Current.Cloud,
	)
}



func FormatAlertMessage(location, timestamp string, alerts []string) string {
	msg := fmt.Sprintf("⚠️ ALERTS Triggered\n📍 %s | 🕒 %s\n---------------------------------\n", location, timestamp)
	for _, alert := range alerts {
		msg += alert + "\n"
	}
	return msg
}

func FormatNoAlertMessage(weather models.WeatherResponse) string {
	return fmt.Sprintf(
		"[🕓 %s] ✅ Weather OK | %.1f°C, Humidity %d%%, Wind %.1f kph",
		weather.Current.LastUpdated,
		weather.Current.TempC,
		weather.Current.Humidity,
		weather.Current.WindKph,
	)
}


func FormatLogEntry(weather models.WeatherResponse, alerts []string) string {
	if len(alerts) == 0 {
		return fmt.Sprintf("[🕓 %s] OK - Temp: %.1f°C | Condition: %s",
			weather.Current.LastUpdated,
			weather.Current.TempC,
			weather.Current.Condition.Text,
		)
	}

	msg := fmt.Sprintf("[🕓 %s] ALERT - ", weather.Current.LastUpdated)
	for i, alert := range alerts {
		if i > 0 {
			msg += " | "
		}
		msg += alert
	}
	return msg
}


func FormatEmailBody(location, timestamp string, alerts []string) string {
	body := fmt.Sprintf("Hello,\n\nAn alert was triggered by your Weather Alert Service:\n\n📍 Location: %s\n🕓 Time: %s\n\n🚨 Condition(s):\n", location, timestamp)
	for _, alert := range alerts {
		body += "- " + alert + "\n"
	}
	body += "\nStay safe,\nWeatherWatch Bot ☁️"
	return body
}
