package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func atoi(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return val
}

func sum(calories []int) int {
	s := 0
	for _, c := range calories {
		s += c
	}
	return s
}

func solve(calories [][]int) int {
	max := 0
	for _, cs := range calories {
		s := sum(cs)
		if s > max {
			max = s
		}
	}
	return max
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	elves := strings.Split(data, "\n\n")
	calories := make([][]int, len(elves))
	for i, e := range elves {
		calories[i] = make([]int, 0)
		for _, c := range strings.Split(e, "\n") {
			calories[i] = append(calories[i], atoi(c))
		}
	}
	fmt.Println(solve(calories))
}
