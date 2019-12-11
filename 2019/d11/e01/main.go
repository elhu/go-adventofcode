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

func getReturnPos(opcodes []int, pos, returnOffset int) int {
	argTypes := codeToArgType(opcodes[pos])
	if len(argTypes) >= returnOffset && argTypes[returnOffset-1] == 2 {
		return relativeBase + opcodes[pos+returnOffset]
	} else if len(argTypes) >= returnOffset && argTypes[returnOffset-1] == 1 {
		return pos + returnOffset
	}
	return opcodes[pos+returnOffset]
}

func add(opcodes []int, pos int, _ IO) int {
	args := getArgs(opcodes, pos, 2)
	returnPos := getReturnPos(opcodes, pos, 3)
	opcodes[returnPos] = args[0] + args[1]
	return pos + 4
}

func multiply(opcodes []int, pos int, _ IO) int {
	args := getArgs(opcodes, pos, 2)
	returnPos := getReturnPos(opcodes, pos, 3)
	opcodes[returnPos] = args[0] * args[1]
	return pos + 4
}

func save(opcodes []int, pos int, io IO) int {
	returnPos := getReturnPos(opcodes, pos, 1)
	// fmt.Printf("[OPCD] Waiting for input\n")
	opcodes[returnPos] = <-io.input
	// fmt.Printf("[OPCD] Received input %d\n", opcodes[returnPos])
	return pos + 2
}

func output(opcodes []int, pos int, io IO) int {
	val := getArgs(opcodes, pos, 1)[0]
	// fmt.Printf("[OPCD] Sending output %d...\n", val)
	io.output <- val
	// fmt.Printf("[OPCD] Output sent\n")
	// fmt.Printf("Output: %d\n", getArgs(opcodes, pos, 1)[0])
	return pos + 2
}

func lessThan(opcodes []int, pos int, _ IO) int {
	args := getArgs(opcodes, pos, 2)
	returnPos := getReturnPos(opcodes, pos, 3)
	if args[0] < args[1] {
		opcodes[returnPos] = 1
	} else {
		opcodes[returnPos] = 0
	}
	return pos + 4
}

func equals(opcodes []int, pos int, _ IO) int {
	args := getArgs(opcodes, pos, 2)
	returnPos := getReturnPos(opcodes, pos, 3)
	if args[0] == args[1] {
		opcodes[returnPos] = 1
	} else {
		opcodes[returnPos] = 0
	}
	return pos + 4
}

func jumpIfTrue(opcodes []int, pos int, _ IO) int {
	args := getArgs(opcodes, pos, 2)
	if args[0] != 0 {
		return args[1]
	}
	return pos + 3
}

func jumpIfFalse(opcodes []int, pos int, _ IO) int {
	args := getArgs(opcodes, pos, 2)
	if args[0] == 0 {
		return args[1]
	}
	return pos + 3
}

func moveRelativeBase(opcodes []int, pos int, _ IO) int {
	args := getArgs(opcodes, pos, 1)
	relativeBase += args[0]
	return pos + 2
}

var instructions = map[int](func([]int, int, IO) int){
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

func run(opcodes []int, io IO) {
	pos := 0
	for opcodes[pos] != 99 {
		code := opcodes[pos] % 100
		if fn, exists := instructions[code]; exists {
			pos = fn(opcodes, pos, io)
		} else {
			check(fmt.Errorf("found unexpected opcode %d at position %d", opcodes[pos], pos))
		}
	}
	fmt.Println("Exit code found, exiting")
	close(io.output)
}

func solve(opcodes []int) int {
	io := IO{make(chan int, 1), make(chan int, 2)}
	pos := coords{0, 0}
	panels := make(map[string]int)
	dir := 0
	io.input <- 0
	go run(opcodes, io)
	i := 0
	for o := range io.output {
		// fmt.Printf("[MAIN] Received output %d\n", o)
		if i%2 == 0 {
			// Handle paint
			// fmt.Printf("Painting %d\n", pos)
			panels[coordsToKey(pos)] = o
		} else {
			// Handle rotate
			if o == 0 {
				dir--
				if dir == -1 {
					dir = 3
				}
			} else {
				dir = (dir + 1) % 4
			}
			pos = addVector(pos, directions[dir])
			// fmt.Printf("[MAIN] Sending input %d...\n", panels[coordsToKey(pos)])
			// fmt.Printf("[MAIN] Input sent %d...\n", panels[coordsToKey(pos)])
			io.input <- panels[coordsToKey(pos)]
		}
		i++
		// fmt.Printf("[MAIN] Waiting for output...\n")
	}
	return len(panels)
}

type IO struct {
	input  chan int
	output chan int
}

type coords struct {
	x, y int
}

func coordsToKey(c coords) string {
	return fmt.Sprintf("%d:%d", c.x, c.y)
}

func addVector(c, v coords) coords {
	return coords{c.x + v.x, c.y + v.y}
}

var directions = map[int]coords{
	0: {0, -1},
	1: {1, 0},
	2: {0, 1},
	3: {-1, 0},
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
	fmt.Println(solve(opcodes))
}
