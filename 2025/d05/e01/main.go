package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inRange(r [2]int, v int) bool {
	return v >= r[0] && v <= r[1]
}

func solve(ranges [][2]int, ingredients []int) int {
	res := 0
	for _, ing := range ingredients {
		valid := false
		for _, r := range ranges {
			if inRange(r, ing) {
				valid = true
				break
			}
		}
		if valid {
			res++
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	parts := strings.Split(data, "\n\n")
	var ranges [][2]int
	for _, p := range strings.Split(parts[0], "\n") {
		bounds := strings.Split(p, "-")
		var r [2]int
		fmt.Sscanf(bounds[0], "%d", &r[0])
		fmt.Sscanf(bounds[1], "%d", &r[1])
		ranges = append(ranges, r)
	}
	var ingredients []int
	for _, line := range strings.Split(parts[1], "\n") {
		ing, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		ingredients = append(ingredients, ing)
	}
	fmt.Println(solve(ranges, ingredients))
}
