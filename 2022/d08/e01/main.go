package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
)

func all(data []int, fn func(a int) bool) bool {
	for _, d := range data {
		if fn(d) == false {
			return false
		}
	}
	return true
}

func isVisible(x, y int, heights [][]int) bool {
	targetHeight := heights[y][x]
	data := make([]int, 0)
	for i := y - 1; i >= 0; i-- {
		data = append(data, heights[i][x])
	}
	if all(data, func(a int) bool { return a < targetHeight }) {
		return true
	}

	data = data[:0]
	for i := y + 1; i < len(heights); i++ {
		data = append(data, heights[i][x])
	}
	if all(data, func(a int) bool { return a < targetHeight }) {
		return true
	}

	data = data[:0]
	for i := x - 1; i >= 0; i-- {
		data = append(data, heights[y][i])
	}
	if all(data, func(a int) bool { return a < targetHeight }) {
		return true
	}

	data = data[:0]
	for i := x + 1; i < len(heights[0]); i++ {
		data = append(data, heights[y][i])
	}
	if all(data, func(a int) bool { return a < targetHeight }) {
		return true
	}

	return false
}

func solve(heights [][]int) int {
	res := 0
	for i := range heights {
		for j := range heights[i] {
			if isVisible(j, i, heights) {
				res++
			} else {
			}
		}
	}
	return res
}

func main() {
	data := files.ReadLines(os.Args[1])
	heights := make([][]int, len(data))
	for i, l := range data {
		heights[i] = make([]int, len(l))
		for j, c := range l {
			heights[i][j] = int(c - '0')
		}
	}
	fmt.Println(solve(heights))
}
