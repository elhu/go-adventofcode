package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	check(err)
	return i
}

type Instruction struct {
	action string
	value  int
}

var angles = map[int][]int{
	0:   []int{1, 0},
	90:  []int{0, 1},
	180: []int{-1, 0},
	270: []int{0, -1},
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func mod(a, b int) int {
	res := a % b
	if res >= 0 {
		return res
	}
	return res + b
}

func solve(instructions []Instruction) int {
	x, y := 0, 0
	direction := 0
	for _, inst := range instructions {
		switch inst.action {
		case "N":
			y -= inst.value
		case "S":
			y += inst.value
		case "W":
			x -= inst.value
		case "E":
			x += inst.value
		case "F":
			x += angles[direction][0] * inst.value
			y += angles[direction][1] * inst.value
		case "R":
			direction = mod(direction+inst.value, 360)
		case "L":
			direction = mod(direction-inst.value, 360)
		}
	}

	return abs(x) + abs(y)
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	instructions := make([]Instruction, len(lines))
	for i, l := range lines {
		instructions[i] = Instruction{l[0:1], atoi(l[1:])}
	}
	fmt.Println(solve(instructions))
}
