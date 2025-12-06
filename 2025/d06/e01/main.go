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

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	problems := make([][]int, len(strings.Fields(lines[0])))
	for _, line := range lines[:len(lines)-1] {
		parts := strings.Fields(line)
		for i, p := range parts {
			problems[i] = append(problems[i], atoi(p))
		}
	}
	operations := strings.Fields(lines[len(lines)-1])
	fmt.Println(solve(problems, operations))
}
