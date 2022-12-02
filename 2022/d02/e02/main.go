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

const (
	lose = "X"
	draw = "Y"
	win  = "Z"

	rock     = "A"
	paper    = "B"
	scissors = "C"
)

var values = map[string]int{
	rock:     1,
	paper:    2,
	scissors: 3,
}

var outcomes = map[string]int{
	lose: 0,
	draw: 3,
	win:  6,
}

var plays = map[string]map[string]string{
	rock: {
		lose: scissors, win: paper, draw: rock,
	},
	paper: {
		lose: rock, win: scissors, draw: paper,
	},
	scissors: {
		lose: paper, win: rock, draw: scissors,
	},
}

func solve(rounds [][2]string) int {
	score := 0
	for _, round := range rounds {
		play := plays[round[0]][round[1]]
		score += values[play]
		score += outcomes[round[1]]
	}
	return score
}

func main() {
	lines := files.ReadLines(os.Args[1])
	rounds := make([][2]string, len(lines))
	for i, l := range lines {
		var a, b string
		fmt.Sscanf(l, "%s %s", &a, &b)
		rounds[i] = [2]string{a, b}
	}
	fmt.Println(solve(rounds))
}
