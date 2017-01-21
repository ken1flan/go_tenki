package main

import fmt "fmt"

func main() {
	var weather_string = getWeatherString()
	fmt.Printf(weather_string)
}

func getWeatherString() string {
	return "☀️\n"
}
