package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	name   string
	cycles int
	fn     func(cpu *CPU, args []string)
}

type InstArgs struct {
	inst *Instruction
	args []string
}

type CPU struct {
	currentCycle int
	x            int
}

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

var instructionsSet = map[string]*Instruction{
	"addx": &Instruction{
		"addx", 2, func(cpu *CPU, args []string) { cpu.x += atoi(args[0]) },
	},
	"noop": &Instruction{
		"noop", 1, func(cpu *CPU, args []string) {},
	},
}

func solve(cpu *CPU, instructions []InstArgs) []byte {
	pixels := make([]byte, 6*40)
	for i := 0; i < len(pixels); i++ {
		pixels[i] = '.'
	}
	for _, inst := range instructions {
		if cpu.currentCycle%40 >= cpu.x-1 && cpu.currentCycle%40 <= cpu.x+1 {
			pixels[cpu.currentCycle] = '#'
		}
		if inst.inst.cycles == 2 && (cpu.currentCycle+1)%40 >= cpu.x-1 && (cpu.currentCycle+1)%40 <= cpu.x+1 {
			pixels[cpu.currentCycle+1] = '#'
		}
		inst.inst.fn(cpu, inst.args)
		cpu.currentCycle += inst.inst.cycles
	}
	return pixels
}

func main() {
	data := files.ReadLines(os.Args[1])
	cpu := &CPU{currentCycle: 0, x: 1}
	instructions := make([]InstArgs, len(data))
	for i, line := range data {
		parts := strings.Split(line, " ")
		instructions[i] = InstArgs{instructionsSet[parts[0]], parts[1:]}
	}
	pixels := solve(cpu, instructions)
	for i := 0; i < len(pixels); i += 40 {
		fmt.Println(string(pixels[i : i+40]))
	}
}
