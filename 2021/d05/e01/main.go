package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"fmt"
	"os"
)

func toKey(c coords2d.Coords2d) int {
	// Consider that a 1M offset is enough to avoid collisions
	return c.Y*1000000 + c.X
}

func sort(a, b coords2d.Coords2d) (coords2d.Coords2d, coords2d.Coords2d) {
	if a.X == b.X {
		if a.Y > b.Y {
			return b, a
		}
		return a, b
	} else {
		if a.X > b.X {
			return b, a
		}
		return a, b
	}
}

func solve(input []string) int {
	vents := make(map[int]int)
	for _, l := range input {
		var start, end coords2d.Coords2d
		fmt.Sscanf(l, "%d,%d -> %d,%d", &start.X, &start.Y, &end.X, &end.Y)
		if start.X != end.X && start.Y != end.Y {
			continue
		}
		start, end = sort(start, end)
		for i := start.X; i <= end.X; i++ {
			for j := start.Y; j <= end.Y; j++ {
				vents[toKey(coords2d.Coords2d{X: i, Y: j})]++
			}
		}
	}
	res := 0
	for _, v := range vents {
		if v > 1 {
			res++
		}
	}
	return res
}

func main() {
	lines := files.ReadLines(os.Args[1])
	fmt.Println(solve(lines))
}
