package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"adventofcode/utils/sets/stringset"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toKey(c coords2d.Coords2d) string {
	return fmt.Sprintf("%d:%d", c.X, c.Y)
}

var headVectors = map[string]coords2d.Coords2d{
	"U": {0, -1},
	"D": {0, 1},
	"L": {-1, 0},
	"R": {1, 0},
}

var tailVectors = map[string]coords2d.Coords2d{
	"2:0":  {1, 0},
	"-2:0": {-1, 0},
	"0:2":  {0, 1},
	"0:-2": {0, -1},

	"2:1": {1, 1}, "1:2": {1, 1}, "2:2": {1, 1},
	"-1:2": {-1, 1}, "-2:1": {-1, 1}, "-2:2": {-1, 1},
	"-2:-1": {-1, -1}, "-1:-2": {-1, -1}, "-2:-2": {-1, -1},
	"1:-2": {1, -1}, "2:-1": {1, -1}, "2:-2": {1, -1},
}

func atoi(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return res
}

func catchUp(head, tail coords2d.Coords2d) coords2d.Coords2d {
	diff := coords2d.Coords2d{head.X - tail.X, head.Y - tail.Y}
	return coords2d.Add(tail, tailVectors[toKey(diff)])
}

func solve(data []string) int {
	knots := make([]coords2d.Coords2d, 10)
	visitedTail := stringset.New()
	visitedTail.Add(toKey(knots[9]))

	for _, line := range data {
		parts := strings.Split(line, " ")
		direction := parts[0]
		count := atoi(parts[1])

		for i := 0; i < count; i++ {
			knots[0] = coords2d.Add(knots[0], headVectors[direction])
			for k := 1; k < len(knots); k++ {
				knots[k] = catchUp(knots[k-1], knots[k])
			}
			visitedTail.Add(toKey(knots[9]))
		}
	}
	return visitedTail.Len()
}

func main() {
	data := files.ReadLines(os.Args[1])
	fmt.Println(solve(data))
}
