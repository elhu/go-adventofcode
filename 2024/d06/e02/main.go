package main

import (
	"adventofcode/utils/coords/coords3d"
	"adventofcode/utils/files"
	"adventofcode/utils/sets"
	"bytes"
	"fmt"
	"os"
)

var directions = [4]coords3d.Coords3d{
	{0, -1, 0},
	{1, 0, 0},
	{0, 1, 0},
	{-1, 0, 0},
}

func findStart(area [][]byte) coords3d.Coords3d {
	for y, line := range area {
		if x := bytes.Index(line, []byte("^")); x != -1 {
			return coords3d.Coords3d{x, y, 0}
		}
	}
	panic("No start found")
}

func hasLoop(area [][]byte, curr coords3d.Coords3d) bool {
	visited := sets.New[coords3d.Coords3d]()
	dir := 0
	for curr.X >= 0 && curr.Y >= 0 && curr.X < len(area[0]) && curr.Y < len(area) {
		if visited.HasMember(curr) {
			return true
		}
		visited.Add(curr)
		newPos := coords3d.Add(curr, directions[dir])
		if newPos.X < 0 || newPos.Y < 0 || newPos.X >= len(area[0]) || newPos.Y >= len(area) {
			return false
		}
		if area[newPos.Y][newPos.X] == '#' {
			newPos = curr
			dir = (dir + 1) % len(directions)
			newPos.Z = dir
		}
		curr = newPos
	}
	return false
}

func makeCopy(area [][]byte) [][]byte {
	cpy := make([][]byte, len(area))
	for i, line := range area {
		cpy[i] = make([]byte, len(line))
		copy(cpy[i], line)
	}
	return cpy
}

func solve(area [][]byte) int {
	curr := findStart(area)
	res := 0
	for y := 0; y < len(area); y++ {
		for x := 0; x < len(area[0]); x++ {
			if area[y][x] != '#' {
				cpy := makeCopy(area)
				cpy[y][x] = '#'
				if hasLoop(cpy, curr) {
					res += 1
				}
			}
		}
	}
	return res
}

func main() {
	data := bytes.TrimRight(files.ReadFile(os.Args[1]), "\n")
	area := bytes.Split(data, []byte("\n"))
	fmt.Println(solve(area))
}
