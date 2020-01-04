package pieces

import "chess/board/location"

// Knight represents a Knight chess piece
type Knight struct {
	loc      location.Location
	color    PieceColor
	hasMoved bool
}

// NewKnight returns a pointer to a new knight struct
func NewKnight(l location.Location, c PieceColor) *Knight {
	return &Knight{
		loc:      l,
		color:    c,
		hasMoved: false,
	}
}

// Location returns the location of the knight
func (k *Knight) Location() location.Location {
	return k.loc
}

// Color returns the color of the knight
func (k *Knight) Color() PieceColor {
	return k.color
}

// HasMoved returns whether or not the knight has moved
func (k *Knight) HasMoved() bool {
	return k.hasMoved
}

// Move moves the knight to a new location and sets hasMoved to true
func (k *Knight) Move(newLocation location.Location) {
	k.loc = newLocation
	k.hasMoved = true
}

// ValidMoves returns a slice of all of the knight's current possible moves
func (k *Knight) ValidMoves(pcs []Piece) []location.Location {
	bearings := []bearing{
		{Row: -1, Col: 2},
		{Row: 1, Col: 2},
		{Row: 2, Col: 1},
		{Row: 2, Col: -1},
		{Row: 1, Col: -2},
		{Row: -1, Col: -2},
		{Row: -2, Col: -1},
		{Row: -2, Col: 1},
	}

	currentRow := k.loc.GetRow()
	currentCol := k.loc.GetCol()

	var validMoves []location.Location
	var loc location.Location

	for _, b := range bearings {
		loc = location.Location{
			Row: currentRow + b.Row,
			Col: currentCol + b.Col,
		}
		// check if valid location
		if isValidLocation(loc) {
			// check if location is vacant or occupied by opponent
			if !isLocationOccupied(loc, pcs) || isLocationOccupiedByOpponent(loc, k.Color(), pcs) {
				validMoves = append(validMoves, loc)
			}
		}
	}

	return validMoves
}
