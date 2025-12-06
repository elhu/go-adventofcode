package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
}

func solve(problems [][]int, operations []string) int {
	res := 0
	for i, op := range operations {
		acc := problems[i][0]
		for _, val := range problems[i][1:] {
			if op == "+" {
				acc += val
			} else if op == "*" {
				acc *= val
			}
		}
		res += acc
	}
	return res
}

func convertProblems(problems [][]string) [][]int {
	intProblems := make([][]int, len(problems))
	for j, p := range problems {
		for i := len(p[0]) - 1; i >= 0; i-- {
			acc := ""
			for _, val := range p {
				if val[i] != ' ' {
					acc += string(val[i])
				}
			}
			intProblems[j] = append(intProblems[j], atoi(acc))
		}
	}
	return intProblems
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	// Pad each line with a space to make parsing easier
	for i, line := range lines {
		lines[i] = line + "  "
	}
	problems := make([][]string, len(strings.Fields(lines[0])))
	for _, line := range lines[:len(lines)-1] {
		parts := strings.Fields(line)
		for i, p := range parts {
			problems[i] = append(problems[i], p)
		}
	}
	paddedProblems := make([][]string, len(problems))
	for i, p := range problems {
		maxLen := 0
		for _, val := range p {
			if maxLen < len(val) {
				maxLen = len(val)
			}
		}
		for j, line := range lines[:len(lines)-1] {
			s := line[0:maxLen]
			paddedProblems[i] = append(paddedProblems[i], s)
			lines[j] = line[maxLen+1:]
		}
	}
	operations := strings.Fields(lines[len(lines)-1])
	fmt.Println(solve(convertProblems(paddedProblems), operations))
}
