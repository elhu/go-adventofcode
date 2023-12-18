package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Instruction struct {
	direction byte
	distance  int
	color     string
}

func parseInstructions(line string) Instruction {
	var inst Instruction
	fmt.Sscanf(line, "%c %d %s", &inst.direction, &inst.distance, &inst.color)
	return inst
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var vecs = map[byte]coords2d.Coords2d{
	'U': {X: 0, Y: -1},
	'D': {X: 0, Y: 1},
	'L': {X: -1, Y: 0},
	'R': {X: 1, Y: 0},
}

func fill(grid [][]byte, start coords2d.Coords2d) {
	queue := []coords2d.Coords2d{start}
	var head coords2d.Coords2d
	for len(queue) > 0 {
		head, queue = queue[0], queue[1:]
		if grid[head.Y][head.X] == '#' {
			continue
		}
		grid[head.Y][head.X] = '#'
		for _, v := range vecs {
			next := coords2d.Add(head, v)
			if grid[next.Y][next.X] == '.' {
				queue = append(queue, next)
			}
		}
	}
}

func solve(instructions []Instruction) int {
	var vertices []coords2d.Coords2d

	pos := coords2d.Coords2d{X: 0, Y: 0}
	vertices = append(vertices, pos)

	for _, inst := range instructions {
		switch inst.direction {
		case 'U':
			pos.Y -= inst.distance
		case 'D':
			pos.Y += inst.distance
		case 'L':
			pos.X -= inst.distance
		case 'R':
			pos.X += inst.distance
		}
		vertices = append(vertices, pos)
	}
	xMin, xMax := vertices[0].X, 0
	yMin, yMax := vertices[0].Y, 0
	for _, v := range vertices {
		xMin, xMax = min(xMin, v.X), max(xMax, v.X)
		yMin, yMax = min(yMin, v.Y), max(yMax, v.Y)
	}
	grid := make([][]byte, yMax-yMin+1)
	for i := range grid {
		grid[i] = bytes.Repeat([]byte("."), xMax-xMin+1)
	}
	pos = coords2d.Coords2d{X: 0, Y: 0}
	for _, inst := range instructions {
		for i := 0; i < inst.distance; i++ {
			pos = coords2d.Add(pos, vecs[inst.direction])
			grid[pos.Y-yMin][pos.X-xMin] = '#'
		}
	}
	fill(grid, coords2d.Coords2d{X: vertices[0].X - xMin + 1, Y: vertices[0].Y - yMin + 1})
	res := 0
	for _, l := range grid {
		res += bytes.Count(l, []byte("#"))
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	var instructions []Instruction
	for _, line := range strings.Split(data, "\n") {
		instructions = append(instructions, parseInstructions(line))
	}
	fmt.Println(solve(instructions))
}
