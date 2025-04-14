// Package ca parses data from Chee Aun: https://github.com/cheeaun/sgbusdata/
package ca

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type StopDetail struct {
	Lng  float64
	Lat  float64
	Desc string
	Road string
}
type Stop map[string]StopDetail

// Usage: Stop["10009"].Desc
func ParseStops() Stop {
	f := openStopsFile()
	defer f.Close()

	dat := readStopsDat(f)
	return toStop(dat)
}
func openStopsFile() *os.File {
	const fn = "../../testdata/cheeaun/stops.json"
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
func readStopsDat(r io.Reader) map[string]any {
	var dat map[string]any
	dec := json.NewDecoder(r)
	if err := dec.Decode(&dat); err != nil {
		log.Fatal(err)
	}
	return dat
}
func toStop(dat map[string]any) Stop {
	var s = make(Stop)
	for k, v := range dat {
		val := v.([]any)
		if len(val) != 4 {
			log.Fatalln("malformed stop data", k, v)
		}
		// fmt.Printf("%s: %v\n", k, val)
		lng, okLng := val[0].(float64)
		lat, okLat := val[1].(float64)
		desc, okDesc := val[2].(string)
		road, okRoad := val[3].(string)
		if !okLng || !okLat || !okDesc || !okRoad {
			log.Fatalln("malformed stop data", k, v)
		}
		sd := StopDetail{
			Lng:  lng,
			Lat:  lat,
			Desc: desc,
			Road: road,
		}
		s[k] = sd
	}
	return s
}
