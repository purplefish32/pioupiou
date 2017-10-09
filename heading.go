package main

import (
        "log"
        "net/http"
        "encoding/json"
        "math"
        "fmt"
)

type Measurement struct {
	Data struct {
		Measurements struct {
			WindHeading float64 `json:"wind_heading"`
		} `json:"measurements"`
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

                var heading = measure.Data.Measurements.WindHeading

                fmt.Print(degToCompass(heading))
        }
}

func degToCompass(num float64) string {
  var val = math.Floor((num / 22.5) + 0.5)
  var heading [16]string
  heading[0] = "N"
  heading[1] = "NNE"
  heading[2] = "NE"
  heading[3] = "ENE"
  heading[4] = "E"
  heading[5] = "ESE"
  heading[6] = "SE"
  heading[7] = "SSE"
  heading[8] = "S"
  heading[9] = "SSW"
  heading[10] = "SW"
  heading[11] = "WSW"
  heading[12] = "W"
  heading[13] = "WNW"
  heading[14] = "NW"
  heading[15] = "NNW"
  return heading[int(val)]
}
