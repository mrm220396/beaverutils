/*
* Package Beta Stage
 */

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"beaverutils/src/weather"
)

var help string = `
 Weather CLI program from Beaver Utils
 
 weather [CITYNAME]			   Current Forecast
 weather --setmetric [OPTION] [k or c or f] Set the temperature measurement unit whether Kelvin, Celcius or Farenheit
 `

func main() {
	var response = os.Args[1:]

	for _, v := range response {
		if strings.Contains(v, "--") {
			CommandChecker(v, response)
		}
	}

	weather.ReturnCityWeather(response)
}

func CommandChecker(value string, response []string) {
	switch value {
	case "--help":
		fmt.Println(help)
	case "--setmetric":
		setMetric(response[1])
	}
}

// setMetric sets the temperature unit is going to be requested
// This keeps the program more effective.
func setMetric(value string) {
	var ap weather.ApiConfig
	file, _ := os.ReadFile("config.json")

	json.Unmarshal(file, &ap)

	if value == "k" {
		ap.Measure = ""
		ap.Code = "K"
	} else if value == "f" {
		ap.Measure = "imperial"
		ap.Code = "F"
	} else {
		ap.Measure = "metric"
		ap.Code = "C"
	}

	writer, err := os.OpenFile("config.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Errorf("Error found %v", err)
	}

	defer writer.Close()

	response, _ := json.MarshalIndent(ap, "", "    ")

	os.WriteFile("config.json", response, 0644)
	writer.Close()
	os.Exit(0)
}
