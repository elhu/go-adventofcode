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

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func solve(instructions []Instruction) int {
	x, y := 0, 0
	var vector = []int{10, -1}
	for _, inst := range instructions {
		switch inst.action {
		case "N":
			vector[1] -= inst.value
		case "S":
			vector[1] += inst.value
		case "W":
			vector[0] -= inst.value
		case "E":
			vector[0] += inst.value
		case "F":
			x += vector[0] * inst.value
			y += vector[1] * inst.value
		case "R":
			for i := 0; i < inst.value/90; i++ {
				vector[0], vector[1] = -vector[1], vector[0]
			}
		case "L":
			for i := 0; i < inst.value/90; i++ {
				vector[0], vector[1] = vector[1], -vector[0]
			}
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
