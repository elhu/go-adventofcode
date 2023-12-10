package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func pad(lines []string, char string) []string {
	padded := make([]string, len(lines)+2)
	padded[0] = strings.Repeat(char, len(lines[0])+2)
	for i, line := range lines {
		padded[i+1] = char + line + char
	}
	padded[len(lines)+1] = strings.Repeat(char, len(lines[0])+2)
	return padded
}

func findStartPos(lines []string) coords2d.Coords2d {
	for i, l := range lines {
		if idx := strings.Index(l, "S"); idx != -1 {
			return coords2d.Coords2d{Y: i, X: idx}
		}
	}
	panic("No starting point")
}

var north = coords2d.Coords2d{X: 0, Y: -1}
var south = coords2d.Coords2d{X: 0, Y: 1}
var west = coords2d.Coords2d{X: -1, Y: 0}
var east = coords2d.Coords2d{X: 1, Y: 0}
var selfVec = coords2d.Coords2d{X: 0, Y: 0}

var vectors = map[byte][2]coords2d.Coords2d{
	'|': {north, south},
	'-': {east, west},
	'L': {north, east},
	'J': {north, west},
	'7': {south, west},
	'F': {south, east},
}

func equals(a, b coords2d.Coords2d) bool {
	return a.X == b.X && a.Y == b.Y
}

func dist(a, b coords2d.Coords2d) coords2d.Coords2d {
	return coords2d.Coords2d{X: a.X - b.X, Y: a.Y - b.Y}
}

func move(pos, prevPos coords2d.Coords2d, lines []string) coords2d.Coords2d {
	diff := dist(prevPos, pos)
	if vec, found := vectors[lines[pos.Y][pos.X]]; found {
		if equals(vec[0], diff) {
			return coords2d.Add(pos, vec[1])
		} else if equals(vec[1], diff) {
			return coords2d.Add(pos, vec[0])
		}
	}
	return coords2d.Coords2d{X: -1, Y: -1}
}

func findLoop(lines []string) (byte, []coords2d.Coords2d) {
	sp := findStartPos(lines)
	var stuck = coords2d.Coords2d{X: -1, Y: -1}
	for pipe := range vectors {
		left, right := coords2d.Add(sp, vectors[pipe][0]), coords2d.Add(sp, vectors[pipe][1])
		var visited []coords2d.Coords2d = []coords2d.Coords2d{sp, left, right}
		prevLeft, prevRight := sp, sp
		for !equals(left, right) {
			newLeft := move(left, prevLeft, lines)
			newRight := move(right, prevRight, lines)
			if equals(newLeft, stuck) || equals(newRight, stuck) {
				break
			}
			prevLeft, prevRight = left, right
			left, right = newLeft, newRight
			visited = append(visited, left, right)
		}
		if equals(left, right) && !equals(left, stuck) {
			return pipe, visited
		}
	}
	panic("Can't find main loop")
}

func buildLoopMap(startPipe byte, loop []coords2d.Coords2d, lines []string) [][]byte {
	lm := make([][]byte, len(lines))
	for i, l := range lines {
		lm[i] = bytes.Repeat([]byte("."), len(l))
	}
	for _, lc := range loop {
		lm[lc.Y][lc.X] = lines[lc.Y][lc.X]
	}
	lm[loop[0].Y][loop[0].X] = startPipe
	return lm
}

var neighbors = []coords2d.Coords2d{
	{X: 0, Y: -1},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
	{X: 1, Y: 0},
	{X: 1, Y: 1},
	{X: -1, Y: 1},
	{X: 1, Y: -1},
	{X: -1, Y: -1},
}

func fill(lm [][]byte, c coords2d.Coords2d, fillVal byte) (bool, []coords2d.Coords2d) {
	hb := false
	if c.Y < 0 || c.Y >= len(lm) || c.X < 0 || c.X >= len(lm[c.Y]) {
		return true, []coords2d.Coords2d{}
	}
	if lm[c.Y][c.X] != '.' {
		return false, []coords2d.Coords2d{}
	}
	filled := []coords2d.Coords2d{c}
	lm[c.Y][c.X] = fillVal
	for _, n := range neighbors {
		b, f := fill(lm, coords2d.Add(c, n), fillVal)
		hb = hb || b
		filled = append(filled, f...)
	}
	return hb, filled
}

func expandMap(loop []coords2d.Coords2d, lm [][]byte) [][]byte {
	em := make([][]byte, len(lm)*3)
	for i := range em {
		em[i] = bytes.Repeat([]byte("."), len(lm[0])*3)
	}
	for _, l := range loop {
		char := lm[l.Y][l.X]
		vecs := vectors[char]
		for _, vec := range append(vecs[:], selfVec) {
			ec := coords2d.Coords2d{X: l.X * 3, Y: l.Y * 3}
			ec = coords2d.Add(ec, vec)
			em[ec.Y][ec.X] = '#'
		}
	}
	return em
}

func expandedContained(pos coords2d.Coords2d, lm [][]byte) bool {
	// Cells are now 3x3 and they must all be included within the loop
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if lm[pos.Y+i][pos.X+j] != 'I' {
				return false
			}
		}
	}
	return true
}

func solve(lines []string) int {
	startPipe, loop := findLoop(lines)
	lm := expandMap(loop, buildLoopMap(startPipe, loop, lines))
	for i := 0; i < len(lm); i++ {
		for j := 0; j < len(lm[i]); j++ {
			fv := byte('I')
			hitsBounds, filled := fill(lm, coords2d.Coords2d{X: j, Y: i}, '?')
			if hitsBounds {
				fv = 'O'
			}
			for _, f := range filled {
				lm[f.Y][f.X] = fv
			}
		}
	}
	res := 0
	for y := 0; y < len(lm); y += 3 {
		for x := 0; x < len(lm[y]); x += 3 {
			if expandedContained(coords2d.Coords2d{Y: y, X: x}, lm) {
				res++
			}
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	lines = pad(lines, ".")
	fmt.Println(solve(lines))
}
