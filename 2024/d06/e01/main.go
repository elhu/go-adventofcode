package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"adventofcode/utils/sets"
	"fmt"
	"os"
	"strings"
)

var directions = [4]coords2d.Coords2d{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func findStart(area []string) coords2d.Coords2d {
	for y, line := range area {
		if x := strings.Index(line, "^"); x != -1 {
			return coords2d.Coords2d{x, y}
		}
	}
	panic("No start found")
}

func solve(area []string) int {
	curr := findStart(area)
	dir := 0
	visited := sets.New[coords2d.Coords2d]()
	for curr.X >= 0 && curr.Y >= 0 && curr.X < len(area[0]) && curr.Y < len(area) {
		visited.Add(curr)
		newPos := coords2d.Add(curr, directions[dir])
		if newPos.X < 0 || newPos.Y < 0 || newPos.X >= len(area[0]) || newPos.Y >= len(area) {
			break
		}
		if area[newPos.Y][newPos.X] == '#' {
			newPos = curr
			dir = (dir + 1) % len(directions)
		}
		curr = newPos
	}
	return visited.Len()
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	area := strings.Split(data, "\n")
	fmt.Println(solve(area))
}
