package location

// Location represents a location on a chess board
type Location struct {
	row int
	col int
}

// GetRow returns the row of the location
func (l *Location) GetRow() int {
	return l.row
}

// GetCol returns the column of the location
func (l *Location) GetCol() int {
	return l.col
}
