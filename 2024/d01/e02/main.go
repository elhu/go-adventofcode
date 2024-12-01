package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	left, right := make(map[int]int), make(map[int]int)
	for _, line := range lines {
		parts := strings.Split(line, "   ")
		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])
		if _, ok := left[l]; !ok {
			left[l] = 0
		}
		left[l]++
		if _, ok := right[r]; !ok {
			right[r] = 0
		}
		right[r]++
	}
	res := 0
	for k, v := range left {
		res += v * k * right[k]
	}
	fmt.Println(res)
}
