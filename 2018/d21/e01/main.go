package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var operations = map[string]func(inA, inB, outC int, registers []int){
	"addr": func(inA, inB, outC int, registers []int) { registers[outC] = registers[inA] + registers[inB] },
	"addi": func(inA, inB, outC int, registers []int) { registers[outC] = registers[inA] + inB },
	"mulr": func(inA, inB, outC int, registers []int) { registers[outC] = registers[inA] * registers[inB] },
	"muli": func(inA, inB, outC int, registers []int) { registers[outC] = registers[inA] * inB },
	"banr": func(inA, inB, outC int, registers []int) { registers[outC] = registers[inA] & registers[inB] },
	"bani": func(inA, inB, outC int, registers []int) { registers[outC] = registers[inA] & inB },
	"borr": func(inA, inB, outC int, registers []int) { registers[outC] = registers[inA] | registers[inB] },
	"bori": func(inA, inB, outC int, registers []int) { registers[outC] = registers[inA] | inB },
	"setr": func(inA, inB, outC int, registers []int) { registers[outC] = registers[inA] },
	"seti": func(inA, inB, outC int, registers []int) { registers[outC] = inA },
	"gtir": func(inA, inB, outC int, registers []int) {
		if inA > registers[inB] {
			registers[outC] = 1
		} else {
			registers[outC] = 0
		}
	},
	"gtri": func(inA, inB, outC int, registers []int) {
		if registers[inA] > inB {
			registers[outC] = 1
		} else {
			registers[outC] = 0
		}
	},
	"gtrr": func(inA, inB, outC int, registers []int) {
		if registers[inA] > registers[inB] {
			registers[outC] = 1
		} else {
			registers[outC] = 0
		}
	},
	"eqir": func(inA, inB, outC int, registers []int) {
		if inA == registers[inB] {
			registers[outC] = 1
		} else {
			registers[outC] = 0
		}
	},
	"eqri": func(inA, inB, outC int, registers []int) {
		if registers[inA] == inB {
			registers[outC] = 1
		} else {
			registers[outC] = 0
		}
	},
	"eqrr": func(inA, inB, outC int, registers []int) {
		if registers[inA] == registers[inB] {
			registers[outC] = 1
		} else {
			registers[outC] = 0
		}
	},
}

var instExp = regexp.MustCompile(`(\w+) (\d+) (\d+) (\d+)`)

type Instruction struct {
	a, b, c int
	opCode  string
}

func parseInstruction(str string) Instruction {
	match := instExp.FindStringSubmatch(str)
	opCode := match[1]
	a, _ := strconv.Atoi(match[2])
	b, _ := strconv.Atoi(match[3])
	c, _ := strconv.Atoi(match[4])
	return Instruction{
		opCode: opCode,
		a:      a,
		b:      b,
		c:      c,
	}
}

func process(ipIndex int, insts []Instruction) []int {
	ip := 0
	registers := make([]int, 6)
	registers[0] = 7224964
	executed := 0
	for {
		registers[ipIndex] = ip
		if ip >= len(insts) {
			break
		}
		inst := insts[ip]
		// fmt.Printf("ip=%d %v %v", ip, registers, inst)
		operations[inst.opCode](inst.a, inst.b, inst.c, registers)
		// fmt.Printf(" %v\n", registers)
		ip = registers[ipIndex] + 1
		executed++
	}
	return registers
}

func parse(lines [][]byte) (int, []Instruction) {
	ipIndex, _ := strconv.Atoi(string(lines[0][4:5]))
	lines = lines[1:]
	instructions := make([]Instruction, 0, len(lines))
	for _, line := range lines {
		instructions = append(instructions, parseInstruction(string(line)))
	}
	return ipIndex, instructions
}

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input = bytes.TrimSuffix(input, []byte{'\n'})
	lines := bytes.Split(input, []byte{'\n'})
	ipIndex, instructions := parse(lines)
	registers := process(ipIndex, instructions)
	fmt.Println(registers)
}
