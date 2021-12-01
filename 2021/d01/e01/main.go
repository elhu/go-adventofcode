package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
)

func atoi(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return res
}

func solve(depths []int) int {
	res := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			res++
		}
	}
	return res
}

func main() {
	lines := files.ReadLines(os.Args[1])
	depths := make([]int, len(lines))
	for i, l := range lines {
		depths[i] = atoi(l)
	}
	fmt.Println(solve(depths))
}
