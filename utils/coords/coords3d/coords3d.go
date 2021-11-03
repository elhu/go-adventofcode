package coords2d

type Coords3d struct {
	X, Y, Z int
}

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

func Distance(a, b Coords3d) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y) + abs(a.Z-b.Z)
}
