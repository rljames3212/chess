package pieces

import (
	"chess/board/location"
	"testing"
)

func TestUnmovedUnblockedBlackPawn(t *testing.T) {
	loc := location.Location{Row: 7, Col: 1}
	p := NewPawn(loc, BLACK)

	var pcs []Piece

	validMoves := p.ValidMoves(pcs)
	expectedMoves := []location.Location{
		{Row: 6, Col: 1},
		{Row: 5, Col: 1},
	}

	evaluate(validMoves, expectedMoves, t)
}

func TestUnmovedUnblockedWhitePawn(t *testing.T) {
	loc := location.Location{Row: 1, Col: 1}
	p := NewPawn(loc, WHITE)

	var pcs []Piece

	validMoves := p.ValidMoves(pcs)
	expectedMoves := []location.Location{
		{Row: 2, Col: 1},
		{Row: 3, Col: 1},
	}

	evaluate(validMoves, expectedMoves, t)
}

func TestUnmovedBlockedBlackPawn(t *testing.T) {
	loc := location.Location{Row: 6, Col: 2}
	p := NewPawn(loc, BLACK)
	p1 := NewPawn(location.Location{Row: 5, Col: 2}, WHITE)
	pcs := []Piece{p1}

	validMoves := p.ValidMoves(pcs)
	var expectedMoves []location.Location

	evaluate(validMoves, expectedMoves, t)
}

func TestUnmovedBlockedWhitePawn(t *testing.T) {
	loc := location.Location{Row: 1, Col: 2}
	p := NewPawn(loc, WHITE)
	p1 := NewPawn(location.Location{Row: 2, Col: 2}, WHITE)
	pcs := []Piece{p1}

	validMoves := p.ValidMoves(pcs)
	var expectedMoves []location.Location

	evaluate(validMoves, expectedMoves, t)
}

func TestMovedUnblockedBlackPawn(t *testing.T) {
	p := NewPawn(location.Location{Row: 6, Col: 0}, BLACK)
	p.Move(location.Location{Row: 5, Col: 0})

	var pcs []Piece
	validMoves := p.ValidMoves(pcs)

	expectedmoves := []location.Location{
		{Row: 4, Col: 0},
	}

	evaluate(validMoves, expectedmoves, t)
}

func TestMovedUnblockedWhitePawn(t *testing.T) {
	p := NewPawn(location.Location{Row: 1, Col: 0}, WHITE)
	p.Move(location.Location{Row: 2, Col: 0})

	var pcs []Piece
	validMoves := p.ValidMoves(pcs)

	expectedmoves := []location.Location{
		{Row: 3, Col: 0},
	}

	evaluate(validMoves, expectedmoves, t)
}

func TestMovedBlockedBlackPawn(t *testing.T) {
	p := NewPawn(location.Location{Row: 7, Col: 0}, BLACK)
	p.Move(location.Location{Row: 6, Col: 0})

	pcs := []Piece{
		NewPawn(location.Location{Row: 5, Col: 0}, BLACK),
	}

	validMoves := p.ValidMoves(pcs)

	var expectedMoves []location.Location

	evaluate(validMoves, expectedMoves, t)
}

func TestMovedBlockedWhitePawn(t *testing.T) {
	p := NewPawn(location.Location{Row: 1, Col: 0}, WHITE)
	p.Move(location.Location{Row: 2, Col: 0})

	pcs := []Piece{
		NewPawn(location.Location{Row: 3, Col: 0}, BLACK),
	}

	validMoves := p.ValidMoves(pcs)

	var expectedMoves []location.Location

	evaluate(validMoves, expectedMoves, t)
}

func TestCaptureBlackPawn(t *testing.T) {
	p := NewPawn(location.Location{Row: 6, Col: 1}, BLACK)

	pcs := []Piece{
		NewPawn(location.Location{Row: 5, Col: 2}, WHITE),
		NewPawn(location.Location{Row: 5, Col: 0}, BLACK),
		NewPawn(location.Location{Row: 4, Col: 1}, WHITE),
		NewPawn(location.Location{Row: 7, Col: 2}, WHITE),
	}

	validMoves := p.ValidMoves(pcs)

	expectedMoves := []location.Location{
		{Row: 5, Col: 2},
		{Row: 5, Col: 1},
	}

	evaluate(validMoves, expectedMoves, t)
}

func TestCaptureWhitePawn(t *testing.T) {
	p := NewPawn(location.Location{Row: 1, Col: 1}, WHITE)

	pcs := []Piece{
		NewPawn(location.Location{Row: 2, Col: 2}, WHITE),
		NewPawn(location.Location{Row: 2, Col: 0}, BLACK),
		NewPawn(location.Location{Row: 4, Col: 1}, WHITE),
		NewPawn(location.Location{Row: 7, Col: 2}, WHITE),
	}

	validMoves := p.ValidMoves(pcs)

	expectedMoves := []location.Location{
		{Row: 2, Col: 0},
		{Row: 2, Col: 1},
		{Row: 3, Col: 1},
	}

	evaluate(validMoves, expectedMoves, t)
}

func TestKnight(t *testing.T) {
	k := NewKnight(location.Location{Row: 2, Col: 6}, WHITE)

	pcs := []Piece{
		NewPawn(location.Location{Row: 2, Col: 5}, WHITE),
		NewPawn(location.Location{Row: 1, Col: 4}, BLACK),
		NewPawn(location.Location{Row: 3, Col: 4}, WHITE),
	}

	validMoves := k.ValidMoves(pcs)

	expectedMoves := []location.Location{
		{Row: 0, Col: 7},
		{Row: 4, Col: 7},
		{Row: 0, Col: 5},
		{Row: 4, Col: 5},
		{Row: 1, Col: 4},
	}

	evaluate(validMoves, expectedMoves, t)
}

func evaluate(moves []location.Location, expectedMoves []location.Location, t *testing.T) {
	t.Helper()

	if len(moves) != len(expectedMoves) {
		t.Errorf("ValidMoves length does not match ExpectedMoves length: %d = %d ", len(expectedMoves), len(moves))
	}

	for _, move := range moves {
		if !contains(expectedMoves, move) {
			t.Errorf("Unexpected move: %v", move)
		}
	}
}

func contains(moves []location.Location, loc location.Location) bool {
	for _, move := range moves {
		if loc.Equals(move) {
			return true
		}
	}
	return false
}
