package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	direction byte
	distance  int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var dirMap = map[byte]byte{
	'0': 'R',
	'1': 'D',
	'2': 'L',
	'3': 'U',
}

func parseInstructions(line string) Instruction {
	var inst Instruction
	parts := strings.Fields(line)
	code := strings.Trim(parts[len(parts)-1], "#()")
	val, err := strconv.ParseInt(code[0:5], 16, 32)
	check(err)

	inst.distance = int(val)
	inst.direction = dirMap[code[5]]
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

func shoelace(vertices []coords2d.Coords2d) int {
	res := 0
	for i := 0; i < len(vertices)-1; i++ {
		res += vertices[i].X*vertices[i+1].Y - vertices[i+1].X*vertices[i].Y
	}
	return res / 2
}

func pick(vertices []coords2d.Coords2d, totalDistance int) int {
	return shoelace(vertices) + totalDistance/2 + 1
}

func solve(instructions []Instruction) int {
	var vertices []coords2d.Coords2d

	pos := coords2d.Coords2d{X: 0, Y: 0}
	vertices = append(vertices, pos)

	totalDistance := 0

	for _, inst := range instructions {
		totalDistance += inst.distance
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
	return pick(vertices, totalDistance)
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	var instructions []Instruction
	for _, line := range strings.Split(data, "\n") {
		instructions = append(instructions, parseInstructions(line))
	}
	fmt.Println(solve(instructions))
}
