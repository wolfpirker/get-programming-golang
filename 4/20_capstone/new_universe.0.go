package main

import (
	"fmt"
	"math/rand"
)

const (
	width  = 80
	height = 15
)

func newUniverse() Universe {
	a := make([][]bool, height)
	for i := range a {
		a[i] = make([]bool, width)
	}
	return a
}

func (u Universe) Show() {
	/*for i := range u {
		fmt.Println(u[i])
	}*/
	fmt.Print("\x0c", u.String())
	fmt.Println()
}

func (u Universe) String() string {
	var b byte
	// why use byte slice, instead of string slice?
	result := make([]byte, 0, (width+1)*height)
	for i := range u {
		for _, isAlive := range u[i] {
			b = '.'
			if isAlive {
				b = '*'
			}
			result = append(result, b)
		}
		result = append(result, '\n')
	}

	return string(result)
}

func (u Universe) Seed() {
	// set randomly around 25% of the cells to be alive
	// see different and faster solution in the book
	// why does it produce the same result always?
	r := 0
	for i, _ := range u {
		for j, _ := range u[i] {
			r = rand.Intn(100)
			if r < 25 {
				u[i][j] = true
			} else {
				u[i][j] = false
			}
		}
	}
}

type Universe [][]bool // True: cell is alive, False: dead
/* use slices instead of arrays so that a universe can be shared with,
 * and modified by functions or methods
 * -> Note: Arrays can be shared using pointers - see lesson 26
 */
func main() {
	// 20.1 A new universe
	u := newUniverse()
	u.Seed()
	u.Show()
}
