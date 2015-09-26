package chuckablast

import (
	"errors"
	"fmt"
)

const (
	pInvalid, pEmpty, pFull = 0, 1, 2
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
		(*b)[i][5] = pFull
		(*b)[i][6] = pFull
		(*b)[i][7] = pFull
	}

	for i := 2; i <= 10; i++ {
		(*b)[5][i] = pFull
		(*b)[6][i] = pFull
		(*b)[7][i] = pFull
	}

	(*b)[6][6] = pEmpty //right at the middle its an empty point
}

// Select returns valid targets from the selected point
func (board *Board) Select(x int, y int) (valid *[][2]int) {

	moves := [][2]int{}

	if board.b[x][y] != pFull {
		//Source point is not full
		return &moves
	}

	if board.b[x][y-1] == pFull && board.b[x][y-2] == pEmpty {
		// Upward possible
		moves = append(moves, [2]int{x, y - 2})
	}

	if board.b[x][y+1] == pFull && board.b[x][y+2] == pEmpty {
		// Downward possible
		moves = append(moves, [2]int{x, y + 2})
	}

	if board.b[x+1][y] == pFull && board.b[x+2][y] == pEmpty {
		// Leftward possible
		moves = append(moves, [2]int{x + 2, y})
	}

	if board.b[x-1][y] == pFull && board.b[x-2][y] == pEmpty {
		// Rightward possible
		moves = append(moves, [2]int{x - 2, y})
	}

	return &moves
}

// Move moves a piece from source to target
// It returns true only if the move is successful
func (board *Board) Move(source Point, target Point) bool {
	if board.b[source[0]][source[1]] != pFull ||
		board.b[target[0]][target[1]] != pEmpty {
		// source has to be full and target has to be empty
		return false
	}

	middle, err := validateMove(source, target)

	if err != nil {
		fmt.Println(err)
		return false
	}

	if board.b[middle[0]][middle[1]] != pFull {
		// middle has to be full
		return false
	}

	board.b[middle[0]][middle[1]] = pEmpty

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
