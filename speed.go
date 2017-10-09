package main

import (
        //"io"
        //"os"
        "log"
        "net/http"
        "encoding/json"
        "time"
      //"strconv"
        "math"
        "fmt"
)

type Measurement struct {
	Doc string `json:"doc"`
	License string `json:"license"`
	Attribution string `json:"attribution"`
	Data struct {
		ID int `json:"id"`
		Meta struct {
			Name string `json:"name"`
		} `json:"meta"`
		Location struct {
			Latitude float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
			Date time.Time `json:"date"`
			Success bool `json:"success"`
		} `json:"location"`
		Measurements struct {
			Date time.Time `json:"date"`
			Pressure interface{} `json:"pressure"`
			WindHeading int `json:"wind_heading"`
			WindSpeedAvg float64 `json:"wind_speed_avg"`
			WindSpeedMax float64 `json:"wind_speed_max"`
			WindSpeedMin float64 `json:"wind_speed_min"`
		} `json:"measurements"`
		Status struct {
			Date time.Time `json:"date"`
			Snr float64 `json:"snr"`
			State string `json:"state"`
		} `json:"status"`
	} `json:"data"`
}

func main() {
        response, err := http.Get("http://api.pioupiou.fr/v1/live/379")
        if err != nil {
                log.Fatal(err)
        } else {
                defer response.Body.Close()

                var measure Measurement

                json.NewDecoder(response.Body).Decode(&measure)

                var knots = measure.Data.Measurements.WindSpeedAvg / 1.852

                fmt.Print(round(knots))
        }
}

func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}
