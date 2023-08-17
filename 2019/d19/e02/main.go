package main

import (
	"fmt"
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

func getArgs(opcodes map[int]int, pos, numArgs int) []int {
	argTypes := codeToArgType(opcodes[pos])
	args := make([]int, 0, numArgs)
	for i := 0; i < numArgs; i++ {
		if len(argTypes) > i && argTypes[i] == 2 {
			args = append(args, opcodes[relativeBase+opcodes[pos+i+1]])
		} else if len(argTypes) > i && argTypes[i] == 1 {
			args = append(args, opcodes[pos+i+1])
		} else {
			args = append(args, opcodes[opcodes[pos+i+1]])
		}
	}
	return args
}

func getReturnPos(opcodes map[int]int, pos, returnOffset int) int {
	argTypes := codeToArgType(opcodes[pos])
	if len(argTypes) >= returnOffset && argTypes[returnOffset-1] == 2 {
		return relativeBase + opcodes[pos+returnOffset]
	} else if len(argTypes) >= returnOffset && argTypes[returnOffset-1] == 1 {
		return pos + returnOffset
	}
	return opcodes[pos+returnOffset]
}

func add(opcodes map[int]int, pos int, _ IO) int {
	args := getArgs(opcodes, pos, 2)
	returnPos := getReturnPos(opcodes, pos, 3)
	opcodes[returnPos] = args[0] + args[1]
	return pos + 4
}

func multiply(opcodes map[int]int, pos int, _ IO) int {
	args := getArgs(opcodes, pos, 2)
	returnPos := getReturnPos(opcodes, pos, 3)
	opcodes[returnPos] = args[0] * args[1]
	return pos + 4
}

func save(opcodes map[int]int, pos int, io IO) int {
	returnPos := getReturnPos(opcodes, pos, 1)
	opcodes[returnPos] = <-io.input
	return pos + 2
}

func output(opcodes map[int]int, pos int, io IO) int {
	val := getArgs(opcodes, pos, 1)[0]
	io.output <- val
	return pos + 2
}

func lessThan(opcodes map[int]int, pos int, _ IO) int {
	args := getArgs(opcodes, pos, 2)
	returnPos := getReturnPos(opcodes, pos, 3)
	if args[0] < args[1] {
		opcodes[returnPos] = 1
	} else {
		opcodes[returnPos] = 0
	}
	return pos + 4
}

func equals(opcodes map[int]int, pos int, _ IO) int {
	args := getArgs(opcodes, pos, 2)
	returnPos := getReturnPos(opcodes, pos, 3)
	if args[0] == args[1] {
		opcodes[returnPos] = 1
	} else {
		opcodes[returnPos] = 0
	}
	return pos + 4
}

func jumpIfTrue(opcodes map[int]int, pos int, _ IO) int {
	args := getArgs(opcodes, pos, 2)
	if args[0] != 0 {
		return args[1]
	}
	return pos + 3
}

func jumpIfFalse(opcodes map[int]int, pos int, _ IO) int {
	args := getArgs(opcodes, pos, 2)
	if args[0] == 0 {
		return args[1]
	}
	return pos + 3
}

func moveRelativeBase(opcodes map[int]int, pos int, _ IO) int {
	args := getArgs(opcodes, pos, 1)
	relativeBase += args[0]
	return pos + 2
}

var instructions = map[int](func(map[int]int, int, IO) int){
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

func run(opcodes map[int]int, io IO) {
	pos := 0
	for opcodes[pos] != 99 {
		code := opcodes[pos] % 100
		if fn, exists := instructions[code]; exists {
			pos = fn(opcodes, pos, io)
		} else {
			check(fmt.Errorf("found unexpected opcode %d at position %d", opcodes[pos], pos))
		}
	}
	close(io.output)
}

type IO struct {
	input  chan int
	output chan int
}

func processBeacon(opcodes map[int]int, x, y int) bool {
	io := IO{input: make(chan int, 2), output: make(chan int, 1)}
	go run(opcodes, io)
	io.input <- x
	io.input <- y
	res := <-io.output
	close(io.input)
	return res == 1
}

func mapCopy(b, a map[int]int) {
	for k, v := range a {
		b[k] = v
	}
}

func solve(opcodes map[int]int) int {
	var cells [10000][10000]bool
	for y := 0; y < 5000; y++ {
		for x := 0; x < 5000; x++ {
			newVM := make(map[int]int)
			mapCopy(newVM, opcodes)
			cells[y][x] = processBeacon(newVM, x, y)
		}
	}
	for i := 0; i < len(cells)-100; i++ {
		for j := 0; j < len(cells[i])-100; j++ {
			if cells[i][j] && cells[i+100][j] && cells[i][j+100] && cells[i+100][j+100] {
				return j*10000 + i
			}
		}
	}
	return 0
}

var relativeBase = 0

func main() {
	data, err := os.ReadFile(os.Args[1])
	check(err)
	opcodesStr := strings.Split(strings.TrimRight(string(data), "\n"), ",")
	opcodes := make(map[int]int)
	for k, s := range opcodesStr {
		i, err := strconv.Atoi(s)
		check(err)
		opcodes[k] = i
	}
	fmt.Println(solve(opcodes))
}
