package main

import (
	"errors"
	"fmt"
	"./color"
)

const (
	rows, columns = 9, 9
	empty = 0
)

type cell struct {
	digit int8
	fixed bool
}

// Grid is a Sudoku grid
type Grid [rows][columns] cell

// Errors that could occur.
var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
	ErrFixed  = errors.New("digit is fixed")
	ErrInRow  = errors.New("digit already in this row")
	ErrInColumn  = errors.New("digit already in this column")
	ErrInRegion  = errors.New("digit already in this region")
)

// NewSudoku makes a new Sudoku grid
func NewSudoku(digits [rows][columns] int8) *Grid {
	var grid Grid
	for r:=0; r<rows; r++ {
		for c := 0; c < columns; c++ {
			d := digits[r][c]
			if d != empty {
				grid[r][c].digit = d
				grid[r][c].fixed = true
			}
		}
	}
	return &grid
}

func (g *Grid) isFixed(row int8, column int8) bool{
	return g[row][column].fixed
}

func (g *Grid) inRow(row int8, digit int8) bool{
	for c := 0; c < columns; c++ {
		if g[row][c].digit == digit {
			return true
		}
	}
	return false
}

func (g *Grid) inColumn(column int8, digit int8) bool{
	for r := 0; r < columns; r++ {
		if g[r][column].digit == digit {
			return true
		}
	}
	return false
}

func (g *Grid) inRegion(row int8, column int8, digit int8) bool{
	sir := row - row % 3 // startindex row
	sic := column - column % 3 // startindex column
	for r := sir; r < sir+3; r++ {
		for c := sic; c < sic+3; c++ {
			if g[r][c].digit == digit {
				return true
			}
		}		
	}
	return false
}

// Set a digit on a Sudoku grid
func (g *Grid) Set(row, column int8, digit int8) error {
	if !inBounds(row, column) {
		return ErrBounds
	}
	if !validDigit(digit) {
		return ErrDigit
	}
	if (g.isFixed(row, column)){
		return ErrFixed
	}
	if (g.inRow(row, digit)){
		return ErrInRow
	}
	if (g.inColumn(column, digit)){
		return ErrInColumn
	}
	if (g.inRegion(row, column, digit)){
		return ErrInRegion
	}

	g[row][column].digit = digit
	return nil
}

// Clear a cell if not fixed
func (g *Grid) Clear(row, column int8) error{
	if !inBounds(row, column){
		return ErrBounds
	} else if g.isFixed(row, column){
		return ErrFixed
	}
	g[row][column].digit = empty
	return nil
}

func inBounds(row, column int8) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func (g *Grid) String() string {
	output := ""
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++{
			if g[r][c].fixed {
				output += fmt.Sprintf("%v%v,%v", color.Cyan, g[r][c].digit, color.Reset)
			} else if g[r][c].digit == 0 {
				output += fmt.Sprintf("%v%v,%v", color.Red, g[r][c].digit, color.Reset)
			} else {
				output += fmt.Sprintf("%v,", g[r][c].digit)
			}
			if c % 3 == 2 {
				output += fmt.Sprintf("|")
			}
		}
		output += fmt.Sprintf("\n")
		if (r % 3) == 2 && r < 8 {
			output += fmt.Sprintf("____________________\n")
		}
	}
	return fmt.Sprintf("%v\n", output)
}

func (g *Grid) hasEmptyCell() bool {
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++{
			if g[r][c].digit == 0 {
				return true
			}
		}
	}
	return false
}

// Backtrack Sudoku (Solver algorithm)
func (g *Grid) Backtrack() bool {
	if !g.hasEmptyCell() {
		return true
	}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if g[r][c].digit == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					err := g.Set(int8(r), int8(c), int8(candidate))
					if err == nil {
						if g.Backtrack() {
							return true
						}
						g[r][c].digit = 0
					} else {
						g[r][c].digit = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func main() {	
	s := NewSudoku([rows][columns] int8{
		{0,3,0, 0,0,0, 0,0,0},
		{0,0,0, 1,9,5, 0,0,0},
		{0,0,8, 0,0,0, 0,6,0},
		{8,0,0, 0,6,0, 0,0,0},
		{4,0,0, 8,0,0, 0,0,1},
		{0,0,0, 0,2,0, 0,0,0},
		{0,6,0, 0,0,0, 2,8,0},
		{0,0,0, 4,1,9, 0,0,5},
		{0,0,0, 0,0,0, 0,7,0},
	})
	
	s.Backtrack()

	fmt.Println(s)
}
