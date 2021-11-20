package coords2d

// Coords2d represents coordinates in 2 dimensions
type Coords2d struct {
	X, Y int
}

// Add returns new Coords2d that are the sum of a and b
func Add(a, b Coords2d) Coords2d {
	return Coords2d{
		a.X + b.X,
		a.Y + b.Y,
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// Returns 2D Manhattan distance between a and b
func Distance(a, b Coords2d) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}
