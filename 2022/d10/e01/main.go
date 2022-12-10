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

func solve(cpu *CPU, instructions []InstArgs) int {
	samples := make([]int, 0, 6)
	sampleFreqs := []int{20, 60, 100, 140, 180, 220}
	for _, inst := range instructions {
		if len(sampleFreqs) == 0 {
			break
		}
		if cpu.currentCycle+inst.inst.cycles >= sampleFreqs[0] {
			samples = append(samples, cpu.x*sampleFreqs[0])
			sampleFreqs = sampleFreqs[1:]
		}
		inst.inst.fn(cpu, inst.args)
		cpu.currentCycle += inst.inst.cycles
	}
	res := samples[0]
	for _, s := range samples[1:] {
		res += s
	}
	return res
}

func main() {
	data := files.ReadLines(os.Args[1])
	cpu := &CPU{currentCycle: 0, x: 1}
	instructions := make([]InstArgs, len(data))
	for i, line := range data {
		parts := strings.Split(line, " ")
		instructions[i] = InstArgs{instructionsSet[parts[0]], parts[1:]}
	}
	fmt.Println(solve(cpu, instructions))
}
