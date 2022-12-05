package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

func parseCrates(lines []string) map[byte][]byte {
	crates := make(map[byte][]byte)
	for i, c := range lines[len(lines)-1] {
		c := byte(c)
		if c != ' ' {
			for j := 0; j < len(lines)-1; j++ {
				if lines[j][i] != ' ' {
					crates[c] = append(crates[c], lines[j][i])
				}
			}
		}
	}
	return crates
}

type Instruction struct {
	From, To byte
	Quantity int
}

func parseInstructions(lines []string) []Instruction {
	instructions := make([]Instruction, len(lines))
	for i, l := range lines {
		inst := Instruction{}
		fmt.Sscanf(l, "move %d from %c to %c", &inst.Quantity, &inst.From, &inst.To)
		instructions[i] = inst
	}
	return instructions
}

func copySlice(slice []byte) []byte {
	newSlice := make([]byte, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func solve(crates map[byte][]byte, instructions []Instruction) string {
	for _, inst := range instructions {
		crates[inst.To] = append(copySlice(crates[inst.From][0:inst.Quantity]), crates[inst.To]...)
		crates[inst.From] = crates[inst.From][inst.Quantity:]
	}
	res := make([]byte, len(crates))
	for i := 0; i < len(crates); i++ {
		if stack := crates[byte(i+1+'0')]; len(stack) > 0 {
			res[i] = stack[0]
		}
	}
	return string(res)
}

func main() {
	data := string(files.ReadFile(os.Args[1]))
	parts := strings.Split(data, "\n\n")
	crates := parseCrates(strings.Split(parts[0], "\n"))
	instructions := parseInstructions(strings.Split(parts[1], "\n"))
	fmt.Println(solve(crates, instructions))
}
