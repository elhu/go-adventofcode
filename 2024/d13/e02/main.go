package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	aocm "adventofcode/utils/math"
	"fmt"
	"math"
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
	prize.X += OFFSET
	prize.Y += OFFSET
	return Machine{a, b, prize}
}

const A_COST = 3
const MAX_COST = 99999999999999999
const OFFSET = 10000000000000

func play(m Machine) int {
	matrix := [][]float64{
		{float64(m.A.X), float64(m.B.X), float64(m.Prize.X)},
		{float64(m.A.Y), float64(m.B.Y), float64(m.Prize.Y)},
	}
	aocm.GaussianElimination(matrix)
	a := matrix[0][len(matrix[0])-1]
	b := matrix[1][len(matrix[1])-1]
	if math.Round(a)*float64(m.A.X)+math.Round(b)*float64(m.B.X) == float64(m.Prize.X) && math.Round(a)*float64(m.A.Y)+math.Round(b)*float64(m.B.Y) == float64(m.Prize.Y) {
		return int(math.Round(a*A_COST) + math.Round(b))
	}
	return 0
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
