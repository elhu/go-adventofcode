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

func add(opcodes []int, pos int, _ *io) int {
	args := getArgs(opcodes, pos, 2)
	opcodes[opcodes[pos+3]] = args[0] + args[1]
	return pos + 4
}

func multiply(opcodes []int, pos int, _ *io) int {
	args := getArgs(opcodes, pos, 2)
	opcodes[opcodes[pos+3]] = args[0] * args[1]
	return pos + 4
}

func save(opcodes []int, pos int, io *io) int {
	opcodes[opcodes[pos+1]] = io.input[0]
	io.input = io.input[1:]
	return pos + 2
}

func output(opcodes []int, pos int, io *io) int {
	val := getArgs(opcodes, pos, 1)[0]
	io.output = val
	return pos + 2
}

func lessThan(opcodes []int, pos int, _ *io) int {
	args := getArgs(opcodes, pos, 2)
	if args[0] < args[1] {
		opcodes[opcodes[pos+3]] = 1
	} else {
		opcodes[opcodes[pos+3]] = 0
	}
	return pos + 4
}

func equals(opcodes []int, pos int, _ *io) int {
	args := getArgs(opcodes, pos, 2)
	if args[0] == args[1] {
		opcodes[opcodes[pos+3]] = 1
	} else {
		opcodes[opcodes[pos+3]] = 0
	}
	return pos + 4
}

func jumpIfTrue(opcodes []int, pos int, _ *io) int {
	args := getArgs(opcodes, pos, 2)
	if args[0] != 0 {
		return args[1]
	}
	return pos + 3
}

func jumpIfFalse(opcodes []int, pos int, _ *io) int {
	args := getArgs(opcodes, pos, 2)
	if args[0] == 0 {
		return args[1]
	}
	return pos + 3
}

var instructions = map[int](func([]int, int, *io) int){
	1: add,
	2: multiply,
	3: save,
	4: output,
	5: jumpIfTrue,
	6: jumpIfFalse,
	7: lessThan,
	8: equals,
}

func run(opcodes []int, io *io) {
	pos := 0
	// fmt.Println(opcodes)
	for opcodes[pos] != 99 {
		// fmt.Printf("Processing pos %d => %d\n", pos, opcodes[pos])
		code := opcodes[pos] % 100
		if fn, exists := instructions[code]; exists {
			pos = fn(opcodes, pos, io)
		} else {
			check(fmt.Errorf("found unexpected opcode %d at position %d", opcodes[pos], pos))
		}
		// fmt.Println(opcodes)
	}
	// fmt.Println("Exit code found, exiting")
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func max(arr []int) int {
	m := arr[0]
	for _, k := range arr {
		if k > m {
			m = k
		}
	}
	return m
}

type io struct {
	input  []int
	output int
}

func solve(opcodes []int) int {
	results := make([]int, 0)
	var inout *io
	for _, phase := range permutations([]int{0, 1, 2, 3, 4}) {
		input := 0
		for i := 0; i < numThrusters; i++ {
			thruster := append(opcodes[:], opcodes...)
			inout = &io{[]int{phase[i], input}, -1}
			run(thruster, inout)
			input = inout.output
		}
		results = append(results, inout.output)
	}
	return max(results)
}

const numThrusters = 5

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
	fmt.Println(solve(opcodes))
}
