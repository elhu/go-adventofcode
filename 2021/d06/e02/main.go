package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
)

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func solve(input map[int]int) int {
	for turn := 0; turn < turns; turn++ {
		newInput := make(map[int]int)
		for k, v := range input {
			if k == 0 {
				newInput[8] += v
				newInput[6] += v
			} else {
				newInput[k-1] += v
			}
		}
		input = newInput
	}
	res := 0
	for _, v := range input {
		res += v
	}
	return res
}

const turns = 256

func main() {
	data := files.ReadLinesWithSeparator(os.Args[1], ",")
	input := make(map[int]int)
	for _, d := range data {
		input[atoi(d)]++
	}
	fmt.Println(solve(input))
}
