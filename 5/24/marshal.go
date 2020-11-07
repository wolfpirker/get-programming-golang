package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// location with a latitude, longitude in decimal degrees.
type location struct {
	Name string `json:"name"`
	Lat coordinate `json:"latitude"`
	Long coordinate `json:"longitude"`
}

// String formats a location with latitude, longitude.
func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.Lat, l.Long)
}

// coordinates in degrees, minutes, seconds in N/S/E/W hemisphere
type coordinate struct {
	d, m, s float64
	h       rune
}

// String formats a location with latitude, longitude.
func (c coordinate) String() string {
	return fmt.Sprintf("%v°%v'\"%v %c", c.d, c.m, c.s, c.h)
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

// see example of implementing MarshalJSON, here:
// https://golang.org/pkg/encoding/json/
func (c coordinate) MarshalJSON() ([]byte, error) {
	var s []byte
	var e error
	s, e = json.Marshal(struct {
		DD float64 `json:"decimal"`
		DMS string `json:"dms"`
		D float64 `json:"degrees"`
		M float64 `json:"minutes"`
		S float64 `json:"seconds"`
		H string `json:"hemisphere"`
	}{
		DD: c.decimal(),
		DMS: c.String(),
		D: c.d,
		M: c.m,
		S: c.s,
		H: string(c.h),
	})
	return s, e
}

func main() {

	elysium := location{
		Name: "Elysium Planitia",
		Lat: coordinate{4.0, 30.0, 0.0, 'N'},
		Long: coordinate{135.0, 54.0, 0.0, 'E'},
	}

	bytes, err := json.MarshalIndent(elysium, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(bytes))
}
