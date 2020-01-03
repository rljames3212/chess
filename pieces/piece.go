package pieces

import "chess/board/location"

const BOARD_SIZE int = 8

func isValidLocation(l location.Location) bool {
	return l.GetRow() >= 0 && l.GetRow() < BOARD_SIZE &&
		l.GetCol() >= 0 && l.GetCol() < BOARD_SIZE
}

func isLocationOccupied(loc location.Location, pcs []Piece) bool {
	for _, p := range pcs {
		if loc.Equals(p.GetLocation()) {
			return true
		}
	}
	return false
}

func isLocationOccupiedByOpponent(loc location.Location, thisColor PieceColor, pcs []Piece) bool {
	for _, p := range pcs {
		if p.GetColor() != thisColor && loc.Equals(p.GetLocation()) {
			return true
		}
	}
	return false
}

// PieceColor is the color of a piece
type PieceColor int32

const (
	BLACK PieceColor = iota
	WHITE
)

func (pc PieceColor) String() string {
	return [...]string{
		"Black",
		"White",
	}[pc]
}

// Piece represents a chess piece
type Piece interface {
	GetColor() PieceColor
	GetLocation() location.Location
	GetHasMoved() bool
	GetValidMoves([]Piece) []location.Location
	Move(location.Location)
}
