package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"sort"
)

func min(numbers []int) int {
	res := numbers[0]
	for _, n := range numbers {
		if res > n {
			res = n
		}
	}
	return res
}

func findBasinSize(i, j int, input [][]int) int {
	res := 1
	input[i][j] = 9
	if i > 0 && input[i-1][j] < 9 {
		res += findBasinSize(i-1, j, input)
	}
	if i < len(input)-1 && input[i+1][j] < 9 {
		res += findBasinSize(i+1, j, input)
	}
	if j > 0 && input[i][j-1] < 9 {
		res += findBasinSize(i, j-1, input)
	}
	if j < len(input[i])-1 && input[i][j+1] < 9 {
		res += findBasinSize(i, j+1, input)
	}
	return res
}

func solve(input [][]int) int {
	basinSizes := make([]int, 0)
	for i := range input {
		for j := range input[i] {
			if input[i][j] < 9 {
				s := findBasinSize(i, j, input)
				basinSizes = append(basinSizes, s)
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func main() {
	data := files.ReadLines(os.Args[1])
	input := make([][]int, len(data))
	for i, l := range data {
		input[i] = make([]int, len(data[i]))
		for j, c := range l {
			input[i][j] = int(c - '0')
		}
	}
	fmt.Println(solve(input))
}
