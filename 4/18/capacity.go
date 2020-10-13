package main

import "fmt"

// program that uses a loop to continuesly append an element to a slice
// prints out the capacity of the slice whenever it changes

// question: Does append always double the capacity,
// when the underlying array runs out of room?

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	lastCap := cap(dwarfs)
	for i := 0; i < 1000; i++ {
		dwarfs = append(dwarfs, "new Planet")
		if lastCap != cap(dwarfs) {
			fmt.Printf("cap: %v\n", cap(dwarfs))
			lastCap = cap(dwarfs)
		}

	}
}
