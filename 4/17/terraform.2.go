package main

import "fmt"

// finished terraform with a terraform method, of an planets object
// attempt on my own
// main differnce to proposed solution: terraform method creates a copied result
// and returns it

type Planets []string

func (p Planets) terraform() Planets {
	result := make([]string, 0, 3)
	for _, planet := range p {
		result = append(result, "New "+planet)
	}
	return result
}

func main() {
	planetsSlice := Planets{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptun", "Pluto"}
	planetsSelection := planetsSlice[3:4]
	planetsSelection = append(planetsSelection, planetsSlice[6], planetsSlice[7])
	result := planetsSelection.terraform()
	fmt.Println(planetsSelection)
	fmt.Println("terraformed planets:", result)

}
