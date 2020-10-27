package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	type location struct {
		Name string  `json:"location"`
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	locations := []location{
		{
			Name: "Bradbury Landing",
			Lat:  -4.5895,
			Long: 137.4417},
		{
			Name: "Columbia Memorial Station",
			Lat:  -14.5684,
			Long: 175.472636},
		{
			Name: "Challenger Memorial Station",
			Lat:  -1.9462,
			Long: 354.4734},
	}

	bytes, err := json.MarshalIndent(locations, "", " ")
	exitOnError(err)

	fmt.Println(string(bytes))
}

// exitOnError prints any errors and exits
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
