package main

import "os"
import "fmt"

import "strconv"
import "text/tabwriter"
import "strings"

const (
	right  = iota
	top    = iota
	left   = iota
	bottom = iota
)

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func showGrid(grid [][]int) {
	w := tabwriter.NewWriter(os.Stdout, 4, 0, 1, ' ', tabwriter.AlignRight)
	for i := 0; i < len(grid); i++ {
		// fmt.Println([]byte(arrayToString(grid[i], "\t") + "\t\n"))
		w.Write([]byte(arrayToString(grid[i], "\t") + "\t\n"))
	}
	w.Flush()
}

func initGrid(size int) [][]int {
	grid := make([][]int, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]int, size)
	}
	grid[size/2][size/2] = 1

	return grid
}

func sumSurroundings(grid [][]int, x int, y int) int {
	sum := 0
	for j := y - 1; j <= y+1; j++ {
		for i := x - 1; i <= x+1; i++ {
			if i != x || j != y {
				sum += grid[j][i]
			}
		}
	}
	return sum
}

func solve(target int) int {
	gridSize := 20
	grid := initGrid(gridSize)
	x := gridSize/2 + 1
	y := gridSize / 2
	current := 0
	ring := 1
	side := right
	posInSide := 0

	for current < target {
		current = sumSurroundings(grid, x, y)
		grid[y][x] = current
		posInSide++
		// Reached end of side, switch side
		if posInSide == ring*2 {
			// Reached the end, switch ring
			if side == bottom {
				x++
				y++ // compensate for the new side
				ring++
			}
			posInSide = 0
			side = (side + 1) % 4
		}
		switch side {
		case right:
			y--
		case top:
			x--
		case left:
			y++
		case bottom:
			x++
		}
	}
	return current
}

func main() {
	target, _ := strconv.Atoi(os.Args[1])

	fmt.Println(solve(target))
}
