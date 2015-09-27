package main

import (
	"github.com/JoelOtter/termloop"
	"github.com/RoonyH/chuckablast"
)

const (
	pl = 3
	pw = 6
)

type point struct {
	outer *termloop.Rectangle
	inner *termloop.Rectangle
}

// Board is a board
type Board struct {
	*chuckablast.Board
	spots *[13][13]point

	selected   chuckablast.Point
	held       bool
	validMoves *[]chuckablast.Point
}

// Draw implements termloop.Drawable.Draw
func (board *Board) Draw(screen *termloop.Screen) {
	points := board.Points()

	for i := 0; i <= 12; i++ {
		for j := 0; j <= 12; j++ {
			if points[i][j] == chuckablast.PInvalid {
				board.drawInvalidP(screen, i, j)
			} else if points[i][j] == chuckablast.PFull {
				board.drawFullP(screen, i, j)
			} else if points[i][j] == chuckablast.PEmpty {
				board.drawEmptyP(screen, i, j)
			}
		}
	}

	if points[board.selected[0]][board.selected[1]] == chuckablast.PEmpty {
		board.drawSelectedEmpty(screen, board.selected[0], board.selected[1])
	} else if board.held {
		board.drawLocked(screen, board.selected[0], board.selected[1])
	} else {
		board.drawSelected(screen, board.selected[0], board.selected[1])
	}

	board.drawValidMoves(screen)
}

// Tick implements termloop.Drawable.Tick
func (board *Board) Tick(ev termloop.Event) {
	if ev.Type == termloop.EventKey {
		switch ev.Key {
		case termloop.KeyArrowRight:
			if board.held {
				if board.move("right") {
					break
				}
			}
			board.selectNextFull("right")
		case termloop.KeyArrowLeft:
			if board.held {
				if board.move("left") {
					break
				}
			}
			board.selectNextFull("left")
		case termloop.KeyArrowUp:
			if board.held {
				if board.move("up") {
					break
				}
			}
			board.selectNextFull("up")
		case termloop.KeyArrowDown:
			if board.held {
				if board.move("down") {
					break
				}
			}
			board.selectNextFull("down")
		case termloop.KeySpace:
			board.hold()
		}
	}
}

func (board *Board) build() {
	for i := 0; i <= 12; i++ {
		for j := 0; j <= 12; j++ {
			board.spots[i][j].outer = termloop.NewRectangle(
				i*pw, j*pl, pw, pl, termloop.ColorWhite)
			board.spots[i][j].inner = termloop.NewRectangle(
				i*pw+1, j*pl+1, pw-2, pl-2, termloop.ColorWhite)
		}
	}
}

// getNextFull returns the next occupied point in the given direction
func (board *Board) selectNextFull(direction string) {
	points := board.Points()
	switch direction {
	case "right":
		if board.selected[0] == 12 {
			board.selected[0] = 0
		} else {
			board.selected[0]++
		}
		if points[board.selected[0]][board.selected[1]] == chuckablast.PInvalid {
			board.selectNextFull("right")
		}
	case "left":
		if board.selected[0] == 0 {
			board.selected[0] = 12
		} else {
			board.selected[0]--
		}
		if points[board.selected[0]][board.selected[1]] == chuckablast.PInvalid {
			board.selectNextFull("left")
		}
	case "up":
		if board.selected[1] == 0 {
			board.selected[1] = 12
		} else {
			board.selected[1]--
		}
		if points[board.selected[0]][board.selected[1]] == chuckablast.PInvalid {
			board.selectNextFull("up")
		}
	case "down":
		if board.selected[1] == 12 {
			board.selected[1] = 0
		} else {
			board.selected[1]++
		}
		if points[board.selected[0]][board.selected[1]] == chuckablast.PInvalid {
			board.selectNextFull("down")
		}
	}
}

func (board *Board) move(direction string) bool {
	success := false
	switch direction {
	case "right":
		success = board.Move(board.selected,
			chuckablast.Point{board.selected[0] + 2, board.selected[1]})
		if success {
			board.selected[0] = board.selected[0] + 2
		}
	case "left":
		success = board.Move(board.selected,
			chuckablast.Point{board.selected[0] - 2, board.selected[1]})
		if success {
			board.selected[0] = board.selected[0] - 2
		}
	case "up":
		success = board.Move(board.selected,
			chuckablast.Point{board.selected[0], board.selected[1] - 2})
		if success {
			board.selected[1] = board.selected[1] - 2
		}
	case "down":
		success = board.Move(board.selected,
			chuckablast.Point{board.selected[0], board.selected[1] + 2})
		if success {
			board.selected[1] = board.selected[1] + 2
		}
	}
	board.hold()
	return success
}

func (board *Board) hold() {
	points := board.Points()
	if points[board.selected[0]][board.selected[1]] != chuckablast.PFull {
		return
	}

	if board.held {
		board.validMoves = &[]chuckablast.Point{}
	} else {
		board.validMoves = board.Select(board.selected[0], board.selected[1])
	}

	board.held = !board.held
}

func main() {
	b := chuckablast.NewBoard()
	board := &Board{b, &[13][13]point{}, chuckablast.Point{5, 2}, false,
		&[]chuckablast.Point{}}

	board.build()

	game := termloop.NewGame()

	level := termloop.NewBaseLevel(termloop.Cell{})
	game.Screen().SetLevel(level)

	level.AddEntity(board)

	game.Start()
}

func (board *Board) drawInvalidP(screen *termloop.Screen, x int, y int) {
	board.spots[x][y].outer.Draw(screen)
}

func (board *Board) drawEmptyP(screen *termloop.Screen, x int, y int) {
	board.spots[x][y].outer.SetColor(termloop.ColorBlack)
	board.spots[x][y].outer.Draw(screen)
}

func (board *Board) drawFullP(screen *termloop.Screen, x int, y int) {
	board.spots[x][y].outer.SetColor(termloop.ColorBlack)
	board.spots[x][y].outer.Draw(screen)
	board.spots[x][y].inner.SetColor(termloop.ColorWhite)
	board.spots[x][y].inner.Draw(screen)
}

func (board *Board) drawSelected(screen *termloop.Screen, x int, y int) {
	board.spots[x][y].outer.SetColor(termloop.ColorRed)
	board.spots[x][y].outer.Draw(screen)
	board.spots[x][y].inner.SetColor(termloop.ColorWhite)
	board.spots[x][y].inner.Draw(screen)
}

func (board *Board) drawLocked(screen *termloop.Screen, x int, y int) {
	board.spots[x][y].outer.SetColor(termloop.ColorRed)
	board.spots[x][y].outer.Draw(screen)
	board.spots[x][y].inner.SetColor(termloop.ColorBlue)
	board.spots[x][y].inner.Draw(screen)
}

func (board *Board) drawSelectedEmpty(screen *termloop.Screen, x int, y int) {
	board.spots[x][y].outer.SetColor(termloop.ColorRed)
	board.spots[x][y].outer.Draw(screen)
}

func (board *Board) drawValidMoves(screen *termloop.Screen) {
	for _, spot := range *board.validMoves {
		board.spots[spot[0]][spot[1]].outer.SetColor(termloop.ColorCyan)
		board.spots[spot[0]][spot[1]].outer.Draw(screen)
	}
}
