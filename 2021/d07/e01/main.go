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

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func computeCost(input []int, point int) int {
	res := 0
	for _, n := range input {
		res += abs(n - point)
	}
	return res
}

func minMax(input []int) (int, int) {
	min := input[1]
	max := 0
	for _, n := range input {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func solve(input []int) int {
	min, max := minMax(input)
	prevCost := 99999999999999999
	var res int
	for i := min; i <= max; i++ {
		cost := computeCost(input, i)
		if cost > prevCost {
			res = prevCost
			break
		}
		prevCost = cost
	}
	return res
}

func main() {
	data := files.ReadLinesWithSeparator(os.Args[1], ",")
	input := make([]int, len(data))
	for i, d := range data {
		input[i] = atoi(d)
	}
	fmt.Println(solve(input))
}
