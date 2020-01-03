package board

import "chess/pieces"

type tileColor int32

const (
	BLACK tileColor = iota
	WHITE
)

type tile struct {
	color tileColor
	piece *pieces.Piece
}

func NewTile(tc tileColor) *tile {
	return &tile{
		color: tc,
		piece: nil,
	}
}

func (t *tile) HasPiece() bool {
	return t.piece == nil
}

func (t *tile) GetPiece() *pieces.Piece {
	return t.piece
}

func (t *tile) GetColor() tileColor {
	return t.color
}

func (t *tile) SetPiece(p *pieces.Piece) {
	t.piece = p
}
