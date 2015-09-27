package chuckablast

import (
	"errors"
	"fmt"
)

const (
	// PInvalid indicate a point is not part of game play
	PInvalid = 0

	// PEmpty indicate a point is free to occupy a piece
	PEmpty = 1

	// PFull indicate a point is occupied by a piece
	PFull = 2
)

// Point represents a point in the board
type Point [2]int

// Board represents chuck a blast board
type Board struct {
	b [13][13]int
}

// NewBoard creates a new board and put points in the traditional way
func NewBoard() (board *Board) {
	board = new(Board)
	initb(&(board.b))
	return
}

func initb(b *[13][13]int) {
	for _, i := range []int{2, 3, 4, 8, 9, 10} {
		(*b)[i][5] = PFull
		(*b)[i][6] = PFull
		(*b)[i][7] = PFull
	}

	for i := 2; i <= 10; i++ {
		(*b)[5][i] = PFull
		(*b)[6][i] = PFull
		(*b)[7][i] = PFull
	}

	(*b)[6][6] = PEmpty //right at the middle its an empty point
}

// Points returns the arrangement of the board
func (board *Board) Points() (points *[13][13]int) {
	return &board.b
}

// Select returns valid targets from the selected point
func (board *Board) Select(x int, y int) (valid *[][2]int) {

	moves := [][2]int{}

	if board.b[x][y] != PFull {
		//Source point is not full
		return &moves
	}

	if board.b[x][y-1] == PFull && board.b[x][y-2] == PEmpty {
		// Upward possible
		moves = append(moves, [2]int{x, y - 2})
	}

	if board.b[x][y+1] == PFull && board.b[x][y+2] == PEmpty {
		// Downward possible
		moves = append(moves, [2]int{x, y + 2})
	}

	if board.b[x+1][y] == PFull && board.b[x+2][y] == PEmpty {
		// Leftward possible
		moves = append(moves, [2]int{x + 2, y})
	}

	if board.b[x-1][y] == PFull && board.b[x-2][y] == PEmpty {
		// Rightward possible
		moves = append(moves, [2]int{x - 2, y})
	}

	return &moves
}

// Move moves a piece from source to target
// It returns true only if the move is successful
func (board *Board) Move(source Point, target Point) bool {
	if board.b[source[0]][source[1]] != PFull ||
		board.b[target[0]][target[1]] != PEmpty {
		// source has to be full and target has to be empty
		return false
	}

	middle, err := validateMove(source, target)

	if err != nil {
		fmt.Println(err)
		return false
	}

	if board.b[middle[0]][middle[1]] != PFull {
		// middle has to be full
		return false
	}

	board.b[middle[0]][middle[1]] = PEmpty

	return true
}

func validateMove(source Point, target Point) (middle Point, err error) {
	if source[0] == target[0] {
		// Move is on y axis
		if target[1] == source[1]+2 {
			// Move is upward
			return Point{source[0], source[1] + 1}, nil
		} else if target[1] == source[1]-2 {
			// Move is downward
			return Point{source[0], source[1] - 1}, nil
		}
	} else if source[1] == target[1] {
		// Move is on x axis
		if target[0] == source[0]+2 {
			// Move is leftward
			return Point{source[0] + 1, source[1]}, nil
		} else if target[0] == source[0]-2 {
			// Move is rightward
			return Point{source[0] - 1, source[1]}, nil
		}
	}

	return Point{0, 0}, errors.New("invalid move")
}
