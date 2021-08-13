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

func process(registers map[string]int, instructions []Instruction) {
	pos := 0
	for pos < len(instructions) {
		inst := instructions[pos]
		switch inst.name {
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
				pos += atoi(inst.args[1])
				// fmt.Printf("[%d] Moving pos by %d because %d != 0\n", pos, atoi(inst.args[1]), registers[inst.args[0]])
			}
		}
	}
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")

	instructions := make([]Instruction, len(input))
	for i, l := range input {
		parts := strings.Split(l, " ")
		instructions[i] = Instruction{
			parts[0],
			parts[1:],
		}
	}
	var registers = map[string]int{
		"a": 0,
		"b": 0,
		"c": 1,
		"d": 0,
	}
	process(registers, instructions)
	fmt.Println(registers["a"])
}
