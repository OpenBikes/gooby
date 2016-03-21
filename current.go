package gooby

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type properties struct {
	Address string  `json:"address"`
	Bikes   float64 `json:"bikes"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Name    string  `json:"name"`
	Stands  float64 `json:"stands"`
	Status  string  `json:"status"`
	Update  string  `json:"update"`
}

type point struct {
	Geometry   geometry   `json:"geometry"`
	Properties properties `json:"properties"`
	Type       string     `json:"type"`
}

// ObGeojson is used for unmarshaling geojson responses from the OpenBikes API
type ObGeojson struct {
	Features []point `json:"features"`
	Type     string  `json:"type"`
	Status   string  `json:"success"`
}

// Current information for a city
func Current(city string) ObGeojson {
	var url = fmt.Sprintf("http://openbikes.co/api/geojson/%s", city)
	var res, err = http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	var geojson ObGeojson
	// Read bytes
	data, err := ioutil.ReadAll(res.Body)
	// Unmarshal
	if err == nil && data != nil {
		err = json.Unmarshal(data, &geojson)
	}
	return geojson
}
