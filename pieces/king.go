package pieces

import (
	"chess/board/location"
	"chess/game/setup"
)

// King represents a king chess piece
type King struct {
	loc      location.Location
	color    PieceColor
	hasMoved bool
}

// NewKing returns a pointer to a new king
func NewKing(l location.Location, c PieceColor) *King {
	return &King{
		loc:      l,
		color:    c,
		hasMoved: false,
	}
}

// Location returns the location of the king
func (k *King) Location() location.Location {
	return k.loc
}

// Color returns the color of the king
func (k *King) Color() PieceColor {
	return k.color
}

// HasMoved returns whether or not the king has moved
func (k *King) HasMoved() bool {
	return k.hasMoved
}

// Move moves the king to a new location and sets hasMoved to true
func (k *King) Move(newLocation location.Location) {
	k.loc = newLocation
	k.hasMoved = true
}

// ValidMoves returns all of the current possible moves the king can make
func (k *King) ValidMoves(pcs []Piece) []location.Location {
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

	currentRow := k.loc.GetRow()
	currentCol := k.loc.GetCol()

	var validMoves []location.Location
	var loc location.Location

	for _, b := range bearings {
		loc = location.Location{Row: currentRow + b.Row, Col: currentCol + b.Col}
		if isValidLocation(loc) {
			// location must either be vacant or occupied by opponent
			if !isLocationOccupied(loc, pcs) || isLocationOccupiedByOpponent(loc, k.color, pcs) {
				// cannot move into check
				if !k.LocationInCheck(loc, pcs) {
					validMoves = append(validMoves, loc)
				}
			}
		}
	}

	// check if can standard castle
	if k.canCastle(false, pcs) {
		if k.color == BLACK {
			loc = location.Location{Row: setup.BlackFirstRank, Col: setup.StandardCastleColumn}
		} else {
			loc = location.Location{Row: setup.WhiteFirstRank, Col: setup.StandardCastleColumn}
		}
		validMoves = append(validMoves, loc)
	}
	// check if can queenside castle
	if k.canCastle(true, pcs) {
		if k.color == BLACK {
			loc = location.Location{Row: setup.BlackFirstRank, Col: setup.QueensideCastleColumn}
		} else {
			loc = location.Location{Row: setup.WhiteFirstRank, Col: setup.QueensideCastleColumn}
		}
		validMoves = append(validMoves, loc)
	}
	return validMoves
}

func (k *King) canCastle(queenside bool, pcs []Piece) bool {
	// if king has moved, cannot castle
	if k.hasMoved {
		return false
	}

	// initialize expected rook location
	var rookStartingLoc location.Location
	if k.color == BLACK {
		rookStartingLoc.Row = setup.BlackFirstRank
	} else {
		rookStartingLoc.Row = setup.WhiteFirstRank
	}

	if queenside {
		rookStartingLoc.Col = 0
	} else {
		rookStartingLoc.Col = BOARD_SIZE - 1
	}

	pieceFound := false
	// check if rook has moved
	for _, p := range pcs {
		if p.Location().Equals(rookStartingLoc) {
			pieceFound = true
			// if the piece at the starting location has moved, then cannot castle
			if p.HasMoved() {
				return false
			}
		}
	}

	// if there is no piece where the rook should be, cannot castle
	if !pieceFound {
		return false
	}

	locationsToCheck := k.getLocationsToCheckForCastle(queenside)
	return k.checkLocationsForCastle(rookStartingLoc, locationsToCheck, pcs)
}

// LocationInCheck returns whether or not a given location would be in check if the king were placed there
func (k *King) LocationInCheck(loc location.Location, pcs []Piece) bool {
	var pieceValidMoves []location.Location
	for _, p := range pcs {
		if p.Color() != k.color {
			pieceValidMoves = p.ValidMoves(pcs)
			// check if location is in the valid moves of the opponent piece
			for _, l := range pieceValidMoves {
				if loc.Equals(l) {
					return true
				}
			}
		}
	}
	return false
}

// InCheck returns whether or not the king is currently in check
func (k *King) InCheck(pcs []Piece) bool {
	var pieceValidMoves []location.Location
	for _, p := range pcs {
		if p.Color() != k.color {
			pieceValidMoves = p.ValidMoves(pcs)
			// check if location is in the valid moves of the opponent piece
			for _, l := range pieceValidMoves {
				if k.loc.Equals(l) {
					return true
				}
			}
		}
	}
	return false
}

func (k *King) getLocationsToCheckForCastle(queenside bool) []location.Location {
	var row, col int
	if k.color == BLACK {
		row = setup.BlackFirstRank
	} else {
		row = setup.WhiteFirstRank
	}
	if queenside {
		col = 0
	} else {
		col = BOARD_SIZE - 1
	}

	var locations []location.Location
	var loc location.Location
	isAtThisLocation := false

	for !isAtThisLocation {
		loc = location.Location{Row: row, Col: col}
		locations = append(locations, loc)
		if k.Location().Equals(loc) {
			isAtThisLocation = true
		}
		if queenside {
			col++
		} else {
			col--
		}
	}
	return locations
}

func (k *King) checkLocationsForCastle(rookLocation location.Location, locations []location.Location, pcs []Piece) bool {
	// check if any pieces are occupying the locations that must be empty in order to castle
	for _, l := range locations {
		for _, p := range pcs {
			// if a piece other than the expected pieces (rook,king) is in the location, cannot castle
			if p.Location().Equals(l) && !(p.Location().Equals(rookLocation)) && !(p.Location().Equals(k.loc)) {
				return false
			}
		}
		if !(l.Equals(rookLocation)) {
			// if location is in check, cannot castle since king cannot castle through check
			if k.LocationInCheck(l, pcs) {
				return false
			}
		}
	}
	return true
}
