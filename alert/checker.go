package alert

import (
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


