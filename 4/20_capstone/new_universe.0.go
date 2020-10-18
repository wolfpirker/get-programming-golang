package main

import (
	"fmt"
	"math/rand"
	"time"
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
	// Note: random, by using a variable Seed - see documentation
	// https://golang.org/pkg/math/rand/#Intn
	r := 0
	for i, _ := range u {
		for j, _ := range u[i] {
			rand.Seed(time.Now().UnixNano())
			r = rand.Intn(100)
			if r < 25 {
				u[i][j] = true
			} else {
				u[i][j] = false
			}
		}
	}
}

// 20.2
func (u Universe) Alive(x, y int) bool {
	// Note: wrap around if the cell is outside,
	// of universe
	x = (x + width) % width
	y = (y + height) % height
	if u[x][y] {
		return true
	}
	return false
}

func (u Universe) Neighbours(x, y int) int {
	// count neighbours, using Alive method
	count := 0
	for x1 := x - 1; x1 <= x+1; x1++ {
		for y1 := y - 1; y1 <= y+1; y1++ {
			if (x1 == x) && (y1 == y) {
				continue
			} else {
				if u.Alive(x1, y1) {
					count++
				}
			}
		}
	}
	return count
}

type Universe [][]bool // True: cell is alive, False: dead
/* use slices instead of arrays so that a universe can be shared with,
 * and modified by functions or methods
 * -> Note: Arrays can be shared using pointers - see lesson 26
 */
func main() {
	// 20.1 A new universe
	// 20.2 Implementing game rules
	u := newUniverse()
	u.Seed()
	u.Show()
	//count := u.Neighbours(-2, -2)
	// panic: runtime error: index out of range!
	count := u.Neighbours(1, 1)
	println(count)
}
