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
	before, after   []int
	opCode, a, b, c int
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
	opCode, _ := strconv.Atoi(match[1])
	a, _ := strconv.Atoi(match[2])
	b, _ := strconv.Atoi(match[3])
	c, _ := strconv.Atoi(match[4])
	return opCode, a, b, c
}

func parseSample(before, after, inst string) sample {
	opCode, a, b, c := parseInst(inst)
	return sample{
		before: parseState(before),
		after:  parseState(after),
		opCode: opCode,
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

func perform(s sample, mapping map[int]map[string]struct{}) {
	original := make([]int, 4)
	copy(original, s.before)
	for opName, inst := range instructions {
		copy(s.before, original)
		inst(s.a, s.b, s.c, s.before)
		if eql(s.before, s.after) {
			mapping[s.opCode][opName] = struct{}{}
		}
	}
}

func reduceDone(mapping map[int]map[string]struct{}) bool {
	for _, v := range mapping {
		if len(v) > 1 {
			return false
		}
	}
	return true
}

func reduce(mapping map[int]map[string]struct{}) map[int]string {
	for !reduceDone(mapping) {
		for opCode, opNames := range mapping {
			if len(opNames) == 1 {
				var opToDelete string
				for opName, _ := range opNames {
					opToDelete = opName
				}
				for code, names := range mapping {
					if code != opCode {
						delete(names, opToDelete)
					}
				}
			}
		}
	}
	res := make(map[int]string)
	for opCode, opNames := range mapping {
		var opName string
		for k, _ := range opNames {
			opName = k
		}
		res[opCode] = opName
	}
	return res
}

func parseTestCase(lines [][]byte) {

}

type instruction struct {
	opCode, a, b, c int
}

func process(opCodeToNames map[int]string, insts []instruction) []int {
	registers := make([]int, 4)
	for _, inst := range insts {
		instructions[opCodeToNames[inst.opCode]](inst.a, inst.b, inst.c, registers)
	}
	return registers
}

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input = bytes.TrimSuffix(input, []byte{'\n'})
	lines := bytes.Split(input, []byte{'\n'})
	samples := parseSamples(lines)
	mapping := make(map[int]map[string]struct{})
	for _, s := range samples {
		if _, exists := mapping[s.opCode]; !exists {
			mapping[s.opCode] = make(map[string]struct{})
		}
		perform(s, mapping)
	}
	opCodeToName := reduce(mapping)
	input, err = ioutil.ReadFile(os.Args[2])
	check(err)
	input = bytes.TrimSuffix(input, []byte{'\n'})
	lines = bytes.Split(input, []byte{'\n'})
	insts := make([]instruction, 0)
	for _, line := range lines {
		opCode, a, b, c := parseInst(string(line))
		inst := instruction{
			a:      a,
			b:      b,
			c:      c,
			opCode: opCode,
		}
		insts = append(insts, inst)
	}
	registers := process(opCodeToName, insts)
	fmt.Println(registers[0])
}
