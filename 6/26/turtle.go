package main

import "fmt"

type turtle struct{
	x int
	y int
}

func (t *turtle) up() {
	t.y--
}

func (t *turtle) down() {
	t.y++
}

func (t *turtle) left() {
	t.x--
}

func (t *turtle) right() {
	t.x++
}

func main() {
	t := turtle{0,0}

	t.up()
	t.up()	
	fmt.Printf("%v\n", t)

	t.down()
	t.right()
	fmt.Printf("%v\n", t)
}
