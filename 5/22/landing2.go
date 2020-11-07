package main

import (
	"fmt"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	Name string
	Lat  coordinate
	Long coordinate
}

func newLocation(name string, lat, long coordinate) location {
	return location{name, lat, long}
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func (loc location) printDecimalDegrees() {
	fmt.Printf("%-32v %5.4f, %5.4f\n", loc.Name, loc.Lat.decimal(), loc.Long.decimal())
}

func main() {

	spirit := newLocation("Columbia Memorial Station", coordinate{14., 34., 6.2, 'S'}, coordinate{175., 28., 21.5, 'E'})
	opportunity := newLocation("Challenger Memorial Station", coordinate{1., 56., 46.3, 'S'}, coordinate{354., 28., 24.2, 'E'})
	curiosity := newLocation("Bradbury Landing", coordinate{4., 35., 22.2, 'S'}, coordinate{137., 26., 30.1, 'E'})
	// insight := ...

	spirit.printDecimalDegrees()
	opportunity.printDecimalDegrees()
	curiosity.printDecimalDegrees()
}
