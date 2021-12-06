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

func solve(input []int) int {
	for turn := 0; turn < 80; turn++ {
		for i, n := range input {
			if n == 0 {
				input = append(input, 8)
				input[i] = 6
			} else {
				input[i]--
			}
		}
	}
	return len(input)
}

func main() {
	data := files.ReadLinesWithSeparator(os.Args[1], ",")
	input := make([]int, len(data))
	for i, d := range data {
		input[i] = atoi(d)
	}
	fmt.Println(solve(input))
}
