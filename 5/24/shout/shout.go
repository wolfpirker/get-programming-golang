package main

import (
	"fmt"
	"strings"
)

type martian struct{}
type rover struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("knrr ", int(l))
}

func (r rover) talk() string {
	return "whir whir"
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	shout(martian{})
	shout(laser(2))
	shout(rover{})

	type crater struct{}
	// crater does not implement talker (missing talk method)
	// shout(crater{})
}
