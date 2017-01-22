package main

import fmt "fmt"
import os "os"

func main() {
	var weather_string = getWeatherString()
	fmt.Printf(weather_string)
}

func getWeatherString() string {
	var api_key = os.Getenv("OPENWEATHERMAP_API_KEY")
	var zip_code = os.Getenv("OPENWEATHERMAP_ZIP")
	var country = os.Getenv("OPENWEATHERMAP_COUNTRY")
	return "☀️\n"
}
