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

func run(instructions []Instruction) int {
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
		}
	}
	return acc
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
	fmt.Println(run(instructions))
}
