package board

import "testing"

func TestNewBoard(t *testing.T) {
	b := NewBoard()

	expectedColor := BLACK
	for row := 0; row < len(*b); row++ {
		for col := 0; col < len((*b)[row]); col++ {
			if (*b)[row][col].GetColor() != expectedColor {
				t.Error("unexpected tileColor")
			}
			if expectedColor == BLACK {
				expectedColor = WHITE
			} else {
				expectedColor = BLACK
			}
		}
	}
}
