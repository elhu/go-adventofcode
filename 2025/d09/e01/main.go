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

func solve(redTiles []coords2d.Coords2d) int {
	maxArea := 0
	for i := 0; i < len(redTiles)-1; i++ {
		for j := i + 1; j < len(redTiles); j++ {
			a := area(redTiles[i], redTiles[j])
			if a > maxArea {
				maxArea = a
			}
		}
	}
	return maxArea
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
