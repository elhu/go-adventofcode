package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
)

func neighbors(i, j int) [][2]int {
	var res [][2]int
	for _, p := range [][2]int{
		{i - 1, j - 1}, {i, j - 1}, {i + 1, j - 1}, {i - 1, j},
		// {i, j},
		{i + 1, j}, {i - 1, j + 1}, {i, j + 1}, {i + 1, j + 1},
	} {
		if p[0] >= 0 && p[0] < 10 && p[1] >= 0 && p[1] < 10 {
			res = append(res, p)
		}
	}
	return res
}

func handleFlash(i, j int, input *[10][10]int, hasFlashed *[10][10]bool) int {
	res := 1
	for _, c := range neighbors(i, j) {
		a, b := c[0], c[1]
		input[a][b]++
		if input[a][b] > 9 && !hasFlashed[a][b] {
			hasFlashed[a][b] = true
			res += handleFlash(a, b, input, hasFlashed)
		}
	}
	return res
}

func solve(input [10][10]int) int {
	res := 0
	for r := 0; r < rounds; r++ {
		var hasFlashed [10][10]bool
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				input[i][j]++
				if input[i][j] > 9 && !hasFlashed[i][j] {
					hasFlashed[i][j] = true
					res += handleFlash(i, j, &input, &hasFlashed)
				}
			}
		}
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if hasFlashed[i][j] {
					input[i][j] = 0
				}
			}
		}
	}
	return res
}

const rounds = 100

func main() {
	data := files.ReadLines(os.Args[1])
	var input [10][10]int
	for i, l := range data {
		for j, c := range l {
			input[i][j] = int(c - '0')
		}
	}
	fmt.Println(solve(input))
}
