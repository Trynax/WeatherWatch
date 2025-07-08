package api

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"github.com/trynax/WeatherWatch/models"
)





func GetCurrrentWeather(city string) (*models.WeatherResponse, error) {

	err:= godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}
	if os.Getenv("API_KEY") == "" {
		return nil, fmt.Errorf("API_KEY not set in .env file")
	}

	baseUrl := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", os.Getenv("API_KEY"), city)

	resp, err := http.Get(baseUrl)
	if err!=nil{
		fmt.Println("Error makeing api call: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)

	if err!=nil {
		fmt.Println("Error reading response")
		return nil, err
	}

	fmt.Println(string(body))

	return nil, nil
}