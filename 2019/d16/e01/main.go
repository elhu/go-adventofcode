package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func inputToDigits(s string) []int {
	res := make([]int, 0)
	for _, c := range s {
		res = append(res, int(c-'0'))
	}
	return res
}

func patternAt(rank int) []int {
	pattern := make([]int, 0, (rank+1)*4)
	for _, n := range []int{0, 1, 0, -1} {
		for i := 0; i < rank+1; i++ {
			pattern = append(pattern, n)
		}
	}
	return pattern
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func compute(digits []int, rank int) int {
	res := 0
	pattern := patternAt(rank)
	for i, n := range digits {
		res += (n * pattern[(i+1)%len(pattern)])
	}
	return abs(res) % 10
}

func solve(digits []int, phases int) []int {
	for step := 0; step < phases; step++ {
		newDigits := make([]int, len(digits))
		for i := range digits {
			newDigits[i] = compute(digits, i)
		}
		digits = append(make([]int, 0), newDigits...)
	}
	return digits
}

func main() {
	fmt.Println(solve(inputToDigits(os.Args[1]), 100)[0:8])
}
