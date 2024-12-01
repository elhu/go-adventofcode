package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"sort"
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
	var left, right []int
	for _, line := range lines {
		parts := strings.Split(line, "   ")
		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])
		left = append(left, l)
		right = append(right, r)
	}
	sort.Ints(left)
	sort.Ints(right)
	res := 0
	for i := 0; i < len(left); i++ {
		res += abs(left[i] - right[i])
	}
	fmt.Println(res)
}
