package main

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

// Point represents a point in the game
type Point [2]int

// Game represents chuck a blast game
type Game struct {
	b [13][13]int
}

// NewGame creates a new game and put points in the traditional way
func NewGame() (game *Game) {
	game = new(Game)
	initb(&(game.b))
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
func (game *Game) Points() (points *[13][13]int) {
	return &game.b
}

// Select returns valid targets from the selected point
func (game *Game) Select(x int, y int) (valid *[]Point) {

	moves := []Point{}

	if game.b[x][y] != PFull {
		//Source point is not full
		return &moves
	}

	if game.b[x][y-1] == PFull && game.b[x][y-2] == PEmpty {
		// Upward possible
		moves = append(moves, Point{x, y - 2})
	}

	if game.b[x][y+1] == PFull && game.b[x][y+2] == PEmpty {
		// Downward possible
		moves = append(moves, [2]int{x, y + 2})
	}

	if game.b[x+1][y] == PFull && game.b[x+2][y] == PEmpty {
		// Leftward possible
		moves = append(moves, [2]int{x + 2, y})
	}

	if game.b[x-1][y] == PFull && game.b[x-2][y] == PEmpty {
		// Rightward possible
		moves = append(moves, [2]int{x - 2, y})
	}

	return &moves
}

// Move moves a piece from source to target
// It returns true only if the move is successful
func (game *Game) Move(source Point, target Point) bool {
	if game.b[source[0]][source[1]] != PFull ||
		game.b[target[0]][target[1]] != PEmpty {
		// source has to be full and target has to be empty
		return false
	}

	middle, err := validateMove(source, target)

	if err != nil {
		fmt.Println(err)
		return false
	}

	if game.b[middle[0]][middle[1]] != PFull {
		// middle has to be full
		return false
	}

	game.b[middle[0]][middle[1]] = PEmpty
	game.b[source[0]][source[1]] = PEmpty
	game.b[target[0]][target[1]] = PFull

	return true
}

// Test tests if the game is over
func (game *Game) Test() bool {
	for i := 0; i <= 12; i++ {
		for j := 0; j <= 12; j++ {
			if game.b[i][j] == PFull {
				if len(*(game.Select(i, j))) > 0 {
					return false
				}
			}
		}
	}

	return true
}

// GetRemaining returns number of pieces on the board
func (game *Game) GetRemaining() int {
	remaining := 0
	for i := 0; i <= 12; i++ {
		for j := 0; j <= 12; j++ {
			if game.b[i][j] == PFull {
				remaining++
			}
		}
	}

	return remaining
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
