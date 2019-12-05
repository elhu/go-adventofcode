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

func add(opcodes []int, pos int) {
	var a, b int
	argTypes := codeToArgType(opcodes[pos])
	if len(argTypes) > 0 && argTypes[0] == 1 {
		a = opcodes[pos+1]
	} else {
		a = opcodes[opcodes[pos+1]]
	}
	if len(argTypes) > 1 && argTypes[1] == 1 {
		b = opcodes[pos+2]
	} else {
		b = opcodes[opcodes[pos+2]]
	}
	// fmt.Printf("Adding %d and %d and putting result at pos %d\n", a, b, opcodes[pos+3])
	opcodes[opcodes[pos+3]] = a + b
}

func multiply(opcodes []int, pos int) {
	var a, b int
	argTypes := codeToArgType(opcodes[pos])
	if len(argTypes) > 0 && argTypes[0] == 1 {
		a = opcodes[pos+1]
	} else {
		a = opcodes[opcodes[pos+1]]
	}
	if len(argTypes) > 1 && argTypes[1] == 1 {
		b = opcodes[pos+2]
	} else {
		b = opcodes[opcodes[pos+2]]
	}
	// fmt.Printf("Multiplying %d and %d and putting result at pos %d\n", a, b, opcodes[pos+3])
	opcodes[opcodes[pos+3]] = a * b
}

func save(opcodes []int, pos int) {
	// fmt.Printf("Saving %d at pos %d\n", input, opcodes[pos+1])
	opcodes[opcodes[pos+1]] = input
}

func output(opcodes []int, pos int) {
	var a int
	argTypes := codeToArgType(opcodes[pos])
	if len(argTypes) > 0 && argTypes[0] == 1 {
		a = opcodes[pos+1]
	} else {
		a = opcodes[opcodes[pos+1]]
	}
	fmt.Printf("Output: %d\n", a)
}

func lessThan(opcodes []int, pos int) {
	var a, b int
	argTypes := codeToArgType(opcodes[pos])
	if len(argTypes) > 0 && argTypes[0] == 1 {
		a = opcodes[pos+1]
	} else {
		a = opcodes[opcodes[pos+1]]
	}
	if len(argTypes) > 1 && argTypes[1] == 1 {
		b = opcodes[pos+2]
	} else {
		b = opcodes[opcodes[pos+2]]
	}
	// fmt.Printf("Checking if %d is less than %d\n", a, b)
	if a < b {
		opcodes[opcodes[pos+3]] = 1
	} else {
		opcodes[opcodes[pos+3]] = 0
	}
}

func equals(opcodes []int, pos int) {
	var a, b int
	argTypes := codeToArgType(opcodes[pos])
	if len(argTypes) > 0 && argTypes[0] == 1 {
		a = opcodes[pos+1]
	} else {
		a = opcodes[opcodes[pos+1]]
	}
	if len(argTypes) > 1 && argTypes[1] == 1 {
		b = opcodes[pos+2]
	} else {
		b = opcodes[opcodes[pos+2]]
	}
	// fmt.Printf("Checking if %d is equal to %d\n", a, b)
	if a == b {
		opcodes[opcodes[pos+3]] = 1
	} else {
		opcodes[opcodes[pos+3]] = 0
	}
}

func jumpIfTrue(opcodes []int, pos int) int {
	var a, b int
	argTypes := codeToArgType(opcodes[pos])
	if len(argTypes) > 0 && argTypes[0] == 1 {
		a = opcodes[pos+1]
	} else {
		a = opcodes[opcodes[pos+1]]
	}
	if len(argTypes) > 1 && argTypes[1] == 1 {
		b = opcodes[pos+2]
	} else {
		b = opcodes[opcodes[pos+2]]
	}
	if a != 0 {
		return b
	}
	return pos + 3
}

func jumpIfFalse(opcodes []int, pos int) int {
	var a, b int
	argTypes := codeToArgType(opcodes[pos])
	if len(argTypes) > 0 && argTypes[0] == 1 {
		a = opcodes[pos+1]
	} else {
		a = opcodes[opcodes[pos+1]]
	}
	if len(argTypes) > 1 && argTypes[1] == 1 {
		b = opcodes[pos+2]
	} else {
		b = opcodes[opcodes[pos+2]]
	}
	if a == 0 {
		return b
	}
	return pos + 3
}

func solve(opcodes []int) {
	pos := 0
	// fmt.Println(opcodes)
	for opcodes[pos] != 99 {
		// fmt.Printf("Processing pos %d => %d\n", pos, opcodes[pos])
		code := opcodes[pos] % 100
		if code == 1 {
			add(opcodes, pos)
			pos += 4
		} else if code == 2 {
			multiply(opcodes, pos)
			pos += 4
		} else if code == 3 {
			save(opcodes, pos)
			pos += 2
		} else if code == 4 {
			output(opcodes, pos)
			pos += 2
		} else if code == 5 {
			pos = jumpIfTrue(opcodes, pos)
		} else if code == 6 {
			pos = jumpIfFalse(opcodes, pos)
		} else if code == 7 {
			lessThan(opcodes, pos)
			pos += 4
		} else if code == 8 {
			equals(opcodes, pos)
			pos += 4
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
