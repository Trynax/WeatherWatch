package api

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/trynax/WeatherWatch/models"
)





func GetCurrentWeatherWithCityName(city string) (*models.WeatherResponse, error) {

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

	body,err := io.ReadAll(resp.Body)

	if err!=nil {
		fmt.Println("Error reading response")
		return nil, err
	}


	weatherResponse :=&models.WeatherResponse{}

	err= json.Unmarshal(body, weatherResponse)
	if err != nil{
		return  nil, err
	}



	return weatherResponse, nil
}


func GetCurrentWeatherWithLongitudeAndLatitude(lat,long float64) (*models.WeatherResponse, error){
		err:= godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}
	if os.Getenv("API_KEY") == "" {
		return nil, fmt.Errorf("API_KEY not set in .env file")
	}

	baseUrl := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%v,%v", os.Getenv("API_KEY"), lat, long)


	resp, err := http.Get(baseUrl) 

	if err !=nil{
		fmt.Println("Error making api call")
		return nil, err
	}

	defer resp.Body.Close()

	body,err := io.ReadAll(resp.Body)

	if err!=nil{
		fmt.Println("Error read the reponse body: ", err)

		return  nil, err
	}

	weatherResponse :=  &models.WeatherResponse{}
	err= json.Unmarshal(body, weatherResponse)
	if err != nil{
		return  nil, err
	}

	return weatherResponse, nil
}


 
