package main

import (
	"fmt"
	"math"
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

// rad converts degrees to radians
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// distance calculation using the spherical law of cosines
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.Lat.decimal()))
	s2, c2 := math.Sincos(rad(p2.Lat.decimal()))
	clong := math.Cos(rad(p1.Long.decimal() - p2.Long.decimal()))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
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

// world with a volumetric mean radius in kilometer
type world struct {
	radius float64
}

var (
	mars = world{radius: 3389.5}
)

func main() {

	spirit := newLocation("Columbia Memorial Station", coordinate{14., 34., 6.2, 'S'}, coordinate{175., 28., 21.5, 'E'})
	opportunity := newLocation("Challenger Memorial Station", coordinate{1., 56., 46.3, 'S'}, coordinate{354., 28., 24.2, 'E'})
	curiosity := newLocation("Bradbury Landing", coordinate{4., 35., 22.2, 'S'}, coordinate{137., 26., 30.1, 'E'})
	// insight := ...

	fmt.Printf("Spirit-Opportunity: %.1f km\n", mars.distance(spirit, opportunity))
	fmt.Printf("Spirit-Curiousity: %.1f km\n", mars.distance(spirit, curiosity))
	fmt.Printf("Opportunity-Curiousity: %.1f km\n", mars.distance(opportunity, curiosity))
}
