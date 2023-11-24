package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type apiConfig struct {
	Nome    string `json:"name"`
	ApiKey  string `json:"key"`
	Version string `json:"version"`
	Measure string `json:"measure"`
	Code    string `json:"code"`
}

func (ap *apiConfig) GetApiInfo(filename string) (*apiConfig, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		err := fmt.Errorf("Error %v", err)
		return nil, err
	}
	err = json.Unmarshal(bytes, &ap)

	if err != nil {
		return nil, err
	}
	return ap, nil
}

type Forecast struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func ReturnCityWeather(cityName []string) {

	var apconf apiConfig

	_, err := apconf.GetApiInfo("config.json")
	if err != nil {
		fmt.Printf("Error found %v\n", err)
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=%v", cityName[0], apconf.ApiKey, apconf.Measure)

	// Send a GET request to OpenWeatherMap API
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching weather data:", err)
		return
	}

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Parse the JSON response
	var weatherData Forecast
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Print the weather information
	var city string
	for _, v := range cityName {
		city += fmt.Sprintf(" %s", v)
	}

	// Creating a return variable called result
	forecast := fmt.Sprintf(`
Weather in %v/%v
Description: %v
Temperature: %v %v
Feels Like: %v %v
`, city, weatherData.Sys.Country, strings.ToUpper(weatherData.Weather[0].Description), weatherData.Main.Temp, apconf.Code, weatherData.Main.FeelsLike, apconf.Code)
	fmt.Println(forecast)
}
