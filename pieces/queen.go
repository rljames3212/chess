package pieces

import "chess/board/location"

// Queen represents a queen chess piece
type Queen struct {
	loc      location.Location
	color    PieceColor
	hasMoved bool
}

// NewQueen returns a pointer to a new queen
func NewQueen(l location.Location, c PieceColor) *Queen {
	return &Queen{
		loc:      l,
		color:    c,
		hasMoved: false,
	}
}

// Location returns the location of the queen
func (q *Queen) Location() location.Location {
	return q.loc
}

// Color returns the color of the queen
func (q *Queen) Color() PieceColor {
	return q.color
}

// HasMoved returns whether or not the queen has moved
func (q *Queen) HasMoved() bool {
	return q.hasMoved
}

// Move moves the queen to a new location and sets hasMoved to true
func (q *Queen) Move(newLocation location.Location) {
	q.loc = newLocation
	q.hasMoved = true
}

// ValidMoves returns all of the possible moves that the queen can currently make
func (q *Queen) ValidMoves(pcs []Piece) []location.Location {
	bearings := []bearing{
		{Row: 0, Col: 1},
		{Row: 1, Col: 1},
		{Row: 1, Col: 0},
		{Row: -1, Col: 1},
		{Row: 0, Col: -1},
		{Row: -1, Col: -1},
		{Row: -1, Col: 0},
		{Row: 1, Col: -1},
	}

	currentRow := q.loc.GetRow()
	currentCol := q.loc.GetCol()

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
				// check if space is vacant
				if !isLocationOccupied(loc, pcs) {
					validMoves = append(validMoves, loc)
				} else {
					// check if location is occupied by opponent
					if isLocationOccupiedByOpponent(loc, q.color, pcs) {
						validMoves = append(validMoves, loc)
					}
					isBlocked = true
				}
			} else {
				// if invalid location, piece is blocked by the edge of the board
				isBlocked = true
			}
		}
	}
	return validMoves
}
