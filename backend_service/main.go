package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func main() {
	mw := multiWeatherProvider{
		openWeatherMap{apikey: "API KEY"},
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		city := strings.SplitN(r.URL.Path, "/", 3)[2]

		temp, err := mw.temprature(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf8")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"city": city,
			"temp": temp,
			"took": time.Since(begin).String(),
		})
	})

	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world..."))
}

// type weatherData struct {
// 	Name string `json:"name"`
// 	Main struct {
// 		Kelvin float64 `json:"temp"`
// 	} `json:"main"`
// }

type weatherProvider interface {
	temprature(city string) (float64, error)
}
type openWeatherMap struct {
	apikey string
}

func (w openWeatherMap) temprature(city string) (float64, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + w.apikey + "&q=" + city)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	var d struct {
		Main struct {
			Kelvin float64 `json:"temp"`
		} `json:"main"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return 0, err
	}

	return d.Main.Kelvin, nil
}

type multiWeatherProvider []weatherProvider

func (w multiWeatherProvider) temprature(city string) (float64, error) {
	sum := 0.0

	for _, provider := range w {
		k, err := provider.temprature(city)
		if err != nil {
			return 0, err
		}

		sum += k
	}
	return sum / float64(len(w)), nil
}
