package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"adventofcode/utils/grids"
	"adventofcode/utils/sets"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const WIDTH, HEIGHT = 71, 71

func atoi(s string) int {
	n, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return n
}

var dirs = grids.ALL_DIRS

func findPath(grid map[coords2d.Coords2d]byte, start, end coords2d.Coords2d) int {
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
	return -1
}

func solve(grid grids.Grid[byte], blocks []coords2d.Coords2d, start, end coords2d.Coords2d) string {
	for i := 1024; i < len(blocks); i++ {
		cpy := grid.Copy()
		for j := 0; j < i; j++ {
			cpy.Set(blocks[j], '#')
		}
		if findPath(cpy, start, end) == -1 {
			return fmt.Sprintf("%d,%d", blocks[i-1].X, blocks[i-1].Y)
		}
	}
	panic("WTF")
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	grid := grids.NewGrid[byte](WIDTH, HEIGHT)
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			grid.Set(coords2d.Coords2d{X: j, Y: i}, '.')
		}
	}
	blocks := []coords2d.Coords2d{}
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, y := atoi(parts[0]), atoi(parts[1])
		blocks = append(blocks, coords2d.Coords2d{X: x, Y: y})
	}
	fmt.Println(solve(grid, blocks, coords2d.Coords2d{X: 0, Y: 0}, coords2d.Coords2d{X: WIDTH - 1, Y: HEIGHT - 1}))
}
