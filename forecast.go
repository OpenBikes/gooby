package gooby

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type prediction struct {
	Quantity float64 `json:"quantity"`
	Std      float64 `json:"std"`
}

// ObForecast is used for unmarshaling forecast responses from the OpenBikes API
type ObForecast struct {
	Bikes     prediction `json:"bikes"`
	City      string     `json:"city"`
	Stands    prediction `json:"spaces"`
	Station   string     `json:"station"`
	Status    string     `json:"status"`
	Timestamp float64    `json:"timestamp"`
}

// Forecast for a station at a certain time
func Forecast(city, station string, timestamp int64) ObForecast {
	var url = fmt.Sprintf("http://openbikes.co/api/prediction/%s/%s/%d", city, station, timestamp)
	var res, err = http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	var forecast ObForecast
	// Read bytes
	data, err := ioutil.ReadAll(res.Body)
	// Unmarshal
	if err == nil && data != nil {
		err = json.Unmarshal(data, &forecast)
	}
	return forecast
}
