package location

// Location represents a location on a chess board
type Location struct {
	Row int
	Col int
}

// GetRow returns the row of the location
func (l Location) GetRow() int {
	return l.Row
}

// GetCol returns the column of the location
func (l Location) GetCol() int {
	return l.Col
}

func (l Location) Equals(other Location) bool {
	return l.Row == other.Row && l.Col == other.Col
}
