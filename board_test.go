package chuckablast

import "testing"

func TestNewBoard(t *testing.T) {
	testBoard := NewBoard()

	if testBoard == nil {
		t.Fatal("Nil board")
	}

	tests := []struct {
		x     int
		y     int
		state int
	}{
		{0, 0, pInvalid},
		{12, 11, pInvalid},
		{3, 3, pInvalid},
		{9, 3, pInvalid},
		{9, 10, pInvalid},
		{3, 9, pInvalid},
		{6, 2, pFull},
		{9, 7, pFull},
		{3, 6, pFull},
		{6, 10, pFull},
		{10, 5, pFull},
		{6, 7, pFull},
		{6, 6, pEmpty},
	}

	for _, test := range tests {
		if testBoard.b[test.x][test.y] != test.state {
			t.Fatal("Wrong state in", test.x, test.y, ". Got",
				testBoard.b[test.x][test.y])
		}
	}
}
