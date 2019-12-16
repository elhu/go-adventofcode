package main

import (
	"fmt"
	"math"
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

func solve(digits []int, phases int) []int {
	for step := 0; step < phases; step++ {
		for i := len(digits) - 2; i >= 0; i-- {
			digits[i] = abs(digits[i]+digits[i+1]) % 10
		}
	}
	return digits
}

func repeatMessage(s []int) []int {
	res := append(make([]int, 0), s...)
	for i := 1; i < 10000; i++ {
		res = append(res, s...)
	}
	return res
}

func digitsToInt(s []int) int {
	res := 0
	for i := range s {
		res += s[len(s)-i-1] * int(math.Pow10(i))
	}
	return res
}

func main() {
	input := repeatMessage(inputToDigits(os.Args[1]))
	offset := digitsToInt(input[0:7])
	fmt.Println(offset)
	fmt.Println(solve(input, 100)[offset : offset+8])
}
