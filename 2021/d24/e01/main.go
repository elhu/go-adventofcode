package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inp(registers map[string]int, args []string) {
	registers[args[0]] = nextInput()
}

func value(registers map[string]int, arg string) int {
	if n, err := strconv.Atoi(arg); err == nil {
		return n
	}
	return registers[arg]
}

func add(registers map[string]int, args []string) {
	registers[args[0]] += value(registers, args[1])
}

func mul(registers map[string]int, args []string) {
	registers[args[0]] *= value(registers, args[1])
}

func div(registers map[string]int, args []string) {
	registers[args[0]] /= value(registers, args[1])
}

func mod(registers map[string]int, args []string) {
	registers[args[0]] %= value(registers, args[1])
}

func eql(registers map[string]int, args []string) {
	if registers[args[0]] == value(registers, args[1]) {
		registers[args[0]] = 1
	} else {
		registers[args[0]] = 0
	}
}

var fns = map[string]func(map[string]int, []string){
	"inp": inp,
	"add": add,
	"mul": mul,
	"div": div,
	"mod": mod,
	"eql": eql,
}

type FuncArgs struct {
	fn   func(map[string]int, []string)
	args []string
}

func parseInstructions(data []string) []FuncArgs {
	var res []FuncArgs
	for _, line := range data {
		parts := strings.Split(line, " ")
		res = append(res, FuncArgs{fns[parts[0]], parts[1:]})
	}
	return res
}

var nextInput func() int

func initInput(number string) {
	currentPos := 0
	nextInput = func() int {
		n := int(number[currentPos] - '0')
		return n
	}
}

func processSerial(data []string, number string, instructions []FuncArgs) bool {
	initInput(number)
	var registers = map[string]int{"w": 0, "x": 0, "y": 0, "z": 0}
	for _, inst := range instructions {
		inst.fn(registers, inst.args)
	}
	return registers["z"] == 0
}

func solve(data []string, instructions []FuncArgs) int {
	for i := 99999999999999; i >= 11111111111111; i-- {
		stri := strconv.Itoa(i)
		if strings.Index(stri, "0") != -1 {
			continue
		}
		if processSerial(data, stri, instructions) {
			fmt.Printf("Found serial number: %s\n", stri)
			return i
		}
	}
	panic("WTF")
}

func solve2(data []string, instructions []FuncArgs) int {
	for i := 11111111111111; i <= 99999999999999; i++ {
		stri := strconv.Itoa(i)
		if strings.Index(stri, "0") != -1 {
			continue
		}
		if processSerial(data, stri, instructions) {
			fmt.Printf("Found serial number: %s\n", stri)
		}
	}
	panic("WTF")
}

func main() {
	data := files.ReadLines(os.Args[1])
	instructions := parseInstructions(data)
	fmt.Println(solve2(data, instructions))
}
