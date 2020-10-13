package main

import (
	"fmt"
	"sort"
)

// no built in set type available in Go
// improvisation by using a map

func main() {
	var temperatures = []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float64]bool)
	for _, t := range temperatures {
		set[t] = true
	}

	if set[-28.0] {
		fmt.Println("set member")
	}

	// before sorting the keys, convert temperatures back to a slice
	fmt.Println(set)
	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
	}
	sort.Float64s(unique)

	fmt.Println(unique)
}
