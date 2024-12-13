package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

type Machine struct {
	A, B, Prize coords2d.Coords2d
}

func parseMachine(data []string) Machine {
	var a, b, prize coords2d.Coords2d
	fmt.Sscanf(data[0], "Button A: X+%d, Y+%d", &a.X, &a.Y)
	fmt.Sscanf(data[1], "Button B: X+%d, Y+%d", &b.X, &b.Y)
	fmt.Sscanf(data[2], "Prize: X=%d, Y=%d", &prize.X, &prize.Y)
	return Machine{a, b, prize}
}

const MAX_PUSHES = 100
const A_COST = 3
const MAX_COST = MAX_PUSHES*A_COST + MAX_PUSHES

func play(m Machine) int {
	min := MAX_COST + 1
	for a := 0; a < MAX_PUSHES; a++ {
		dX := m.Prize.X - m.A.X*a
		dY := m.Prize.Y - m.A.Y*a
		if dX%m.B.X == 0 && dY%m.B.Y == 0 && dX/m.B.X == dY/m.B.Y {
			b := dX / m.B.X
			if a*A_COST+b < min {
				min = a*A_COST + b
			}
		}
	}
	if min == MAX_COST+1 {
		return 0
	}
	return min
}

func solve(machines []Machine) int {
	res := 0
	for _, m := range machines {
		res += play(m)
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	var machines []Machine
	for _, rm := range strings.Split(data, "\n\n") {
		m := parseMachine(strings.Split(rm, "\n"))
		machines = append(machines, m)
	}
	fmt.Println(solve(machines))
}
