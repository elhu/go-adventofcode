package coords2d

type Coords2d struct {
	X, Y int
}

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

func Distance(a, b Coords2d) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}
