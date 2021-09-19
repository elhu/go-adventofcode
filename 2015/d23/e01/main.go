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

func atoi(str string) int {
	i, e := strconv.Atoi(str)
	check(e)
	return i
}

type Instruction struct {
	name string
	args []string
}

func parse(input []string) []Instruction {
	var insts []Instruction
	for _, l := range input {
		l = strings.Replace(l, ",", "", 1)
		fields := strings.Fields(l)
		insts = append(insts, Instruction{fields[0], fields[1:]})
	}
	return insts
}

func solve(insts []Instruction) uint {
	registers := map[string]uint{"a": 0, "b": 0}
	for i := 0; i < len(insts); {
		inst := insts[i]
		switch inst.name {
		case "hlf":
			registers[inst.args[0]] /= 2
			i++
		case "tpl":
			registers[inst.args[0]] *= 3
			i++
		case "inc":
			registers[inst.args[0]]++
			i++
		case "jmp":
			i += atoi(inst.args[0])
		case "jie":
			if registers[inst.args[0]]%2 == 0 {
				i += atoi(inst.args[1])
			} else {
				i++
			}
		case "jio":
			if registers[inst.args[0]] == 1 {
				i += atoi(inst.args[1])
			} else {
				i++
			}
		}
	}
	return registers["b"]
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	insts := parse(input)
	fmt.Println(solve(insts))
}
