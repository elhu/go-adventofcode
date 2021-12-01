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

func sumSlice(s []int) int {
	res := 0
	for _, n := range s {
		res += n
	}
	return res
}

func solve(depths []int) int {
	res := 0
	for i := 3; i < len(depths); i++ {
		if sumSlice(depths[i-2:i+1]) > sumSlice(depths[i-3:i]) {
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
