package main

import (
	"fmt"
	"image"
	"time"
)

func main() {
	go worker()
	time.Sleep(10 * time.Second)
}
func worker() {
	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 2}
	delay := time.Second
	next := time.After(delay)
	for {
		select {
		case <-next:
			pos = pos.Add(direction)
			delay += 500 * time.Millisecond
			fmt.Println("current position is ", pos)
			next = time.After(delay)
		}
	}
}
