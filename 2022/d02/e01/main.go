package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
)

func atoi(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return res
}

const loss = 0
const draw = 3
const win = 6

var values = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"X": 1,
	"Y": 2,
	"Z": 3,
}

const rock = 1
const paper = 2
const scissors = 3

func outcome(a, b int) int {
	if a == b {
		return draw
	}
	if a == rock {
		if b == scissors {
			return win
		}
		return loss
	}
	if a == paper {
		if b == rock {
			return win
		}
		return loss
	}
	if a == scissors {
		if b == paper {
			return win
		}
		return loss
	}
	panic("wtf")
}

func solve(rounds [][2]int) int {
	score := 0
	for _, round := range rounds {
		score += outcome(round[1], round[0])
		score += round[1]
	}
	return score
}

func main() {
	lines := files.ReadLines(os.Args[1])
	rounds := make([][2]int, len(lines))
	for i, l := range lines {
		var a, b string
		fmt.Sscanf(l, "%s %s", &a, &b)
		rounds[i] = [2]int{values[a], values[b]}
	}
	fmt.Println(solve(rounds))
}
