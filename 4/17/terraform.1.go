package main

import "fmt"

// first iteration of terraform with a function
// attempt on my own

func terraform(planets ...string) []string {
	result := make([]string, 0, 3)
	for _, planet := range planets {
		result = append(result, "New "+planet)
	}
	return result
}

func main() {
	planetsSlice := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptun", "Pluto"}
	planetsSelection := planetsSlice[3:4]
	planetsSelection = append(planetsSelection, planetsSlice[6], planetsSlice[7])
	result := terraform(planetsSelection...)
	fmt.Println(planetsSelection)
	fmt.Println("terraformed planets:", result)

}
