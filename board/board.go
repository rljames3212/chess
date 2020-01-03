package board

import "chess/board/location"

const (
	DEFAULT_BOARD_SIZE int = 8
)

// Board represents a chess board
type Board [][]*tile

// NewBoard returns a new chess boards
func NewBoard() *Board {
	// create and initialize board tiles
	board := make(Board, DEFAULT_BOARD_SIZE)
	for row := 0; row < DEFAULT_BOARD_SIZE; row++ {
		board[row] = make([]*tile, DEFAULT_BOARD_SIZE)
	}

	isWhiteTile := false
	for row := 0; row < DEFAULT_BOARD_SIZE; row++ {
		for col := 0; col < DEFAULT_BOARD_SIZE; col++ {
			if isWhiteTile {
				board[row][col] = NewTile(WHITE)
			} else {
				board[row][col] = NewTile(BLACK)
			}
			isWhiteTile = !isWhiteTile
		}
	}
	return &board
}

// GetTile returns a tile at a given location
func (b *Board) GetTile(loc location.Location) *tile {
	return (*b)[loc.GetRow()][loc.GetCol()]
}

// IsPieceAtLocation returns true if there is a piece at a given location and false otherwise
func (b *Board) IsPieceAtLocation(loc location.Location) bool {
	return (*b)[loc.GetRow()][loc.GetCol()].HasPiece()
}
