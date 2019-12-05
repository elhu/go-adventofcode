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
	for i := 0; i < numArgs; i++ {
		if len(argTypes) > i && argTypes[i] == 1 {
			args = append(args, opcodes[pos+i+1])
		} else {
			args = append(args, opcodes[opcodes[pos+i+1]])
		}
	}
	return args
}

func add(opcodes []int, pos int) int {
	args := getArgs(opcodes, pos, 2)
	opcodes[opcodes[pos+3]] = args[0] + args[1]
	return pos + 4
}

func multiply(opcodes []int, pos int) int {
	args := getArgs(opcodes, pos, 2)
	opcodes[opcodes[pos+3]] = args[0] * args[1]
	return pos + 4
}

func save(opcodes []int, pos int) int {
	opcodes[opcodes[pos+1]] = input
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

var instructions = map[int](func([]int, int) int){
	1: add,
	2: multiply,
	3: save,
	4: output,
	5: jumpIfTrue,
	6: jumpIfFalse,
	7: lessThan,
	8: equals,
}

func solve(opcodes []int) {
	pos := 0
	// fmt.Println(opcodes)
	for opcodes[pos] != 99 {
		// fmt.Printf("Processing pos %d => %d\n", pos, opcodes[pos])
		code := opcodes[pos] % 100
		if fn, exists := instructions[code]; exists {
			pos = fn(opcodes, pos)
		} else {
			check(fmt.Errorf("found unexpected opcode %d at position %d", opcodes[pos], pos))
		}
		// fmt.Println(opcodes)
	}
	fmt.Println("Exit code found, exiting")
}

const input = 5

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	opcodesStr := strings.Split(strings.TrimRight(string(data), "\n"), ",")
	opcodes := make([]int, 0, len(opcodesStr))
	for _, s := range opcodesStr {
		i, err := strconv.Atoi(s)
		check(err)
		opcodes = append(opcodes, i)
	}
	solve(opcodes)
}
