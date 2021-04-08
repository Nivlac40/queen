package main

import (
	"fmt"
)

type space int

const (
	empty space = iota
	queen
	hit
)

func main() {
	recurse([8][8]space{{0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0}}, 0)
}

func recurse(board [8][8]space, row int) bool {

	if row >= 8 {
		printBoard(board)
		return true
	}

	for col := 0; col < 8; col++ {
		if board[row][col] == empty {
			bc := board
			mark(&bc, row, col)
			if recurse(bc, row+1) {
				return true
			}
		}
	}

	return false
}

func mark(board *[8][8]space, x, y int) {
	board[x][y] = queen

	for _, xdelta := range []int{-1, 0, 1} {
		for _, ydelta := range []int{-1, 0, 1} {
			_mark(board, x, y, xdelta, ydelta)
		}
	}
}

func _mark(board *[8][8]space, x int, y int, xdelta int, ydelta int) {
	if xdelta == 0 && ydelta == 0 {
		return
	}

	for {
		x += xdelta
		y += ydelta
		if y < 0 || y >= 8 || x < 0 || x >= 8 {
			return
		}
		board[x][y] = hit
	}
}

func printBoard(board [8][8]space) {
	for _, spots := range board {
		for _, s := range spots {
			switch s {
			case empty:
				fmt.Print("# ")
			case hit:
				fmt.Print("x ")
			case queen:
				fmt.Print("Q ")
			default:
				panic("Invalid Space ID")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
