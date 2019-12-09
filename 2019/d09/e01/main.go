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

func codeToArgType(code int) []int {
	res := make([]int, 0)
	// Remove code
	code /= 100
	for code > 0 {
		res = append(res, code%10)
		code /= 10
	}
	return res
}

func getArgs(opcodes []int, pos, numArgs int) []int {
	argTypes := codeToArgType(opcodes[pos])
	args := make([]int, 0, numArgs)
	fmt.Printf("Argtypes: %v\n", argTypes)
	for i := 0; i < numArgs; i++ {
		if len(argTypes) > i && argTypes[i] == 2 {
			args = append(args, opcodes[relativeBase+opcodes[pos+i+1]])
		} else if len(argTypes) > i && argTypes[i] == 1 {
			args = append(args, opcodes[pos+i+1])
		} else {
			// fmt.Printf("1. Trying to access offset %d\n", pos+i+1)
			// fmt.Printf("2. Trying to access offset %d\n", opcodes[pos+i+1])
			args = append(args, opcodes[opcodes[pos+i+1]])
		}
	}
	return args
}

func add(opcodes []int, pos int) int {
	args := getArgs(opcodes, pos, 3)
	fmt.Printf("[%d,%d,%d,%d] %d <=> %d\n", opcodes[pos], opcodes[pos+1], opcodes[pos+2], opcodes[pos+3], opcodes[pos+3], args[2])
	opcodes[args[2]] = args[0] + args[1]
	return pos + 4
}

func multiply(opcodes []int, pos int) int {
	args := getArgs(opcodes, pos, 2)
	opcodes[opcodes[pos+3]] = args[0] * args[1]
	return pos + 4
}

func save(opcodes []int, pos int) int {
	args := getArgs(opcodes, pos, 2)
	opcodes[args[0]] = input
	// fmt.Printf("Saving %d at pos %d\n", input, opcodes[pos+1])
	return pos + 2
}

func output(opcodes []int, pos int) int {
	fmt.Printf("Output: %d\n", getArgs(opcodes, pos, 1)[0])
	return pos + 2
}

func lessThan(opcodes []int, pos int) int {
	args := getArgs(opcodes, pos, 2)
	if args[0] < args[1] {
		opcodes[opcodes[pos+3]] = 1
	} else {
		opcodes[opcodes[pos+3]] = 0
	}
	return pos + 4
}

func equals(opcodes []int, pos int) int {
	args := getArgs(opcodes, pos, 2)
	if args[0] == args[1] {
		opcodes[opcodes[pos+3]] = 1
	} else {
		opcodes[opcodes[pos+3]] = 0
	}
	return pos + 4
}

func jumpIfTrue(opcodes []int, pos int) int {
	args := getArgs(opcodes, pos, 2)
	if args[0] != 0 {
		return args[1]
	}
	return pos + 3
}

func jumpIfFalse(opcodes []int, pos int) int {
	args := getArgs(opcodes, pos, 2)
	if args[0] == 0 {
		return args[1]
	}
	return pos + 3
}

func moveRelativeBase(opcodes []int, pos int) int {
	args := getArgs(opcodes, pos, 1)
	// fmt.Printf("Moving relative base by %d (was %d)\n", args[0], relativeBase)
	relativeBase += args[0]
	return pos + 2
}

var instructions = map[int](func([]int, int) int){
	1: add,
	2: multiply,
	3: save,
	4: output,
	5: jumpIfTrue,
	6: jumpIfFalse,
	7: lessThan,
	8: equals,
	9: moveRelativeBase,
}

func solve(opcodes []int) {
	pos := 0
	// fmt.Println(opcodes)
	for opcodes[pos] != 99 {
		// fmt.Printf("Evaluating intcode %d at position %d\n", opcodes[pos], pos)
		// fmt.Printf("Processing pos %d => %d\n", pos, opcodes[pos])
		code := opcodes[pos] % 100
		if fn, exists := instructions[code]; exists {
			pos = fn(opcodes, pos)
		} else {
			// check(fmt.Errorf("found unexpected opcode %d at position %d", opcodes[pos], pos))
		}
		// fmt.Println(opcodes)
		// fmt.Printf("Current relative base %d\n", relativeBase)
	}
	fmt.Println("Exit code found, exiting")
}

const input = 1

var relativeBase = 0

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	opcodesStr := strings.Split(strings.TrimRight(string(data), "\n"), ",")
	opcodes := make([]int, len(opcodesStr)*128)
	for k, s := range opcodesStr {
		i, err := strconv.Atoi(s)
		check(err)
		opcodes[k] = i
	}
	solve(opcodes)
}
