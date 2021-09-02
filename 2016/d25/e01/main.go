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
	name string
	args []string
}

func atoi(str string) int {
	val, err := strconv.Atoi(str)
	check(err)
	return val
}

func process(registers map[string]int, instructions []*Instruction) bool {
	pos := 0
	prevEmitted := 1
	emittedCount := 0
	for pos < len(instructions) {
		inst := instructions[pos]
		switch inst.name {
		case "out":
			var value int
			if strings.ContainsAny(inst.args[0], "abcd") {
				value = registers[inst.args[0]]
			} else {
				value = atoi(inst.args[0])
			}
			if (value < 0 || value > 1) || value == prevEmitted {
				return false
			}
			prevEmitted = value
			emittedCount++
			if emittedCount > 50 {
				return true
			}
			pos++
		case "cpy":
			var value int
			if strings.ContainsAny(inst.args[0], "abcd") {
				value = registers[inst.args[0]]
			} else {
				value = atoi(inst.args[0])
			}
			registers[inst.args[1]] = value
			pos++
		case "inc":
			registers[inst.args[0]]++
			pos++
		case "dec":
			registers[inst.args[0]]--
			pos++
		case "jnz":
			var value int
			if strings.ContainsAny(inst.args[0], "abcd") {
				value = registers[inst.args[0]]
			} else {
				value = atoi(inst.args[0])
			}
			if value == 0 {
				pos++
			} else {
				var arg int
				if strings.ContainsAny(inst.args[1], "abcd") {
					arg = registers[inst.args[1]]
				} else {
					arg = atoi(inst.args[1])
				}
				pos += arg
			}
		}
	}
	return false
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")

	instructions := make([]*Instruction, len(input))
	for i, l := range input {
		parts := strings.Split(l, " ")
		instructions[i] = &Instruction{
			parts[0],
			parts[1:],
		}
	}
	for i := 0; ; i++ {
		var registers = map[string]int{
			"a": i,
			"b": 0,
			"c": 0,
			"d": 0,
		}
		if process(registers, instructions) {
			fmt.Println(i)
			return
		}
	}
}
