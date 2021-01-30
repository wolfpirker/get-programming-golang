package main

import (
	"image"
	"log"
	"math/rand"
	"sync"
	"time"
)

var num_rover int

type cell struct {
	occupied bool
}

type cord struct {
	X int
	Y int
}

// surface on Mars, may be used
// concurrently by different goroutines
// Note: only allows positive coordinates
type MarsGrid struct {
	mu     sync.Mutex
	bounds cord
	grid   [][]cell
}

// Occupy occupies a cell at the given point
func (g *MarsGrid) Occupy(p image.Point) *Occupier {
	g.mu.Lock()
	defer g.mu.Unlock()
	occ := &Occupier{
		grid: g,
		pos:  p,
	}
	if occ == nil {
		return nil
	}
	occ = &Occupier{
		grid: g,
		pos:  p,
	}
	return occ
}

// Occupier represents a occupied cell in the grid
type Occupier struct {
	grid *MarsGrid
	pos  image.Point
}

// Move moves the occupier to a different cell in the grid
// reports back whether the move was successful
// could fail due to moving out of the grid or
// cell already occupied - for fail: occupier stay at pos.
func (g *Occupier) Move(p image.Point) bool {
	g.grid.mu.Lock()
	defer g.grid.mu.Unlock()

	if p.X < 0 || p.X > g.grid.bounds.X || p.Y < 0 || p.Y > g.grid.bounds.Y {
		return false
	} else if g.grid.grid[p.X][p.Y].occupied {
		return false
	}
	// unoccupy current position
	g.grid.grid[p.X][p.Y].occupied = false
	g.pos = p
	g.grid.grid[p.X][p.Y].occupied = true
	return true
}

func NewMarsGrid(size int) *MarsGrid {
	c := cord{X: size, Y: size}
	g := &MarsGrid{
		bounds: c,
		grid:   make([][]cell, size+1),
	}
	return g
}

func main() {
	grid := NewMarsGrid(24)

	num_rover = 1
	r1 := NewRoverDriver(grid)
	r1.Left()
	time.Sleep(2 * time.Second)
	r2 := NewRoverDriver(grid)
	r2.RandMove()
	time.Sleep(time.Second)
	r3 := NewRoverDriver(grid)
	r3.RandMove()
	time.Sleep(time.Second)
	r2.RandMove()
	r1.Stop()
	time.Sleep(2 * time.Second)
	r3.RandMove()
	r1.Start()
	time.Sleep(time.Second)
	r1.Right()
	time.Sleep(3 * time.Second)
}

// RoverDriver drives a rover around the surface of Mars.
type RoverDriver struct {
	id       int
	commandc chan command
}

// NewRoverDriver starts a new RoverDriver and returns it.
func NewRoverDriver(g *MarsGrid) *RoverDriver {
	r := &RoverDriver{
		id:       num_rover,
		commandc: make(chan command),
	}
	num_rover += 1
	go r.drive()
	return r
}

type command int

const (
	right     = command(0)
	left      = command(1)
	start     = command(2)
	stop      = command(3)
	rand_move = command(4)
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
			case rand_move:
				prev_cmnd = command(rand.Intn(2))
			}

			if (prev_cmnd == left) || (prev_cmnd == right) {
				log.Println("rover", r.id, " new direction ", direction)
			} else if prev_cmnd == stop {
				log.Println("rover", r.id, " stopped")
			} else if prev_cmnd == start {
				log.Println("rover", r.id, " started")
			}

		case <-nextMove:
			pos = pos.Add(direction)
			if (prev_cmnd == left) || (prev_cmnd == right) {
				log.Println(r.id, " moved to ", pos)
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

func (r *RoverDriver) RandMove() {
	r.commandc <- rand_move
}
