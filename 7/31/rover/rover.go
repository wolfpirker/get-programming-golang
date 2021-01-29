package main

import (
	"image"
	"log"
	"time"
)

func main() {
	r := NewRoverDriver()
	time.Sleep(2 * time.Second)
	r.Left()
	time.Sleep(time.Second)
	r.Stop()
	time.Sleep(2 * time.Second)
	r.Start()
	time.Sleep(time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
}

// RoverDriver drives a rover around the surface of Mars.
type RoverDriver struct {
	commandc chan command
}

// NewRoverDriver starts a new RoverDriver and returns it.
func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

type command int

const (
	right = command(0)
	left  = command(1)
	start = command(2)
	stop  = command(3)
)

// drive is responsible for driving the rover. It
// is expected to be started in a goroutine.
func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	prev_cmnd := right
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
				prev_cmnd = right
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
				prev_cmnd = left
			case start:
				prev_cmnd = start
			case stop:
				prev_cmnd = stop
			}
			if (prev_cmnd == left) || (prev_cmnd == right) {
				log.Printf("new direction %v", direction)
			} else if prev_cmnd == stop {
				log.Printf("rover stopped")
			} else if prev_cmnd == start {
				log.Printf("rover started")
			}

		case <-nextMove:
			pos = pos.Add(direction)
			if (prev_cmnd == left) || (prev_cmnd == right) {
				log.Printf("moved to %v", pos)
			}

			nextMove = time.After(updateInterval)
		}
	}
}

// Left turns the rover left (90° counterclockwise).
func (r *RoverDriver) Left() {
	r.commandc <- left
}

// Right turns the rover right (90° clockwise).
func (r *RoverDriver) Right() {
	r.commandc <- right
}

func (r *RoverDriver) Start() {
	r.commandc <- start
}

func (r *RoverDriver) Stop() {
	r.commandc <- stop
}
