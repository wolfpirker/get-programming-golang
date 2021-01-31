package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// Message sent from Mars to Earth
type Message struct {
	Pos  cord
	Life int
	Rid  int // rover id
}

type cell struct {
	life     int // rand number between 0-1000
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
func (g *MarsGrid) Occupy(p cord) *Occupier {
	g.mu.Lock()
	defer g.mu.Unlock()
	occ := &Occupier{
		grid: g,
		Pos:  p,
	}
	if occ == nil {
		return nil
	}
	occ = &Occupier{
		grid: g,
		Pos:  p,
	}
	return occ
}

// Occupier represents a occupied cell in the grid
type Occupier struct {
	grid *MarsGrid
	Pos  cord
}

// Size returns a Point representing the size of the grid
func (g *MarsGrid) Size() cord {
	return cord{X: g.bounds.X, Y: g.bounds.Y}
}

func startDriver(id int, grid *MarsGrid) *RoverDriver {
	var o *Occupier
	// get a random point, continue until one is
	// found which is not occupied
	for o == nil {
		startPoint := cord{X: rand.Intn(grid.Size().X), Y: rand.Intn(grid.Size().Y)}
		o = grid.Occupy(startPoint)
	}
	return NewRoverDriver(id, o)
}

// Move moves the occupier to a different cell in the grid
// reports back whether the move was successful
// could fail due to moving out of the grid or
// cell already occupied - for fail: occupier stay at pos.
func (g *Occupier) MoveTo(p cord) bool {
	g.grid.mu.Lock()
	defer g.grid.mu.Unlock()

	// FIXME!
	// PANIC: index out of range!
	// if p.X < 0 || p.X > g.grid.bounds.X || p.Y < 0 || p.Y > g.grid.bounds.Y {
	// 	return false
	// } else if g.grid.grid[p.X][p.Y].occupied {
	// 	return false
	// }
	// // unoccupy current position
	// g.grid.grid[p.X][p.Y].occupied = false
	// g.Pos = p
	// g.grid.grid[p.X][p.Y].occupied = true

	return true
}

func NewMarsGrid(size int) *MarsGrid {
	c := cord{X: size, Y: size}
	g := &MarsGrid{
		bounds: c,
		grid:   make([][]cell, c.Y+1),
	}
	return g
}

func main() {
	grid := NewMarsGrid(15)
	rover := make([]*RoverDriver, 3)
	for i := range rover {
		rover[i] = startDriver(i, grid)
	}
	time.Sleep(60 * time.Second)
}

// RoverDriver drives a rover around the surface of Mars.
type RoverDriver struct {
	id       int
	commandc chan command
	occupier *Occupier
}

// NewRoverDriver starts a new RoverDriver and returns it.
func NewRoverDriver(n int, o *Occupier) *RoverDriver {
	r := &RoverDriver{
		id:       n,
		commandc: make(chan command),
		occupier: o,
	}
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
	send_msg  = command(5)
)

// drive is responsible for driving the rover. It
// is expected to be started in a goroutine.
func (r *RoverDriver) drive() {
	direction := cord{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = cord{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				direction = cord{
					X: direction.Y,
					Y: -direction.X,
				}
			}

			log.Printf("%v new direction %v", r.id, direction)
		case <-nextMove:
			nextMove = time.After(updateInterval)
			newPos := r.occupier.Pos.Add(direction)
			if r.occupier.MoveTo(newPos) {
				log.Printf("%v moved to %v", r.id, newPos)
				break
			}
			log.Printf("%v blocked trying to move from %v to %v", r.id, r.occupier.Pos, newPos)
			// Pick one other direction randomly
			// next cycle try to move in the new direction
			dir := rand.Intn(3) + 1
			for i := 0; i < dir; i++ {
				direction = cord{
					X: -direction.Y,
					Y: direction.X,
				}
			}
			log.Printf("%v new random direction %v", r.id, direction)

		}
	}
}

func (c *cord) Add(d cord) cord {
	return cord{X: c.X + d.X, Y: c.Y + d.Y}
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
