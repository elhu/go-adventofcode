package coords2d

// Coords3d represents coordinates in 3 dimensions
type Coords3d struct {
	X, Y, Z int
}

// Add returns new Coords3d that are the sum of a and b
func Add(a, b Coords3d) Coords3d {
	return Coords3d{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// Returns 3D Manhattan distance between a and b
func Distance(a, b Coords3d) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y) + abs(a.Z-b.Z)
}
