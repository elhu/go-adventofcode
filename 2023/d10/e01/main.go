package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
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

func solve(lines []string) int {
	sp := findStartPos(lines)
	var stuck = coords2d.Coords2d{X: -1, Y: -1}
	for pipe := range vectors {
		left, right := coords2d.Add(sp, vectors[pipe][0]), coords2d.Add(sp, vectors[pipe][1])
		prevLeft, prevRight := sp, sp
		res := 1
		for ; !equals(left, right); res++ {
			newLeft := move(left, prevLeft, lines)
			newRight := move(right, prevRight, lines)
			if equals(newLeft, stuck) || equals(newRight, stuck) {
				break
			}
			prevLeft, prevRight = left, right
			left, right = newLeft, newRight
		}
		if equals(left, right) && !equals(left, stuck) {
			return res
		}
	}
	return 0
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	lines = pad(lines, ".")
	fmt.Println(solve(lines))
}
