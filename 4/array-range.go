package main

import "fmt"

func main() {
	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	for _, dwarf := range dwarfs {
		fmt.Println(dwarf)
	}
	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}
}
