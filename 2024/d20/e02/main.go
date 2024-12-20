package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"adventofcode/utils/sets"
	"fmt"
	"os"
	"strings"
)

func copyMap(src map[coords2d.Coords2d]byte) map[coords2d.Coords2d]byte {
	dst := make(map[coords2d.Coords2d]byte)
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

var DIRS = []coords2d.Coords2d{{X: 0, Y: 1}, {X: 0, Y: -1}, {X: 1, Y: 0}, {X: -1, Y: 0}}

func findPath(grid map[coords2d.Coords2d]byte, start, end coords2d.Coords2d) []coords2d.Coords2d {
	res := []coords2d.Coords2d{}
	queue := []coords2d.Coords2d{start}
	visited := sets.New[coords2d.Coords2d]()
	from := make(map[coords2d.Coords2d]coords2d.Coords2d)
	var curr coords2d.Coords2d
	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		if curr == end {
			break
		}
		if visited.HasMember(curr) {
			continue
		}
		visited.Add(curr)
		for _, dir := range DIRS {
			if next := coords2d.Add(curr, dir); grid[next] == '.' && !visited.HasMember(next) {
				from[next] = curr
				queue = append(queue, next)
			}
		}
	}
	if _, ok := from[end]; !ok {
		panic("WTF")
	}
	for curr := end; curr != start; curr = from[curr] {
		res = append(res, curr)
	}
	res = append(res, start)
	return res
}

const MIN_SAVE = 100
const MAX_CHEAT = 20

type FT struct {
	From, To int
}

func solve(grid map[coords2d.Coords2d]byte, start, end coords2d.Coords2d) int {
	path := findPath(grid, start, end)
	res := 0
	for i := 0; i < len(path); i++ {
		for j := i + MIN_SAVE; j < len(path); j++ {
			dist := coords2d.Distance(path[i], path[j])
			if dist <= MAX_CHEAT && j-i-dist >= MIN_SAVE {
				res++
			}
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	grid := make(map[coords2d.Coords2d]byte)

	var start, end coords2d.Coords2d
	for y, line := range strings.Split(data, "\n") {
		for x, cell := range line {
			if cell == 'E' {
				end = coords2d.Coords2d{X: x, Y: y}
				cell = '.'
			}
			if cell == 'S' {
				start = coords2d.Coords2d{X: x, Y: y}
				cell = '.'
			}
			grid[coords2d.Coords2d{X: x, Y: y}] = byte(cell)
		}
	}
	fmt.Println(solve(grid, start, end))
}
