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

var values = map[string]int{
	"A": 1, // rock
	"B": 2, // paper
	"C": 3, // scissors
}

var outcomes = map[string]int{
	"X": 0, // lose
	"Y": 3, // draw
	"Z": 6, // win
}

const lose = "X"
const draw = "Y"
const win = "Z"

var plays = map[string]map[string]string{
	"A": {
		lose: "C", win: "B", draw: "A",
	},
	"B": {
		lose: "A", win: "C", draw: "B",
	},
	"C": {
		lose: "B", win: "A", draw: "C",
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
