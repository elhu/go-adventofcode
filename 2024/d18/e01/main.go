package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"adventofcode/utils/sets"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// const WIDTH, HEIGHT = 7, 7
// const SIMUL_STEPS = 12

const WIDTH, HEIGHT = 71, 71
const SIMUL_STEPS = 1024

func atoi(s string) int {
	n, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return n
}

var dirs = []coords2d.Coords2d{{X: 0, Y: -1}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 1, Y: 0}}

func solve(grid map[coords2d.Coords2d]byte, start, end coords2d.Coords2d) int {
	visited := sets.New[coords2d.Coords2d]()
	queue := []coords2d.Coords2d{start}
	nextQueue := []coords2d.Coords2d{}
	var curr coords2d.Coords2d
	dist := 0
	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		if curr == end {
			return dist
		}
		if !visited.HasMember(curr) {
			visited.Add(curr)
			for _, dir := range dirs {
				next := coords2d.Add(curr, dir)
				if _, ok := grid[next]; ok && grid[next] == '.' {
					nextQueue = append(nextQueue, next)
				}
			}
		}
		if len(queue) == 0 {
			dist++
			queue, nextQueue = nextQueue, []coords2d.Coords2d{}
		}
	}
	panic("no path found")
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	grid := make(map[coords2d.Coords2d]byte)
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			grid[coords2d.Coords2d{X: j, Y: i}] = '.'
		}
	}
	for i, line := range lines {
		if i >= SIMUL_STEPS {
			break
		}
		parts := strings.Split(line, ",")
		x, y := atoi(parts[0]), atoi(parts[1])
		grid[coords2d.Coords2d{X: x, Y: y}] = '#'
	}
	fmt.Println(solve(grid, coords2d.Coords2d{X: 0, Y: 0}, coords2d.Coords2d{X: WIDTH - 1, Y: HEIGHT - 1}))
}
