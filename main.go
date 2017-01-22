package main

import (
	json "encoding/json"
	fmt "fmt"
	"net/http"
	os "os"
)

func main() {
	var weather_string = getWeatherString()
	fmt.Printf(weather_string)
}

func getWeatherString() string {
	api_key := os.Getenv("OPENWEATHERMAP_API_KEY")
	zip_code := os.Getenv("OPENWEATHERMAP_ZIP")
	country := os.Getenv("OPENWEATHERMAP_COUNTRY")
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?zip=%s,%s&appid=%s", zip_code, country, api_key)
	resp, err := http.Get(url)
	if err != nil {
		return "💀\n"
	}
	defer resp.Body.Close()

	var target map[string]interface{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&target); err != nil {
		return "💀\n"
	}
	fmt.Printf("%v\n", target)
	return "☀️\n"
}
