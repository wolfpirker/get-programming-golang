package main

import (
	"fmt"
	"math/rand"
)

type description string
type food string

type animal interface {
	move() description
	eat() food	
}

type cow struct {
	name string
}

type duck struct {
	name string
}

func (a cow) String() string {
	return a.name
}

func (a duck) String() string {
	return a.name
}

func (a cow) move() description {
	switch rand.Intn(3) {
	case 0:
		return "escapes from danger"
	case 1:
		return "runs towards food source"
	default:
		return "goes slowly around"
	}
}

func (a cow) eat() food {
	switch rand.Intn(3) {
	case 0:
		return "oat"
	case 1:
		return "hay"
	default:
		return "fresh gras"
	}
}

func (a duck) move() description {
	switch rand.Intn(3) {
	case 0:
		return "takes a bath"
	default:
		return "goes for a hike"
	}
}

func (a duck) eat() food {
	switch rand.Intn(3) {
	case 0:
		return "snails"
	case 1:
		return "bugs"
	default:
		return "fresh gras"
	}
}

func randomAction(a animal) {
	switch rand.Intn(2){
	case 0:
		fmt.Printf("%v %v.\n", a, a.move())
	case 1:
		fmt.Printf("%v eats %v.\n", a, a.eat())
	}
}

func main() {
	simulationDuration := 72
	const sunset  = 19
	const sunrise = 6

	bless := cow{name: "Bless"}
	berta := cow{name: "Berta"}
	quack := duck{name: "Quack"}

	// run simulation
	for h:=0; h < simulationDuration; h++ {
		if (h % 24) >= sunrise && (h % 24) < sunset {
			if ((h % 24) == 6) {
				fmt.Printf("sunrise, day %v\n", h/24 + 1)
			}
			fmt.Printf("current hour: %v\n", h)
			switch rand.Intn(3){
			case 0:
				randomAction(bless)
			case 1:
				randomAction(berta)
			case 2:
				randomAction(quack)
			}
		} else{
			if ((h % 24) == 19){
				fmt.Printf("sunset at hour %v -> animals sleep\n", h)
			}
			
		}
		
	}

}
