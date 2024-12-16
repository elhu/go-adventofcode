package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

var DIRS = map[byte]coords2d.Coords2d{'>': {1, 0}, '<': {-1, 0}, '^': {0, -1}, 'v': {0, 1}}

const BOX = 'O'
const EMPTY = '.'
const WALL = '#'

func push(pos coords2d.Coords2d, dir coords2d.Coords2d, grid map[coords2d.Coords2d]byte) bool {
	// Check if the box can be pushed
	pushable := false
	var empty coords2d.Coords2d
	for p := pos; ; p = coords2d.Add(p, dir) {
		c, ok := grid[p]
		if !ok || c == WALL {
			break
		}
		if c == EMPTY {
			pushable = true
			empty = p
			break
		}
	}
	if !pushable {
		return false
	}
	// Push the box
	grid[empty] = BOX
	grid[pos] = EMPTY
	return true
}

func solve(pos coords2d.Coords2d, grid map[coords2d.Coords2d]byte, moves string) int {
	for _, move := range moves {
		p := coords2d.Add(pos, DIRS[byte(move)])
		if grid[p] == BOX {
			if !push(p, DIRS[byte(move)], grid) {
				p = pos
			}
		} else if grid[p] == WALL {
			p = pos
		}
		pos = p
	}
	res := 0
	for pos, c := range grid {
		if c == BOX {
			res += pos.X + pos.Y*100
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	parts := strings.Split(data, "\n\n")
	grid := make(map[coords2d.Coords2d]byte)
	var start coords2d.Coords2d
	for i, line := range strings.Split(parts[0], "\n") {
		for j, c := range line {
			if c == '@' {
				start = coords2d.Coords2d{X: j, Y: i}
				c = EMPTY
			}
			grid[coords2d.Coords2d{X: j, Y: i}] = byte(c)
		}
	}
	moves := strings.Join(strings.Split(parts[1], "\n"), "")
	fmt.Println(solve(start, grid, moves))
}
