package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

var DIRS = map[byte]coords2d.Coords2d{'>': {1, 0}, '<': {-1, 0}, '^': {0, -1}, 'v': {0, 1}}

const BOX_L = '['
const BOX_R = ']'
const BOX = 'O'
const EMPTY = '.'
const WALL = '#'

func opposite(dir coords2d.Coords2d) coords2d.Coords2d {
	return coords2d.Coords2d{X: -dir.X, Y: -dir.Y}
}

func pushable(pos coords2d.Coords2d, dir coords2d.Coords2d, grid map[coords2d.Coords2d]byte) bool {
	p := coords2d.Add(pos, dir)
	c, ok := grid[p]
	if !ok || c == WALL {
		return false
	} else if c == EMPTY {
		return true
	} else if c == BOX_L {
		return pushable(p, dir, grid) && pushable(coords2d.Add(p, coords2d.Coords2d{1, 0}), dir, grid)
	} else if c == BOX_R {
		return pushable(p, dir, grid) && pushable(coords2d.Add(p, coords2d.Coords2d{-1, 0}), dir, grid)
	} else {
		panic("WTF")
	}
}

func doPushH(pos coords2d.Coords2d, dir coords2d.Coords2d, grid map[coords2d.Coords2d]byte) {
	next := coords2d.Add(pos, dir)
	if grid[next] == BOX_L {
		doPushH(next, dir, grid)
		doPushH(coords2d.Add(next, coords2d.Coords2d{1, 0}), dir, grid)
	} else if grid[next] == BOX_R {
		doPushH(next, dir, grid)
		doPushH(coords2d.Add(next, coords2d.Coords2d{-1, 0}), dir, grid)
	}
	grid[next] = grid[pos]
	grid[pos] = EMPTY
}

func pushV(pos coords2d.Coords2d, dir coords2d.Coords2d, grid map[coords2d.Coords2d]byte) bool {
	// Check if both sides of the box can be pushed
	var l, r coords2d.Coords2d
	if grid[pos] == BOX_L {
		l = pos
		r = coords2d.Add(pos, coords2d.Coords2d{X: 1, Y: 0})
	} else {
		l = coords2d.Add(pos, coords2d.Coords2d{X: -1, Y: 0})
		r = pos
	}
	if !pushable(l, dir, grid) || !pushable(r, dir, grid) {
		return false
	}
	doPushH(l, dir, grid)
	doPushH(r, dir, grid)
	return true
}

func pushH(pos coords2d.Coords2d, dir coords2d.Coords2d, grid map[coords2d.Coords2d]byte) bool {
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
	op := opposite(dir)
	for p := empty; p != pos; p = coords2d.Add(p, op) {
		grid[p] = grid[coords2d.Add(p, op)]
	}
	grid[pos] = EMPTY
	return true
}

func solve(pos coords2d.Coords2d, grid map[coords2d.Coords2d]byte, moves string) int {
	for _, move := range moves {
		p := coords2d.Add(pos, DIRS[byte(move)])
		if grid[p] == BOX_L || grid[p] == BOX_R {
			push := pushV
			if move == '<' || move == '>' {
				push = pushH
			}
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
		if c == BOX_L {
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
				start = coords2d.Coords2d{X: j * 2, Y: i}
				c = EMPTY
				grid[coords2d.Coords2d{X: j * 2, Y: i}] = byte(c)
				grid[coords2d.Coords2d{X: j*2 + 1, Y: i}] = EMPTY
			} else if c == EMPTY || c == WALL {
				grid[coords2d.Coords2d{X: j * 2, Y: i}] = byte(c)
				grid[coords2d.Coords2d{X: j*2 + 1, Y: i}] = byte(c)
			} else if c == BOX {
				grid[coords2d.Coords2d{X: j * 2, Y: i}] = '['
				grid[coords2d.Coords2d{X: j*2 + 1, Y: i}] = ']'
			} else {
				panic("Invalid character")
			}
		}
	}
	moves := strings.Join(strings.Split(parts[1], "\n"), "")
	fmt.Println(solve(start, grid, moves))
}
