package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type WeatherForecast map[string]interface{}

func GetWeatherString() string {
	weather_forecast, err := getWeatherForecast()
	if err != nil {
		return "💀\n"
	}

	weather_mark := map[int]string{
		// Thunderstorm
		200: "⚡️☂️",   // thunderstorm with light rain
		201: "⚡️☂️",   // thunderstorm with rain
		202: "⚡️☂️☂️", // thunderstorm with heavy rain
		210: "⚡️☂️",   // light thunderstorm
		211: "⚡️☂️",   // thunderstorm
		212: "⚡️⚡️☂️", // heavy thunderstorm
		221: "⚡️☂️",   // ragged thunderstorm
		230: "⚡️☂️🌫",  // thunderstorm with light drizzle
		231: "⚡️☂️🌫",  // thunderstorm with drizzle
		232: "⚡️☂️🌫🌫", // thunderstorm with heavy drizzle
		// Drizzle
		300: "🌫",     // light intensity drizzle
		301: "🌫",     // drizzle
		302: "🌫🌫",    // heavy intensity drizzle
		310: "🌫☂️",   // light intensity drizzle rain
		311: "🌫☂️",   // drizzle rain
		312: "🌫🌫☂️",  // heavy intensity drizzle rain
		313: "🌫☂️",   // shower rain and drizzle
		314: "🌫☂️☂️", // heavy shower rain and drizzle
		321: "🌫☂️",   // shower drizzle
		// Rain
		500: "☂️",       // light rain
		501: "☂️",       // moderate rain
		502: "☂️☂️",     // heavy intensity rain
		503: "☂️☂️☂️",   // very heavy rain
		504: "☂️☂️☂️☂️", // extreme rain
		511: "☂️⛄️",     // freezing rain
		520: "☂️",       // light intensity shower rain
		521: "☂️",       // shower rain
		522: "☂️☂️",     // heavy intensity shower rain
		531: "☂️☂️☂️",   // ragged shower rain
		// Snow
		600: "⛄️",     // light snow
		601: "⛄️",     // snow
		602: "⛄️⛄️",   // heavy snow
		611: "⛄️☂️",   // sleet
		612: "⛄️☂️",   // shower sleet
		615: "⛄️☂️",   // light rain and snow
		616: "⛄️☂️",   // rain and snow
		620: "⛄️☂️",   // light shower snow
		621: "⛄️☂️",   // shower snow
		622: "⛄️☂️☂️", // heavy shower snow
		// Atmosphere
		701: "🌫",                 // mist
		711: "🌫",                 // smoke
		721: "🌫",                 // haze
		731: "sand, dust whirls", // sand, dust whirls
		741: "🌫",                 // fog
		751: "sand",              // sand
		761: "dust",              // dust
		762: "volcanic ash",      // volcanic ash
		771: "💨☂️",               // squalls
		781: "🌪",                 // tornado
		// Clouds
		800: "☀️",   // clear sky
		801: "☀️",   // few clouds
		802: "☁️",   // scattered clouds
		803: "☁️",   // broken clouds
		804: "☁️☁️", // overcast clouds
		// Extreme
		900: "🌪",     // tornado
		901: "💨☂️",   // tropical storm
		902: "🌪☂️",   // hurricane
		903: "cold",  // cold
		904: "hot",   // hot
		905: "windy", // windy
		906: "hail",  // hail
		// Additional
		951: "calm", // calm
		952: "🌬",    // light breeze
		953: "🌬",    // gentle breeze
		954: "🌬",    // moderate breeze
		955: "🌬",    // fresh breeze
		956: "💨",    // strong breeze
		957: "💨",    // high wind, near gale
		958: "💨",    // gale
		959: "💨💨",   // severe gale
		960: "🌪☂️",  // storm
		961: "🌪☂️",  // violent storm
		962: "🌪☂️",  // hurricane
	}

	weather_id := int(weather_forecast["weather"].([]interface{})[0].(map[string]interface{})["id"].(float64))

	return weather_mark[weather_id]
}

func getWeatherForecast() (weather_forecast WeatherForecast, err error) {
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
