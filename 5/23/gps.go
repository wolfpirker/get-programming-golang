package main

import (
	"fmt"
	"math"
)

type sol int

// world with a volumetric mean radius in kilometers
type world struct {
	radius float64
}

type rover struct {
	gps
}

type location struct {
	name      string
	lat, long float64
}

func (l location) description() string {
	return fmt.Sprintf("%s is located at %f, %f", l.name, l.lat, l.long)
}

type gps struct {
	current, destination location
	world                world
}

// rad converts degrees to radians
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// distance calculation using the spherical law of cosines
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func (s gps) distance() float64 {
	return s.world.distance(s.current, s.destination)
}

func (s gps) message() string {
	return fmt.Sprintf("%f km left to go", s.distance())
}

func main() {
	mars := world{radius: 3389.5}
	bradbury := location{
		"Bradburry Landing",
		-4.58,
		137.44,
	}
	elysium := location{"Elysium Planitia", 4.5, 135.9}

	system := gps{
		current:     bradbury,
		destination: elysium,
		world:       mars,
	}

	curiousity := rover{
		gps: system,
	}

	fmt.Println(curiousity.current.description())
	fmt.Println(curiousity.destination.description())
	fmt.Println(curiousity.message())
}
