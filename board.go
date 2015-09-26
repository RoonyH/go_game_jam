package chuckablast

const (
	pInvalid, pEmpty, pFull = 0, 1, 2
)

// Point represents a point in the board
type Point struct {
	state int
}

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
