package pieces

import "chess/board/location"

// Bishop represents a bishop chess piece
type Bishop struct {
	loc      location.Location
	color    PieceColor
	hasMoved bool
}

// NewBishop returns a pointer to a new bishop
func NewBishop(l location.Location, c PieceColor) *Bishop {
	return &Bishop{
		loc:      l,
		color:    c,
		hasMoved: false,
	}
}

// Location returns the location of the bishop
func (b *Bishop) Location() location.Location {
	return b.loc
}

// Color returns the color of the bishop
func (b *Bishop) Color() PieceColor {
	return b.color
}

// HasMoved returns whether the bishop has moved or not
func (b *Bishop) HasMoved() bool {
	return b.hasMoved
}

// Move moves the bishop to a new location and sets hasMoved to true
func (b *Bishop) Move(newLocation location.Location) {
	b.loc = newLocation
	b.hasMoved = true
}

// ValidMoves returns all of the current possible moves for the bishop
func (b *Bishop) ValidMoves(pcs []Piece) []location.Location {
	bearings := []bearing{
		{Row: 1, Col: 1},
		{Row: -1, Col: 1},
		{Row: 1, Col: -1},
		{Row: -1, Col: -1},
	}

	currentRow := b.loc.GetRow()
	currentCol := b.loc.GetCol()

	var validMoves []location.Location
	var loc location.Location
	isBlocked := false

	for _, bear := range bearings {
		loc = location.Location{Row: currentRow, Col: currentCol}
		isBlocked = false
		for !isBlocked {
			loc = location.Location{Row: loc.GetRow() + bear.Row, Col: loc.GetCol() + bear.Col}
			// check if location is valid
			if isValidLocation(loc) {
				// check if location is vacant
				if !isLocationOccupied(loc, pcs) {
					validMoves = append(validMoves, loc)
				} else {
					// check if location is occupied by opponent
					if isLocationOccupiedByOpponent(loc, b.color, pcs) {
						validMoves = append(validMoves, loc)
					}
					isBlocked = true
				}
			} else {
				// if location is invalid, piece is blocked by the edge of the board
				isBlocked = true
			}
		}
	}
	return validMoves
}
