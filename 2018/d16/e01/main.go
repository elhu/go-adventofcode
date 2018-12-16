package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var instructions = map[string]func(inA, inB, outC int, registers []int){
	"addr": func(inA, inB, outC int, registers []int) { registers[outC] = registers[inA] + registers[inB] },
	"addi": func(inA, inB, outC int, registers []int) { registers[outC] = registers[inA] + inB },
	"mulr": func(inA, inB, outC int, registers []int) { registers[outC] = registers[inA] * registers[inB] },
	"muli": func(inA, inB, outC int, registers []int) { registers[outC] = registers[inA] * inB },
	"banr": func(inA, inB, outC int, registers []int) { registers[outC] = registers[inA] & registers[inB] },
	"bani": func(inA, inB, outC int, registers []int) { registers[outC] = registers[inA] & inB },
	"borr": func(inA, inB, outC int, registers []int) { registers[outC] = registers[inA] | registers[inB] },
	"bori": func(inA, inB, outC int, registers []int) { registers[outC] = registers[inA] | inB },
	"setr": func(inA, inB, outC int, registers []int) { registers[outC] = registers[inA] },
	"seti": func(inA, inB, outC int, registers []int) { registers[outC] = inA },
	"gtir": func(inA, inB, outC int, registers []int) {
		if inA > registers[inB] {
			registers[outC] = 1
		} else {
			registers[outC] = 0
		}
	},
	"gtri": func(inA, inB, outC int, registers []int) {
		if registers[inA] > inB {
			registers[outC] = 1
		} else {
			registers[outC] = 0
		}
	},
	"gtrr": func(inA, inB, outC int, registers []int) {
		if registers[inA] > registers[inB] {
			registers[outC] = 1
		} else {
			registers[outC] = 0
		}
	},
	"eqir": func(inA, inB, outC int, registers []int) {
		if inA == registers[inB] {
			registers[outC] = 1
		} else {
			registers[outC] = 0
		}
	},
	"eqri": func(inA, inB, outC int, registers []int) {
		if registers[inA] == inB {
			registers[outC] = 1
		} else {
			registers[outC] = 0
		}
	},
	"eqrr": func(inA, inB, outC int, registers []int) {
		if registers[inA] == registers[inB] {
			registers[outC] = 1
		} else {
			registers[outC] = 0
		}
	},
}

type sample struct {
	before, after []int
	a, b, c       int
}

var stateExp = regexp.MustCompile(`\[(\d+), (\d+), (\d+), (\d+)\]`)
var instExp = regexp.MustCompile(`(\d+) (\d+) (\d+) (\d+)`)

func parseState(str string) []int {
	match := stateExp.FindStringSubmatch(str)
	state := make([]int, 0, 4)
	for _, m := range match[1:] {
		register, _ := strconv.Atoi(m)
		state = append(state, register)
	}
	return state
}

func parseInst(str string) (int, int, int, int) {
	match := instExp.FindStringSubmatch(str)
	inst, _ := strconv.Atoi(match[1])
	a, _ := strconv.Atoi(match[2])
	b, _ := strconv.Atoi(match[3])
	c, _ := strconv.Atoi(match[4])
	return inst, a, b, c
}

func parseSample(before, after, inst string) sample {
	_, a, b, c := parseInst(inst)
	return sample{
		before: parseState(before),
		after:  parseState(after),
		a:      a,
		b:      b,
		c:      c,
	}
}

func parseSamples(lines [][]byte) []sample {
	samples := make([]sample, 0)
	for i := 0; i < len(lines); i++ {
		line := string(lines[i])
		if strings.HasPrefix(line, "Before: ") {
			samples = append(samples, parseSample(line, string(lines[i+2]), string(lines[i+1])))
			i += 2
		}
	}
	return samples
}

func eql(a, b []int) bool {
	if len(a) == len(b) {
		for i, j := range a {
			if j != b[i] {
				return false
			}
		}
		return true
	}
	return false
}

func perform(s sample) int {
	matchCount := 0
	original := make([]int, 4)
	copy(original, s.before)
	for _, inst := range instructions {
		copy(s.before, original)
		inst(s.a, s.b, s.c, s.before)
		if eql(s.before, s.after) {
			matchCount++
		}
	}
	return matchCount
}

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input = bytes.TrimSuffix(input, []byte{'\n'})
	lines := bytes.Split(input, []byte{'\n'})
	samples := parseSamples(lines)
	res := 0
	for _, s := range samples {
		if perform(s) >= 3 {
			res++
		}
	}
	fmt.Println(res)
}
