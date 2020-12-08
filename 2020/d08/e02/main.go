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

type Instruction struct {
	name     string
	arg      int
	executed bool
}

func unsafeAtoi(str string) int {
	i, err := strconv.Atoi(str)
	check(err)
	return i
}

func run(instructions []Instruction) (int, bool) {
	acc := 0
	pos := 0
	for !instructions[pos].executed {
		instructions[pos].executed = true
		switch instructions[pos].name {
		case "acc":
			acc += instructions[pos].arg
			pos++
		case "jmp":
			pos += instructions[pos].arg
		case "nop":
			pos++
		case "mgc":
			return acc, true
		}
	}
	return acc, false
}

func solve(instructions []Instruction) int {
	for i, inst := range instructions {
		newInst := ""
		if inst.name == "nop" {
			newInst = "jmp"
		} else if inst.name == "jmp" {
			newInst = "nop"
		} else {
			continue
		}
		cpy := make([]Instruction, len(instructions))
		copy(cpy, instructions)
		cpy[i].name = newInst
		if value, success := run(cpy); success {
			return value
		}
	}
	return -42
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	var instructions []Instruction
	for _, l := range lines {
		parts := strings.Split(l, " ")
		instructions = append(instructions, Instruction{parts[0], unsafeAtoi(parts[1]), false})
	}
	// Enqueue magic instruction at the end to terminate the loop
	instructions = append(instructions, Instruction{"mgc", 42, false})
	fmt.Println(solve(instructions))
}
