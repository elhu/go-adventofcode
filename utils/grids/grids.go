package grids

import "adventofcode/utils/coords/coords2d"

type Grid[K any] struct {
	Width, Height int
	Points        map[coords2d.Coords2d]K
}

func NewGrid[K any](width, height int) *Grid[K] {
	return &Grid[K]{Width: width, Height: height, Points: make(map[coords2d.Coords2d]K)}
}

func NewGridFromStringArray(lines []string) *Grid[byte] {
	g := NewGrid[byte](len(lines[0]), len(lines))
	for y, line := range lines {
		for x, c := range line {
			g.Points[coords2d.Coords2d{X: x, Y: y}] = byte(c)
		}
	}
	return g
}

func (g *Grid[K]) Get(c coords2d.Coords2d) K {
	return g.Points[c]
}

func (g *Grid[K]) Set(c coords2d.Coords2d, v K) {
	g.Points[c] = v
}

func (g *Grid[K]) Copy() *Grid[K] {
	newGrid := NewGrid[K](g.Width, g.Height)
	for k, v := range g.Points {
		newGrid.Points[k] = v
	}
	return newGrid
}

func (g *Grid[K]) FindPath(from, to coords2d.Coords2d) []coords2d.Coords2d {
	visited := make(map[coords2d.Coords2d]bool)
	queue := []coords2d.Coords2d{from}
	nextQueue := []coords2d.Coords2d{}
	parents := make(map[coords2d.Coords2d]coords2d.Coords2d)
	var curr coords2d.Coords2d
	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		if curr == to {
			break
		}
		if !visited[curr] {
			visited[curr] = true
			for _, dir := range ALL_DIRS {
				next := coords2d.Add(curr, dir)
				if g.Get(next) == '.' {
					nextQueue = append(nextQueue, next)
					parents[next] = curr
				}
			}
		}
		if len(queue) == 0 {
			queue, nextQueue = nextQueue, []coords2d.Coords2d{}
		}
	}
	path := []coords2d.Coords2d{to}
	for curr := to; curr != from; curr = parents[curr] {
		path = append(path, parents[curr])
	}
	return path
}

var (
	EAST  = coords2d.Coords2d{X: 1, Y: 0}
	SOUTH = coords2d.Coords2d{X: 0, Y: 1}
	WEST  = coords2d.Coords2d{X: -1, Y: 0}
	NORTH = coords2d.Coords2d{X: 0, Y: -1}
)

var ALL_DIRS = []coords2d.Coords2d{EAST, SOUTH, WEST, NORTH}
