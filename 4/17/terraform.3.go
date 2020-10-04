package main

import "fmt"

// finished terraform with a terraform method, of an planets object
// used solution of the book

type Planets []string

func (p Planets) terraform() {
	for i := range p {
		p[i] = "New " + p[i]
	}
}

func main() {
	planetsSlice := Planets{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptun", "Pluto"}
	planetsSelection := planetsSlice[3:4]
	planetsSelection = append(planetsSelection, planetsSlice[6], planetsSlice[7])
	planetsSelection.terraform()
	fmt.Println(planetsSlice)
}
