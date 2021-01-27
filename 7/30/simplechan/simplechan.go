package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// channel to send integer values:
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	for i := 0; i < 5; i++ {
		gopherID := <-c // left arrow operator
		// receive, point away of the channel
		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}
}

func sleepyGopher(id int, c chan int) {
	rt := time.Duration(rand.Intn(7))
	fmt.Println("... ", id, " snore ...")
	time.Sleep(rt * time.Second)
	c <- id // // to send: point to the channel
}
