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
	maxX, maxY, maxSize := 0, 0, 0
	for squareSize := 1; squareSize <= 300; squareSize++ {
		for i := 1; i < gridSize-squareSize; i++ {
			for j := 1; j < gridSize-squareSize; j++ {
				total := 0
				for k := 0; k < squareSize; k++ {
					total += sum(grid[i+k][j : j+squareSize])
				}
				if total > max {
					maxX = j
					maxY = i
					maxSize = squareSize
					max = total
				}
			}
		}
	}
	fmt.Printf("%d,%d,%d\n", maxX, maxY, maxSize)
}
