package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type WeatherResult struct {
	City        string
	Unit        string
	Temperature float64
}

func HandlerCurrent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch r.Method {
	case http.MethodGet:
		api_url, setBool := os.LookupEnv("API_URL_CURRENT")
		if !setBool {
			fmt.Println("API_URL_FORECAST env not found")
		}
		var currentWeather WeatherResult
		var cityName string
		r.ParseForm()
		for key, value := range r.Form {
			if key == "city" {
				cityName = value[0]
			}
		}
		client := http.Client{}
		api_url = fmt.Sprintf(api_url, cityName)
		fmt.Println(api_url)
		req, err := http.NewRequest("GET", api_url, nil)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		var result map[string]interface{}

		json.NewDecoder(resp.Body).Decode(&result)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result, "\n\n\n\n")
		cod, ok := result["cod"].(float64)
		if ok && cod == 200 {
			currentWeather.Temperature = result["main"].(map[string]interface{})["temp"].(float64)
			currentWeather.City = result["name"].(string)
			currentWeather.Unit = "Celsius"
		}

		fmt.Print(currentWeather, "\n\n\n")

		w.Header().Set("Content-type", "text/html")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(currentWeather)
		//w.Write(currentWeather)

	default:
		println("default")
	}
}

func HandlerForecast(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch r.Method {
	case http.MethodGet:
		api_url, setBool := os.LookupEnv("API_URL_FORECAST")
		if !setBool {
			fmt.Println("API_URL_FORECAST env not found")
		}
		var forecastWeather WeatherResult
		var cityName string
		var timeStamp int
		var nearDataWeather map[string]interface{}
		r.ParseForm()
		for key, value := range r.Form {
			if key == "city" {
				cityName = value[0]
			}
			if key == "dt" {
				dt, e := strconv.Atoi(value[0])
				timeStamp = dt
				fmt.Println(timeStamp)
				if e != nil {
					log.Fatal(e)
				}
			}
		}
		client := http.Client{}

		api_url = fmt.Sprintf(api_url, cityName)
		req, err := http.NewRequest("GET", api_url, nil)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		var result map[string]interface{}

		json.NewDecoder(resp.Body).Decode(&result)

		if err != nil {
			log.Fatal(err)
		}

		cod, ok := result["cod"].(string)
		if !ok || cod != "200" {
			fmt.Println("hello there")
			w.Header().Set("Content-type", "text/html")
			w.WriteHeader(http.StatusNoContent)
		}
		for _, v := range result["list"].([]interface{}) {
			if int(v.(map[string]interface{})["dt"].(float64)) >= timeStamp {
				nearDataWeather = v.(map[string]interface{})
				break
			}
		}
		fmt.Println(nearDataWeather)
		forecastWeather.Temperature = nearDataWeather["main"].(map[string]interface{})["temp"].(float64)
		forecastWeather.City = result["city"].(map[string]interface{})["name"].(string)
		forecastWeather.Unit = "Celsius"

		fmt.Print(forecastWeather, "\n\n\n")

		w.Header().Set("Content-type", "text/html")
		json.NewEncoder(w).Encode(forecastWeather)

	default:
		println("default")
	}
}

func Server() {
	portStr, setBool := os.LookupEnv("LISTEN_PORT")
	if !setBool {
		fmt.Println("unset")
	}
	fmt.Println("Server start... LISTEN_PORT=", portStr)

	http.HandleFunc("/v1/forecast/", HandlerForecast)
	http.HandleFunc("/v1/current/", HandlerCurrent)

	log.Fatal(http.ListenAndServe(portStr, nil))
}
