package main

import "fmt"

const min = 134564
const max = 585159

func matchesRules(digits []int) bool {
	prev := digits[0]
	foundAdjacent := false
	for _, i := range digits[1:] {
		if i < prev {
			return false
		}
		if i == prev {
			foundAdjacent = true
		}
		prev = i
	}
	return foundAdjacent
}

func numberToDigits(num int) []int {
	res := make([]int, 0)
	for num != 0 {
		res = append(res, num%10)
		num /= 10
	}
	for i := 0; i < len(res)/2; i++ {
		j := len(res) - 1 - i
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func solve(min, max int) int {
	res := 0
	for i := min; i <= max; i++ {
		if matchesRules(numberToDigits(i)) {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(solve(min, max))
}
