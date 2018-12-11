package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sum(numbers []int) int {
	res := 0
	for _, i := range numbers {
		res += i
	}
	return res
}

const gridSize = 301
const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

func nthDigit(number, n int) int {
	if number/int(math.Pow10(n)) == 0 {
		return 0
	}
	number /= int(math.Pow10(n))
	return number % 10
}

func computeScore(x, y, serial int) int {
	rackID := x + 10
	res := rackID*y + serial
	res *= rackID
	res = nthDigit(res, 2)
	return res - 5
}

func main() {
	serial, _ := strconv.Atoi(os.Args[1])
	grid := make([][]int, gridSize)
	for i := 1; i < len(grid); i++ {
		grid[i] = make([]int, gridSize)
		for j := 1; j < len(grid); j++ {
			grid[i][j] = computeScore(j, i, serial)
		}
	}
	max := minInt
	maxX, maxY := 0, 0
	for i := 1; i < gridSize-3; i++ {
		for j := 1; j < gridSize-3; j++ {
			total := sum(grid[i][j:j+3]) + sum(grid[i+1][j:j+3]) + sum(grid[i+2][j:j+3])
			if total > max {
				maxX = j
				maxY = i
				max = total
			}
		}
	}
	fmt.Printf("%d,%d\n", maxX, maxY)
}
