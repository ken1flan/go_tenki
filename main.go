package main

import (
	fmt "fmt"

	"./weather"
)

func main() {
	var weather_string = weather.GetWeatherString()
	fmt.Printf(weather_string)
}
