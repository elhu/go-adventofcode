package main

import (
	"adventofcode/utils/files"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func decodeCombo(val int, regs map[byte]int) int {
	if val < 4 {
		return val
	} else if val == 4 {
		return regs['A']
	} else if val == 5 {
		return regs['B']
	} else if val == 6 {
		return regs['C']
	} else {
		panic("WTF")
	}
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

type State struct {
	regs   map[byte]int
	ip     int
	output []int
}

var instFns = []func(int, *State){
	func(v int, s *State) { // 0: adv
		s.regs['A'] = s.regs['A'] / powInt(2, decodeCombo(v, s.regs))
	},
	func(v int, s *State) { // 1: bxl
		s.regs['B'] = s.regs['B'] ^ v
	},
	func(v int, s *State) { // 2: bst
		s.regs['B'] = decodeCombo(v, s.regs) % 8
	},
	func(v int, s *State) { // 3: jnz
		if s.regs['A'] != 0 {
			s.ip = v - 2 // consider offset at end of loop
		}
	},
	func(v int, s *State) { // 4: bxc
		s.regs['B'] = s.regs['B'] ^ s.regs['C']
	},
	func(v int, s *State) { // 5: out
		s.output = append(s.output, decodeCombo(v, s.regs)%8)
	},
	func(v int, s *State) { // 6: bdv
		s.regs['B'] = s.regs['A'] / powInt(2, decodeCombo(v, s.regs))
	},
	func(v int, s *State) { // 7: cdv
		s.regs['C'] = s.regs['A'] / powInt(2, decodeCombo(v, s.regs))
	},
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func parseInsts(data string) []int {
	insts := []int{}
	data = data[len("Program: "):]
	ns := strings.Split(data, ",")
	for _, n := range ns {
		insts = append(insts, atoi(n))
	}
	return insts
}

func parseState(data []string) State {
	regs := map[byte]int{}
	regs['A'] = atoi(data[0][len("Register A: "):])
	regs['B'] = atoi(data[1][len("Register B: "):])
	regs['C'] = atoi(data[1][len("Register C: "):])
	return State{regs: regs, ip: 0, output: []int{}}
}

func eq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, n := range a {
		if n != b[i] {
			return false
		}
	}
	return true
}

func runMachine(s State, insts []int) []int {
	for s.ip < len(insts) {
		instFns[insts[s.ip]](insts[s.ip+1], &s)
		s.ip += 2
	}
	return s.output
}

func tailMatch(out []int, insts []int) bool {
	for i := 1; i <= len(out); i++ {
		if out[len(out)-i] != insts[len(insts)-i] {
			return false
		}
	}
	return true
}

// The machine processes A in 3 bits increments
// Started by printing i every time the tail of the output matches the tail of the instructions
// Each time there is a match, the next i that has one more output "right" is at least 8 times the previous one (3 bits increment)
func solve(s State, insts []int) int {
	for i := 0; ; i++ {
		s.regs = map[byte]int{'A': i, 'B': 0, 'C': 0}
		s.ip = 0
		out := runMachine(s, insts)

		if eq(out, insts) {
			return i
		}
		if tailMatch(out, insts) {
			i = i*8 - 1
		}
	}
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	parts := strings.Split(data, "\n\n")
	state := parseState(strings.Split(parts[0], "\n"))
	instructions := parseInsts(parts[1])
	fmt.Println(solve(state, instructions))
}
