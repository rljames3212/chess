package pieces

import "chess/board/location"

// Pawn represents a chess pawn
type Pawn struct {
	loc      location.Location
	hasMoved bool
	color    PieceColor
}

const (
	BLACK_DIRECTION int = -1
	WHITE_DIRECTION int = 1
)

// NewPawn creates and returns a new pawn
func NewPawn(l location.Location, pc PieceColor) *Pawn {
	return &Pawn{
		loc:      l,
		color:    pc,
		hasMoved: false,
	}
}

// Color returns the color of the pawn
func (p *Pawn) Color() PieceColor {
	return p.color
}

// HasMoved returns whether or not the piece has moved
func (p *Pawn) HasMoved() bool {
	return p.hasMoved
}

// Location returns the location of the pawn on the board
func (p *Pawn) Location() location.Location {
	return p.loc
}

// ValidMoves returns a slice of all moves the pawn can make given a piece configuration
func (p *Pawn) ValidMoves(pcs []Piece) []location.Location {
	currentRow := p.loc.GetRow()
	currentCol := p.loc.GetCol()

	var validMoves []location.Location

	var movementDirection int
	if p.color == BLACK {
		movementDirection = BLACK_DIRECTION
	} else {
		movementDirection = WHITE_DIRECTION
	}

	var loc location.Location
	// check locations diagonally in front of pawn to see if a capture can be made
	// check first diagonal
	loc = location.Location{Row: currentRow + movementDirection, Col: currentCol + 1}
	if isValidLocation(loc) && isLocationOccupiedByOpponent(loc, p.Color(), pcs) {
		validMoves = append(validMoves, loc)
	}
	// check second diagonal
	loc = location.Location{Row: currentRow + movementDirection, Col: currentCol - 1}
	if isValidLocation(loc) && isLocationOccupiedByOpponent(loc, p.Color(), pcs) {
		validMoves = append(validMoves, loc)
	}
	// check location directly in front of pawn
	loc = location.Location{Row: currentRow + movementDirection, Col: currentCol}
	if isValidLocation(loc) && !isLocationOccupied(loc, pcs) {
		validMoves = append(validMoves, loc)
	} else {
		return validMoves
	}

	// if piece hasn't moved, check 2 spaces ahead of this piece
	loc = location.Location{Row: currentRow + 2*movementDirection, Col: currentCol}
	if !p.hasMoved && isValidLocation(loc) && !isLocationOccupied(loc, pcs) {
		validMoves = append(validMoves, loc)
	}
	return validMoves
}

// Move sets the location of the pawn to a new location and sets hasMoved member to true
func (p *Pawn) Move(newLocation location.Location) {
	p.loc = newLocation
	p.hasMoved = true
}
