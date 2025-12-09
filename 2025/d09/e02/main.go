package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func area(a, b coords2d.Coords2d) int {
	return (abs(b.X-a.X) + 1) * (abs(b.Y-a.Y) + 1)
}

func arrangeSegment(a, b coords2d.Coords2d) [2]coords2d.Coords2d {
	if a.X == b.X {
		if a.Y < b.Y {
			return [2]coords2d.Coords2d{a, b}
		}
		return [2]coords2d.Coords2d{b, a}
	}
	if a.X < b.X {
		return [2]coords2d.Coords2d{a, b}
	}
	return [2]coords2d.Coords2d{b, a}
}

func makeSides(a, b coords2d.Coords2d) [4][2]coords2d.Coords2d {
	return [4][2]coords2d.Coords2d{
		arrangeSegment(coords2d.Coords2d{X: a.X, Y: a.Y}, coords2d.Coords2d{X: a.X, Y: b.Y}), // Left side
		arrangeSegment(coords2d.Coords2d{X: b.X, Y: b.Y}, coords2d.Coords2d{X: b.X, Y: a.Y}), // Right side
		arrangeSegment(coords2d.Coords2d{X: a.X, Y: b.Y}, coords2d.Coords2d{X: b.X, Y: b.Y}), // Bottom side
		arrangeSegment(coords2d.Coords2d{X: b.X, Y: a.Y}, coords2d.Coords2d{X: a.X, Y: a.Y}), // Top side
	}
}

func rectangleCorners(a, b coords2d.Coords2d) [4]coords2d.Coords2d {
	return [4]coords2d.Coords2d{
		{X: a.X, Y: a.Y},
		{X: a.X, Y: b.Y},
		{X: b.X, Y: b.Y},
		{X: b.X, Y: a.Y},
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func onSegment(point, start, end coords2d.Coords2d) bool {
	return min(start.X, end.X) <= point.X && point.X <= max(start.X, end.X) &&
		min(start.Y, end.Y) <= point.Y && point.Y <= max(start.Y, end.Y)
}

func pointInPolygon(point coords2d.Coords2d, redTiles []coords2d.Coords2d) bool {
	inside := false
	for i := 0; i < len(redTiles); i++ {
		j := (i + 1) % len(redTiles)
		start, end := redTiles[i], redTiles[j]

		if onSegment(point, start, end) {
			return true
		}
		if (start.Y > point.Y) != (end.Y > point.Y) {
			xIntersect := ((end.X-start.X)*(point.Y-start.Y))/(end.Y-start.Y) + start.X
			if point.X < xIntersect {
				inside = !inside
			}
		}
	}
	return inside
}

func segmentsIntersect(a, b, c, d coords2d.Coords2d) bool {
	// Check if segments are parallel
	if a.X == b.X && c.X == d.X {
		return false
	}
	if a.Y == b.Y && c.Y == d.Y {
		return false
	}

	// Check for general case
	return a.X > min(c.X, d.X) && a.X < max(c.X, d.X) &&
		a.Y < max(c.Y, d.Y) && b.Y > min(c.Y, d.Y)
}

func solve(redTiles []coords2d.Coords2d) int {
	maxSize := 0
	for i := 0; i < len(redTiles)-1; i++ {
		for j := i + 1; j < len(redTiles); j++ {
			corners := rectangleCorners(redTiles[i], redTiles[j])
			allInside := true
			for _, corner := range corners {
				if !pointInPolygon(corner, redTiles) {
					allInside = false
					break
				}
			}
			noSideIntersect := true
			for _, side := range makeSides(redTiles[i], redTiles[j]) {
				for k := 0; k < len(redTiles); k++ {
					l := (k + 1) % len(redTiles)
					if segmentsIntersect(side[0], side[1], redTiles[k], redTiles[l]) {
						noSideIntersect = false
						break
					}
				}
			}
			if allInside && noSideIntersect {
				a := area(redTiles[i], redTiles[j])
				maxSize = max(maxSize, a)
			}
		}
	}
	return maxSize
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	var redTiles []coords2d.Coords2d
	for _, line := range lines {
		var tile coords2d.Coords2d
		fmt.Sscanf(line, "%d,%d", &tile.X, &tile.Y)
		redTiles = append(redTiles, tile)
	}
	fmt.Println(solve(redTiles))
}
