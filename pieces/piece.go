package pieces

import "chess/board/location"

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
	GetValidMoves([]*Piece) []location.Location
}
