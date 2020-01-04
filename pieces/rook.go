package pieces

import "chess/board/location"

// Rook represents a rook chess piece
type Rook struct {
	loc      location.Location
	color    PieceColor
	hasMoved bool
}

// NewRook returns a pointer to a new rook
func NewRook(l location.Location, c PieceColor) *Rook {
	return &Rook{
		loc:      l,
		color:    c,
		hasMoved: false,
	}
}

// Location returns the location of the rook
func (r *Rook) Location() location.Location {
	return r.loc
}

// Color returns the color of the rook
func (r *Rook) Color() PieceColor {
	return r.color
}

// HasMoved returns whether or not the rook has moved
func (r *Rook) HasMoved() bool {
	return r.hasMoved
}

// Move moves the rook to a new location and sets hasMoved to true
func (r *Rook) Move(newLocation location.Location) {
	r.loc = newLocation
	r.hasMoved = true
}

// ValidMoves returns all of the locations the rook can currently move to
func (r *Rook) ValidMoves(pcs []Piece) []location.Location {
	bearings := []bearing{
		{Row: 0, Col: 1},
		{Row: 1, Col: 0},
		{Row: 0, Col: -1},
		{Row: -1, Col: 0},
	}

	currentRow := r.loc.GetRow()
	currentCol := r.loc.GetCol()

	var validMoves []location.Location
	var loc location.Location
	isBlocked := false

	for _, b := range bearings {
		isBlocked = false
		loc = location.Location{Row: currentRow, Col: currentCol}
		for !isBlocked {
			loc = location.Location{Row: loc.GetRow() + b.Row, Col: loc.GetCol() + b.Col}
			// check if location is valid
			if isValidLocation(loc) {
				// check if location is vacant
				if !isLocationOccupied(loc, pcs) {
					validMoves = append(validMoves, loc)
				} else {
					// check if location is occupied by opponent
					if isLocationOccupiedByOpponent(loc, r.color, pcs) {
						validMoves = append(validMoves, loc)
					}
					isBlocked = true
				}
			} else {
				// if the location is invalid, piece is blocked by the edge of the board
				isBlocked = true
			}
		}
	}
	return validMoves
}
