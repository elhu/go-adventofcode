package main

import (
	"adventofcode/utils/files"
	set "adventofcode/utils/sets"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	check(err)
	return n
}

func parseLine(line string) (*set.Set[int], *set.Set[int]) {
	line = strings.Split(line, ": ")[1]
	parts := strings.Split(line, " | ")
	left, right := set.New[int](), set.New[int]()
	for _, num := range strings.Fields(parts[0]) {
		left.Add(atoi(num))
	}
	for _, num := range strings.Fields(parts[1]) {
		right.Add(atoi(num))
	}
	return left, right
}

func pow(n, p int) int {
	res := 1
	for i := 0; i < p; i++ {
		res *= n
	}
	return res
}

func solve(lines []string) int {
	var res int
	for _, line := range lines {
		left, right := parseLine(line)
		interLen := left.Intersection(right).Len()
		if interLen > 0 {
			res += pow(2, interLen-1)
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	fmt.Println(solve(lines))
}
