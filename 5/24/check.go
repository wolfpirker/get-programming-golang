package main

import "fmt"

// location with a latitude, longitude in decimal degrees.
type location struct {
	lat, long coordinate
}

// String formats a location with latitude, longitude.
func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

// location with a latitude, longitude in decimal degrees.
type coordinate struct {
	d, m, s float64
	h       rune
}

// String formats a location with latitude, longitude.
func (c coordinate) String() string {
	return fmt.Sprintf("%v°%v'\"%v %c", c.d, c.m, c.s, c.h)
}

func main() {
	lat := coordinate{4.0, 30.0, 0.0, 'N'}
	long := coordinate{135.0, 54.0, 0.0, 'E'}
	elysium := location{lat, long}
	/* or like this...
	elysium := location{
		lat: coordinate{4.0, 30.0, 0.0, 'N'}
		long: coordinate{135.0, 54.0, 0.0, 'E'}
	}
	*/

	fmt.Println("Elysium Planitia is at", elysium)
}
