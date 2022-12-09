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

	"2:1": {1, 1}, "1:2": {1, 1},
	"-1:2": {-1, 1}, "-2:1": {-1, 1},
	"-2:-1": {-1, -1}, "-1:-2": {-1, -1},
	"1:-2": {1, -1}, "2:-1": {1, -1},
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
	var head, tail coords2d.Coords2d
	visitedTail := stringset.New()
	visitedTail.Add(toKey(tail))

	for _, line := range data {
		parts := strings.Split(line, " ")
		direction := parts[0]
		count := atoi(parts[1])

		for i := 0; i < count; i++ {
			head = coords2d.Add(head, headVectors[direction])
			tail = catchUp(head, tail)
			// fmt.Printf("Adding %s to visited set -- head is at %s\n", toKey(tail), toKey(head))
			visitedTail.Add(toKey(tail))
		}
	}
	return visitedTail.Len()
}

func main() {
	data := files.ReadLines(os.Args[1])
	fmt.Println(solve(data))
}
