package main

import (
	json "encoding/json"
	fmt "fmt"
	"net/http"
	os "os"
)

type WetherForecast map[string]interface{}

func main() {
	var weather_string = getWeatherString()
	fmt.Printf(weather_string)
}

func getWeatherString() string {
	weather_forecast, err := getWeatherForecast()
	if err != nil {
		return "💀\n"
	}
	fmt.Printf("%v\n", weather_forecast)
	return "☀️\n"
}

func getWeatherForecast() (weather_forecast WetherForecast, err error) {
	api_key := os.Getenv("OPENWEATHERMAP_API_KEY")
	zip_code := os.Getenv("OPENWEATHERMAP_ZIP")
	country := os.Getenv("OPENWEATHERMAP_COUNTRY")
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?zip=%s,%s&appid=%s", zip_code, country, api_key)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&weather_forecast); err != nil {
		return nil, err
	}
	fmt.Printf("%v", weather_forecast)
	return weather_forecast, nil
}
