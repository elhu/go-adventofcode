package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
)

type Range struct {
	min, max int
}

func (p *Range) contains(o *Range) bool {
	return p.min <= o.min && p.max >= o.max
}

func solve(pairs [][2]*Range) int {
	res := 0
	for _, p := range pairs {
		if p[0].contains(p[1]) || p[1].contains(p[0]) {
			res++
		}
	}
	return res
}

func main() {
	lines := files.ReadLines(os.Args[1])
	pairs := make([][2]*Range, len(lines))
	for i, l := range lines {
		left, right := &Range{}, &Range{}
		fmt.Sscanf(l, "%d-%d,%d-%d", &left.min, &left.max, &right.min, &right.max)
		pairs[i] = [2]*Range{left, right}
	}

	fmt.Println(solve(pairs))
}
