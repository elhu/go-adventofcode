package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func lt(a, b int) bool {
	return a < b
}

func gt(a, b int) bool {
	return a > b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isSafe(report []int) bool {
	var cmp func(int, int) bool
	if report[0] > report[1] {
		cmp = gt
	} else if report[0] < report[1] {
		cmp = lt
	} else {
		return false
	}
	for i := 1; i < len(report); i++ {
		if !cmp(report[i-1], report[i]) || abs(report[i-1]-report[i]) < 1 || abs(report[i-1]-report[i]) > 3 {
			return false
		}
	}
	return true
}

func combinations(report []int) [][]int {
	res := [][]int{}
	for i := 0; i < len(report); i++ {
		var comb []int
		for j := 0; j < len(report); j++ {
			if i != j {
				comb = append(comb, report[j])
			}
		}
		res = append(res, comb)
	}
	return res
}

func anySafe(report []int) bool {
	for _, comb := range combinations(report) {
		if isSafe(comb) {
			return true
		}
	}
	return false
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	reports := strings.Split(data, "\n")
	res := 0
	for _, report := range reports {
		parts := strings.Split(report, " ")
		var report []int
		for _, part := range parts {
			r, _ := strconv.Atoi(part)
			report = append(report, r)
		}
		if anySafe(report) {
			res += 1
		}
	}
	fmt.Println(res)
}
