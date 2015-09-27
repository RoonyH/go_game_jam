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
		{0, 0, PInvalid},
		{12, 11, PInvalid},
		{3, 3, PInvalid},
		{9, 3, PInvalid},
		{9, 10, PInvalid},
		{3, 9, PInvalid},
		{6, 2, PFull},
		{9, 7, PFull},
		{3, 6, PFull},
		{6, 10, PFull},
		{10, 5, PFull},
		{6, 7, PFull},
		{6, 6, PEmpty},
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
	testBoard.b[4][5] = PEmpty
	testBoard.b[5][5] = PEmpty
	testBoard.b[4][6] = PEmpty
	testBoard.b[5][6] = PEmpty
	testBoard.b[5][10] = PEmpty

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

func TestMove(t *testing.T) {
	testBoard := NewBoard()

	if testBoard == nil {
		t.Fatal("Nil board")
	}

	// Mark some points as empty
	testBoard.b[7][6] = PEmpty
	testBoard.b[7][4] = PEmpty

	tests := []struct {
		source  Point
		target  Point
		middle  Point
		success bool
	}{
		{Point{7, 6}, Point{7, 4}, Point{7, 5}, false},
		{Point{9, 5}, Point{9, 7}, Point{9, 6}, false},
		{Point{9, 5}, Point{3, 5}, Point{}, false},
		{Point{9, 6}, Point{7, 6}, Point{8, 6}, true},
		{Point{7, 2}, Point{7, 4}, Point{7, 3}, true},
		{Point{6, 4}, Point{6, 6}, Point{6, 5}, true},
	}

	for _, test := range tests {
		success := testBoard.Move(test.source, test.target)

		if success != test.success {
			t.Fatal("Wrong success state for", test.source, test.target,
				"Got", success, "Expected", test.success)
		}

		if !success {
			continue
		}

		if testBoard.b[test.middle[0]][test.middle[1]] != PEmpty {
			t.Fatal("Middle piece not removed", test.source, test.target)
		}
	}
}
