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

func TestSelect(t *testing.T) {
	testBoard := NewBoard()

	if testBoard == nil {
		t.Fatal("Nil board")
	}

	tests := []struct {
		x     int
		y     int
		moves [][2]int
	}{
		{6, 2, [][2]int{}},
		{2, 5, [][2]int{}},
		{5, 9, [][2]int{}},
		{4, 6, [][2]int{{6, 6}}},
		{6, 4, [][2]int{{6, 6}}},
	}

	for _, test := range tests {
		moves := *testBoard.Select(test.x, test.y)
		if len(moves) != len(test.moves) {
			t.Fatal("Not correct valid moves for", test.x, test.y, ". Got", moves,
				"Expected", test.moves)
		}

		for i := 0; i < len(moves); i++ {
			if moves[i] != test.moves[i] {
				t.Fatal("Not a correct move. Got", moves[i], "Expected", test.moves[i])
			}
		}
	}

	// Mark some points as empty
	testBoard.b[4][5] = pEmpty
	testBoard.b[5][5] = pEmpty
	testBoard.b[4][6] = pEmpty
	testBoard.b[5][6] = pEmpty
	testBoard.b[5][10] = pEmpty

	tests = []struct {
		x     int
		y     int
		moves [][2]int
	}{
		{6, 6, [][2]int{}},
		{6, 5, [][2]int{}},
		{5, 7, [][2]int{}},
		{2, 6, [][2]int{{4, 6}}},
		{5, 8, [][2]int{{5, 6}, {5, 10}}},
	}

	for _, test := range tests {
		moves := *testBoard.Select(test.x, test.y)
		if len(moves) != len(test.moves) {
			t.Fatal("Not correct valid moves for", test.x, test.y, ". Got", moves,
				"Expected", test.moves)
		}

		for i := 0; i < len(moves); i++ {
			if moves[i] != test.moves[i] {
				t.Fatal("Not a correct move for", test.x, test.y, "Got", moves[i],
					"Expected", test.moves[i])
			}
		}
	}
}
