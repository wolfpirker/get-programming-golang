package main

import "fmt"

// TODO: function to display board nicely

func main() {
	var board [8][8]rune

	// black pieces
	board[0][0] = 'r'
	board[0][1] = 'n'
	board[0][2] = 'b'
	board[0][3] = 'q'
	board[0][4] = 'k'
	board[0][5] = 'b'
	board[0][6] = 'n'
	board[0][7] = 'r'

	for column := range board[1] {
		board[1][column] = 'p'
		board[6][column] = 'p'
	}

	// white pieces
	for i := 0; i <= 7; i++ {
		board[7][i] = board[0][i] + 'A' - 'a'
	}

	// board[7][0] = 'R'
	// board[7][1] = 'N'
	// board[7][2] = 'B'
	// board[7][3] = 'Q'
	// board[7][4] = 'K'
	// board[7][5] = 'B'
	// board[7][6] = 'N'
	// board[7][7] = 'R'

	// TODO: function to display board nicely
	for _, row := range board {
		fmt.Println(row)
	}
	//fmt.Print(board)
}
