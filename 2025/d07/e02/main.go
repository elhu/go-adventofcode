package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

func solve(start coords2d.Coords2d, lines []string) int {
	cache := make(map[coords2d.Coords2d]int)
	return visitPath(lines, start, cache)
}

func visitPath(lines []string, loc coords2d.Coords2d, cache map[coords2d.Coords2d]int) int {
	if cached, found := cache[loc]; found {
		return cached
	}
	if loc.Y+1 == len(lines) {
		return 1
	}
	if lines[loc.Y+1][loc.X] == '^' {
		leftPaths := visitPath(lines, coords2d.Coords2d{X: loc.X - 1, Y: loc.Y + 1}, cache)
		rightPaths := visitPath(lines, coords2d.Coords2d{X: loc.X + 1, Y: loc.Y + 1}, cache)
		cache[loc] = leftPaths + rightPaths
		return leftPaths + rightPaths
	} else {
		downPaths := visitPath(lines, coords2d.Coords2d{X: loc.X, Y: loc.Y + 1}, cache)
		cache[loc] = downPaths
		return downPaths
	}
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	sLoc := coords2d.Coords2d{X: strings.Index(lines[0], "S"), Y: 0}
	fmt.Println(solve(sLoc, lines))
}
