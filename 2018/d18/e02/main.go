package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	open       = '.'
	tree       = '|'
	lumberyard = '#'
)

type TypeCounts struct {
	open, trees, lumberyards int
}

func surroundingCount(plan [][]byte, x, y int) TypeCounts {
	res := TypeCounts{}
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i >= 0 && i < len(plan) && j >= 0 && j < len(plan[i]) && !(x == j && i == y) {
				switch plan[i][j] {
				case open:
					res.open++
				case tree:
					res.trees++
				case lumberyard:
					res.lumberyards++
				}
			}
		}
	}
	return res
}

// Found by printing results as I went
const cycleSize = 28

func solve(plan [][]byte) int {
	scores := make(map[int]int)
	cpy := make([][]byte, len(plan))
	for i := 0; i < len(plan); i++ {
		cpy[i] = make([]byte, len(plan[i]))
	}
	for m := 0; m < 1000000000; m++ {
		for i := 0; i < len(plan); i++ {
			for j := 0; j < len(plan[i]); j++ {
				counts := surroundingCount(plan, j, i)
				switch plan[i][j] {
				case open:
					if counts.trees >= 3 {
						cpy[i][j] = tree
					} else {
						cpy[i][j] = open
					}
				case tree:
					if counts.lumberyards >= 3 {
						cpy[i][j] = lumberyard
					} else {
						cpy[i][j] = tree
					}
				case lumberyard:
					if counts.lumberyards >= 1 && counts.trees >= 1 {
						cpy[i][j] = lumberyard
					} else {
						cpy[i][j] = open
					}
				}
			}
		}
		plan, cpy = cpy, plan
		score := computeResult(plan)
		if minute, seen := scores[score]; seen {
			if (1000000000-(minute+1))%cycleSize == 0 {
				return score
			}
		}
		scores[computeResult(plan)] = m
	}
	return computeResult(plan)
}

func printPlan(plan [][]byte) {
	for i := 0; i < len(plan); i++ {
		fmt.Println(string(plan[i]))
	}
	fmt.Println("")
}

func computeResult(plan [][]byte) int {
	treeCount := 0
	lumberyardCount := 0
	for i := 0; i < len(plan); i++ {
		for j := 0; j < len(plan[i]); j++ {
			if plan[i][j] == tree {
				treeCount++
			} else if plan[i][j] == lumberyard {
				lumberyardCount++
			}
		}
	}
	return treeCount * lumberyardCount
}

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input = bytes.TrimSuffix(input, []byte{'\n'})
	plan := bytes.Split(input, []byte{'\n'})
	fmt.Println(solve(plan))
}
