package main

import (
	"github.com/JoelOtter/termloop"
	"github.com/RoonyH/chuckablast"
)

const (
	pl = 3
	pw = 6
)

func main() {
	board := chuckablast.NewBoard()
	game := termloop.NewGame()

	level := termloop.NewBaseLevel(termloop.Cell{})
	game.Screen().SetLevel(level)

	drawBoard(level, board)

	game.Start()
}

func drawBoard(level termloop.Level, board *chuckablast.Board) {
	points := board.Points()

	for i := 0; i <= 12; i++ {
		for j := 0; j <= 12; j++ {
			if points[i][j] == chuckablast.PInvalid {
				drawInvalidP(level, i, j)
			} else if points[i][j] == chuckablast.PFull {
				drawFullP(level, i, j)
			}
		}
	}
}

func drawInvalidP(level termloop.Level, x int, y int) {
	level.AddEntity(
		termloop.NewRectangle(x*pw, y*pl, pw, pl, termloop.ColorWhite))
}

func drawEmptyP(level termloop.Level, x int, y int) {
	level.AddEntity(
		termloop.NewRectangle(x*pw, y*pl, pw, pl, termloop.ColorBlack))
}

func drawFullP(level termloop.Level, x int, y int) {
	level.AddEntity(
		termloop.NewRectangle(x*pw+1, y*pl+1, pw-2, pl-2, termloop.ColorWhite))
}
