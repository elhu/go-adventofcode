package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/draffensperger/golp"
)

type Machine struct {
	buttons [][]int
	joltage []int
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func parseMachine(line string) Machine {
	var machine Machine
	parts := strings.Split(line, " ")
	for i := 1; i < len(parts)-1; i++ {
		var button []int
		part := parts[i][1 : len(parts[i])-1]
		positions := strings.Split(part, ",")
		for _, pos := range positions {
			button = append(button, atoi(pos))
		}
		machine.buttons = append(machine.buttons, button)
	}
	joltageParts := strings.Split(parts[len(parts)-1][1:len(parts[len(parts)-1])-1], ",")
	for _, jolt := range joltageParts {
		machine.joltage = append(machine.joltage, atoi(jolt))
	}
	return machine
}

func solveMachine(machine Machine) int {
	btnCount := len(machine.buttons)
	joltageCount := len(machine.joltage)

	lp := golp.NewLP(0, btnCount)
	lp.SetVerboseLevel(golp.NEUTRAL)

	coefs := make([]float64, btnCount)
	for i := 0; i < btnCount; i++ {
		coefs[i] = 1.0
	}
	lp.SetObjFn(coefs)

	for i := 0; i < btnCount; i++ {
		lp.SetInt(i, true)
		lp.SetBounds(i, 0.0, float64(1000)) // Arbitrary upper bound
	}

	for i := 0; i < joltageCount; i++ {
		var entries []golp.Entry
		for j, btn := range machine.buttons {
			if slices.Contains(btn, i) {
				entries = append(entries, golp.Entry{Col: j, Val: 1.0})
			}
		}
		targetValue := float64(machine.joltage[i])
		if err := lp.AddConstraintSparse(entries, golp.EQ, targetValue); err != nil {
			panic(err)
		}
	}
	status := lp.Solve()

	if status != golp.OPTIMAL {
		return 0
	}
	solution := lp.Variables()
	totalPushes := 0
	for _, val := range solution {
		totalPushes += int(val + 0.5)
	}
	return totalPushes
}

func solve(machines []Machine) int {
	res := 0
	for i, machine := range machines {
		fmt.Printf("Solving machine %d/%d with target %v\n", i+1, len(machines), machine.joltage)
		r := solveMachine(machine)
		res += r
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	var machines []Machine
	for _, line := range lines {
		machines = append(machines, parseMachine(line))
	}
	fmt.Println(solve(machines))
}
