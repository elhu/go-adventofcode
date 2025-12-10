package main

import (
	"adventofcode/utils/files"
	"adventofcode/utils/sets"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	target  string
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
	machine.target = parts[0][1 : len(parts[0])-1]
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

type State struct {
	toPush       int
	currentState []byte
}

func buildNewState(current []byte, button []int) []byte {
	newState := make([]byte, len(current))
	copy(newState, current)
	for _, pos := range button {
		if newState[pos] == '.' {
			newState[pos] = '#'
		} else {
			newState[pos] = '.'
		}
	}
	return newState
}

func solveMachine(machine Machine) int {
	rounds := 0
	visited := sets.New[string]()
	visited.Add(strings.Repeat(".", len(machine.target)))
	queue := make([]State, len(machine.buttons))
	for i := range queue {
		queue[i] = State{toPush: i, currentState: []byte(strings.Repeat(".", len(machine.target)))}
	}
	var nextQueue []State
	var head State
	for len(queue) > 0 {
		head, queue = queue[0], queue[1:]
		button := machine.buttons[head.toPush]
		visited.Add(string(head.currentState))
		newState := buildNewState(head.currentState, button)
		stateStr := string(newState)
		if visited.HasMember(stateStr) {
			if len(queue) == 0 {
				rounds++
				queue = nextQueue
				nextQueue = []State{}
			}
			continue
		}
		if stateStr == machine.target {
			return rounds + 1
		}
		for i := range machine.buttons {
			if i != head.toPush {
				nextQueue = append(nextQueue, State{toPush: i, currentState: newState})
			}
		}
		if len(queue) == 0 {
			rounds++
			queue = nextQueue
			nextQueue = []State{}
		}
	}
	fmt.Println("Visited states:", visited.Len(), "in rounds:", rounds)
	panic("No solution found")
}

func solve(machines []Machine) int {
	res := 0
	for _, machine := range machines {
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
