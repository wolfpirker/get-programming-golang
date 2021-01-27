package main

import (
	"fmt"
)

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go removeDouble(c0, c1)
	printGopher(c1)
}
func sourceGopher(downstream chan string) {

	for _, v := range []string{"hello world", "hello world", "goodbye all"} {
		downstream <- v
	}
	close(downstream)
}

func removeDouble(upstream, downstream chan string) {
	previous_value := ""
	for item := range upstream {
		if previous_value == item {
			continue
		}
		previous_value = item
		downstream <- item
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}
