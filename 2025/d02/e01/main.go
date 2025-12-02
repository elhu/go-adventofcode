package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(ranges [][2]int) int {
	var res int
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			a := strconv.Itoa(i)
			if len(a)%2 == 0 {
				mid := len(a) / 2
				left, right := a[:mid], a[mid:]
				if left == right {
					res += i
				}
			}
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	rawRanges := strings.Split(data, ",")
	ranges := make([][2]int, 0, len(rawRanges))
	for _, r := range rawRanges {
		bounds := strings.Split(r, "-")
		lower, err := strconv.Atoi(bounds[0])
		if err != nil {
			panic(err)
		}
		upper, err := strconv.Atoi(bounds[1])
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, [2]int{lower, upper})
	}
	fmt.Println(solve(ranges))
}
