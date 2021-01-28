package main

import (
	"fmt"
	"time"
)

func main() {
	go worker()
	time.Sleep(5 * time.Second)
}
func worker() {
	n := 0
	next := time.After(time.Second) // make timer channel
	for {
		select {
		case <-next: // wait for the timer to fire
			n++
			fmt.Println(n)
			next = time.After(time.Second)
			// -> timer channel for another event
		}
	}
}
